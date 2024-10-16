package gsiTypes

import "encoding/json"

type Weapons map[string]DefaultWeapon

func (w *Weapons) UnmarshalJSON(data []byte) error {
	var m map[string]json.RawMessage
	if err := json.Unmarshal(data, &m); err != nil {
		return err
	}

	*w = make(map[string]DefaultWeapon)
	for k, v := range m {
		var weaponType struct {
			Type string `json:"type"`
		}
		if err := json.Unmarshal(v, &weaponType); err != nil {
			return err
		}

		var weapon DefaultWeapon
		switch weaponType.Type {
		case "":
			weapon = &WeaponUnknown{} // error here
		case string(WeaponTypeRifle):
			weapon = &WeaponRifle{}

		case string(WeaponTypeSniperRifle):
			weapon = &WeaponSniperRifle{}

		case string(WeaponTypeSubmachineGun):
			weapon = &WeaponSubmachineGun{}

		case string(WeaponTypeShotgun):
			weapon = &WeaponShotgun{}

		case string(WeaponTypeMachineGun):
			weapon = &WeaponMachineGun{}

		case string(WeaponTypePistol):
			weapon = &WeaponPistol{}

		case string(WeaponTypeKnife):
			weapon = &WeaponKnife{}

		case string(WeaponTypeGrenade):
			weapon = &WeaponGrenade{}

		case string(WeaponTypeC4):
			weapon = &WeaponC4{}

		default:
			weapon = &WeaponUnknown{} // error here
		}

		if err := json.Unmarshal(v, weapon); err != nil {
			return err
		}

		(*w)[k] = weapon
	}

	return nil
}

type DefaultWeapon interface {
	IsDefaultWeapon() bool
}

type WeaponType string

const (
	WeaponTypeUnknown       WeaponType = ""
	WeaponTypeRifle         WeaponType = "Rifle"
	WeaponTypeSniperRifle   WeaponType = "SniperRifle"
	WeaponTypeSubmachineGun WeaponType = "SubmachineGun"
	WeaponTypeShotgun       WeaponType = "Shotgun"
	WeaponTypeMachineGun    WeaponType = "MachineGun"
	WeaponTypePistol        WeaponType = "Pistol"
	WeaponTypeKnife         WeaponType = "Knife"
	WeaponTypeGrenade       WeaponType = "Grenade"
	WeaponTypeC4            WeaponType = "C4"
	WeaponStackableItem     WeaponType = "StackableItem"
	WeaponTable             WeaponType = "Table"
	WeaponFists             WeaponType = "Fists"
	WeaponBreachCharge      WeaponType = "BreachCharge"
	WeponMelee              WeaponType = "Melee"
)

type WeaponState string

const (
	WeaponStateUnknown   WeaponState = ""
	WeaponStateActive    WeaponState = "active"
	WeaponStateHolstered WeaponState = "holstered"
	WeaponStateReloading WeaponState = "reloading"
)

// Weapons
type WeaponUnknown struct {
	Type    string `json:"type"`
	NameVal string `json:"name"`
}

func (w WeaponUnknown) IsDefaultWeapon() bool { return true }

type WeaponRifle struct {
	Type        string       `json:"type"`
	NameVal     string       `json:"name"`
	PaintKit    string       `json:"paintkit"`
	AmmoClip    int          `json:"ammo_clip"`
	AmmoClipMax int          `json:"ammo_clip_max"`
	AmmoReserve int          `json:"ammo_reserve"`
	State       *WeaponState `json:"state"`
}

func (w WeaponRifle) IsDefaultWeapon() bool { return true }

type WeaponSniperRifle struct {
	Type        string       `json:"type"`
	NameVal     string       `json:"name"`
	PaintKit    string       `json:"paintkit"`
	AmmoClip    int          `json:"ammo_clip"`
	AmmoClipMax int          `json:"ammo_clip_max"`
	AmmoReserve int          `json:"ammo_reserve"`
	State       *WeaponState `json:"state"`
}

func (w WeaponSniperRifle) IsDefaultWeapon() bool { return true }

type WeaponSubmachineGun struct {
	Type        string       `json:"type"`
	NameVal     string       `json:"name"`
	PaintKit    string       `json:"paintkit"`
	AmmoClip    int          `json:"ammo_clip"`
	AmmoClipMax int          `json:"ammo_clip_max"`
	AmmoReserve int          `json:"ammo_reserve"`
	State       *WeaponState `json:"state"`
}

func (w WeaponSubmachineGun) IsDefaultWeapon() bool { return true }

type WeaponShotgun struct {
	Type        string       `json:"type"`
	NameVal     string       `json:"name"`
	PaintKit    string       `json:"paintkit"`
	AmmoClip    int          `json:"ammo_clip"`
	AmmoClipMax int          `json:"ammo_clip_max"`
	AmmoReserve int          `json:"ammo_reserve"`
	State       *WeaponState `json:"state"`
}

func (w WeaponShotgun) IsDefaultWeapon() bool { return true }

type WeaponMachineGun struct {
	Type        string       `json:"type"`
	NameVal     string       `json:"name"`
	PaintKit    string       `json:"paintkit"`
	AmmoClip    int          `json:"ammo_clip"`
	AmmoClipMax int          `json:"ammo_clip_max"`
	AmmoReserve int          `json:"ammo_reserve"`
	State       *WeaponState `json:"state"`
}

func (w WeaponMachineGun) IsDefaultWeapon() bool { return true }

type WeaponPistol struct {
	Type        string       `json:"type"`
	NameVal     string       `json:"name"`
	PaintKit    string       `json:"paintkit"`
	AmmoClip    int          `json:"ammo_clip"`
	AmmoClipMax int          `json:"ammo_clip_max"`
	AmmoReserve int          `json:"ammo_reserve"`
	State       *WeaponState `json:"state"`
}

func (w WeaponPistol) IsDefaultWeapon() bool { return true }

type WeaponKnife struct {
	Type     string       `json:"type"`
	NameVal  string       `json:"name"`
	PaintKit string       `json:"paintkit"`
	State    *WeaponState `json:"state"`
}

func (w WeaponKnife) IsDefaultWeapon() bool { return true }

type WeaponGrenade struct {
	Type        string       `json:"type"`
	NameVal     string       `json:"name"`
	PaintKit    string       `json:"paintkit"`
	AmmoReserve int          `json:"ammo_reserve"`
	State       *WeaponState `json:"state"`
}

func (w WeaponGrenade) IsDefaultWeapon() bool { return true }

type WeaponC4 struct {
	Type     string       `json:"type"`
	NameVal  string       `json:"name"`
	PaintKit string       `json:"paintkit"`
	State    *WeaponState `json:"state"`
}

func (w WeaponC4) IsDefaultWeapon() bool { return true }
