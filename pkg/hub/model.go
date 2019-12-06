package hub

type ObjectModel interface {
	GetAttribute(key string) (string, error)
	Encode(string) (string, error)
	Dencode(string) (string, error)
}
