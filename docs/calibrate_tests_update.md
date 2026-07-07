## calibrate tests update

Update test

### Synopsis

Update a test's name, config, and/or evaluator links.

```
calibrate tests update [flags]
```

### Examples

```
  calibrate tests update --test-uuid b1c2d3e4-f5a6-7890-bcde-f12345678901
```

### Options

```
      --body string               Request body as JSON (alternative to individual flags). Can also be provided via stdin.
  -c, --config-param string       Replacement calibrate config. Omit to leave unchanged
  -e, --evaluators conversation   Replacement evaluator links (replaces the existing set). Omit to leave links unchanged; an empty list clears them (**rejected for conversation tests**)
  -h, --help                      help for update
  -n, --name string               New test name. Omit to leave unchanged
      --test-uuid string          Test to update [required]
      --type string               Test type. Immutable — may only echo the existing value; a different value is rejected (400). Omit to leave unchanged (options: response, tool_call, conversation)
      --x-api-key string          string value
      --x-org-uuid string         string value
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

* [calibrate tests](calibrate_tests.md)	 - Operations for tests
