# Devtool-Cli!!

> A production-quality developer productivity tool built in Go with modular architecture, powered by [Cobra](https://github.com/spf13/cobra).

Automate common developer tasks — fetch APIs, format JSON, and organize messy folders — all from your terminal with colored output, rich error handling, and debug mode.

---

## ✨ Features

| Command | Description |
|---|---|
| `devtool get <url>` | Fetch API data with status codes, response time, and optional save-to-file |
| `devtool format <file>` | Validate and pretty-print JSON files with save support |
| `devtool clean <folder>` | Organize files into categorized subfolders with a detailed summary |
| `devtool version` | Display version, commit, and build info |

### Flags

| Flag | Scope | Description |
|---|---|---|
| `--save`, `-s` | `get`, `format` | Save output to a file instead of printing to terminal |
| `--verbose`, `-v` | Global | Enable debug logging (response headers, file paths, etc.) |

---

## 🛡️ Error Handling

This tool handles real-world edge cases gracefully:

```bash
# Invalid URL
$ devtool get not-a-url
✖ URL must include scheme and host (e.g., https://example.com), got: not-a-url

# Missing file
$ devtool format missing.json
✖ file not found: 'missing.json'

# Wrong file type
$ devtool format README.md
✖ expected a .json file, got '.md'

# Non-existent folder
$ devtool clean fakefolder
✖ directory not found: 'fakefolder'

# DNS failure
$ devtool get https://doesnotexist.invalid
✖ DNS resolution failed — host not found

# Request timeout (10s limit)
$ devtool get https://slow-server.example.com
✖ request timed out after 10s
```

---

## 📁 Project Structure

```
devtool-cli/
├── main.go                      # Minimal entry point
├── cmd/                         # Cobra command handlers
│   ├── root.go                  # Base CLI + global --verbose flag
│   ├── get.go                   # HTTP fetch command
│   ├── format.go                # JSON formatter command
│   ├── clean.go                 # File organizer command
│   └── version.go               # Version info command
├── internal/utils/              # Private, reusable helper functions
│   ├── http.go                  # HTTP client (timeout, validation, structured response)
│   ├── json.go                  # JSON parser (validation, formatting)
│   ├── file.go                  # File organizer (categorization, move summary)
│   └── logger.go                # Colored terminal output (ANSI)
├── go.mod
└── go.sum
```

---

## 🚀 Installation

### Prerequisites

- [Go 1.21+](https://go.dev/dl/)

### Build from source

```bash
git clone https://github.com/GhanshyamJha05/Devtool-cli.git
cd Devtool-cli
go build -o devtool .
```

---

## 📖 Usage

### Fetch API Data

```bash
# Print response to terminal (with status code & response time)
devtool get https://jsonplaceholder.typicode.com/posts/1

# Save response to a file
devtool get https://api.github.com/users/octocat --save user.json

# Debug mode — see headers, content-type, content-length
devtool get https://httpbin.org/get --verbose
```

### Format JSON

```bash
# Pretty-print a JSON file
devtool format config.json

# Save formatted output to a new file
devtool format data.json --save pretty.json
```

### Clean / Organize a Folder

```bash
# Organize files by type (Images, Documents, Code, Videos, etc.)
devtool clean ./downloads
```

Output:
```
ℹ Scanning folder: ./downloads...

✔ Images       → 5 file(s)
✔ Documents    → 3 file(s)
✔ Code         → 8 file(s)
✔ Archives     → 2 file(s)
⚠ Skipped 1 file(s) with no extension

✔ Done! Organized 18 file(s) into 4 categories
```

### Version

```bash
devtool version
# devtool-cli v1.0.0
#   commit : abc1234
#   built  : 2026-05-01
```

---

## 🏗️ Tech Stack

- **Language:** Go (Golang)
- **CLI Framework:** [Cobra](https://github.com/spf13/cobra)
- **Architecture:** Clean, modular — `cmd/` for commands, `internal/utils/` for business logic
- **Error Handling:** URL validation, file checks, timeouts, DNS errors, empty files
- **Output:** Colored terminal output using ANSI codes (zero dependencies)
- **Dependencies:** Only Cobra + pflag (minimal footprint)

---

## 🔮 Future Improvements

- [ ] `devtool encode/decode` — Base64 and URL encoding
- [ ] `devtool hash <file>` — MD5/SHA checksums
- [ ] `devtool serve <folder>` — Local HTTP file server
- [ ] Config file support (`.devtool.yaml`)
- [ ] Unit tests for all utilities
- [ ] Publish via `go install` for global installation

---

## 📄 License

MIT
