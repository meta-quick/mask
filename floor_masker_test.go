package mask

import (
	"fmt"
	"testing"
	"time"
)

func Test_mask_floor(t *testing.T) {
	var floor = NewFloorMasker()
	fmt.Println(floor.MaskFloat32(3.1234))
	fmt.Println(floor.MaskFloat64(5.12334134))
	d,_ := floor.MaskTime(time.Now(),"YMDHms")
	fmt.Println(d)
}
