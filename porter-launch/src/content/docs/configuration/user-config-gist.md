---
title: "User config and secrets"
description: "D1-backed user settings and encrypted provider secrets used during task execution."
slug: "docs/configuration/user-config-gist"
category: "configuration"
order: 1
updated: 2026-02-19
sidebar:
  order: 1
---

Porter stores per-user settings and credentials in Cloudflare D1.

## Required shape

```json
{
  "version": "1.0.0",
  "execution_mode": "cloud",
  "fly_app_name": "porter-prod",
  "fly_token": "...",
  "anthropic_api_key": "...",
  "amp_api_key": "...",
  "default_agent": "opencode",
  "settings": {
    "max_retries": 3,
    "task_timeout": 90,
    "poll_interval": 10
  }
}
```

## Field expectations

- `fly_token`: token with permission to create and manage machines.
- `fly_app_name`: target Fly app where Porter starts machines.
- `anthropic_api_key`: required when running Anthropic-backed agent paths.
- `amp_api_key`: required for Amp execution.
- `default_agent`: fallback when command does not specify an agent.

## Data model

```ts
interface UserConfig {
  version: string;
  executionMode: "cloud" | "priority";
  flyToken: string;
  flyAppName?: string;
  agents: Record<string, { enabled: boolean; priority?: "low" | "normal" | "high" }>;
  providerCredentials?: Record<string, Record<string, string>>;
  settings: { maxRetries: number; taskTimeout: number; pollInterval: number };
}
```

## Runtime usage

- Porter loads this configuration after webhook validation.
- Porter injects required credentials into machine environment variables.
- Credentials are encrypted at rest and are not hardcoded into repositories.

## Security recommendations

- Use Porter settings while authenticated as the same GitHub user that will invoke commands.
- Rotate credentials regularly and immediately after suspected exposure.
- Do not commit raw key values to repository files.

## Validation checklist

- JSON is valid and fields are correctly named.
- Keys have no leading/trailing whitespace.
- `default_agent` matches a supported command name.

## Failure symptoms and fixes

### Credentials not configured

Ensure you are signed in as the expected GitHub identity and that required keys are saved.

### Auth errors in worker

Re-check token values and provider account status; invalid keys typically fail before agent output.
