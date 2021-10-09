Main:
msgbox,
(
TwitchBanFromList v1.211008b
-----------------------------------
This script can BAN or BLOCK (or both) 
a large amount of users from a txt file. 

Instructions:
	1. Click 'OK' on this box (do this last or this box disappears)
	2. Go to your twitch chat click in the text entry field
	2. Push F12 to Ban OR Push F11 to Block OR F10 for both
	3. DO NOT TOUCH YOUR MOUSE OR 
	   KEYBOARD WHILE THE SCRIPT RUNS 
	4. Watch as all of the names on the list get banned 
	   one after the other in your chat box!
	5. At the end of the banlist, Whitelist unbanning will occur
	5. A box will pop up at the end when the script 
	   completes! (script will exit)
)	
return

Esc::ExitApp

F12::
FileRead, banlist, %A_ScriptDir%\banlist.txt
Loop, Parse, banlist, `n, `r
{
	if (A_LoopField = "BANLISTEND")
	{
		WhitelistRun()
		msgbox, Everyone on the list has been banned! `nExiting!
		ExitApp
	}	
	clipboard := "/ban " . A_LoopField
	PasteEnterFcn()
}
Return

F11::	
FileRead, banlist, %A_ScriptDir%\banlist.txt
Loop, Parse, banlist, `n, `r
{
	if (A_LoopField = "BANLISTEND")
	{
		WhitelistRun()
		msgbox, Everyone on the list has been [[BLOCKED]]! `nExiting!
		ExitApp
	}	
	clipboard := "/block " . A_LoopField	
	PasteEnterFcn()
}
Return

F10::	
FileRead, banlist, %A_ScriptDir%\banlist.txt
Loop, Parse, banlist, `n, `r
{
	if (A_LoopField = "BANLISTEND")
	{
		WhitelistRun()
		msgbox, Everyone on the list has been [[BANNED]] AND [[BLOCKED]]! `nExiting!
		ExitApp
	}	
	clipboard := "/ban " . A_LoopField
	PasteEnterFcn()
	clipboard := "/block " . A_LoopField
	PasteEnterFcn()
}
Return

PasteEnterFcn()
{
	sleep 100
	send ^v
	sleep 50
	Send {Enter}
	sleep 100
}
WhitelistRun()
{
FileRead, whitelist, %A_ScriptDir%\whitelist.txt
	Loop, Parse, whitelist, `n, `r
	{
		if (A_LoopField = "LISTEND")
		{
			break
		}	
		clipboard := "/unban " . A_LoopField
		PasteEnterFcn()
		clipboard := "/unblock " . A_LoopField
		PasteEnterFcn()
	}
}
