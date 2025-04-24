# core/http/app.go  
Package: http  
  
Imports:  
- "embed"  
- "errors"  
- "fmt"  
- "net/http"  
- "github.com/dave-gray101/v2keyauth"  
- "github.com/mudler/LocalAI/pkg/utils"  
- "github.com/mudler/LocalAI/core/http/endpoints/localai"  
- "github.com/mudler/LocalAI/core/http/endpoints/openai"  
- "github.com/mudler/LocalAI/core/http/middleware"  
- "github.com/mudler/LocalAI/core/http/routes"  
- "github.com/mudler/LocalAI/core/application"  
- "github.com/mudler/LocalAI/core/schema"  
- "github.com/mudler/LocalAI/core/services"  
- "github.com/gofiber/contrib/fiberzerolog"  
- "github.com/gofiber/fiber/v2"  
- "github.com/gofiber/fiber/v2/middleware/cors"  
- "github.com/gofiber/fiber/v2/middleware/csrf"  
- "github.com/gofiber/fiber/v2/middleware/favicon"  
- "github.com/gofiber/fiber/v2/middleware/filesystem"  
- "github.com/gofiber/fiber/v2/middleware/recover"  
- "github.com/rs/zerolog/log"  
  
External data and input sources:  
- Application configuration settings  
- User input through API requests  
  
TODOs:  
- Add support for custom error handlers  
- Implement a mechanism for handling API rate limits  
- Improve documentation and examples  
  
Summary:  
- The `API` function initializes and configures a Fiber web server for the LocalAI application.  
- It sets up middleware for handling CORS, CSRF, and authentication.  
- It registers routes for various API endpoints, including health checks, user management, and model interactions.  
- The function also handles custom error responses and provides a default 404 handler.  
- The `embedDirStatic` variable embeds a directory of static files, which are served by the web server.  
- The `renderEngine` function is responsible for rendering HTML templates.  
- The `notFoundHandler` function handles requests for non-existent endpoints.  
- The `LocalAIMetricsAPIMiddleware` middleware is used to collect metrics about the API usage.  
- The `GetKeyAuthConfig` function retrieves the KeyAuth configuration settings.  
- The `LoadConfig` function loads configuration files for various components of the application.  
- The `NewGalleryService` function creates a new instance of the Gallery service.  
- The `NewRequestExtractor` function creates a new instance of the RequestExtractor, which is used to extract information from API requests.  
- The `RegisterElevenLabsRoutes`, `RegisterLocalAIRoutes`, `RegisterOpenAIRoutes`, and `RegisterUIRoutes` functions register routes for various API endpoints.  
- The `RegisterJINARoutes` function registers routes for the JINA AI framework.  
  
  
  
# core/http/explorer.go  
## Package: http  
  
Imports:  
- "net/http"  
- "github.com/gofiber/fiber/v2"  
- "github.com/gofiber/fiber/v2/middleware/favicon"  
- "github.com/gofiber/fiber/v2/middleware/filesystem"  
- "github.com/mudler/LocalAI/core/explorer"  
- "github.com/mudler/LocalAI/core/http/middleware"  
- "github.com/mudler/LocalAI/core/http/routes"  
  
External data, input sources:  
- The code uses a database instance from the "explorer" package.  
  
TODOs:  
- There are no TODO comments in the code.  
  
### Explorer function  
This function takes a database instance as input and returns a Fiber application. It sets up the Fiber application with a custom configuration, registers explorer routes, and adds middleware for handling favicons and static files. The function also defines a custom 404 handler.  
  
### Fiber configuration  
The code configures the Fiber application with a custom view engine, disables the startup message, and overrides the default error handler.  
  
### Middleware  
The code uses middleware to strip the path prefix, handle favicons, and serve static files.  
  
### Routes  
The code registers explorer routes with the Fiber application.  
  
### Custom 404 handler  
The code defines a custom 404 handler to handle requests that don't match any registered routes.  
  
### Static files  
The code serves static files from the embedded directory "static".  
  
### Favicons  
The code handles favicon requests by serving the favicon from the embedded directory "static".  
  
### Database interaction  
The code interacts with the database instance to retrieve and update data related to the explorer.  
  
### Error handling  
The code overrides the default error handler to provide custom error handling.  
  
### Logging  
The code logs information about the Fiber application's startup and any errors that occur.  
  
### Security  
The code uses middleware to strip the path prefix and handle favicons, which can help to improve the security of the application.  
  
### Performance  
The code uses a custom view engine and overrides the default error handler to improve the performance of the application.  
  
### Maintainability  
The code is well-organized and easy to understand, which can help to improve the maintainability of the application.  
  
### Scalability  
The code is designed to be scalable, as it uses a database to store and retrieve data.  
  
### Testability  
The code is designed to be testable, as it uses a database to store and retrieve data.  
  
# core/http/render.go  
Package: http  
  
Imports:  
embed  
fmt  
html/template  
net/http  
github.com/Masterminds/sprig/v3  
github.com/gofiber/fiber/v2  
github.com/gofiber/template/html/v2  
github.com/microcosm-cc/bluemonday  
github.com/mudler/LocalAI/core/http/utils  
github.com/mudler/LocalAI/core/schema  
github.com/russross/blackfriday  
  
External data, input sources:  
- Views directory (viewsfs) containing HTML templates.  
  
TODOs:  
- None found.  
  
Summary:  
The http package provides functionality for handling HTTP requests and rendering HTML templates. It includes a custom 404 handler that returns a JSON response if the client accepts JSON or an HTML response if the client accepts HTML. The package also defines a custom template engine that uses Sprig's template functions and a Markdown-to-HTML converter.  
  
- The notFoundHandler function handles 404 errors by returning a JSON or HTML response depending on the client's preferences.  
- The renderEngine function creates a new Fiber template engine that uses the views directory as the source for templates. It also adds Sprig's template functions and a custom Markdown-to-HTML converter to the engine.  
- The markDowner function converts Markdown text to HTML using the Blackfriday library and sanitizes the output using the Bluemonday library.  
  
  
  
