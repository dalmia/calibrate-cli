## calibrate annotation-tasks

Operations for annotation-tasks

### Synopsis

Operations for annotation-tasks

```
calibrate annotation-tasks [flags]
```

### Options

```
  -h, --help   help for annotation-tasks
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

* [calibrate](calibrate.md)	 - Calibrate Public API: Programmatic API for CI/automation
* [calibrate annotation-tasks add-items](calibrate_annotation-tasks_add-items.md)	 - Bulk create items
* [calibrate annotation-tasks create](calibrate_annotation-tasks_create.md)	 - Create annotation task
* [calibrate annotation-tasks create-evaluator-run](calibrate_annotation-tasks_create-evaluator-run.md)	 - Run evaluators on items
* [calibrate annotation-tasks get](calibrate_annotation-tasks_get.md)	 - Get annotation task
* [calibrate annotation-tasks get-agreement](calibrate_annotation-tasks_get-agreement.md)	 - Get task agreement
* [calibrate annotation-tasks get-evaluator-run](calibrate_annotation-tasks_get-evaluator-run.md)	 - Get evaluator run
* [calibrate annotation-tasks get-summary](calibrate_annotation-tasks_get-summary.md)	 - Get task summary
* [calibrate annotation-tasks list](calibrate_annotation-tasks_list.md)	 - List annotation tasks
* [calibrate annotation-tasks set-evaluators](calibrate_annotation-tasks_set-evaluators.md)	 - Update task evaluators
* [calibrate annotation-tasks update-items](calibrate_annotation-tasks_update-items.md)	 - Bulk update items
