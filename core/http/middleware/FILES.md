# core/http/middleware/auth.go  
Package name: middleware  
  
Imports:  
"crypto/subtle"  
"errors"  
"github.com/dave-gray101/v2keyauth"  
"github.com/gofiber/fiber/v2"  
"github.com/gofiber/fiber/v2/middleware/keyauth"  
"github.com/mudler/LocalAI/core/config"  
"github.com/mudler/LocalAI/core/http/utils"  
  
External data, input sources:  
- Application configuration settings from the config package.  
  
TODOs:  
- Investigate the possibility of migrating to v3 of the fiber/keyauth middleware once it stabilizes.  
  
Summary:  
This file contains the configuration generators and handler functions that are used along with the fiber/keyauth middleware. It provides functions to handle API key validation, error handling, and filtering.  
  
- getApiKeyErrorHandler: This function handles errors related to API key validation. It checks if the error is due to a missing or malformed API key and returns a 401 Unauthorized status code if it is. If opaque errors are enabled, it returns a 500 Internal Server Error status code.  
- getApiKeyValidationFunction: This function validates the API key provided by the client. It checks if the key is present in the list of valid API keys and returns true if it is, otherwise it returns false and an error.  
- getApiKeyRequiredFilterFunction: This function determines whether the API key is required for the current request. It checks if the request method is GET and if the endpoint is exempted from the API key requirement. If both conditions are met, it returns false, otherwise it returns true.  
- GetKeyAuthConfig: This function generates the configuration for the fiber/keyauth middleware. It sets up the custom key lookup, error handler, validation function, and authentication scheme.  
  
# core/http/middleware/request.go  
## Package: middleware  
  
This package contains middleware functions for handling requests and extracting information from them.  
  
### External Data and Input Sources:  
  
- The code relies on the `context` package for managing request context and cancellation.  
- It uses the `fiber/v2` package for handling HTTP requests and responses.  
- It interacts with the `config` package to load backend configuration files.  
- It uses the `model` package to load and manage models.  
- It leverages the `schema` package to define data structures for localAI requests.  
- It depends on the `services` package for interacting with backend services.  
- It utilizes the `functions` package to handle function calls.  
- It employs the `templates` package for templating content.  
- It uses the `utils` package for various utility functions.  
  
### TODOs:  
  
- Refactor `setModelNameFromRequest` to not return an error if unchanged.  
- If context and cancel above belong on all methods, move that part of `SetModelAndConfig` into `SetModelAndConfig`. Otherwise, it's in its own method below for now.  
  
### Major Code Parts:  
  
1. **Request Extractor:**  
   - The `RequestExtractor` struct is responsible for extracting information from incoming requests and setting up the necessary context for processing.  
   - It includes methods for setting the model name, setting the OpenAI request, and building middleware handlers for default model selection and filtering.  
  
2. **Model Name Handling:**  
   - The `setModelNameFromRequest` method extracts the model name from the request parameters, query string, or authorization header.  
   - The `BuildConstantDefaultModelNameMiddleware` and `BuildFilteredFirstAvailableDefaultModelName` methods provide middleware handlers for setting default model names.  
  
3. **Context and Configuration Setup:**  
   - The `SetModelAndConfig` method sets the model name and configuration in the request context.  
   - The `SetOpenAIRequest` method sets the OpenAI request and backend configuration in the request context.  
  
4. **OpenAI Request Handling:**  
   - The `mergeOpenAIRequestAndBackendConfig` function merges the OpenAI request and backend configuration, handling various parameters and settings.  
  
5. **Utility Functions:**  
   - The code includes utility functions for handling correlation IDs, setting up request context, and merging request data with backend configuration.  
  
# core/http/middleware/strippathprefix.go  
Package: middleware  
  
Imports:  
- "strings"  
- "github.com/gofiber/fiber/v2"  
  
External data, input sources:  
- The middleware reads the X-Forwarded-Prefix HTTP request header to determine the path prefix to strip.  
  
TODOs:  
- There are no TODO comments in the code.  
  
Summary:  
- The StripPathPrefix function is a middleware that strips a path prefix from the request path. It reads the X-Forwarded-Prefix HTTP request header to determine the path prefix to strip. If the path prefix is found in the request path, it strips the prefix from the path. If the path prefix is not found in the request path, it redirects the request to the path prefix.  
  
- The middleware is implemented as a function that takes a fiber.Ctx object as input and returns a fiber.Handler object. The fiber.Ctx object contains information about the incoming HTTP request, such as the request path and headers. The fiber.Handler object is a function that is executed when the middleware is invoked.  
  
- The middleware first checks if the X-Forwarded-Prefix header is present in the request headers. If it is, it iterates over the values of the header and checks if the prefix is not empty. If the prefix is not empty, it extracts the path from the fiber.Ctx object and checks if the path starts with the prefix. If it does, it strips the prefix from the path and updates the path in the fiber.Ctx object. If the prefix is not found in the path, it redirects the request to the prefix.  
  
- If the X-Forwarded-Prefix header is not present in the request headers, the middleware calls the Next() method on the fiber.Ctx object to pass control to the next middleware in the chain.  
  
- The middleware is designed to be used in a Fiber application to strip a path prefix from the request path. This can be useful in situations where the application is deployed behind a reverse proxy that adds a path prefix to the request path.  
  
- The middleware is simple to use and can be easily integrated into a Fiber application. It is a useful tool for developers who need to strip a path prefix from the request path in a Fiber application.  
  
# core/http/middleware/strippathprefix_test.go  
Package name: middleware  
  
Imports:  
"net/http/httptest"  
"testing"  
"github.com/gofiber/fiber/v2"  
"github.com/stretchr/testify/require"  
  
External data, input sources:  
- The code uses the 'X-Forwarded-Prefix' header to determine the prefix to strip.  
  
TODOs:  
- There are no TODO comments in the code.  
  
Summary:  
- The code defines a middleware function called StripPathPrefix() that strips a specified prefix from the request path.  
- The middleware function checks for the 'X-Forwarded-Prefix' header and compares it to the request path.  
- If the header is present and matches the request path, the prefix is stripped from the path.  
- If the header is not present or does not match the request path, the original path is used.  
- The code includes unit tests to verify the functionality of the middleware function.  
- The tests cover various scenarios, including cases where the prefix is present and absent, and cases where the header matches and does not match the request path.  
- The tests assert that the middleware function correctly strips the prefix from the path and returns the expected status code.  
  
  
  
