# swagger/docs.go  
```go  
package main  
  
import (  
	"fmt"  
	"net/http"  
	"os"  
	"os/signal"  
	"syscall"  
	"time"  
  
	"github.com/gorilla/mux"  
	"github.com/rs/zerolog"  
)  
  
func main() {  
	// Initialize logger  
	logger := zerolog.New(os.Stdout).With().Timestamp().Logger()  
  
	// Create a new router  
	router := mux.NewRouter()  
  
	// Register API endpoints  
	router.HandleFunc("/health", healthHandler).Methods("GET")  
	router.HandleFunc("/version", versionHandler).Methods("GET")  
  
	// Start the server  
	server := &http.Server{  
		Addr:         ":8080",  
		Handler:      router,  
		ReadTimeout:  5 * time.Second,  
		WriteTimeout: 10 * time.Second,  
		IdleTimeout:  120 * time.Second,  
	}  
  
	go func() {  
		logger.Info().Msg("Starting server")  
		if err := server.ListenAndServe(); err != nil {  
			logger.Error().Err(err).Msg("Error starting server")  
		}  
	}()  
  
	// Graceful shutdown  
	quit := make(chan os.Signal, 1)  
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)  
	<-quit  
	logger.Info().Msg("Shutting down server")  
  
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)  
	defer cancel()  
	if err := server.Shutdown(ctx); err != nil {  
		logger.Error().Err(err).Msg("Error shutting down server")  
	}  
}  
  
func healthHandler(w http.ResponseWriter, r *http.Request) {  
	w.WriteHeader(http.StatusOK)  
	fmt.Fprintf(w, "OK")  
}  
  
func versionHandler(w http.ResponseWriter, r *http.Request) {  
	w.WriteHeader(http.StatusOK)  
	fmt.Fprintf(w, "LocalAI API v1.0.0")  
}  
```  
  
This code sets up a basic REST API with two endpoints:  
  
1. `/health`: Returns a 200 OK status to indicate the server is healthy.  
2. `/version`: Returns the API version.  
  
The server listens on port 8080 and uses Gorilla Mux for routing. It also includes graceful shutdown functionality to allow the server to shut down cleanly when a SIGINT or SIGTERM signal is received.  
  
To run this code:  
  
1. Save the code as `main.go`.  
2. Install the required dependencies: `go get github.com/gorilla/mux github.com/rs/zerolog`  
3. Run the code: `go run main.go`  
  
Once the server is running, you can access the API endpoints at `http://localhost:8080/health` and `http://localhost:8080/version`.  
  
