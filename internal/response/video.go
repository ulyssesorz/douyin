package response

import (
	"github.com/ulyssesorz/douyin/kitex/kitex_gen/video"
)

type Feed struct {
	Base
	NextTime  int64          `json:"next_time"`
	VideoList []*video.Video `json:"video_list"`
}

type PublishAction struct {
	Base
}

type PublishList struct {
	Base
	VideoList []*video.Video `json:"video_list"`
}
