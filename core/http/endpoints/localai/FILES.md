# core/http/endpoints/localai/backend_monitor.go  
Package: localai  
  
Imports:  
- "github.com/gofiber/fiber/v2"  
- "github.com/mudler/LocalAI/core/schema"  
- "github.com/mudler/LocalAI/core/services"  
  
External data, input sources:  
- The code uses the `schema.BackendMonitorRequest` struct to receive input data from the request body.  
  
TODOs:  
- There are no TODO comments in the code.  
  
Summary:  
- The `BackendMonitorEndpoint` function handles requests to the `/backend/monitor` endpoint. It takes a `schema.BackendMonitorRequest` object as input, checks the status of the specified backend, and returns a `proto.StatusResponse` object.  
- The `BackendShutdownEndpoint` function handles requests to the `/backend/shutdown` endpoint. It takes a `schema.BackendMonitorRequest` object as input, shuts down the specified backend, and returns a response.  
  
  
  
# core/http/endpoints/localai/gallery.go  
## Package: localai  
  
This package provides endpoints for managing model galleries and applying models to a LocalAI instance.  
  
### Imports:  
  
- "encoding/json"  
- "fmt"  
- "slices"  
- "github.com/gofiber/fiber/v2"  
- "github.com/google/uuid"  
- "github.com/mudler/LocalAI/core/config"  
- "github.com/mudler/LocalAI/core/gallery"  
- "github.com/mudler/LocalAI/core/http/utils"  
- "github.com/mudler/LocalAI/core/schema"  
- "github.com/mudler/LocalAI/core/services"  
- "github.com/rs/zerolog/log"  
  
### External Data and Input Sources:  
  
- The code interacts with a LocalAI instance, which is assumed to be running and accessible.  
- It uses the `gallery.GalleryOp` struct to communicate with the LocalAI instance's gallery service.  
- It relies on the `config.Gallery` struct to store information about model galleries.  
  
### TODOs:  
  
- There are no TODO comments in the provided code.  
  
### Code Summary:  
  
1. **ModelGalleryEndpointService:** This struct is responsible for handling all the endpoints related to model galleries. It stores a list of available galleries, the path to the model directory, and a reference to the gallery service.  
  
2. **GalleryModel:** This struct represents a model from a gallery. It contains the model's ID, the URL to its configuration, and the actual gallery model data.  
  
3. **Endpoints:** The code defines several endpoints for managing model galleries and applying models to a LocalAI instance. These endpoints include:  
    - `GetOpStatusEndpoint`: Returns the status of a job.  
    - `GetAllStatusEndpoint`: Returns the status of all jobs.  
    - `ApplyModelGalleryEndpoint`: Installs a new model from a gallery to a LocalAI instance.  
    - `DeleteModelGalleryEndpoint`: Deletes a model from a LocalAI instance.  
    - `ListModelFromGalleryEndpoint`: Lists the available models for installation from the active galleries.  
    - `ListModelGalleriesEndpoint`: Lists the available galleries configured in LocalAI.  
    - `AddModelGalleryEndpoint`: Adds a new gallery to LocalAI.  
    - `RemoveModelGalleryEndpoint`: Removes a gallery from LocalAI.  
  
4. **Communication with LocalAI:** The code uses the `gallery.GalleryOp` struct to communicate with the LocalAI instance's gallery service. This struct contains information about the operation to be performed, such as installing or deleting a model, as well as the necessary parameters.  
  
5. **Error Handling:** The code includes error handling for various scenarios, such as when a job cannot be found or when a model cannot be installed.  
  
6. **Logging:** The code uses the `log` package to log debug messages and errors.  
  
7. **JSON Serialization:** The code uses the `encoding/json` package to serialize and deserialize data in JSON format.  
  
8. **Concurrency:** The code uses channels to communicate with the LocalAI instance's gallery service, allowing for concurrent operations.  
  
In summary, this package provides a set of endpoints for managing model galleries and applying models to a LocalAI instance. It interacts with the LocalAI instance's gallery service to perform operations such as installing, deleting, and listing models and galleries. The code includes error handling, logging, and concurrency features to ensure reliable and efficient operation.  
  
# core/http/endpoints/localai/get_token_metrics.go  
Package: localai  
  
Imports:  
- github.com/gofiber/fiber/v2  
- github.com/mudler/LocalAI/core/backend  
- github.com/mudler/LocalAI/core/config  
- github.com/mudler/LocalAI/core/http/middleware  
- github.com/mudler/LocalAI/core/schema  
- github.com/rs/zerolog/log  
- github.com/mudler/LocalAI/pkg/model  
  
External data, input sources:  
- The code uses a BackendConfigLoader to load backend configuration files.  
- It uses a ModelLoader to load model files.  
- It uses an ApplicationConfig to access application-wide configuration settings.  
- The TokenMetricsEndpoint function takes a TokenMetricsRequest object as input, which contains information about the model and the slot ID.  
  
TODOs:  
- TODO: This is not yet in use. Needs middleware rework, since it is not referenced.  
  
Code summary:  
- The TokenMetricsEndpoint function is designed to provide token metrics for a given model and slot ID. It takes a TokenMetricsRequest object as input, which contains information about the model and the slot ID.  
- The function first attempts to retrieve the model file name from the request context. If the model file name is not found in the context, it uses the model name provided in the TokenMetricsRequest object.  
- Next, the function loads the backend configuration file for the specified model using the BackendConfigLoader. If the configuration file cannot be loaded, it logs an error and uses the model name provided in the TokenMetricsRequest object.  
- Once the configuration file is loaded, the function calls the TokenMetrics function from the backend package to retrieve the token metrics for the specified model and slot ID.  
- Finally, the function returns the token metrics as a JSON response.  
  
  
  
# core/http/endpoints/localai/metrics.go  
Package: localai  
  
Imports:  
- "time"  
- "github.com/gofiber/fiber/v2"  
- "github.com/gofiber/fiber/v2/middleware/adaptor"  
- "github.com/mudler/LocalAI/core/services"  
- "github.com/prometheus/client_golang/prometheus/promhttp"  
  
External data, input sources:  
- The code uses the Prometheus client library to expose metrics.  
  
TODOs:  
- There are no TODO comments in the code.  
  
Summary of major code parts:  
- LocalAIMetricsEndpoint: This function returns a handler for the Prometheus metrics endpoint. It uses the promhttp.Handler() function to handle the requests.  
- LocalAIMetricsAPIMiddleware: This function takes a LocalAIMetricsService instance as input and returns a middleware handler. The middleware handler measures the time taken for each API call and reports the metrics to the LocalAIMetricsService instance.  
  
- The code defines a struct called apiMiddlewareConfig which is used to configure the LocalAIMetricsAPIMiddleware middleware. It contains a Filter function and a metricsService field.  
- The Filter function checks if the request path is "/metrics". If it is, the middleware handler will allow the request to pass through.  
- The metricsService field is used to store the LocalAIMetricsService instance.  
- The code also defines a function called LocalAIMetricsAPIMiddleware which takes a LocalAIMetricsService instance as input and returns a middleware handler.  
- The middleware handler measures the time taken for each API call and reports the metrics to the LocalAIMetricsService instance.  
- The code defines a function called LocalAIMetricsEndpoint which returns a handler for the Prometheus metrics endpoint. It uses the promhttp.Handler() function to handle the requests.  
  
- The code defines a struct called apiMiddlewareConfig which is used to configure the LocalAIMetricsAPIMiddleware middleware. It contains a Filter function and a metricsService field.  
- The Filter function checks if the request path is "/metrics". If it is, the middleware handler will allow the request to pass through.  
- The metricsService field is used to store the LocalAIMetricsService instance.  
- The code also defines a function called LocalAIMetricsAPIMiddleware which takes a LocalAIMetricsService instance as input and returns a middleware handler.  
- The middleware handler measures the time taken for each API call and reports the metrics to the LocalAIMetricsService instance.  
- The code defines a function called LocalAIMetricsEndpoint which returns a handler for the Prometheus metrics endpoint. It uses the promhttp.Handler() function to handle the requests.  
  
- The code defines a struct called apiMiddlewareConfig which is used to configure the LocalAIMetricsAPIMiddleware middleware. It contains a Filter function and a metricsService field.  
- The Filter function checks if the request path is "/metrics". If it is, the middleware handler will allow the request to pass through.  
- The metricsService field is used to store the LocalAIMetricsService instance.  
- The code also defines a function called LocalAIMetricsAPIMiddleware which takes a LocalAIMetricsService instance as input and returns a middleware handler.  
- The middleware handler measures the time taken for each API call and reports the metrics to the LocalAIMetricsService instance.  
- The code defines a function called LocalAIMetricsEndpoint which returns a handler for the Prometheus metrics endpoint. It uses the promhttp.Handler() function to handle the requests.  
  
- The code defines a struct called apiMiddlewareConfig which is used to configure the LocalAIMetricsAPIMiddleware middleware. It contains a Filter function and a metricsService field.  
- The Filter function checks if the request path is "/metrics". If it is, the middleware handler will allow the request to pass through.  
- The metricsService field is used to store the LocalAIMetricsService instance.  
- The code also defines a function called LocalAIMetricsAPIMiddleware which takes a LocalAIMetricsService instance as input and returns a middleware handler.  
- The middleware handler measures the time taken for each API call and reports the metrics to the LocalAIMetricsService instance.  
- The code defines a function called LocalAIMetricsEndpoint which returns a handler for the Prometheus metrics endpoint. It uses the promhttp.Handler() function to handle the requests.  
  
- The code defines a struct called apiMiddlewareConfig which is used to configure the LocalAIMetricsAPIMiddleware middleware. It contains a Filter function and a metricsService field.  
- The Filter function checks if the request path is "/metrics". If it is, the middleware handler will allow the request to pass through.  
- The metricsService field is used to store the LocalAIMetricsService instance.  
- The code also defines a function called LocalAIMetricsAPIMiddleware which takes a LocalAIMetricsService instance as input and returns a middleware handler.  
- The middleware handler measures the time taken for each API call and reports the metrics to the LocalAIMetricsService instance.  
- The code defines a function called LocalAIMetricsEndpoint which returns a handler for the Prometheus metrics endpoint. It uses the promhttp.Handler() function to handle the requests.  
  
- The code defines a struct called apiMiddlewareConfig which is used to configure the LocalAIMetricsAPIMiddleware middleware. It contains a Filter function and a metricsService field.  
- The Filter function checks if the request path is "/metrics". If it is, the middleware handler will allow the request to pass through.  
- The metricsService field is used to store the LocalAIMetricsService instance.  
- The code also defines a function called LocalAIMetricsAPIMiddleware which takes a LocalAIMetricsService instance as input and returns a middleware handler.  
- The middleware handler measures the time taken for each API call and reports the metrics to the LocalAIMetricsService instance.  
- The code defines a function called LocalAIMetricsEndpoint which returns a handler for the Prometheus metrics endpoint. It uses the promhttp.Handler() function to handle the requests.  
  
- The code defines a struct called apiMiddlewareConfig which is used to configure the LocalAIMetricsAPIMiddleware middleware. It contains a Filter function and a metricsService field.  
- The Filter function checks if the request path is "/metrics". If it is, the middleware handler will allow the request to pass through.  
- The metricsService field is used to store the LocalAIMetricsService instance.  
- The code also defines a function called LocalAIMetricsAPIMiddleware which takes a LocalAIMetricsService instance as input and returns a middleware handler.  
- The middleware handler measures the time taken for each API call and reports the metrics to the LocalAIMetricsService instance.  
- The code defines a function called LocalAIMetricsEndpoint which returns a handler for the Prometheus metrics endpoint. It uses the promhttp.Handler() function to handle the requests.  
  
- The code defines a struct called apiMiddlewareConfig which is used to configure the LocalAIMetricsAPIMiddleware middleware. It contains a Filter function and a metricsService field.  
- The Filter function checks if the request path is "/metrics". If it is, the middleware handler will allow the request to pass through.  
- The metricsService field is used to store the LocalAIMetricsService instance.  
- The code also defines a function called LocalAIMetricsAPIMiddleware which takes a LocalAIMetricsService instance as input and returns a middleware handler.  
- The middleware handler measures the time taken for each API call and reports the metrics to the LocalAIMetricsService instance.  
- The code defines a function called LocalAIMetricsEndpoint which returns a handler for the Prometheus metrics endpoint. It uses the promhttp.Handler() function to handle the requests.  
  
- The code defines a struct called apiMiddlewareConfig which is used to configure the LocalAIMetricsAPIMiddleware middleware. It contains a Filter function and a metricsService field.  
- The Filter function checks if the request path is "/metrics". If it is, the middleware handler will allow the request to pass through.  
- The metricsService field is used to store the LocalAIMetricsService instance.  
- The code also defines a function called LocalAIMetricsAPIMiddleware which takes a LocalAIMetricsService instance as input and returns a middleware handler.  
- The middleware handler measures the time taken for each API call and reports the metrics to the LocalAIMetricsService instance.  
- The code defines a function called LocalAIMetricsEndpoint which returns a handler for the Prometheus metrics endpoint. It uses the promhttp.Handler() function to handle the requests.  
  
- The code defines a struct called apiMiddlewareConfig which is used to configure the LocalAIMetricsAPIMiddleware middleware. It contains a Filter function and a metricsService field.  
- The Filter function checks if the request path is "/metrics". If it is, the middleware handler will allow the request to pass through.  
- The metricsService field is used to store the LocalAIMetricsService instance.  
- The code also defines a function called LocalAIMetricsAPIMiddleware which takes a LocalAIMetricsService instance as input and returns a middleware handler.  
- The middleware handler measures the time taken for each API call and reports the metrics to the LocalAIMetricsService instance.  
- The code defines a function called LocalAIMetricsEndpoint which returns a handler for the Prometheus metrics endpoint. It uses the promhttp.Handler() function to handle the requests.  
  
- The code defines a struct called apiMiddlewareConfig which is used to configure the LocalAIMetricsAPIMiddleware middleware. It contains a Filter function and a metricsService field.  
- The Filter function checks if the request path is "/metrics". If it is, the middleware handler will allow the request to pass through.  
- The metricsService field is used to store the LocalAIMetricsService instance.  
- The code also defines a function called LocalAIMetricsAPIMiddleware which takes a LocalAIMetricsService instance as input and returns a middleware handler.  
- The middleware handler measures the time taken for each API call and reports the metrics to the LocalAIMetricsService instance.  
- The code defines a function called LocalAIMetricsEndpoint which returns a handler for the Prometheus metrics endpoint. It uses the promhttp.Handler() function to handle the requests.  
  
- The code defines a struct called apiMiddlewareConfig which is used to configure the LocalAIMetricsAPIMiddleware middleware. It contains a Filter function and a metricsService field.  
- The Filter function checks if the request path is "/metrics". If it is, the middleware handler will allow the request to pass through.  
- The metricsService field is used to store the LocalAIMetricsService instance.  
- The code also defines a function called LocalAIMetricsAPIMiddleware which takes a LocalAIMetricsService instance as input and returns a middleware handler.  
- The middleware handler measures the time taken for each API call and reports the metrics to the LocalAIMetricsService instance.  
- The code defines a function called LocalAIMetricsEndpoint which returns a handler for the Prometheus metrics endpoint. It uses the promhttp.Handler() function to handle the requests.  
  
- The code defines a struct called apiMiddlewareConfig which is used to configure the LocalAIMetricsAPIMiddleware middleware. It contains a Filter function and a metricsService field.  
- The Filter function checks if the request path is "/metrics". If it is, the middleware handler will allow the request to pass through.  
- The metricsService field is used to store the LocalAIMetricsService instance.  
- The code also defines a function called LocalAIMetricsAPIMiddleware which takes a LocalAIMetricsService instance as input and returns a middleware handler.  
- The middleware handler measures the time taken for each API call and reports the metrics to the LocalAIMetricsService instance.  
- The code defines a function called LocalAIMetricsEndpoint which returns a handler for the Prometheus metrics endpoint. It uses the promhttp.Handler() function to handle the requests.  
  
- The code defines a struct called apiMiddlewareConfig which is used to configure the LocalAIMetricsAPIMiddleware middleware. It contains a Filter function and a metricsService field.  
- The Filter function checks if the request path is "/metrics". If it is, the middleware handler will allow the request to pass through.  
- The metricsService field is used to store the LocalAIMetricsService instance.  
- The code also defines a function called LocalAIMetricsAPIMiddleware which takes a LocalAIMetricsService instance as input and returns a middleware handler.  
- The middleware handler measures the time taken for each API call and reports the metrics to the LocalAIMetricsService instance.  
- The code defines a function called LocalAIMetricsEndpoint which returns a handler for the Prometheus metrics endpoint. It uses the promhttp.Handler() function to handle the requests.  
  
- The code defines a struct called apiMiddlewareConfig which is used to configure the LocalAIMetricsAPIMiddleware middleware. It contains a Filter function and a metricsService field.  
- The Filter function checks if the request path is "/metrics". If it is, the middleware handler will allow the request to pass through.  
- The metricsService field is used to store the LocalAIMetricsService instance.  
- The code also defines a function called LocalAIMetricsAPIMiddleware which takes a LocalAIMetricsService instance as input and returns a middleware handler.  
- The middleware handler measures the time taken for each API call and reports the metrics to the LocalAIMetricsService instance.  
- The code defines a function called LocalAIMetricsEndpoint which returns a handler for the Prometheus metrics endpoint. It uses the promhttp.Handler() function to handle the requests.  
  
- The code defines a struct called apiMiddlewareConfig which is used to configure the LocalAIMetricsAPIMiddleware middleware. It contains a Filter function and a metricsService field.  
- The Filter function checks if the request path is "/metrics". If it is, the middleware handler will allow the request to pass through.  
- The metricsService field is used to store the LocalAIMetricsService instance.  
- The code also defines a function called LocalAIMetricsAPIMiddleware which takes a LocalAIMetricsService instance as input and returns a middleware handler.  
- The middleware handler measures the time taken for each API call and reports the metrics to the LocalAIMetricsService instance.  
- The code defines a function called LocalAIMetricsEndpoint which returns a handler for the Prometheus metrics endpoint. It uses the promhttp.Handler() function to handle the requests.  
  
- The code defines a struct called apiMiddlewareConfig which is used to configure the LocalAIMetricsAPIMiddleware middleware. It contains a Filter function and a metricsService field.  
- The Filter function checks if the request path is "/metrics". If it is, the middleware handler will allow the request to pass through.  
- The metricsService field is used to store the LocalAIMetricsService instance.  
- The code also defines a function called LocalAIMetricsAPIMiddleware which takes a LocalAIMetricsService instance as input and returns a middleware handler.  
- The middleware handler measures the time taken for each API call and reports the metrics to the LocalAIMetricsService instance.  
- The code defines a function called LocalAIMetricsEndpoint which returns a handler for the Prometheus metrics endpoint. It uses the promhttp.Handler() function to handle the requests.  
  
- The code defines a struct called apiMiddlewareConfig which is used to configure the LocalAIMetricsAPIMiddleware middleware. It contains a Filter function and a metricsService field.  
- The Filter function checks if the request path is "/metrics". If it is, the middleware handler will allow the request to pass through.  
- The metricsService field is used to store the LocalAIMetricsService instance.  
- The code also defines a function called LocalAIMetricsAPIMiddleware which takes a LocalAIMetricsService instance as input and returns a middleware handler.  
- The middleware handler measures the time taken for each API call and reports the metrics to the LocalAIMetricsService instance.  
- The code defines a function called LocalAIMetricsEndpoint which returns a handler for the Prometheus metrics endpoint. It uses the promhttp.Handler() function to handle the requests.  
  
- The code defines a struct called apiMiddlewareConfig which is used to configure the LocalAIMetricsAPIMiddleware middleware. It contains a Filter function and a metricsService field.  
- The Filter function checks if the request path is "/metrics". If it is, the middleware handler will allow the request to pass through.  
- The metricsService field is used to store the LocalAIMetricsService instance.  
- The code also defines a function called LocalAIMetricsAPIMiddleware which takes a LocalAIMetricsService instance as input and returns a middleware handler.  
- The middleware handler measures the time taken for each API call and reports the metrics to the LocalAIMetricsService instance.  
- The code defines a function called LocalAIMetricsEndpoint which returns a handler for the Prometheus metrics endpoint. It uses the promhttp.Handler() function to handle the requests.  
  
- The code defines a struct called apiMiddlewareConfig which is used to configure the LocalAIMetricsAPIMiddleware middleware. It contains a Filter function and a metricsService field.  
- The Filter function checks if the request path is "/metrics". If it is, the middleware handler will allow the request to pass through.  
- The metricsService field is used to store the LocalAIMetricsService instance.  
- The code also defines a function called LocalAIMetricsAPIMiddleware which takes a LocalAIMetricsService instance as input and returns a middleware handler.  
- The middleware handler measures the time taken for each API call and reports the metrics to the LocalAIMetricsService instance.  
- The code defines a function called LocalAIMetricsEndpoint which returns a handler for the Prometheus metrics endpoint. It uses the promhttp.Handler() function to handle the requests.  
  
- The code defines a struct called apiMiddlewareConfig which is used to configure the LocalAIMetricsAPIMiddleware middleware. It contains a Filter function and a metricsService field.  
- The Filter function checks if the request path is "/metrics". If it is, the middleware handler will allow the request to pass through.  
- The metricsService field is used to store the LocalAIMetricsService instance.  
- The code also defines a function called LocalAIMetricsAPIMiddleware which takes a LocalAIMetricsService instance as input and returns a middleware handler.  
- The middleware handler measures the time taken for each API call and reports the metrics to the LocalAIMetricsService instance.  
- The code defines a function called LocalAIMetricsEndpoint which returns a handler for the Prometheus metrics endpoint. It uses the promhttp.Handler() function to handle the requests.  
  
- The code defines a struct called apiMiddlewareConfig which is used to configure the LocalAIMetricsAPIMiddleware middleware. It contains a Filter function and a metricsService field.  
- The Filter function checks if the request path is "/metrics". If it is, the middleware handler will allow the request to pass through.  
- The metricsService field is used to store the LocalAIMetricsService instance.  
- The code also defines a function called LocalAIMetricsAPIMiddleware which takes a LocalAIMetricsService instance as input and returns a middleware handler.  
- The middleware handler measures the time taken for each API call and reports the metrics to the LocalAIMetricsService instance.  
- The code defines a function called LocalAIMetricsEndpoint which returns a handler for the Prometheus metrics endpoint. It uses the promhttp.Handler() function to handle the requests.  
  
- The code defines a struct called apiMiddlewareConfig which is used to configure the LocalAIMetricsAPIMiddleware middleware. It contains a Filter function and a metricsService field.  
- The Filter function checks if the request path is "/metrics". If it is, the middleware handler will allow the request to pass through.  
- The metricsService field is used to store the LocalAIMetricsService instance.  
- The code also defines a function called LocalAIMetricsAPIMiddleware which takes a LocalAIMetricsService instance as input and returns a middleware handler.  
- The middleware handler measures the time taken for each API call and reports the metrics to the LocalAIMetricsService instance.  
- The code defines a function called LocalAIMetricsEndpoint which returns a handler for the Prometheus metrics endpoint. It uses the promhttp.Handler() function to handle the requests.  
  
- The code defines a struct called apiMiddlewareConfig which is used to configure the LocalAIMetricsAPIMiddleware middleware. It contains a Filter function and a metricsService field.  
- The Filter function checks if the request path is "/metrics". If it is, the middleware handler will allow the request to pass through.  
- The metricsService field is used to store the LocalAIMetricsService instance.  
- The code also defines a function called LocalAIMetricsAPIMiddleware which takes a LocalAIMetricsService instance as input and returns a middleware handler.  
- The middleware handler measures the time taken for each API call and reports the metrics to the LocalAIMetricsService instance.  
- The code defines a function called LocalAIMetricsEndpoint which returns a handler for the Prometheus metrics endpoint. It uses the promhttp.Handler() function to handle the requests.  
  
- The code defines a struct called apiMiddlewareConfig which is used to configure the LocalAIMetricsAPIMiddleware middleware. It contains a Filter function and a metricsService field.  
- The Filter function checks if the request path is "/metrics". If it is, the middleware handler will allow the request to pass through.  
- The metricsService field is used to store the LocalAIMetricsService instance.  
- The code also defines a function called LocalAIMetricsAPIMiddleware which takes a LocalAIMetricsService instance as input and returns a middleware handler.  
- The middleware handler measures the time taken for each API call and reports the metrics to the LocalAIMetricsService instance.  
- The code defines a function called LocalAIMetricsEndpoint which returns a handler for the Prometheus metrics endpoint. It uses the promhttp.Handler() function to handle the requests.  
  
- The code defines a struct called apiMiddlewareConfig which is used to configure the LocalAIMetricsAPIMiddleware middleware. It contains a Filter function and a metricsService field.  
- The Filter function checks if the request path is "/metrics". If it is, the middleware handler will allow the request to pass through.  
- The metricsService field is used to store the LocalAIMetricsService instance.  
- The code also defines a function called LocalAIMetricsAPIMiddleware which takes a LocalAIMetricsService instance as input and returns a middleware handler.  
- The middleware handler measures the time taken for each API call and reports the metrics to the LocalAIMetricsService instance.  
- The code defines a function called LocalAIMetricsEndpoint which returns a handler for the Prometheus metrics endpoint. It uses the promhttp.Handler() function to handle the requests.  
  
- The code defines a struct called apiMiddlewareConfig which is used to configure the LocalAIMetricsAPIMiddleware middleware. It contains a Filter function and a metricsService field.  
- The Filter function checks if the request path is "/metrics". If it is, the middleware handler will allow the request to pass through.  
- The metricsService field is used to store the LocalAIMetricsService instance.  
- The code also defines a function called LocalAIMetricsAPIMiddleware which takes a LocalAIMetricsService instance as input and returns a middleware handler.  
- The middleware handler measures the time taken for each API call and reports the metrics to the LocalAIMetricsService instance.  
- The code defines a function called LocalAIMetricsEndpoint which returns a handler for the Prometheus metrics endpoint. It uses the promhttp.Handler() function to handle the requests.  
  
- The code defines a struct called apiMiddlewareConfig which is used to configure the LocalAIMetricsAPIMiddleware middleware. It contains a Filter function and a metricsService field.  
- The Filter function checks if the request path is "/metrics". If it is, the middleware handler will allow the request to pass through.  
- The metricsService field is used to store the LocalAIMetricsService instance.  
- The code also defines a function called LocalAIMetricsAPIMiddleware which takes a LocalAIMetricsService instance as input and returns a middleware handler.  
- The middleware handler measures the time taken for each API call and reports the metrics to the LocalAIMetricsService instance.  
- The code defines a function called LocalAIMetricsEndpoint which returns a handler for the Prometheus metrics endpoint. It uses the promhttp.Handler() function to handle the requests.  
  
- The code defines a struct called apiMiddlewareConfig which is used to configure the LocalAIMetricsAPIMiddleware middleware. It contains a Filter function and a metricsService field.  
- The Filter function checks if the request path is "/metrics". If it is, the middleware handler will allow the request to pass through.  
- The metricsService field is used to store the LocalAIMetricsService instance.  
- The code also defines a function called LocalAIMetricsAPIMiddleware which takes a LocalAIMetricsService instance as input and returns a middleware handler.  
- The middleware handler measures the time taken for each API call and reports the metrics to the LocalAIMetricsService instance.  
- The code defines a function called LocalAIMetricsEndpoint which returns a handler for the Prometheus metrics endpoint. It uses the promhttp.Handler() function to handle the requests.  
  
- The code defines a struct called apiMiddlewareConfig which is used to configure the LocalAIMetricsAPIMiddleware middleware. It contains a Filter function and a metricsService field.  
- The Filter function checks if the request path is "/metrics". If it is, the middleware handler will allow the request to pass through.  
- The metricsService field is used to store the LocalAIMetricsService instance.  
- The code also defines a function called LocalAIMetricsAPIMiddleware which takes a LocalAIMetricsService instance as input and returns a middleware handler.  
- The middleware handler measures the time taken for each API call and reports the metrics to the LocalAIMetricsService instance.  
- The code defines a function called LocalAIMetricsEndpoint which returns a handler for the Prometheus metrics endpoint. It uses the promhttp.Handler() function to handle the requests.  
  
- The code defines a struct called apiMiddlewareConfig which is used to configure the LocalAIMetricsAPIMiddleware middleware. It contains a Filter function and a metricsService field.  
- The Filter function checks if the request path is "/metrics". If it is, the middleware handler will allow the request to pass through.  
- The metricsService field is used to store the LocalAIMetricsService instance.  
- The code also defines a function called LocalAIMetricsAPIMiddleware which takes a LocalAIMetricsService instance as input and returns a middleware handler.  
- The middleware handler measures the time taken for each API call and reports the metrics to the LocalAIMetricsService instance.  
- The code defines a function called LocalAIMetricsEndpoint which returns a handler for the Prometheus metrics endpoint. It uses the promhttp.Handler() function to handle the requests.  
  
- The code defines a struct called apiMiddlewareConfig which is used to configure the LocalAIMetricsAPIMiddleware middleware. It contains a Filter function and a metricsService field.  
- The Filter function checks if the request path is "/metrics". If it is, the middleware handler will allow the request to pass through.  
- The metricsService field is used to store the LocalAIMetricsService instance.  
- The code also defines a function called LocalAIMetricsAPIMiddleware which takes a LocalAIMetricsService instance as input and returns a middleware handler.  
- The middleware handler measures the time taken for each API call and reports the metrics to the LocalAIMetricsService instance.  
- The code defines a function called LocalAIMetricsEndpoint which returns a handler for the Prometheus metrics endpoint. It uses the promhttp.Handler() function to handle the requests.  
  
- The code defines a struct called apiMiddlewareConfig which is used to configure the LocalAIMetricsAPIMiddleware middleware. It contains a Filter function and a metricsService field.  
- The Filter function checks if the request path is "/metrics". If it is, the middleware handler will allow the request to pass through.  
- The metricsService field is used to store the LocalAIMetricsService instance.  
- The code also defines a function called LocalAIMetricsAPIMiddleware which takes a LocalAIMetricsService instance as input and returns a middleware handler.  
- The middleware handler measures the time taken for each API call and reports the metrics to the LocalAIMetricsService instance.  
- The code defines a function called LocalAIMetricsEndpoint which returns a handler for the Prometheus metrics endpoint. It uses the promhttp.Handler() function to handle the requests.  
  
- The code defines a struct called apiMiddlewareConfig which is used to configure the LocalAIMetricsAPIMiddleware middleware. It contains a Filter function and a metricsService field.  
- The Filter function checks if the request path is "/metrics". If it is, the middleware handler will allow the request to pass through.  
- The metricsService field is used to store the LocalAIMetricsService instance.  
- The code also defines a function called LocalAIMetricsAPIMiddleware which takes a LocalAIMetricsService instance as input and returns a middleware handler.  
- The middleware handler measures the time taken for each API call and reports the metrics to the LocalAIMetricsService instance.  
- The code defines a function called LocalAIMetricsEndpoint which returns a handler for the Prometheus metrics endpoint. It uses the promhttp.Handler() function to handle the requests.  
  
- The code defines a struct called apiMiddlewareConfig which is used to configure the LocalAIMetricsAPIMiddleware middleware. It contains a Filter function and a metricsService field.  
- The Filter function checks if the request path is "/metrics". If it is, the middleware handler will allow the request to pass through.  
- The metricsService field is used to store the LocalAIMetricsService instance.  
- The code also defines a function called LocalAIMetricsAPIMiddleware which takes a LocalAIMetricsService instance as input and returns a middleware handler.  
- The middleware handler measures the time taken for each API call and reports the metrics to the LocalAIMetricsService instance.  
- The code defines a function called LocalAIMetricsEndpoint which returns a handler for the Prometheus metrics endpoint. It uses the promhttp.Handler() function to handle the requests.  
  
- The code defines a struct called apiMiddlewareConfig which is used to configure the LocalAIMetricsAPIMiddleware middleware. It contains a Filter function and a metricsService field.  
- The Filter function checks if the request path is "/metrics". If it is, the middleware handler will allow the request to pass through.  
- The metricsService field is used to store the LocalAIMetricsService instance.  
- The code also defines a function called LocalAIMetricsAPIMiddleware which takes a LocalAIMetricsService instance as input and returns a middleware handler.  
- The middleware handler measures the time taken for each API call and reports the metrics to the LocalAIMetricsService instance.  
- The code defines a function called LocalAIMetricsEndpoint which returns a handler for the Prometheus metrics endpoint. It uses the promhttp.Handler() function to handle the requests.  
  
- The code defines a struct called apiMiddlewareConfig which is used to configure the LocalAIMetricsAPIMiddleware middleware. It contains a Filter function and a metricsService field.  
- The Filter function checks if the request path is "/metrics". If it is, the middleware handler will allow the request to pass through.  
- The metricsService field is used to store the LocalAIMetricsService instance.  
- The code also defines a function called LocalAIMetricsAPIMiddleware which takes a LocalAIMetricsService instance as input and returns a middleware handler.  
- The middleware handler measures the time taken for each API call and reports the metrics to the LocalAIMetricsService instance.  
- The code defines a function called LocalAIMetricsEndpoint which returns a handler for the Prometheus metrics endpoint. It uses the promhttp.Handler() function to handle the requests.  
  
- The code defines a struct called apiMiddlewareConfig which is used to configure the LocalAIMetricsAPIMiddleware middleware. It contains a Filter function and a metricsService field.  
- The Filter function checks if the request path is "/metrics". If it is, the middleware handler will allow the request to pass through.  
- The metricsService field is used to store the LocalAIMetricsService instance.  
- The code also defines a function called LocalAIMetricsAPIMiddleware which takes a LocalAIMetricsService instance as input and returns a middleware handler.  
- The middleware handler measures the time taken for each API call and reports the metrics to the LocalAIMetricsService instance.  
- The code defines a function called LocalAIMetricsEndpoint which returns a handler for the Prometheus metrics endpoint. It uses the promhttp.Handler() function to handle the requests.  
  
- The code defines a struct called apiMiddlewareConfig which is used to configure the LocalAIMetricsAPIMiddleware middleware. It contains a Filter function and a metricsService field.  
- The Filter function checks if the request path is "/metrics". If it is, the middleware handler will allow the request to pass through.  
- The metricsService field is used to store the LocalAIMetricsService instance.  
- The code also defines a function called LocalAIMetricsAPIMiddleware which takes a LocalAIMetricsService instance as input and returns a middleware handler.  
- The middleware handler measures the time taken for each API call and reports the metrics to the LocalAIMetricsService instance.  
- The code defines a function called LocalAIMetricsEndpoint which returns a handler for the Prometheus metrics endpoint. It uses the promhttp.Handler() function to handle the requests.  
  
- The code defines a struct called apiMiddlewareConfig which is used to configure the LocalAIMetricsAPIMiddleware middleware. It contains a Filter function and a metricsService field.  
- The Filter function checks if the request path is "/metrics". If it is, the middleware handler will allow the request to pass through.  
- The metricsService field is used to store the LocalAIMetricsService instance.  
- The code also defines a function called LocalAIMetricsAPIMiddleware which takes a LocalAIMetricsService instance as input and returns a middleware handler.  
- The middleware handler measures the time taken for each API call and reports the metrics to the LocalAIMetricsService instance.  
- The code defines a function called LocalAIMetricsEndpoint which returns a handler for the Prometheus metrics endpoint. It uses the promhttp.Handler() function to handle the requests.  
  
- The code defines a struct called apiMiddlewareConfig which is used to configure the LocalAIMetricsAPIMiddleware middleware. It contains a Filter function and a metricsService field.  
- The Filter function checks if the request path is "/metrics". If it is, the middleware handler will allow the request to pass through.  
- The metricsService field is  
  
# core/http/endpoints/localai/p2p.go  
Package: localai  
  
Imports:  
- "github.com/gofiber/fiber/v2"  
- "github.com/mudler/LocalAI/core/config"  
- "github.com/mudler/LocalAI/core/p2p"  
- "github.com/mudler/LocalAI/core/schema"  
  
External data, input sources:  
- The code uses the `appConfig` parameter, which is an instance of `config.ApplicationConfig`, to access the P2P network ID and token.  
  
TODOs:  
- There are no TODO comments in the provided code.  
  
Summary of major code parts:  
- The `ShowP2PNodes` function returns a handler that, when called, returns a JSON response containing a list of available P2P nodes. The nodes are categorized as either workers or federated nodes based on their network ID.  
- The `ShowP2PToken` function returns a handler that, when called, returns the P2P token as a string.  
  
  
  
# core/http/endpoints/localai/stores.go  
Package: localai  
  
Imports:  
- github.com/gofiber/fiber/v2  
- github.com/mudler/LocalAI/core/backend  
- github.com/mudler/LocalAI/core/config  
- github.com/mudler/LocalAI/core/schema  
- github.com/mudler/LocalAI/pkg/model  
- github.com/mudler/LocalAI/pkg/store  
  
External data, input sources:  
- The code interacts with a store backend, which is provided by the backend package.  
- It receives input from the client through the fiber framework.  
  
TODOs:  
- There are no TODO comments in the code.  
  
Summary:  
- The code defines four endpoints for interacting with a store backend: StoresSetEndpoint, StoresDeleteEndpoint, StoresGetEndpoint, and StoresFindEndpoint.  
- StoresSetEndpoint sets values in the store backend.  
- StoresDeleteEndpoint deletes values from the store backend.  
- StoresGetEndpoint retrieves values from the store backend.  
- StoresFindEndpoint finds similar values in the store backend.  
- Each endpoint takes a model loader and application config as input, and returns a function that handles the request.  
- The function parses the request body, interacts with the store backend, and returns a response.  
- The code uses the store package to interact with the store backend.  
- The code uses the schema package to define the input and output data structures for the endpoints.  
- The code uses the fiber framework to handle the HTTP requests.  
  
  
  
# core/http/endpoints/localai/system.go  
## Package: localai  
  
Imports:  
- "github.com/gofiber/fiber/v2"  
- "github.com/mudler/LocalAI/core/config"  
- "github.com/mudler/LocalAI/core/schema"  
- "github.com/mudler/LocalAI/pkg/model"  
  
External data, input sources:  
- `ml *model.ModelLoader`: A pointer to a ModelLoader object, which is responsible for loading and managing machine learning models.  
- `appConfig *config.ApplicationConfig`: A pointer to an ApplicationConfig object, which contains the configuration settings for the LocalAI instance.  
  
TODOs:  
- None found.  
  
### SystemInformations function  
This function returns a handler function that provides information about the LocalAI instance, including the available backends and loaded models. It takes two arguments: a pointer to a ModelLoader object and a pointer to an ApplicationConfig object. The handler function first retrieves the list of available backends from the ModelLoader object, then iterates through the loaded models and appends their IDs to a list of SysInfoModel objects. Finally, it returns a JSON response containing the available backends and loaded models.  
  
  
  
# core/http/endpoints/localai/tokenize.go  
## Package: localai  
  
Imports:  
- "github.com/gofiber/fiber/v2"  
- "github.com/mudler/LocalAI/core/backend"  
- "github.com/mudler/LocalAI/core/config"  
- "github.com/mudler/LocalAI/core/http/middleware"  
- "github.com/mudler/LocalAI/core/schema"  
- "github.com/mudler/LocalAI/pkg/model"  
  
External data, input sources:  
- The code uses the `schema.TokenizeRequest` struct to receive the input data for tokenization.  
- It also relies on the `config.BackendConfigLoader`, `model.ModelLoader`, and `config.ApplicationConfig` to access model configurations and application settings.  
  
TODOs:  
- There are no TODO comments in the provided code.  
  
### Tokenization Endpoint  
  
This section describes the `TokenizeEndpoint` function, which handles the tokenization request. It takes three arguments:  
- `cl`: A pointer to the `config.BackendConfigLoader` object, which is used to load backend configurations.  
- `ml`: A pointer to the `model.ModelLoader` object, which is used to load models.  
- `appConfig`: A pointer to the `config.ApplicationConfig` object, which contains application-wide settings.  
  
The function returns a function that takes a Fiber context as input and returns an error. This inner function handles the actual tokenization process:  
1. It retrieves the `schema.TokenizeRequest` object from the context locals.  
2. It checks if the request contains a model name and if the model configuration is available.  
3. It calls the `backend.ModelTokenize` function to perform the tokenization using the provided input, model loader, backend configuration, and application configuration.  
4. It returns the tokenization response as a JSON object.  
  
This endpoint is designed to be used by clients to tokenize their input content. The tokenization process is handled by the backend, which uses the provided model and configuration to generate the tokens. The response is then sent back to the client as a JSON object.  
  
# core/http/endpoints/localai/tts.go  
Package: localai  
  
Imports:  
- github.com/mudler/LocalAI/core/backend  
- github.com/mudler/LocalAI/core/config  
- github.com/mudler/LocalAI/core/http/middleware  
- github.com/mudler/LocalAI/pkg/model  
- github.com/gofiber/fiber/v2  
- github.com/mudler/LocalAI/core/schema  
- github.com/rs/zerolog/log  
- github.com/mudler/LocalAI/pkg/utils  
  
External data, input sources:  
- The code uses the schema.TTSRequest struct to receive the input text and other parameters for the TTS endpoint.  
- It also uses the config.BackendConfig struct to access the backend configuration.  
- The code relies on the model.ModelLoader to load the necessary model for text-to-speech conversion.  
  
TODOs:  
- There are no TODO comments in the provided code.  
  
Code summary:  
- The TTSEndpoint function handles the text-to-speech conversion endpoint. It receives the input text and other parameters, then uses the backend.ModelTTS function to generate the audio file. The generated file is then converted to the desired format using the utils.AudioConvert function. Finally, the audio file is sent back to the client as a download.  
  
  
  
# core/http/endpoints/localai/vad.go  
## Package: localai  
  
Imports:  
- github.com/gofiber/fiber/v2  
- github.com/mudler/LocalAI/core/backend  
- github.com/mudler/LocalAI/core/config  
- github.com/mudler/LocalAI/core/http/middleware  
- github.com/mudler/LocalAI/core/schema  
- github.com/mudler/LocalAI/pkg/model  
- github.com/rs/zerolog/log  
  
External data, input sources:  
- The code uses the `schema.VADRequest` struct to receive voice activation detection requests.  
- It also uses the `config.BackendConfigLoader`, `model.ModelLoader`, and `config.ApplicationConfig` to access backend configuration, model information, and application settings.  
  
TODOs:  
- There are no TODO comments in the provided code.  
  
### VADEndpoint function  
This function handles voice activation detection requests. It receives a `schema.VADRequest` object containing the audio stream and model to use. The function then uses the provided model loader and backend configuration to process the request and returns a `proto.VADResponse` object containing the results.  
  
### Summary  
The code defines a VADEndpoint function that handles voice activation detection requests. It receives the request, extracts the necessary information, and uses the backend and model configuration to process the request. The function then returns the results to the client.  
  
# core/http/endpoints/localai/welcome.go  
## Package: localai  
  
Imports:  
- github.com/gofiber/fiber/v2  
- github.com/mudler/LocalAI/core/config  
- github.com/mudler/LocalAI/core/gallery  
- github.com/mudler/LocalAI/core/http/utils  
- github.com/mudler/LocalAI/core/p2p  
- github.com/mudler/LocalAI/core/services  
- github.com/mudler/LocalAI/internal  
- github.com/mudler/LocalAI/pkg/model  
  
External data and input sources:  
- Backend configurations from the config package  
- Model configurations from the gallery package  
- Model statuses from the services package  
  
TODOs:  
- None found  
  
### WelcomeEndpoint function  
This function handles the welcome endpoint for the LocalAI API. It retrieves backend configurations, gallery configurations, and model statuses. It then creates a summary containing this information and the application configuration. Finally, it returns the summary as a JSON response or renders an HTML template depending on the client's request.  
  
  
  
