package msgcli

import (
	"fmt"
	"strconv"

	"github.com/kralamoure/retroproto"
)

type EmotesSetDirection struct {
	Dir int
}

func NewEmotesSetDirection(extra string) (EmotesSetDirection, error) {
	var m EmotesSetDirection

	if err := m.Deserialize(extra); err != nil {
		return EmotesSetDirection{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m EmotesSetDirection) MessageId() retroproto.MsgCliId {
	return retroproto.EmotesSetDirection
}

func (m EmotesSetDirection) MessageName() string {
	return "EmotesSetDirection"
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
