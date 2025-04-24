## Package: library

Imports:
- errors
- fmt
- os
- path/filepath
- runtime
- github.com/rs/zerolog/log

External data, input sources:
- The code reads the environment variable "LOCALAI_SKIP_LIBRARY_PATH" to determine whether to skip loading libraries.
- It also checks for the presence of a "ld.so" file in the asset directory to handle compatibility issues on Linux systems.

TODOs:
- Investigate whether "DYLD_FALLBACK_LIBRARY_PATH" should be "DYLD_LIBRARY_PATH" on macOS.

### LoadExtractedLibs function
This function loads libraries from the asset directory. It first checks if the "LOCALAI_SKIP_LIBRARY_PATH" environment variable is set. If it is, the function returns without loading any libraries. Otherwise, it iterates through two library directories and calls the LoadExternal function for each directory.

### LoadLDSO function
This function checks if there is a "ld.so" file in the asset directory and, if so, prefixes the grpc process with it. This is done to ensure compatibility on Linux systems. If the "LOCALAI_SKIP_LIBRARY_PATH" environment variable is set or the operating system is not Linux, the function returns without modifying the arguments or the grpc process.

### LoadExternal function
This function sets the LD_LIBRARY_PATH environment variable to include the given directory. It first checks if the "LOCALAI_SKIP_LIBRARY_PATH" environment variable is set. If it is, the function returns without modifying the environment variable. Otherwise, it retrieves the current value of the LD_LIBRARY_PATH environment variable and appends the given directory to it. Finally, it sets the LD_LIBRARY_PATH environment variable to the updated value.

- dynaload.go
- pkg/library/dynaload.go

This package provides functionality for loading libraries from the asset directory. It handles compatibility issues on Linux systems by checking for the presence of a "ld.so" file and modifying the grpc process accordingly. The package also allows for the loading of libraries from multiple directories by setting the LD_LIBRARY_PATH environment variable.

The LoadExtractedLibs function is responsible for loading libraries from the asset directory. It checks for the "LOCALAI_SKIP_LIBRARY_PATH" environment variable and, if it is set, skips loading any libraries. Otherwise, it iterates through two library directories and calls the LoadExternal function for each directory.

The LoadLDSO function checks for the presence of a "ld.so" file in the asset directory and, if found, prefixes the grpc process with it. This is done to ensure compatibility on Linux systems. If the "LOCALAI_SKIP_LIBRARY_PATH" environment variable is set or the operating system is not Linux, the function returns without modifying the arguments or the grpc process.

The LoadExternal function sets the LD_LIBRARY_PATH environment variable to include the given directory. It checks for the "LOCALAI_SKIP_LIBRARY_PATH" environment variable and, if it is set, returns without modifying the environment variable. Otherwise, it retrieves the current value of the LD_LIBRARY_PATH environment variable, appends the given directory to it, and sets the LD_LIBRARY_PATH environment variable to the updated value.

The package relies on the "LOCALAI_SKIP_LIBRARY_PATH" environment variable to determine whether to skip loading libraries. It also checks for the presence of a "ld.so" file in the asset directory to handle compatibility issues on Linux systems.

The package is designed to be used in a grpc server environment, where it can load libraries from the asset directory and ensure compatibility on Linux systems.