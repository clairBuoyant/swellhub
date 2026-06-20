# Issue tracker: GitHub

Issues and PRDs for this repo live as GitHub issues in `clairBuoyant/swellhub`. Use the
`gh` CLI for all operations. Every issue is also tracked on the org-level roadmap board
(see "Org Project mapping" below).

## Conventions

- **Create an issue**: `gh issue create --title "..." --body "..."`. Use a heredoc for multi-line bodies. After creating, add it to the roadmap project (see below).
- **Read an issue**: `gh issue view <number> --comments`, filtering comments by `jq` and also fetching labels.
- **List issues**: `gh issue list --state open --json number,title,body,labels,comments --jq '[.[] | {number, title, body, labels: [.labels[].name], comments: [.comments[].body]}]'` with appropriate `--label` and `--state` filters.
- **Comment on an issue**: `gh issue comment <number> --body "..."`
- **Apply / remove labels**: `gh issue edit <number> --add-label "..."` / `--remove-label "..."`
- **Close**: `gh issue close <number> --comment "..."`

Infer the repo from `git remote -v` — `gh` does this automatically when run inside a clone.

## Org Project mapping

This repo's issues are tracked on a shared org-level GitHub Project (Projects v2):

- **Project**: #5 — "clairBuoyant Roadmap"
- **Owner**: `clairBuoyant` (organization)
- **URL**: https://github.com/orgs/clairBuoyant/projects/5
- **Project ID**: `PVT_kwDOBQKLk84BawLC`
- **Scope**: shared across clairBuoyant repos; the board's built-in **Repository** field distinguishes swellhub items from other repos'.

New issues are **not** auto-added to the board, so add them explicitly after creation:

```bash
gh project item-add 5 --owner clairBuoyant --url <issue-url>
```

(`gh issue create` prints the new issue's URL; pass it straight through.)

**Status field** (single-select `Status`): `Todo` → `In Progress` → `Done`. New items
have no status by default; set `Todo` when adding to the backlog and advance as work
proceeds. There's also a `Priority` single-select. Status/Priority can be set in the
project UI or via `gh project item-edit`, which needs item, field, and option IDs — fetch
them with:

```bash
gh project item-list  5 --owner clairBuoyant --format json
gh project field-list 5 --owner clairBuoyant --format json
```

Requires the `project` scope on your `gh` token (already present for this checkout).

## Pull requests as a triage surface

**PRs as a request surface: no.** This is a private roadmap repo; PRs are implementation
work (one PR per ticket), not incoming feature requests. `/triage` processes issues only
and leaves PRs alone.

## When a skill says "publish to the issue tracker"

Create a GitHub issue, then add it to Project #5 (see "Org Project mapping").

## When a skill says "fetch the relevant ticket"

Run `gh issue view <number> --comments`.
