## calibrate tests bulk-create

Bulk create tests

### Synopsis

Create many test cases at once and link them to your agents

```
calibrate tests bulk-create [flags]
```

### Examples

```
  calibrate tests bulk-create --type tool_call --tests '[{"name":"<value>","conversation_history":[{"role":"tool"}],"evaluators":[{"evaluator_uuid":"f47ac10b-58cc-4372-a567-0e02b2c3d479"}]}]'
```

### Options

```
  -a, --agent-uuids string                  Agents (IDs) to link every created test to. Omit to link none
      --body string                         Request body as JSON (alternative to individual flags). Can also be provided via stdin.
  -h, --help                                help for bulk-create
  -l, --language config.settings.language   Language written to each test's config.settings.language. Omit to leave unset
      --tests string                        Test items to create (non-empty, max 500 per request, names unique within the batch) [required]
      --type response                       What the test judges:
                                            
                                            - response: judges the generated reply
                                            - `tool_call`: diffs the generated tool calls
                                            - `conversation`: judges the full conversation
                                            
                                            Applied to every test in the batch. (options: response, tool_call, conversation) [required]
      --x-api-key string                    string value
      --x-org-uuid string                   string value
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

* [calibrate tests](calibrate_tests.md)	 - Operations for tests
