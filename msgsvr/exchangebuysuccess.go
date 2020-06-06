// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1proto"
)

type ExchangeBuySuccess struct{}

func (m ExchangeBuySuccess) ProtocolId() d1proto.MsgSvrId {
	return d1proto.ExchangeBuySuccess
}

func (m ExchangeBuySuccess) Serialized() (string, error) {
	return "", nil
}

func (m *ExchangeBuySuccess) Deserialize(extra string) error {
	return nil
}
