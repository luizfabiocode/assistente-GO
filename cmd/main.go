package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"botgo/internal/config"
	"botgo/internal/llm"
	"botgo/internal/memory"
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

	if err := memory.AppendHistory("Bot iniciado"); err != nil {
		fmt.Fprintf(os.Stderr, "Erro ao registrar no histórico: %v\n", err)
		os.Exit(1)
	}

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Você: ")
		if !scanner.Scan() {
			break
		}

		userInput := strings.TrimSpace(scanner.Text())
		if userInput == "" {
			continue
		}

		if strings.EqualFold(userInput, "sair") {
			fmt.Println("Encerrando o bot...")
			break
		}

		memContext, err := memory.ReadMemory()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Erro ao ler memória: %v\n", err)
			continue
		}

		reply, err := llm.SendMessage(memContext, userInput, cfg.OpenRouterAPIKey, cfg.ModelName)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Erro ao comunicar com LLM: %v\n", err)
			continue
		}

		fmt.Printf("Bot: %s\n", reply)

		interaction := fmt.Sprintf("User: %s / Bot: %s", userInput, reply)
		if err := memory.AppendHistory(interaction); err != nil {
			fmt.Fprintf(os.Stderr, "Erro ao salvar interação no histórico: %v\n", err)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "Erro na leitura da entrada: %v\n", err)
		os.Exit(1)
	}
}
