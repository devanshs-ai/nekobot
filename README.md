<div align="center">
  <h1>🐈 nekobotic</h1>
  <p><b>A cinematic, system-aware terminal companion for developers.</b></p>
  
  [![Go Report Card](https://goreportcard.com/badge/github.com/devanshs-ai/nekobot)](https://goreportcard.com/report/github.com/devanshs-ai/nekobot)
  [![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)
</div>

`nekobotic` is a lightweight, high-performance terminal daemon written in Go. Inspired by the *Ghost in the Shell* and *Interstellar* aesthetics, it lives in your terminal prompt, monitoring your system's "health" and tracking your research/coding sessions.

---

## 🌌 Features

- **Reactive ASCII Pet**: Your companion's mood shifts based on system load. Watch it get stressed when your CPU peaks or lonely when you haven't committed code.
- **XAI Diagnostics**: Unlike basic monitors, `nekobotic` performs a "feature importance" scan to identify the specific process (Brave, Chrome, Docker) causing system pressure.
- **Research Mode**: Automatically tracks active coding time per project. It recognizes which folder you're in and logs your "Neural Sync" time.
- **Minimalist Aesthetic**: Designed with a high-contrast **Cyber Teal** and **Vengeance Red** palette.

---

## 🛠️ Tech Stack

- **Language**: [Go (Golang)](https://go.dev/) for near-instant execution.
- **Hardware Telemetry**: [gopsutil](https://github.com/shirou/gopsutil) for cross-platform resource monitoring.
- **Persistence**: Centralized JSON-based ledger for lifetime activity tracking.

---

## 🚀 Quick Start

### Prerequisites
- [Go](https://go.dev/doc/install) (1.21 or higher)
- Windows (Git Bash / PowerShell), Linux, or macOS

### Installation

1. **Clone the repository:**
   ```bash
   git clone https://github.com/devanshs-ai/neko-bot.git
   cd neko-bot
   ```

2. **Build the binary:**
   ```bash
   go build -o nekobotic.exe ./cmd/nekobotic/main.go
   ```

### Configuration

**For Bash / Zsh:**
Add the daemon to your prompt by appending this to your `~/.bashrc` or `~/.zshrc`:
```bash
# Run nekobotic on every command
export PROMPT_COMMAND="/absolute/path/to/nekobotic.exe; $PROMPT_COMMAND"
```
Reload your shell:
```bash
source ~/.bashrc
```

**For PowerShell:**
Add the following snippet to your PowerShell `$PROFILE` (open by typing `notepad $PROFILE` in PowerShell):
```powershell
function prompt {
    # Replace with the actual absolute path to your binary
    & "C:\absolute\path\to\nekobotic.exe" 
    return "PS $($executionContext.SessionState.Path.CurrentLocation)$('>' * ($nestedPromptLevel + 1)) "
}
```

---

## 🧪 Diagnostic Logic

The Neko uses a priority-weighted scan to determine the system's operational state:

- **Memory Threshold (>75%)**: Identifies and flags top RAM-consuming processes.
- **CPU Threshold (>50%)**: Identifies and flags top CPU-consuming tasks.
- **Productivity Check**: Scans local Git history to verify recent developer activity.

---

## 📂 Project Structure

```text
nekobotic/
├── cmd/nekobotic/    # Entry point
├── internal/
│   ├── ascii/        # Visual frames & color logic
│   ├── stats/        # Telemetry & process monitoring
│   └── config/       # (Coming Soon) YAML configuration
└── README.md
```

---

## 🤝 Contributing

Contributions, issues, and feature requests are welcome! We are currently looking to add:

- [ ] **YAML Configuration support** (Custom thresholds and colors)
- [ ] **Discord Rich Presence integration**
- [ ] **More ASCII pet skins** (Tachikoma, Bit-Ghost, etc.)

---

<div align="center">
  <i>"The Net is vast and infinite..."</i>
</div>