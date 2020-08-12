# go-music

Control the golang for Music and Spotify.

## Using

```
package main

import (
	"github.com/polidog/go-music"
	"fmt"
)

func main() {
	res, err := music.Track()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("player_state:", res.PlayerState)

	fmt.Println("----------track-----------")
    fmt.Println("Name:", res.Track.Name)
	fmt.Println("Album:", res.Track.Album)
	fmt.Println("Artist:", res.Track.Artist)
}
```