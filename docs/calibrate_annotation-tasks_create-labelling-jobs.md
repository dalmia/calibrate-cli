## calibrate annotation-tasks create-labelling-jobs

Create labelling jobs

### Synopsis

Assign items to annotators, creating one labelling job per annotator

```
calibrate annotation-tasks create-labelling-jobs [flags]
```

### Examples

```
  calibrate annotation-tasks create-labelling-jobs --task-uuid f47ac10b-58cc-4372-a567-0e02b2c3d479 --annotator-ids '["f47ac10b-58cc-4372-a567-0e02b2c3d479"]'
```

### Options

```
  -a, --annotator-ids stringArray   Annotator IDs to assign, creating one labelling job for each annotator. Must be in your workspace [required]
      --body string                 Request body as JSON (alternative to individual flags). Can also be provided via stdin.
  -c, --comments-enabled true       When true, the labelling form lets the annotator leave a comment on each item (default true)
  -e, --evaluator-ids None          Subset of the task's linked evaluators to show in these jobs. Must be a subset of the current links, an empty list gives a 400. Applies to every annotator's job. Omit (None) to snapshot every linked evaluator
  -h, --help                        help for create-labelling-jobs
  -i, --item-ids select_all=false   Item IDs to assign. **Required when select_all=false**. Ignored when `select_all=true`
      --q payload.name              Case-insensitive substring filter on payload.name. Applies only when `select_all=true`
  -r, --reasoning-mode optional     How the labelling form treats the reasoning box on each judgement. optional shows it, `required` shows it and asks the annotator to fill it in, `hidden` leaves it out (options: optional, required, hidden) (default "optional")
  -s, --select-all true             When true, assign every item in the task and ignore `item_ids`. Set `q` to assign only items whose name matches it
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
