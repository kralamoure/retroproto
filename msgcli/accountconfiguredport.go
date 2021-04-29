package msgcli

import (
	"fmt"
	"strconv"

	"github.com/kralamoure/d1encoding"
)

type AccountConfiguredPort struct {
	Port int
}

func (m AccountConfiguredPort) ProtocolId() d1encoding.MsgCliId {
	return d1encoding.AccountConfiguredPort
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
