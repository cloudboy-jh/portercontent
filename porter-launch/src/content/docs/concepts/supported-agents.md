---
title: "Supported Agents"
description: "Agent CLI compatibility, invocation patterns, and selection guidance."
slug: "docs/concepts/supported-agents"
category: "concepts"
order: 7
updated: 2026-02-09
sidebar:
  order: 7
---

Porter supports any CLI agent that can run non-interactively in a container.

## Current support

| Agent | Package | Headless command |
| --- | --- | --- |
| Opencode | `opencode-ai` | `opencode run --model anthropic/claude-sonnet-4 "prompt"` |
| Claude Code | `@anthropic-ai/claude-code` | `claude -p "prompt" --dangerously-skip-permissions` |
| Amp | `@sourcegraph/amp` | `amp -x "prompt" --dangerously-allow-all` |

## Command syntax

Use one of the supported agent names in issue comments:

```text
@porter opencode
@porter claude
@porter amp
```

## How to pick an agent

- `opencode`: balanced default for general issue implementation.
- `claude`: strong for broad refactors and architecture-sensitive edits.
- `amp`: useful for test and codebase-wide assistance workflows.

## Compatibility requirements

- CLI must support non-interactive invocation.
- CLI must run cleanly in Linux container runtime.
- CLI should return machine-readable or clearly parseable output for callbacks.

## Operational tips

- Keep an explicit default agent in user config to avoid ambiguous behavior.
- Validate agent-specific credentials before production rollout.
- Track completion quality by agent type to tune defaults over time.
