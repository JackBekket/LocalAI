# cli_api

This package provides a command-line interface for interacting with the P2P stack.

## Project package structure:
- core/cli/api/p2p.go

## Code summary:
The `StartP2PStack` function initializes and starts the P2P stack. It checks if the federated mode is enabled and if so, exposes a service to the local instance. If the token is provided, it starts the service discovery process. The function also handles the discovery of available nodes and sets the `LLAMACPP_GRPC_SERVERS` environment variable with the tunnel addresses of the discovered nodes.

## Configuration:
- `LLAMACPP_GRPC_SERVERS`: Environment variable containing the tunnel addresses of the discovered nodes.

## Edge cases:
- If the federated mode is enabled, the service will be exposed to the local instance.
- If the token is provided, the service discovery process will be started.

## Unclear places:
- It is unclear how the `LLAMACPP_GRPC_SERVERS` environment variable is used by the P2P stack.

