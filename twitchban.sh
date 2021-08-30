#!/bin/bash

SOURCE="$(dirname ${BASH_SOURCE[0]})/banlist.txt"

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
    sleep 0.2
done < "$SOURCE"
