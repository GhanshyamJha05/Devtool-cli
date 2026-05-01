# devtool-cli

A production-quality, modular CLI developer tool built in Go using [Cobra](https://github.com/spf13/cobra).

Automate common developer tasks — fetch APIs, format JSON files, and organize messy folders — all from your terminal.

---

## Features

| Command | Description |
|---|---|
| `devtool get <url>` | Fetch API data and display the response |
| `devtool format <file.json>` | Pretty-print a JSON file |
| `devtool clean <folder>` | Organize files into categorized subfolders |

### Flags

| Flag | Scope | Description |
|---|---|---|
| `--save`, `-s` | `get`, `format` | Save output to a file instead of printing to terminal |
| `--verbose`, `-v` | Global | Enable debug logging for all commands |

---

## Project Structure

```
devtool-cli/
├── main.go                  # Entry point
├── cmd/                     # Cobra command handlers
│   ├── root.go              # Base CLI setup & global flags
│   ├── get.go               # HTTP fetch command
│   ├── format.go            # JSON formatter command
│   └── clean.go             # File organizer command
├── internal/utils/          # Reusable helper functions
│   ├── http.go              # HTTP client utility
│   ├── json.go              # JSON formatting utility
│   ├── file.go              # File organization utility
│   └── logger.go            # Colored terminal output
├── go.mod
└── go.sum
```

---

## Installation

### Prerequisites

- [Go 1.21+](https://go.dev/dl/)

### Build from source

```bash
git clone https://github.com/YOUR_USERNAME/devtool-cli.git
cd devtool-cli
go build -o devtool .
```

The binary `devtool` (or `devtool.exe` on Windows) will be created in the project root.

---

## Usage

### Fetch API data

```bash
# Print response to terminal
devtool get https://jsonplaceholder.typicode.com/posts/1

# Save response to a file
devtool get https://jsonplaceholder.typicode.com/posts/1 --save output.json
```

### Format JSON

```bash
# Pretty-print a JSON file
devtool format data.json

# Save formatted output to a new file
devtool format data.json --save pretty.json
```

### Clean / Organize a folder

```bash
# Organize files in ~/Downloads by type
devtool clean ~/Downloads
```

Files are moved into subfolders like `Images/`, `Documents/`, `Code/`, `Videos/`, `Audio/`, `Archives/`, and `Others/`.

### Debug mode

Add `--verbose` (or `-v`) to any command to see debug output:

```bash
devtool get https://api.example.com/data --verbose
```

---

## Tech Stack

- **Language:** Go (Golang)
- **CLI Framework:** [Cobra](https://github.com/spf13/cobra)
- **Architecture:** Clean, modular — commands and utilities are fully separated
- **Dependencies:** Only Cobra + pflag (zero bloat)

---

## Future Improvements

- [ ] Add `devtool encode/decode` for Base64 and URL encoding
- [ ] Add `devtool hash <file>` for MD5/SHA checksums
- [ ] Add `devtool serve <folder>` to spin up a local HTTP file server
- [ ] Add config file support (`.devtool.yaml`)
- [ ] Publish via `go install` for global installation
- [ ] Add unit tests for all utilities

---

## License

MIT
