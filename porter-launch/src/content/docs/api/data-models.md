---
title: "Data models"
description: "Canonical task and configuration shapes used throughout Porter runtime."
slug: "docs/api/data-models"
category: "api"
order: 3
updated: 2026-02-09
sidebar:
  order: 3
---

## Task

```ts
interface Task {
  id: string;
  status: "queued" | "running" | "complete" | "failed";
  repo: string;
  issueNumber: number;
  agent: string;
  prompt: string;
  machineId?: string;
  prUrl?: string;
  createdAt: Date;
  completedAt?: Date;
  error?: string;
}
```

Field intent:

- `id`: globally unique task identifier.
- `status`: current lifecycle phase.
- `machineId`: Fly runtime handle for operational introspection.
- `prUrl`: set when task results in a pull request.
- `error`: terminal failure details for UI and debugging.

## User config

```ts
interface UserConfig {
  flyToken: string;
  anthropicKey: string;
  ampKey?: string;
  openaiKey?: string;
  defaultAgent: string;
}
```

## Notes

- `status` tracks the end-to-end execution lifecycle.
- `machineId` links a task to Fly runtime state.
- `prUrl` is set when agent execution results in a pull request.

## Modeling recommendations

- Keep status enum small and explicit.
- Use nullable optional fields for runtime-dependent metadata.
- Store timestamps in UTC and render locale at presentation time.

## Extension ideas

- Add `durationMs` for performance analytics.
- Add `attempt` for retries and replay workflows.
- Add `actor` metadata to correlate request ownership and audit trails.
