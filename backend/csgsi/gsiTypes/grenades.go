package gsiTypes

import "encoding/json"

/**
looks like this:
{
	"<player that threw grenade>": {"type": "<grenadetype>", ...rest of grenade data}
		"<player that threw grenade>": {"type": "<grenadetype>", ...rest of grenade data}
}
*/

type Grenades map[string]GrenadeType

func (g *Grenades) UnmarshalJSON(data []byte) error {
	var m map[string]json.RawMessage
	if err := json.Unmarshal(data, &m); err != nil {
		return err
	}

	*g = make(map[string]GrenadeType)
	for k, v := range m {
		var grenadeType struct {
			Type string `json:"type"`
		}
		if err := json.Unmarshal(v, &grenadeType); err != nil {
			return err
		}

		var grenade GrenadeType
		switch grenadeType.Type {
		case "decoy", "smoke":
			grenade = &DecoySmokeGrenade{}
		case "firebomb":
			grenade = &FireBombGrenade{}
		default:
			grenade = &DefaultGrenade{}
		}

		if err := json.Unmarshal(v, grenade); err != nil {
			return err
		}

		(*g)[k] = grenade
	}

	return nil

}

type GrenadeType interface {
	Type() string
	Owner() string
}

type DecoySmokeGrenade struct {
	TypeVal  string `json:"type"`
	OwnerVal string `json:"owner"`

	Position   string `json:"position"`
	Velocity   string `json:"velocity"`
	Lifetime   string `json:"lifetime"`
	EffectTime string `json:"effecttime"`
}

func (d DecoySmokeGrenade) Type() string  { return d.TypeVal }
func (d DecoySmokeGrenade) Owner() string { return d.OwnerVal }

type DefaultGrenade struct {
	TypeVal  string `json:"type"`
	OwnerVal string `json:"owner"`
	Position string `json:"position"`
	Velocity string `json:"velocity"`
	Lifetime string `json:"lifetime"`
}

func (d DefaultGrenade) Type() string  { return d.TypeVal }
func (d DefaultGrenade) Owner() string { return d.OwnerVal }

type FireBombGrenade struct {
	TypeVal  string `json:"type"`
	OwnerVal string `json:"owner"`
	Lifetime string `json:"lifetime"`
	Flames   map[string]*string
}

func (f FireBombGrenade) Type() string  { return f.TypeVal }
func (f FireBombGrenade) Owner() string { return f.OwnerVal }
