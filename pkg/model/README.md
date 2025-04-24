# Package: model

Imports:
- context
- errors
- fmt
- os
- path/filepath
- slices
- strings
- time
- github.com/klauspost/cpuid/v2
- github.com/mudler/LocalAI/pkg/grpc
- github.com/mudler/LocalAI/pkg/library
- github.com/mudler/LocalAI/pkg/utils
- github.com/mudler/LocalAI/pkg/xsysinfo
- github.com/phayes/freeport
- github.com/rs/zerolog/log
- github.com/elliotchance/orderedmap/v2
- process "github.com/mudler/go-processmanager"

External data, input sources:
- Asset directory containing backend binaries and configurations.
- Environment variables for backend-specific settings and paths.
- User-provided options for model loading and backend selection.

TODOs:
- Add support for more backends and model types.
- Improve error handling and reporting.
- Implement a mechanism for automatic backend selection based on system capabilities.

Summary:
The model package provides a comprehensive solution for managing and interacting with models across various backends. It enables users to easily load and use models without worrying about the underlying backend implementation.

The package is responsible for loading and managing models for various backends. It provides functionalities to list available backends, load models, and interact with them through a gRPC interface.

The package includes several key components:
1. Backend Management:
   - `backendsInAssetDir`: Scans the asset directory for available backends and their configurations.
   - `orderBackends`: Orders the backends based on priority and availability.
   - `selectGRPCProcessByHostCapabilities`: Selects the appropriate backend process based on system capabilities.
2. Model Loading:
   - `grpcModel`: Loads the model using the gRPC interface and handles communication with the backend.
   - `LoadModel`: Loads the model with the specified backend and options.
3. Backend Interaction:
   - `Model`: Represents a loaded model and provides methods for interacting with it through the gRPC interface.
   - `HealthCheck`: Checks the health status of the backend.
   - `LoadModel`: Loads the model with the specified options.

4. Model Management:
   - `ModelLoader`: Manages the loading and unloading of models for different backends.
   - `ListAvailableBackends`: Lists the available backends in the asset directory.
   - `Load`: Loads the model with the specified backend and options.
   - `Close`: Closes the connection to the backend and releases resources.

5. Utility Functions:
   - `attemptLoadingOnFailure`: Attempts to load the model with a fallback backend if the initial attempt fails.
   - `stopActiveBackends`: Stops all active backends except the one being loaded.
   - `lockBackend`: Acquires a lock to ensure only one backend is active at a time.

This package provides a comprehensive solution for managing and interacting with models across various backends. It enables users to easily load and use models without worrying about the underlying backend implementation.

pkg/model/filters.go
Package: model

Imports:
- process "github.com/mudler/go-processmanager"

External data, input sources:
- None

TODOs:
- None

Summary:
The model package provides a GRPCProcessFilter type and two functions, all and allExcept, for filtering processes. The all function returns true for all processes, while the allExcept function takes a string argument and returns true for all processes except the one with the given ID.

pkg/model/initializers.go
## Package: model

Imports:
- context
- errors
- fmt
- os
- path/filepath
- slices
- strings
- time
- github.com/klauspost/cpuid/v2
- github.com/mudler/LocalAI/pkg/grpc
- github.com/mudler/LocalAI/pkg/library
- github.com/mudler/LocalAI/pkg/utils
- github.com/mudler/LocalAI/pkg/xsysinfo
- github.com/phayes/freeport
- github.com/rs/zerolog/log
- github.com/elliotchance/orderedmap/v2

External Data and Input Sources:
- Asset directory containing backend binaries and configurations.
- Environment variables for backend-specific settings and paths.
- User-provided options for model loading and backend selection.

TODOs:
- Add support for more backends and model types.
- Improve error handling and reporting.
- Implement a mechanism for automatic backend selection based on system capabilities.

### Summary of Major Code Parts:

1. Backend Management:
   - `backendsInAssetDir`: Scans the asset directory for available backends and their configurations.
   - `orderBackends`: Orders the backends based on priority and availability.
   - `selectGRPCProcessByHostCapabilities`: Selects the appropriate backend process based on system capabilities.
2. Model Loading:
   - `grpcModel`: Loads the model using the gRPC interface and handles communication with the backend.
   - `LoadModel`: Loads the model with the specified backend and options.
3. Backend Interaction:
   - `Model`: Represents a loaded model and provides methods for interacting with it through the gRPC interface.
   - `HealthCheck`: Checks the health status of the backend.
   - `LoadModel`: Loads the model with the specified options.

4. Model Management:
   - `ModelLoader`: Manages the loading and unloading of models for different backends.
   - `ListAvailableBackends`: Lists the available backends in the asset directory.
   - `Load`: Loads the model with the specified backend and options.
   - `Close`: Closes the connection to the backend and releases resources.

5. Utility Functions:
   - `attemptLoadingOnFailure`: Attempts to load the model with a fallback backend if the initial attempt fails.
   - `stopActiveBackends`: Stops all active backends except the one being loaded.
   - `lockBackend`: Acquires a lock to ensure only one backend is active at a time.

This package provides a comprehensive solution for managing and interacting with models across various backends. It enables users to easily load and use models without worrying about the underlying backend implementation.

pkg/model/loader.go
Package: model

Imports:
- context
- fmt
- os
- path/filepath
- strings
- sync
- time
- github.com/mudler/LocalAI/pkg/utils
- github.com/rs/zerolog/log

External data, input sources:
- The code reads files from the modelPath directory.
- It uses the WatchDog to monitor the modelPath directory for changes.

TODOs:
- Split ModelLoader and TemplateLoader? Just to keep things more organized. Left together to share a mutex until I look into that. Would split if we separate directories for .bin/.yaml and .tmpl

Summary:
- The ModelLoader struct is responsible for loading and managing models. It keeps track of loaded models in a map and provides methods for loading, checking availability, and shutting down models.
- The NewModelLoader function initializes a new ModelLoader instance with the given modelPath and singleActiveBackend flag.
- The SetWatchDog function sets the WatchDog instance for the ModelLoader.
- The ExistsInModelPath function checks if a given string exists in the modelPath directory.
- The ListFilesInModelPath function lists all files in the modelPath directory, excluding certain known files and directories.
- The ListModels function returns a list of all loaded models.
- The LoadModel function loads a model with the given modelID and modelName, using the provided loader function.
- The ShutdownModel function shuts down a model with the given modelName.
- The CheckIsLoaded function checks if a model with the given modelID is already loaded and available.
- The deleteProcess function is used to stop and delete a process associated with a model.

pkg/model/loader_options.go
## Package: model

Imports:
- "context"
- "github.com/mudler/LocalAI/pkg/grpc/proto"

External data, input sources:
- The code uses a context.Context to manage the lifecycle of the model.
- It interacts with external backends through a map of strings.
- It uses a grpc.ModelOptions struct to configure the model.

TODOs:
- There are no TODO comments in the code.

### Options struct
The Options struct holds various configuration options for the model. It includes fields for the backend string, model file, model ID, asset directory, context, gRPC options, external backends, gRPC attempts, gRPC attempts delay, and parallel requests.

### Option functions
The code defines several Option functions that can be used to set the configuration options for the model. These functions include WithExternalBackend, WithGRPCAttempts, WithGRPCAttemptsDelay, WithBackendString, WithDefaultBackendString, WithModel, WithLoadGRPCLoadModelOpts, WithAssetDir, WithContext, and WithModelID.

### NewOptions function
The NewOptions function creates a new Options struct with default values and allows the user to customize the options using the Option functions.

### Summary
The model package provides a way to configure and manage the model's settings. It allows users to specify the backend, model file, model ID, asset directory, context, gRPC options, external backends, gRPC attempts, gRPC attempts delay, and parallel requests. The package also includes functions to set these options and create a new Options struct with the desired configuration.

pkg/model/model.go
## Package: model

Imports:
- "sync"
- "github.com/mudler/LocalAI/pkg/grpc"
- "github.com/mudler/go-processmanager"

External data, input sources:
- The code interacts with a gRPC server, which is an external data source.

TODOs:
- There are no TODO comments in the code.

### Model struct
The Model struct represents a model with an ID, address, gRPC client, and process. It has a mutex to ensure thread safety when accessing the client.

### NewModel function
The NewModel function creates a new Model instance with the given ID, address, and process.

### Process function
The Process function returns the process associated with the model.

### GRPC function
The GRPC function returns a gRPC client for the model. If the client is not already initialized, it creates a new client with the given address, parallel flag, watchdog, and enableWD flag. The watchdog and enableWD flag are used to control the watchdog behavior.

### WatchDog struct
The WatchDog struct is used to monitor the gRPC client and restart it if it fails. It is not defined in this file, but it is assumed to be defined elsewhere in the package.

pkg/model/process.go
Package: model

Imports:
	"errors"
	"fmt"
	"os"
	"os/signal"
	"path/filepath"
	"strconv"
	"strings"
	"syscall"
	"time"
	"github.com/hpcloud/tail"
	process "github.com/mudler/go-processmanager"
	"github.com/rs/zerolog/log"

External data, input sources:
	- Environment variable "LOCALAI_FORCE_BACKEND_SHUTDOWN"

TODOs:
	- None

Summary:
	- The code defines a ModelLoader struct with methods for managing and interacting with GRPC processes.
	- The deleteProcess method handles the deletion of a GRPC process, including stopping the process and cleaning up resources.
	- The StopGRPC method stops a specific GRPC process based on a filter, while StopAllGRPC stops all GRPC processes.
	- The GetGRPCPID method retrieves the process ID of a GRPC process.
	- The startProcess method starts a new GRPC process, sets up logging, and handles process termination.
	- The code also includes error handling and logging mechanisms to ensure proper operation and debugging.

	- The code is designed to manage the lifecycle of GRPC processes, including starting, stopping, and monitoring them.
	- It provides a robust and reliable way to interact with GRPC processes, ensuring proper resource management and error handling.
	- The logging and error handling features help in debugging and troubleshooting any issues that may arise during the process lifecycle.

pkg/model/watchdog.go
Package: model

Imports:
- "sync"
- "time"
- process "github.com/mudler/go-processmanager"
- "github.com/rs/zerolog/log"

External data, input sources:
- None

TODOs:
- None

Summary:
- The WatchDog struct is responsible for monitoring the status of backend processes and taking action if they become unresponsive or idle for too long. It tracks the time each backend has been busy or idle and uses this information to determine when to intervene.
- The NewWatchDog function initializes a new WatchDog instance with the specified timeout durations and checks for busy and idle states.
- The Shutdown function stops the watchdog by closing the stop channel.
- The AddAddressModelMap function adds a mapping between an address and a model name to the watchdog's internal map.
- The Add function adds a process to the watchdog's address map.
- The Mark function records the time a backend becomes busy.
- The UnMark function records the time a backend becomes idle.
- The Run function starts the watchdog's monitoring loop, which checks for busy and idle connections at regular intervals.
- The checkIdle function checks for idle connections and shuts down models that have been idle for too long.
- The checkBusy function checks for busy connections and shuts down models that have been busy for too long.

]