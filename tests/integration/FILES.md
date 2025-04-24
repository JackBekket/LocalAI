# tests/integration/integration_suite_test.go  
Package: integration_test  
  
Imports:  
- "os"  
- "testing"  
- "github.com/onsi/ginkgo/v2"  
- "github.com/onsi/gomega"  
- "github.com/rs/zerolog"  
- "github.com/rs/zerolog/log"  
  
External data, input sources:  
- None  
  
TODOs:  
- None  
  
Summary:  
The integration_test package contains a single test suite for the LocalAI component. The test suite is executed using the Ginkgo testing framework and the Gomega assertion library. The test suite is configured to output logs to the console using the zerolog logging library.  
  
The test suite is designed to test the LocalAI component in a local environment. It is executed by running the `go test` command in the integration_test package directory. The test suite will run all the tests defined in the package and report any failures.  
  
The test suite is a valuable tool for ensuring the quality and stability of the LocalAI component. By running the test suite regularly, developers can catch issues early and prevent them from impacting users.  
  
  
  
# tests/integration/stores_test.go  
## Package: integration_test  
  
This package contains integration tests for the stores backend(s) and internal APIs.  
  
### Imports:  
  
- context  
- embed  
- math  
- math/rand  
- os  
- path/filepath  
- github.com/onsi/ginkgo/v2  
- github.com/onsi/gomega  
- github.com/rs/zerolog  
- github.com/mudler/LocalAI/core/config  
- github.com/mudler/LocalAI/pkg/assets  
- github.com/mudler/LocalAI/pkg/grpc  
- github.com/mudler/LocalAI/pkg/model  
- github.com/mudler/LocalAI/pkg/store  
  
### External Data and Input Sources:  
  
- The tests use a temporary directory to store the backend assets.  
- The tests use a mock store backend to simulate the behavior of the actual backend.  
  
### TODOs:  
  
- None found.  
  
### Summary:  
  
The package contains integration tests for the stores backend(s) and internal APIs. The tests cover various aspects of the backend, including setting and getting keys, deleting keys, finding similar keys, and ensuring the triangle inequality holds for normalized values. The tests use a temporary directory to store the backend assets and a mock store backend to simulate the behavior of the actual backend.  
  
