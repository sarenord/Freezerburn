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
	Guild string
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
	//ignore the bot's own messages
	if m.Author.ID == s.State.User.ID {
		return
	}
	t, err := m.Timestamp.Parse()
	tString := time.Time.String(t)
	stamp := strings.Split(tString, " ")
	tString = strings.Join(strings.Split(stamp[1], "")[:8], "")
	
	fmt.Println(t, m.Author.Username,":", m.Content)
	if err != nil {
		fmt.Println(err)
		return
	}
}
