package main

import (
	"MusicEntry"
	"bufio"
	"fmt"
	"os"
	"playMusic"
	"strconv"
	"strings"
)

var lib MusicEntry.MusicManager
var id int = 1
var crtl, signal chan int

func handleLibCommands(tokens []string) {
	switch tokens[1] {
	case "list":
		for i := 0; i < lib.Len(); i++ {
			e, _ := lib.Get(i)
			fmt.Println(i+1, ":", e.Name, e.Artist, e.Source, e.Type)
		}
	case "add":
		if len(tokens) == 6 {
			id++
			lib.Add(&MusicEntry.MusicEntry{Id: strconv.Itoa(id), Name: tokens[2], Artist: tokens[3], Source: tokens[4], Type: tokens[5]})
		} else {
			fmt.Println("USAGE: lib add <name><artist><source><type>")
		}
	case "remove":
		if len(tokens) == 3 {
			lib.RemoveByName(tokens[2])
		} else {
			fmt.Println("USAGE: lib remove <name>")
		}
	default:
		fmt.Println("Unrecognized lib command: ", tokens[1])
	}
}

func handlerPlayCommand(tokens []string) {
	if len(tokens) != 2 {
		fmt.Println("Usage play <name>")
		return
	}

	e := lib.Find(tokens[1])
	if e != nil {
		fmt.Println("The music", tokens[1], "does not exist.")
		return
	}
	playMusic.Play(e.Source, e.Type)
}

func main() {
	fmt.Println(`
	Enter follwing command to control player:
lib list <name><artist><source><type> -- Add a music to the music lib
lib remove <name> -- Remove the specifed music from the lib
play <name> -- Play the specify music
`)

	lib = *MusicEntry.NewMusicManager()
	r := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("Enter command ->")

		rawLine,_,_ := r.ReadLine()
		line := string(rawLine)

		if line == "q" || line == "e" {
			break
		}
		tokens := strings.Split(line, " ")

		if tokens[0] == "lib" {
			handleLibCommands(tokens)
		}else if tokens[0] == "play" {
			handlerPlayCommand(tokens)
		} else {
			fmt.Println("Unrecognized command: ", tokens[0])
		}
	}

}
