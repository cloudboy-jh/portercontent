---
title: "Worker runtime"
description: "Container build, entrypoint flow, and execution guarantees for worker tasks."
slug: "docs/concepts/worker-runtime"
category: "concepts"
order: 6
updated: 2026-02-09
sidebar:
  order: 6
---

## Container image

```dockerfile
FROM node:20-slim

RUN apt-get update && apt-get install -y git curl && \
    npm install -g opencode-ai @anthropic-ai/claude-code @sourcegraph/amp && \
    apt-get clean && rm -rf /var/lib/apt/lists/*

WORKDIR /workspace

COPY entrypoint.sh /entrypoint.sh
RUN chmod +x /entrypoint.sh

ENTRYPOINT ["/entrypoint.sh"]
```

Porter builds this image once and reuses it for all tasks.

Runtime guarantees:

- Same base image across all tasks.
- Agent CLIs installed in build stage, not at request time.
- Predictable startup path and dependency availability.

## Entrypoint behavior

```bash
#!/bin/bash
set -e

git clone "https://${GITHUB_TOKEN}@github.com/${REPO_FULL_NAME}.git" .
git checkout -b "porter/${TASK_ID}"

case "$AGENT" in
  opencode)
    opencode run --model anthropic/claude-sonnet-4 "$PROMPT"
    ;;
  claude)
    claude -p "$PROMPT" --dangerously-skip-permissions
    ;;
  amp)
    amp -x "$PROMPT" --dangerously-allow-all
    ;;
esac

curl -X POST "$CALLBACK_URL" \
  -H "Content-Type: application/json" \
  -d "{\"task_id\": \"$TASK_ID\", \"status\": \"complete\"}"
```

## Execution phases

1. Clone repository with GitHub token.
2. Create isolated task branch.
3. Run selected agent with enriched prompt.
4. Capture result and notify callback endpoint.

## Security notes

- Worker token scope should be minimal for required Git operations.
- Treat agent runtime flags as high-trust execution; use only in controlled containers.
- Avoid writing secrets to logs or committed files.

## Hardening recommendations

- Pin image tags for deterministic rollouts.
- Add health and completion telemetry for each phase.
- Use structured callback payloads with error details when non-zero exit codes occur.
