# swellhub

Agent instructions for `clairBuoyant/swellhub`. Add project-specific guidance above
the section below; the `## Agent skills` block is managed by the
`setup-matt-pocock-skills` skill and points at the per-topic docs under `docs/agents/`.

## CI conventions

- **GitHub token format:** the Actions `GITHUB_TOKEN` is rolling out as a longer (~520-char), `ghs_`-prefixed JWT ([2026-05-15 changelog](https://github.blog/changelog/2026-05-15-github-app-installation-tokens-per-request-override-header/)). We pass it opaquely, so there's nothing to change today — but never assume a token's length or shape in new CI tooling, regexes, or token-storing DB columns (allow ≥520 chars and the dotted `ghs_…` form). No deprecation date announced yet.

## Agent skills

### Issue tracker

Issues live in GitHub Issues for `clairBuoyant/swellhub` (via the `gh` CLI) and are
tracked on org Project #5 "clairBuoyant Roadmap". External PRs are not a triage
surface. See `docs/agents/issue-tracker.md`.

### Triage labels

Canonical triage vocabulary (`needs-triage`, `needs-info`, `ready-for-agent`,
`ready-for-human`, `wontfix`); created on first use. See `docs/agents/triage-labels.md`.

### Domain docs

Single-context: one `CONTEXT.md` + `docs/adr/` at the repo root. See `docs/agents/domain.md`.
