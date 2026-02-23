# AGENTS.md - Developer Guidelines

This document provides guidelines for agentic coding agents working in this repository.

## Project Overview

- **Language**: Go 1.21
- **Web Framework**: chi/v5
- **HTML Templates**: a-h/templ (generated .templ.go files)
- **CSS**: Tailwind CSS
- **Server**: HTTP on port 8080

## Build, Run, and Test Commands

### Building the Application

```bash
# Build Go application
go build -o ./tmp/main ./cmd/app/main.go

# Generate templ components (required after modifying .templ files)
templ generate -path ./templates

# Build Tailwind CSS
tailwind -i ./templates/input.css -o ./static/styles/output.css
```

### Running the Application

```bash
# Run with hot-reload using air (recommended for development)
air

# Or run directly
go run ./cmd/app/main.go
```

The server starts on `http://localhost:8080`.

### Running Tests

```bash
# Run all tests
go test ./...

# Run a single test
go test -run TestName ./package/path

# Run with verbose output
go test -v ./...
```

Note: Currently there are no test files in this project (`*_test.go`).

### Linting and Formatting

```bash
# Format code (gofmt)
gofmt -w .

# Run go vet
go vet ./...

# Download dependencies
go mod download
```

### Docker

```bash
# Build Docker image
docker build -t portfolio .

# Run container
docker run -p 8080:8080 portfolio
```

## Code Style Guidelines

### General Conventions

- Follow standard Go project layout: `cmd/` for binaries, packages in root
- Use Go modules (`go.mod`) for dependency management
- Run `gofmt` before committing

### Naming Conventions

- **Variables/Functions**: Use PascalCase for exported, camelCase for unexported
- **Files**: Use lowercase with underscores: `job_experience.go`
- **Packages**: Use short, lowercase names: `portfolio`, `templates`
- **Types**: Use PascalCase: `JobExperienceItemProps`

### Import Organization

Imports should be grouped (standard library first, then external):

```go
import (
	"log"
	"net/http"
	"slices"
	"strings"

	"github.com/a-h/templ"
	"github.com/felipefbs/portfolio/portfolio"
	"github.com/felipefbs/portfolio/templates"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)
```

### Error Handling

- Return errors explicitly; avoid ignoring them with `_`
- Use `log.Fatal(err)` for fatal errors in `main()`
- Handle render errors from templ components:

```go
if err != nil {
	return err
}
```

### Types and Interfaces

- Define props structs for template data (e.g., `JobExperienceItemProps`)
- Use pointers for struct types when mutability is needed
- Use generics for reusable utility functions:

```go
func removeDuplicate[T comparable](sliceList []T) []T
```

### Templ Guidelines

- Write `.templ` source files in `templates/` directory
- Run `templ generate` after modifying `.templ` files
- Do NOT edit generated `*_templ.go` files directly
- Generated files include lint ignores for certain rules

### Tailwind CSS

- Source CSS: `templates/input.css`
- Output: `static/styles/output.css`
- Rebuild CSS with: `tailwind -i ./templates/input.css -o ./static/styles/output.css`

### Routing (chi)

- Use `chi.NewRouter()` for router instance
- Define routes with `router.Get()`, `router.Handle()`, `router.HandleFunc()`
- Use middleware with `router.Use()`

## File Structure

```
.
├── cmd/app/main.go          # Application entry point
├── portfolio/               # Data models and content
│   ├── job.go
│   ├── project.go
│   └── education.go
├── templates/               # HTML templates
│   ├── *.templ              # Source template files (edit these)
│   ├── *_templ.go           # Generated (do not edit)
│   └── input.css            # Tailwind input
├── static/                  # Static assets
│   └── styles/output.css    # Generated CSS
├── .air.toml                # Air hot-reload config
├── go.mod                   # Go module definition
└── Dockerfile               # Docker configuration
```

## Common Tasks

### Adding a New Page

1. Create new `.templ` file in `templates/`
2. Run `templ generate -path ./templates`
3. Add route in `cmd/app/main.go`
4. Rebuild CSS if needed: `tailwind -i ./templates/input.css -o ./static/styles/output.css`

### Adding Portfolio Data

1. Edit files in `portfolio/` package (e.g., `job.go`)
2. Create or update corresponding templ component if needed

### Modifying Styles

1. Edit `templates/input.css`
2. Rebuild: `tailwind -i ./templates/input.css -o ./static/styles/output.css`

## Dependencies

- `github.com/a-h/templ` - HTML templating
- `github.com/go-chi/chi/v5` - HTTP routing
