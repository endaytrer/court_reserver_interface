package court_reserver_interface

import "github.com/google/uuid"

type TrackPoint struct {
	X    int    `json:"x"`
	Y    int    `json:"y"`
	Type string `json:"type"`
	T    int    `json:"t"`
}
type ChallengeMessage struct {
	BgImageHeight     int          `json:"bgImageHeight"`
	BgImageWidth      int          `json:"bgImageWidth"`
	EndSlidingTime    string       `json:"endSlidingTime"`
	SliderImageHeight int          `json:"sliderImageHeight"`
	SliderImageWidth  int          `json:"sliderImageWidth"`
	StartSlidingTime  string       `json:"startSlidingTime"`
	TrackList         []TrackPoint `json:"trackList"`
}
type SolveResult struct {
	Uuid         uuid.UUID
	Msg          ChallengeMessage
	ChallengeUrl string
}

type CaptchaInfo struct {
	Id      string `json:"id"`
	Captcha struct {
		// blog of image
		BackgroundImage       string      `json:"backgroundImage"`
		SliderImage           string      `json:"sliderImage"`
		BackgroundImageWidth  uint        `json:"backgroundImageWidth"`
		BackgroundImageHeight uint        `json:"backgroundImageHeight"`
		SliderImageWidth      uint        `json:"sliderImageWidth"`
		SliderImageHeight     uint        `json:"sliderImageHeight"`
		Data                  interface{} `json:"data"`
	} `json:"captcha"`
}

type CaptchaSolver interface {
	Solve(CaptchaInfo) SolveResult
}
