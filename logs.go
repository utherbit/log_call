package log_call

import (
	"fmt"
	"log"
)

var LogCall = logCallHandlerGroups{state: make(map[string]*logCallHandler)}

type logCallHandlerGroups struct {
	state map[string]*logCallHandler
}

type logCallHandler struct {
	group string
	state bool
}

func (b logCallHandler) Log(i interface{}) {
	fmt.Printf("Log, b %#v", b.state)

	if b.state {
		log.Println(b.group, ": g", i)
	}
}

func (g *logCallHandlerGroups) SetConfig(config map[string]bool) {
	for s, b := range config {
		g.state[s] = &logCallHandler{
			group: s,
			state: b,
		}
	}
}
func (g logCallHandlerGroups) AddGroup(group string, state ...bool) *logCallHandler {
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
