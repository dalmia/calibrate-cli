## calibrate

Calibrate Public API: Programmatic API for CI/automation

### Synopsis

Calibrate Public API: Programmatic API for CI/automation. Pass your key in the `X-API-Key` header.

```
calibrate [flags]
```

### Options

```
      --agent-mode             Enable structured errors and default TOON output for AI coding agents. Automatically enabled when a known agent environment is detected (CLAUDE_CODE, CURSOR_AGENT, etc.). Use --agent-mode=false to disable.
      --api-key-auth string    Workspace API key. Create one under Workspace settings → API keys.
      --color string           Control colored output: auto (color when output is a TTY), always, or never. Respects NO_COLOR and FORCE_COLOR env vars. (default "auto")
  -d, --debug                  Log request and response diagnostics to stderr
      --dry-run                Preview the request that would be sent without executing it (output to stderr)
  -H, --header stringArray     Set a custom HTTP request header (format: "Key: Value"). Can be specified multiple times.
  -h, --help                   help for calibrate
      --include-headers        Include HTTP response headers in the output
  -q, --jq string              Filter and transform output using a jq expression (e.g., '.name', '.items[] | .id')
      --no-interactive         Disable all interactive features (auto-prompting, explorer auto-launch, TUI forms)
  -o, --output-format string   Specify the output format. Options: pretty, json, yaml, table, toon. (default "pretty")
      --server string          Select a server by index (for indexed servers) or name (for named servers)
      --server-url string      Override the default server URL
      --timeout string         HTTP request timeout (e.g., 30s, 5m, 100ms)
      --usage                  Print the CLI Usage schema in KDL format
```

### SEE ALSO

* [calibrate agent-tests](calibrate_agent-tests.md)	 - Operations for agent-tests
* [calibrate agents](calibrate_agents.md)	 - Operations for agents
* [calibrate annotation-tasks](calibrate_annotation-tasks.md)	 - Operations for annotation-tasks
* [calibrate auth](calibrate_auth.md)	 - Manage authentication credentials
* [calibrate configure](calibrate_configure.md)	 - Configure authentication credentials and preferences
* [calibrate evaluators](calibrate_evaluators.md)	 - Operations for evaluators
* [calibrate explore](calibrate_explore.md)	 - Interactively browse and run commands
* [calibrate tests](calibrate_tests.md)	 - Operations for tests
* [calibrate version](calibrate_version.md)	 - Print the CLI version
* [calibrate whoami](calibrate_whoami.md)	 - Display current authentication configuration
