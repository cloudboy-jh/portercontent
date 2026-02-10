---
title: "GitHub App Permissions"
description: "Minimum repository permissions and event scope needed for Porter workflows."
slug: "docs/api/github-app-permissions"
category: "api"
order: 4
updated: 2026-02-09
sidebar:
  order: 4
---

Porter GitHub App requires the following permissions:

- Contents: read/write
- Issues: read/write
- Pull requests: read/write
- Metadata: read

These permissions let Porter parse issue comments, trigger execution, and post PR outcomes back into GitHub.

## Permission rationale

- `Contents` is needed for branch and commit operations in worker tasks.
- `Issues` is needed to parse commands and post status updates.
- `Pull requests` is needed to create and annotate result PRs.
- `Metadata` is needed for repository/app context resolution.

## Security posture

- Start with repository-level installation before org-wide rollout.
- Grant access only to repositories where automation is expected.
- Review app permission changes as part of release governance.

## Event scope

At minimum, enable issue comment events. Additional lifecycle events (for example issue close) can support cancellation and cleanup logic.
