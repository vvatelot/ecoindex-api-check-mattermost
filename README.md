# Go Ecoindex Health checker

The aim of this project is to call the Health endpoint of an [Ecoindex API](https://github.com/cnumr/ecoindex_api), and send a Mattermost notification when there is an error on a dedicated incoming webhook.

You can configure this executable file to run as a scheduled task. The first time the API encounters an error, it will send a notification on Mattermost. When the server comes back to normal, it will send a notification of end of failure.

## Prerequisite

It works on linux/386, linux/amd64, linux/arm64, windows/386, windows/amd64, darwin/amd64, darwin/arm64. You can download your platform package on the [release]("./release") page.

You obviously need a running Ecoindex API that you want to test and a Mattermost instance with a [incoming webhook](https://developers.mattermost.com/integrate/webhooks/incoming/) configured.

## Execute

```text
NAME:
   Ecoindex Bot - Check Ecoindex API Status and alert in case of emergency

USAGE:
   Ecoindex Bot [global options] command [command options] [arguments...]

COMMANDS:
   help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --ecoindex_url value    base url of the ecoindex API to check (default: "http://localhost:8001")
   --mattermost_url value  mattermost webhook url used to send message
   --name value            name of the ecoindex API. It can bes something like 'production', 'backup' (default: "Local")
   --help, -h              show help (default: false)
```

## Basic example

Run it on a localhost API instance:

```bash
./ecoindex_bot --name "My Local Ecoindex Instance" --mattermost_url "https://my.mattermost.com/hooks/myincomingwebhookid"
```

Run it on a remote API instance:

```bash
./ecoindex_bot --name "My Remote Ecoindex Instance" --mattermost_url "https://my.mattermost.com/hooks/myincomingwebhookid" --ecoindex_url "https://my.ecoindex.api"
```
