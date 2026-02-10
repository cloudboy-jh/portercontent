---
title: "Architecture"
description: "How Porter converts issue comments into isolated, auditable worker execution."
slug: "docs/concepts/architecture"
category: "concepts"
order: 1
updated: 2026-02-09
sidebar:
  order: 1
---

Porter is a GitHub-native orchestration layer for coding agents. It keeps user interaction in issues while moving compute into short-lived cloud workers.

## Design goals

- Keep GitHub as the control surface for developers.
- Isolate each execution in an ephemeral runtime.
- Preserve an auditable record through issue comments and pull requests.
- Avoid hard-coding credentials in source repositories.

## System design

```text
GitHub (Issues, PRs, Webhooks)
           |
           v
    Porter (SvelteKit)
           |
           v
    Fly Machines API
           |
           v
    Docker Container
    (Agent CLI runs, creates PR)
           |
           v
    Callback to Porter
```

## Component responsibilities

### GitHub

- Emits webhook events for issue comments.
- Stores issue and pull request timeline history.
- Hosts the repository that the worker clones and modifies.

### Porter web app

- Verifies webhook authenticity.
- Parses `@porter <agent>` commands.
- Loads user-owned config and credentials from GitHub Gist.
- Builds the execution prompt and provisions workers.
- Receives completion callbacks and posts final issue updates.

### Fly Machines workers

- Start from an immutable container image.
- Clone repository and create a task branch.
- Run selected CLI agent in non-interactive mode.
- Exit and self-destroy after completion.

## Why this architecture works

Porter standardizes one flow across agents and repos:

- Trigger from GitHub comments
- Enrich with issue and repo context
- Execute in short-lived cloud containers
- Report back to the issue with result and PR

This separation keeps the web tier lightweight and the execution tier disposable.

## Security boundaries

- User credentials remain in user-owned private Gists.
- Secrets are injected at runtime as environment variables per task.
- Worker lifetime is bounded to the task, reducing long-lived exposure.
- GitHub App permissions are scoped to only required repository actions.

## Tech stack

| Layer | Technology |
| --- | --- |
| Web app | SvelteKit + Bun |
| Hosting | Vercel |
| Auth | GitHub OAuth |
| Execution | Fly Machines |
| Container | Docker (Node 20 + agent CLIs) |
| Config storage | GitHub Gist (user-owned) |

## Failure model

- If webhook validation fails, no task is created.
- If runtime provisioning fails, task status moves to failed with reason.
- If agent execution fails, Porter posts failure details back to the originating issue.
- All failures are visible in GitHub where the request started.
