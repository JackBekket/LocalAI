# core/gallery/gallery.go  
## Package: gallery  
  
This package provides functionality for managing models from the gallery.  
  
### External Data and Input Sources  
  
The package relies on the following external data and input sources:  
  
1. A list of galleries, which are YAML files hosted on a remote server.  
2. A base path where models are stored.  
3. A model name to search for.  
4. A request schema containing additional files and overrides.  
5. A download status callback function.  
6. A flag to enforce scanning for vulnerabilities.  
  
### TODOs  
  
1. Determine if using the override method with a blank cfg yaml is worse than using the URL method for gallery models.  
2. Implement a way to handle cases where a model doesn't install a config file.  
  
### Code Summary  
  
1. **InstallModelFromGallery:** This function installs a model from the gallery. It first searches for the model in the available gallery models. If the model is found, it applies the model configuration and installs the model.  
2. **FindModel:** This function searches for a model in the list of available gallery models. It takes the model name and base path as input and returns the model if found.  
3. **AvailableGalleryModels:** This function lists all available models from the given galleries. It iterates through each gallery and retrieves the models from the gallery YAML files.  
4. **getGalleryModels:** This function retrieves the models from a single gallery. It downloads the gallery YAML file and parses it to extract the models.  
5. **GetLocalModelConfiguration:** This function retrieves the configuration for a locally installed model. It reads the model's YAML configuration file and returns the configuration data.  
6. **DeleteModelFromSystem:** This function deletes a model and its associated files from the system. It first reads the model's configuration file to identify the files to be deleted, then deletes the files.  
7. **SafetyScanGalleryModels:** This function scans all installed gallery models for known vulnerabilities. It iterates through each model and calls the SafetyScanGalleryModel function to scan the model.  
8. **SafetyScanGalleryModel:** This function scans a single gallery model for known vulnerabilities. It checks the model's additional files for unsafe content and reports any findings.  
  
### Summary  
  
The gallery package provides a comprehensive set of functions for managing models from the gallery. It allows users to install, update, and delete models, as well as scan them for vulnerabilities. The package relies on external data sources such as YAML files and a base path for storing models.  
  
# core/gallery/models.go  
## Package: gallery  
  
This package is responsible for handling the gallery of models, including downloading, installing, and managing their configurations.  
  
### Imports:  
  
- "errors"  
- "fmt"  
- "os"  
- "path/filepath"  
- "dario.cat/mergo"  
- "github.com/mudler/LocalAI/core/config"  
- "github.com/mudler/LocalAI/pkg/downloader"  
- "github.com/mudler/LocalAI/pkg/utils"  
- "github.com/rs/zerolog/log"  
- "gopkg.in/yaml.v2"  
  
### External Data and Input Sources:  
  
- The package interacts with the gallery endpoint to retrieve model configurations.  
- It reads YAML files containing model configurations.  
- It downloads model files from the specified URLs.  
  
### TODOs:  
  
- There are no TODO comments in the provided code.  
  
### Code Summary:  
  
1. **Model Configuration (Config):**  
   - The `Config` struct represents the model configuration, containing details such as description, license, URLs, files, and prompt templates.  
   - It is used to store and manage the model's metadata and configuration.  
  
2. **File and Prompt Template Structures:**  
   - The `File` struct stores information about each file associated with the model, including filename, SHA256 hash, and URI.  
   - The `PromptTemplate` struct contains the name and content of each prompt template used by the model.  
  
3. **Fetching Gallery Configuration:**  
   - The `GetGalleryConfigFromURL` function retrieves the model configuration from the given URL.  
   - It downloads the YAML file containing the configuration and parses it into a `Config` struct.  
  
4. **Reading Configuration File:**  
   - The `ReadConfigFile` function reads the YAML configuration file from the specified path.  
   - It parses the YAML data into a `Config` struct and returns it.  
  
5. **Installing Model:**  
   - The `InstallModel` function handles the installation of the model.  
   - It downloads the model files, verifies their SHA256 hashes, and saves them to the specified base path.  
   - It also writes the prompt template contents to separate files and updates the configuration file with any provided overrides.  
  
6. **Utility Functions:**  
   - The `galleryFileName` function generates the filename for the gallery file based on the model name.  
  
In summary, the gallery package provides functionalities for managing model configurations, downloading and installing models, and handling prompt templates. It interacts with the gallery endpoint to retrieve model information and downloads the necessary files for the model to function.  
  
# core/gallery/op.go  
Package/Component name: gallery  
  
Imports:  
github.com/mudler/LocalAI/core/config  
  
External data, input sources:  
- GalleryModelName: The name of the gallery model.  
- ConfigURL: The URL of the configuration file.  
- Delete: A boolean value indicating whether the operation is a deletion.  
  
TODOs:  
- There are no TODO comments in the provided code.  
  
Summary of major code parts:  
- GalleryOp: This struct represents a gallery operation. It contains information about the gallery model, the configuration URL, and whether the operation is a deletion.  
- GalleryOpStatus: This struct represents the status of a gallery operation. It contains information about the deletion status, file name, error, processed status, message, progress, total file size, downloaded file size, and gallery model name.  
  
  
  
# core/gallery/request.go  
Package: gallery  
  
Imports:  
- "fmt"  
- "strings"  
- "github.com/mudler/LocalAI/core/config"  
  
External data, input sources:  
- The code interacts with the LocalAI core configuration package to access the Gallery struct.  
  
TODOs:  
- There are no TODO comments in the provided code.  
  
Summary:  
- The GalleryModel struct represents a model in the gallery, containing metadata, configuration, and override information.  
- The GalleryModels type provides methods for searching, finding by name, and paginating a collection of GalleryModel instances.  
- The ID() method returns a unique identifier for a GalleryModel instance.  
- The Search() method filters a collection of GalleryModel instances based on a search term.  
- The FindByName() method finds a GalleryModel instance by its name.  
- The Paginate() method paginates a collection of GalleryModel instances.  
  
  
  
