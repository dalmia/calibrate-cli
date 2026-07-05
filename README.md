# Calibrate CLI

Command-line client for [Calibrate](https://calibrate.artpark.ai), a framework for evaluating AI agents which lets you move from slow, manual testing to a fast, automated, and repeatable testing process for your entire agent stack. Use it to run and manage your evals from the terminal — including from coding agents like Claude Code or Codex — without navigating the UI.

## Install

```bash
brew install dalmia/tap/calibrate
```

## Usage

Log in once with your API key, which you can get from the [UI](https://calibrate.artpark.ai/workspace-settings?tab=api-keys):

```bash
calibrate auth login
```

Then run any command, for example:

```bash
calibrate agents list
```

## Self-hosting

For self-hosted Calibrate deployments, set `CALIBRATE_SERVER_URL` to point the CLI at your API instance before running commands:

```bash
export CALIBRATE_SERVER_URL=https://calibrate.my-company.internal
```

## Resources

- `calibrate --help` — all commands
- `calibrate explore` — interactive command browser
- [`docs/`](docs/) — per-command reference (auto-generated; may change on release)
