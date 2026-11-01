# 🤖 Assistente Pessoal em Go (Golang)

![Go Version](https://img.shields.io/badge/Go-1.22+-00ADD8?style=for-the-badge&logo=go)
![Docker](https://img.shields.io/badge/Docker-Enabled-2496ED?style=for-the-badge&logo=docker)
![OpenRouter](https://img.shields.io/badge/LLM-OpenRouter-8A2BE2?style=for-the-badge)
![Telegram](https://img.shields.io/badge/Telegram-Bot-26A5E4?style=for-the-badge&logo=telegram)

Um agente de Inteligência Artificial leve, eficiente e com memória persistente, construído do zero em **Go (Golang)**. O projeto utiliza a API do **OpenRouter** para acesso a Modelos de Linguagem (LLMs) gratuitos e atua como um bot no **Telegram**, projetado para rodar de forma contínua em ambientes de nuvem (como o Botkeep).

## 🚀 Funcionalidades

- **Integração com Telegram**: Interface de conversação remota usando a API oficial do Telegram via _Long Polling_ (implementado com a biblioteca padrão `net/http`, sem dependências externas).
- **Memória Multi-usuário (Multi-tenant)**: Isolamento de contexto e histórico por usuário. O sistema utiliza o `chat_id` do Telegram para criar diretórios privados (`data/users/<chat_id>/`), garantindo que as conversas nunca se misturem.
- **Memória Persistente**: Sistema de memória de longo prazo dividido em arquivos Markdown (`MEMORY.md` e `HISTORY.md`), permitindo que o bot lembre de contextos passados.
- **Tratamento de Limites (Chunking)**: Lógica inteligente de fatiamento de mensagens que respeita o limite de 4096 caracteres do Telegram, cortando o texto por quebras de linha sem corromper caracteres especiais ou emojis.
- **Integração com LLM**: Conexão com a API do OpenRouter para respostas de IA.
- **Containerização**: Pronto para deploy via Docker com _multi-stage build_, gerando uma imagem final ultra-leve.

## 🛠️ Tecnologias Utilizadas

- **Linguagem**: Go (Golang)
- **Inteligência Artificial**: OpenRouter (Modelos gratuitos como Nemotron e Llama)
- **Mensageria**: API Oficial do Telegram (Long Polling)
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
│   ├── memory/           # Sistema de persistência (multi-usuário)
│   └── telegram/         # Lógica de polling e envio de mensagens
├── data/                 # Criado em runtime (contém as pastas dos usuários)
├── Dockerfile            # Build otimizado para produção
├── go.mod
└── README.md
```

## ⚙️ Como Rodar Localmente

### Pré-requisitos

- Go 1.22+ instalado.
- Uma chave de API gratuita do [OpenRouter](https://openrouter.ai/).
- Um Token de Bot gerado pelo [@BotFather](https://t.me/BotFather) no Telegram.

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
export TELEGRAM_TOKEN="seu-token-do-telegram"

# No Windows (PowerShell)
$env:OPENROUTER_API_KEY="sua-chave-aqui"
$env:MODEL_NAME="openrouter/free"
$env:TELEGRAM_TOKEN="seu-token-do-telegram"
```

3. Execute a aplicação:

```bash
go run cmd/main.go
```

_(O terminal ficará em modo de escuta. Vá até o seu bot no Telegram e envie uma mensagem)._

## 🐳 Como Rodar com Docker

O projeto conta com um `Dockerfile` otimizado que gera uma imagem final baseada em Alpine Linux com menos de 20MB.

```bash
docker build -t assistente-go .
docker run -e OPENROUTER_API_KEY="sua-chave" -e MODEL_NAME="openrouter/free" -e TELEGRAM_TOKEN="seu-token" assistente-go
```

## 🗺️ Próximos Passos (Roadmap)

- [x] Implementação de memória persistente.
- [x] Integração com o Telegram para conversação remota (Long Polling).
- [x] Tratamento de limite de caracteres do Telegram (Chunking).
- [x] Isolamento de memória multi-usuário.
- [ ] Implementação de sistema de "Skills" (habilidades customizadas via Markdown).
- [ ] Adição de tarefas proativas (Heartbeat/Cron).
- [ ] Deploy em nuvem (Botkeep).

## 👨‍💻 Autor

Desenvolvido por **Luiz Fabio**
[LinkedIn](https://www.linkedin.com/in/luizfabiocode/) |

---

_Este projeto foi desenvolvido como parte dos estudos de arquitetura Backend com Go, Cloud Computing e integração de IA._
