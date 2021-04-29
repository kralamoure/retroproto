package msgcli

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/kralamoure/d1proto"
)

type InfosSendScreenInfo struct {
	Width        int
	Height       int
	DisplayState int
}

func (m InfosSendScreenInfo) ProtocolId() d1proto.MsgCliId {
	return d1proto.InfosSendScreenInfo
}

func (m InfosSendScreenInfo) Serialized() (string, error) {
	return fmt.Sprintf("%d;%d;%d", m.Width, m.Height, m.DisplayState), nil
}

func (m *InfosSendScreenInfo) Deserialize(extra string) error {
	sli := strings.Split(extra, ";")
	if len(sli) < 3 {
		return d1proto.ErrInvalidMsg
	}

	width, err := strconv.ParseInt(sli[0], 10, 32)
	if err != nil {
		return err
	}
	m.Width = int(width)

	height, err := strconv.ParseInt(sli[1], 10, 32)
	if err != nil {
		return err
	}
	m.Height = int(height)

	displayState, err := strconv.ParseInt(sli[2], 10, 32)
	if err != nil {
		return err
	}
	m.DisplayState = int(displayState)

	return nil
}
