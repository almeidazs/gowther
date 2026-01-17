# Contributing

Thanks for your interest in contributing to Serenity.

This repository contains a Go linter and migration tool focused on
performance, correctness, and explicit configuration.

## Requirements

- Go >= 1.22
- Linux, macOS or Windows
- Git

## Development setup

```bash
git clone https://github.com/serenitysz/serenity.git
cd serenity
go mod download
```

## Running locally

```bash
go test ./...
go run .
```

## Code style

- Follow standard Go formatting (`gofmt`)
- Prefer explicit code over abstractions
- Avoid unnecessary allocations and reflection
- Keep logic simple and predictable

## Adding or modifying rules

- Each rule should be:
  - deterministic
  - fast (no unnecessary allocations)
  - well-scoped

- Migration logic must be explicit and reversible
- Avoid magic defaults

## Commit messages

Use clear, concise commit messages.
Conventional Commits are recommended but not required.

Examples:

- `feat: add new rule for max params`
- `fix: handle empty revive config`
- `chore: update dependencies`

## Pull requests

- Keep PRs focused
- Avoid unrelated refactors
- Add tests when changing behavior
- Explain why the change is needed
