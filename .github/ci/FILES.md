# .github/ci/modelslist.go  
## Package/Component: LocalAI model gallery  
  
This package is responsible for displaying a gallery of available LocalAI models. It reads a YAML file containing model information, sanitizes the data, and then renders an HTML page with the model details.  
  
### External Data and Input Sources  
  
1. A YAML file containing model information is read as input. This file should contain a list of models, each with a name, URLs, icon, and description.  
  
### TODOs  
  
There are no TODO comments in the provided code.  
  
### Code Summary  
  
1. **Model Data Handling:**  
    - The code reads a YAML file containing model information.  
    - It unmarshals the YAML data into a slice of `GalleryModel` structs.  
    - It sanitizes the model names and descriptions using the `bluemonday` library to prevent cross-site scripting (XSS) vulnerabilities.  
  
2. **Template Rendering:**  
    - The code defines a template string containing the HTML structure for the model gallery page.  
    - It creates a `data` struct containing the sanitized model data and the number of available models.  
    - It uses the `template` package to parse the template string and execute it with the `data` struct.  
    - The rendered HTML is written to standard output.  
  
3. **Output:**  
    - The code outputs an HTML page containing a list of models with their names, descriptions, and icons.  
    - Each model has a "More info" button that, when clicked, displays a modal with additional information about the model, including its URLs.  
    - The page includes a live search feature that allows users to filter the list of models by keyword.  
  
4. **Dependencies:**  
    - The code depends on the `bluemonday`, `yaml`, and `template` packages.  
  
5. **Functionality:**  
    - The code provides a user-friendly interface for browsing and searching a collection of LocalAI models.  
    - It ensures that the displayed model data is safe from XSS vulnerabilities.  
    - It allows users to easily access detailed information about each model.  
  
# assets.go  
Package: main  
  
Imports:  
- "embed"  
  
External data, input sources:  
- backend-assets/* (embedded files)  
  
TODOs:  
- None  
  
Summary:  
This code defines a package named "main" that uses the "embed" package to embed files from the "backend-assets" directory into the binary. The embedded files can be accessed using the "backendAssets" variable, which is of type "embed.FS". This is useful for serving static assets without the need for a separate file system.  
  
  
  
# backend/go/bark/gobark.go  
Package/Component name: main  
  
Imports:  
- "C"  
- "fmt"  
- "unsafe"  
- "github.com/mudler/LocalAI/pkg/grpc/base"  
- "github.com/mudler/LocalAI/pkg/grpc/proto"  
  
External data, input sources:  
- Model file specified in the `opts.ModelFile` field of the `pb.ModelOptions` struct.  
- Text to be synthesized specified in the `opts.Text` field of the `pb.TTSRequest` struct.  
- Destination file path specified in the `opts.Dst` field of the `pb.TTSRequest` struct.  
  
TODOs:  
- None found.  
  
Summary:  
The `main` package provides a Go wrapper for the C++ Bark text-to-speech (TTS) engine. It defines a `Bark` struct that implements the `base.SingleThread` interface, allowing it to be used in a multi-threaded environment.  
  
The `Load` method of the `Bark` struct takes a `pb.ModelOptions` struct as input and loads the specified Bark model. It uses the CGO interface to call the `load_model` function from the C++ library.  
  
The `TTS` method takes a `pb.TTSRequest` struct as input and performs text-to-speech synthesis using the loaded Bark model. It uses the CGO interface to call the `tts` function from the C++ library.  
  
The package is designed to be used in a gRPC server, where the `Bark` struct can be used to handle TTS requests.  
  
  
  
# backend/go/bark/main.go  
Package/Component name: main  
  
Imports:  
- "flag"  
- "github.com/mudler/LocalAI/pkg/grpc"  
  
External data, input sources:  
- The code connects to a server at the address specified by the "addr" flag, which defaults to "localhost:50051".  
  
TODOs:  
- There are no TODO comments in the code.  
  
Summary:  
The main function of this code is to start a server for the Bark model. It does this by parsing command-line flags, starting a gRPC server, and registering the Bark model with the server. The server will then be accessible at the address specified by the "addr" flag.  
  
  
  
# backend/go/image/stablediffusion-ggml/gosd.go  
Package/Component name: main  
  
Imports:  
- "C"  
- "fmt"  
- "os"  
- "path/filepath"  
- "strings"  
- "unsafe"  
- "github.com/mudler/LocalAI/pkg/grpc/base"  
- "github.com/mudler/LocalAI/pkg/grpc/proto"  
- "github.com/mudler/LocalAI/pkg/utils"  
  
External data, input sources:  
- Model file path provided in the `opts.ModelFile` field of the `pb.ModelOptions` struct.  
- Options provided in the `opts.Options` field of the `pb.ModelOptions` struct.  
  
TODOs:  
- None found.  
  
Summary:  
- The code defines a struct `SDGGML` that implements the `base.SingleThread` interface.  
- The `Load` method of the `SDGGML` struct takes a `pb.ModelOptions` struct as input and loads the Stable Diffusion model.  
- The `GenerateImage` method of the `SDGGML` struct takes a `pb.GenerateImageRequest` struct as input and generates an image using the loaded model.  
- The code uses the CGO feature to interact with the C++ Stable Diffusion implementation.  
- The code handles the loading of the model, configuration, and generation of images.  
  
  
  
# backend/go/image/stablediffusion-ggml/main.go  
Package/Component name: main  
  
Imports:  
- "flag"  
- "github.com/mudler/LocalAI/pkg/grpc"  
  
External data, input sources:  
- The code connects to a server at the address specified by the "addr" flag, which defaults to "localhost:50051".  
  
TODOs:  
- There are no TODO comments in the code.  
  
Summary:  
The main function of this code is to start a server for the SDGGML model. It does this by parsing command-line flags, starting a gRPC server, and registering the SDGGML model with the server. The server will then be accessible at the address specified by the "addr" flag.  
  
  
  
# backend/go/llm/langchain/langchain.go  
Package name: main  
  
Imports:  
- "fmt"  
- "os"  
- "github.com/mudler/LocalAI/pkg/grpc/base"  
- "github.com/mudler/LocalAI/pkg/grpc/proto"  
- "github.com/mudler/LocalAI/pkg/langchain"  
  
External data, input sources:  
- HUGGINGFACEHUB_API_TOKEN environment variable  
  
TODOs:  
- None  
  
Summary:  
- The code defines a struct named LLM which implements the GRPC service interface. It uses the langchain package to interact with HuggingFace models.  
- The Load method initializes the LLM by creating a new HuggingFace instance and setting the model name.  
- The Predict method takes a prompt and prediction options as input and returns the prediction result.  
- The PredictStream method streams the prediction result to a channel.  
  
# backend/go/llm/langchain/main.go  
Package/Component name: main  
  
Imports:  
- "flag"  
- "github.com/mudler/LocalAI/pkg/grpc"  
  
External data, input sources:  
- The code connects to a server at the address specified by the "addr" flag, which defaults to "localhost:50051".  
  
TODOs:  
- There are no TODO comments in the code.  
  
Summary:  
The main function of this code is to start a server for a Large Language Model (LLM). It uses the grpc package to establish a connection to the server at the specified address. The server is started using the StartServer function, which takes the address and an instance of the LLM struct as arguments.  
  
  
  
# backend/go/llm/llama/llama.go  
## Package: main  
  
This package contains a wrapper for the GRPC service interface, designed to be used by the main executable that acts as the server for a specific backend type (falcon, gpt3, etc.).  
  
### Imports:  
  
- "fmt"  
- "path/filepath"  
- "github.com/go-skynet/go-llama.cpp"  
- "github.com/mudler/LocalAI/pkg/grpc/base"  
- "github.com/mudler/LocalAI/pkg/grpc/proto"  
  
### External Data and Input Sources:  
  
- The code relies on the "go-llama.cpp" library for interacting with the LLaMA language model.  
- It uses the "grpc/base" package for handling the base functionality of the GRPC service.  
- The "grpc/proto" package provides the protocol buffer definitions for communication.  
  
### TODOs:  
  
- There are no TODO comments in the provided code.  
  
### Code Summary:  
  
1. **LLM struct:**  
   - This struct represents the main LLM object, inheriting from the SingleThread struct from the "grpc/base" package.  
   - It contains two fields: "llama" and "draftModel," both of type *llama.LLama. These represent the main LLaMA model and an optional draft model, respectively.  
  
2. **Load function:**  
   - This function is responsible for loading the LLaMA model and setting up the necessary parameters based on the provided ModelOptions.  
   - It handles various options such as rope frequency base and scale, Lora adapter and base, context size, F16 memory, and more.  
   - The function returns an error if any issues occur during the loading process.  
  
3. **buildPredictOptions function:**  
   - This function takes PredictOptions as input and builds a list of llama.PredictOption objects.  
   - It sets various parameters for the prediction process, such as temperature, top P, top K, tokens, threads, grammar, rope frequency base and scale, and more.  
   - The function returns the list of llama.PredictOption objects.  
  
4. **Predict function:**  
   - This function performs the actual prediction using the loaded LLaMA model.  
   - It takes PredictOptions as input and uses the buildPredictOptions function to prepare the prediction options.  
   - If a draft model is available, it uses the speculative sampling method; otherwise, it uses the standard prediction method.  
   - The function returns the generated text and an error if any issues occur.  
  
5. **PredictStream function:**  
   - This function enables streaming of prediction results.  
   - It takes PredictOptions as input and uses the buildPredictOptions function to prepare the prediction options.  
   - The function sets up a token callback to send each generated token to the provided results channel.  
   - It then starts a goroutine to perform the prediction and close the results channel when finished.  
  
6. **Embeddings function:**  
   - This function calculates the embeddings for the given input text.  
   - It takes PredictOptions as input and uses the buildPredictOptions function to prepare the prediction options.  
   - If embedding tokens are provided, it uses the TokenEmbeddings method; otherwise, it uses the Embeddings method.  
   - The function returns the calculated embeddings and an error if any issues occur.  
  
7. **TokenizeString function:**  
   - This function tokenizes the input string and returns the length and tokens.  
   - It takes PredictOptions as input and uses the buildPredictOptions function to prepare the prediction options.  
   - The function returns a TokenizationResponse object containing the length and tokens, and an error if any issues occur.  
  
### Summary:  
  
This package provides a wrapper for the GRPC service interface, allowing users to interact with the LLaMA language model. It handles loading the model, setting up various parameters, and performing predictions, embeddings, and tokenization tasks. The code is well-structured and documented, making it easy to understand and use.  
  
# backend/go/llm/llama/main.go  
Package/Component name: main  
  
Imports:  
- "flag"  
- "github.com/mudler/LocalAI/pkg/grpc"  
  
External data, input sources:  
- The code uses the address provided by the flag "addr" to connect to the server.  
  
TODOs:  
- There are no TODO comments in the code.  
  
Summary:  
The main function of this code is to start a GRPC server for the LocalAI project. It uses the address provided by the flag "addr" to connect to the server. The server is allocated for each model and is started internally by LocalAI.  
  
  
  
# backend/go/stores/debug.go  
Package: main  
  
Imports:  
github.com/rs/zerolog/log  
  
External data, input sources:  
None  
  
TODOs:  
None  
  
Summary:  
The code defines a function called assert that takes a boolean condition and a message string as input. If the condition is false, the function logs a fatal error with the provided message and stack trace. This function is likely used for debugging purposes, as it is only compiled in debug builds.  
  
  
  
# backend/go/stores/main.go  
Package/Component name: main  
  
Imports:  
- "flag"  
- "os"  
- "github.com/mudler/LocalAI/pkg/grpc"  
- "github.com/rs/zerolog"  
- "github.com/rs/zerolog/log"  
  
External data, input sources:  
- The code connects to a server at the address specified by the "addr" flag, which defaults to "localhost:50051".  
  
TODOs:  
- There are no TODO comments in the code.  
  
Summary of major code parts:  
- The code starts a gRPC server for each store.  
- It uses the "grpc" package to start the server and the "NewStore()" function to create a new store instance.  
- The server address is configurable via the "addr" flag.  
- The code sets up logging using the "zerolog" package.  
  
  
  
# backend/go/stores/production.go  
Package: main  
  
Imports:  
- none  
  
External data, input sources:  
- none  
  
TODOs:  
- none  
  
Summary:  
- The code defines a function called assert that takes a boolean condition and a string message as input. If the condition is false, the function will print the message to the console. This function is useful for debugging purposes, as it allows developers to check the values of variables and the flow of execution in their code.  
  
- The code also includes a build constraint that prevents the assert function from being compiled when the debug flag is enabled. This ensures that the assert function is only used in non-debug builds, which can help to improve the performance of the application.  
  
- The assert function is a simple but effective way to add debugging information to a Go program. By using the assert function, developers can quickly identify and fix issues in their code.  
  
# backend/go/stores/store.go  
## Package: main  
  
This package contains a wrapper to satisfy the GRPC service interface. It is meant to be used by the main executable that is the server for the specific backend type (falcon, gpt3, etc).  
  
### External Data and Input Sources:  
  
- The code interacts with a gRPC service, which is an external data source.  
- It receives input from the gRPC service in the form of key-value pairs, where the keys are vectors of float32 and the values are byte slices.  
  
### TODOs:  
  
- TODO: Only used for sorting using Go's builtin implementation. The interfaces are columnar because that's theoretically best for memory layout and cache locality, but this isn't optimized yet.  
- TODO: This we could replace with handwritten SIMD code  
- TODO: Should we normalize incoming keys if they are not instead?  
  
### Summary of Major Code Parts:  
  
1. **Store struct:** This struct represents the in-memory store for key-value pairs. It contains the sorted keys and values, as well as information about the keys, such as whether they are normalized and their length.  
2. **Pair struct:** This struct is used to store a key-value pair. It is used for sorting the keys and values.  
3. **compareSlices function:** This function compares two slices of float32 values.  
4. **hasKey function:** This function checks if a given key is present in a slice of keys.  
5. **findInSortedSlice function:** This function searches for a given key in a sorted slice of keys.  
6. **isSortedPairs function:** This function checks if a slice of key-value pairs is sorted by key.  
7. **isSortedKeys function:** This function checks if a slice of keys is sorted.  
8. **sortIntoKeySlicese function:** This function sorts a slice of keys into a slice of float32 values.  
9. **Load function:** This function loads the store from a data source.  
10. **StoresSet function:** This function adds a set of key-value pairs to the store.  
11. **StoresDelete function:** This function deletes a set of keys from the store.  
12. **StoresGet function:** This function retrieves a set of key-value pairs from the store.  
13. **isNormalized function:** This function checks if a key is normalized.  
14. **normalizedCosineSimilarity function:** This function calculates the normalized cosine similarity between two keys.  
15. **PriorityQueue struct:** This struct is used to store a priority queue of key-value pairs.  
16. **StoresFindNormalized function:** This function finds the top-k most similar keys to a given key, assuming the keys are normalized.  
17. **cosineSimilarity function:** This function calculates the cosine similarity between two keys.  
18. **StoresFindFallback function:** This function finds the top-k most similar keys to a given key, even if the keys are not normalized.  
19. **StoresFind function:** This function finds the top-k most similar keys to a given key, using either the normalized or fallback method depending on the keys in the store.  
  
### Summary:  
  
The code implements a key-value store that can be used to store and retrieve data. It supports adding, deleting, and retrieving key-value pairs, as well as finding the most similar keys to a given key. The store is designed to be used in a gRPC service, and it interacts with the service to receive input and return results.  
  
# backend/go/transcribe/whisper/main.go  
Package/Component name: main  
  
Imports:  
- "flag"  
- "github.com/mudler/LocalAI/pkg/grpc"  
  
External data, input sources:  
- The code connects to a server at the address specified by the "addr" flag, which defaults to "localhost:50051".  
  
TODOs:  
- There are no TODO comments in the code.  
  
Summary:  
The main function of this code is to start a server for the Whisper model. It does this by parsing command-line flags, starting a gRPC server, and registering the Whisper model with the server. The server will then be accessible at the address specified by the "addr" flag.  
  
  
  
# backend/go/transcribe/whisper/whisper.go  
Package/Component name: main  
  
Imports:  
os  
path/filepath  
github.com/ggerganov/whisper.cpp/bindings/go/pkg/whisper  
github.com/go-audio/wav  
github.com/mudler/LocalAI/pkg/grpc/base  
github.com/mudler/LocalAI/pkg/grpc/proto  
github.com/mudler/LocalAI/pkg/utils  
  
External data, input sources:  
- Model files for Whisper, specified by the ModelFile field in the ModelOptions message.  
- Audio files to be transcribed, provided as input to the AudioTranscription method.  
  
TODOs:  
- None found.  
  
Summary:  
The Whisper component is a wrapper that implements the GRPC service interface for the Whisper speech recognition model. It is designed to be used by the main executable, which acts as the server for a specific backend type.  
  
The component provides two main functionalities:  
1. Loading the Whisper model: The Load method takes a ModelOptions message as input, which contains the path to the directory containing the model files. It initializes a Whisper model instance and returns any errors encountered during the process.  
2. Audio transcription: The AudioTranscription method takes a TranscriptRequest message as input, which contains the path to the audio file to be transcribed, the desired language, and other transcription options. It converts the audio file to WAV format, processes the audio samples using the Whisper model, and returns the transcription result as a TranscriptResult message.  
  
The Whisper component relies on several external libraries, including the Whisper C++ bindings, the Go WAV library, and the LocalAI GRPC base and proto packages. It also uses the utils package for audio conversion and other utility functions.  
  
# backend/go/tts/main.go  
Package/Component name: main  
  
Imports:  
- "flag"  
- "github.com/mudler/LocalAI/pkg/grpc"  
  
External data, input sources:  
- The code connects to a server at the address specified by the "addr" flag, which defaults to "localhost:50051".  
  
TODOs:  
- There are no TODO comments in the code.  
  
Summary:  
The main function of this code is to start a server for the Piper component. It takes an address as input and uses the grpc package to start the server. The Piper component is responsible for handling requests from clients and interacting with the LocalAI system.  
  
  
  
# backend/go/tts/piper.go  
Package/Component name: main  
  
Imports:  
"fmt"  
"os"  
"path/filepath"  
"github.com/mudler/LocalAI/pkg/grpc/base"  
"github.com/mudler/LocalAI/pkg/grpc/proto"  
"github.com/mudler/go-piper"  
  
External data, input sources:  
- Model files in ONNX format, specified by the ModelFile field in the ModelOptions message.  
- Library search path, specified by the LibrarySearchPath field in the ModelOptions message.  
  
TODOs:  
- None found.  
  
Summary:  
The main package provides a wrapper for the GRPC service interface, specifically designed for backends like Falcon and GPT3. It implements the SingleThread interface from the base package and uses the PiperB struct to handle text-to-speech (TTS) requests.  
  
The Piper struct is responsible for loading the model and handling TTS requests. It first checks if the model file has the .onnx extension. If not, it returns an error. Then, it initializes the PiperB struct with the provided library search path.  
  
The PiperB struct is responsible for the actual TTS process. It takes the input text, model, and destination path as arguments and uses the piper.TextToWav function to generate the speech output.  
  
The package also includes a New function that creates a new PiperB instance and checks if the asset directory exists.  
  
This package is designed to be used by the main executable that serves as the backend server for various AI models. It provides a simple and efficient way to handle TTS requests and interact with the underlying Piper library.  
  
# backend/go/vad/silero/main.go  
Package/Component name: main  
  
Imports:  
- "flag"  
- "github.com/mudler/LocalAI/pkg/grpc"  
  
External data, input sources:  
- The code connects to a server at the address specified by the "addr" flag, which defaults to "localhost:50051".  
  
TODOs:  
- There are no TODO comments in the code.  
  
Summary:  
The main function of this code is to start a server for a Voice Activity Detection (VAD) model. It uses the grpc package to establish a connection to the server and then starts the server. The server address is configurable via the "addr" flag, which defaults to "localhost:50051".  
  
  
  
# backend/go/vad/silero/vad.go  
Package/Component name: main  
  
Imports:  
- "fmt"  
- "github.com/mudler/LocalAI/pkg/grpc/base"  
- "github.com/mudler/LocalAI/pkg/grpc/proto"  
- "github.com/streamer45/silero-vad-go/speech"  
  
External data, input sources:  
- Model file path provided in the ModelOptions struct.  
- Audio data provided in the VADRequest struct.  
  
TODOs:  
- None found.  
  
Summary:  
- The code defines a VAD struct that implements the SingleThread interface from the base package.  
- The Load method initializes the speech detector using the provided model file path and configuration.  
- The VAD method takes an audio input and returns a list of VAD segments, indicating the start and end times of speech detected in the audio.  
- The VAD struct is designed to be used as a GRPC service for voice activity detection.  
  
# core/dependencies_manager/manager.go  
Package/Component name: main  
  
Imports:  
- "fmt"  
- "os"  
- "path/filepath"  
- "github.com/mudler/LocalAI/pkg/downloader"  
- "github.com/mudler/LocalAI/pkg/utils"  
- "gopkg.in/yaml.v3"  
  
External data, input sources:  
- YAML file containing a list of assets (filename, URL, SHA)  
- Destination path for downloading assets  
  
TODOs:  
- None  
  
Summary:  
The main function reads a YAML file containing a list of assets, downloads them to the specified destination path, and prints a confirmation message. The YAML file contains a list of Asset structs, each containing the filename, URL, and SHA of the asset. The code iterates through the list of assets, downloads each asset using the downloader package, and handles any errors that occur during the download process.  
  
- The code first reads the YAML file and unmarshals the data into a slice of Asset structs.  
- Then, it iterates through the slice of Asset structs and downloads each asset using the downloader package.  
- Finally, it prints a confirmation message indicating that the assets have been downloaded.  
  
This code is responsible for downloading assets required by the LocalAI package. It reads a YAML file containing a list of assets, downloads them to the specified destination path, and handles any errors that occur during the download process.  
  
# main.go  
Package/Component name: main  
  
Imports:  
os  
os/signal  
path/filepath  
syscall  
github.com/alecthomas/kong  
github.com/joho/godotenv  
github.com/mudler/LocalAI/core/cli  
github.com/mudler/LocalAI/internal  
github.com/rs/zerolog  
github.com/rs/zerolog/log  
github.com/mudler/LocalAI/swagger  
  
External data, input sources:  
- Environment variables loaded from .env files: .env, localai.env, ~/.config/localai.env, /etc/localai.env  
  
TODOs:  
-  
  
Summary:  
- The code initializes zerolog at an INFO level and sets up a signal handler to catch OS exit requests.  
- It loads environment variables from .env files and parses CLI options using the kong library.  
- The code configures the logging level based on the CLI options and populates the application with embedded backend assets.  
- Finally, it runs the application and handles any errors.  
  
