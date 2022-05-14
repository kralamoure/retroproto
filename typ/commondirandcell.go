package typ

import (
	"fmt"
	"strings"

	"github.com/kralamoure/retroproto"
)

type CommonDirAndCell struct {
	DirId  int
	CellId int
}

func NewCommonDirAndCell(extra string) (CommonDirAndCell, error) {
	var m CommonDirAndCell

	if err := m.Deserialize(extra); err != nil {
		return CommonDirAndCell{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m CommonDirAndCell) Serialized() (string, error) {
	n1 := m.DirId & 7
	n2 := (m.CellId & 4032) >> 6
	n3 := m.CellId & 63

	ch1, err := retroproto.Encode64(n1)
	if err != nil {
		return "", err
	}
	ch2, err := retroproto.Encode64(n2)
	if err != nil {
		return "", err
	}
	ch3, err := retroproto.Encode64(n3)
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

	dirId, err := retroproto.Decode64(ch1)
	if err != nil {
		return err
	}
	m.DirId = dirId

	n2, err := retroproto.Decode64(ch2)
	if err != nil {
		return err
	}
	n3, err := retroproto.Decode64(ch3)
	if err != nil {
		return err
	}

	cellId := (n2&15)<<6 | n3
	m.CellId = cellId

	return nil
}
