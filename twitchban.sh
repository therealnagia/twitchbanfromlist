#!/bin/bash

SOURCE="$(dirname ${BASH_SOURCE[0]})/banlist.txt"

echo "Obtaining latest banlist..."
rm -f banlist.txt
wait
wget -q https://raw.githubusercontent.com/therealnagia/twitchbanfromlist/main/banlist.txt -O banlist.txt
wait

if [ -f "banlist.txt" ]; then
    echo "Latest banlist obtained"
else
    echo "Failed to obtain latest banlist. Do you have an internet connection or permission to write to this folder?"
    echo "Exiting."
    exit 1
fi

echo -n "Starting in 5 seconds..."
for i in `seq 5`; do
    sleep 1
    echo -n .
done
echo

while IFS="" read -r p || [ -n "$p" ]; do
    if [ "$p" == "BANLISTEND" ]; then
        echo "Everyone on the list has been banned!"
        exit 0
    fi
    xdotool type "/ban $p"
    xdotool key Return
    sleep 0.4
done < "$SOURCE"
