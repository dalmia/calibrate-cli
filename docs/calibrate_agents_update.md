## calibrate agents update

Update agent

### Synopsis

Update an agent's configuration

```
calibrate agents update [flags]
```

### Examples

```
  calibrate agents update --agent-uuid f47ac10b-58cc-4372-a567-0e02b2c3d479
```

### Options

```
  -a, --agent-uuid string   The agent to update. [required]
      --body string         Request body as JSON (alternative to individual flags). Can also be provided via stdin.
  -c, --config-param type   Agent behavioral config. The keys depend on type.
                            
                            **`type=agent`** (built inside Calibrate):
                            - `system_prompt` (string): the agent's instructions
                            - `llm.model` (string): `provider/model`, e.g. `openai/gpt-4.1` or `google/gemini-2.5-flash`
                            - `stt.provider` (string): `deepgram`, `openai`, `cartesia`, `elevenlabs`, `google`, `sarvam`, or `smallest`
                            - `tts.provider` (string): `cartesia`, `openai`, `google`, `elevenlabs`, `sarvam`, or `smallest`
                            - `settings.agent_speaks_first` (bool), `settings.max_assistant_turns` (int)
                            - `system_tools.end_call` (bool, optional): let the agent end the call
                            - `data_extraction_fields` (array, optional): `[{name, type, description, required}]`
                            
                            ```json
                            {
                              "system_prompt": "You are a helpful support agent.",
                              "llm": {"model": "openai/gpt-4.1"},
                              "stt": {"provider": "deepgram"},
                              "tts": {"provider": "elevenlabs"},
                              "settings": {"agent_speaks_first": true, "max_assistant_turns": 50}
                            }
                            ```
                            
                            **`type=connection`** (your own HTTP endpoint):
                            - `agent_url` (string, required): public HTTPS endpoint the agent is called at
                            - `agent_headers` (object, optional): headers sent on each request, e.g. auth
                            - `benchmark_provider` (string, optional): `openrouter` (default), `openai`, `google`, `anthropic`, `meta-llama`, `mistralai`, `deepseek`, `x-ai`, `cohere`, `qwen`, or `ai21`
                            
                            ```json
                            {
                              "agent_url": "https://api.example.com/agent",
                              "agent_headers": {"Authorization": "Bearer <token>"},
                              "benchmark_provider": "openrouter"
                            }
                            ```
                            
                            Replaces the stored config. Omit to leave unchanged.
                            
                            For `type=connection`, changing `agent_url` or `agent_headers` resets the connection and benchmark verification flags.
  -h, --help                help for update
  -n, --name string         New agent name. Omit to leave the name unchanged
      --x-api-key string    string value
      --x-org-uuid string   string value
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

* [calibrate agents](calibrate_agents.md)	 - Operations for agents
