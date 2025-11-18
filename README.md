# Nuke node_modules

[![Release](https://img.shields.io/github/v/release/vukan322/nuke-node_modules)](https://github.com/vukan322/nuke-node_modules/releases)
![Downloads](https://img.shields.io/github/downloads/vukan322/nuke-node_modules/total)
[![License: MIT](https://img.shields.io/badge/License-MIT-blue.svg)](https://opensource.org/licenses/MIT)
[![Go Version](https://img.shields.io/badge/Go-1.23+-00ADD8?logo=go)](https://go.dev/)

A fast CLI tool to find and delete stale `node_modules` directories, freeing up disk space.

## Why?

JavaScript projects accumulate `node_modules` folders that eat disk space. Backing up or syncing directories with thousands of tiny files in `node_modules` is painfully slow. This tool helps you reclaim space by finding and deleting old ones you're not actively using.

## Features

- **Smart filtering**: Only targets `node_modules` folders not modified in N days (default: 14)
- **Safe by default**: Skips hidden directories (`.nvm`, `.cache`, `.vscode`, etc.) to protect system tools
- **Dry run first**: `scan` command shows what would be deleted before you commit
- **Fast**: Skips symlinks, stays on same filesystem, uses approximate size calculation
- **Cross-platform**: Works on Linux, macOS, and Windows

## Installation

Download the latest binary from [releases](https://github.com/vukan322/nuke-node_modules/releases):

**Linux:**
```bash
curl -L https://github.com/vukan322/nuke-node_modules/releases/latest/download/nukenm-linux-amd64 -o nukenm
chmod +x nukenm
sudo mv nukenm /usr/local/bin/
```

**macOS (Intel):**
```bash
curl -L https://github.com/vukan322/nuke-node_modules/releases/latest/download/nukenm-darwin-amd64 -o nukenm
chmod +x nukenm
sudo mv nukenm /usr/local/bin/
```

**macOS (Apple Silicon):**
```bash
curl -L https://github.com/vukan322/nuke-node_modules/releases/latest/download/nukenm-darwin-arm64 -o nukenm
chmod +x nukenm
sudo mv nukenm /usr/local/bin/
```

**Windows:**
```bash
curl -L https://github.com/vukan322/nuke-node_modules/releases/latest/download/nukenm-windows-amd64.exe -o nukenm.exe
```

Then add to PATH or run directly: `.\nukenm.exe scan C:\Users\YourName\Documents`

## Usage

**Scan for old node_modules:**
```bash
nukenm scan ~/Documents
```

**Delete them:**
```bash
nukenm nuke ~/Documents
```

**Show verbose output:**
```bash
nukenm scan ~/Documents --verbose
```

**Change age threshold:**
```bash
nukenm scan ~/Documents --days 30
```

**Skip confirmation prompt:**
```bash
nukenm nuke ~/Documents -y
```

**Include hidden directories (dangerous):**
```bash
nukenm scan ~/ --include-hidden
```

## Flags

| Flag | Short | Default | Description |
|------|-------|---------|-------------|
| `--days` | - | `14` | Only process folders not modified in N days |
| `--verbose` | `-v` | `false` | Show detailed output during scan |
| `--yes` | `-y` | `false` | Skip confirmation prompt (nuke only) |
| `--include-hidden` | - | `false` | Include hidden directories (e.g., `.nvm`, `.cache`) |

## Safety

By default, `nukenm` protects your system by:
- **Skipping hidden directories** - Won't touch `.nvm`, `.npm`, `.cache`, `.vscode`, etc.
- **Age filtering** - Only deletes folders older than 14 days (configurable)
- **Confirmation prompt** - Asks before deleting (unless `-y` flag is used)
- **Same filesystem only** - Won't cross into mounted drives
- **Skips symlinks** - Avoids following symbolic links

⚠️ **Warning:** Using `--include-hidden` can break system tools like Node Version Manager. Only use it if you know what you're doing.

## Examples

**Find all old node_modules in your projects folder:**
```bash
nukenm scan ~/Projects
```

**Delete them after review:**
```bash
nukenm nuke ~/Projects
```

**Aggressive cleanup (0 days, no hidden protection):**
```bash
nukenm nuke ~/old-stuff --days 0 --include-hidden -y
```

**Check what's taking space before a backup:**
```bash
nukenm scan ~/ --verbose
```

## Build from Source
```bash
git clone https://github.com/vukan322/nuke-node_modules.git
cd nuke-node_modules
go build -o nukenm
```

OR

```bash
git clone https://github.com/vukan322/nuke-node_modules.git
cd nuke-node_modules
make install
```

**Other make targets:**
- `make build` - Build binary only
- `make man` - Generate man pages only
- `make release VERSION=v1.0.0` - Build release binaries for all platforms
- `make clean` - Remove build artifacts
- `make uninstall` - Remove from system

## License

MIT License - see [LICENSE](LICENSE) for details.

## Contributing

Issues and pull requests welcome.
