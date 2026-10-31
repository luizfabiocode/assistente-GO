# 🤖 Assistente Pessoal em Go (Golang)

![Go Version](https://img.shields.io/badge/Go-1.22+-00ADD8?style=for-the-badge&logo=go)
![Docker](https://img.shields.io/badge/Docker-Enabled-2496ED?style=for-the-badge&logo=docker)
![OpenRouter](https://img.shields.io/badge/LLM-OpenRouter-8A2BE2?style=for-the-badge)

Um agente de Inteligência Artificial leve, eficiente e com memória persistente, construído do zero em **Go (Golang)**. O projeto utiliza a API do **OpenRouter** para acesso a Modelos de Linguagem (LLMs) gratuitos e foi arquitetado para rodar de forma contínua em ambientes de nuvem (como o Botkeep).

## 🚀 Funcionalidades

- **Memória Persistente**: Sistema de memória de longo prazo dividido em dois arquivos Markdown (`MEMORY.md` e `HISTORY.md`), permitindo que o bot lembre de contextos e interações passadas.
- **Integração com LLM**: Conexão com a API do OpenRouter usando apenas a biblioteca padrão do Go (`net/http`), sem dependências externas pesadas.
- **Arquitetura Modular**: Código rigorosamente separado em pacotes internos (`internal/config`, `internal/memory`, `internal/llm`).
- **Interface CLI**: Loop de conversação interativo via terminal (fase inicial).
- **Containerização**: Pronto para deploy via Docker com _multi-stage build_, gerando uma imagem final ultra-leve.

## 🛠️ Tecnologias Utilizadas

- **Linguagem**: Go (Golang)
- **Inteligência Artificial**: OpenRouter (Modelos gratuitos)
- **Infraestrutura**: Docker (Multi-stage build)
- **Arquitetura**: Clean Architecture voltada para microsserviços

## 📂 Estrutura do Projeto

```text
assistente-GO/
├── cmd/
│   └── main.go           # Ponto de entrada da aplicação
├── internal/
│   ├── config/           # Gerenciamento de variáveis de ambiente
│   ├── llm/              # Integração com a API do OpenRouter
│   └── memory/           # Sistema de persistência de memória (Markdown)
├── data/                 # Criado em runtime (MEMORY.md e HISTORY.md)
├── Dockerfile            # Build otimizado para produção
├── go.mod
└── README.md
```

## ⚙️ Como Rodar Localmente

### Pré-requisitos

- Go 1.22+ instalado.
- Uma chave de API gratuita do [OpenRouter](https://openrouter.ai/).

### Passo a passo

1. Clone o repositório:

```bash
git clone https://github.com/luizfabiocode/assistente-GO.git
cd assistente-GO
```

2. Configure as variáveis de ambiente:

```bash
# No Linux/macOS
export OPENROUTER_API_KEY="sua-chave-aqui"
export MODEL_NAME="openrouter/free"

# No Windows (PowerShell)
$env:OPENROUTER_API_KEY="sua-chave-aqui"
$env:MODEL_NAME="openrouter/free"
```

3. Execute a aplicação:

```bash
go run cmd/main.go
```

## 🐳 Como Rodar com Docker

O projeto conta com um `Dockerfile` otimizado que gera uma imagem final baseada em Alpine Linux com menos de 20MB.

```bash
docker build -t assistente-go .
docker run -e OPENROUTER_API_KEY="sua-chave" -e MODEL_NAME="openrouter/free" assistente-go
```

## 🗺️ Próximos Passos (Roadmap)

- [ ] Integração com o Telegram para conversação remota (Long Polling).
- [ ] Implementação de sistema de "Skills" (habilidades customizadas via Markdown).
- [ ] Adição de tarefas proativas (Heartbeat/Cron).
- [ ] Migração da interface CLI para uma API REST.

## 👨‍💻 Autor

Desenvolvido por **Luiz Fabio**
[LinkedIn](https://www.linkedin.com/in/luizfabiocode/) | [GitHub](https://github.com/luizfabiocode)

---

_Este projeto foi desenvolvido como parte dos estudos de arquitetura Backend com Go, Cloud Computing e integração de IA._
