package d1proto

type Typ interface {
	Serialized() (extra string, err error)
}
