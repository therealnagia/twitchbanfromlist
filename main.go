package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gempir/go-twitch-irc/v2"
)

var Auth struct {
	ChannelName string
	OauthToken  string
}
var Delay int64
var exitChannel chan os.Signal

func init() {
	// Parse Commandline parameter
	flag.StringVar(&Auth.ChannelName, "channel", "", "the channel where to ban")
	flag.StringVar(&Auth.OauthToken, "token", "", "your token to log into chat")
	flag.Int64Var(&Delay, "delay", 3, "time in seconds between ban messages")
	flag.Parse()
}

func main() {
	// Check if parameters set
	if Auth.ChannelName != "" || Auth.OauthToken != "" {
		log.Println("[Info] Token and Channel set")
		func() {
			log.Println("[Info] Starting Client Connection")
			client := twitch.NewClient(Auth.ChannelName, Auth.OauthToken)

			client.OnConnect(func() {
				log.Println("[Info] Connected")
				log.Println("[Info] Read BanList")
				lines, err := readLines("./banlist.txt")
				if err != nil {
					log.Println(err)
					return
				}
				log.Println("[Info] Start ban process")
				for i, v := range lines {
					log.Printf("Banned: %d/%d %s", i+1, len(lines)+1, v)
					client.Say(Auth.ChannelName, fmt.Sprintf("/ban %s Banned by Banlist", v))
					time.Sleep(time.Second * time.Duration(Delay))
				}
			})

			client.Join(Auth.ChannelName)

			log.Println("[Info] Connecting....")
			err := client.Connect()
			if err != nil {
				log.Println(err)
				return
			}
		}()
	} else {
		log.Println("[Error]: No Channelname or Token provided")
		fmt.Println("visit https://twitchapps.com/tmi/ to get an oAuth token")
		fmt.Println("use '-channel <channelname> -token <token>' in commandline")
		fmt.Println("Example:\ntwitchbanfromlist.exe -channel noodlesoup -token oauth:dwiuaofhmcmiouehjiuwqehx")
		fmt.Println("or use '-help' for more options")
	}

	// Wait for Terminate Process
	log.Println("press CTRL-C to exit")
	exitChannel = make(chan os.Signal, 1)
	signal.Notify(exitChannel, syscall.SIGTERM, syscall.SIGINT)
	signal := <-exitChannel
	if signal == syscall.SIGTERM {
		log.Println("exit")
	} else if signal == syscall.SIGINT {
		log.Println("exit")
	}
}

func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}
