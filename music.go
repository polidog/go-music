package music

import (
	"encoding/json"
	"fmt"

	"github.com/polidog/go-music/script"
)

// Track is 再生しているトラック情報
type Track struct {
	Album    string `json:"album"`
	Artist   string `json:"artist"`
	Name     string `json:"name"`
	Category string `json:"category"`
	Time     string `json:"time"`
}

// Result 取得データ
type Result struct {
	Status      bool   `json:"status"`
	PlayerState string `json:"player_state"`
	Track       Track  `json:"track"`
}

// Player is music player
type Player struct {
	script script.Script
}

// NewPlayer is create a Player struct.
func NewPlayer(service string) (Player, error) {
	script, err := script.NewScript(service)
	if err != nil {
		return Player{}, err
	}

	return Player{
		script: script,
	}, nil
}

func (p *Player) Track() (Result, error) {
	return p.exec("track")
}

func (p *Player) Play() (Result, error) {
	return p.exec("play")
}

func (p *Player) Pause() (Result, error) {
	return p.exec("pause")
}

func (p *Player) Next() (Result, error) {
	return p.exec("next")
}

func (p *Player) Previous() (Result, error) {
	return p.exec("previous")
}

func (p *Player) State() (Result, error) {
	return p.exec("state")
}

func (p *Player) IsPlaying() (Result, error) {
	return p.exec("isPlaying")
}

func (p *Player) exec(command string) (Result, error) {
	str, err := p.script.Exec(command)
	result := Result{
		Track: Track{},
	}

	if err != nil {
		return result, err
	}

	jsonErr := json.Unmarshal(str, &result)
	if jsonErr != nil {
		fmt.Println(command)
		fmt.Println(str)
		return result, jsonErr
	}
	return result, nil
}
