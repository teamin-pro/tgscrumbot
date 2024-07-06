# SCRUM Poker Bot

Simplest Telegram bot for [Scrum Planning Poker](https://en.wikipedia.org/wiki/Planning_poker):
```shell
Scrum-master: Ready? 3... 2... 1... GO!
Alice: 2
Bob: 10
Carol: 5
Scrum-master: ---
TgScrumBot: Num votes: 3, average: 6
```

## Install
1. Create new group chat in Telegram
2. Add [@TgScrumBot](https://t.me/TgScrumBot)
3. Make [@TgScrumBot](https://t.me/TgScrumBot) group chat admin

## Usage
1. just make votes in chat. A vote is a message containing only a positive integer
2. to get results or to reset current vote send a message with any number of `-` symbols (e.g `-`, `----`, `------`)

Text messages will be ignored. If you want to change your vote, just send a new value to the chat.

```shell
Alice: 222222222222222
Alice: Sorry, cat on keyboard...
Alice: 2
Bob: 4
Scrum-master: ---
TgScrumBot: Num votes: 2, average: 3
```

## Playground

This is real demo chat with curious persons:
https://t.me/+oIw9EKVrg8s5MGJi

## Run on server 
This step is for those who want to run their own bot on their server. It may be useful for security reasons or 
if you want to customize the bot.

First register new bot in Telegram and get token. Then run:

```shell
go get -u github.com/teamin-pro/tgscrumbot
$HOME/go/bin/tgscrumbot -token '0000000000:xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx' -vote-message='Num votes: %d, average: %d' 
```

Also you can use Ansible to deploy bot on your server:
- Clone this repo
- Modify files in `ansible/` and `.github/workflows` directories if needed
- Change secrets and variables in repository settings
- Commit changes
