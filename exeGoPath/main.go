package main

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

func main() {
	fmt.Println("Hello, Please check the mod file for for GoPath!!! Very Important")
	discord, err := discordgo.New("Bot " + "authentication token")

	fmt.Println(discord)
	fmt.Println("--------")
	fmt.Println(err)
}
