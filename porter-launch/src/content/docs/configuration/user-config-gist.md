---
title: "User config gist"
description: "Secure per-user runtime configuration pattern using private GitHub Gists."
slug: "docs/configuration/user-config-gist"
category: "configuration"
order: 1
updated: 2026-02-09
sidebar:
  order: 1
---

Porter reads per-user configuration from a private GitHub Gist.

## Required shape

```json
{
  "fly_token": "...",
  "anthropic_api_key": "...",
  "amp_api_key": "...",
  "default_agent": "opencode"
}
```

## Field expectations

- `fly_token`: token with permission to create and manage machines.
- `anthropic_api_key`: required when running Anthropic-backed agent paths.
- `amp_api_key`: required for Amp execution.
- `default_agent`: fallback when command does not specify an agent.

## Data model

```ts
interface UserConfig {
  flyToken: string;
  anthropicKey: string;
  ampKey?: string;
  openaiKey?: string;
  defaultAgent: string;
}
```

## Runtime usage

- Porter loads this config after webhook validation.
- Porter injects required credentials into machine environment variables.
- Credentials remain user-owned and are not hardcoded into repos.

## Security recommendations

- Keep Gist private and owned by the same GitHub user invoking Porter.
- Rotate credentials regularly and immediately after suspected exposure.
- Do not commit raw key values to repository files.

## Validation checklist

- JSON is valid and fields are correctly named.
- Keys have no leading/trailing whitespace.
- `default_agent` matches a supported command name.

## Failure symptoms and fixes

### Gist not found

Ensure Porter account is linked to the same GitHub identity that owns the Gist.

### Auth errors in worker

Re-check token values and provider account status; invalid keys typically fail before agent output.
