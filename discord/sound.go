package discord

import (
	"encoding/binary"
	"fmt"
    "io"
	"os"
	"time"

	"github.com/bwmarrin/discordgo"
)

type Player struct {
	buffer [][]byte
}

// LoadSound attempts to load an encoded sound file from disk.
func LoadSound() (*Player, error) {
	file, err := os.Open("discord/airhorn.dca")
	if err != nil {
		return nil, fmt.Errorf("error opening file: %v", err)
	}

	var opuslen int16

	var buffer = make([][]byte, 0)

	for {
		// Read opus frame length from dca file.
		err = binary.Read(file, binary.LittleEndian, &opuslen)

		// If this is the end of the file, just return.
		if err == io.EOF || err == io.ErrUnexpectedEOF {
			err := file.Close()
			if err != nil {
				return nil, err
			}
			return &Player{
				buffer: buffer,
			}, nil
		}

		if err != nil {
			fmt.Println("Error reading from dca file :", err)
			return nil, err
		}

		// Read encoded pcm from dca file.
		InBuf := make([]byte, opuslen)
		err = binary.Read(file, binary.LittleEndian, &InBuf)

		// Should not be any end of file errors
		if err != nil {
			fmt.Println("Error reading from dca file :", err)
			return nil, err
		}

		// Append encoded pcm data to the buffer.
		buffer = append(buffer, InBuf)
	}
}

// PlaySound plays the current buffer to the provided channel.
func (p *Player) PlaySound(s *discordgo.Session, guildID, channelID string) (err error) {

	// Join the provided voice channel.
	vc, err := s.ChannelVoiceJoin(guildID, channelID, false, true)
	if err != nil {
		return err
	}

	// Start speaking.
	vc.Speaking(true)

	// Send the buffer data.
	for _, buff := range p.buffer {
		vc.OpusSend <- buff
	}

	// Stop speaking
	vc.Speaking(false)

	// Sleep for a specificed amount of time before ending.
	time.Sleep(250 * time.Millisecond)

	// Disconnect from the provided voice channel.
	vc.Disconnect()

	return nil
}
