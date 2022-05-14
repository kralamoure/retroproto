package msgsvr

import (
	"fmt"
	"strconv"

	"github.com/kralamoure/retroproto"
)

type ExchangeMountStorageRemove struct {
	MountId int
}

func NewExchangeMountStorageRemove(extra string) (ExchangeMountStorageRemove, error) {
	var m ExchangeMountStorageRemove

	if err := m.Deserialize(extra); err != nil {
		return ExchangeMountStorageRemove{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m ExchangeMountStorageRemove) MessageId() retroproto.MsgSvrId {
	return retroproto.ExchangeMountStorageRemove
}

func (m ExchangeMountStorageRemove) MessageName() string {
	return "ExchangeMountStorageRemove"
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
