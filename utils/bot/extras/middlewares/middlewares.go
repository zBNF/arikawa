package middlewares

import (
	"github.com/zBNF/arikawa/v3/discord"
	"github.com/zBNF/arikawa/v3/utils/bot"
	"github.com/zBNF/arikawa/v3/utils/bot/extras/infer"
)

func AdminOnly(ctx *bot.Context) func(interface{}) error {
	return func(ev interface{}) error {
		var channelID = infer.ChannelID(ev)
		if !channelID.IsValid() {
			return bot.Break
		}

		var userID = infer.UserID(ev)
		if !userID.IsValid() {
			return bot.Break
		}

		p, err := ctx.Permissions(channelID, userID)
		if err == nil && p.Has(discord.PermissionAdministrator) {
			return nil
		}

		return bot.Break
	}
}

func GuildOnly(ctx *bot.Context) func(interface{}) error {
	return func(ev interface{}) error {
		// Try and infer the GuildID.
		if guildID := infer.GuildID(ev); guildID.IsValid() {
			return nil
		}

		var channelID = infer.ChannelID(ev)
		if !channelID.IsValid() {
			return bot.Break
		}

		c, err := ctx.Channel(channelID)
		if err != nil || !c.GuildID.IsValid() {
			return bot.Break
		}

		return nil
	}
}
