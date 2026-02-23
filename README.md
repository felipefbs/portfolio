# Portfolio

Meu portfólio pessoal desenvolvido em Go.

## Pré-requisitos

- Go 1.21+
- Docker (opcional)

## Instalação das ferramentas

```bash
# Instalar templ (gerador de templates HTML)
go install github.com/a-h/templ/cmd/templ@latest

# Instalar air (hot-reload para desenvolvimento)
go install github.com/cosmtrek/air@latest

# Instalar Tailwind CSS (binário standalone)
wget -O tailwind https://github.com/tailwindlabs/tailwindcss/releases/download/v3.4.1/tailwindcss-linux-x64
chmod +x tailwind

# Ou via npm (alternativo)
npm install -g tailwindcss
```

## Como rodar

### Desenvolvimento

```bash
# Inicia o servidor com hot-reload
air
```

O servidor estará disponível em `http://localhost:8080`.

### Produção

```bash
# Build
go build -o ./tmp/main ./cmd/app/main.go

# Gerar templates (se necessário)
templ generate -path ./templates

# Compilar CSS
tailwind -i ./templates/input.css -o ./static/styles/output.css

# Executar
./tmp/main
```

### Com Docker

```bash
# Build da imagem
docker build -t portfolio .

# Executar container
docker run -p 8080:8080 portfolio
```

## Estrutura do projeto

```
├── cmd/app/main.go      # Ponto de entrada da aplicação
├── portfolio/           # Dados do portfólio (experiência, projetos, etc.)
├── templates/           # Templates HTML (.templ) e CSS de entrada
├── static/              # Arquivos estáticos (CSS, imagens)
└── Dockerfile           # Configuração Docker
```

## Tecnologias

- **Backend**: Go com chi/v5
- **Templates**: a-h/templ
- **Frontend**: Tailwind CSS
- **Hot-reload**: air
