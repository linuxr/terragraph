package utils

import (
	"github.com/google/uuid"
)

func GenUUID() string {
	u4 := uuid.New()
	return u4.String()
}

func RemoveDuplicatedArr(strs []string) []string {
	m := make(map[string]struct{})

	for _, a := range strs {
		m[a] = struct{}{}
	}

	arr := make([]string, 0)
	for k := range m {
		arr = append(arr, k)
	}

	return arr
}
