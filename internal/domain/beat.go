package domain

type Beat struct {
	ID        string
	Title     string
	Producer  string
	Stream    string
	Rithm     string
	CreatedAt string
	IDType    string
	Type      *BeatType
}
