package _case

import "fmt"

// iota表示从0开始的一个数，后续默认 + 1
const (
	B = 1 << (10 * iota)
	KB
	MB
	GB
	TB
)

type Gender uint

const (
	FEMALE Gender = iota
	MALE
	THIRD
)

func ConstAndEnumCase() {
	const (
		A = 10
		B = 20
	)
	size := MB
	var gender Gender = MALE
	fmt.Println(A, B, gender, size)
}
