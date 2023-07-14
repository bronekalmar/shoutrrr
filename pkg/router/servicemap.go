package router

import (
	"github.com/bronekalmar/shoutrrr/pkg/services/bark"
	"github.com/bronekalmar/shoutrrr/pkg/services/discord"
	"github.com/bronekalmar/shoutrrr/pkg/services/generic"
	"github.com/bronekalmar/shoutrrr/pkg/services/googlechat"
	"github.com/bronekalmar/shoutrrr/pkg/services/gotify"
	"github.com/bronekalmar/shoutrrr/pkg/services/ifttt"
	"github.com/bronekalmar/shoutrrr/pkg/services/join"
	"github.com/bronekalmar/shoutrrr/pkg/services/logger"
	"github.com/bronekalmar/shoutrrr/pkg/services/matrix"
	"github.com/bronekalmar/shoutrrr/pkg/services/mattermost"
	"github.com/bronekalmar/shoutrrr/pkg/services/ntfy"
	"github.com/bronekalmar/shoutrrr/pkg/services/opsgenie"
	"github.com/bronekalmar/shoutrrr/pkg/services/pushbullet"
	"github.com/bronekalmar/shoutrrr/pkg/services/pushover"
	"github.com/bronekalmar/shoutrrr/pkg/services/rocketchat"
	"github.com/bronekalmar/shoutrrr/pkg/services/slack"
	"github.com/bronekalmar/shoutrrr/pkg/services/smtp"
	"github.com/bronekalmar/shoutrrr/pkg/services/teams"
	"github.com/bronekalmar/shoutrrr/pkg/services/telegram"
	"github.com/bronekalmar/shoutrrr/pkg/services/zulip"
	t "github.com/bronekalmar/shoutrrr/pkg/types"
)

var serviceMap = map[string]func() t.Service{
	"bark":       func() t.Service { return &bark.Service{} },
	"discord":    func() t.Service { return &discord.Service{} },
	"generic":    func() t.Service { return &generic.Service{} },
	"gotify":     func() t.Service { return &gotify.Service{} },
	"googlechat": func() t.Service { return &googlechat.Service{} },
	"hangouts":   func() t.Service { return &googlechat.Service{} },
	"ifttt":      func() t.Service { return &ifttt.Service{} },
	"join":       func() t.Service { return &join.Service{} },
	"logger":     func() t.Service { return &logger.Service{} },
	"matrix":     func() t.Service { return &matrix.Service{} },
	"mattermost": func() t.Service { return &mattermost.Service{} },
	"ntfy":       func() t.Service { return &ntfy.Service{} },
	"opsgenie":   func() t.Service { return &opsgenie.Service{} },
	"pushbullet": func() t.Service { return &pushbullet.Service{} },
	"pushover":   func() t.Service { return &pushover.Service{} },
	"rocketchat": func() t.Service { return &rocketchat.Service{} },
	"slack":      func() t.Service { return &slack.Service{} },
	"smtp":       func() t.Service { return &smtp.Service{} },
	"teams":      func() t.Service { return &teams.Service{} },
	"telegram":   func() t.Service { return &telegram.Service{} },
	"zulip":      func() t.Service { return &zulip.Service{} },
}
