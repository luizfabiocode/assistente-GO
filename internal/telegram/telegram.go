package telegram

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"

	"botgo/internal/llm"
	"botgo/internal/memory"
)

const (
	apiBaseURL       = "https://api.telegram.org/bot"
	maxMessageLength = 4000
)

type UpdateResponse struct {
	OK     bool     `json:"ok"`
	Result []Update `json:"result"`
}

type Update struct {
	UpdateID int64    `json:"update_id"`
	Message  *Message `json:"message,omitempty"`
}

type Message struct {
	MessageID int64  `json:"message_id"`
	Chat      Chat   `json:"chat"`
	Text      string `json:"text"`
}

type Chat struct {
	ID int64 `json:"id"`
}

type sendMessagePayload struct {
	ChatID int64  `json:"chat_id"`
	Text   string `json:"text"`
}

func StartPolling(token string, apiKey string, model string) {
	var offset int64

	fmt.Println("Iniciando long polling do Telegram...")

	for {
		updates, err := getUpdates(token, offset)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Erro ao buscar atualizações do Telegram: %v\n", err)
			time.Sleep(5 * time.Second)
			continue
		}

		for _, update := range updates {
			offset = update.UpdateID + 1

			if update.Message == nil || update.Message.Text == "" {
				continue
			}

			chatID := update.Message.Chat.ID
			userText := update.Message.Text

			handleMessage(token, apiKey, model, chatID, userText)
		}
	}
}

func handleMessage(token string, apiKey string, model string, chatID int64, userText string) {
	if err := memory.InitUser(chatID); err != nil {
		fmt.Fprintf(os.Stderr, "Erro ao inicializar diretório do usuário %d: %v\n", chatID, err)
		return
	}

	memContext, err := memory.ReadMemory(chatID)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Erro ao ler memória do usuário %d: %v\n", chatID, err)
		return
	}

	reply, err := llm.SendMessage(memContext, userText, apiKey, model)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Erro no LLM para o usuário %d: %v\n", chatID, err)
		return
	}

	interaction := fmt.Sprintf("User: %s / Bot: %s", userText, reply)
	if err := memory.AppendHistory(chatID, interaction); err != nil {
		fmt.Fprintf(os.Stderr, "Erro ao salvar histórico do usuário %d: %v\n", chatID, err)
	}

	if err := sendMessage(token, chatID, reply); err != nil {
		fmt.Fprintf(os.Stderr, "Erro ao enviar mensagem ao chat %d: %v\n", chatID, err)
	}
}

func getUpdates(token string, offset int64) ([]Update, error) {
	url := fmt.Sprintf("%s%s/getUpdates?timeout=30&offset=%d", apiBaseURL, token, offset)

	client := &http.Client{
		Timeout: 40 * time.Second,
	}

	resp, err := client.Get(url)
	if err != nil {
		return nil, fmt.Errorf("falha na requisição getUpdates: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("erro ao ler corpo da resposta: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("status inesperado %d: %s", resp.StatusCode, string(body))
	}

	var updateResp UpdateResponse
	if err := json.Unmarshal(body, &updateResp); err != nil {
		return nil, fmt.Errorf("erro ao decodificar JSON de updates: %w", err)
	}

	if !updateResp.OK {
		return nil, fmt.Errorf("API do Telegram retornou ok=false: %s", string(body))
	}

	return updateResp.Result, nil
}

func splitMessage(text string, maxLength int) []string {
	runes := []rune(text)
	if len(runes) <= maxLength {
		return []string{text}
	}

	var chunks []string
	for len(runes) > 0 {
		if len(runes) <= maxLength {
			chunks = append(chunks, string(runes))
			break
		}

		window := string(runes[:maxLength])
		cutIndex := maxLength

		// Tenta cortar na última quebra de linha dentro do limite
		if lastNewline := strings.LastIndex(window, "\n"); lastNewline > 0 {
			cutIndex = len([]rune(window[:lastNewline]))
		} else if lastSpace := strings.LastIndex(window, " "); lastSpace > 0 {
			// Fallback: corta no último espaço para não quebrar palavras no meio
			cutIndex = len([]rune(window[:lastSpace]))
		}

		chunk := strings.TrimRight(string(runes[:cutIndex]), "\n ")
		if chunk != "" {
			chunks = append(chunks, chunk)
		}

		// Avança o slice removendo a quebra de linha ou espaço usado como ponto de corte
		runes = runes[cutIndex:]
		for len(runes) > 0 && (runes[0] == '\n' || runes[0] == ' ') {
			runes = runes[1:]
		}
	}

	return chunks
}

func sendMessage(token string, chatID int64, text string) error {
	url := fmt.Sprintf("%s%s/sendMessage", apiBaseURL, token)
	chunks := splitMessage(text, maxMessageLength)

	for _, chunk := range chunks {
		payload := sendMessagePayload{
			ChatID: chatID,
			Text:   chunk,
		}

		bodyBytes, err := json.Marshal(payload)
		if err != nil {
			return fmt.Errorf("erro ao serializar payload de envio: %w", err)
		}

		resp, err := http.Post(url, "application/json", bytes.NewReader(bodyBytes))
		if err != nil {
			return fmt.Errorf("falha na requisição sendMessage: %w", err)
		}

		if resp.StatusCode != http.StatusOK {
			respBody, _ := io.ReadAll(resp.Body)
			resp.Body.Close()
			return fmt.Errorf("erro na API do Telegram ao enviar mensagem (status %d): %s", resp.StatusCode, string(respBody))
		}
		resp.Body.Close()
	}

	return nil
}
