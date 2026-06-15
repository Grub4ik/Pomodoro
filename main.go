package main

func main() {
	// В main
	commandCh := make(chan Command)
	updateCh := make(chan Update)

	// go inputLoop(commandCh)
	// go timer.Run(commandCh, updateCh)

	Timer := &Timer{}
}
