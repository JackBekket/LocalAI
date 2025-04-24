# pkg/library/dynaload.go  
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
  
  
  
