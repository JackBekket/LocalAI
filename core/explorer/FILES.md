# core/explorer/database.go  
Package: explorer  
  
Imports:  
- "encoding/json"  
- "os"  
- "sort"  
- "sync"  
- "github.com/gofrs/flock"  
  
External data, input sources:  
- The database reads and writes data from a JSON file specified by the path provided when creating a new Database instance.  
  
TODOs:  
- There are no TODO comments in the provided code.  
  
Summary:  
- The code defines a simple JSON database for storing and retrieving p2p network tokens and their associated name and description.  
- The database is implemented using a map to store the tokens and their data, and it uses a file lock to ensure that only one process can write to the database at a time.  
- The database supports adding, retrieving, and deleting tokens, as well as listing all available tokens.  
- The database is designed to be used in a concurrent environment, with locks to protect shared resources.  
- The database is implemented using the JSON format for data storage and retrieval.  
- The database is designed to be used in a distributed environment, with a file lock to ensure that only one process can write to the database at a time.  
- The database is designed to be used in a concurrent environment, with locks to protect shared resources.  
  
# core/explorer/discovery.go  
Package: explorer  
  
Imports:  
- context  
- fmt  
- strings  
- sync  
- time  
- github.com/rs/zerolog/log  
- github.com/mudler/LocalAI/core/p2p  
- github.com/mudler/edgevpn/pkg/blockchain  
  
External data, input sources:  
- The code interacts with a database to store and retrieve information about network tokens and their associated clusters.  
- It uses the p2p package to connect to the network and retrieve data from the ledger.  
  
TODOs:  
- There are no TODO comments in the provided code.  
  
Summary:  
- The DiscoveryServer struct is responsible for discovering and maintaining information about the network. It keeps track of the number of nodes, the clusters they belong to, and the workers within each cluster.  
- The NewDiscoveryServer function initializes a new DiscoveryServer instance with the given database, connection timeout, and failure threshold.  
- The runBackground function is responsible for collecting network data and updating the database accordingly. It connects to the network, retrieves the ledger, and extracts information about the clusters and workers.  
- The failedToken function is called when a token fails to connect to the network. It increments the failure count for the token in the database.  
- The deleteFailedConnections function removes tokens from the database if their failure count exceeds the threshold.  
- The retrieveNetworkData function retrieves network data from the ledger and populates the clusters map with the extracted information.  
- The Start function starts the discovery server and runs it in a goroutine. It continuously collects data and updates the database.  
  
  
  
# core/http/endpoints/explorer/dashboard.go  
Package: explorer  
  
Imports:  
- "encoding/base64"  
- "sort"  
- "github.com/gofiber/fiber/v2"  
- "github.com/mudler/LocalAI/core/explorer"  
- "github.com/mudler/LocalAI/core/http/utils"  
- "github.com/mudler/LocalAI/internal"  
  
External data, input sources:  
- The code interacts with a database to store and retrieve network information.  
  
TODOs:  
- TODO: check if token is valid, otherwise reject  
  
Summary:  
- The `Dashboard` function provides a dashboard with information about the LocalAI API, including its title, version, and base URL. It returns the information as JSON or renders an HTML template depending on the client's request.  
- The `ShowNetworks` function retrieves a list of networks from the database and returns them as JSON. It filters the networks to include only those with workers and orders them by the number of clusters.  
- The `AddNetwork` function handles requests to add a new network to the database. It parses the request body, validates the provided token, name, and description, and stores the network information in the database.  
  
