---
title: "Endpoints"
description: "HTTP surface for ingestion, task lifecycle, callbacks, and user configuration."
slug: "docs/api/endpoints"
category: "api"
order: 1
updated: 2026-02-09
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

## Tasks

```http
GET    /api/tasks
POST   /api/tasks
GET    /api/tasks/:id
DELETE /api/tasks/:id
```

- `GET /api/tasks`: list tasks
- `POST /api/tasks`: create task (internal)
- `GET /api/tasks/:id`: fetch task status
- `DELETE /api/tasks/:id`: cancel task

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
```

- `GET /api/config`: load user config from Gist
- `PUT /api/config`: update user config in Gist

## Error handling guidance

- Use `400` for malformed request payloads.
- Use `401`/`403` for auth and permission failures.
- Use `404` when task or config target is missing.
- Use `5xx` only for genuine server/runtime failures.
