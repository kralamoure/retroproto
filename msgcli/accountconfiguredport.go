package msgcli

import (
	"fmt"
	"strconv"

	"github.com/kralamoure/retroproto"
)

type AccountConfiguredPort struct {
	Port int
}

func (m AccountConfiguredPort) MessageId() retroproto.MsgCliId {
	return retroproto.AccountConfiguredPort
}

func (m AccountConfiguredPort) MessageName() string {
	return "AccountConfiguredPort"
}

func (m AccountConfiguredPort) Serialized() (string, error) {
	return fmt.Sprint(m.Port), nil
}

func (m *AccountConfiguredPort) Deserialize(extra string) error {
	port, err := strconv.Atoi(extra)
	if err != nil {
		return err
	}
	m.Port = port

	return nil
}
