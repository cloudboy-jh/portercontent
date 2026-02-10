---
title: "Resource Limits And Cost"
description: "Execution limits, concurrency boundaries, and practical cost expectations."
slug: "docs/concepts/resource-limits-and-cost"
category: "concepts"
order: 5
updated: 2026-02-09
sidebar:
  order: 5
---

## Resource limits

| Resource | Limit |
| --- | --- |
| Timeout | 10 minutes |
| Memory | 2 GB |
| CPU | 2 shared cores |
| Concurrent tasks per user | 5 |

Limits protect platform stability and keep execution behavior predictable.

## Cost model

Porter is free to use. Users pay underlying provider costs:

| Service | Cost |
| --- | --- |
| Fly Machines | ~$0.01-0.05 per task |
| Anthropic/OpenAI | Per-token pricing |
| Amp | Per Amp pricing |

## Design intent

- Keep workers short-lived and predictable.
- Let users choose their own model providers.
- Avoid hidden platform margin on execution.

## Cost planning guidance

- Start with smaller tasks to baseline average runtime.
- Track provider token usage by task type (bug fix, refactor, test generation).
- Set team-level guidelines for when to use each agent/model tier.

## Concurrency behavior

- User-level concurrency prevents one actor from saturating worker capacity.
- Queued tasks should execute in arrival order unless manual prioritization is introduced.

## Common overrun causes

- Issues with broad scope and no acceptance boundary.
- Repositories with heavy dependency install overhead.
- Long-running test suites executed as part of every task.

## Suggested controls

- Add issue templates that force clear task boundaries.
- Keep CI setup efficient to reduce worker idle time.
- Tune timeout and resource profiles as usage patterns become clear.
