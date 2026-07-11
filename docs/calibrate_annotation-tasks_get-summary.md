## calibrate annotation-tasks get-summary

Get task summary

### Synopsis

Get a paginated summary table of items, evaluator runs, and human annotations for a task

```
calibrate annotation-tasks get-summary [flags]
```

### Examples

```
  calibrate annotation-tasks get-summary --task-uuid f47ac10b-58cc-4372-a567-0e02b2c3d479
```

### Options

```
  -c, --compact rows.payload   Return a compact response that omits heavy detail fields (rows.payload, `rows.evaluator_reasoning`, `rows.annotations.reasoning`, `evaluators.versions.system_prompt`, `evaluators.versions.output_config`, `evaluators.versions.variables`, `item_comments`), keeping only the lightweight decision fields. Omit for full detail
      --disagreement-only      When true, keep only rows where the evaluator disagreed with at least one annotator
  -h, --help                   help for get-summary
  -i, --item-id annotators     Filter rows to a single item. The full task-wide annotator union is still returned in annotators
      --limit int              Maximum number of items to return (default 50)
      --live-only              When true, emit only one row for each (item, evaluator) pair using the evaluator's live version. Versions other than the live one that have runs are excluded
      --offset int             Number of items to skip before returning results
      --order string           Sort direction (options: asc, desc) (default "desc")
      --q payload.name         Case-insensitive substring search on payload.name. Blank is a no-op
  -s, --sort-by string         Sort key for the results (default "created_at")
  -t, --task-uuid string       Annotation task to act on [required]
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
