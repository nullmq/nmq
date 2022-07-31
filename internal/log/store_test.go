package log

import (
	"fmt"
	"testing"
	"unsafe"
)

func Test(t *testing.T) {
	i := uint64(64)
	fmt.Println(unsafe.Sizeof(i))
	t.FailNow()
}
