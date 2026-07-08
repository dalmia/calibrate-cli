## calibrate evaluators create-version

Create evaluator version

### Synopsis

Add a new version to an evaluator you created

```
calibrate evaluators create-version [flags]
```

### Examples

```
  calibrate evaluators create-version --evaluator-uuid f47ac10b-58cc-4372-a567-0e02b2c3d479 --judge-model <value> --system-prompt <value>
```

### Options

```
      --body string                    Request body as JSON (alternative to individual flags). Can also be provided via stdin.
  -e, --evaluator-uuid string          Evaluator to add a version to [required]
  -h, --help                           help for create-version
  -j, --judge-model openai/gpt-4.1     The model that runs the judge, named the way its provider does, for example openai/gpt-4.1 or `anthropic/claude-sonnet-4` [required]
  -m, --make-live true                 When true, immediately point the evaluator's live version at this new version
      --output-config rating           The scale points and their labels. Required for a rating evaluator. A `binary` evaluator uses the default Correct/Wrong labels unless you set your own
  -s, --system-prompt { {variable} }   Judge system prompt. May contain { {variable} } placeholders [required]
  -v, --variables string               Declared prompt variables. Omit if the prompt has none. After the first version the variable names are fixed. You can change a variable's description or default, but not add, remove, or rename one
  -x, --x-api-key string               string value
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
