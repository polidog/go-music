class Player {
    constructor(app) {
        this.app = app
    }

    track() {
        return {
            'track': this._track(),
            'player_state': this.app.playerState()
        }
    }

    _track() {
        const t = this.app.currentTrack
        return {
            name: t.name(),
            artist: t.artist(),
            album: t.album()
        }
    }

    play() {
        if (!this.isPlaying()) {
            this.app.play();
        }
        return {
            'track': this._track(),
            'player_state': 'playing'
        }
    }

    start() {
        return this.play()
    }

    pause() {
        if (this.isPlaying()) {
            this.app.pause();
        }
        return {
            'track': this._track(),
            'player_state': 'paused',
        }
    }

    stop() {
        return this.pause()
    }

    async next() {
        this.app.nextTrack();
        return {
            'track': this._track(),
            'player_state': this.app.playerState(),
        }
    }

    previous() {
        this.app.previousTrack();
        return {
            'track': this._track(),
            'player_state': this.app.playerState(),
        }
    }

    state() {
        return {
            'player_state': this.app.playerState()
        }
    }

    isPlaying() {
        return this.app.playerState() === 'playing'
    }
}


function run(argv) {
    const service = argv[0].substring(0, 1).toUpperCase() + argv[0].substring(1);
    const command = argv[1]
    const app = Application(service)
    const player = new Player(app)

    if (player[command] !== undefined) {
        return JSON.stringify(player[command]());
    }
    return JSON.stringify({"message": "argument error", "arg": argv})
}