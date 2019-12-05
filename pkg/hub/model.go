package hub

type ObjectModel interface {
	GetAttribute(string) (string, error)
	Encode(string) (string, error)
	Dencdoe(string) (string, error)
}
