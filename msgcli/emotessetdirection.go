package msgcli

import (
	"fmt"
	"strconv"

	"github.com/kralamoure/d1encoding"
)

type EmotesSetDirection struct {
	Dir int
}

func (m EmotesSetDirection) ProtocolId() d1encoding.MsgCliId {
	return d1encoding.EmotesSetDirection
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
