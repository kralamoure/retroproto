// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1proto"
)

type ExchangeCraftLoopEnd struct{}

func (m ExchangeCraftLoopEnd) ProtocolId() d1proto.MsgSvrId {
	return d1proto.ExchangeCraftLoopEnd
}

func (m ExchangeCraftLoopEnd) Serialized() (string, error) {
	return "", nil
}

func (m *ExchangeCraftLoopEnd) Deserialize(extra string) error {
	return nil
}
