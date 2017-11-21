package main
import(
	"github.com/bwmarrin/discordgo"
	"fmt"
)

func (bot *Object) ClearChan(s *discordgo.Session, m *discordgo.Message) {
	idar, err := s.ChannelMessages(m.ChannelID, 100, "", "", "")

	for _, message := range idar  {
		s.ChannelMessageDelete(m.ChannelID, message.ID)
	}
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("done")
}
