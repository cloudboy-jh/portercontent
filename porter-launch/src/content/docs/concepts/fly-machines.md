---
title: "Fly Machines execution"
description: "Provisioning model, runtime configuration, and lifecycle of ephemeral workers."
slug: "docs/concepts/fly-machines"
category: "concepts"
order: 3
updated: 2026-02-09
sidebar:
  order: 3
---

## Machine provisioning

Porter creates one Fly Machine per task using the Machines API.

```http
POST https://api.machines.dev/v1/apps/{app}/machines
Authorization: Bearer {FLY_API_TOKEN}
Content-Type: application/json
```

Porter treats machine creation as a deterministic step: one task maps to one worker instance.

## Machine config

```json
{
  "config": {
    "image": "registry.fly.io/porter-worker:latest",
    "auto_destroy": true,
    "env": {
      "TASK_ID": "task_abc123",
      "REPO_FULL_NAME": "user/repo",
      "AGENT": "opencode",
      "PROMPT": "Fix the bug described in issue #42...",
      "GITHUB_TOKEN": "ghp_xxx",
      "ANTHROPIC_API_KEY": "sk-ant-xxx",
      "AMP_API_KEY": "amp_xxx",
      "CALLBACK_URL": "https://porter.dev/api/callbacks/complete"
    },
    "guest": {
      "cpu_kind": "shared",
      "cpus": 2,
      "memory_mb": 2048
    }
  }
}
```

## Runtime behavior

- Workers are immutable container starts from one shared image.
- `auto_destroy` removes machine state after completion.
- Environment variables are injected per task from validated runtime input.

## Why Fly Machines

- Fast cold starts for short-lived automation tasks.
- Per-machine resource sizing (CPU and memory) per workload profile.
- Clean lifecycle control aligned to event-driven jobs.

## Operational safeguards

- Use per-user Fly tokens so billing and ownership boundaries stay clear.
- Keep machine timeouts strict to prevent hung agent sessions.
- Prefer immutable worker images and version tags for reproducibility.

## Troubleshooting

### Provisioning fails before runtime

Check Fly API token validity, app name, and region/resource limits.

### Worker starts but cannot complete

Inspect callback payload and worker logs for clone/auth/provider failures.

### Intermittent performance variance

Review CPU/memory sizing and concurrent task volume for burst scenarios.
