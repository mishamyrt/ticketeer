# Ticketeer [![Quality assurance](https://github.com/mishamyrt/ticketeer/actions/workflows/quality-assurance.yaml/badge.svg)](https://github.com/mishamyrt/ticketeer/actions/workflows/quality-assurance.yaml) [![Coverage Status](https://coveralls.io/repos/github/mishamyrt/ticketeer/badge.svg?branch=main)](https://coveralls.io/github/mishamyrt/ticketeer?branch=main)

<img src="./docs/logo.svg" align="right" width="100" />

The Ticketeer is designed to streamline your development workflow by automatically appending the ticket number to your commit messages. This ensures that every commit is linked to a specific issue or task, making it easier to track changes and maintain a clear history.

Utility to insert task ticket id into commit message.

- **Simple**. Does not need to be configured;
- **Environment agnostic**. Works with any platform that can run git;
- **Fast**. Won't slow down your commit process as actions are performed instantly.

## How It Works

When you make a commit in a branch containing the ticket number (e.g., `feature/XXX-123`), ticketeer automatically detects the ticket number (`XXX-123`) and modifies the commit message.

If you commit with the message `commit text` and the branch name is `feature/XXX-123`, the resulting commit message will be:

```
commit text

XXX-123
```

By default, ticketeer ticket id is appended to the body of the commit message. In my humble opinion, it makes the story more definitive. But if you think otherwise and want the task number in the header prefix — use the configuration file:

```yaml
# .ticketeer.yaml
message:
  location: title
```

With this configuration, the resulting commit message will be:

```
XXX-123: commit text
```

## Installation

### go

```bash
go install github.com/mishamyrt/ticketeer@latest
```

### pnpm

```bash
pnpm add --save-dev ticketeer
```

### npm

```bash
npm install --save-dev ticketeer
```

### yarn

```bash
yarn add --dev ticketeer
```

### pip

```bash
pip install ticketeer
```

## Usage

### Standalone

Ticketeer can be used as a standalone tool. To install hook, run:

```bash
ticketeer install
```

The utility will check for installed hooks and install its own.

### Runner

If you are already using runner hooks on your project and want to keep everything under their control, add a `ticketeer apply` call to your runner configuration.

#### [Lefthook](https://github.com/evilmartians/lefthook)

Add the following to your `lefthook.yml` file:

```yaml
prepare-commit-msg:
  commands:
    append-ticket-id:
      run: ticketeer apply
```

If ticketeer is installed using NodeJS package manager, prefix command with your package manager runner, e.g. `pnpm ticketeer apply`, `yarn ticketeer apply` or `npx ticketeer apply`.

## Configuration

Ticketeer is configured by default to be as comfortable as possible for most, but the settings can be overridden. All options are optional (hehe), if you don't specify your own - default value will be used.

- `message`
  - `location` - where to insert ticket id into commit message (`title` or `body`).
  - `template` - template to use when inserting ticket id into commit message. Must include `{ticket}` variable.
- `ticket`
  - `format` - format of ticket id (`numeric`, `alphanumeric`, `alphanumeric-small`, `alphanumeric-caps`).
  - `allow-empty` - allow empty ticket id. If set to `false`, commit will be aborted if branch name doesn't contain ticket id.
- `branch`
  - `format` - format of branch name (`git-flow`, `git-flow-typeless`, `ticket-id`).
  - `ignore` - Additional list of branch names to ignore. Default value is `["main", "master", "develop", "dev", "release/*"]`.

Default configuration is as follows:

```yaml
message:
  location: body
  template: "{ticket}"
ticket:
  format: alphanumeric-caps
  allow-empty: true
branch:
  format: git-flow
```
