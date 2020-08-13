# go-music

Control the golang for Music and Spotify.

## Using

```
package main

import (
	"fmt"

	"github.com/polidog/go-music"
)

func main() {

	p, err := player.NewPlayer("spotify")
	if err != nil {
		fmt.Println(err)
		return
	}

	res, errTrack := p.Track()
	if errTrack != nil {
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