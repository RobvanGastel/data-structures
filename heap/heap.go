package heap

import (
	"math"
	"fmt"
)

func Parent(i int) int {
	return int(math.Floor(float64(i-1)/2.0))
}

func Left(i int) int {
	return 2*i+1
}

func Right(i int) int {
	return 2*i+2
}

func MaxHeapify() {

}

func IsMaxHeap(size int) {
	// for int(math.Ceil(float(size)/2.0)) < 2 {

	// }
}

func BuildMaxHeap() {
	var i = int(math.Ceil(float64(2000)/2.0))
	for i < 2 {
		fmt.Println(i)
	}
}