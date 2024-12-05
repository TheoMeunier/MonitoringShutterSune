package utils

import "time"

func IsTimeToAct(now, targetTime time.Time, actionType string) bool {
	diff := now.Sub(targetTime)
	return diff >= 0 && diff < time.Minute*1
}
