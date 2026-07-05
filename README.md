# Calibrate CLI

Command-line client for the [Calibrate](https://calibrate.artpark.ai) **public API** — run agent tests and manage agents from CI or your terminal.

> **Not** the offline evaluation engine. For local STT/TTS/LLM evals, install [`calibrate-agent`](https://pypi.org/project/calibrate-agent/) (`calibrate-agent` on the CLI).

## Install

**Homebrew** (after the tap is published):

```bash
brew install dalmia/tap/calibrate
```

**Install script** (macOS / Linux, from GitHub Releases):

```bash
curl -fsSL https://raw.githubusercontent.com/dalmia/calibrate-cli/main/scripts/install.sh | bash
```

**From source** (requires Go 1.25+):

```bash
go install github.com/dalmia/calibrate-cli/cmd/calibrate@latest
```

## Authentication

Create an org-scoped API key in Calibrate (**Settings → API keys**). Keys look like `sk_…`.

Pass it on every command:

```bash
export CALIBRATE_API_KEY_AUTH="sk_your_key_here"   # or use the flag below
calibrate --api-key-auth "$CALIBRATE_API_KEY_AUTH" agents list
```

Or store it interactively (keychain / config file):

```bash
calibrate auth login
calibrate whoami
```

For a self-hosted backend, point at your server:

```bash
calibrate --server-url https://your-calibrate.example.com agents list
```

## Examples

**List agents in your org:**

```bash
calibrate --api-key-auth "$CALIBRATE_API_KEY_AUTH" agents list
```

**Resolve agent names to UUIDs:**

```bash
calibrate --api-key-auth "$CALIBRATE_API_KEY_AUTH" \
  agents resolve --names my-agent --names another-agent
```

**Run tests for an agent** (returns a `task_id` to poll):

```bash
calibrate --api-key-auth "$CALIBRATE_API_KEY_AUTH" \
  agent-tests run --agent-uuid "<agent-uuid>"
```

**Poll until the run finishes:**

```bash
calibrate --api-key-auth "$CALIBRATE_API_KEY_AUTH" \
  agent-tests get-run --task-id "<task-id>"
```

**JSON output** (useful in scripts):

```bash
calibrate --api-key-auth "$CALIBRATE_API_KEY_AUTH" \
  -o json agents list | jq '.[0].name'
```

## More

- `calibrate --help` — all commands
- `calibrate explore` — interactive command browser
- [`docs/`](docs/) — per-command reference (auto-generated; may change on release)
