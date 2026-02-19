---
title: "Data models"
description: "Canonical task and configuration shapes used throughout Porter runtime."
slug: "docs/api/data-models"
category: "api"
order: 3
updated: 2026-02-19
sidebar:
  order: 3
---

## Task

```ts
interface Task {
  id: string;
  status: "queued" | "running" | "success" | "failed" | "timed_out";
  repoOwner: string;
  repoName: string;
  issueNumber: number;
  issueTitle: string;
  issueBody: string;
  agent: string;
  priority: number;
  progress: number;
  createdBy: string;
  branch?: string;
  prUrl?: string;
  callbackAttempts?: number;
  callbackMaxAttempts?: number;
  callbackLastHttpCode?: number;
  createdAt: string;
  startedAt?: string;
  completedAt?: string;
  errorMessage?: string;
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
  version: string;
  executionMode: "cloud" | "priority";
  flyToken: string;
  flyAppName?: string;
  agents: Record<string, { enabled: boolean; priority?: "low" | "normal" | "high" }>;
  providerCredentials?: Record<string, Record<string, string>>;
  settings: {
    maxRetries: number;
    taskTimeout: number;
    pollInterval: number;
  };
}
```

## Notes

- `status` tracks the end-to-end execution lifecycle.
- `branch` and `prUrl` link issue execution to Git state changes.
- callback fields support retry/idempotency and callback diagnostics.

## Modeling recommendations

- Keep status enum small and explicit.
- Use nullable optional fields for runtime-dependent metadata.
- Store timestamps in UTC and render locale at presentation time.

## Extension ideas

- Add `durationMs` for performance analytics.
- Add `attempt` for retries and replay workflows.
- Add `actor` metadata to correlate request ownership and audit trails.
