package csgsi

import (
	"changeme/backend/csgsi/gsiTypes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"math"
	"net/http"
	"reflect"
	"strconv"
	"time"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

type GameEvents struct {
	RoundOver     chan bool            // Emits: true when round is over
	GameMap       chan string          // Emits: Current map
	GamePhase     chan string          // Emits: Current game phase
	GameRounds    chan int             // Emits: Current round / phase
	GameCTScore   chan int             // Emits: CT Score
	GameTScore    chan int             // Emits: T Score
	Player        chan gsiTypes.Player // Emits: Current spectated player
	RoundPhase    chan string          // Emits: Current round phase
	RoundWinTeam  chan string          // Emits: The winning team, if round is over
	BombPlanted   chan bool            // Emits: true when bomb is planted
	BombDefused   chan bool            // Emits: true when bomb is defused
	BombExploded  chan bool            // Emits: true when bomb is exploded
	BombState     chan string          // Emits: The bomb state
	BombCountDown chan int             // Emits: The bomb countdown
}

type TrackedState struct {
	IsBombPlanted bool          // If the bomb is planted
	BombTickQuit  chan struct{} // Bomb ticker quit channel
}

type Game struct {
	// Channel for game state data.
	ctx          context.Context
	Channel      chan gsiTypes.GameState
	RawChannel   chan interface{}
	Events       GameEvents
	TrackedState *TrackedState
}

// Returns a new Game object.
func New(ctx context.Context) *Game {
	channelSize := 100
	game := &Game{
		Channel:    make(chan gsiTypes.GameState, channelSize),
		RawChannel: make(chan interface{}, channelSize),
		Events: GameEvents{
			RoundOver:     make(chan bool, channelSize),
			GameMap:       make(chan string, channelSize),
			GamePhase:     make(chan string, channelSize),
			GameRounds:    make(chan int, channelSize),
			GameCTScore:   make(chan int, channelSize),
			GameTScore:    make(chan int, channelSize),
			Player:        make(chan gsiTypes.Player, channelSize),
			RoundPhase:    make(chan string, channelSize),
			RoundWinTeam:  make(chan string, channelSize),
			BombPlanted:   make(chan bool, channelSize),
			BombDefused:   make(chan bool, channelSize),
			BombExploded:  make(chan bool, channelSize),
			BombState:     make(chan string, channelSize),
			BombCountDown: make(chan int, channelSize),
		},
		TrackedState: &TrackedState{
			IsBombPlanted: false,
		},
	}
	game.ctx = ctx

	game.StartEventHandlers()
	// game.DebugListenAllChannels()

	return game
}

// Starts listening to address provided.
// If a POST request is received, it sends it through
// the Game.Channel channel.
func (gs *Game) Listen(ctx context.Context, addr string) error {
	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		if req.Method == "POST" {

			bodyBytes, err := ioutil.ReadAll(req.Body)
			if err != nil {
				io.WriteString(res, "bad")
				return
			}

			// Decode into the rawState variable
			var rawState interface{}
			if err := json.Unmarshal(bodyBytes, &rawState); err != nil {
				io.WriteString(res, "bad")
				return
			}

			sendDropOld(&gs.RawChannel, rawState)
			state := &gsiTypes.GameState{}
			if err := json.Unmarshal(bodyBytes, &state); err != nil {
				stringError := fmt.Sprintf("%s", err)
				runtime.LogError(ctx, stringError)
				io.WriteString(res, stringError)
				return
			}

			sendDropOld(&gs.Channel, *state)

			io.WriteString(res, "ok")
		}
	})

	if err := http.ListenAndServe(addr, nil); err != nil {
		return err
	}

	return nil
}

func (gs *Game) DebugListenAllChannels() {

	// iterates over all channels and prints the value
	v := reflect.ValueOf(gs.Events)
	t := v.Type()

	for i := 0; i < v.NumField(); i++ {
		fieldName := t.Field(i).Name
		field := v.Field(i)

		if field.Kind() == reflect.Chan {
			go func(name string, ch reflect.Value) {
				for {
					value, ok := ch.Recv()
					if !ok {
						break
					}
					runtime.LogDebug(gs.ctx, fmt.Sprintf("%s: %v", name, value))
				}
			}(fieldName, field)
		}
	}

}

func (gs *Game) StartEventHandlers() {
	go func() {
		for state := range gs.Channel {

			// Map events
			if state.Map != nil {
				sendDropOld(&gs.Events.GameMap, *state.Map.Name)
				sendDropOld(&gs.Events.GamePhase, *state.Map.Phase)
				sendDropOld(&gs.Events.GameRounds, *state.Map.Round)
				sendDropOld(&gs.Events.GameCTScore, *state.Map.TeamCT.Score)
				sendDropOld(&gs.Events.GameTScore, *state.Map.TeamT.Score)
			}

			// Player events
			if state.Player != nil {
				sendDropOld(&gs.Events.Player, *state.Player)
			}

			// Round events
			if state.Round != nil {
				sendDropOld(&gs.Events.RoundPhase, *state.Round.Phase)

				if state.Round.Bomb != nil {
					sendDropOld(&gs.Events.BombState, *state.Round.Bomb)

					switch *state.Round.Bomb {
					case "planted":
						if !gs.TrackedState.IsBombPlanted && state.PhaseCountDowns != nil && *state.PhaseCountDowns.Phase == "bomb" {
							phaseEndsIn, err := strconv.ParseFloat(*state.PhaseCountDowns.PhaseEndsIn, 64)
							if err != nil {
								runtime.LogError(gs.ctx, "Error converting timeLeft to int")
							}
							timeLeft := 40 - (40 - int(math.Round(phaseEndsIn)))
							gs.StartBombTimer(timeLeft)
							gs.TrackedState.IsBombPlanted = true
							sendDropOld(&gs.Events.BombPlanted, true)
						}
					case "defused":
						sendDropOld(&gs.Events.BombDefused, true)
						if gs.TrackedState.IsBombPlanted == true {
							gs.StopBombTimer()
							gs.TrackedState.IsBombPlanted = false
						}
					case "exploded":
						sendDropOld(&gs.Events.BombExploded, true)
						if gs.TrackedState.IsBombPlanted == true {
							gs.StopBombTimer()
							gs.TrackedState.IsBombPlanted = false
						}
					default:
						break
					}
				}

				switch *state.Round.Phase {
				case "live":
					break
				case "freezetime":
					break
				case "over":
					sendDropOld(&gs.Events.RoundWinTeam, *state.Round.WinTeam)
					sendDropOld(&gs.Events.RoundOver, true)
					if gs.TrackedState.IsBombPlanted == true {
						gs.StopBombTimer()
						gs.TrackedState.IsBombPlanted = false
					}
					break
				default:
					break
				}
			}

		}
	}()
}

func (gs *Game) StartBombTimer(timeLeft int) {
	// start a countdown timer for the bomb
	runtime.LogDebug(gs.ctx, "Starting bomb timer...")
	bombCountdown := timeLeft

	ticker := time.NewTicker(time.Second)
	gs.TrackedState.BombTickQuit = make(chan struct{})
	go func() {
		for {
			select {
			case <-ticker.C:
				if bombCountdown > 0 {
					bombCountdown--
					gs.Events.BombCountDown <- bombCountdown
				} else {
					gs.Events.BombCountDown <- 0
					ticker.Stop()
					return
				}
			case <-gs.TrackedState.BombTickQuit:
				ticker.Stop()
				return
			}
		}
	}()
}

func (gs *Game) StopBombTimer() {
	close(gs.TrackedState.BombTickQuit)
}

// Sends a value to a channel, dropping the oldest value if the channel is full.
func sendDropOld[T any](ch *chan T, value T) {
	select {
	case *ch <- value:
	default:
		<-*ch
		*ch <- value
	}
}
