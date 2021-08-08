FROM golang:1.16-alpine as build
WORKDIR /app
COPY . .
RUN go build -o /slack-notification

FROM alpine
RUN apk add -u ca-certificates
COPY --from=build /slack-notification /app/

LABEL "name"="Slack Notifier"
LABEL "maintainer"="Gergely Brautigam <gergely@gergelybrautigam.com>"
LABEL "version"="0.0.2"

LABEL "com.github.actions.name"="Slack Notifier"
LABEL "com.github.actions.description"="Send messages on Slack using a Slack App and a Bot token."
LABEL "com.github.actions.icon"="package"
LABEL "com.github.actions.color"="purple"

WORKDIR /app/
ENTRYPOINT [ "/app/slack-notification" ]
