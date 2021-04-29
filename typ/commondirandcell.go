package typ

import (
	"fmt"
	"strings"

	"github.com/kralamoure/d1encoding"
)

type CommonDirAndCell struct {
	DirId  int
	CellId int
}

func (m CommonDirAndCell) Serialized() (string, error) {
	n1 := m.DirId & 7
	n2 := (m.CellId & 4032) >> 6
	n3 := m.CellId & 63

	ch1, err := d1encoding.Encode64(n1)
	if err != nil {
		return "", err
	}
	ch2, err := d1encoding.Encode64(n2)
	if err != nil {
		return "", err
	}
	ch3, err := d1encoding.Encode64(n3)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%v%v%v", ch1, ch2, ch3), nil
}

func (m *CommonDirAndCell) Deserialize(extra string) error {
	r := strings.NewReader(extra)

	ch1, _, err := r.ReadRune()
	if err != nil {
		return err
	}
	ch2, _, err := r.ReadRune()
	if err != nil {
		return err
	}
	ch3, _, err := r.ReadRune()
	if err != nil {
		return err
	}

	dirId, err := d1encoding.Decode64(ch1)
	if err != nil {
		return err
	}
	m.DirId = dirId

	n2, err := d1encoding.Decode64(ch2)
	if err != nil {
		return err
	}
	n3, err := d1encoding.Decode64(ch3)
	if err != nil {
		return err
	}

	cellId := (n2&15)<<6 | n3
	m.CellId = cellId

	return nil
}
