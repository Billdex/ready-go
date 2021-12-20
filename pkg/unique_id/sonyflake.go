package unique_id

import (
	"github.com/sony/sonyflake"
	"strconv"
)

var flack = sonyflake.NewSonyflake(sonyflake.Settings{})

func SonyFlackSetup(st sonyflake.Settings) {
	flack = sonyflake.NewSonyflake(st)
}

func SonyFlakeID() (string, error) {
	id, err := flack.NextID()
	if err != nil {
		return "", err
	}
	return strconv.FormatUint(id, 10), nil
}
