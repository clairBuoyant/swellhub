# Development Scripts

These scripts are provided for development of [clairBuoyant](https://www.github.com/clairBuoyant). The script names are standardized across all repositories for clairBuoyant to simplify the development experience.

### Available Commands

- `bootstrap`: resolve all system dependencies the application needs to run.
- `clean`: remove all unnecessary build artifacts.
- `init`: run bootstrap and setup.
- `setup`: install node dependencies and githooks.
- `start`: start app locally.
- `uninstall`: remove node dependencies and build artifacts.

### Usage

These scripts can be used directly or with `npm` (**recommended**).

1. Run with npm: `npm run <command_name>` (e.g., `npm run init`).
2. Run script directly:
  - `./scripts/<command_name>` (e.g., `./scripts/init`).
  - Run `. ./aliases` in your terminal in order to run any script just by `<command_name>` (e.g., `init` or `start`). <sup>1</sup>

#### Note

1. This script needs to be re-run every time you start a new terminal session. But, it saves you from prepending `npm run` every time! :)

### Attribution

Styled after GitHub's ["Scripts to Rule Them All"](https://github.com/github/scripts-to-rule-them-all).
