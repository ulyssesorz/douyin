package redis

import (
	"context"
	"fmt"
	"strconv"
	"time"
)

// redis存储了上次消息的最晚时间戳，key为token_userid，获取聊天记录时先从redis拿到时间戳
func GetMessageTimestamp(ctx context.Context, token string, toUserID int64) (int, error) {
	key := fmt.Sprintf("%s_%d", token, toUserID)
	if ec, err := GetRedisHelper().Exists(ctx, key).Result(); err != nil {
		return -1, err
	} else if ec == 0 {
		return -1, nil
	}

	val, err := GetRedisHelper().Get(ctx, key).Result()
	if err != nil {
		return -1, err
	}

	return strconv.Atoi(val)
}

func SetMessageTimestamp(ctx context.Context, token string, toUserID int64, timestamp int) error {
	key := fmt.Sprintf("%s_%d", token, toUserID)
	return GetRedisHelper().Set(ctx, key, timestamp, 2*time.Second).Err()
}
