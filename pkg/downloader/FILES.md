# pkg/downloader/downloader_suite_test.go  
Package: downloader  
  
Imports:  
- "testing"  
- "github.com/onsi/ginkgo/v2"  
- "github.com/onsi/gomega"  
  
External data, input sources:  
- None  
  
TODOs:  
- None  
  
Summary:  
- The code defines a test suite for the downloader package. It registers a fail handler and runs the test specs.  
  
  
  
# pkg/downloader/huggingface.go  
Package name: downloader  
  
Imports:  
encoding/json  
errors  
fmt  
io  
net/http  
strings  
  
External data, input sources:  
The code interacts with the HuggingFace API to scan models for unsafe files. It uses the URI of the model to fetch scan results from the API.  
  
TODOs:  
There are no TODO comments in the provided code.  
  
Summary of major code parts:  
The code defines a HuggingFaceScanResult struct to store the results of the scan. It also defines two error variables: ErrNonHuggingFaceFile and ErrUnsafeFilesFound. The HuggingFaceScan function takes a URI as input and returns a HuggingFaceScanResult and an error. It first checks if the URI points to a HuggingFace repository. If it does, it fetches the scan results from the API and parses the JSON response. If the scan results indicate that unsafe files were found, it returns the HuggingFaceScanResult and the ErrUnsafeFilesFound error. Otherwise, it returns the HuggingFaceScanResult and a nil error.  
  
  
  
# pkg/downloader/progress.go  
Package: downloader  
  
Imports:  
hash  
  
External data, input sources:  
- The code uses a hash.Hash object to calculate the hash of the downloaded data.  
  
TODOs:  
- There are no TODO comments in the code.  
  
Summary:  
The code defines a progressWriter struct that implements the io.Writer interface. This struct is used to track the progress of a download and update a downloadStatus function with the current progress. The progressWriter struct keeps track of the file name, total size, current progress, and a hash object to calculate the hash of the downloaded data. The Write method of the progressWriter struct updates the hash object with the data being written and calculates the percentage of the download that has been completed. If the download is a multi-file download, the percentage is adjusted to reflect the progress of the whole download. The downloadStatus function is then called with the file name, the amount of data written, the total size of the download, and the percentage of the download that has been completed.  
  
The code also includes a formatBytes function that is used to format the size of the data being downloaded in bytes, kilobytes, megabytes, or gigabytes.  
  
The progressWriter struct and the formatBytes function are useful for tracking the progress of a download and providing feedback to the user.  
  
# pkg/downloader/uri.go  
## Package: downloader  
  
This package provides functionality for downloading files from various sources, including URLs, GitHub repositories, and HuggingFace. It also includes support for downloading OCI and Ollama images.  
  
### External Data and Input Sources  
  
The package relies on external data and input sources such as:  
  
1. URLs: The package downloads files from URLs provided as input.  
2. GitHub repositories: The package can download files from GitHub repositories using the repository URL and branch information.  
3. HuggingFace repositories: The package can download files from HuggingFace repositories using the repository URL and branch information.  
4. OCI and Ollama images: The package can download OCI and Ollama images using the image URL.  
  
### TODOs  
  
There are no TODO comments in the provided code.  
  
### Code Summary  
  
1. **URI struct:** This struct represents a URL and provides methods for downloading files, checking if the URL looks like a specific type (e.g., OCI, Ollama), and resolving the URL to its full form.  
2. **DownloadWithCallback and DownloadWithAuthorizationAndCallback methods:** These methods download a file from a given URL and execute a callback function with the downloaded data.  
3. **FilenameFromUrl method:** This method extracts the filename from a given URL.  
4. **LooksLikeURL and LooksLikeOCI methods:** These methods check if a given URL looks like a specific type (e.g., OCI, Ollama).  
5. **ResolveURL method:** This method resolves a given URL to its full form.  
6. **removePartialFile and calculateHashForPartialFile methods:** These methods handle partial downloads and calculate the hash of the downloaded data.  
7. **checkSeverSupportsRangeHeader method:** This method checks if the server supports the Range header for partial downloads.  
8. **DownloadFile method:** This method downloads a file from a given URL, handles partial downloads, and verifies the SHA hash of the downloaded file.  
9. **formatBytes and calculateSHA methods:** These methods format the size of a file and calculate the SHA hash of a given file.  
  
The package provides a comprehensive set of tools for downloading files from various sources, handling partial downloads, and verifying the integrity of the downloaded data.  
  
