## calibrate tests create

Create test

### Synopsis

Create a test that runs your agent against a conversation and evaluates its answer quality or the tools it calls

```
calibrate tests create [flags]
```

### Examples

```
  calibrate tests create --name <value> --type tool_call
```

### Options

```
      --body string            Request body as JSON (alternative to individual flags). Can also be provided via stdin.
  -c, --config-param history   The calibrate test config. Three top-level keys.
                               
                               - history: the required conversation up to the agent's turn. Each item is `{role, content}` with `role` one of `user`, `assistant`, `tool`. A `tool` message also carries `tool_call_id` and `name`.
                               - `evaluation`: the required `{type, ...}`, where `type` matches the test's `type` below.
                               - `settings`: an optional object, e.g. `{"language": "en"}`.
                               
                               `evaluation` by test type:
                               - `response`: judge the agent's reply, graded by the linked evaluators. `{"type": "response"}`
                               - `conversation`: append the reply and judge the whole conversation. `{"type": "conversation"}`
                               - `tool_call`: diff the agent's tool calls against expected ones. Add `tool_calls`, a list of `{tool, arguments, accept_any_arguments?}`.
                               
                               For `tool_call`, each expected argument value is one of:
                               - `{"match_type": "exact", "value": <any>}`: must equal `value`
                               - `{"match_type": "llm_judge", "criteria": "..."}`: judged against the criteria
                               - `{"match_type": "any"}`: any value, only checks the argument was passed
                               
                               `response` / `conversation` example:
                               ```json
                               {
                                 "history": [{"role": "user", "content": "What is your return policy?"}],
                                 "evaluation": {"type": "response"},
                                 "settings": {"language": "en"}
                               }
                               ```
                               
                               `tool_call` example:
                               ```json
                               {
                                 "history": [{"role": "user", "content": "Book room 101 for tomorrow"}],
                                 "evaluation": {
                                   "type": "tool_call",
                                   "tool_calls": [
                                     {
                                       "tool": "book_room",
                                       "arguments": {
                                         "room": {"match_type": "exact", "value": "101"},
                                         "date": {"match_type": "llm_judge", "criteria": "tomorrow's date"}
                                       },
                                       "accept_any_arguments": false
                                     }
                                   ]
                                 }
                               }
                               ```
                               
                               Evaluators are linked via the separate `evaluators` field, not inside `config`.
                               
                               Omit to create the test with no config and fill it in later via update
  -e, --evaluators response    Evaluators to link. Used by response and `conversation` tests
  -h, --help                   help for create
  -n, --name string            Name of the test, unique within the workspace [required]
  -t, --type response          What the test judges:
                               
                               - response: judges the generated reply
                               - `tool_call`: diffs the generated tool calls
                               - `conversation`: judges the full conversation
                                (options: response, tool_call, conversation) [required]
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

* [calibrate tests](calibrate_tests.md)	 - Operations for tests
