---
title: "Core Flow"
description: "Lifecycle from `@porter` command to callback, issue update, and pull request."
slug: "docs/concepts/core-flow"
category: "concepts"
order: 2
updated: 2026-02-09
sidebar:
  order: 2
---

## End-to-end execution

1. A user comments `@porter <agent>` on a GitHub issue.
2. GitHub sends an `issue_comment.created` webhook to Porter.
3. Porter validates signature, actor, repository scope, and command syntax.
4. Porter loads user config from private GitHub Gist.
5. Porter enriches the prompt with issue, repository, and optional `AGENTS.md` context.
6. Porter creates a Fly Machine with the task payload.
7. Worker container clones the repository and checks out a task branch.
8. Selected agent runs in non-interactive mode.
9. Worker pushes changes and opens a pull request when changes are produced.
10. Worker calls Porter completion callback with status and metadata.
11. Porter finalizes task state and comments on the issue with result details.
12. Machine auto-destroys after process exit.

## Why this flow

- Keeps GitHub as the user-facing control surface.
- Runs compute in isolated, short-lived workers.
- Preserves an auditable trail in issues and PRs.

## State transitions

Typical status progression:

`queued -> running -> complete` or `queued -> running -> failed`

Status is persisted so API consumers and UI components can render reliable task progress.

## Retry behavior

- Webhook-level failures can be retried by re-posting the command comment.
- Worker runtime failures should include actionable error output in the issue comment.
- Duplicate commands create separate tasks, which can be useful for model comparison.

## Operational notes

- One issue can have multiple Porter tasks over time.
- Agent choice is explicit (`@porter opencode`) or defaults from user config.
- Cancellation flows can be bound to issue lifecycle events like close/reopen.
