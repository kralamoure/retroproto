package typ

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/kralamoure/retroproto"
)

type AccountHostsHost struct {
	Id         int
	State      int
	Completion int
	CanLog     bool
}

func NewAccountHostsHost(extra string) (AccountHostsHost, error) {
	var m AccountHostsHost

	if err := m.Deserialize(extra); err != nil {
		return AccountHostsHost{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m AccountHostsHost) Serialized() (string, error) {
	var canLog int
	if m.CanLog {
		canLog = 1
	}

	return fmt.Sprintf("%d;%d;%d;%d", m.Id, m.State, m.Completion, canLog), nil
}

func (m *AccountHostsHost) Deserialize(extra string) error {
	sli := strings.Split(extra, ";")
	if len(sli) < 4 {
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
		state, err := strconv.ParseInt(sli[1], 10, 32)
		if err != nil {
			return err
		}
		m.State = int(state)
	}

	if sli[2] != "" {
		completion, err := strconv.ParseInt(sli[2], 10, 32)
		if err != nil {
			return err
		}
		m.Completion = int(completion)
	}

	if sli[3] != "" {
		canLog, err := strconv.ParseBool(sli[3])
		if err != nil {
			return err
		}
		m.CanLog = canLog
	}

	return nil
}
