// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1proto"
)

type GameFightChallengeUpdateError struct{}

func (m GameFightChallengeUpdateError) ProtocolId() d1proto.MsgSvrId {
	return d1proto.GameFightChallengeUpdateError
}

func (m GameFightChallengeUpdateError) Serialized() (string, error) {
	return "", nil
}

func (m *GameFightChallengeUpdateError) Deserialize(extra string) error {
	return nil
}
