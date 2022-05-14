package typ

import (
	"fmt"
	"strconv"
	"strings"
)

type CommonTitle struct {
	Id    int
	Param string
}

func NewCommonTitle(extra string) (CommonTitle, error) {
	var m CommonTitle

	if err := m.Deserialize(extra); err != nil {
		return CommonTitle{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m CommonTitle) Serialized() (string, error) {
	sb := &strings.Builder{}

	if m.Id != 0 {
		sb.WriteString(fmt.Sprintf("%d", m.Id))

		if m.Param != "" {
			sb.WriteString(fmt.Sprintf("*%s", m.Param))
		}
	}

	return sb.String(), nil
}

func (m *CommonTitle) Deserialize(extra string) error {
	sli := strings.Split(extra, "*")

	if sli[0] != "" {
		id, err := strconv.ParseInt(sli[0], 10, 32)
		if err != nil {
			return err
		}
		m.Id = int(id)
	}

	if len(sli) < 2 {
		return nil
	}

	m.Param = sli[1]

	return nil
}
