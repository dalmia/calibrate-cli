# Calibrate CLI

Command-line client for the [Calibrate](https://calibrate.artpark.ai) Cloud **API** — run and manage your AI evals from CI. This can be used by your coding agents like Claude Code/Codex to directly communicate with Calibrate without you having to navigate the UI. 

## Install

```bash
brew install dalmia/tap/calibrate
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
