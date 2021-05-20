package bot

import (
	"fmt"
	"strings"
	"test/testBot/config"

	"github.com/bwmarrin/discordgo"
)

var BotID string

func Init() {
	testBot, err := discordgo.New("Bot " + config.Token)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	u, err := testBot.User("@me")

	if err != nil {
		fmt.Println(err.Error())
	}

	BotID = u.ID

	testBot.AddHandler(messageHandler)

	err = testBot.Open()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("Bot is running!")

	<-make(chan struct{})
	return
}

func messageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	fmt.Println(m.Author, " said '", m.Content, "'")

	if strings.HasPrefix(m.Content, config.BotPrefix) {

		if m.Author.ID == BotID {
			return
		}

		if m.Content == "!ping" {
			_, _ = s.ChannelMessageSend(m.ChannelID, "pong")
		}
	}
}
