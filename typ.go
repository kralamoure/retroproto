package d1proto

type Typ interface {
	Serialized() (string, error)
	Deserialize(string) error
}
