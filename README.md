# Porter Content Engine

This repo contains the base scaffolding for the Porter content engine.

## Structure

- `porter-launch/` - Astro monorepo for landing + docs
- `social-gen/` - Go CLI for social content generation
- `doc-md/` - Product docs and specs for reference

## Quick start

```bash
cd porter-launch
bun install
bun run dev
```

```bash
cd social-gen
go run main.go generate
```
