# core/http/utils/baseurl.go  
Package/Component name: utils  
  
Imports:  
- "strings"  
- "github.com/gofiber/fiber/v2"  
  
External data, input sources:  
- HTTP request context  
  
TODOs:  
- None  
  
Summary:  
The `utils` package provides utility functions for working with HTTP requests. The `BaseURL` function is designed to determine the base URL for a given HTTP request context, taking into account the possibility of the application being exposed by a reverse-proxy under a different protocol, host, and path. It ensures that the returned URL ends with a forward slash and should be used in conjunction with the StripPathPrefix middleware.  
  
The function works by comparing the original URL with the current path. If the original URL contains the current path as a suffix, it extracts the path prefix and appends it to the base URL. Otherwise, it simply returns the base URL with a trailing slash. This approach ensures that the returned URL is consistent and accurate, regardless of the application's deployment environment.  
  
# core/http/utils/baseurl_test.go  
Package: utils  
  
Imports:  
- "net/http/httptest"  
- "testing"  
- "github.com/gofiber/fiber/v2"  
- "github.com/stretchr/testify/require"  
  
External data, input sources:  
- None  
  
TODOs:  
- None  
  
Summary:  
- The `utils` package provides utility functions for working with Fiber applications.  
- The `TestBaseUrl` function tests the `BaseUrl` function, which returns the base URL of a Fiber context.  
- The test cases cover scenarios with and without a prefix in the URL.  
- The test function sets up a Fiber application, registers a test route, and asserts that the `BaseUrl` function returns the expected base URL.  
  
  
  
# pkg/utils/base64.go  
Package/Component name: utils  
  
Imports:  
- "encoding/base64"  
- "fmt"  
- "io"  
- "net/http"  
- "strings"  
- "time"  
  
External data, input sources:  
- The code downloads content from URLs provided as input.  
  
TODOs:  
- There are no TODO comments in the code.  
  
Summary of major code parts:  
- The `GetContentURIAsBase64` function takes a string as input and checks if it's a URL. If it is, it downloads the content, encodes it in base64, and returns the base64 string. If the string is not a URL but starts with a specific prefix, it removes the prefix and returns the remaining string. Otherwise, it returns an error.  
  
- The `base64DownloadClient` variable is a http.Client with a timeout of 30 seconds, used for downloading content from URLs.  
  
- The code uses the `io.ReadAll` function to read the downloaded content into memory.  
  
- The code uses the `base64.StdEncoding.EncodeToString` function to encode the downloaded content in base64.  
  
- The code uses the `strings.HasPrefix` function to check if the input string starts with a specific prefix.  
  
- The code uses the `strings.ReplaceAll` function to remove the prefix from the input string.  
  
- The code uses the `fmt.Errorf` function to create an error if the input string is not valid.  
  
- The code uses the `http.Client.Get` function to download content from URLs.  
  
- The code uses the `time.Second` constant to define the timeout for the http.Client.  
  
- The code uses the `defer` keyword to ensure that the response body is closed after it's been read.  
  
- The code uses the `strings.HasPrefix` function to check if the input string starts with a specific prefix.  
  
- The code uses the `strings.ReplaceAll` function to remove the prefix from the input string.  
  
- The code uses the `fmt.Errorf` function to create an error if the input string is not valid.  
  
- The code uses the `http.Client.Get` function to download content from URLs.  
  
- The code uses the `time.Second` constant to define the timeout for the http.Client.  
  
- The code uses the `defer` keyword to ensure that the response body is closed after it's been read.  
  
# pkg/utils/config.go  
Package: utils  
  
Imports:  
- "encoding/json"  
- "os"  
- "path/filepath"  
- "github.com/rs/zerolog/log"  
  
External data, input sources:  
- Configuration files are loaded and saved from the file system.  
  
TODOs:  
- There are no TODO comments in the code.  
  
Summary:  
The utils package provides functions for saving and loading configuration files. The SaveConfig function takes a file path, file name, and an object as input, and saves the object's JSON representation to the specified file. The LoadConfig function takes a file path, file name, and an object as input, and loads the JSON data from the specified file into the object.  
  
The package uses the encoding/json package to handle JSON encoding and decoding, the os package to interact with the file system, and the path/filepath package to work with file paths. The github.com/rs/zerolog/log package is used for logging.  
  
The package is designed to be used by other components that need to store and retrieve configuration data. It provides a simple and efficient way to manage configuration files, making it easy for developers to work with.  
  
# pkg/utils/ffmpeg.go  
Package: utils  
  
Imports:  
- "fmt"  
- "os"  
- "os/exec"  
- "strings"  
  
External data, input sources:  
- Audio files to be converted  
  
TODOs:  
- In AudioToWav function: use https://github.com/mccoyst/ogg?  
- In AudioConvert function: handle pcm to have 100% parity of supported format from OpenAI  
  
Summary:  
- The `ffmpegCommand` function executes ffmpeg commands with the given arguments and returns the output and any errors.  
- The `AudioToWav` function converts audio files to the WAV format using ffmpeg.  
- The `AudioConvert` function converts WAV files to other audio formats (opus, mp3, aac, flac) using ffmpeg.  
  
  
  
# pkg/utils/hash.go  
Package: utils  
  
Imports:  
- "crypto/md5"  
- "fmt"  
  
External data, input sources:  
- The code takes a string as input.  
  
TODOs:  
- There are no TODO comments in the code.  
  
Summary:  
- The `MD5` function takes a string as input and returns its MD5 hash as a hexadecimal string.  
  
  
  
# pkg/utils/json.go  
## Package: utils  
  
Imports:  
- "regexp"  
  
External data, input sources:  
- None  
  
TODOs:  
- None  
  
### Function: EscapeNewLines  
  
This function takes a string as input and replaces all newline characters with the escape sequence "\n" within any double-quoted strings. It uses regular expressions to find double-quoted strings and then replaces any newline characters within those strings with the escape sequence.  
  
  
  
# pkg/utils/logging.go  
Package/Component name: utils  
  
Imports:  
- "time"  
- "github.com/rs/zerolog/log"  
  
External data, input sources:  
- None  
  
TODOs:  
- None  
  
Summary:  
The utils package provides utility functions for handling download progress and displaying download information. The ResetDownloadTimers function resets the timers used to track download progress. The DisplayDownloadFunction function displays the download progress information, including the file name, current and total size, percentage, and estimated time of arrival (ETA). The function calculates the ETA based on the percentage of the download completed and the elapsed time.  
  
  
  
# pkg/utils/path.go  
## Package: utils  
  
Imports:  
- "fmt"  
- "os"  
- "path/filepath"  
- "strings"  
  
External data, input sources:  
- The code interacts with the file system to check for the existence of files and directories.  
  
TODOs:  
- There are no TODO comments in the code.  
  
### File system interaction functions  
  
This section covers functions that interact with the file system.  
  
- `ExistsInPath(path string, s string) bool`: This function checks if a file or directory exists in a given path.  
- `InTrustedRoot(path string, trustedRoot string) error`: This function checks if a given path is within a trusted root directory.  
- `VerifyPath(path, basePath string) error`: This function verifies that a given path is based in a base path.  
- `SanitizeFileName(fileName string) string`: This function sanitizes a given file name to ensure it is safe to use.  
- `GenerateUniqueFileName(dir, baseName, ext string) string`: This function generates a unique file name in a given directory.  
  
These functions are useful for handling file system operations in a secure and reliable manner.  
  
# pkg/utils/strings.go  
Package name: utils  
  
Imports:  
math/rand  
time  
  
External data, input sources:  
None  
  
TODOs:  
None  
  
Summary:  
The utils package provides utility functions for generating random strings and finding unique elements in a string array.  
  
The RandString function generates a random string of a given length. It uses the math/rand package to generate random numbers, which are then used to select characters from a set of letters (both uppercase and lowercase).  
  
The Unique function takes a string array as input and returns a new array containing only the unique elements from the input array. It uses a map to keep track of the elements that have already been seen, and only adds new elements to the result array if they are not already present in the map.  
  
# pkg/utils/untar.go  
## Package: utils  
  
Imports:  
- "fmt"  
- "os"  
- "github.com/mholt/archiver/v3"  
  
External data, input sources:  
- The code interacts with files on the filesystem.  
  
TODOs:  
- There are no TODO comments in the code.  
  
### Archive Handling  
  
This section of the code provides functions for working with archives. The `IsArchive` function checks if a given file is an archive based on its extension. The `ExtractArchive` function extracts the contents of an archive to a specified destination directory. It handles various archive formats and ensures that any symlinks within the archive are not extracted.  
  
### Summary  
  
The utils package provides utility functions for handling archives, including checking if a file is an archive and extracting its contents. It supports various archive formats and includes error handling for symlinks within archives.  
  
