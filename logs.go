package log_call

import "log"

var LogCall LogCallHandlerGroups

type LogCallHandlerGroups struct {
	state map[string]*logCallHandler
}

type logCallHandler struct {
	group string
	state bool
}

func (b logCallHandler) Log(i interface{}) {
	if b.state {
		log.Println(i)
	}
}

func InitLogCall(config ...map[string]bool) {
	LogCall = LogCallHandlerGroups{
		state: map[string]*logCallHandler{},
	}
	if len(config) > 0 {
		for s, b := range config[0] {
			LogCall.state[s] = &logCallHandler{
				group: s,
				state: b,
			}
		}
	}

}
func (g *LogCallHandlerGroups) SetConfig(config map[string]bool) {
	state := map[string]*logCallHandler{}
	for s, b := range config {
		state[s] = &logCallHandler{
			group: s,
			state: b,
		}
	}
}
func (g LogCallHandlerGroups) AddGroup(group string, state ...bool) *logCallHandler {
	if h, ok := g.state[group]; ok {
		if len(state) > 0 {
			h.state = state[0]
		}
		return h
	}

	h := &logCallHandler{
		group: group,
	}
	if len(state) > 0 {
		h.state = state[0]
	}
	g.state[group] = h
	return h
}
