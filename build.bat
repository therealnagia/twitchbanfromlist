@echo off
@title Build

set filename=twitchbanfromlist

echo Build Windows Release
set GOOS=windows
set GOARCH=amd64
go build -o Builds/%filename%-windows.exe

echo Build Linux Release
set GOOS=linux
set GOARCH=amd64
go build -o Builds/%filename%-linux_64

echo Build RaspberryPi Release
set GOOS=linux
set GOARCH=arm
set GOARM=5
go build -o Builds/%filename%-pi

echo Build Mac Release
echo DERP this will never happen DERP