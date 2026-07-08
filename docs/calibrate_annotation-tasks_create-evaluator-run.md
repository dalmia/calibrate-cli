## calibrate annotation-tasks create-evaluator-run

Run evaluators on items

### Synopsis

Run evaluators on task items as a background job

```
calibrate annotation-tasks create-evaluator-run [flags]
```

### Examples

```
  calibrate annotation-tasks create-evaluator-run --task-uuid f47ac10b-58cc-4372-a567-0e02b2c3d479 --evaluators '[]'
```

### Options

```
      --body string                 Request body as JSON (alternative to individual flags). Can also be provided via stdin.
  -e, --evaluators string           The evaluators to run. Each must be linked to the task [required]
  -h, --help                        help for create-evaluator-run
  -i, --item-ids select_all=false   Item IDs to run on. **Required when select_all=false**. Ignored when `select_all=true`
      --q payload.name              Case-insensitive substring filter on payload.name. Applies only when `select_all=true`
  -s, --select-all true             When true, run on every item in the task. Set `q` to run only items whose name matches it
  -t, --task-uuid string            Annotation task to act on [required]
  -x, --x-api-key string            string value
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
