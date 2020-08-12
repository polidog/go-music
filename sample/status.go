package main

import (
	"flag"
	"fmt"
	"go-music"
)

func main() {
	var command string
	flag.StringVar(&command, "f", "", "config file name.")
	flag.Parse()

	music, err := music.NewPlayer("spotify")

	if err != nil {
		fmt.Println(err)
		return
	}

	ret, errTrack := music.Track()
	if errTrack != nil {
		fmt.Println(errTrack)
	}

	fmt.Println(ret.Track.Name)

}
