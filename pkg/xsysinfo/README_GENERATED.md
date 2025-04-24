# xsysinfo

pkg/xsysinfo/cpu.go
Package name: xsysinfo

Imports:
- "sort"
- "github.com/jaypipes/ghw"
- "github.com/klauspost/cpuid/v2"

External data, input sources:
- The code uses the "github.com/jaypipes/ghw" package to access hardware information, including CPU capabilities.
- It also uses the "github.com/klauspost/cpuid/v2" package to retrieve CPU-specific information, such as supported features and physical core count.

TODOs:
- There are no TODO comments in the provided code.

Summary of major code parts:
- CPUCapabilities(): This function retrieves a list of CPU capabilities from the system's hardware information. It uses the "github.com/jaypipes/ghw" package to access the CPU information and then extracts the capabilities. The function returns a sorted list of capabilities as strings.
- HasCPUCaps(): This function checks if the CPU supports a given set of CPUID features. It uses the "github.com/klauspost/cpuid/v2" package to access the CPUID information and then checks if the specified features are supported.
- CPUPhysicalCores(): This function returns the number of physical CPU cores. It uses the "github.com/klauspost/cpuid/v2" package to access the CPU information and returns the number of physical cores. If the number of physical cores is not available, it defaults to 1.

pkg/xsysinfo/gpu.go
Package name: xsysinfo

Imports:
- "github.com/jaypipes/ghw"
- "github.com/jaypipes/ghw/pkg/gpu"

External data, input sources:
- The code uses the ghw package to access hardware information.

TODOs:
- There are no TODO comments in the code.

Summary:
The xsysinfo package provides a function called GPUs() that returns a list of graphics cards and any associated errors. It uses the ghw package to access hardware information and specifically targets the GPU component. The function first retrieves the GPU information using the ghw.GPU() function. If there is an error during this process, it returns nil and the error. Otherwise, it returns the list of graphics cards obtained from the GPU information and nil for the error.

The xsysinfo package provides functions to retrieve information about the system's CPU and GPU. The CPU-related functions allow you to get a list of CPU capabilities, check if the CPU supports specific features, and determine the number of physical cores. The GPU-related function enables you to obtain a list of graphics cards.