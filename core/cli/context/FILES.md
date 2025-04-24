# core/cli/context/context.go  
Package/Component name: cliContext  
  
Imports:  
embed  
  
External data, input sources:  
- Environment variables: LOCALAI_DEBUG, DEBUG, LOCALAI_LOG_LEVEL  
  
TODOs:  
- DEPRECATED, use --log-level=debug instead. Enable debug logging  
  
Summary:  
The cliContext package provides a Context struct that stores configuration settings for the application. The Context struct includes fields for debug mode and log level, which can be set using environment variables. The BackendAssets field is used to store embedded assets, such as configuration files or templates.  
  
The Debug field is a boolean value that indicates whether debug mode is enabled. If the LOCALAI_DEBUG or DEBUG environment variables are set to true, the Debug field will be set to true.  
  
The LogLevel field is a string value that specifies the level of logs to output. The allowed values are error, warn, info, debug, and trace. The LogLevel field can be set using the LOCALAI_LOG_LEVEL environment variable.  
  
The BackendAssets field is an embed.FS object that stores embedded assets. This field is not a command line argument/flag, and it is excluded from the parsed CLI.  
  
The Context struct is used to store and manage the application's configuration settings. It provides a centralized location for storing and accessing configuration data, making it easy to manage and update the application's settings.  
  
