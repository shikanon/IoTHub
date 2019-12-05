package hub

type ObjectModel interface {
	GetAttribute(key string) (string, error)
	Encode(string) (string, error)
	Dencdoe(string) (string, error)
}
