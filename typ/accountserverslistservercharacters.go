package typ

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/kralamoure/d1encoding"
)

type AccountServersListServerCharacters struct {
	Id  int
	Qty int
}

func (m AccountServersListServerCharacters) Serialized() (string, error) {
	return fmt.Sprintf("%d,%d", m.Id, m.Qty), nil
}

func (m *AccountServersListServerCharacters) Deserialize(extra string) error {
	sli := strings.Split(extra, ",")
	if len(sli) < 2 {
		return d1encoding.ErrInvalidMsg
	}

	if sli[0] != "" {
		id, err := strconv.ParseInt(sli[0], 10, 32)
		if err != nil {
			return err
		}
		m.Id = int(id)
	}

	if sli[1] != "" {
		qty, err := strconv.ParseInt(sli[1], 10, 32)
		if err != nil {
			return err
		}
		m.Qty = int(qty)
	}

	return nil
}
