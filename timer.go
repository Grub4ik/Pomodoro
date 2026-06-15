package main

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

func (t *Timer) Next() {
	currentState := t.state
	switch currentState {

	case StateWork:
		if t.workCount == 4 {
			t.state = StateLongBreak
		} else {
			t.state = StateBreak
		}

	case StateBreak:
		t.state = StateWork
		t.workCount++

	case StateLongBreak:
		t.state = StateWork
		t.workCount = 0
	}
}

func (t *Timer) Start() {
	if t.running {
		return
	}
	t.state = StateWork
	t.running = true
	t.remaining = t.state.Seconds()

}

func (t *Timer) Run() {

}

func (cn Command) String() string {
	if commandName, ok := commandName[cn]; ok {
		return commandName
	}
	return "unknown"
}
