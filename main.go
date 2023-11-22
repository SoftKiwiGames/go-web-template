package main

import (
	"SoftKiwiGames/go-web-template/web"
	"os"

	_ "embed"
)

var (
	//go:embed assets/images/logo.svg
	logo []byte
)

func main() {
	srv := &web.Server{
		EmbeddableResources: web.EmbeddableResources{
			Favicon: logo,
			Logo:    logo,
		},
	}
	srv.Run(os.Args)
}
