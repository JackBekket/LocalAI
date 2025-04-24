# services

This package contains various services for managing and monitoring models and galleries.

## File structure:
- backend_monitor.go
- gallery.go
- list_models.go
- metrics.go

## Code summary:
- BackendMonitorService: This service monitors the status of backend models and provides information about their resource usage. It uses the config.BackendConfigLoader to load backend configuration data and the model.ModelLoader to load model data. The service provides methods to check the status of a backend model, sample its resource usage, and shut down the model.
- GalleryService: This service is responsible for managing the installation, deletion, and updates of models and galleries. It receives requests from users and processes them accordingly. The service uses the gallery package to interact with the gallery system and the startup package to install models.
- ListModels: This function lists all available models, considering both configured and loose files. It takes a BackendConfigLoader, a ModelLoader, a filter function, and a LooseFilePolicy as input. The function first iterates through the configured models and adds them to the list if they match the specified policy. Then, it iterates through the loose files and adds them to the list if they match the filter and are not skipped according to the policy. Finally, it returns the list of models and any errors encountered.
- CheckIfModelExists: This function checks if a specific model exists, considering both configured and loose files. It takes a BackendConfigLoader, a ModelLoader, the model name, and a LooseFilePolicy as input. The function first builds a filter function based on the model name and then calls the ListModels function to get a list of models matching the filter. Finally, it returns true if the list contains any models and false otherwise.
- LocalAIMetricsService: This struct is responsible for collecting and reporting metrics related to API calls. It uses OpenTelemetry to collect data and export it to a Prometheus exporter.

## Relations between code entities:
- BackendMonitorService and CheckIfModelExists are both part of the services package and work together to monitor and manage backend models.
- ListModels and CheckIfModelExists are both functions that work with models and galleries.
- LocalAIMetricsService is responsible for collecting and reporting metrics related to API calls.

## Unclear places:
- The purpose of the Shutdown method in the LocalAIMetricsService is unclear, as the implementation is incomplete and only logs a warning message.

## Dead code:
- None found.

