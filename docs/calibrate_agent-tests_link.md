## calibrate agent-tests link

Link tests to agent

### Synopsis

Link one or more tests to an agent. Tests that are already linked are skipped.

```
calibrate agent-tests link [flags]
```

### Examples

```
  calibrate agent-tests link --agent-uuid f47ac10b-58cc-4372-a567-0e02b2c3d479 --test-uuids '["b1c2d3e4-f5a6-7890-bcde-f12345678901"]'
```

### Options

```
  -a, --agent-uuid string        Agent to link tests to [required]
      --body string              Request body as JSON (alternative to individual flags). Can also be provided via stdin.
  -h, --help                     help for link
  -t, --test-uuids stringArray   Tests to link. Any that are already linked are skipped [required]
  -x, --x-api-key string         string value
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

* [calibrate agent-tests](calibrate_agent-tests.md)	 - Operations for agent-tests
