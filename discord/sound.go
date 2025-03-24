package discord

import (
	"encoding/binary"
	"fmt"
    "io"
	"os"
	"time"

	"github.com/bwmarrin/discordgo"
)

// PlaySound sends the airhorn sound to the provided channel.
func PlaySound(s *discordgo.Session, guildID, channelID, fileName string) error {
	var err error
	f, err := os.Open(fmt.Sprintf("discord/%s.dca", fileName))
	if err != nil {
		return fmt.Errorf("couldn't open file: %v", err)
	}
	defer f.Close()

	vc, err := s.ChannelVoiceJoin(guildID, channelID, false, true)
	if err != nil {
		return err
	}

	vc.Speaking(true)

	var opusLen int16
	for {
		// Read opus frame length from dca file.
		err = binary.Read(f, binary.LittleEndian, &opusLen)

		// If this is the end of the file, just return.
		if err == io.EOF || err == io.ErrUnexpectedEOF {
			err = nil
			break
		}

		if err != nil {
			break
		}

		// Read encoded pcm from dca file.
		inBuf := make([]byte, opusLen)
		err = binary.Read(f, binary.LittleEndian, &inBuf)
		if err != nil {
			break
		}

	    vc.OpusSend <- inBuf
	}
	if err != nil {
		return fmt.Errorf("couldn't read dca file: %v", err)
	}

	// Stop speaking
	vc.Speaking(false)

	// Sleep for a specificed amount of time before ending.
	time.Sleep(100 * time.Millisecond)

	// Disconnect from the provided voice channel.
	vc.Disconnect()

	return nil
}
