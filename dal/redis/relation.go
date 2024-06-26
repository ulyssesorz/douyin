package redis

import (
	"context"
	"fmt"
	"strconv"
	"strings"
)

type RelationCache struct {
	UserID     uint `json:"user_id" redis:"user_id"`
	ToUserID   uint `json:"to_user_id" redis:"to_user_id"`
	ActionType uint `json:"action_type" redis:"action_type"`
	CreatedAt  uint `json:"created_at" redis:"created_at"`
}

// 更新关系，和点赞逻辑相同
func UpdateRelation(ctx context.Context, relation *RelationCache) error {
	errLock := LockByMutex(ctx, RelationMutex)
	if errLock != nil {
		zapLogger.Errorf("lock failed: %s", errLock.Error())
		return errLock
	}

	keyRelationRead := fmt.Sprintf("user::%d::to_user::%d::r", relation.UserID, relation.ToUserID)
	keyRelationWrite := fmt.Sprintf("user::%d::to_user::%d::w", relation.UserID, relation.ToUserID)
	valueRedis := fmt.Sprintf("%d::%d", relation.CreatedAt, relation.ActionType)

	readExisted, err := GetRedisHelper().Exists(ctx, keyRelationWrite).Result()
	if err != nil {
		errUnlock := UnlockByMutex(ctx, RelationMutex)
		if errUnlock != nil {
			zapLogger.Errorf("unlock failed: %s", errUnlock.Error())
			return errUnlock
		}
		zapLogger.Errorf("Get Redis data error: %v", err.Error())
		return err
	}
	if readExisted == 0 {
		// redis中不存在，直接加入
		err := setKey(ctx, keyRelationRead, valueRedis, ExpireTime, RelationMutex)
		if err != nil {
			zapLogger.Errorln(err.Error())
			return err
		}

		errLock := LockByMutex(ctx, RelationMutex)
		if errLock != nil {
			zapLogger.Errorf("lock failed: %s", errLock.Error())
			return errLock
		}
		err = setKey(ctx, keyRelationWrite, valueRedis, 0, RelationMutex)
		if err != nil {
			zapLogger.Errorln(err.Error())
			return err
		}
	} else {
		res, _ := GetRedisHelper().Get(ctx, keyRelationRead).Result()
		vSplit := strings.Split(res, "::")
		redis_ct, redis_at := vSplit[0], vSplit[1]
		if redis_at == strconv.Itoa(int(relation.ActionType)) {
			// 若新增的action_type不变，则直接返回
			errUnlock := UnlockByMutex(ctx, RelationMutex)
			if errUnlock != nil {
				zapLogger.Errorf("unlock failed: %s", errUnlock.Error())
				return errUnlock
			}
			return nil
		} else if strconv.Itoa(int(relation.CreatedAt)) > redis_ct {
			// 若action_type改变，且该消息创建时间晚于redis中的消息时间，则替换
			err := setKey(ctx, keyRelationRead, valueRedis, ExpireTime, RelationMutex)
			if err != nil {
				zapLogger.Errorln(err.Error())
				return err
			}

			errLock := LockByMutex(ctx, RelationMutex)
			if errLock != nil {
				zapLogger.Errorf("lock failed: %s", errLock.Error())
				return errLock
			}
			err = setKey(ctx, keyRelationWrite, valueRedis, 0, RelationMutex)
			if err != nil {
				zapLogger.Errorln(err.Error())
				return err
			}
		} else {
			errUnlock := UnlockByMutex(ctx, RelationMutex)
			if errUnlock != nil {
				zapLogger.Errorf("unlock failed: %s", errUnlock.Error())
				return errUnlock
			}
		}
	}

	return nil
}

// 根据用户ID获取粉丝ID列表
func GetFollowerIDs(ctx context.Context, userID int64) (*[]int64, error) {
	key := fmt.Sprintf("follower::%d", userID)
	results, err := GetRedisHelper().SMembers(ctx, key).Result()
	if err != nil {
		zapLogger.Errorln(err.Error())
		return nil, err
	}
	userIDs := make([]int64, 0)
	for _, result := range results {
		id, _ := strconv.ParseInt(result, 10, 64)
		userIDs = append(userIDs, id)
	}
	return &userIDs, nil
}

// 根据用户ID获取关注者ID列表
func GetFollowingIDs(ctx context.Context, userID int64) (*[]int64, error) {
	key := fmt.Sprintf("following::%d", userID)
	results, err := GetRedisHelper().SMembers(ctx, key).Result()
	if err != nil {
		zapLogger.Errorln(err.Error())
		return nil, err
	}
	userIDs := make([]int64, 0)
	for _, result := range results {
		id, _ := strconv.ParseInt(result, 10, 64)
		userIDs = append(userIDs, id)
	}
	return &userIDs, nil
}

// 根据该用户的ID和从Redis获取后的userIDs，获取该用户的粉丝RelationCache列表
func GetUserFollowers(ctx context.Context, userID int64, userIDs []int64) (*[]*RelationCache, error) {
	relations := make([]*RelationCache, 0)
	for _, id := range userIDs {
		relationCache, err := GetRedisHelper().Get(ctx, fmt.Sprintf("user::%d::to_user::%d", id, userID)).Result()
		if err != nil {
			zapLogger.Errorln(err.Error())
			return nil, err
		}
		actionType, err := strconv.ParseInt(relationCache, 10, 64)
		if err != nil {
			zapLogger.Errorln(err.Error())
			return nil, err
		}
		relations = append(relations, &RelationCache{
			UserID:     uint(userID),
			ToUserID:   uint(id),
			ActionType: uint(actionType),
		})
	}
	return &relations, nil
}

// GetUserFollowings 根据该用户的ID和从Redis获取后的userIDs，获取该用户的关注者RelationCache列表
func GetUserFollowings(ctx context.Context, userID int64, userIDs []int64) (*[]*RelationCache, error) {
	relations := make([]*RelationCache, 0)
	for _, id := range userIDs {
		relationCache, err := GetRedisHelper().Get(ctx, fmt.Sprintf("user::%d::to_user::%d", userID, id)).Result()
		if err != nil {
			zapLogger.Errorln(err.Error())
			return nil, err
		}
		actionType, err := strconv.ParseInt(relationCache, 10, 64)
		if err != nil {
			zapLogger.Errorln(err.Error())
			return nil, err
		}
		relations = append(relations, &RelationCache{
			UserID:     uint(userID),
			ToUserID:   uint(id),
			ActionType: uint(actionType),
		})
	}
	return &relations, nil
}
