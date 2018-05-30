package tool

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
