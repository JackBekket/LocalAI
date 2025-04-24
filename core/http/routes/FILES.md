# core/http/routes/elevenlabs.go  
Package: routes  
  
Imports:  
- "github.com/gofiber/fiber/v2"  
- "github.com/mudler/LocalAI/core/config"  
- "github.com/mudler/LocalAI/core/http/endpoints/elevenlabs"  
- "github.com/mudler/LocalAI/core/http/middleware"  
- "github.com/mudler/LocalAI/core/schema"  
- "github.com/mudler/LocalAI/pkg/model"  
  
External data, input sources:  
- The code uses the `config.BackendConfigLoader` to load backend configuration.  
- It uses the `model.ModelLoader` to load models.  
- It uses the `config.ApplicationConfig` to access application-wide configuration.  
  
TODOs:  
- There are no TODO comments in the code.  
  
Summary:  
- The `RegisterElevenLabsRoutes` function registers routes for ElevenLabs endpoints.  
- It sets up two endpoints: one for text-to-speech and one for sound generation.  
- For each endpoint, it uses middleware to filter and select the appropriate model based on the use case.  
- It then calls the corresponding endpoint handler function from the `elevenlabs` package.  
  
  
  
# core/http/routes/explorer.go  
Package/Component name: routes  
  
Imports:  
- "github.com/gofiber/fiber/v2"  
- "github.com/mudler/LocalAI/core/explorer"  
- "github.com/mudler/LocalAI/core/http/endpoints/explorer"  
  
External data, input sources:  
- The code interacts with a database instance of type coreExplorer.Database, which is passed as an argument to the RegisterExplorerRoutes function.  
  
TODOs:  
- There are no TODO comments in the provided code.  
  
Summary of major code parts:  
- The RegisterExplorerRoutes function registers routes for the explorer component. It takes an instance of a fiber.App and a coreExplorer.Database as input.  
- The function registers three routes:  
    - A GET route at "/" that returns the explorer dashboard.  
    - A POST route at "/network/add" that adds a new network to the database.  
    - A GET route at "/networks" that shows a list of networks in the database.  
  
  
  
# core/http/routes/health.go  
Package/Component name: routes  
  
Imports:  
github.com/gofiber/fiber/v2  
  
External data, input sources:  
None  
  
TODOs:  
None  
  
Summary:  
The `HealthRoutes` function in the `routes` package is responsible for setting up health check endpoints for a service. It defines two endpoints, `/healthz` and `/readyz`, which return a 200 OK status code when accessed. These endpoints can be used by monitoring systems to check the health and readiness of the service.  
  
The function takes an instance of the Fiber web framework's `App` struct as input, which represents the application's router. It then registers two GET routes, one for each health check endpoint, using the `app.Get` method. The `ok` function is used as the handler for both endpoints, which simply returns a 200 OK status code to the client.  
  
By providing these health check endpoints, the `HealthRoutes` function enables monitoring systems to easily determine the status of the service and ensure that it is functioning correctly.  
  
  
  
# core/http/routes/jina.go  
Package: routes  
  
Imports:  
- "github.com/mudler/LocalAI/core/config"  
- "github.com/mudler/LocalAI/core/http/endpoints/jina"  
- "github.com/mudler/LocalAI/core/http/middleware"  
- "github.com/mudler/LocalAI/core/schema"  
- "github.com/gofiber/fiber/v2"  
- "github.com/mudler/LocalAI/pkg/model"  
  
External data, input sources:  
- The code uses the `config.BackendConfigLoader` to load backend configuration.  
- It uses the `model.ModelLoader` to load models.  
- It uses the `config.ApplicationConfig` to access application-wide configuration.  
  
TODOs:  
- There are no TODO comments in the code.  
  
Summary:  
The `RegisterJINARoutes` function registers a POST endpoint for reranking requests. It uses the provided `RequestExtractor`, `BackendConfigLoader`, `ModelLoader`, and `ApplicationConfig` to handle the request and perform the reranking operation. The endpoint is designed to mimic the behavior of the JINA framework.  
  
  
  
# core/http/routes/localai.go  
## Package/Component: routes  
  
### Imports:  
- "github.com/gofiber/fiber/v2"  
- "github.com/gofiber/swagger"  
- "github.com/mudler/LocalAI/core/config"  
- "github.com/mudler/LocalAI/core/http/endpoints/localai"  
- "github.com/mudler/LocalAI/core/http/middleware"  
- "github.com/mudler/LocalAI/core/p2p"  
- "github.com/mudler/LocalAI/core/schema"  
- "github.com/mudler/LocalAI/core/services"  
- "github.com/mudler/LocalAI/internal"  
- "github.com/mudler/LocalAI/pkg/model"  
  
### External Data and Input Sources:  
- The code uses a `config.BackendConfigLoader` to load backend configuration.  
- It uses a `model.ModelLoader` to load models.  
- It uses a `config.ApplicationConfig` to access application-wide configuration.  
- It uses a `services.GalleryService` to interact with the gallery service.  
  
### TODOs:  
- TODO: Should these use standard middlewares? Refactor later, they are extremely simple.  
  
### Summary:  
This file contains the implementation for registering various API endpoints related to LocalAI. It handles endpoints for model management, voice activity detection, text-to-speech, and other functionalities.  
  
1. **Swagger Documentation:**  
   - The code registers a swagger endpoint for API documentation.  
  
2. **Model Gallery Endpoints:**  
   - If the gallery endpoint is not disabled, it registers endpoints for managing model galleries, including adding, deleting, and listing models and galleries.  
   - It also provides endpoints for retrieving the status of model gallery operations.  
  
3. **TTS Endpoint:**  
   - It registers an endpoint for text-to-speech functionality, using the first available default model for the TTS use case.  
  
4. **VAD Endpoint:**  
   - It registers an endpoint for voice activity detection, using the first available default model for the VAD use case.  
  
5. **Stores Endpoints:**  
   - It registers endpoints for managing stores, including setting, deleting, getting, and finding stores.  
  
6. **Metrics Endpoint:**  
   - If the metrics endpoint is not disabled, it registers an endpoint for retrieving backend statistics.  
  
7. **Backend Monitor and Shutdown Endpoints:**  
   - It registers endpoints for monitoring the backend and shutting it down.  
  
8. **P2P Endpoints:**  
   - If P2P is enabled, it registers endpoints for showing P2P nodes and tokens.  
  
9. **Version and System Information Endpoints:**  
   - It registers endpoints for retrieving the version and system information.  
  
10. **Tokenize Endpoint:**  
   - It registers an endpoint for tokenization, using the first available default model for the tokenize use case.  
  
This file provides a comprehensive set of API endpoints for interacting with the LocalAI system, covering various functionalities and configurations.  
  
# core/http/routes/openai.go  
Package/Component name: routes  
  
Imports:  
- "github.com/gofiber/fiber/v2"  
- "github.com/mudler/LocalAI/core/application"  
- "github.com/mudler/LocalAI/core/config"  
- "github.com/mudler/LocalAI/core/http/endpoints/localai"  
- "github.com/mudler/LocalAI/core/http/endpoints/openai"  
- "github.com/mudler/LocalAI/core/http/middleware"  
- "github.com/mudler/LocalAI/core/schema"  
  
External data, input sources:  
- The code uses the application, BackendLoader, ModelLoader, TemplatesEvaluator, and ApplicationConfig from the application package.  
- It also uses the RequestExtractor from the middleware package.  
  
TODOs:  
- There are no TODO comments in the code.  
  
Summary:  
- The code defines a set of routes for the openAI compatible API endpoint.  
- It includes endpoints for chat, edit, completion, assistant, files, embeddings, audio, images, and list models.  
- Each endpoint is associated with a specific handler chain that performs the necessary operations, such as filtering models, setting configurations, and executing the corresponding endpoint function.  
- The code uses the RequestExtractor to extract the request and set the appropriate model and configuration for each endpoint.  
- The endpoints are registered with the fiber app, which handles the incoming requests and routes them to the appropriate handlers.  
  
- The chat endpoint handles chat completions, allowing users to interact with the AI model in a conversational manner.  
- The edit endpoint enables users to edit existing text using the AI model.  
- The completion endpoint provides completions for given text prompts.  
- The assistant endpoint allows users to manage and interact with AI assistants.  
- The files endpoint handles file uploads, downloads, and management.  
- The embeddings endpoint generates embeddings for given text inputs.  
- The audio endpoint handles audio transcriptions and text-to-speech conversions.  
- The images endpoint generates images based on text prompts.  
- The list models endpoint provides a list of available AI models.  
  
- The code ensures that the appropriate models and configurations are used for each endpoint, based on the user's request and the application's settings.  
- It also handles any necessary preprocessing and postprocessing of the input and output data.  
  
- By providing a comprehensive set of endpoints, the code enables users to interact with the AI models in a variety of ways, making it a versatile and powerful tool for various applications.  
  
# core/http/routes/ui.go  
```json  
{  
  "package/component": "routes",  
  "imports": [  
    "fmt",  
    "html/template",  
    "math",  
    "sort",  
    "strconv",  
    "strings",  
    "github.com/mudler/LocalAI/core/config",  
    "github.com/mudler/LocalAI/core/gallery",  
    "github.com/mudler/LocalAI/core/http/elements",  
    "github.com/mudler/LocalAI/core/http/endpoints/localai",  
    "github.com/mudler/LocalAI/core/http/utils",  
    "github.com/mudler/LocalAI/core/p2p",  
    "github.com/mudler/LocalAI/core/services",  
    "github.com/mudler/LocalAI/internal",  
    "github.com/mudler/LocalAI/pkg/model",  
    "github.com/mudler/LocalAI/pkg/xsync",  
    "github.com/gofiber/fiber/v2",  
    "github.com/google/uuid",  
    "github.com/microcosm-cc/bluemonday",  
    "github.com/rs/zerolog/log"  
  ],  
  "external_data": [],  
  "input_sources": [],  
  "TODOs": [],  
  "summary": {  
    "modelOpCache": "This struct is used to store the status of models that are being installed or deleted from the UI. It uses a map to store the status of each model, and provides methods to set, get, and delete the status of models.",  
    "RegisterUIRoutes": "This function registers the UI routes for the LocalAI application. It takes the fiber app, config loader, model loader, application config, and gallery service as input. It sets up routes for the welcome page, p2p dashboard, models page, chat page, text2image page, and tts page. It also handles the installation and deletion of models, and provides progress updates to the UI."  
  },  
  "end_of_output": ""  
}  
```  
  
