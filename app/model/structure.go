package model

type Structure struct {
	ID       string
	Name     string
	Settings []Setting
}

type Setting struct {
	ID       string
	Name     string
	Tier     Tier
	Datatype string
}

type Tier int

const (
	TierOne Tier = 1 + iota
	TierTwo
	TierThree
)
