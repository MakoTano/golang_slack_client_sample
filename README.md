# golang_slack_client_sample
A sample for command line slack post client.

# usage
- Setup your slack incoming webhook. see https://api.slack.com/incoming-webhooks
- Set `config/config.toml` (you should set `url` and `channel`)
- execute `go run slack_client.go post {some message}` from command line.