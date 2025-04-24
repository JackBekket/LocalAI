# core/cli/api/p2p.go  
Package/Component name: cli_api  
  
Imports:  
	"context"  
	"fmt"  
	"net"  
	"os"  
	"strings"  
	"github.com/mudler/LocalAI/core/p2p"  
	"github.com/mudler/edgevpn/pkg/node"  
	"github.com/rs/zerolog/log"  
  
External data, input sources:  
	- The code uses the address, token, networkID, and federated flag provided as input parameters.  
  
TODOs:  
	- There are no TODO comments in the code.  
  
Summary:  
	- The StartP2PStack function initializes and starts the P2P stack. It checks if the federated mode is enabled and if so, exposes a service to the local instance. If the token is provided, it starts the service discovery process. The function also handles the discovery of available nodes and sets the LLAMACPP_GRPC_SERVERS environment variable with the tunnel addresses of the discovered nodes.  
  
  
  
