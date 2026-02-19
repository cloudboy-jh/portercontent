---
title: "Endpoints"
description: "HTTP surface for ingestion, task lifecycle, callbacks, and user configuration."
slug: "docs/api/endpoints"
category: "api"
order: 1
updated: 2026-02-19
sidebar:
  order: 1
---

This page describes the functional endpoint groups and what each route is responsible for.

## Webhooks

```http
POST /api/webhooks/github
```

Receives GitHub issue comment events and dispatches task creation.

Expected behavior:

- Valid signature + valid command -> create task.
- Invalid signature or unsupported payload -> reject request.

## Auth and session

```http
GET  /api/auth/github
GET  /api/auth/github/callback
GET  /api/auth/github/installed
GET  /api/auth/diagnostics
POST /api/auth/logout
```

These routes handle GitHub OAuth sign-in, installation handoff, and auth diagnostics.

## Tasks

```http
GET    /api/tasks
POST   /api/tasks
GET    /api/tasks/history
GET    /api/tasks/:id/status
POST   /api/tasks/:id/retry
POST   /api/tasks/:id/stop
```

- `GET /api/tasks`: list tasks
- `POST /api/tasks`: create task (internal)
- `GET /api/tasks/history`: task feed/history view
- `GET /api/tasks/:id/status`: fetch current task status
- `POST /api/tasks/:id/retry`: retry a failed task
- `POST /api/tasks/:id/stop`: request cancellation/stop

Recommended response model:

- Return consistent task identifiers and status enum values.
- Include timestamps for queued, started, and completed states.
- Include error payload on failure transitions.

## Callback

```http
POST /api/callbacks/complete
```

Receives task completion signals from worker containers.

Expected callback payload fields typically include task id, status, and optional metadata such as PR URL.

## Config

```http
GET /api/config
PUT /api/config
PUT /api/config/credentials
GET /api/config/provider-credentials
PUT /api/config/provider-credentials
PUT /api/config/fly
GET /api/config/providers
POST /api/config/validate/anthropic
POST /api/config/validate/fly
```

- `GET /api/config`: load user config from D1
- `PUT /api/config`: update user config in D1
- Provider credential routes: configure key status and secret writes
- Validation routes: verify provider and Fly credentials before runtime

## GitHub data

```http
GET /api/github/summary
GET /api/github/repositories
GET /api/github/installations
GET /api/github/orgs
GET /api/github/profile
GET /api/github/issues/:owner/:repo/:number
```

These routes drive authenticated UI data for installations, repositories, and issue context.

## Ops and health

```http
GET /api/startup/check
GET /api/rate-limit
```

- `GET /api/startup/check`: required env and D1 readiness checks
- `GET /api/rate-limit`: current GitHub API rate limit information

## Error handling guidance

- Use `400` for malformed request payloads.
- Use `401`/`403` for auth and permission failures.
- Use `404` when task or config target is missing.
- Use `5xx` only for genuine server/runtime failures.
