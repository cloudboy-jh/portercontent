---
title: "Prompt enrichment"
description: "Prompt contract Porter constructs from issue payloads and repository metadata."
slug: "docs/concepts/prompt-enrichment"
category: "concepts"
order: 4
updated: 2026-02-09
sidebar:
  order: 4
---

Porter transforms issue payloads into a consistent prompt format before agent execution.

## Prompt template

```markdown
## Task
{issue title}

## Description
{issue body}

## Repository Context
{from AGENTS.md if present}

## Instructions
Complete this GitHub issue by making the necessary code changes.
Create a branch, make commits, and open a pull request.
Reference issue #{issue_number} in the PR description.
```

## Inputs Porter adds

- Issue title and body
- Repository name and issue number
- Optional `AGENTS.md` guidance from repository root
- Agent selection from `@porter <agent>` command

This gives each supported agent a stable execution contract while preserving repo-specific guidance.

## Enrichment goals

- Reduce ambiguity so agents start with complete context.
- Keep output conventions stable across different agent CLIs.
- Preserve repository-specific standards without duplicating prompts manually.

## Recommended prompt quality checks

- Include explicit acceptance criteria from the issue body when present.
- Preserve file paths, error messages, and reproduction steps verbatim.
- Avoid adding speculative requirements that are not in issue context.

## Failure cases to guard against

### Missing issue body

Fallback to title plus minimal instruction block; do not skip execution automatically.

### Oversized context

Apply truncation strategy for large issue threads while keeping required details and links.

### Conflicting instructions

Prefer explicit issue content over generalized defaults and annotate conflict in final output.

## Example enriched prompt intent

Good enriched prompts tell the agent what to do, where to do it, and what done looks like. This is the main driver of reliable first-pass PR quality.
