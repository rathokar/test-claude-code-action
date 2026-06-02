# test-claude-code-action

## Claude Server Status Check

This repository contains a Go program that checks whether the Claude server at https://status.claude.com/ is reachable.

### `check_claude_status.go`

Hits the Claude status page with a 10-second timeout and prints the result:

- **Exit 0** — server responded with HTTP 2xx/3xx (`[UP]`)
- **Exit 1** — connection failed or HTTP 4xx/5xx (`[DOWN]`)

#### Run locally

```bash
go run check_claude_status.go
```

### GitHub Actions: Status Check Workflow

A scheduled workflow (`.github/workflows/check_claude_status.yml`) runs this check automatically:

- **Schedule:** every 15 minutes
- **Manual trigger:** `workflow_dispatch`
- **On failure:** GitHub notifies repository watchers; the workflow exits non-zero

To add the workflow, create `.github/workflows/check_claude_status.yml` with the content shown in [issue #2](https://github.com/rathokar/test-claude-code-action/issues/2).
