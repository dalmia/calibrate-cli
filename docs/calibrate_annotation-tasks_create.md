## calibrate annotation-tasks create

Create annotation task

### Synopsis

Create an annotation task for annotators to label items against evaluators

```
calibrate annotation-tasks create [flags]
```

### Examples

```
  calibrate annotation-tasks create --name <value> --type conversation
```

### Options

```
      --body string            Request body as JSON (alternative to individual flags). Can also be provided via stdin.
      --description string     A description for the task. Omit for none
  -e, --evaluator-ids string   IDs of evaluators to link when the task is created, in order. Each must be one you created or a built-in default. Omit to create with no linked evaluators
  -h, --help                   help for create
  -n, --name string            Task name, unique within your workspace [required]
  -t, --type stt               Task type. Determines the shape of each item's payload.
                               - stt: judge a transcript on its own
                               - `llm`: judge one response with its conversation history
                               - `llm-general`: judge a standalone `input -> output` pair
                               - `conversation`: judge a full conversation (options: stt, llm, llm-general, conversation) [required]
  -x, --x-api-key string       string value
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

* [calibrate annotation-tasks](calibrate_annotation-tasks.md)	 - Operations for annotation-tasks
