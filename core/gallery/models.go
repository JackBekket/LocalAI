package gallery

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"dario.cat/mergo"
	"github.com/mudler/LocalAI/core/config"
	lconfig "github.com/mudler/LocalAI/core/config"
	"github.com/mudler/LocalAI/pkg/downloader"
	"github.com/mudler/LocalAI/pkg/utils"

	"github.com/rs/zerolog/log"
	"gopkg.in/yaml.v2"
)

/*

description: |
    foo
license: ""

urls:
-
-

name: "bar"

config_file: |
    # Note, name will be injected. or generated by the alias wanted by the user
    threads: 14

files:
    - filename: ""
      sha: ""
      uri: ""

prompt_templates:
    - name: ""
      content: ""

*/
// ModelConfig is the model configuration which contains all the model details
// This configuration is read from the gallery endpoint and is used to download and install the model
// It is the internal structure, separated from the request
type ModelConfig struct {
	Description     string           `yaml:"description"`
	Icon            string           `yaml:"icon"`
	License         string           `yaml:"license"`
	URLs            []string         `yaml:"urls"`
	Name            string           `yaml:"name"`
	ConfigFile      string           `yaml:"config_file"`
	Files           []File           `yaml:"files"`
	PromptTemplates []PromptTemplate `yaml:"prompt_templates"`
}

type File struct {
	Filename string `yaml:"filename" json:"filename"`
	SHA256   string `yaml:"sha256" json:"sha256"`
	URI      string `yaml:"uri" json:"uri"`
}

type PromptTemplate struct {
	Name    string `yaml:"name"`
	Content string `yaml:"content"`
}

// Installs a model from the gallery
func InstallModelFromGallery(galleries []config.Gallery, name string, basePath string, req GalleryModel, downloadStatus func(string, string, string, float64), enforceScan bool) error {

	applyModel := func(model *GalleryModel) error {
		name = strings.ReplaceAll(name, string(os.PathSeparator), "__")

		var config ModelConfig

		if len(model.URL) > 0 {
			var err error
			config, err = GetGalleryConfigFromURL[ModelConfig](model.URL, basePath)
			if err != nil {
				return err
			}
			config.Description = model.Description
			config.License = model.License
		} else if len(model.ConfigFile) > 0 {
			// TODO: is this worse than using the override method with a blank cfg yaml?
			reYamlConfig, err := yaml.Marshal(model.ConfigFile)
			if err != nil {
				return err
			}
			config = ModelConfig{
				ConfigFile:  string(reYamlConfig),
				Description: model.Description,
				License:     model.License,
				URLs:        model.URLs,
				Name:        model.Name,
				Files:       make([]File, 0), // Real values get added below, must be blank
				// Prompt Template Skipped for now - I expect in this mode that they will be delivered as files.
			}
		} else {
			return fmt.Errorf("invalid gallery model %+v", model)
		}

		installName := model.Name
		if req.Name != "" {
			installName = req.Name
		}

		// Copy the model configuration from the request schema
		config.URLs = append(config.URLs, model.URLs...)
		config.Icon = model.Icon
		config.Files = append(config.Files, req.AdditionalFiles...)
		config.Files = append(config.Files, model.AdditionalFiles...)

		// TODO model.Overrides could be merged with user overrides (not defined yet)
		if err := mergo.Merge(&model.Overrides, req.Overrides, mergo.WithOverride); err != nil {
			return err
		}

		if err := InstallModel(basePath, installName, &config, model.Overrides, downloadStatus, enforceScan); err != nil {
			return err
		}

		return nil
	}

	models, err := AvailableGalleryModels(galleries, basePath)
	if err != nil {
		return err
	}

	model := FindGalleryElement(models, name, basePath)
	if model == nil {
		return fmt.Errorf("no model found with name %q", name)
	}

	return applyModel(model)
}

func InstallModel(basePath, nameOverride string, config *ModelConfig, configOverrides map[string]interface{}, downloadStatus func(string, string, string, float64), enforceScan bool) error {
	// Create base path if it doesn't exist
	err := os.MkdirAll(basePath, 0750)
	if err != nil {
		return fmt.Errorf("failed to create base path: %v", err)
	}

	if len(configOverrides) > 0 {
		log.Debug().Msgf("Config overrides %+v", configOverrides)
	}

	// Download files and verify their SHA
	for i, file := range config.Files {
		log.Debug().Msgf("Checking %q exists and matches SHA", file.Filename)

		if err := utils.VerifyPath(file.Filename, basePath); err != nil {
			return err
		}

		// Create file path
		filePath := filepath.Join(basePath, file.Filename)

		if enforceScan {
			scanResults, err := downloader.HuggingFaceScan(downloader.URI(file.URI))
			if err != nil && errors.Is(err, downloader.ErrUnsafeFilesFound) {
				log.Error().Str("model", config.Name).Strs("clamAV", scanResults.ClamAVInfectedFiles).Strs("pickles", scanResults.DangerousPickles).Msg("Contains unsafe file(s)!")
				return err
			}
		}
		uri := downloader.URI(file.URI)
		if err := uri.DownloadFile(filePath, file.SHA256, i, len(config.Files), downloadStatus); err != nil {
			return err
		}
	}

	// Write prompt template contents to separate files
	for _, template := range config.PromptTemplates {
		if err := utils.VerifyPath(template.Name+".tmpl", basePath); err != nil {
			return err
		}
		// Create file path
		filePath := filepath.Join(basePath, template.Name+".tmpl")

		// Create parent directory
		err := os.MkdirAll(filepath.Dir(filePath), 0750)
		if err != nil {
			return fmt.Errorf("failed to create parent directory for prompt template %q: %v", template.Name, err)
		}
		// Create and write file content
		err = os.WriteFile(filePath, []byte(template.Content), 0600)
		if err != nil {
			return fmt.Errorf("failed to write prompt template %q: %v", template.Name, err)
		}

		log.Debug().Msgf("Prompt template %q written", template.Name)
	}

	name := config.Name
	if nameOverride != "" {
		name = nameOverride
	}

	if err := utils.VerifyPath(name+".yaml", basePath); err != nil {
		return err
	}

	// write config file
	if len(configOverrides) != 0 || len(config.ConfigFile) != 0 {
		configFilePath := filepath.Join(basePath, name+".yaml")

		// Read and update config file as map[string]interface{}
		configMap := make(map[string]interface{})
		err = yaml.Unmarshal([]byte(config.ConfigFile), &configMap)
		if err != nil {
			return fmt.Errorf("failed to unmarshal config YAML: %v", err)
		}

		configMap["name"] = name

		if err := mergo.Merge(&configMap, configOverrides, mergo.WithOverride); err != nil {
			return err
		}

		// Write updated config file
		updatedConfigYAML, err := yaml.Marshal(configMap)
		if err != nil {
			return fmt.Errorf("failed to marshal updated config YAML: %v", err)
		}

		backendConfig := lconfig.BackendConfig{}
		err = yaml.Unmarshal(updatedConfigYAML, &backendConfig)
		if err != nil {
			return fmt.Errorf("failed to unmarshal updated config YAML: %v", err)
		}
		if !backendConfig.Validate() {
			return fmt.Errorf("failed to validate updated config YAML")
		}

		err = os.WriteFile(configFilePath, updatedConfigYAML, 0600)
		if err != nil {
			return fmt.Errorf("failed to write updated config file: %v", err)
		}

		log.Debug().Msgf("Written config file %s", configFilePath)
	}

	// Save the model gallery file for further reference
	modelFile := filepath.Join(basePath, galleryFileName(name))
	data, err := yaml.Marshal(config)
	if err != nil {
		return err
	}

	log.Debug().Msgf("Written gallery file %s", modelFile)

	return os.WriteFile(modelFile, data, 0600)

	//return nil
}

func galleryFileName(name string) string {
	return "._gallery_" + name + ".yaml"
}

func GetLocalModelConfiguration(basePath string, name string) (*ModelConfig, error) {
	name = strings.ReplaceAll(name, string(os.PathSeparator), "__")
	galleryFile := filepath.Join(basePath, galleryFileName(name))
	return ReadConfigFile[ModelConfig](galleryFile)
}

func DeleteModelFromSystem(basePath string, name string, additionalFiles []string) error {
	// os.PathSeparator is not allowed in model names. Replace them with "__" to avoid conflicts with file paths.
	name = strings.ReplaceAll(name, string(os.PathSeparator), "__")

	configFile := filepath.Join(basePath, fmt.Sprintf("%s.yaml", name))

	galleryFile := filepath.Join(basePath, galleryFileName(name))

	for _, f := range []string{configFile, galleryFile} {
		if err := utils.VerifyPath(f, basePath); err != nil {
			return fmt.Errorf("failed to verify path %s: %w", f, err)
		}
	}

	var err error
	// Delete all the files associated to the model
	// read the model config
	galleryconfig, err := ReadConfigFile[ModelConfig](galleryFile)
	if err != nil {
		log.Error().Err(err).Msgf("failed to read gallery file %s", configFile)
	}

	var filesToRemove []string

	// Remove additional files
	if galleryconfig != nil {
		for _, f := range galleryconfig.Files {
			fullPath := filepath.Join(basePath, f.Filename)
			filesToRemove = append(filesToRemove, fullPath)
		}
	}

	for _, f := range additionalFiles {
		fullPath := filepath.Join(filepath.Join(basePath, f))
		filesToRemove = append(filesToRemove, fullPath)
	}

	filesToRemove = append(filesToRemove, configFile)
	filesToRemove = append(filesToRemove, galleryFile)

	// skip duplicates
	filesToRemove = utils.Unique(filesToRemove)

	// Removing files
	for _, f := range filesToRemove {
		if e := os.Remove(f); e != nil {
			err = errors.Join(err, fmt.Errorf("failed to remove file %s: %w", f, e))
		}
	}

	return err
}

// This is ***NEVER*** going to be perfect or finished.
// This is a BEST EFFORT function to surface known-vulnerable models to users.
func SafetyScanGalleryModels(galleries []config.Gallery, basePath string) error {
	galleryModels, err := AvailableGalleryModels(galleries, basePath)
	if err != nil {
		return err
	}
	for _, gM := range galleryModels {
		if gM.Installed {
			err = errors.Join(err, SafetyScanGalleryModel(gM))
		}
	}
	return err
}

func SafetyScanGalleryModel(galleryModel *GalleryModel) error {
	for _, file := range galleryModel.AdditionalFiles {
		scanResults, err := downloader.HuggingFaceScan(downloader.URI(file.URI))
		if err != nil && errors.Is(err, downloader.ErrUnsafeFilesFound) {
			log.Error().Str("model", galleryModel.Name).Strs("clamAV", scanResults.ClamAVInfectedFiles).Strs("pickles", scanResults.DangerousPickles).Msg("Contains unsafe file(s)!")
			return err
		}
	}
	return nil
}
