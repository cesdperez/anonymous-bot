# anonymous-bot

Slack bot for playing around with golang. It sends "anonymous" messages to a user or channel.

## Running with go

1. `go get github.com/cesdperez/anonymous-bot`
2. `go build github.com/cesdperez/anonymous-bot`
3. `SLACK_TOKEN=<your-token> go run github.com/cesdperez/anonymous-bot` 

## Running with Docker

2. `docker run -e "SLACK_TOKEN=<your-token>" -d infortino/anonymous-bot`
