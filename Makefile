.PHONY: build test install clean run

APP_NAME=tk
CMD_PATH=./cmd/toolkit

# Comando padrão
all: build

# Construção do ecossistema local na raiz do projeto
build:
	@echo "Construindo o binário $(APP_NAME)..."
	go build -o $(APP_NAME) $(CMD_PATH)

# Rodar testes unitários e lint básico
test:
	@echo "Rodando a bateria de testes unitários..."
	go test ./... -v -count=1

# Instalar app no diretório GOPATH global do ambiente do dev/usuário
install:
	@echo "Instalando $(APP_NAME) no GOPATH/bin..."
	go install $(CMD_PATH)

# Limpar binários
clean:
	@echo "Limpando artefatos criados..."
	rm -f $(APP_NAME) $(APP_NAME).exe

# Executar a CLI em tempo de compilação rápida (bom pro fluxo de desenvolvimento)
run:
	@go run $(CMD_PATH)/main.go
