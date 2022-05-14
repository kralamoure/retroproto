package msgcli

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/kralamoure/retroproto"
)

type AccountVersion struct {
	Major     int
	Minor     int
	Patch     int
	Beta      int
	Streaming bool
	Electron  bool
}

func NewAccountVersion(extra string) (AccountVersion, error) {
	var m AccountVersion

	if err := m.Deserialize(extra); err != nil {
		return AccountVersion{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m AccountVersion) MessageId() retroproto.MsgCliId {
	return retroproto.AccountVersion
}

func (m AccountVersion) MessageName() string {
	return "AccountVersion"
}

func (m AccountVersion) Serialized() (string, error) {
	var beta string
	if m.Beta > 0 {
		beta = fmt.Sprintf(".%d", m.Beta)
	}

	var streaming string
	if m.Streaming {
		streaming = "s"
	}

	var electron string
	if m.Electron {
		electron = "e"
	}

	return fmt.Sprintf("%d.%d.%d%s%s%s", m.Major, m.Minor, m.Patch, beta, streaming, electron), nil
}

func (m *AccountVersion) Deserialize(extra string) error {
	if strings.HasSuffix(extra, "e") {
		m.Electron = true
		extra = strings.TrimSuffix(extra, "e")
	}

	if strings.HasSuffix(extra, "s") {
		m.Streaming = true
		extra = strings.TrimSuffix(extra, "s")
	}

	sli := strings.Split(extra, ".")
	switch len(sli) {
	default:
		return retroproto.ErrInvalidMsg
	case 4:
		beta, err := strconv.ParseInt(sli[3], 10, 32)
		if err != nil {
			return err
		}
		m.Beta = int(beta)

		fallthrough
	case 3:
		major, err := strconv.ParseInt(sli[0], 10, 32)
		if err != nil {
			return err
		}
		m.Major = int(major)

		minor, err := strconv.ParseInt(sli[1], 10, 32)
		if err != nil {
			return err
		}
		m.Minor = int(minor)

		patch, err := strconv.ParseInt(sli[2], 10, 32)
		if err != nil {
			return err
		}
		m.Patch = int(patch)
	}

	return nil
}
