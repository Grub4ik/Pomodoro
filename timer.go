package main

import (
	"bufio"
	"os"
	"strings"
	"time"

	"github.com/gen2brain/beeep"
)

func (ts TimerState) String() string {
	if name, ok := stateName[ts]; ok {
		return name
	}
	return "unknown"
}

func (ts TimerState) Seconds() int {
	if dur, ok := stateDuration[ts]; ok {
		return dur * 60
	}
	return -1
}

func (t *Timer) Next() {

	switch t.state {
	case StateWork:
		if t.workCount == 3 {
			beeep.Beep(100, 500)
			t.state = StateLongBreak
		} else {
			beeep.Beep(400, 500)
			t.state = StateBreak
		}
	case StateBreak:
		beeep.Beep(800, 500)
		t.state = StateWork
		t.workCount++
	case StateLongBreak:
		beeep.Beep(800, 500)
		t.state = StateWork
		t.workCount = 0
	}
}

func (t *Timer) Start() {
	if t.running {
		return
	}
	t.running = true
}

func (t *Timer) Pause() {
	t.running = false
}

func (t *Timer) Reset() {
	t.state = StateWork
	t.remaining = t.state.Seconds()
	t.workCount = 0
	t.running = true
}

func (t *Timer) Skip() {
	if t.state == StateIdle {
		return
	}
	t.Next()
	t.remaining = t.state.Seconds()
}

func (t *Timer) getUpdate() Update {
	return Update{
		State:     t.state,
		Remaining: t.remaining,
		WorkCount: t.workCount,
		Running:   t.running,
	}
}

func (t *Timer) Run(commandCh <-chan Command, updateCh chan<- Update) {
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			if t.running && t.state != StateIdle && t.state != StatePaused {
				t.remaining--
				if t.remaining <= 0 {
					t.Next()
					t.remaining = t.state.Seconds()
				}
				updateCh <- t.getUpdate()
			}

		case cmd := <-commandCh:
			if cmd == CmdQuit {
				close(updateCh)
				return
			}
			ExecuteCommand(t, cmd)
			updateCh <- t.getUpdate()
		}
	}
}

func inputLoop(commandCh chan<- Command) {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		text := strings.ToLower(scanner.Text())
		cmd, ok := commandDefenition[text]
		if !ok {
			continue
		}
		commandCh <- cmd
		if cmd == CmdQuit {
			close(commandCh)
			return
		}
	}
}

func ExecuteCommand(t *Timer, cmd Command) {
	switch cmd {
	case CmdStart:
		t.Start()
	case CmdPause:
		t.Pause()
	case CmdReset:
		t.Reset()
	case CmdSkip:
		t.Skip()
	case CmdQuit:
	}
}

func (cn Command) String() string {
	if name, ok := commandName[cn]; ok {
		return name
	}
	return "unknown"
}
