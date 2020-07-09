package msgcli

import (
	"fmt"
	"strconv"

	"github.com/kralamoure/d1proto"
)

type ExchangeBigStoreType struct {
	Type int
}

func (m ExchangeBigStoreType) ProtocolId() d1proto.MsgCliId {
	return d1proto.ExchangeBigStoreType
}

func (m ExchangeBigStoreType) Serialized() (string, error) {
	return fmt.Sprintf("%d", m.Type), nil
}

func (m *ExchangeBigStoreType) Deserialize(extra string) error {
	v, err := strconv.ParseInt(extra, 10, 32)
	if err != nil {
		return err
	}
	m.Type = int(v)
	return nil
}
