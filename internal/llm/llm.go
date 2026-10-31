package llm

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

const openRouterURL = "https://openrouter.ai/api/v1/chat/completions"

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type Request struct {
	Model    string    `json:"model"`
	Messages []Message `json:"messages"`
}

type Response struct {
	Choices []struct {
		Message Message `json:"message"`
	} `json:"choices"`
}

func SendMessage(memoryContext string, userMessage string, apiKey string, model string) (string, error) {
	reqPayload := Request{
		Model: model,
		Messages: []Message{
			{
				Role:    "system",
				Content: memoryContext,
			},
			{
				Role:    "user",
				Content: userMessage,
			},
		},
	}

	bodyBytes, err := json.Marshal(reqPayload)
	if err != nil {
		return "", fmt.Errorf("erro ao serializar requisição JSON: %w", err)
	}

	req, err := http.NewRequest(http.MethodPost, openRouterURL, bytes.NewReader(bodyBytes))
	if err != nil {
		return "", fmt.Errorf("erro ao criar requisição HTTP: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+apiKey)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", fmt.Errorf("erro ao enviar requisição ao OpenRouter: %w", err)
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("erro ao ler resposta do OpenRouter: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("erro na API do OpenRouter (status %d): %s", resp.StatusCode, string(respBody))
	}

	var llmResp Response
	if err := json.Unmarshal(respBody, &llmResp); err != nil {
		return "", fmt.Errorf("erro ao fazer parse da resposta JSON: %w", err)
	}

	if len(llmResp.Choices) == 0 {
		return "", errors.New("resposta da API do OpenRouter não contém choices")
	}

	return llmResp.Choices[0].Message.Content, nil
}

