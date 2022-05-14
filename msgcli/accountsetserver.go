package msgcli

import (
	"fmt"
	"strconv"

	"github.com/kralamoure/retroproto"
)

type AccountSetServer struct {
	Id int
}

func NewAccountSetServer(extra string) (AccountSetServer, error) {
	var m AccountSetServer

	if err := m.Deserialize(extra); err != nil {
		return AccountSetServer{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m AccountSetServer) MessageId() retroproto.MsgCliId {
	return retroproto.AccountSetServer
}

func (m AccountSetServer) MessageName() string {
	return "AccountSetServer"
}

func (m AccountSetServer) Serialized() (string, error) {
	return fmt.Sprintf("%d", m.Id), nil
}

func (m *AccountSetServer) Deserialize(extra string) error {
	id, err := strconv.ParseInt(extra, 10, 32)
	if err != nil {
		return err
	}
	m.Id = int(id)

	return nil
}
