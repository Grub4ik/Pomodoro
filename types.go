package main

type TimerState int

type Timer struct {
	state     TimerState
	remaining int
	workCount int
	running   bool
}

const (
	StateWork TimerState = iota
	StateBreak
	StateLongBreak
	StatePaused
	StateIdle
)

var stateName = map[TimerState]string{
	StateWork:      "work",
	StateBreak:     "break",
	StateLongBreak: "long break",
	StatePaused:    "paused",
	StateIdle:      "idle",
}

var stateDuration = map[TimerState]int{
	StateWork:      25,
	StateBreak:     5,
	StateLongBreak: 20,
}

type Command int

const (
	CmdStart Command = iota
	CmdPause
	CmdReset
	CmdSkip
	CmdQuit
)

var commandName = map[Command]string{
	CmdStart: "Start",
	CmdPause: "Pause",
	CmdReset: "Reset",
	CmdSkip:  "Skip",
	CmdQuit:  "Quit",
}

var commandDefenition = map[string]Command{
	"s": CmdStart,
	"p": CmdPause,
	"r": CmdReset,
	"k": CmdSkip,
	"q": CmdQuit,
}

type Update struct {
	State     TimerState
	Remaining int
	WorkCount int
	Running   bool
}
