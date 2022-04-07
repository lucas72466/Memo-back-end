package public

import (
	"github.com/sony/sonyflake"
	"github.com/spf13/viper"
	"time"
)

var iIDGenerator *sonyflake.Sonyflake

func InitIDGenerator() {
	startTime, err := time.Parse(viper.GetString("idGenerator.timeLayout"), viper.GetString("idGenerator.startTime"))
	if err != nil {
		panic(err)
	}

	generator := sonyflake.NewSonyflake(sonyflake.Settings{
		StartTime:      startTime,
		MachineID:      nil,
		CheckMachineID: nil,
	})

	if generator == nil {
		panic("id generator init fail")
	}

	iIDGenerator = generator
}

func GetUniqueID() uint64 {
	// NextID will return err only when Sonyflake time is over the limit
	id, _ := iIDGenerator.NextID()
	return id
}
