---
title: "Webhooks and callbacks"
description: "Inbound GitHub trigger contracts and outbound worker completion signaling."
slug: "docs/api/webhooks-and-callbacks"
category: "api"
order: 2
updated: 2026-02-09
sidebar:
  order: 2
---

## GitHub webhook events

Porter listens for:

- `issue_comment.created` to trigger execution from `@porter <agent>` mentions.
- `issues.closed` to cancel related running tasks.

Webhook validation should include signature checks, repository scope checks, and command parser checks.

## Command parsing

Supported commands:

```text
@porter <agent>
@porter opencode
@porter claude
@porter amp
```

Parser rules should:

- Ignore comments without `@porter` prefix.
- Normalize spacing and handle minor formatting variation.
- Reject unknown agent names with a clear issue reply.

## Completion callback contract

Workers call Porter when execution completes:

```bash
curl -X POST "$CALLBACK_URL" \
  -H "Content-Type: application/json" \
  -d '{"task_id": "task_abc123", "status": "complete"}'
```

Porter uses this callback to finalize task state and comment back on the issue.

## Callback statuses

Typical statuses include:

- `complete`
- `failed`
- `cancelled`

Each status should map to a deterministic update in task state and issue comment messaging.

## Reliability practices

- Require callback authentication to prevent spoofed completion events.
- Treat callbacks as idempotent by task id and terminal status.
- Persist callback payload for audit and debugging workflows.
