package discord

import (
	"fmt"
	"strings"

	"github.com/bwmarrin/discordgo"
)

// Ready will be called by AddHandler when the bot receives
// the "ready" event from Discord.
func Ready(s *discordgo.Session, _ *discordgo.Ready) {

	// Set the playing status.
	s.UpdateGameStatus(0, "!airhorn")
}

func MessageCreateWithFileName(fileName string) func(*discordgo.Session, *discordgo.MessageCreate) {
	return func (s *discordgo.Session, m *discordgo.MessageCreate) {
		// Ignore all messages created by the bot itself
		// This isn't required in this specific example but it's a good practice.
		if m.Author.ID == s.State.User.ID {
			return
		}

		// check if the message is "!airhorn"
		if strings.HasPrefix(m.Content, "!"+fileName) {

			// Find the channel that the message came from.
			c, err := s.State.Channel(m.ChannelID)
			if err != nil {
				// Could not find channel.
				return
			}

			// Find the guild for that channel.
			g, err := s.State.Guild(c.GuildID)
			if err != nil {
				// Could not find guild.
				return
			}

			// Look for the message sender in that guild's current voice states.
			for _, vs := range g.VoiceStates {
				if vs.UserID == m.Author.ID {
					err = PlaySound(s, g.ID, vs.ChannelID, fileName)
					if err != nil {
						fmt.Println("Error playing sound:", err)
					}

					return
				}
			}
		}
	}
}

// GuildCreate is called by AddHandler every time a new
// guild is joined.
func GuildCreate(s *discordgo.Session, event *discordgo.GuildCreate) {

	if event.Guild.Unavailable {
		return
	}

	for _, channel := range event.Guild.Channels {
		if channel.ID == event.Guild.ID {
			_, _ = s.ChannelMessageSend(channel.ID, "Airhorn is ready! Type !airhorn while in a voice channel to play a sound.")
			return
		}
	}
}
