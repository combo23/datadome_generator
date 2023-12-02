package main

import (
	"flag"
	"os"

	"github.com/combo23/datadome_generator/api"
	"github.com/joho/godotenv"
)

func main() {
	flag.Parse()
	if os.Getenv("RELEASE_MODE") == "" {
		err := godotenv.Load()
		if err != nil {
			panic(err)
		}
	}
	api.Serve(os.Getenv("RELEASE_MODE"))
}
