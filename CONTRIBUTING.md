# Contributing to Go CLI Toolkit

> [!CAUTION]
> **This project is ARCHIVED.**
> It serves as a personal learning record for the Go language. As such, I am no longer actively reviewing Pull Requests or maintaining the codebase. You are encouraged to fork this repository for your own study and experimentation.

---

This document remains here for historical reference, documenting the standards and workflow used during the development of this project.

## Prerequisites

| Tool | Version | Purpose |
|------|---------|---------|
| Go | >= 1.25 | Language runtime |
| Git | >= 2.x | Version control |
| Make | any | Build automation |

## Setting Up the Development Environment

```bash
git clone https://github.com/ESousa97/go-cli-toolkit.git
cd go-cli-toolkit
go mod download
make build
```

Verify the setup by running the test suite:

```bash
make test
```

## Code Style and Conventions

This project follows standard Go conventions. Before submitting code, review these references:

- [Effective Go](https://go.dev/doc/effective_go)
- [Go Code Review Comments](https://github.com/golang/go/wiki/CodeReviewComments)
- [Standard Go Project Layout](https://github.com/golang-standards/project-layout)

### Key Rules

1. **Single responsibility per file.** UI components or command files exceeding 100 lines should be split.
2. **No magic values.** Extract strings, colors, timeouts, and configuration to named constants or config files.
3. **Doc comments on all exports.** Every exported function, type, constant, and variable must have a doc comment starting with the item name.
4. **Error wrapping.** Use `fmt.Errorf("context: %w", err)` for all error returns.

## Running Quality Checks

```bash
# Unit tests with verbose output
make test

# Static analysis
go vet ./...

# Build verification
make build
```

## Branch Naming (Historical)

| Prefix | Purpose | Example |
|--------|---------|---------|
| `feat/` | New features | `feat/add-csv-formatter` |
| `fix/` | Bug fixes | `fix/ping-timeout-handling` |
| `docs/` | Documentation only | `docs/improve-readme` |
| `refactor/` | Code reorganization | `refactor/extract-http-client` |
| `test/` | Test additions or fixes | `test/add-config-edge-cases` |

## Commit Convention

Follow [Conventional Commits](https://www.conventionalcommits.org/en/v1.0.0/):

```
<type>: <short description>

[optional body]
```

## License

This project is licensed under the **MIT License** — see the [LICENSE](LICENSE) file for details.
