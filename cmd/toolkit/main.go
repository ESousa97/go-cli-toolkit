package main

import (
	"log"

	"github.com/sousa/go-cli-toolkit/internal/commands"
)

func main() {
	if err := commands.Execute(); err != nil {
		log.Fatalf("Erro ao executar CLI: %v", err)
	}
}
