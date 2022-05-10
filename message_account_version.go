package retroproto

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/kralamoure/retropb/gen/go/retropb"
)

type AccountVersion struct {
	*retropb.AccountVersion
}

func NewAccountVersion(extra string) (AccountVersion, error) {
	var m retropb.AccountVersion

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
		return AccountVersion{}, errInvalidMessage
	case 4:
		beta, err := strconv.ParseInt(sli[3], 10, 32)
		if err != nil {
			return AccountVersion{}, fmt.Errorf("could not parse int: %w", err)
		}
		m.Beta = int32(beta)

		fallthrough
	case 3:
		major, err := strconv.ParseInt(sli[0], 10, 32)
		if err != nil {
			return AccountVersion{}, fmt.Errorf("could not parse int: %w", err)
		}
		m.Major = int32(major)

		minor, err := strconv.ParseInt(sli[1], 10, 32)
		if err != nil {
			return AccountVersion{}, fmt.Errorf("could not parse int: %w", err)
		}
		m.Minor = int32(minor)

		patch, err := strconv.ParseInt(sli[2], 10, 32)
		if err != nil {
			return AccountVersion{}, fmt.Errorf("could not parse int: %w", err)
		}
		m.Patch = int32(patch)
	}

	return AccountVersion{AccountVersion: &m}, nil
}

func (m AccountVersion) MessageId() string {
	return ""
}

func (m AccountVersion) String() string {
	var beta string
	if m.GetBeta() > 0 {
		beta = fmt.Sprintf(".%d", m.GetBeta())
	}

	var streaming string
	if m.GetStreaming() {
		streaming = "s"
	}

	var electron string
	if m.GetElectron() {
		electron = "e"
	}

	return fmt.Sprintf("%d.%d.%d%s%s%s", m.GetMajor(), m.GetMinor(), m.GetPatch(), beta, streaming, electron)
}
