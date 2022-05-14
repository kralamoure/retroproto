package msgcli

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/kralamoure/retroproto"
)

type InfosSendScreenInfo struct {
	Width        int
	Height       int
	DisplayState int
}

func (m InfosSendScreenInfo) MessageId() retroproto.MsgCliId {
	return retroproto.InfosSendScreenInfo
}

func (m InfosSendScreenInfo) MessageName() string {
	return "InfosSendScreenInfo"
}

func (m InfosSendScreenInfo) Serialized() (string, error) {
	return fmt.Sprintf("%d;%d;%d", m.Width, m.Height, m.DisplayState), nil
}

func (m *InfosSendScreenInfo) Deserialize(extra string) error {
	sli := strings.Split(extra, ";")
	if len(sli) < 3 {
		return retroproto.ErrInvalidMsg
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
