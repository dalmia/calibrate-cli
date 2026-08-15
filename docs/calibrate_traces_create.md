## calibrate traces create

Create trace

### Synopsis

Store a production agent turn and its conversation history for later review

```
calibrate traces create [flags]
```

### Examples

```
  calibrate traces create --agent-id <id> --input '[{"role":"<value>"}]' --output-param '{}'
```

### Options

```
  -a, --agent-id string          ID of the agent that produced the turn. Must be an agent in your workspace [required]
      --body string              Request body as JSON (alternative to individual flags). Can also be provided via stdin.
  -c, --conversation-id string   Your own ID for the conversation this turn belongs to, stored for reference only. Omit if you have none
  -h, --help                     help for create
  -i, --input string             Conversation history up to the reported output, oldest turn first, in OpenAI chat format [required]
      --message-id input         Your own ID for the last user message in input, stored for reference only. Omit if you have none
      --metadata gen_ai.*        Key-value pairs stored with the trace. Prefer OTel gen_ai.* key names where they fit. Omit if you have none
      --output-param string      [required]
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

* [calibrate traces](calibrate_traces.md)	 - Operations for traces
