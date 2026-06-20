# Triage Labels

The skills speak in terms of five canonical triage roles. This file maps those roles to
the actual label strings used in this repo's issue tracker.

| Label in mattpocock/skills | Label in our tracker | Meaning                                  |
| -------------------------- | -------------------- | ---------------------------------------- |
| `needs-triage`             | `needs-triage`       | Maintainer needs to evaluate this issue  |
| `needs-info`               | `needs-info`         | Waiting on reporter for more information |
| `ready-for-agent`          | `ready-for-agent`    | Fully specified, ready for an AFK agent  |
| `ready-for-human`          | `ready-for-human`    | Requires human implementation            |
| `wontfix`                  | `wontfix`            | Will not be actioned                     |

When a skill mentions a role (e.g. "apply the AFK-ready triage label"), use the
corresponding label string from this table.

## These labels don't exist in the repo yet

swellhub currently uses descriptive labels (`bug`, `feature`, `enhancement`, `tech debt`,
`in-progress`, `question`, …) — none of which are part of this triage state machine. The
five labels above are created on first use. To pre-create them:

```bash
gh label create needs-triage    --description "Maintainer needs to evaluate"   --color 0e8a16
gh label create needs-info      --description "Waiting on reporter for info"   --color fbca04
gh label create ready-for-agent --description "Fully specified, AFK-ready"      --color 1d76db
gh label create ready-for-human --description "Requires human implementation"  --color 5319e7
gh label create wontfix         --description "Will not be actioned"           --color cfd3d7
```

Edit the right-hand column above if you later adopt different label strings.
