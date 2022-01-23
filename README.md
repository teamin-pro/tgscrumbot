# SCRUM Poker Bot

Simplest bot for Scrum poker.

### Install
1. Create new group chat in Telegram
2. Add @ShantiScrumBot
3. Make @ShantiScrumBot group chat admin

### Usage
1. just make votes in chat. Vote is message with positive integer number
2. for results or reset write message with any number of `-` symbols (e.g `-`, `----`, `------`)

Text messages will be ignored. If you want change vote, just send new value to chat.

Example:
```shell
Alice: 2222222
Bob: 10
Carol: 5
Alice: Sorry, cat on keyboard...
Alice: 2
Scrum-master: ---
ShantiScrumBot: Num votes: 3, average: 17
```

### Run on server
```shell
go get -u github.com/shantilabs/shantiscrumbot
$HOME/go/bin/shantiscrumbot -token '0000000000:xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx' -vote-message='Num votes: %d, average: %d' 
```

### History
Based on bot for [tada.team](https://tada.team) task messenger: https://github.com/tada-team/tdhooks
