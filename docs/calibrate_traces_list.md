## calibrate traces list

List traces

### Synopsis

List ingested traces, newest first

```
calibrate traces list [flags]
```

### Examples

```
  calibrate traces list
```

### Options

```
  -a, --agent-id string                    Return only traces from this agent
      --conversation-id agent_id           Return only the turns of this conversation. Requires agent_id
      --created-after string               Return only traces created at or after this time (ISO 8601). A time with no timezone is read as UTC
      --created-before string              Return only traces created before this time (ISO 8601). A time with no timezone is read as UTC
  -h, --help                               help for list
      --include-content limit              Also return each trace's full input, output and metadata. Use limit and `offset` to page through them
      --include-score-averages             Also return each evaluator's mean score over every matching trace, not just this page. Costs a pass over those traces, so ask for it when the filters change rather than on every page
      --labels string                      Return only traces carrying at least one of these labels. Repeat the parameter for each label
      --limit int                          Maximum number of items to return (default 50)
      --offset int                         Number of items to skip before returning results
      --output-type response               Return only traces whose output is of this kind. response covers every trace carrying a reply, including one that also issued tool calls. `tool_call` covers traces that only issued tool calls (options: response, tool_call)
      --q string                           Return only traces containing this text in their message ID, conversation ID, conversation history, output, or metadata
      --score <evaluator ID>:<condition>   Return only traces whose latest scoring run matches this condition. Each is written as <evaluator ID>:<condition>, where the condition is `passed`, `failed`, or one of `=`, `!=`, `<`, `<=`, `>`, `>=` followed by a number, such as `>=4`. Repeat the parameter to require every condition at once
      --sort-by-evaluator string           Order traces by this evaluator's score on each trace's latest scoring run. Traces it has not scored come last whichever direction you pick
      --sort-order sort_by_evaluator       Whether sort_by_evaluator puts the highest scores first or the lowest. Ignored without it (options: asc, desc) (default "desc")
  -x, --x-api-key string                   string value
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
