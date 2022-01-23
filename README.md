# SCRUM Poker Bot

Simplest Telegram bot for Scrum poker:
```shell
Scrum-master: 3... 2... 1... GO!
Alice: 2222222
Bob: 10
Carol: 5
Alice: Sorry, cat on keyboard...
Alice: 2
Scrum-master: ---
ShantiScrumBot: Num votes: 3, average: 6
```

### Install
1. Create new group chat in Telegram
2. Add [@ShantiScrumBot](https://t.me/ShantiScrumBot)
3. Make [@ShantiScrumBot](https://t.me/ShantiScrumBot) group chat admin

### Usage
1. just make votes in chat. A vote is a message containing only a positive integer
2. to get results or to reset current vote send a message with any number of `-` symbols (e.g `-`, `----`, `------`)

Text messages will be ignored. If you want to change your vote, just send a new value to the chat.

### Run on server (if you want own instance of this bot)
```shell
go get -u github.com/shantchat/tgscrumbot
$HOME/go/bin/tgscrumbot -token '0000000000:xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx' -vote-message='Num votes: %d, average: %d' 
```

### Playground
https://t.me/+oIw9EKVrg8s5MGJi
