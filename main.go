package main

import "fmt"

func main() {
	timer := &Timer{}
	timer.Reset()

	commandCh := make(chan Command)
	updateCh := make(chan Update)

	go inputLoop(commandCh)
	go timer.Run(commandCh, updateCh)

	for update := range updateCh {
		display(update)
	}
}

func display(upd Update) {
	fmt.Print("\033[H\033[2J")
	fmt.Printf("%s\n", upd.State)
	fmt.Printf("%02d:%02d\n", upd.Remaining/60, upd.Remaining%60)
	fmt.Printf("WorkCount: %d/4\n", upd.WorkCount)
	fmt.Printf("Running: %v\n", upd.Running)
	fmt.Println("\nCommands: s-start p-pause r-reset k-skip q-quit")
}
