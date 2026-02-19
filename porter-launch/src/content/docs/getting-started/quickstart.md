---
title: "Quick start"
description: "From first sign-in to validated issue-to-PR execution in one guided path."
slug: "docs/getting-started/quickstart"
category: "getting-started"
order: 1
updated: 2026-02-19
sidebar:
  order: 1
---

Use this page when you want a reliable first run, not just a demo. You will verify every stage so you know the full workflow is wired correctly.

## Before you start

- A GitHub repository where you can create issues and open pull requests.
- A GitHub account that can install apps for that repository.
- API credentials for your model provider(s) and Fly.
- Access to the Porter web app.

## Step 1: Sign in with GitHub

Authenticate in the Porter web app using the same GitHub identity that owns or contributes to your target repository.

Why this matters:

- Porter links webhook events, user config, and execution ownership through your GitHub identity.
- If the wrong account signs in, Porter may not find your repository or saved configuration.

## Step 2: Install the Porter GitHub App

Install the Porter GitHub App and grant repository access for repos where Porter should respond to issue comments.

Recommended scope:

- Start with one test repository.
- Expand permissions to more repos after your first successful run.

## Step 3: Configure credentials in Porter settings

Open Porter settings and save your execution credentials.

```json
{
  "fly_token": "...",
  "anthropic_api_key": "...",
  "amp_api_key": "...",
  "default_agent": "opencode"
}
```

Notes:

- Credentials are encrypted and stored server-side in Porter (D1-backed).
- Optional gist mirroring can exist, but task execution does not depend on it.
- Use valid keys only; malformed values are a common cause of failed first runs.
- `default_agent` is used when a command omits an explicit agent name.

## Step 4: Run your first task

Create a GitHub issue with a simple request, then add one command comment:

- `@porter opencode`
- `@porter claude`
- `@porter amp`

Porter receives the webhook, provisions a Fly Machine, runs the selected agent in a container, and posts the result back to the issue.

## What success looks like

You should observe this sequence:

1. Porter acknowledges the command in the issue timeline.
2. Task transitions to running state.
3. A branch is created with `porter/<task-id>` style naming.
4. A pull request is opened (or a detailed failure is reported).
5. Final issue comment includes status and PR link.

## Common first-run failures

### App installed on wrong repo

Symptom: no task starts after `@porter <agent>`.

Fix: confirm the GitHub App has access to the exact repository where the comment was posted.

### Invalid or missing credentials

Symptom: task starts then fails before agent execution.

Fix: re-check `fly_token` and provider API keys in Porter settings, then run validation again.

### Unsupported command format

Symptom: webhook received, but no execution.

Fix: use one of the documented command formats exactly (for example, `@porter opencode`).
