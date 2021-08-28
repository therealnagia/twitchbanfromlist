Esc::ExitApp

F12::
FileRead, banlist, %A_ScriptDir%\banlist.txt
Loop, Parse, banlist, `n, `r
{
	if (A_LoopField = "BANLISTEND")
	{
		msgbox, Everyone on the list has been banned! `nExiting!
		Exit
	}	
	Send, /ban{Space}
	SendRaw %A_LoopField%
	Send {Enter}
	sleep 200
}
Return


