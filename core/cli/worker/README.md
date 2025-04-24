# worker

This package provides a way to start a LocalAI llama.cpp worker in either P2P mode or standalone mode.

## File structure:
- worker.go
- worker_llamacpp.go
- worker_nop2p.go
- worker_p2p.go

## Code summary:
The Worker struct contains two fields, P2P and LLamaCPP, which represent the two modes of operation.

- P2P mode: Starts a LocalAI llama.cpp worker in P2P mode, which requires a token.
- Standalone mode: Starts a llama.cpp worker in standalone mode.

The package allows users to easily configure and start a llama.cpp worker for their specific needs.

## Configuration:
- Environment variables: LOCALAI_BACKEND_ASSETS_PATH, BACKEND_ASSETS_PATH, LOCALAI_EXTRA_LLAMA_CPP_ARGS, EXTRA_LLAMA_CPP_ARGS

## Edge cases:
- If the p2p mode is not enabled in the current build, the P2P struct's Run method will return an error.

## Unclear places:
- The package does not explicitly mention how the llama-cpp-rpc-server process is terminated when the worker is stopped.

## Dead code:
- None found.

## Relations between code entities:
- The Worker struct contains two fields, P2P and LLamaCPP, which represent the two modes of operation.
- The P2P struct implements the Run method, which handles the setup and execution of the P2P worker.
- The LLamaCPP struct implements the Run method, which handles the setup and execution of the standalone worker.

## Possible improvements:
- Add support for custom runner configuration in the P2P mode.
- Provide a way to terminate the llama-cpp-rpc-server process when the worker is stopped.