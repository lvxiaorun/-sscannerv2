package tool

import (
	"math/rand"
	"time"
)

//记录出现的牌面以及数量 并且返回有没有数量大于2的
func Int8SliceToMap(slice []int8) map[int8]int8 {
	m := map[int8]int8{}
	for _, item := range slice {
		if _, ok := m[item]; ok {
			m[item] = m[item] + 1
		} else {
			m[item] = 1
		}
	}
	return m
}

//打乱数组顺序 高纳德置乱算法

func FisherYates(since []string) []string {
	if len(since) <= 1 {
		return since
	}
	rand.Seed(time.Now().UnixNano())
	l := len(since)
	for l > 1 {
		i := rand.Intn(l)
		since[i], since[l-1] = since[l-1], since[i]
		l--
	}
	return since
}
