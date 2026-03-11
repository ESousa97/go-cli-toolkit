# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.1.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [1.0.0] - 2026-03-11

### Added
- Comprehensive test suite using `testify` for all core packages.
- `Makefile` for automated build, test, and installation workflows.
- `config.yaml` support via Viper for personalizing toolkit behavior.
- Advanced UI rendering with `lipgloss` for the `ping` command results.
- Proper concurrency model using Goroutines and Channels for multi-host pings.
- Support for pretty-printing JSON through the `format json` command.
- Detailed architecture and internal documentation (godoc comments).
- Professional project governance files (`LICENSE`, `CONTRIBUTING.md`, `CODE_OF_CONDUCT.md`, `SECURITY.md`).

### Changed
- Refactored `ping` command to support concurrent execution and multiple targets.
- Standardized project layout following `cmd/` and `internal/` pattern.
- Updated project requirements to Go 1.25.

## [0.1.0] - 2026-03-10

### Added
- Initial project structure with Cobra CLI.
- Basic `ping` command for single host checking.
- Basic JSON validation command.
