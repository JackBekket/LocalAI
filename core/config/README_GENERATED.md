## Package: config

This package is responsible for handling the application configuration. It defines the `ApplicationConfig` struct, which stores all the necessary settings for the application. The package also provides functions to create and modify the configuration.

### External Data and Input Sources

The package relies on several external data sources and input sources, including:

1. Configuration files: The `ApplicationConfig` struct can be loaded from configuration files, which contain settings for the application.
2. Environment variables: The package can read environment variables to override default settings.
3. Command-line arguments: The package can accept command-line arguments to modify the configuration.

### TODOs

There are no TODO comments in the provided code.

### Code Summary

1. **ApplicationConfig Struct:**
   - The `ApplicationConfig` struct stores all the necessary settings for the application.
   - It includes settings for the model path, library path, upload limit, number of threads, context size, and more.
   - The struct also includes options for enabling or disabling various features, such as the watchdog, CORS, and CSRF protection.

2. **AppOption Function Type:**
   - The `AppOption` function type is used to create and modify the `ApplicationConfig` struct.
   - It takes a pointer to the `ApplicationConfig` struct as input and modifies its fields.

3. **NewApplicationConfig Function:**
   - The `NewApplicationConfig` function creates a new instance of the `ApplicationConfig` struct with default settings.
   - It takes a variable number of `AppOption` functions as input to customize the configuration.

4. **AppOption Functions:**
   - The package provides numerous `AppOption` functions to customize the `ApplicationConfig` struct.
   - These functions allow users to set the model path, library path, upload limit, number of threads, context size, and more.
   - They also allow users to enable or disable various features, such as the watchdog, CORS, and CSRF protection.

5. **ToConfigLoaderOptions Function:**
   - The `ToConfigLoaderOptions` function converts the `ApplicationConfig` struct into a slice of `ConfigLoaderOption` values.
   - This allows the configuration to be used by the `ConfigLoader` component.

In summary, the config package provides a comprehensive set of tools for managing the application configuration. It allows users to customize the settings and features of the application through a variety of options and functions.

core/config/backend_config.go
## Package/Component Name: config

This package is responsible for handling configuration settings for the LocalAI project. It includes data structures and functions to manage backend configurations, such as the BackendConfig struct, which stores various settings related to the backend, including prediction options, backend-specific parameters, and file download information.

## External Data and Input Sources:

- The package relies on YAML files for storing configuration data.
- It may also interact with external sources to download files specified in the configuration.

## TODOs:

- Add more detailed documentation for the BackendConfig struct and its fields.
- Implement error handling for file download failures.
- Add support for loading configuration from environment variables.

## Summary of Major Code Parts:

1. Data Structures:
    - BackendConfig: Stores backend-specific configuration settings, including prediction options, backend parameters, and file download information.
    - File: Represents a file to be downloaded, including filename, SHA256 hash, and URI.
    - FeatureFlag: A map of feature flags, which can be enabled or disabled.
    - GRPC: Configuration for GRPC options.
    - Diffusers: Configuration for diffusers.
    - LLMConfig: Configuration for LLM backends.
    - LimitMMPerPrompt: Configuration for limiting memory usage per prompt.
    - TemplateConfig: Configuration for templating systems.

2. Functions:
    - UnmarshalYAML: Decodes YAML data into a BackendConfig struct.
    - SetFunctionCallString: Sets the function call string for the backend.
    - SetFunctionCallNameString: Sets the function call name string for the backend.
    - ShouldUseFunctions: Checks if the backend should use functions.
    - ShouldCallSpecificFunction: Checks if the backend should call a specific function.
    - MMProjFileName: Returns the filename of the MMProj file.
    - IsMMProjURL: Checks if the MMProj is a URL.
    - IsModelURL: Checks if the model is a URL.
    - ModelFileName: Returns the filename of the model.
    - FunctionToCall: Returns the function to call.
    - SetDefaults: Sets default values for the backend configuration.
    - Validate: Validates the backend configuration.
    - HasTemplate: Checks if the backend has a template.
    - HasUsecases: Checks if the backend has the specified use cases.
    - GuessUsecases: Guesses the use cases for the backend.

3. Constants:
    - RAND_SEED: A constant value for the random number generator seed.

4. Interfaces:
    - ConfigLoaderOption: An interface for configuration loader options.

5. Types:
    - BackendConfigUsecases: An enumeration of backend use cases.

6. Utilities:
    - guessDefaultsFromFile: A utility function to guess default values from a file.

This package provides a comprehensive set of tools for managing backend configurations, enabling users to easily customize and control the behavior of their LocalAI applications.

core/config/backend_config_filter.go
Package: config

Imports:
regexp

External data, input sources:
- BackendConfigFilterFn: A function type that takes a string and a BackendConfig pointer as input and returns a boolean value.
- BackendConfigUsecases: An enum type that represents the use cases of a backend configuration.

TODOs:
- Potentially make the return value of the BuildUsecaseFilterFn function a parameter, for now, no known use case to include.

Summary:
The config package provides functions for filtering backend configurations. The BuildNameFilterFn function takes a regular expression as input and returns a filter function that matches backend configurations based on their name. The BuildUsecaseFilterFn function takes a set of use cases as input and returns a filter function that matches backend configurations based on their use cases. Both filter functions can be used to select specific backend configurations from a list.



core/config/backend_config_loader.go
Package: config

Imports:
- "errors"
- "fmt"
- "io/fs"
- "os"
- "path/filepath"
- "sort"
- "strings"
- "sync"
- "github.com/charmbracelet/glamour"
- "github.com/mudler/LocalAI/core/schema"
- "github.com/mudler/LocalAI/pkg/downloader"
- "github.com/mudler/LocalAI/pkg/utils"
- "github.com/rs/zerolog/log"
- "gopkg.in/yaml.v3"

External data and input sources:
- Configuration files in YAML format

TODOs:
- Merge the functions readMultipleBackendConfigsFromFile and readBackendConfigFromFile into a single function that looks at the first few characters of the file to determine if we need to deserialize to []BackendConfig or BackendConfig.

Summary:
- The BackendConfigLoader struct is responsible for loading and managing backend configurations. It stores a map of BackendConfig objects, each representing a model's configuration.
- The LoadBackendConfigFileByName function loads a configuration file for a given model name. If a model-specific config file is not found, it attempts to load a default config file.
- The LoadBackendConfigFileByNameDefaultOptions function loads a configuration file with default options.
- The LoadMultipleBackendConfigsSingleFile function loads multiple backend configurations from a single file.
- The LoadBackendConfig function loads a backend configuration from a file.
- The GetBackendConfig function retrieves a backend configuration by its name.
- The GetAllBackendConfigs function returns a list of all loaded backend configurations.
- The GetBackendConfigsByFilter function returns a list of backend configurations that match a given filter.
- The RemoveBackendConfig function removes a backend configuration by its name.
- The Preload function prepares models by downloading and verifying files if they are not local.
- The LoadBackendConfigsFromPath function loads backend configurations from a given path.

This package provides a comprehensive solution for managing backend configurations, including loading, storing, and retrieving configurations, as well as handling the preload process for models.

core/config/backend_config_test.go
Package: config

Imports:
- "io"
- "net/http"
- "os"
- . "github.com/onsi/ginkgo/v2"
- . "github.com/onsi/gomega"

External data, input sources:
- Configuration files (e.g., config.yaml)

TODOs:
- None found.

Summary:
- The package contains test cases for configuration-related functions.
- It includes tests for reading and validating configuration files.
- The tests cover various scenarios, such as checking if a configuration is valid and handling backend use case matching.
- The package also includes functions for handling backend use case matching, such as checking if a configuration has specific use cases.



core/config/config_test.go
Package name: config

Imports:
os
github.com/onsi/ginkgo/v2
github.com/onsi/gomega

External data, input sources:
- CONFIG_FILE environment variable
- MODELS_PATH environment variable

TODOs:
- None

Summary:
- The code defines test cases for configuration-related functions.
- It tests the functionality of reading configuration files and loading backend configurations from a given path.
- The tests cover scenarios such as reading multiple backend configurations from a file, loading backend configurations from a path, and loading backend configurations from a temporary directory.
- The tests verify that the loaded configurations contain the expected model names and descriptions.



core/config/gallery.go
Package: config

Imports:
- "encoding/json"
- "gopkg.in/yaml.v2"

External data, input sources:
- The code reads configuration data from JSON or YAML files.

TODOs:
- There are no TODO comments in the code.

Summary:
- The `Gallery` struct is used to store information about a gallery, including its URL and name. It is designed to be used with JSON and YAML configuration files.



core/config/gguf.go
```
Package: config

Imports:
- "strings"
- "github.com/rs/zerolog/log"
- "github.com/thxcode/gguf-parser-go"

External data and input sources:
- GGUF files

TODOs:
- Add support for more model families.

Summary:
- The config package provides functionality for configuring and managing settings for various model families.
- It includes a set of default settings for each supported model family, which can be used to customize the behavior of the model.
- The package also includes functions for identifying the model family and guessing the appropriate settings based on the GGUF file.
- The guessDefaultsFromFile function is responsible for identifying the model family and guessing the appropriate settings based on the GGUF file.
- The identifyFamily function is responsible for identifying the model family based on the GGUF file.
- The package also includes a set of known templates that can be used to identify the model family.
- The package is designed to be flexible and extensible, allowing users to add support for new model families and settings.

- The package includes a set of default settings for each supported model family, which can be used to customize the behavior of the model.
- The package also includes functions for identifying the model family and guessing the appropriate settings based on the GGUF file.
- The guessDefaultsFromFile function is responsible for identifying the model family and guessing the appropriate settings based on the GGUF file.
- The identifyFamily function is responsible for identifying the model family based on the GGUF file.
- The package also includes a set of known templates that can be used to identify the model family.
- The package is designed to be flexible and extensible, allowing users to add support for new model families and settings.

core/config/guesser.go
Package: config

Imports:
- "os"
- "path/filepath"
- "github.com/rs/zerolog/log"
- "github.com/thxcode/gguf-parser-go"

External data, input sources:
- The code reads the value of the environment variable "LOCALAI_DISABLE_GUESSING".
- It also reads the value of the "modelPath" variable, which is assumed to be provided by the caller.

TODOs:
- There are no TODO comments in the code.

Summary:
The code defines a function called guessDefaultsFromFile that attempts to guess default values for a BackendConfig object based on a given model path. It first checks if guessing is disabled by an environment variable. If not, it tries to parse a GGUF file from the given model path. If the parsing is successful, it calls another function to guess the GGUF values. If the parsing fails, it sets the context size to a default value if it hasn't been set already.

This function is likely used to automatically configure a backend configuration object based on the available model files. It demonstrates the use of environment variables and external file parsing to determine default values for the configuration.

]

%!s(MISSING)
%!s(MISSING)