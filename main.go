package main

// import packages
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

// define variables
var Auth struct {
	ChannelName string
	OauthToken  string
}
var Delay int64
var exitChannel chan os.Signal

// this function is always called first on start
func init() {
	// Parse Commandline parameter
	flag.StringVar(&Auth.ChannelName, "channel", "", "the channel where to ban")
	flag.StringVar(&Auth.OauthToken, "token", "", "your token to log into chat")
	flag.Int64Var(&Delay, "delay", 3, "time in seconds between ban messages")
	flag.Parse()
}

// this is the main function here is the entry point
func main() {
	// Check if parameters set
	if Auth.ChannelName != "" || Auth.OauthToken != "" {
		log.Println("[Info] Token and Channel set")

		// calling an async process to let it run whithout
		// obstructing anything else
		go func() {
			log.Println("[Info] Starting Client Connection")
			// create connection to twitch
			client := twitch.NewClient(Auth.ChannelName, Auth.OauthToken)

			// this gets called when the client connects to twitch
			client.OnConnect(func() {
				log.Println("[Info] Connected")
				log.Println("[Info] Read BanList")

				// calls the readline function (defined at the bottom of this file)
				lines, err := readLines("./banlist.txt")
				if err != nil {
					log.Println(err)
					return
				}
				log.Println("[Info] Start ban process")

				// a for each for every line in the text file
				// i = current array position
				// v = current array position data
				for i, v := range lines {
					log.Printf("Banned: %d/%d %s", i+1, len(lines)+1, v)
					client.Say(Auth.ChannelName,
						fmt.Sprintf("/ban %s Banned by Banlist", v), // format message for ban
					)
					// wait for x seconds to avoid errors
					// there is an limit for messages/
					time.Sleep(time.Second * time.Duration(Delay))
				}
			})

			// here is the channel to join defined
			client.Join(Auth.ChannelName)

			log.Println("[Info] Connecting....")
			// connect to twitch
			err := client.Connect()
			if err != nil {
				log.Println(err)
				return
			}
		}()
	} else {
		// this get printed if no channelname or token is set
		log.Println("[Error]: No Channelname or Token provided")
		fmt.Println("visit https://twitchapps.com/tmi/ to get an oAuth token")
		fmt.Println("use '-channel <channelname> -token <token>' in commandline")
		fmt.Println("Example:\ntwitchbanfromlist.exe -channel noodlesoup -token oauth:dwiuaofhmcmiouehjiuwqehx")
		fmt.Println("or use '-help' for more options")
	}

	// this keeps the process running because we called
	// the connection to twitch in async-function
	// important if some new stuff gets implemented
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

// read line function from file
func readLines(path string) ([]string, error) {
	// open the file
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	// when this function exit this will free the file
	// important because if its not called there can be dragons
	defer file.Close()

	// and here read the file line by line in an array
	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	// return the array and if there is an error that message too
	return lines, scanner.Err()
}
