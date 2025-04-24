# core/services/backend_monitor.go  
Package: services  
  
Imports:  
- context  
- fmt  
- strings  
- github.com/mudler/LocalAI/core/config  
- github.com/mudler/LocalAI/core/schema  
- github.com/mudler/LocalAI/pkg/grpc/proto  
- github.com/mudler/LocalAI/pkg/model  
- github.com/rs/zerolog/log  
- github.com/shirou/gopsutil/v3/process  
  
External data, input sources:  
- Backend configuration data is loaded using the config.BackendConfigLoader.  
- Model data is loaded using the model.ModelLoader.  
  
TODOs:  
- None found.  
  
Code summary:  
- BackendMonitorService: This service monitors the status of backend models and provides information about their resource usage.  
    - getModelLoaderIDFromModelName: This method retrieves the backend ID for a given model name.  
    - SampleLocalBackendProcess: This method samples the resource usage of a local backend process.  
    - CheckAndSample: This method checks if a backend model is loaded and samples its status and resource usage.  
    - ShutdownModel: This method shuts down a backend model.  
  
- BackendMonitorService is responsible for monitoring the status of backend models and providing information about their resource usage. It uses the config.BackendConfigLoader to load backend configuration data and the model.ModelLoader to load model data. The service provides methods to check the status of a backend model, sample its resource usage, and shut down the model.  
  
# core/services/gallery.go  
Package/Component name: services  
  
Imports:  
"context"  
"encoding/json"  
"fmt"  
"os"  
"path/filepath"  
"sync"  
"github.com/mudler/LocalAI/core/config"  
"github.com/mudler/LocalAI/core/gallery"  
"github.com/mudler/LocalAI/pkg/startup"  
"github.com/mudler/LocalAI/pkg/utils"  
"gopkg.in/yaml.v2"  
  
External data, input sources:  
- Configuration files for models and galleries  
- User requests for model installation, deletion, or updates  
  
TODOs:  
- Add support for model updates  
- Implement error handling for model installation and deletion  
- Add logging for model operations  
  
Summary:  
The GalleryService is responsible for managing the installation, deletion, and updates of models and galleries. It receives requests from users and processes them accordingly. The service uses the gallery package to interact with the gallery system and the startup package to install models.  
  
The GalleryService has a channel (C) to receive requests for model operations. When a request is received, the service updates the status of the operation and processes the request. The service uses the gallery package to install or delete models, and it uses the startup package to install models from galleries.  
  
The service also has a method to apply a gallery from a file or a string. This method takes the model path, the gallery content, and a flag to enforce predownload scans as input. It then processes the gallery content and installs the models.  
  
The GalleryService is designed to be used in a multi-threaded environment, and it uses a mutex to protect its internal state. The service is also designed to be extensible, and it can be easily modified to support new model formats or galleries.  
  
# core/services/list_models.go  
## Package: services  
  
Imports:  
- "github.com/mudler/LocalAI/core/config"  
- "github.com/mudler/LocalAI/pkg/model"  
  
External data, input sources:  
- Backend configurations from the `config.BackendConfigLoader`  
- Model files from the `model.ModelLoader`  
  
TODOs:  
- None found.  
  
### ListModels function  
This function lists all available models, considering both configured and loose files. It takes a `BackendConfigLoader`, a `ModelLoader`, a filter function, and a `LooseFilePolicy` as input. The function first iterates through the configured models and adds them to the list if they match the specified policy. Then, it iterates through the loose files and adds them to the list if they match the filter and are not skipped according to the policy. Finally, it returns the list of models and any errors encountered.  
  
### CheckIfModelExists function  
This function checks if a specific model exists, considering both configured and loose files. It takes a `BackendConfigLoader`, a `ModelLoader`, the model name, and a `LooseFilePolicy` as input. The function first builds a filter function based on the model name and then calls the `ListModels` function to get a list of models matching the filter. Finally, it returns true if the list contains any models and false otherwise.  
  
# core/services/metrics.go  
## Package: services  
  
Imports:  
- context  
- github.com/rs/zerolog/log  
- go.opentelemetry.io/otel/attribute  
- go.opentelemetry.io/otel/exporters/prometheus  
- go.opentelemetry.io/otel/metric  
- go.opentelemetry.io/otel/sdk/metric  
  
External data, input sources:  
- The code uses OpenTelemetry to collect metrics about API calls.  
  
TODOs:  
- Not sure how to actually do this: setupOTelSDK bootstraps the OpenTelemetry pipeline. If it does not return an error, make sure to call shutdown for proper cleanup.  
  
### LocalAIMetricsService  
  
This struct is responsible for collecting and reporting metrics related to API calls. It uses OpenTelemetry to collect data and export it to a Prometheus exporter.  
  
### NewLocalAIMetricsService  
  
This function initializes a new LocalAIMetricsService instance. It sets up the OpenTelemetry pipeline and creates a meter and a histogram to track API call durations.  
  
### Shutdown  
  
This method is intended to shut down the OpenTelemetry pipeline and release resources. However, the implementation is incomplete and only logs a warning message.  
  
### ObserveAPICall  
  
This method records the duration of an API call, along with the method and path, using the OpenTelemetry API.  
  
