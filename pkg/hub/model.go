package hub

type ObjectModel interface {
	GetAttribute(string) (string, error)
	Encode(string) (string, error)
	Dencode(string) (string, error)
}
