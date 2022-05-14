// generated (unrevised)

package msgsvr

import (
	"fmt"

	"github.com/kralamoure/retroproto"
)

type ConquestWorldData struct{}

func NewConquestWorldData(extra string) (ConquestWorldData, error) {
	var m ConquestWorldData

	if err := m.Deserialize(extra); err != nil {
		return ConquestWorldData{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m ConquestWorldData) MessageId() retroproto.MsgSvrId {
	return retroproto.ConquestWorldData
}

func (m ConquestWorldData) MessageName() string {
	return "ConquestWorldData"
}

func (m ConquestWorldData) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *ConquestWorldData) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
