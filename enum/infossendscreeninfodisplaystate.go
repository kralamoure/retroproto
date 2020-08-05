package enum

var InfosSendScreenInfoDisplayState = newInfosSendScreenInfoDisplayState()

func newInfosSendScreenInfoDisplayState() *infosSendScreenInfoDisplayState {
	return &infosSendScreenInfoDisplayState{
		Normal:     0,
		FullScreen: 1,
		Other:      2,
	}
}

type infosSendScreenInfoDisplayState struct {
	Normal     int
	FullScreen int
	Other      int
}
