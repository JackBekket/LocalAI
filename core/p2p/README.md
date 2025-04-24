# p2p

This package provides a p2p server implementation for the edgevpn project.

## Files

- `core/p2p/federated.go`
- `core/p2p/federated_server.go`
- `core/p2p/node.go`
- `core/p2p/p2p.go`
- `core/p2p/p2p_common.go`
- `core/p2p/p2p_disabled.go`

## Summary

The `p2p` package provides functionalities for peer-to-peer networking and service discovery. It includes functions for generating tokens, starting the federated server, discovering services, exposing services, and creating new nodes.

The `FederatedServer` struct is responsible for starting the p2p server and handling incoming connections. The `Start` method initializes a new node, starts it, and then discovers other nodes in the network. The `proxy` method listens for incoming connections and forwards them to the appropriate p2p node.

The `NodeData` struct stores information about a node, including its name, ID, tunnel address, service ID, and last seen time. The `GetAvailableNodes` function returns a list of available nodes for a given service ID. If no service ID is provided, it defaults to "services". The `GetNode` function returns a specific node for a given service ID and node ID. If the node is not found, it returns an empty `NodeData` struct and false. The `AddNode` function adds a new node to the system for a given service ID. If the service ID is not provided, it defaults to "services".

The `p2p_common` package defines a package-level variable `logLevel`, which is initialized to the value of the environment variable `LOCALAI_P2P_LOGLEVEL`. If the environment variable is not set, the log level defaults to "info".

The `p2p_disabled` package provides functions for generating tokens, starting the federated server, discovering services, exposing services, and creating new nodes.

## Configuration

- Environment variable `LOCALAI_P2P_LOGLEVEL` can be used to set the log level.

## Edge Cases

- If the environment variable `LOCALAI_P2P_LOGLEVEL` is not set, the log level defaults to "info".