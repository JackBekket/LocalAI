# templates

pkg/templates/cache.go
pkg/templates/evaluator.go
pkg/templates/evaluator_test.go
pkg/templates/multimodal.go
pkg/templates/multimodal_test.go
pkg/templates/templates_suite_test.go

The package provides a template cache that stores and manages templates for different template types. It supports both Go templates and Jinja templates. The cache is thread-safe and ensures that templates are loaded only once. The package includes functions to load, evaluate, and manage templates. It also includes security checks to prevent loading templates from outside the specified path. The package is designed to be used in a multi-threaded environment, where multiple components may need to access and use templates concurrently. The use of a cache ensures that templates are loaded only once, improving performance and reducing resource consumption.

The package also includes functionality for evaluating prompt templates and chat message templates. It defines various data structures to represent prompt template data and chat message template data. The Evaluator struct is responsible for evaluating templates based on the provided template type and configuration. The EvaluateTemplateForPrompt function handles the evaluation of prompt templates, while the evaluateTemplateForChatMessage function handles chat message templates. The TemplateMessages function takes a list of messages, configuration, functions, and a flag indicating whether to use functions, and returns a string containing the templated input.

The package uses Jinja templates for templating and supports various template types, such as completion, edit, chat, and functions. It also includes logic for handling system prompts and suppressing them if necessary.

The package also includes functionality for handling multimodal content, such as images, audio, and video. It defines a MultiModalOptions struct to store the total and message-specific counts of these media elements. The TemplateMultiModal function takes a template string, MultiModalOptions, and text as input. It compiles the template using the sprig library for additional functions. The function then generates lists of MultimodalContent objects for videos, audios, and images based on the provided MultiModalOptions. Finally, it executes the template with the generated content and text, returning the resulting string.

In essence, the templates package enables the use of templates to dynamically generate prompts and chat messages based on the provided input and configuration. It also provides support for handling multimodal content and generating templates that incorporate this content.

