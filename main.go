package main

import (
	"github.com/majormilan/miranda-browser/backend"
	"github.com/majormilan/miranda-browser/frontend"
)

func main() {
	backend.Initialize() // backend/init.go
	frontend.Run()       // frontend/frontend.go
}
