package main

import (
	"log"
	"os"

	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
)

func main() {
	// Try to load .env file, or skip loading if it doesn't exist
	_, err := os.Stat(".env")
	if !os.IsNotExist(err) {
		err = godotenv.Load()
		if err != nil {
			log.Fatal(err)
		}
	}

	// Parse environment variables from command line and falling back to .env
	botToken := os.Getenv("BOT_TOKEN")
	channelId := os.Getenv("CHANNEL_ID")
	message := os.Getenv("MESSAGE")

	// Log in as a bot
	discord, err := discordgo.New("Bot " + botToken)
	if err != nil {
		log.Fatal(err)
	}

	// Send the message
	_, err = discord.ChannelMessageSend(channelId, message)
	if err != nil {
		log.Fatal(err)
	}
}
