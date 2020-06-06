package msgcli

import (
	"fmt"
	"strconv"

	"github.com/kralamoure/d1proto"
)

type EmotesSetDirection struct {
	Dir int
}

func (m EmotesSetDirection) ProtocolId() d1proto.MsgCliId {
	return d1proto.EmotesSetDirection
}

func (m EmotesSetDirection) Serialized() (string, error) {
	return fmt.Sprintf("%d", m.Dir), nil
}

func (m *EmotesSetDirection) Deserialize(extra string) error {
	dir, err := strconv.ParseInt(extra, 10, 32)
	if err != nil {
		return err
	}
	m.Dir = int(dir)

	return nil
}
