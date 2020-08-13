package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/polidog/go-music"
)

func main() {
	flag.Parse()
	args := flag.Args()
	os.Exit(run(args))
}

func run(args []string) int {
	if len(args) < 2 {
		fmt.Println("argument error")
		return 1
	}

	player, err := music.NewPlayer(args[0])
	if err != nil {
		fmt.Println(err)
		return 1
	}

	res, errExec := player.Exec(args[1])

	if errExec != nil {
		fmt.Println(err)
		return 1
	}

	fmt.Println("player_state:", res.PlayerState)

	fmt.Println("----------track-----------")
	fmt.Println("Name:", res.Track.Name)
	fmt.Println("Album:", res.Track.Album)
	fmt.Println("Artist:", res.Track.Artist)
	return 0
}
