package msgsvr

import (
	"fmt"
	"strconv"

	"github.com/kralamoure/d1encoding"
)

type ExchangeMountStorageRemove struct {
	MountId int
}

func (m ExchangeMountStorageRemove) ProtocolId() d1encoding.MsgSvrId {
	return d1encoding.ExchangeMountStorageRemove
}

func (m ExchangeMountStorageRemove) Serialized() (string, error) {
	return fmt.Sprint(m.MountId), nil
}

func (m *ExchangeMountStorageRemove) Deserialize(extra string) error {
	mountId, err := strconv.Atoi(extra)
	if err != nil {
		return err
	}
	m.MountId = mountId

	return nil
}
