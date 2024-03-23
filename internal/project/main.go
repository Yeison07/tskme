package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
	"github.com/rs/zerolog/log"
	"github.com/yeison07/tskme/internal/common/server"
)

func main() {

	err := godotenv.Load("../../.env")

	if err != nil {
		log.Fatal().Msg("Couldnt find env file")
		panic(err)
	}
	project := CreateNewProject("el pepe", "ete sech", nil, nil)

	serverType := strings.ToLower(os.Getenv("SERVER_TO_RUN"))
	switch serverType {
	case "http":
		server.RunHTTPServer(func(router chi.Router) http.Handler {
			router.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusCreated)
				json.NewEncoder(w).Encode(project)

			})
			return router
		})
	default:
		panic(fmt.Sprintf("server type '%s' is not supported", serverType))

	}

}
