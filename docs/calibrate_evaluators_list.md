## calibrate evaluators list

List evaluators

### Synopsis

List your evaluators

```
calibrate evaluators list [flags]
```

### Examples

```
  calibrate evaluators list
```

### Options

```
      --data-type string        Filter by modality. Omit for all (options: text, audio)
  -e, --evaluator-type string   Filter by what the evaluator judges. Omit for all types (options: tts, stt, llm, llm-general, conversation)
  -h, --help                    help for list
  -i, --include-defaults true   When true, include the built-in default evaluators alongside the ones you created (default true)
  -x, --x-api-key string        string value
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

* [calibrate evaluators](calibrate_evaluators.md)	 - Operations for evaluators
