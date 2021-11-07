package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/ntsd/alpha-interface-backend/pkg/config"
	"github.com/ntsd/alpha-interface-backend/pkg/server"
	"github.com/spf13/cobra"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Can not load .env file.")
	}

	cfg := config.GetConfig()

	cmd := &cobra.Command{
		Use:   "serve",
		Short: "Serve the http server",
		Run: func(_ *cobra.Command, _ []string) {
			err := server.Serve(cfg)
			if err != nil {
				log.Fatal(err)
			}
		},
	}

	err = cmd.Execute()
	if err != nil {
		log.Fatal(err)
	}
}
