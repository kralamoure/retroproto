package msgsvr

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/kralamoure/retroproto"
)

type EmotesList struct {
	Emotes []int
}

func (m EmotesList) MessageId() retroproto.MsgSvrId {
	return retroproto.EmotesList
}

func (m EmotesList) Serialized() (string, error) {
	var emotes int

	for _, v := range m.Emotes {
		if v == 0 || v >= 32 {
			return "", retroproto.ErrInvalidMsg
		}
		emotes += 1 << (v - 1)
	}

	return fmt.Sprintf("%d|%d", emotes, 0), nil
}

func (m *EmotesList) Deserialize(extra string) error {
	sli := strings.Split(extra, "|")
	if len(sli) < 2 {
		return retroproto.ErrInvalidMsg
	}

	emotesMap := make(map[int]struct{})

	if sli[0] != "" {
		emotes, err := strconv.ParseInt(sli[0], 10, 32)
		if err != nil {
			return err
		}

		if emotes > 0 {
			for i := 0; i < 32; i++ {
				if int(emotes)>>i&1 == 1 {
					id := i + 1
					emotesMap[id] = struct{}{}
				}
			}
		}
	}

	if sli[1] != "" {
		emotes, err := strconv.ParseInt(sli[1], 10, 32)
		if err != nil {
			return err
		}

		if emotes > 0 {
			for i := 0; i < 32; i++ {
				if int(emotes)>>i&1 == 1 {
					id := i + 1
					emotesMap[id] = struct{}{}
				}
			}
		}
	}

	var emotes []int
	for emote := range emotesMap {
		emotes = append(emotes, emote)
	}
	sort.Slice(emotes, func(i, j int) bool {
		return i < j
	})
	m.Emotes = emotes

	return nil
}
