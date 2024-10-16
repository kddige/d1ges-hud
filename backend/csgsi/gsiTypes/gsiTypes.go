package gsiTypes

type GameState struct {
	Provider        *Provider   `json:"provider"`
	Auth            *Auth       `json:"auth"`
	Player          *Player     `json:"player"`
	AllPlayers      *AllPlayers `json:"allplayers"`
	Round           *Round      `json:"round"`
	PhaseCountDowns *Phase      `json:"phase_countdowns"`
	Grenades        *Grenades   `json:"grenades"`
	Previously      *Previously `json:"previously"`
	// Added           *AddedPlaying `json:"added"`
	// Bomb            *Bomb         `json:"bomb"`
	Map *Map `json:"map"`
}

// type previously, is the same as gamestate, without the previously field
type Previously struct {
	AllPlayers      *AllPlayers `json:"allplayers"`
	PhaseCountDowns *Phase      `json:"phase_countdowns"`
}

type Provider struct {
	Name      *string `json:"name"`
	AppID     *int    `json:"appid"`
	Version   *int    `json:"version"`
	SteamID   *string `json:"steamid"`
	Timestamp *int64  `json:"timestamp"`
}

type Auth struct {
	Token *string `json:"token"`
}

type Player struct {
	Clan         *string      `json:"clan"`
	SteamID      *string      `json:"steamid"`
	Name         *string      `json:"name"`
	ObserverSlot *int         `json:"observer_slot"`
	Team         *string      `json:"team"`
	Activity     *string      `json:"activity"`
	State        *PlayerState `json:"state"`
	Position     *string      `json:"position"`
	Forward      *string      `json:"forward"`
	Spectarget   *string      `json:"spectarget"`
}

type PlayerState struct {
	Health        *int     `json:"health"`
	Armor         *int     `json:"armor"`
	Helmet        *bool    `json:"helmet"`
	Flashed       *float32 `json:"flashed"`
	Smoked        *float32 `json:"smoked"`
	Burning       *float32 `json:"burning"`
	Money         *int     `json:"money"`
	RoundKills    *int     `json:"round_kills"`
	RoundKillhs   *int     `json:"round_killhs"`
	RoundTotalDmg *int     `json:"round_totaldmg"`
	EquipValue    *int     `json:"equip_value"`
	Defusekit     *bool    `json:"defusekit"`
}

type AllPlayers map[string]*PlayerList

type PlayerList struct {
	Clan         *string      `json:"clan"`
	Name         *string      `json:"name"`
	ObserverSlot *int         `json:"observer_slot"`
	Team         *string      `json:"team"`
	State        *PlayerState `json:"state"`
	MatchStats   *PlayerStats `json:"match_stats"`
	Weapons      *Weapons     `json:"weapons"` // TODO: implement
	Position     *string      `json:"position"`
	Forward      *string      `json:"forward"`
}

type PlayerStats struct {
	Kills   *int `json:"kills"`
	Assists *int `json:"assists"`
	Deaths  *int `json:"deaths"`
	Mvps    *int `json:"mvps"`
	Score   *int `json:"score"`
}

type Round struct {
	Phase   *string `json:"phase"`
	Bomb    *string `json:"bomb"`
	WinTeam *string `json:"win_team"`
}

type Phase struct {
	Phase       *string `json:"phase"`
	PhaseEndsIn *string `json:"phase_ends_in"`
}

type Map struct {
	Mode                  *string `json:"mode"`
	Name                  *string `json:"name"`
	NumMatchesToWinSeries *int    `json:"num_matches_to_win_series"`
	Phase                 *string `json:"phase"`
	Round                 *int    `json:"round"`
	TeamCT                *Team   `json:"team_ct"`
	TeamT                 *Team   `json:"team_t"`
}
type Team struct {
	ConsecutiveRoundLosses *int `json:"consecutive_round_losses"`
	MatchesWonThisSeries   *int `json:"matches_won_this_series"`
	Score                  *int `json:"score"`
	TimeoutsRemaining      *int `json:"timeouts_remaining"`
}
