# pkg/startup/model_preload.go  
Package: startup  
  
Imports:  
- "errors"  
- "fmt"  
- "os"  
- "path/filepath"  
- "strings"  
- "github.com/mudler/LocalAI/core/config"  
- "github.com/mudler/LocalAI/core/gallery"  
- "github.com/mudler/LocalAI/pkg/downloader"  
- "github.com/mudler/LocalAI/pkg/utils"  
- "github.com/rs/zerolog/log"  
  
External data, input sources:  
- Galleries: A list of galleries to search for models.  
- Model path: The directory where models should be stored.  
- Enforce scan: A boolean indicating whether to force a scan for models.  
- Download status: A function to display the progress of model downloads.  
- Models: A list of model URLs or local file paths.  
  
TODOs:  
- None  
  
Code summary:  
- InstallModels: This function takes a list of model URLs or local file paths, a model path, and a download status function. It downloads and installs the models, handling both OCI and URL-based models. If a model is not found in the given galleries, it will attempt to download it from the remote library.  
- installModel: This function searches for a model in the given galleries and installs it if found. It takes a list of galleries, a model name, a model path, a download status function, and a boolean indicating whether to force a scan for models. If the model is found, it will be installed. If not, it will return an error.  
  
  
  
