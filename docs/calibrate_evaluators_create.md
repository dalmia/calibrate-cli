## calibrate evaluators create

Create evaluator

### Synopsis

Create an evaluator along with its first version, which is set live

```
calibrate evaluators create [flags]
```

### Examples

```
  calibrate evaluators create --name <value> --version-param '{"judge_model":"<value>","system_prompt":"<value>"}'
```

### Options

```
      --body string            Request body as JSON (alternative to individual flags). Can also be provided via stdin.
      --data-type text         The modality the judge reads:
                               
                               - text
                               - `audio`
                                (options: text, audio) (default "text")
      --description string     Description. Omit to leave blank
  -e, --evaluator-type tts     What the evaluator judges:
                               
                               - tts: TTS audio
                               - `stt`: one transcript
                               - `llm`: a reply with its conversation history
                               - `llm-general`: a standalone input and output pair
                               - `conversation`: a full conversation
                                (options: tts, stt, llm, llm-general, conversation) (default "llm")
  -h, --help                   help for create
  -n, --name string            Evaluator name, unique within your workspace [required]
      --output-type binary     How the evaluator scores:
                               
                               - binary: pass or fail
                               - `rating`: a numeric score, using the scale in `output_config`
                                (options: binary, rating) (default "binary")
  -v, --version-param string   One version of an evaluator: its judge prompt, model, variables, and rubric. [required]
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

* [calibrate evaluators](calibrate_evaluators.md)	 - Operations for evaluators
