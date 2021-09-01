# TwitchBanFromList

## AHK text entry script that bans people from a text file.
## AutoBan from File

You can even run it before stream to make sure that you have the latest ban list (if you trust the list below)

## Instructions:
1. Go to the releases page, downloaded the latest exe
2. Get an login token from [https://twitchapps.com/tmi/](https://twitchapps.com/tmi/)
3. Run the exe from powershell
> Example: `.\twitchbanfromlist.exe -channel yourchannel -token oauth:wfaest45z4wtg9waf92`

## Linux Instructions:
### Instructions:
1. Go to the releases page, downloaded the latest exe
2. Get an login token from [https://twitchapps.com/tmi/](https://twitchapps.com/tmi/)
3. Run the exe from powershell
> Example: `.\twitchbanfromlist -channel yourchannel -token oauth:wfaest45z4wtg9waf92`

## Build from Source
1. Go to [https://golang.org/](https://golang.org/) download the latest version
2. Clone the Repository
3. Run `go mod download`
4. Run `go build`
5. Run `twitchbanfromlist` like the examples above
##	Credits
###	Ban list source: 
https://docs.google.com/document/d/1_F3qKiwkECmHYJvHv4hevkOYWundzNewpC_PcSGBj1I/edit
###	Linux script port: iguanaonmystack

If latest release is after last updated date/time then it should have the latest, otherwise, notify me, and/or just submit a PR

