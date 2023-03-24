package main

import (
	"github.com/Luftalian/writers_app/infrastructure"
)

func main() {
	infrastructure.Router.Logger.Fatal(infrastructure.Router.Start(":8080"))
}
