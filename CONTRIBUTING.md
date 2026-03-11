# Contributing to Go CLI Toolkit

Thank you for considering a contribution. This document describes the development setup, conventions, and workflow required for submitting changes.

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

All checks must pass before opening a Pull Request.

## Branch Naming

Use the following prefixes:

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

Examples:

```
feat: add CSV output format to ping command
fix: handle empty stdin gracefully in format json
docs: document config.yaml schema
test: add edge cases for concurrent ping
```

## Pull Request Process

1. Fork the repository and create your branch from `main`.
2. Make your changes following the conventions above.
3. Run `make test` and `go vet ./...` locally.
4. Open a Pull Request with a clear title and description.
5. Wait for review. Address feedback with fixup commits, then squash before merge.

## Areas Where Contributions Are Welcome

- New formatter subcommands (CSV, YAML, TOML)
- Output export options (file, clipboard)
- Additional ping protocols (TCP, ICMP)
- Performance improvements in concurrent operations
- Documentation improvements and translations
- Test coverage expansion
