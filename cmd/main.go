package main

import (
	"fmt"
	"os"

	"botgo/internal/config"
	"botgo/internal/memory"
	"botgo/internal/telegram"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Erro ao carregar configuração: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("Configuração carregada com sucesso!")

	if err := memory.Init(); err != nil {
		fmt.Fprintf(os.Stderr, "Erro ao inicializar sistema de memória: %v\n", err)
		os.Exit(1)
	}

	telegram.StartPolling(cfg.TelegramToken, cfg.OpenRouterAPIKey, cfg.ModelName)
}
