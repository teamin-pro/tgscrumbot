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
1. just make votes in chat. Vote is message with positive integer number
2. for results or reset write message with any number of `-` symbols (e.g `-`, `----`, `------`)

Text messages will be ignored. If you want change vote, just send new value to chat.

### Run on server (if you want own instance of this bot)
```shell
go get -u github.com/shantilabs/shantiscrumbot
$HOME/go/bin/shantiscrumbot -token '0000000000:xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx' -vote-message='Num votes: %d, average: %d' 
```

### History
Based on bot for [tada.team](https://tada.team) task messenger: https://github.com/tada-team/tdhooks

### Playground
https://t.me/+oIw9EKVrg8s5MGJi
