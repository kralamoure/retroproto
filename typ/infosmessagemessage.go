package typ

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/kralamoure/retroproto"
)

type InfosMessageMessage struct {
	Id   int
	Args []string
}

func NewInfosMessageMessage(extra string) (InfosMessageMessage, error) {
	var m InfosMessageMessage

	if err := m.Deserialize(extra); err != nil {
		return InfosMessageMessage{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m InfosMessageMessage) Serialized() (string, error) {
	return fmt.Sprintf("%d;%s", m.Id, strings.Join(m.Args, "~")), nil
}

func (m *InfosMessageMessage) Deserialize(extra string) error {
	sli := strings.Split(extra, ";")
	if len(sli) < 2 {
		return retroproto.ErrInvalidMsg
	}

	if sli[0] != "" {
		id, err := strconv.ParseInt(sli[0], 10, 32)
		if err != nil {
			return err
		}
		m.Id = int(id)
	}

	if sli[1] != "" {
		m.Args = strings.Split(sli[1], "~")
	}

	return nil
}
