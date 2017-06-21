# Bot

Bot is a [alertmanager](https://github.com/prometheus/alertmanager) webhook.

###### Supported endpoints
* IRC
* WebSocket
* Telegram

WebSocket and Telegram endpoints can only receive alerts.

###### Receiver groups

When alert is received a message is broadcasted to configured endpoints and mapped user of the current ISO week is highlighted on IRC. If user does not respond within 5min user will be marked as failed and proceed to next user.

Failed state can be reset with the `!reset` command.

###### Week to user mapping

ISO Week | User
-----|-----
W1 | user1
W2 | user2
W3 | user3
W4 | user1
W5 | user2
...|...

###### bot.json (optional)

```json
{
	"receivers": {
		"group1": ["user1", "user2", "user3"],
		"group2": ["user1", "user4"],
		"group3": ["user5", "user3"]
	},
	"telegram": {
		"botid": "botid-xxx",
		"chatid": "chatid-xxx"
	}
}
```
###### IRC Commands
`!clear` clears all alerts,
`!reset` resets failed state.


###### Help

``` text
 ./bot-daemon -h
 Usage of ./bot-daemon:
  -alertmanager.addr string
    	alertmanager webhook listen address (default "127.0.0.1:9500")
  -cfg string
    	bot configuration file (default "bot.json")
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
