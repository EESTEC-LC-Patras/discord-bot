package bot

import (
	"fmt"
	"math/rand"
	"strings"
	"test/testBot/config"
	"time"

	"github.com/bwmarrin/discordgo"
)

var BotID string

func Init() {
	testBot, err := discordgo.New("Bot " + config.Token)
	rand.Seed(time.Now().UTC().UnixNano())

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

	fmt.Println("Τρέχει ο βότος!")

	<-make(chan struct{})
	return
}

func messageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	if strings.HasPrefix(m.Content, config.BotPrefix) {

		if m.Author.ID == BotID {
			return
		}

		switch m.Content {
		case "!greet":
			if m.Author.Username == "chriskats" {
				_, _ = s.ChannelMessageSend(m.ChannelID, "Γεια σου ρε μάγκα.")
				break
			}
			if time.Now().UnixNano()/1000/2%2 == 0 {
				_, _ = s.ChannelMessageSend(m.ChannelID, "Γεια σου, "+m.Author.Username+"!")
			} else {
				_, _ = s.ChannelMessageSend(m.ChannelID, "Πώς είσαι, "+m.Author.Username+";")
			}
		case "!help":
			_, _ = s.ChannelMessageSend(m.ChannelID, printHelp(config.BotPrefix))
		case "!lang":
			fmt.Println("Ορίστηκε γλώσσα: " + config.Language)
			if config.Language == "GR" {
				config.Language = "EN"
			} else {
				config.Language = "GR"
			}
		case "!quest":
			//todo
		}
	}
}

func printHelp(botPrefix string) string {
	return "Αυτές είναι οι εντολές που εκτελώ!\n\t•" + botPrefix + "greet\n\t•" + botPrefix + "help\n\t•" + botPrefix + "lang"
}
