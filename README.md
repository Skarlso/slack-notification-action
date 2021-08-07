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