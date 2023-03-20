package log_call

import (
	"fmt"
	"time"
)

var LogCall = LogCallHandlerGroups{state: make(map[string]*LogCallHandler), cfg: cfgDefault()}

type LogCallHandlerGroups struct {
	cfg   Config
	state map[string]*LogCallHandler
}

const (
	ColorRed    = "\033[31m"
	ColorGreen  = "\033[32m"
	ColorYellow = "\033[33m"
	ColorBlue   = "\033[34m"
	ColorPurple = "\033[35m"
	ColorCyan   = "\033[36m"
	ColorWhite  = "\033[37m"

	ColorReset = "\033[0m"
)

func cfgDefault() Config {
	return Config{
		Display: &ConfigDisplay{
			Time:    true,
			Group:   true,
			Message: true,
		},
		ColorsTime:    ColorGreen,
		ColorsGroup:   ColorBlue,
		ColorsMessage: ColorWhite,

		TimeFormat: "_2 Jan 15:04:05",
	}
}

type ConfigDisplay struct {
	Time    bool // Выключение лога для времени
	Group   bool // Выключение лога для группы
	Message bool // Выключение лога для сообщения
}
type Config struct {
	Groups map[string]bool // Выключение логов для перечисленных групп

	Display *ConfigDisplay
	//DisplayTime    bool // Выключение лога для времени
	//DisplayGroup   bool // Выключение лога для группы
	//DisplayMessage bool // Выключение лога для сообщения

	TimeFormat string

	ColorsTime    string // Цвет лога для времени
	ColorsGroup   string // Цвет лога для группы
	ColorsMessage string // Цвет лога для сообщения

}

type LogCallHandler struct {
	root  *LogCallHandlerGroups
	group string
	state bool
}

func (b LogCallHandler) Log(i ...interface{}) {
	if b.state {
		print("\n")
		// 20 Mar 12:50:23
		if b.root.cfg.Display.Time {
			print(b.root.cfg.ColorsTime, time.Now().Format(b.root.cfg.TimeFormat))
		}

		// Authorization:
		if b.root.cfg.Display.Group {
			print(b.root.cfg.ColorsGroup, " ", b.group, ": ")
		}

		// HandlerLogin
		print(b.root.cfg.ColorsMessage)
		fmt.Print(i...)

		// the end, set default color cursor
		print(ColorReset)
	}
}

func (g *LogCallHandlerGroups) SetConfig(config Config) {
	if config.ColorsTime != "" {
		g.cfg.ColorsTime = config.ColorsTime
	}
	if config.ColorsGroup != "" {
		g.cfg.ColorsGroup = config.ColorsGroup
	}
	if config.ColorsMessage != "" {
		g.cfg.ColorsMessage = config.ColorsMessage
	}
	if config.TimeFormat != "" {
		g.cfg.TimeFormat = config.TimeFormat
	}
	if config.Display != nil {
		g.cfg.Display = config.Display
	}

	g.cfg.Groups = config.Groups

	for s, b := range config.Groups {
		g.state[s] = &LogCallHandler{
			group: s,
			state: b,
		}
	}
}
func (g *LogCallHandlerGroups) AddGroup(group string, state ...bool) *LogCallHandler {

	if h, ok := g.state[group]; ok {
		if len(state) > 0 {
			h.state = state[0]
		}
		return h
	}

	h := &LogCallHandler{
		root:  g,
		group: group,
	}
	if len(state) > 0 {
		h.state = state[0]
	} else {
		h.state = true
	}
	g.state[group] = h
	return h
}
