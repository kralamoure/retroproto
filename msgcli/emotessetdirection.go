package msgcli

import (
	"fmt"
	"strconv"

	"github.com/kralamoure/retroproto"
)

type EmotesSetDirection struct {
	Dir int
}

func (m EmotesSetDirection) MessageId() retroproto.MsgCliId {
	return retroproto.EmotesSetDirection
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
