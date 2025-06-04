package random

import "time"

func Elememt(array []string) (index int64, element string) {
	index = time.Now().Unix() % int64(len(array))
	return index, array[index]
}
