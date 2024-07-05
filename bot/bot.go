package bot

import (
	"fmt"
	"log"
	"github.com/bwmarrin/discordgo"
	"github.com/trynax/inj-price-checker/config"
	"github.com/trynax/inj-price-checker/price"
)

var BotID string
var goBot *discordgo.Session


func Start() {
	var err error

	goBot, err = discordgo.New("Bot " + config.Token)
	if err != nil {
		log.Fatalf("error creating Discord session: %v", err)
		return
	}


	goBot.Identify.Intents = discordgo.IntentsGuildMessages | discordgo.IntentsDirectMessages | discordgo.IntentMessageContent
	u, err := goBot.User("@me")
	if err != nil {
		log.Fatalf("error retrieving account information: %v", err)
		return
	}

	BotID = u.ID
	goBot.AddHandler(messageHandler)
	goBot.AddHandler(guildCreateHandler)

	err = goBot.Open()
	if err != nil {
		log.Fatalf("error opening connection: %v", err)
		return
	}
	fmt.Println("Bot is now running....")
}

func guildCreateHandler(s *discordgo.Session, event *discordgo.GuildCreate) {
	if event.Guild.Unavailable {
		return
	}

	for _, channel := range event.Guild.Channels {
		if channel.Type == discordgo.ChannelTypeGuildText {
			_, err := s.ChannelMessageSend(channel.ID, "Hello! I am now added to this server.")
			if err != nil {
				log.Printf("error sending welcome message: %v", err)
			} else {
				fmt.Printf("Sent welcome message to channel %s of guild %s\n", channel.ID, event.Guild.ID)
			}
			break
		}
	}
}

func messageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == BotID {
		return
	}

	fmt.Printf("Message received in channel %s (guild %s): %s\n", m.ChannelID, m.GuildID, m.Content) // Log the channel ID, guild ID, and message content

	if m.Content == config.BotPrefix+"INJ" {
		fmt.Println("INJ command received") 
		price := price.CheckPrice("INJ")
		message := fmt.Sprintf("Current INJ price: %.2f USD", price)
		fmt.Printf("Sending message: %s\n", message)
		_, err := s.ChannelMessageSend(m.ChannelID, message)
		if err != nil {
			log.Printf("error sending message: %v", err)
		}
	}

	if m.Content == config.BotPrefix + "QUNT" {
		fmt.Println("QUNT commend received")
		price := price.CheckPrice("QUNT")
		message := fmt.Sprintf("Current QUNT price: %.8f USD", price)
		fmt.Printf("Sending message: %s\n", message)	
		_, err := s.ChannelMessageSend(m.ChannelID, message)
		if err!= nil {
            log.Printf("error sending message: %v", err)
        }
	}
}
