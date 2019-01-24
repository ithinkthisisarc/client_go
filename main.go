package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"github.com/fatih/color"
)

// Variables used for command line parameters
var (
	Token    string
	Help     string
	msg      string
	inp      string
	servers  []*discordgo.Guild
	channels []*discordgo.Channel
)

func init() {
	flag.StringVar(&Help, "h", "f", "Display helper ids")
	flag.StringVar(&Token, "t", "NDQxMzQyMDYxMjc1MzgxNzYz.DyCt9Q.cXcEH2JH8phxCHTTSdSO88JHImE", "Bot Token")
	flag.Parse()
}

func main() {
	// Create a new Discord session using the provided bot token.
	dg, err := discordgo.New(Token)
	if err != nil {
		fmt.Println("error creating Discord session,", err)
		return
	}

	fmt.Println("Welcome to Client_Go, a discord command line client written in go!")
	fmt.Print("Enter the ID of the channel you would like to talk in: ")
	if Help == "t" {
		fmt.Print("\n\nSaved Ids:\n DN:general:\t422293824770146306\n PH:general:\t469851459966730262\n FC:general:\t439871916082331650\n: ")
	}
	inp = read()
	fmt.Println("Connecting... (this may take a while)")

	// Register the messageCreate func as a callback for MessageCreate events.
	dg.AddHandler(messageCreate)

	// Open a websocket connection to Discord and begin listening.
	err = dg.Open()
	if err != nil {
		fmt.Println("error opening connection,", err)
		return
	}

	// Wait here until CTRL-C or other term signal is received.
	fmt.Print("Bot is now running.  Press CTRL-C to exit.\n\n-----------------------------------\n")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	// Cleanly close down the Discord session.
	dg.Close()
}

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	loc := m.ChannelID
	if loc == inp {
		// If you sent the message, don't show it
		if m.Author.ID != s.State.User.ID {
			temp := color.New(color.FgCyan).PrintfFunc()
			author := "\n >>> " + m.Author.Username
			temp("%s", author)
			fmt.Println(": " + m.Content)
		} else {
			temp := color.New(color.FgRed).PrintfFunc()
			author := "\n >>> You"
			temp("%s", author)
			fmt.Println(": " + m.Content)
		}
		msg = read()
		if strings.HasPrefix(msg, ">>") {
			check_for_cmd(msg)
		} else {
			s.ChannelMessageSend(inp, msg)
		}
	}
}
