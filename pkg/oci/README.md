# pkg/oci

This package provides functions for working with OCI (Open Container Initiative) images and interacting with the Ollama registry.

pkg/oci/blob.go
Package: oci

Imports:
	"context"
	"fmt"
	"io"
	"os"
	"github.com/opencontainers/image-spec/specs-go/v1"
	"oras.land/oras-go/v2"
	"oras.land/oras-go/v2/registry/remote"

External data, input sources:
	- The code interacts with a remote repository to fetch image blobs.
	- It takes the repository URL, image reference, and destination path as input.

TODOs:
	- There are no TODO comments in the code.

Summary:
	- The code defines a function called FetchImageBlob that fetches an image blob from a remote repository.
	- It first creates a file store for the output and connects to the remote repository.
	- Then, it uses the oras library to fetch the image blob and writes it to the file store.
	- The function also takes an optional status reader to track the progress of the download.

pkg/oci/image.go
Package: oci

Imports:
- context
- errors
- fmt
- io
- net/http
- runtime
- strings
- syscall
- time
- github.com/containerd/containerd/archive
- github.com/docker/docker/api/types/registry
- github.com/google/go-containerregistry/pkg/authn
- github.com/google/go-containerregistry/pkg/logs
- github.com/google/go-containerregistry/pkg/name
- github.com/google/go-containerregistry/pkg/v1
- github.com/google/go-containerregistry/pkg/v1/mutate
- github.com/google/go-containerregistry/pkg/v1/remote
- github.com/google/go-containerregistry/pkg/v1/remote/transport

External data, input sources:
- The code interacts with container images from Docker Hub or other registries.

TODOs:
- There are no TODO comments in the code.

Summary:
- The package provides functions for working with OCI (Open Container Initiative) images.
- The ExtractOCIImage function extracts a given targetImage into a given targetDestination.
- The ParseImageParts function parses an image string and returns the tag, repository, and dstimage.
- The GetImage function retrieves an image from a registry, either locally or remotely.
- The GetOCIImageSize function returns the size of an OCI image.

- The code includes a custom authentication mechanism for Docker registries.
- It also implements retry logic for handling network errors during image retrieval.
- The package is designed to be used in environments where OCI images are commonly used, such as container orchestration platforms.

pkg/oci/ollama.go
Package: oci

Imports:
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"github.com/opencontainers/image-spec/specs-go/v1"

External data, input sources:
	- The code interacts with the Ollama registry (https://registry.ollama.ai) to fetch model manifests and blobs.

TODOs:
	- There are no TODO comments in the provided code.

Summary:
	- The code defines a Manifest struct to represent the JSON data of an Ollama model manifest. It includes the schema version, media type, configuration details, and layer information.
	- The OllamaModelManifest function takes an image string as input and returns the Manifest struct for the corresponding model. It parses the repository and tag from the image string, constructs a request to the Ollama registry, and decodes the JSON response into the Manifest struct.
	- The OllamaModelBlob function takes an image string as input and returns the digest of the model blob. It first retrieves the Manifest struct using OllamaModelManifest and then searches for the layer with the media type "application/vnd.ollama.image.model".
	- The OllamaFetchModel function takes an image string and an output path as input and downloads the model blob to the specified output path. It first retrieves the blob ID using OllamaModelBlob and then uses the FetchImageBlob function to download the blob from the Ollama registry.

pkg/oci/oci_suite_test.go
Package: oci

Imports:
	"testing"
	"github.com/stretchr/testify/suite"
	"github.com/opencontainers/image-spec/specs-go/v1"
	"oras.land/oras-go/v2"
	"oras.land/oras-go/v2/registry/remote"

External data, input sources:
	- The code interacts with a remote repository to fetch image blobs.
	- It takes the repository URL, image reference, and destination path as input.

TODOs:
	- There are no TODO comments in the code.

Summary:
	- The code defines a suite of tests for the OCI package.
	- It tests the FetchImageBlob function by fetching an image blob from a remote repository and verifying that the downloaded blob matches the expected content.

pkg/oci/ollama_test.go
Package: oci

Imports:
	"testing"
	"github.com/stretchr/testify/suite"
	"github.com/opencontainers/image-spec/specs-go/v1"
	"oras.land/oras-go/v2"
	"oras.land/oras-go/v2/registry/remote"

External data, input sources:
	- The code interacts with the Ollama registry (https://registry.ollama.ai) to fetch model manifests and blobs.

TODOs:
	- There are no TODO comments in the provided code.

Summary:
	- The code defines a suite of tests for the Ollama package.
	- It tests the OllamaModelManifest, OllamaModelBlob, and OllamaFetchModel functions by fetching model manifests and blobs from the Ollama registry and verifying that the retrieved data matches the expected content.

