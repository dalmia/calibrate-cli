## calibrate agent-tests get-run-case

Get test case result

### Synopsis

Get the full result of one test case in a run

```
calibrate agent-tests get-run-case [flags]
```

### Examples

```
  calibrate agent-tests get-run-case --task-id a3b2c1d0-e5f4-3210-abcd-ef1234567890 --test-uuid b1c2d3e4-f5a6-7890-bcde-f12345678901
```

### Options

```
  -h, --help                  help for get-run-case
  -m, --model string          Which model's answer to read. Required for a benchmark, which runs every test once per model
      --task-id string        Test run or benchmark the case was run in [required]
      --test-uuid test_uuid   The test whose result to read, as test_uuid on the case [required]
  -x, --x-api-key string      string value
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
