package bot

import (
	"fmt"
	"strings"

	"github.com/Pradumnasaraf/discord-bot/config"
	"github.com/bwmarrin/discordgo"
)

var (
	BotID string
	goBot *discordgo.Session
)

type Answer struct {
	OriginalChannelID string
	FavFood           string
	FavGame           string
}

var responses = make(map[string]Answer)

func Start() {
	fmt.Println("Bot is starting...")

	goBot, err := discordgo.New("Bot " + config.Token)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	user, err := goBot.User("@me")
	if err != nil {
		fmt.Println(err.Error())
	}

	BotID = user.ID

	goBot.AddHandler(messageHandler)

	goBot.Identify.Intents = discordgo.IntentsAllWithoutPrivileged

	err = goBot.Open()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	defer goBot.Close()

	fmt.Println("Bot is online...")

}

func messageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {

	if m.Author.ID == BotID || m.Author.Bot {
		return
	}
	
	args := strings.Split(m.Content, " ")
	if args[0] != config.BotPrefix {
		return
	}

	// /oss
	if len(args) == 1 {
		_, _ = s.ChannelMessageSend(m.ChannelID, "Yes, I am listening")
		return
	}

	// /oss ping
	if args[1] == "ping" {
		_, err := s.ChannelMessageSend(m.ChannelID, "Pong")
		if err != nil {
			fmt.Println("Error sending message:", err)
			return
		}
	}

	if args[1] == "author" {

		author := discordgo.MessageEmbedAuthor{
			Name: "Pradumna Saraf",
			URL:  "https://os.pradumnasaraf.dev",
		}

		embed := discordgo.MessageEmbed{
			Title:       "Open Source Bot",
			Author:      &author,
			Color:       0x00ff00,
			Description: "This is a bot to help you with open source",
		}

		_, err := s.ChannelMessageSendEmbed(m.ChannelID, &embed)
		if err != nil {
			fmt.Println("Error sending embed message:", err)
			return
		}

	}

	if args[1] == "prompt" {
		channel, err := s.UserChannelCreate(m.Author.ID)
		if err != nil {
			fmt.Println("Error creating channel:", err)
			return
		}

		// This will create a new entry in the map if the user doesn't already exist
		if _, ok := responses[m.Author.ID]; !ok {
			responses[m.Author.ID] = Answer{
				OriginalChannelID: m.ChannelID,
				FavFood:           "",
				FavGame:           "",
			}
		}
		s.ChannelMessageSend(channel.ID, "Hey, here are some questions for you")
		s.ChannelMessageSend(channel.ID, "What is your favorite food?")
	} else {
		s.ChannelMessageSend(m.ChannelID, "We are still working on this command")
	}

}
