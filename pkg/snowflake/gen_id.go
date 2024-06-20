package snowflake

import (
	"time"

	"github.com/sony/sonyflake"
)

var (
	sonyFlake     *sonyflake.Sonyflake // 实例
	sonyMachineID uint16               // 机器ID
)

func getMachineID() (uint16, error) {
	return sonyMachineID, nil
}

func Init(machineId uint16) (err error) {
	sonyMachineID = machineId
	t, _ := time.Parse("2006-01-02", "2022-02-09")
	settings := sonyflake.Settings{
		StartTime: t,
		MachineID: getMachineID,
	}
	sonyFlake = sonyflake.NewSonyflake(settings)
	return
}

// GetID 返回生成的id值
func GetID() (uint64) {
	id, _ := sonyFlake.NextID()
	return id
}