## calibrate agent-tests get-run

Get test run status

### Synopsis

Poll a test run for its status and evaluation results.

```
calibrate agent-tests get-run [flags]
```

### Examples

```
  calibrate agent-tests get-run --task-id a3b2c1d0-e5f4-3210-abcd-ef1234567890
```

### Options

```
  -c, --compact results.output   Return a compact response that omits heavy detail fields (results.output, `results.test_case`, `results.judge_results`, `results.reasoning`, `evaluators.output_config`), keeping only the lightweight decision fields. Omit for full detail
  -h, --help                     help for get-run
      --only-failed              Return only failing test cases. Omit to return every case
  -t, --task-id string           Test run to poll for status and results [required]
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
