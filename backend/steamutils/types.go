package steamutils

import "encoding/xml"

type ProfileXML struct {
	XMLName          xml.Name `xml:"profile"`
	Text             string   `xml:",chardata"`
	SteamID64        string   `xml:"steamID64"`
	SteamID          string   `xml:"steamID"`
	OnlineState      string   `xml:"onlineState"`
	StateMessage     string   `xml:"stateMessage"`
	PrivacyState     string   `xml:"privacyState"`
	VisibilityState  string   `xml:"visibilityState"`
	AvatarIcon       string   `xml:"avatarIcon"`
	AvatarMedium     string   `xml:"avatarMedium"`
	AvatarFull       string   `xml:"avatarFull"`
	VacBanned        string   `xml:"vacBanned"`
	TradeBanState    string   `xml:"tradeBanState"`
	IsLimitedAccount string   `xml:"isLimitedAccount"`
	CustomURL        string   `xml:"customURL"`
	MemberSince      string   `xml:"memberSince"`
	SteamRating      string   `xml:"steamRating"`
	HoursPlayed2Wk   string   `xml:"hoursPlayed2Wk"`
	Headline         string   `xml:"headline"`
	Location         string   `xml:"location"`
	Realname         string   `xml:"realname"`
	Summary          string   `xml:"summary"`
	Groups           struct {
		Text  string `xml:",chardata"`
		Group []struct {
			Text          string `xml:",chardata"`
			IsPrimary     string `xml:"isPrimary,attr"`
			GroupID64     string `xml:"groupID64"`
			GroupName     string `xml:"groupName"`
			GroupURL      string `xml:"groupURL"`
			Headline      string `xml:"headline"`
			Summary       string `xml:"summary"`
			AvatarIcon    string `xml:"avatarIcon"`
			AvatarMedium  string `xml:"avatarMedium"`
			AvatarFull    string `xml:"avatarFull"`
			MemberCount   string `xml:"memberCount"`
			MembersInChat string `xml:"membersInChat"`
			MembersInGame string `xml:"membersInGame"`
			MembersOnline string `xml:"membersOnline"`
		} `xml:"group"`
	} `xml:"groups"`
}
