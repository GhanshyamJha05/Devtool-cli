# devtool-cli

A small Go CLI for everyday developer chores.

It can fetch API responses, format JSON files, and organize a messy folder by file type. I built it as a simple terminal tool, with readable output and clear errors.

## Preview



![image](https://github.com/user-attachments/assets/7b2c3e42-ac62-418b-b5f8-39e4e0b40d0d)


Live site:

```text
https://ghanshyamjha05.github.io/devtool-cli/
```

## What it does

| Command | What it does |
| --- | --- |
| `get` | Fetches data from an HTTP/HTTPS URL |
| `format` | Validates and pretty-prints a JSON file |
| `clean` | Moves files into folders like Images, Documents, Code, Videos, etc. |
| `version` | Prints the installed version |

## Install

You need Go installed first.

```bash
go install github.com/GhanshyamJha05/devtool-cli@latest
```

On Windows, Go usually installs the binary here:

```text
C:\Users\<your-name>\go\bin
```

If the command is not recognized, add that folder to your PATH and reopen the terminal.

The installed command name is:

```bash
devtool-cli
```

## Usage

Check the available commands:

```bash
devtool-cli --help
```

Fetch an API response:

```bash
devtool-cli get https://jsonplaceholder.typicode.com/posts/1
```

Save the response to a file:

```bash
devtool-cli get https://jsonplaceholder.typicode.com/posts/1 --save post.json
```

Format a JSON file:

```bash
devtool-cli format ugly.json
```

Save formatted JSON:

```bash
devtool-cli format ugly.json --save pretty.json
```

Clean a folder:

```bash
devtool-cli clean "C:\Users\you\Downloads"
```

Use quotes when the folder path has spaces.

Show extra details:

```bash
devtool-cli clean "C:\Users\you\Downloads" --verbose
```

## Using it on another system

Keep these things in mind:

- Install Go first.
- Use the same install command: `go install github.com/GhanshyamJha05/devtool-cli@latest`.
- Make sure Go's `bin` folder is in PATH.
- On Windows, the binary will usually be `devtool-cli.exe`.
- On macOS/Linux, it will usually be `devtool-cli`.
- Internet is needed for `go install` and for the `get` command.
- The `clean` command moves files, so test it on a sample folder first.
- If a path has spaces, wrap it in quotes.

Quick check after installing:

```bash
devtool-cli version
devtool-cli --help
```

## Build from source

```bash
git clone https://github.com/GhanshyamJha05/devtool-cli.git
cd devtool-cli
go build -o devtool-cli .
```

Then run:

```bash
./devtool-cli --help
```

On Windows Command Prompt:

```cmd
devtool-cli.exe --help
```

## Project structure

```text
devtool-cli/
  cmd/              command setup
  internal/utils/   core logic
  docs/             GitHub Pages site
  testdata/         sample files
  main.go
  go.mod
```

## Notes

This is still a small project. Good next upgrades would be tests, a dry-run option for `clean`, and better release automation.

## License

MIT
