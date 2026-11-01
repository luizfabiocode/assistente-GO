package memory

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"time"
)

const (
	usersDir    = "data/users"
	memoryFile  = "MEMORY.md"
	historyFile = "HISTORY.md"
)

func Init() error {
	if err := os.MkdirAll(usersDir, 0755); err != nil {
		return fmt.Errorf("erro ao criar diretório %s: %w", usersDir, err)
	}
	return nil
}

func userDirPath(chatID int64) string {
	return filepath.Join(usersDir, strconv.FormatInt(chatID, 10))
}

func InitUser(chatID int64) error {
	dir := userDirPath(chatID)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return fmt.Errorf("erro ao criar diretório do usuário %s: %w", dir, err)
	}

	memPath := filepath.Join(dir, memoryFile)
	if err := ensureFile(memPath, "# Memória de Longo Prazo\n\n"); err != nil {
		return err
	}

	histPath := filepath.Join(dir, historyFile)
	if err := ensureFile(histPath, "# Histórico de Interações\n\n"); err != nil {
		return err
	}

	return nil
}

func ensureFile(path string, defaultHeader string) error {
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		if err := os.WriteFile(path, []byte(defaultHeader), 0644); err != nil {
			return fmt.Errorf("erro ao criar arquivo %s: %w", path, err)
		}
		return nil
	}
	if err != nil {
		return fmt.Errorf("erro ao verificar arquivo %s: %w", path, err)
	}
	return nil
}

func ReadMemory(chatID int64) (string, error) {
	if err := InitUser(chatID); err != nil {
		return "", err
	}

	memPath := filepath.Join(userDirPath(chatID), memoryFile)
	data, err := os.ReadFile(memPath)
	if err != nil {
		return "", fmt.Errorf("erro ao ler arquivo %s: %w", memPath, err)
	}
	return string(data), nil
}

func AppendHistory(chatID int64, event string) error {
	if err := InitUser(chatID); err != nil {
		return err
	}

	histPath := filepath.Join(userDirPath(chatID), historyFile)
	f, err := os.OpenFile(histPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return fmt.Errorf("erro ao abrir arquivo %s: %w", histPath, err)
	}
	defer f.Close()

	timestamp := time.Now().Format("2006-01-02 15:04")
	line := fmt.Sprintf("[%s] - %s\n", timestamp, event)

	if _, err := f.WriteString(line); err != nil {
		return fmt.Errorf("erro ao escrever no arquivo %s: %w", histPath, err)
	}

	return nil
}
