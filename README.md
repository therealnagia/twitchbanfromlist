# TwitchBanFromList
## AHK text entry script that bans people from a text file.

Very hacked together but "simple" way of using AHK to ban a bunch of identified bot accounts, 
You can even run it before stream to make sure that you have the latest ban list (if you trust the list below)


=======
### CAUTION:
- While there are many happy people that have used this script, PLEASE USE THIS SCRIPT AT YOUR OWN RISK
- Per [Issue #15](https://github.com/therealnagia/twitchbanfromlist/issues/15), some users report experiencing temporary shadow bans from using this script with lists that are very large. 
- This script will be reworked to add more safety options in the near future. 
- Since there have been a few (2-3) usernames that have been brought to my attention that are legit accounts, a whitelist.txt has been created. 
- - It will run after every ban command, empty it if you don't want to whitelist anyone. 	


### Features:
	* Ban from list with F12 Key
	* Block from list with F11 Key
	* Perform both ban and block with F10 key
	* Runs through a whitelist after banning and blocking

## Instructions:
	1. Go to the releases page, downloaded the latest exe
	2. Run the exe.
	3. Go to your twitch chat click in the text entry field
	4. Push F12 to Ban OR Push F11 to Block 
	5. DO NOT TOUCH YOUR MOUSE OR KEYBOARD WHILE THE SCRIPT RUNS 
	6. Watch as all of the names on the list get banned one after the other in your chat box!
	7. A box will pop up at the end when the script completes! (script will exit)
	
### Linux Instructions:
	Required packages: xdotool
	1. Run bash script via 'bash twitchban.sh' (5 second timer starts)
	2. Click into your twitch chat text entry field
	3. After 5 seconds of running twitchban.sh, it should automatically start running through the ban list
	

##	Credits
###	Ban list sources: 
https://docs.google.com/document/d/1_F3qKiwkECmHYJvHv4hevkOYWundzNewpC_PcSGBj1I/edit

twitter@MissRayvenn: https://docs.google.com/document/d/1F5tMM2oTDlJJ41320vJFd8MVDW_mGKnU_Q9RFWdY-fM/edit 

twitter@MissRayvenn: https://pastebin.com/FyxGesEw 

gh@iguanaonmystack / gh@biosniper: Linux script port 

### More list contributors!
twitter@OpalStreams, gh@liadala, gh@biosniper, twitter@ItsRogueRen


### Similar names but legit accounts:
Sources: 
https://twitter.com/MissRayvenn/status/1432409056459776000
https://twitter.com/StygianStyx/status/1446779342818889729

If latest release is after last updated date/time then it should have the latest, otherwise, notify me, and/or just submit a PR

