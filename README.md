# Bot

Bot is a [alertmanager](https://github.com/prometheus/alertmanager) webhook.

## IRC Commands
`!clear` clears all current active alerts,
`!reset` resets failed state for the caller.

## Configuration

#### -cfg
A non-configured feature will be ignored, this file is optional.
```json
{
	"receivers": {
		"group1": ["user1", "user2", "user3"],
		"group2": ["user1", "user4"],
		"group3": ["user5", "user3"]
	}
}
```
#### -cfg.template
Alerts sent to endpoints can be customized with a template, see the default [template](https://github.com/multimfi/bot/blob/master/http/template.go) for an example, newlines are replaced with spaces before parsing.


## Receiver groups

When an alert is received a message is broadcasted via SSE and the current user is highlighted on IRC. If user does not respond in time user will be marked as failed and proceed to the next user.

Failed state can be reset with the `!reset` command.

## Help

``` text
 $ bot-daemon -h
 Usage of bot-daemon:
  -alertmanager.addr string
    	alertmanager webhook listen address (default "127.0.0.1:9500")
  -cfg string
    	bot configuration file (default "bot.json")
  -cfg.template string
        template file (default "template.tmpl")
  -irc.channel string
    	irc channel to join (default "#test")
  -irc.nick string
    	irc nickname (default "bot")
  -irc.server string
    	irc server address (default "127.0.0.1:6667")
  -irc.user string
    	irc username (default "Bot")
  -version
    	version
```
