package d1encoding

type Typ interface {
	Serialized() (extra string, err error)
}
