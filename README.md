# Slack notification

This is a simple Slack notification action which runs using a Bot token.

## Example Action

A simple example on how to use this action:

```yaml
- name: slack-notification
  uses: skarlso/slack-notification-action
  with:
    token: ${{ secrets.SLACK_BOT_TOKEN }} # xoxb-...
    message: Build successful.
    channel: CD123456
```

Special characters to format the message are also allowed:

```yaml
- name: slack-notification
  uses: skarlso/slack-notification-action
  with:
    token: ${{ secrets.SLACK_BOT_TOKEN }} # xoxb-...
    message: "This is an awesome :wink: formatted message with `extra` characters."
    channel: CD123456
```

## Update a previous message

It's possible to update a previous message with a new status.
In order to do that, use the output parameter `timestamp`.

```yaml
name: Release

on:
  push:
    tags: [ 'v*' ]

jobs:
  goreleaser:
    runs-on: ubuntu-latest
    steps:
      - name: slack-notification
        id: slackstart
        uses: skarlso/slack-notification-action
        with:
          token: ${{ secrets.SLACK_BOT_TOKEN }} # xoxb-...
          message: "Running the release process... Current status is: :running:"
          channel: CD123456
      - name: Checkout
        uses: actions/checkout@v2
      - name: Unshallow
        run: git fetch --prune --unshallow
      - name: Setup Go
        uses: actions/setup-go@v2
        with:
          go-version: 1.16.x
      - name: Prepare
        id: prep
        run: |
          VERSION=sha-${GITHUB_SHA::8}
          if [[ $GITHUB_REF == refs/tags/* ]]; then
            VERSION=${GITHUB_REF/refs\/tags\//}
          fi
          echo ::set-output name=BUILD_DATE::$(date -u +'%Y-%m-%dT%H:%M:%SZ')
          echo ::set-output name=VERSION::${VERSION}
      - name: Run GoReleaser
        uses: goreleaser/goreleaser-action@v1
        with:
          version: latest
          args: release --release-notes=docs/release_notes/${{ steps.prep.outputs.VERSION }}.md --skip-validate
        env:
          GITHUB_TOKEN: ${{ secrets.GITHUB_TOKEN }}
      - name: slack-notification
        uses: skarlso/slack-notification-action
        with:
          token: ${{ secrets.SLACK_BOT_TOKEN }} # xoxb-...
          message: "Running the release process... Current status is: :done:"
          channel: CD123456
          timestamp: ${{ steps.slackstart.outputs.timestamp }}
```