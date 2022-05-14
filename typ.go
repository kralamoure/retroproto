package retroproto

type Typ interface {
	Serialized() (extra string, err error)
	Deserialize(extra string) error
}
