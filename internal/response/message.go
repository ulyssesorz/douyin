package response

import (
	"github.com/ulyssesorz/douyin/kitex/kitex_gen/message"
)

type MessageChat struct {
	Base
	MessageList []*message.Message `json:"message_list"`
}

type MessageAction struct {
	Base
}
