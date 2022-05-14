package msgcli

import (
	"fmt"
	"strconv"

	"github.com/kralamoure/retroproto"
)

type AccountSetServer struct {
	Id int
}

func (m AccountSetServer) MessageId() retroproto.MsgCliId {
	return retroproto.AccountSetServer
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
