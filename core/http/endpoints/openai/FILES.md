# core/http/endpoints/openai/assistant.go  
## Package: openai  
  
This package provides functionalities related to the OpenAI Assistant API. It includes endpoints for creating, listing, deleting, and modifying assistants, as well as managing files associated with them.  
  
### External Data and Input Sources:  
  
- The package relies on the `config` package to load backend and model configurations.  
- It uses the `schema` package to define data structures for API responses.  
- The `services` package is used to interact with the backend and retrieve model information.  
- The `model` package provides access to model-related data.  
- The `utils` package contains utility functions for saving configurations.  
  
### TODOs:  
  
- Implement proper error handling for API requests.  
- Add support for pagination in list endpoints.  
- Improve documentation and code comments.  
  
### Summary of Major Code Parts:  
  
1. **Data Structures:**  
   - `ToolType`: Defines the type of tool for an assistant.  
   - `Assistant`: Represents the structure of an assistant object.  
   - `AssistantRequest`: Contains the request data for creating an assistant.  
   - `AssistantFile`: Represents a file associated with an assistant.  
  
2. **Endpoints:**  
   - `CreateAssistantEndpoint`: Creates a new assistant with a model and instructions.  
   - `ListAssistantsEndpoint`: Lists available assistants.  
   - `DeleteAssistantEndpoint`: Deletes an assistant.  
   - `GetAssistantEndpoint`: Retrieves an assistant's data.  
   - `CreateAssistantFileEndpoint`: Creates a new file for an assistant.  
   - `ListAssistantFilesEndpoint`: Lists files associated with an assistant.  
   - `ModifyAssistantEndpoint`: Modifies an existing assistant.  
   - `DeleteAssistantFileEndpoint`: Deletes a file associated with an assistant.  
   - `GetAssistantFileEndpoint`: Retrieves a file's data.  
  
3. **Utility Functions:**  
   - `modelExists`: Checks if a model exists in the configuration.  
   - `filterAssistantsAfterID`: Filters assistants with IDs after a given ID.  
   - `filterAssistantsBeforeID`: Filters assistants with IDs before a given ID.  
  
4. **Configuration Management:**  
   - `Assistants`: A slice to store assistant objects.  
   - `AssistantsConfigFile`: The configuration file for assistants.  
   - `AssistantFiles`: A slice to store assistant file objects.  
   - `AssistantsFileConfigFile`: The configuration file for assistant files.  
  
This package provides a comprehensive set of functionalities for managing OpenAI Assistants and their associated files. It covers the creation, listing, modification, and deletion of assistants, as well as the management of files attached to them.  
  
# core/http/endpoints/openai/assistant_test.go  
## Package: openai  
  
This package provides functionalities related to the openai API.  
  
### External Data and Input Sources:  
  
- The code interacts with the openai API to perform various tasks.  
- It uses the `fiber` web framework to handle HTTP requests and responses.  
- It relies on the `schema` package for data structures related to assistants, files, and other related entities.  
  
### TODOs:  
  
- There are no explicit TODO comments in the code.  
  
### Code Summary:  
  
1. **Test Setup:**  
   - The code sets up a test environment by creating a mock backend configuration loader, a model loader, and an application configuration.  
   - It also creates a test server using the `fiber` web framework.  
  
2. **Assistant Endpoints:**  
   - The code tests various endpoints related to assistants, such as creating, listing, deleting, and modifying assistants.  
   - It also tests endpoints for managing assistant files, such as uploading, listing, deleting, and retrieving files associated with assistants.  
  
3. **File Endpoints:**  
   - The code tests endpoints related to files, such as uploading, deleting, and retrieving files.  
  
4. **Cleanup:**  
   - The code includes cleanup functions to remove any created assistants, files, and other resources after the tests are completed.  
  
5. **Helper Functions:**  
   - The code includes helper functions to create assistants, files, and other related entities.  
   - It also includes functions to clean up any created resources.  
  
6. **Test Execution:**  
   - The code executes the tests and asserts that the expected results are obtained.  
  
### Summary:  
  
The code tests various endpoints related to assistants and files in the openai API. It uses the `fiber` web framework to handle HTTP requests and responses, and it relies on the `schema` package for data structures related to assistants, files, and other related entities. The code includes cleanup functions to remove any created resources after the tests are completed.  
  
# core/http/endpoints/openai/chat.go  
## Package: openai  
  
This package provides an endpoint for the OpenAI Completion API. It handles chat completions for a given prompt and model.  
  
### External Data and Input Sources:  
  
- The package relies on the `schema.OpenAIRequest` struct to receive input parameters, such as the model, prompt, and any additional context.  
- It uses the `config.BackendConfig` struct to access backend configuration settings, including grammar and function definitions.  
- The `model.ModelLoader` is used to load and interact with the language model.  
  
### TODOs:  
  
- Change generated BNF grammar to be compliant with the schema so that the result token by token can be streamed.  
  
### Summary of Major Code Parts:  
  
1. **ChatEndpoint Function:** This function is the main entry point for handling chat completion requests. It takes the backend configuration, model loader, evaluator, and startup options as input.  
  
2. **Process Function:** This function is responsible for processing the input prompt and generating a response using the language model. It handles both streaming and non-streaming modes.  
  
3. **HandleQuestion Function:** This function is called when the response from the language model needs to be processed further, such as when functions are involved. It takes the backend configuration, input request, model loader, startup options, function results, and the prompt as input.  
  
4. **ComputeChoices Function:** This function is used to compute the choices for the chat completion response. It takes the input request, prompt, backend configuration, startup options, model loader, and two callback functions as input.  
  
5. **TemplateMessages Function:** This function is used to process the input messages and prepare them for the language model. It takes the input messages, backend configuration, function definitions, and a flag indicating whether functions should be used as input.  
  
6. **Finetune Function:** This function is used to fine-tune the response from the language model based on the prompt and context. It takes the backend configuration, prompt, and response as input.  
  
7. Stream Response Handling: The code includes logic for handling streaming responses, where the client can receive chunks of the response as they are generated.  
  
8. Function Handling: The code supports handling functions, allowing users to define custom actions that can be executed by the language model.  
  
9. Grammar and Schema Support: The code includes support for grammars and schemas, allowing users to define the structure of the input and output data.  
  
10. Error Handling: The code includes error handling mechanisms to ensure that the endpoint can handle unexpected situations gracefully.  
  
By providing a comprehensive set of features and functionalities, this package enables users to interact with the OpenAI Completion API and leverage the power of large language models for various applications.  
  
# core/http/endpoints/openai/completion.go  
## Package: openai  
  
This package provides the OpenAI Completion API endpoint.  
  
### External Data and Input Sources:  
  
- The code relies on the `schema.OpenAIRequest` struct to receive the user's request.  
- It uses the `config.BackendConfigLoader`, `model.ModelLoader`, and `templates.Evaluator` to access backend configurations, load models, and evaluate templates, respectively.  
- The `config.ApplicationConfig` is used to access application-level configurations.  
  
### TODOs:  
  
- There are no TODO comments in the provided code.  
  
### Code Summary:  
  
1. **CompletionEndpoint Function:**  
This function handles the OpenAI Completion API endpoint. It takes the backend configuration, model loader, evaluator, and application config as input.  
- It processes the incoming request, extracts the model and prompt, and prepares the input for the completion process.  
- It uses the `ComputeChoices` function to generate completions based on the input and configuration.  
- It formats the response according to the OpenAI specification and returns it to the client.  
- If the request specifies streaming, it sends the response in chunks as the completions are generated.  
  
2. **ComputeChoices Function:**  
This function is responsible for generating completions based on the input prompt and model. It takes the request, prompt, configuration, application config, model loader, and callback functions as input.  
- It uses the provided model loader to load the specified model.  
- It processes the prompt and generates completions using the loaded model.  
- It calls the provided callback functions to handle the generated completions and token usage.  
  
3. **Template Evaluation:**  
The code uses the `templates.Evaluator` to evaluate templates for prompt processing. It takes the prompt, system prompt, and configuration as input and returns the modified prompt.  
  
4. **Response Formatting:**  
The code formats the response according to the OpenAI specification, including the ID, created timestamp, model, choices, object, and usage information.  
  
5. **Token Usage Tracking:**  
The code tracks token usage during the completion process and includes it in the response.  
  
6. **Error Handling:**  
The code includes error handling for cases such as invalid input or errors during the completion process.  
  
# core/http/endpoints/openai/edit.go  
Package: openai  
  
Imports:  
	"encoding/json"  
	"time"  
	"github.com/mudler/LocalAI/core/backend"  
	"github.com/mudler/LocalAI/core/config"  
	"github.com/mudler/LocalAI/core/http/middleware"  
	"github.com/gofiber/fiber/v2"  
	"github.com/google/uuid"  
	"github.com/mudler/LocalAI/core/schema"  
	"github.com/mudler/LocalAI/pkg/model"  
	"github.com/mudler/LocalAI/pkg/templates"  
	"github.com/rs/zerolog/log"  
  
External data, input sources:  
	- The code uses the schema.OpenAIRequest struct to receive input from the user.  
	- It also uses the config.BackendConfig struct to access backend configuration.  
  
TODOs:  
	- There are no TODO comments in the code.  
  
Code summary:  
	- The EditEndpoint function handles the OpenAI edit API endpoint. It takes the user's input, model, and configuration as parameters.  
	- The function first checks if the input is valid and if the model is specified.  
	- Then, it iterates through the input strings and uses the evaluator to generate a templated input for each string.  
	- For each templated input, it calls the ComputeChoices function to get the choices.  
	- The ComputeChoices function calculates the choices based on the input, configuration, and model.  
	- The function then calculates the total token usage and creates a response object containing the choices, usage, and other relevant information.  
	- Finally, the response object is marshalled into JSON format and returned to the user.  
  
	- The ComputeChoices function is responsible for calculating the choices based on the input, configuration, and model.  
	- It takes the input, input string, configuration, application configuration, model loader, and two callback functions as parameters.  
	- The function first checks if the input string is empty. If it is, it returns an empty slice of choices.  
	- Then, it calls the model's predict function to get the predictions.  
	- The function then iterates through the predictions and calls the callback function for each prediction.  
	- Finally, the function returns the slice of choices.  
  
	- The code also includes a function to evaluate the template for the prompt. This function takes the template, configuration, and prompt template data as parameters.  
	- It first checks if the template is empty. If it is, it returns the input string.  
	- Then, it uses the evaluator to evaluate the template and returns the result.  
  
	- The code also includes a function to compute the choices. This function takes the input, input string, configuration, application configuration, model loader, and two callback functions as parameters.  
	- It first checks if the input string is empty. If it is, it returns an empty slice of choices.  
	- Then, it calls the model's predict function to get the predictions.  
	- The function then iterates through the predictions and calls the callback function for each prediction.  
	- Finally, the function returns the slice of choices.  
  
	- The code also includes a function to compute the token usage. This function takes the input, input string, configuration, application configuration, model loader, and two callback functions as parameters.  
	- It first checks if the input string is empty. If it is, it returns an empty token usage object.  
	- Then, it calls the model's predict function to get the predictions.  
	- The function then iterates through the predictions and calls the callback function for each prediction.  
	- Finally, the function returns the token usage object.  
  
	- The code also includes a function to compute the timing token generation. This function takes the input, input string, configuration, application configuration, model loader, and two callback functions as parameters.  
	- It first checks if the input string is empty. If it is, it returns an empty timing token generation object.  
	- Then, it calls the model's predict function to get the predictions.  
	- The function then iterates through the predictions and calls the callback function for each prediction.  
	- Finally, the function returns the timing token generation object.  
  
	- The code also includes a function to compute the timing prompt processing. This function takes the input, input string, configuration, application configuration, model loader, and two callback functions as parameters.  
	- It first checks if the input string is empty. If it is, it returns an empty timing prompt processing object.  
	- Then, it calls the model's predict function to get the predictions.  
	- The function then iterates through the predictions and calls the callback function for each prediction.  
	- Finally, the function returns the timing prompt processing object.  
  
	- The code also includes a function to compute the prompt tokens. This function takes the input, input string, configuration, application configuration, model loader, and two callback functions as parameters.  
	- It first checks if the input string is empty. If it is, it returns an empty prompt tokens object.  
	- Then, it calls the model's predict function to get the predictions.  
	- The function then iterates through the predictions and calls the callback function for each prediction.  
	- Finally, the function returns the prompt tokens object.  
  
	- The code also includes a function to compute the completion tokens. This function takes the input, input string, configuration, application configuration, model loader, and two callback functions as parameters.  
	- It first checks if the input string is empty. If it is, it returns an empty completion tokens object.  
	- Then, it calls the model's predict function to get the predictions.  
	- The function then iterates through the predictions and calls the callback function for each prediction.  
	- Finally, the function returns the completion tokens object.  
  
	- The code also includes a function to compute the total tokens. This function takes the input, input string, configuration, application configuration, model loader, and two callback functions as parameters.  
	- It first checks if the input string is empty. If it is, it returns an empty total tokens object.  
	- Then, it calls the model's predict function to get the predictions.  
	- The function then iterates through the predictions and calls the callback function for each prediction.  
	- Finally, the function returns the total tokens object.  
  
	- The code also includes a function to compute the timing token generation. This function takes the input, input string, configuration, application configuration, model loader, and two callback functions as parameters.  
	- It first checks if the input string is empty. If it is, it returns an empty timing token generation object.  
	- Then, it calls the model's predict function to get the predictions.  
	- The function then iterates through the predictions and calls the callback function for each prediction.  
	- Finally, the function returns the timing token generation object.  
  
	- The code also includes a function to compute the timing prompt processing. This function takes the input, input string, configuration, application configuration, model loader, and two callback functions as parameters.  
	- It first checks if the input string is empty. If it is, it returns an empty timing prompt processing object.  
	- Then, it calls the model's predict function to get the predictions.  
	- The function then iterates through the predictions and calls the callback function for each prediction.  
	- Finally, the function returns the timing prompt processing object.  
  
	- The code also includes a function to compute the prompt tokens. This function takes the input, input string, configuration, application configuration, model loader, and two callback functions as parameters.  
	- It first checks if the input string is empty. If it is, it returns an empty prompt tokens object.  
	- Then, it calls the model's predict function to get the predictions.  
	- The function then iterates through the predictions and calls the callback function for each prediction.  
	- Finally, the function returns the prompt tokens object.  
  
	- The code also includes a function to compute the completion tokens. This function takes the input, input string, configuration, application configuration, model loader, and two callback functions as parameters.  
	- It first checks if the input string is empty. If it is, it returns an empty completion tokens object.  
	- Then, it calls the model's predict function to get the predictions.  
	- The function then iterates through the predictions and calls the callback function for each prediction.  
	- Finally, the function returns the completion tokens object.  
  
	- The code also includes a function to compute the total tokens. This function takes the input, input string, configuration, application configuration, model loader, and two callback functions as parameters.  
	- It first checks if the input string is empty. If it is, it returns an empty total tokens object.  
	- Then, it calls the model's predict function to get the predictions.  
	- The function then iterates through the predictions and calls the callback function for each prediction.  
	- Finally, the function returns the total tokens object.  
  
	- The code also includes a function to compute the timing token generation. This function takes the input, input string, configuration, application configuration, model loader, and two callback functions as parameters.  
	- It first checks if the input string is empty. If it is, it returns an empty timing token generation object.  
	- Then, it calls the model's predict function to get the predictions.  
	- The function then iterates through the predictions and calls the callback function for each prediction.  
	- Finally, the function returns the timing token generation object.  
  
	- The code also includes a function to compute the timing prompt processing. This function takes the input, input string, configuration, application configuration, model loader, and two callback functions as parameters.  
	- It first checks if the input string is empty. If it is, it returns an empty timing prompt processing object.  
	- Then, it calls the model's predict function to get the predictions.  
	- The function then iterates through the predictions and calls the callback function for each prediction.  
	- Finally, the function returns the timing prompt processing object.  
  
	- The code also includes a function to compute the prompt tokens. This function takes the input, input string, configuration, application configuration, model loader, and two callback functions as parameters.  
	- It first checks if the input string is empty. If it is, it returns an empty prompt tokens object.  
	- Then, it calls the model's predict function to get the predictions.  
	- The function then iterates through the predictions and calls the callback function for each prediction.  
	- Finally, the function returns the prompt tokens object.  
  
	- The code also includes a function to compute the completion tokens. This function takes the input, input string, configuration, application configuration, model loader, and two callback functions as parameters.  
	- It first checks if the input string is empty. If it is, it returns an empty completion tokens object.  
	- Then, it calls the model's predict function to get the predictions.  
	- The function then iterates through the predictions and calls the callback function for each prediction.  
	- Finally, the function returns the completion tokens object.  
  
	- The code also includes a function to compute the total tokens. This function takes the input, input string, configuration, application configuration, model loader, and two callback functions as parameters.  
	- It first checks if the input string is empty. If it is, it returns an empty total tokens object.  
	- Then, it calls the model's predict function to get the predictions.  
	- The function then iterates through the predictions and calls the callback function for each prediction.  
	- Finally, the function returns the total tokens object.  
  
	- The code also includes a function to compute the timing token generation. This function takes the input, input string, configuration, application configuration, model loader, and two callback functions as parameters.  
	- It first checks if the input string is empty. If it is, it returns an empty timing token generation object.  
	- Then, it calls the model's predict function to get the predictions.  
	- The function then iterates through the predictions and calls the callback function for each prediction.  
	- Finally, the function returns the timing token generation object.  
  
	- The code also includes a function to compute the timing prompt processing. This function takes the input, input string, configuration, application configuration, model loader, and two callback functions as parameters.  
	- It first checks if the input string is empty. If it is, it returns an empty timing prompt processing object.  
	- Then, it calls the model's predict function to get the predictions.  
	- The function then iterates through the predictions and calls the callback function for each prediction.  
	- Finally, the function returns the timing prompt processing object.  
  
	- The code also includes a function to compute the prompt tokens. This function takes the input, input string, configuration, application configuration, model loader, and two callback functions as parameters.  
	- It first checks if the input string is empty. If it is, it returns an empty prompt tokens object.  
	- Then, it calls the model's predict function to get the predictions.  
	- The function then iterates through the predictions and calls the callback function for each prediction.  
	- Finally, the function returns the prompt tokens object.  
  
	- The code also includes a function to compute the completion tokens. This function takes the input, input string, configuration, application configuration, model loader, and two callback functions as parameters.  
	- It first checks if the input string is empty. If it is, it returns an empty completion tokens object.  
	- Then, it calls the model's predict function to get the predictions.  
	- The function then iterates through the predictions and calls the callback function for each prediction.  
	- Finally, the function returns the completion tokens object.  
  
	- The code also includes a function to compute the total tokens. This function takes the input, input string, configuration, application configuration, model loader, and two callback functions as parameters.  
	- It first checks if the input string is empty. If it is, it returns an empty total tokens object.  
	- Then, it calls the model's predict function to get the predictions.  
	- The function then iterates through the predictions and calls the callback function for each prediction.  
	- Finally, the function returns the total tokens object.  
  
	- The code also includes a function to compute the timing token generation. This function takes the input, input string, configuration, application configuration, model loader, and two callback functions as parameters.  
	- It first checks if the input string is empty. If it is, it returns an empty timing token generation object.  
	- Then, it calls the model's predict function to get the predictions.  
	- The function then iterates through the predictions and calls the callback function for each prediction.  
	- Finally, the function returns the timing token generation object.  
  
	- The code also includes a function to compute the timing prompt processing. This function takes the input, input string, configuration, application configuration, model loader, and two callback functions as parameters.  
	- It first checks if the input string is empty. If it is, it returns an empty timing prompt processing object.  
	- Then, it calls the model's predict function to get the predictions.  
	- The function then iterates through the predictions and calls the callback function for each prediction.  
	- Finally, the function returns the timing prompt processing object.  
  
	- The code also includes a function to compute the prompt tokens. This function takes the input, input string, configuration, application configuration, model loader, and two callback functions as parameters.  
	- It first checks if the input string is empty. If it is, it returns an empty prompt tokens object.  
	- Then, it calls the model's predict function to get the predictions.  
	- The function then iterates through the predictions and calls the callback function for each prediction.  
	- Finally, the function returns the prompt tokens object.  
  
	- The code also includes a function to compute the completion tokens. This function takes the input, input string, configuration, application configuration, model loader, and two callback functions as parameters.  
	- It first checks if the input string is empty. If it is, it returns an empty completion tokens object.  
	- Then, it calls the model's predict function to get the predictions.  
	- The function then iterates through the predictions and calls the callback function for each prediction.  
	- Finally, the function returns the completion tokens object.  
  
	- The code also includes a function to compute the total tokens. This function takes the input, input string, configuration, application configuration, model loader, and two callback functions as parameters.  
	- It first checks if the input string is empty. If it is, it returns an empty total tokens object.  
	- Then, it calls the model's predict function to get the predictions.  
	- The function then iterates through the predictions and calls the callback function for each prediction.  
	- Finally, the function returns the total tokens object.  
  
	- The code also includes a function to compute the timing token generation. This function takes the input, input string, configuration, application configuration, model loader, and two callback functions as parameters.  
	- It first checks if the input string is empty. If it is, it returns an empty timing token generation object.  
	- Then, it calls the model's predict function to get the predictions.  
	- The function then iterates through the predictions and calls the callback function for each prediction.  
	- Finally, the function returns the timing token generation object.  
  
	- The code also includes a function to compute the timing prompt processing. This function takes the input, input string, configuration, application configuration, model loader, and two callback functions as parameters.  
	- It first checks if the input string is empty. If it is, it returns an empty timing prompt processing object.  
	- Then, it calls the model's predict function to get the predictions.  
	- The function then iterates through the predictions and calls the callback function for each prediction.  
	- Finally, the function returns the timing prompt processing object.  
  
	- The code also includes a function to compute the prompt tokens. This function takes the input, input string, configuration, application configuration, model loader, and two callback functions as parameters.  
	- It first checks if the input string is empty. If it is, it returns an empty prompt tokens object.  
	- Then, it calls the model's predict function to get the predictions.  
	- The function then iterates through the predictions and calls the callback function for each prediction.  
	- Finally, the function returns the prompt tokens object.  
  
	- The code also includes a function to compute the completion tokens. This function takes the input, input string, configuration, application configuration, model loader, and two callback functions as parameters.  
	- It first checks if the input string is empty. If it is, it returns an empty completion tokens object.  
	- Then, it calls the model's predict function to get the predictions.  
	- The function then iterates through the predictions and calls the callback function for each prediction.  
	- Finally, the function returns the completion tokens object.  
  
	- The code also includes a function to compute the total tokens. This function takes the input, input string, configuration, application configuration, model loader, and two callback functions as parameters.  
	- It first checks if the input string is empty. If it is, it returns an empty total tokens object.  
	- Then, it calls the model's predict function to get the predictions.  
	- The function then iterates through the predictions and calls the callback function for each prediction.  
	- Finally, the function returns the total tokens object.  
  
	- The code also includes a function to compute the timing token generation. This function takes the input, input string, configuration, application configuration, model loader, and two callback functions as parameters.  
	- It first checks if the input string is empty. If it is, it returns an empty timing token generation object.  
	- Then, it calls the model's predict function to get the predictions.  
	- The function then iterates through the predictions and calls the callback function for each prediction.  
	- Finally, the function returns the timing token generation object.  
  
	- The code also includes a function to compute the timing prompt processing. This function takes the input, input string, configuration, application configuration, model loader, and two callback functions as parameters.  
	- It first checks if the input string is empty. If it is, it returns an empty timing prompt processing object.  
	- Then, it calls the model's predict function to get the predictions.  
	- The function then iterates through the predictions and calls the callback function for each prediction.  
	- Finally, the function returns the timing prompt processing object.  
  
	- The code also includes a function to compute the prompt tokens. This function takes the input, input string, configuration, application configuration, model loader, and two callback functions as parameters.  
	- It first checks if the input string is empty. If it is, it returns an empty prompt tokens object.  
	- Then, it calls the model's predict function to get the predictions.  
	- The function then iterates through the predictions and calls the callback function for each prediction.  
	- Finally, the function returns the prompt tokens object.  
  
	- The code also includes a function to compute the completion tokens. This function takes the input, input string, configuration, application configuration, model loader, and two callback functions as parameters.  
	- It first checks if the input string is empty. If it is, it returns an empty completion tokens object.  
	- Then, it calls the model's predict function to get the predictions.  
	- The function then iterates through the predictions and calls the callback function for each prediction.  
	- Finally, the function returns the completion tokens object.  
  
	- The code also includes a function to compute the total tokens. This function takes the input, input string, configuration, application configuration, model loader, and two callback functions as parameters.  
	- It first checks if the input string is empty. If it is, it returns an empty total tokens object.  
	- Then, it calls the model's predict function to get the predictions.  
	- The function then iterates through the predictions and calls the callback function for each prediction.  
	- Finally, the function returns the total tokens object.  
  
	- The code also includes a function to compute the timing token generation. This function takes the input, input string, configuration, application configuration, model loader, and two callback functions as parameters.  
	- It first checks if the input string is empty. If it is, it returns an empty timing token generation object.  
	- Then, it calls the model's predict function to get the predictions.  
	- The function then iterates through the predictions and calls the callback function for each prediction.  
	- Finally, the function returns the timing token generation object.  
  
	- The code also includes a function to compute the timing prompt processing. This function takes the input, input string, configuration, application configuration, model loader, and two callback functions as parameters.  
	- It first checks if the input string is empty. If it is, it returns an empty timing prompt processing object.  
	- Then, it calls the model's predict function to get the predictions.  
	- The function then iterates through the predictions and calls the callback function for each prediction.  
	- Finally, the function returns the timing prompt processing object.  
  
	- The code also includes a function to compute the prompt tokens. This function takes the input, input string, configuration, application configuration, model loader, and two callback functions as parameters.  
	- It first checks if the input string is empty. If it is, it returns an empty prompt tokens object.  
	- Then, it calls the model's predict function to get the predictions.  
	- The function then iterates through the predictions and calls the callback function for each prediction.  
	- Finally, the function returns the prompt tokens object.  
  
	- The code also includes a function to compute the completion tokens. This function takes the input, input string, configuration, application configuration, model loader, and two callback functions as parameters.  
	- It first checks if the input string is empty. If it is, it returns an empty completion tokens object.  
	- Then, it calls the model's predict function to get the predictions.  
	- The function then iterates through the predictions and calls the callback function for each prediction.  
	- Finally, the function returns the completion tokens object.  
  
	- The code also includes a function to compute the total tokens. This function takes the input, input string, configuration, application configuration, model loader, and two callback functions as parameters.  
	- It first checks if the input string is empty. If it is, it returns an empty total tokens object.  
	- Then, it calls the model's predict function to get the predictions.  
	- The function then iterates through the predictions and calls the callback function for each prediction.  
	- Finally, the function returns the total tokens object.  
  
	- The code also includes a function to compute the timing token generation. This function takes the input, input string, configuration, application configuration, model loader, and two callback functions as parameters.  
	- It first checks if the input string is empty. If it is, it returns an empty timing token generation object.  
	- Then, it calls the model's predict function to get the predictions.  
	- The function then iterates through the predictions and calls the callback function for each prediction.  
	- Finally, the function returns the timing token generation object.  
  
	- The code also includes a function to compute the timing prompt processing. This function takes the input, input string, configuration, application configuration, model loader, and two callback functions as parameters.  
	- It first checks if the input string is empty. If it is, it returns an empty timing prompt processing object.  
	- Then, it calls the model's predict function to get the predictions.  
	- The function then iterates through the predictions and calls the callback function for each prediction.  
	- Finally, the function returns the timing prompt processing object.  
  
	- The code also includes a function to compute the prompt tokens. This function takes the input, input string, configuration, application configuration, model loader, and two callback functions as parameters.  
	- It first checks if the input string is empty. If it is, it returns an empty prompt tokens object.  
	- Then, it calls the model's predict function to get the predictions.  
	- The function then iterates through the predictions and calls the callback function for each prediction.  
	- Finally, the function returns the prompt tokens object.  
  
	- The code also includes a function to compute the completion tokens. This function takes the input, input string, configuration, application configuration, model loader, and two callback functions as parameters.  
	- It first checks if the input string is empty. If it is, it returns an empty completion tokens object.  
	- Then, it calls the model's predict function to get the predictions.  
	- The function then iterates through the predictions and calls the callback function for each prediction.  
	- Finally, the function returns the completion tokens object.  
  
	- The code also includes a function to compute the total tokens. This function takes the input, input string, configuration, application configuration, model loader, and two callback functions as parameters.  
	- It first checks if the input string is empty. If it is, it returns an empty total tokens object.  
	- Then, it calls the model's predict function to get the predictions.  
	- The function then iterates through the predictions and calls the callback function for each prediction.  
	- Finally, the function returns the total tokens object.  
  
	- The code also includes a function to compute the timing token generation. This function takes the input, input string, configuration, application configuration, model loader, and two callback functions as parameters.  
	- It first checks if the input string is empty. If it is, it returns an empty timing token generation object.  
	- Then, it calls the model's predict function to get the predictions.  
	- The function then iterates through the predictions and calls the callback function for each prediction.  
	- Finally, the function returns the timing token generation object.  
  
	- The code also includes a function to compute the timing prompt processing. This function takes the input, input string, configuration, application configuration, model loader, and two callback functions as parameters.  
	- It first checks if the input string is empty. If it is, it returns an empty timing prompt processing object.  
	- Then, it calls the model's predict function to get the predictions.  
	- The function then iterates through the predictions and calls the callback function for each prediction.  
	- Finally, the function returns the timing prompt processing object.  
  
	- The code also includes a function to compute the prompt tokens. This function takes the input, input string, configuration, application configuration, model loader, and two callback functions as parameters.  
	- It first checks if the input string is empty. If it is, it returns an empty prompt tokens object.  
	- Then, it calls the model's predict function to get the predictions.  
	- The function then iterates through the predictions and calls the callback function for each prediction.  
	- Finally, the function returns the prompt tokens object.  
  
	- The code also includes a function to compute the completion tokens. This function takes the input, input string, configuration, application configuration, model loader, and two callback functions as parameters.  
	- It first checks if the input string is empty. If it is, it returns an empty completion tokens object.  
	- Then, it calls the model's predict function to get the predictions.  
	- The function then iterates through the predictions and calls the callback function for each prediction.  
	- Finally, the function returns the completion tokens object.  
  
	- The code also includes a function to compute the total tokens. This function takes the input, input string, configuration, application configuration, model loader, and two callback functions as parameters.  
	- It first checks if the input string is empty. If it is, it returns an empty total tokens object.  
	- Then, it calls the model's predict function to get the predictions.  
	- The function then iterates through the predictions and calls the callback function for each prediction.  
	- Finally, the function returns the total tokens object.  
  
	- The code also includes a function to compute the timing token generation. This function takes the input, input string, configuration, application configuration, model loader, and two callback functions as parameters.  
	- It first checks if the input string is empty. If it is, it returns an empty timing token generation object.  
	- Then, it calls the model's predict function to get the predictions.  
	- The function then iterates through the predictions and calls the callback function for each prediction.  
	- Finally, the function returns the timing token generation object.  
  
	- The code also includes a function to compute the timing prompt processing. This function takes the input, input string, configuration, application configuration, model loader, and two callback functions as parameters.  
	- It first checks if the input string is empty. If it is, it returns an empty timing prompt processing object.  
	- Then, it calls the model's predict function to get the predictions.  
	- The function then iterates through the predictions and calls the callback function for each prediction.  
	- Finally, the function returns the timing prompt processing object.  
  
	- The code also includes a function to compute the prompt tokens. This function takes the input, input string, configuration, application configuration, model loader, and two callback functions as parameters.  
	- It first checks if the input string is empty. If it is, it returns an empty prompt tokens object.  
	- Then, it calls the model's predict function to get the predictions.  
	- The function then iterates through the predictions and calls the callback function for each prediction.  
	- Finally, the function returns the prompt tokens object.  
  
	- The code also includes a function to compute the completion tokens. This function takes the input,  
  
# core/http/endpoints/openai/embeddings.go  
Package: openai  
  
Imports:  
- "encoding/json"  
- "time"  
- "github.com/mudler/LocalAI/core/backend"  
- "github.com/mudler/LocalAI/core/config"  
- "github.com/mudler/LocalAI/core/http/middleware"  
- "github.com/mudler/LocalAI/pkg/model"  
- "github.com/google/uuid"  
- "github.com/mudler/LocalAI/core/schema"  
- "github.com/gofiber/fiber/v2"  
- "github.com/rs/zerolog/log"  
  
External data, input sources:  
- The code uses the OpenAI Embeddings API endpoint to get vector representations of given inputs.  
- It relies on the LocalAI backend and model loaders to access the necessary models and configurations.  
  
TODOs:  
- There are no TODO comments in the provided code.  
  
Code summary:  
- The EmbeddingsEndpoint function handles requests to the OpenAI Embeddings API endpoint. It takes the request body, model loader, and application configuration as input.  
- The function first checks if the request contains a valid model and input. If not, it returns a bad request error.  
- Then, it iterates through the input tokens and strings, calling the appropriate model embedding function for each.  
- The embeddings are collected and stored in a list of schema.Item objects.  
- Finally, the function creates a schema.OpenAIResponse object containing the ID, creation timestamp, model, data, and object type.  
- The response is then marshalled into JSON format and returned to the client.  
  
  
  
# core/http/endpoints/openai/files.go  
Package: openai  
  
Imports:  
- "errors"  
- "fmt"  
- "os"  
- "path/filepath"  
- "sync/atomic"  
- "time"  
- "github.com/microcosm-cc/bluemonday"  
- "github.com/mudler/LocalAI/core/config"  
- "github.com/mudler/LocalAI/core/schema"  
- "github.com/gofiber/fiber/v2"  
- "github.com/mudler/LocalAI/pkg/utils"  
  
External data, input sources:  
- UploadedFilesFile: "uploadedFiles.json"  
- UploadedFiles: []schema.File  
  
TODOs:  
- TODO put in purpose dirs  
  
Summary:  
- UploadFilesEndpoint: This function handles the upload of files to the server. It checks the file size, purpose, and filename for validity. If the file is valid, it saves the file to the server and updates the UploadedFiles list.  
- ListFilesEndpoint: This function lists all uploaded files or files with a specific purpose.  
- GetFilesEndpoint: This function retrieves information about a specific file based on its ID.  
- DeleteFilesEndpoint: This function deletes a file from the server and updates the UploadedFiles list.  
- GetFilesContentsEndpoint: This function retrieves the contents of a specific file based on its ID.  
  
- UploadedFiles: This is a global variable that stores a list of all uploaded files.  
- getNextFileId: This function generates a unique ID for each uploaded file.  
- getFileFromRequest: This function retrieves a file from the request based on its ID.  
  
# core/http/endpoints/openai/files_test.go  
Package: openai  
  
Imports:  
	"encoding/json"  
	"fmt"  
	"io"  
	"mime/multipart"  
	"net/http"  
	"net/http/httptest"  
	"os"  
	"path/filepath"  
	"strings"  
	"github.com/rs/zerolog/log"  
	"github.com/mudler/LocalAI/core/config"  
	"github.com/mudler/LocalAI/core/schema"  
	"github.com/gofiber/fiber/v2"  
	utils2 "github.com/mudler/LocalAI/pkg/utils"  
	"github.com/stretchr/testify/assert"  
	"testing"  
  
External data and input sources:  
	- Configuration settings from the config package  
	- Uploaded files from the user  
  
TODOs:  
	- Add more tests for different file types and sizes  
	- Implement error handling for file upload failures  
	- Add support for file deletion and retrieval  
  
Summary:  
	- The package provides a set of functions for handling file uploads, storage, and retrieval.  
	- It includes endpoints for uploading, listing, and deleting files.  
	- The package also includes helper functions for creating test files and handling multipart file uploads.  
	- The code is well-documented and includes unit tests to ensure proper functionality.  
  
	- The `startUpApp` function sets up a test server with endpoints for file upload, listing, and deletion.  
	- The `TestUploadFileExceedSizeLimit` function tests the file upload endpoint with various scenarios, such as exceeding the upload limit, missing purpose parameter, and file already existing.  
	- The `CallListFilesEndpoint` function tests the endpoint for listing files with and without purpose parameter.  
	- The `CallFilesContentEndpoint` function tests the endpoint for getting file content.  
	- The `CallFilesUploadEndpoint` function tests the file upload endpoint with different file sizes and purposes.  
	- The `CallFilesUploadEndpointWithCleanup` function tests the file upload endpoint and cleans up the uploaded file after the test.  
	- The `CallFilesDeleteEndpoint` function tests the endpoint for deleting files.  
	- The `newMultipartFile` function creates a multipart file for testing purposes.  
	- The `createTestFile` function creates a test file with a specified size.  
	- The `bodyToString` function converts the response body to a string.  
	- The `bodyToByteArray` function converts the response body to a byte array.  
	- The `responseToFile` function converts the response to a File object.  
	- The `responseToListFile` function converts the response to a ListFiles object.  
  
	- The package is designed to be used in a local AI application, where users can upload and manage their own files.  
	- It provides a simple and efficient way to handle file uploads and storage, making it easy for users to manage their data.  
  
# core/http/endpoints/openai/image.go  
Package: openai  
  
Imports:  
	"bufio"  
	"encoding/base64"  
	"encoding/json"  
	"fmt"  
	"io"  
	"net/http"  
	"os"  
	"path/filepath"  
	"strconv"  
	"strings"  
	"time"  
	"github.com/google/uuid"  
	"github.com/mudler/LocalAI/core/config"  
	"github.com/mudler/LocalAI/core/http/middleware"  
	"github.com/mudler/LocalAI/core/schema"  
	"github.com/mudler/LocalAI/core/backend"  
	"github.com/gofiber/fiber/v2"  
	model "github.com/mudler/LocalAI/pkg/model"  
	"github.com/rs/zerolog/log"  
  
External data, input sources:  
	- The code uses the `http` package to make GET requests to download files from URLs.  
	- It reads input from the user through the command line.  
	- It reads configuration settings from a file.  
  
TODOs:  
	- Add support for different image formats.  
	- Add support for different image sizes.  
	- Add support for different image quality settings.  
  
Summary:  
	- The `downloadFile` function downloads a file from a given URL and returns the path to the downloaded file.  
	- The `ImageEndpoint` function handles requests to the OpenAI Image generation API endpoint. It takes a prompt, a file, and other parameters as input, and returns an image based on the prompt and file.  
	- The function first checks if the input file is a URL or a base64 encoded string. If it's a URL, it downloads the file and saves it to a temporary location. If it's a base64 encoded string, it decodes the string and saves the file to a temporary location.  
	- The function then checks the configuration settings to determine the backend to use for image generation.  
	- Based on the backend, it calls the appropriate function to generate the image.  
	- Finally, the function returns the generated image to the user.  
  
	- The code also includes functions to handle different image generation parameters, such as size, mode, and step.  
	- It also includes functions to handle the response format, such as returning the image as a base64 encoded string or a URL to the generated image.  
  
# core/http/endpoints/openai/inference.go  
Package: openai  
  
Imports:  
- github.com/mudler/LocalAI/core/backend  
- github.com/mudler/LocalAI/core/config  
- github.com/mudler/LocalAI/core/schema  
- github.com/mudler/LocalAI/pkg/model  
  
External data, input sources:  
- req: A pointer to a schema.OpenAIRequest object, which contains the user's request and context.  
- predInput: A string representing the input to the prediction function.  
- config: A pointer to a config.BackendConfig object, which contains backend-specific configuration.  
- o: A pointer to a config.ApplicationConfig object, which contains application-wide configuration.  
- loader: A pointer to a model.ModelLoader object, which is responsible for loading the model.  
- cb: A callback function that takes a string and a pointer to a slice of schema.Choice objects as input.  
- tokenCallback: A callback function that takes a string and a backend.TokenUsage object as input.  
  
TODOs:  
- None  
  
Summary:  
The ComputeChoices function takes a user's request, input to the prediction function, backend and application configuration, model loader, and callback functions as input. It then extracts images, videos, and audios from the request, calls the backend's ModelInference function to get the model function to call for the result, and iterates through the number of completions requested. For each completion, it calls the model function, updates the token usage, and calls the callback function with the finetuned response and the result slice. Finally, it returns the result slice, token usage, and any errors encountered.  
  
  
  
# core/http/endpoints/openai/list.go  
Package: openai  
  
Imports:  
- github.com/gofiber/fiber/v2  
- github.com/mudler/LocalAI/core/config  
- github.com/mudler/LocalAI/core/schema  
- github.com/mudler/LocalAI/core/services  
- github.com/mudler/LocalAI/pkg/model  
  
External data, input sources:  
- The code uses the `config.BackendConfigLoader` and `model.ModelLoader` to access configuration and model data.  
- It also uses the `fiber.Ctx` object to get query parameters from the incoming request.  
  
TODOs:  
- TODO: give more options to the user?  
  
Summary:  
The `ListModelsEndpoint` function is responsible for handling the OpenAI Models API endpoint. It takes a `BackendConfigLoader`, `ModelLoader`, and `ApplicationConfig` as input. The function first retrieves the filter parameter from the request query. If the filter parameter is not provided, no filter is applied. The function then determines the loose file policy based on the `excludeConfigured` query parameter. If the parameter is true, the policy is set to `SKIP_IF_CONFIGURED`, otherwise it is set to `ALWAYS_INCLUDE`. The function then uses the filter parameter and loose file policy to list the available models. Finally, the function returns a JSON response containing a list of OpenAIModel objects, each with an ID and an object type of "model".  
  
  
  
# core/http/endpoints/openai/transcription.go  
Package: openai  
  
Imports:  
	"io"  
	"net/http"  
	"os"  
	"path"  
	"path/filepath"  
	"github.com/mudler/LocalAI/core/backend"  
	"github.com/mudler/LocalAI/core/config"  
	"github.com/mudler/LocalAI/core/http/middleware"  
	"github.com/mudler/LocalAI/core/schema"  
	model "github.com/mudler/LocalAI/pkg/model"  
	"github.com/gofiber/fiber/v2"  
	"github.com/rs/zerolog/log"  
  
External data, input sources:  
	- The code uses the OpenAI Whisper API endpoint to transcribe audio into the input language.  
	- It retrieves the file data from the request.  
  
TODOs:  
	- TODO: handle different outputs here  
  
Summary:  
	- The TranscriptEndpoint function handles the OpenAI Whisper API endpoint for transcribing audio. It takes the model, file, and language as input and returns the transcribed text.  
	- The function first retrieves the file data from the request and saves it to a temporary directory.  
	- Then, it uses the backend.ModelTranscription function to transcribe the audio using the specified model and language.  
	- Finally, it returns the transcribed text as a JSON response.  
	- The TODO comment indicates that there is a need to handle different outputs from the transcription process.  
  
  
  
