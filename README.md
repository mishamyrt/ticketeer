# Ticketeer [![Quality assurance](https://github.com/mishamyrt/ticketeer/actions/workflows/quality-assurance.yaml/badge.svg)](https://github.com/mishamyrt/ticketeer/actions/workflows/quality-assurance.yaml)

<img src="./docs/logo.svg" align="right" width="100" />

Utility to insert task ticket id into commit message.

- **Simple**. Does not need to be configured;
- **Environment agnostic**. Works with any platform that can run git;
- **Fast**. Won't slow down your commit process as actions are performed instantly.

## Installation

### Via `go`

```bash
go install github.com/mishamyrt/ticketeer@latest
```

## Usage

Add a `ticketeer apply` command call to the `prepare-commit-msg` hook in your hook runner.
For example, for [lefthook](https://github.com/evilmartians/lefthook/tree/master) it would look like this:

```yaml
prepare-commit-msg:
  commands:
    append-ticket-id:
      run: ticketeer apply
```

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
