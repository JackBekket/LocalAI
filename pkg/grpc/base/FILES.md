# pkg/grpc/base/base.go  
Package: base  
  
Imports:  
- "fmt"  
- "os"  
- "github.com/mudler/LocalAI/pkg/grpc/proto"  
- "github.com/shirou/gopsutil/v3/process"  
  
External data, input sources:  
- The code uses the gopsutil library to get memory usage information.  
  
TODOs:  
- There are no TODO comments in the code.  
  
Summary:  
- The Base struct is a base class for all backends to implement. It provides default implementations for all the methods required by the GRPC service interface.  
- The memoryUsage function calculates the memory usage of the backend process and returns a MemoryUsageData object.  
- The Status method returns the memory usage of the backend process.  
- The other methods are not implemented and return an error.  
  
# pkg/grpc/base/singlethread.go  
Package: base  
  
Imports:  
- "sync"  
- "github.com/mudler/LocalAI/pkg/grpc/proto"  
  
External data, input sources:  
- The code uses the memoryUsage() function, which is not provided in the given code snippet. It's assumed that this function is defined elsewhere and provides information about memory usage.  
  
TODOs:  
- There are no TODO comments in the provided code.  
  
Summary:  
The code defines a SingleThread struct, which represents a backend that does not support multiple requests. This struct inherits from the Base struct (not shown in the provided code) and implements the Locking(), Lock(), Unlock(), and Busy() methods. The Locking() method returns true, indicating that the backend needs to lock resources. The Lock() and Unlock() methods are used to acquire and release the lock, respectively. The Busy() method checks if the backend is currently busy by attempting to acquire the lock. If the lock is acquired, it is immediately released, and the method returns true. Otherwise, it returns false.  
  
The SingleThread struct also implements the Status() method, which returns a StatusResponse object containing information about the backend's state and memory usage. The state is determined by checking if the backend is busy using the Busy() method. The memory usage is obtained by calling the memoryUsage() function.  
  
The SingleThread struct is designed for models that are not thread-safe and cannot run multiple requests concurrently. It ensures that only one request is being served at a time by using a mutex to protect shared resources.  
  
  
  
