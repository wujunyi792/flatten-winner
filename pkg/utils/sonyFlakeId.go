package utils

import (
	"github.com/sony/sonyflake"
	"github.com/wujunyi792/flatten-winner/internal/logger"
)

var flake *sonyflake.Sonyflake

func init() {
	flake = sonyflake.NewSonyflake(sonyflake.Settings{})
}

func GenSonyFlakeId() (int64, error) {

	id, err := flake.NextID()
	if err != nil {
		logger.Warning.Println("flake NextID failed: ", err)
		return 0, err
	}

	return int64(id), nil
}
