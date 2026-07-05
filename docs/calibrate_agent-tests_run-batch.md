## calibrate agent-tests run-batch

Run Tests Batch

### Synopsis

Run every linked test for a set of agents, one ``llm-unit-test`` job per agent.

Scope is driven by the optional ``agent_names`` payload:

- **Provided (non-empty)** — run only those agents. Names are unique per org
  and **all are validated up front**: if any doesn't resolve to a
  (non-deleted) agent in the caller's org, the call 404s with the offending
  names and NO jobs are created.
- **Omitted / null / empty** — run every agent in the caller's org.

For each selected agent, its linked tests are launched as one job. Agents
with no linked tests or an unverified connection are reported under
``skipped`` instead of failing the batch. Subject to the normal per-org
concurrency queue, so over-limit jobs come back ``queued``.

Auth accepts a JWT (frontend) or an `sk_` API key (programmatic clients).
Returns one ``runs`` entry per launched agent with ``agent_name``,
``agent_uuid``, ``task_id``, and ``status``.

```
calibrate agent-tests run-batch [flags]
```

### Examples

```
  calibrate agent-tests run-batch
```

### Options

```
      --body string         Request body as JSON (alternative to individual flags). Can also be provided via stdin.
  -b, --body-param string   JSON object
  -h, --help                help for run-batch
      --x-api-key string    string value
      --x-org-uuid string   string value
```

### Options inherited from parent commands

```
      --agent-mode             Enable structured errors and default TOON output for AI coding agents. Automatically enabled when a known agent environment is detected (CLAUDE_CODE, CURSOR_AGENT, etc.). Use --agent-mode=false to disable.
      --api-key-auth string    Org-scoped API key. Create one under Settings → API keys.
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

* [calibrate agent-tests](calibrate_agent-tests.md)	 - Operations for agent-tests
