## calibrate agents update

Update agent

### Synopsis

Update an agent's name and/or config. Changing `agent_url` or `agent_headers` resets connection and benchmark verification flags.

```
calibrate agents update [flags]
```

### Examples

```
  calibrate agents update --agent-uuid f47ac10b-58cc-4372-a567-0e02b2c3d479
```

### Options

```
  -a, --agent-uuid string                         The agent to update. Must be in your workspace. [required]
  -b, --benchmark-models-verified string          Directly set the per-model benchmark verification map inside config. Omit to leave it untouched
      --body string                               Request body as JSON (alternative to individual flags). Can also be provided via stdin.
      --config-param agent_url                    Replacement config, stored as-is (no deep-merge on update). Changing agent_url or `agent_headers` resets all connection/benchmark verification flags. Omit to leave config unchanged
      --connection-verified connection_verified   Directly set the connection_verified flag inside config. Omit to leave it untouched
  -h, --help                                      help for update
  -n, --name string                               New agent name. Omit to leave the name unchanged
      --x-api-key string                          string value
      --x-org-uuid string                         string value
```

### Options inherited from parent commands

```
      --agent-mode             Enable structured errors and default TOON output for AI coding agents. Automatically enabled when a known agent environment is detected (CLAUDE_CODE, CURSOR_AGENT, etc.). Use --agent-mode=false to disable.
      --api-key-auth string    Workspace API key. Create one under Workspace settings → API keys.
      --color string           Control colored output: auto (color when output is a TTY), always, or never. Respects NO_COLOR and FORCE_COLOR env vars. (default "auto")
  -d, --debug                  Log request and response diagnostics to stderr
      --dry-run                Preview the request that would be sent without executing it (output to stderr)
  -H, --header stringArray     Set a custom HTTP request header (format: "Key: Value"). Can be specified multiple times.
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

* [calibrate agents](calibrate_agents.md)	 - Operations for agents
