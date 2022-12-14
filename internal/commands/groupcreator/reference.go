package groupcreator

import (
	"errors"
	"fmt"

	"github.com/NilsPonsard/verbosity"
	"github.com/bwmarrin/discordgo"
	configurechannel "github.com/do3-2021/booty-mover/internal/commands/configure-channel"
	"github.com/do3-2021/booty-mover/internal/database"
	"github.com/do3-2021/booty-mover/internal/guild"
)

// Add a message to the configured role channel
func ReferenceRoleInChannel(s *discordgo.Session, i *discordgo.InteractionCreate, group string, description string, roleID string) (err error) {
	db, err := database.GetDB()

	if err != nil {
		verbosity.Error(err.Error())
		return errors.New("could not contact database to get the role listing channel's ID.\nCancelling Group's creation")
	}

	channel, err := guild.GetGroupChannel(db, i.GuildID)
	if err != nil {
		verbosity.Error(err.Error())
		return errors.New("could not find the role listing channel's.\nDid you run the " + configurechannel.Descriptor.Command.Name + " command?")
	}

	s.ChannelMessageSendComplex(channel, &discordgo.MessageSend{
		Content: fmt.Sprintf("%v: %v", group, description),
		Components: []discordgo.MessageComponent{
			discordgo.ActionsRow{
				Components: []discordgo.MessageComponent{
					discordgo.Button{
						Label:    "Join now!",
						CustomID: "create-group-" + roleID,
					},
				},
			},
		},
	})

	return
}
