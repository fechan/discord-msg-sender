package main

import (
	"flag"
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
	botTokenPtr := flag.String("bot_token", os.Getenv("BOT_TOKEN"), "ID of the channel to send to")
	channelIdPtr := flag.String("channel_id", os.Getenv("CHANNEL_ID"), "ID of the channel to send to")
	messagePtr := flag.String("message", os.Getenv("MESSAGE"), "The message to send")
	flag.Parse()

	// Log in as a bot
	discord, err := discordgo.New("Bot " + *botTokenPtr)
	if err != nil {
		log.Fatal(err)
	}

	// Send the message
	_, err = discord.ChannelMessageSend(*channelIdPtr, *messagePtr)
	if err != nil {
		log.Fatal(err)
	}
}
