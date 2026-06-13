package main

// TimerState

type TimerState int

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

func (ts TimerState) String() string {
	if stateName, ok := stateName[ts]; ok {
		return stateName
	}
	return "unknown"
}

func (ts TimerState) Seconds() int {
	if stateDuration, ok := stateDuration[ts]; ok {
		return stateDuration * 60
	}
	return -1
}

//Timer

func (t *Timer) Next() {
	currentState := t.state
	switch currentState {
	case StateWork:
		t.state = StateBreak
	case StateBreak:
		t.state = StateWork
	}

}

type Timer struct {
	state     TimerState
	remaining int
}

func main() {

}
