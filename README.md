<div align="center">

# Go CLI Toolkit

[![CodeFactor](https://img.shields.io/codefactor/grade/github/ESousa97/go-cli-toolkit?style=flat&logo=codefactor&logoColor=white)](https://www.codefactor.io/repository/github/ESousa97/go-cli-toolkit)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg?style=flat&logo=opensourceinitiative&logoColor=white)](https://opensource.org/licenses/MIT)
[![Status](https://img.shields.io/badge/Status-Active-brightgreen.svg?style=flat&logo=check&logoColor=white)](#)

**Projeto educacional para prática e construção de uma Interface de Linha de Comando (CLI) utilitária em Go — construído com o framework Cobra CLI, seguindo as premissas do Standard Go Project Layout. Organizado com ponto de entrada isolado em `cmd/` e lógica encapsulada em `internal/`, promovendo modularização extrema e arquitetura stateless.**

</div>

---

## Índice

- [Sobre o Projeto](#sobre-o-projeto)
- [Funcionalidades](#funcionalidades)
- [Tecnologias](#tecnologias)
- [Arquitetura](#arquitetura)
- [Estrutura do Projeto](#estrutura-do-projeto)
- [Começando](#começando)
  - [Pré-requisitos](#pré-requisitos)
  - [Instalação](#instalação)
  - [Uso](#uso)
- [Licença](#licença)
- [Contato](#contato)

---

## Sobre o Projeto

Projeto em Go para construção de uma Interface de Linha de Comando (CLI) com foco na implementação inicial estruturada seguindo os princípios absolutos de modularização extrema. O repositório foi organizado com padrão de produção, isolando dependências externas e lógica de negócio e entrada da aplicação.

O repositório prioriza:

- **Organização por Bounded Contexts** — Código fonte dividido em pacotes lógicos (`cmd/` para inicialização e `internal/commands/` para comandos CLI), evitando exportação de lógicas dependentes da aplicação.
- **Isolamento de Ponto de Entrada** — O `main.go` apenas invoca a CLI. Toda a configuração semântica de comandos fica restrita ao componente filho.
- **Gestão de Comandos com Cobra** — Gerenciador de comandos hierárquico, permitindo evolução rápida na adoção de subcomandos e _flags_.
- **Sem Magic Values** — Todas as definições dos comandos (uso, mensagem curta e longa, etc.) são providas via constantes fortemente tipadas.

---

## Funcionalidades

- **Comando Raiz (`toolkit`)** — Configuração inicial do entrypoint.
- **Subcomando `ping`** — Verifica se uma URL está acessível através de uma requisição HTTP GET com timeout controlado.
- **Subcomando `format json`** — Lê um JSON (via arquivo ou stdin), valida sua estrutura e o imprime formatado (Pretty Print).

---

## Tecnologias

![Go](https://img.shields.io/badge/go-%2300ADD8.svg?style=flat&logo=go&logoColor=white)
![Cobra](https://img.shields.io/badge/Cobra_CLI-E04E39?style=flat&logo=go&logoColor=white)

---

## Arquitetura

```mermaid
graph TD
    subgraph "Entrypoint"
        A[cmd/toolkit/main.go]
    end

    subgraph "Lógica de Comandos (internal)"
        A -- Invokes --> B[commands.Execute]
        B --> C[root.go]
        C -- Registers --> D[ping.go]
        C -- Registers --> G[format.go]
    end

    subgraph "Core Business"
        D -- Executes --> E[pingHost]
        G -- Executes --> H[runFormatJSON]
        E -- Uses --> F[net/http]
        H -- Uses --> I[encoding/json]
    end
```

### Pacotes e Responsabilidades

| Pacote                 | Responsabilidade                                                                                   |
| ---------------------- | -------------------------------------------------------------------------------------------------- |
| `cmd/toolkit/main.go`  | Entrypoint do binário. Isola a função main() de regras de negócio.                                 |
| `internal/commands`    | Organiza os comandos e subcomandos utilizando Cobra CLI.                                           |
| `net/http` e `context` | Bibliotecas standard usadas para controle da rede com segurança (Timeout estrito contra gargalos). |

---

## Estrutura do Projeto

```
go-cli-toolkit/
├── cmd/
│   └── toolkit/
│       └── main.go                     # Entrypoint principal
├── internal/
│   └── commands/
│       ├── root.go                     # Comando base da CLI (Cobra Setup)
│       ├── ping.go                     # Implementação de 'ping'
│       └── format.go                   # Implementação de 'format json'
├── go.mod                              # Manifesto de dependências do Go
└── go.sum                              # Lock de checksum
```

---

## Começando

### Pré-requisitos

- Go 1.21+ (ou versão superior instalada localmente)
- Terminal/Prompt de Comando para interação

### Instalação

```bash
git clone https://github.com/sousa/go-cli-toolkit.git
cd go-cli-toolkit
go mod download
```

### Compilação do Binário

**Compilar na raiz do ecossistema:**

```bash
go build -o tk.exe ./cmd/toolkit
```

_(No Linux/macOS remova o `.exe`)_

### Uso

Para rodar ajuda da ferramenta raiz:

```bash
./tk.exe --help
```

### Ping

Executar o subcomando `ping` em uma URL válida:

```bash
./tk.exe ping https://www.google.com
# O host https://www.google.com está ONLINE (Status: 200)
```

### Format JSON

Formatar um JSON bagunçado via arquivo:

```bash
./tk.exe format json --file raw.json
```

Ou via pipe stdin:

```bash
echo '{"name":"toolkit"}' | ./tk.exe format json
```

Exemplo de teste completo (criação, execução e limpeza):

```powershell
echo '{"name": "teste_final", "status": true}' > test.json; .\tk.exe format json --file test.json; rm test.json
```

Output esperado:

```json
{
  "name": "teste_final",
  "status": true
}
```

Testando caso de falha:

```bash
./tk.exe ping https://site.que.nao.existe
```

---

## Licença

Este projeto está sob a licença MIT. Veja o arquivo [LICENSE](LICENSE) para mais detalhes.

```
MIT License - você pode usar, copiar, modificar e distribuir este código.
```

---

## Contato

**Enoque Sousa**

[![LinkedIn](https://img.shields.io/badge/LinkedIn-0077B5?style=flat&logo=linkedin&logoColor=white)](https://www.linkedin.com/in/enoque-sousa-bb89aa168/)
[![GitHub](https://img.shields.io/badge/GitHub-100000?style=flat&logo=github&logoColor=white)](https://github.com/ESousa97)
[![Portfolio](https://img.shields.io/badge/Portfolio-FF5722?style=flat&logo=todoist&logoColor=white)](https://enoquesousa.vercel.app)

---

<div align="center">

**[⬆ Voltar ao topo](#go-cli-toolkit)**

Feito com ❤️ por [Enoque Sousa](https://github.com/ESousa97)

**Status do Projeto:** Ativo — Em constante atualização

</div>
