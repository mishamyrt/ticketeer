# Ticketeer [![Quality assurance](https://github.com/mishamyrt/ticketeer/actions/workflows/quality-assurance.yaml/badge.svg)](https://github.com/mishamyrt/ticketeer/actions/workflows/quality-assurance.yaml)

Utility to insert task ticket id into commit message.

## Usage

Add a `ticketeer apply` command call to the `prepare-commit-msg` hook in your hook runner.
For example, for [lefthook](https://github.com/evilmartians/lefthook/tree/master) it would look like this:

```yaml
prepare-commit-msg:
  commands:
    append-ticket-id:
      run: ticketeer apply
```