# core/cli/worker/worker.go  
Package/Component name: worker  
  
Imports:  
- P2P  
- LLamaCPP  
  
External data, input sources:  
- Environment variables: LOCALAI_BACKEND_ASSETS_PATH, BACKEND_ASSETS_PATH, LOCALAI_EXTRA_LLAMA_CPP_ARGS, EXTRA_LLAMA_CPP_ARGS  
  
TODOs:  
- None  
  
Summary:  
The worker package provides a way to start a LocalAI llama.cpp worker in either P2P mode or standalone mode. It defines a WorkerFlags struct to store configuration options, such as the path to backend assets and extra arguments for llama-cpp-rpc-server. The Worker struct contains two fields, P2P and LLamaCPP, which represent the two modes of operation.  
  
- P2P mode: Starts a LocalAI llama.cpp worker in P2P mode, which requires a token.  
- Standalone mode: Starts a llama.cpp worker in standalone mode.  
  
The package allows users to easily configure and start a llama.cpp worker for their specific needs.  
  
# core/cli/worker/worker_llamacpp.go  
Package/Component name: worker  
  
Imports:  
"fmt"  
"os"  
"strings"  
"syscall"  
"github.com/mudler/LocalAI/core/cli/context"  
"github.com/mudler/LocalAI/pkg/assets"  
"github.com/mudler/LocalAI/pkg/library"  
"github.com/rs/zerolog/log"  
  
External data, input sources:  
- Backend assets files are extracted from the embedded FS and placed in the specified path.  
- The user provides the llama-rpc-server arguments through command line arguments.  
  
TODOs:  
- None found.  
  
Summary of major code parts:  
- The LLamaCPP struct embeds the WorkerFlags struct and is responsible for running the llama-cpp-rpc server.  
- The Run method extracts backend assets files, checks for the required command line arguments, and then executes the llama-cpp-rpc-server process with the provided arguments.  
- The method uses the library.LoadLDSO function to load the necessary shared libraries for the server to function correctly.  
  
  
  
# core/cli/worker/worker_nop2p.go  
Package/Component name: worker  
  
Imports:  
- "fmt"  
- "github.com/mudler/LocalAI/core/cli/context"  
  
External data, input sources:  
- None  
  
TODOs:  
- None  
  
Summary:  
The worker package contains a P2P struct that implements the Run method. The Run method returns an error indicating that the p2p mode is not enabled in the current build.  
  
  
  
# core/cli/worker/worker_p2p.go  
Package/Component name: worker  
  
Imports:  
- context  
- fmt  
- os  
- os/exec  
- strings  
- time  
- github.com/mudler/LocalAI/core/cli/context  
- github.com/mudler/LocalAI/core/p2p  
- github.com/mudler/LocalAI/pkg/assets  
- github.com/mudler/LocalAI/pkg/library  
- github.com/phayes/freeport  
- github.com/rs/zerolog/log  
  
External data, input sources:  
- Environment variables: LOCALAI_TOKEN, LOCALAI_P2P_TOKEN, TOKEN, LOCALAI_NO_RUNNER, NO_RUNNER, LOCALAI_RUNNER_ADDRESS, RUNNER_ADDRESS, LOCALAI_RUNNER_PORT, RUNNER_PORT, LOCALAI_P2P_NETWORK_ID, P2P_NETWORK_ID  
  
TODOs:  
- Add support for custom runner configuration.  
  
Summary:  
- The P2P struct defines the configuration for the P2P mode, including the token, network ID, and runner address and port.  
- The Run method handles the setup and execution of the P2P worker. It extracts backend assets, checks for the token, and starts the llama-cpp-rpc-server or exposes the service depending on the configuration.  
- If the NoRunner flag is set, the method exposes the service using the provided address and port. Otherwise, it starts the llama-cpp-rpc-server with the specified address and port, and then exposes the service.  
- The method also includes a loop that keeps the worker running indefinitely.  
  
