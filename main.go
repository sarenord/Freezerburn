package main
import (
	"fmt"
	"strings"
	"io/ioutil"
	"encoding/json"
	"os"
	"os/signal"
	"time"
	"github.com/bwmarrin/discordgo"
)

type Object struct {
	Token string
	CommandChar string
}
var bot Object
var err error
var start int64

func main() {
	start=time.Now().UnixNano()
	//load config
	io, err := ioutil.ReadFile("config.json")
	if err != nil {
		fmt.Println(err)
		return
	}
	json.Unmarshal(io, &bot)

	//generate the session
	discordSession, err := discordgo.New("Bot " + bot.Token)
	if err != nil {
		fmt.Println(err)
		return
	}
	//handlers
	discordSession.AddHandler(ready)
	discordSession.AddHandler(messageCreate)

	//make the socket
	err = discordSession.Open()
	if err !=  nil{
		fmt.Println(err)
		return
	}


	//somehow this kills the program on CTRL+C
	fmt.Println("bot is now running.")
	//this creates a channel to send the signal
	//to the actual program
	sc := make(chan os.Signal, 1)
	//if the signal is os.Interrupt, pass it to the program? maybe?
	signal.Notify(sc, os.Interrupt)
	<-sc
	discordSession.Close()
}

func ready(s *discordgo.Session, event *discordgo.Ready) {
	fmt.Println("received ready signal")
	return
}

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	logGen(m)
	//ignore the bot's own messages
	if m.Author.ID == s.State.User.ID {
		return
	}
	if strings.Contains(m.Content, bot.CommandChar){
		command := strings.Split(strings.TrimLeft(m.Content, bot.CommandChar), " ")
		switch command[0] {
			case "Clearchan":
			s.ChannelMessageSend(m.ChannelID, "clearing channel")
			bot.ClearChan(s, m.Message)
			s.ChannelMessageSend(m.ChannelID, "Channel Cleared"
			break;
		default:
			return
		}
	}
}

func logGen(m *discordgo.MessageCreate) { 
	t, err := m.Timestamp.Parse()
	if err != nil {
		fmt.Println(err)
		return
	}
	tString := strings.Split(t.String(), " ")
	stamp := []string{tString[0], strings.Join(strings.Split(tString[1], "")[:8], "")}
	fmt.Println(stamp)
	logline := []string{stamp[0], stamp[1], strings.Join([]string{m.Author.Username, ":"}, ""), m.Content}
	fmt.Println(strings.Join(logline, " "))
}

