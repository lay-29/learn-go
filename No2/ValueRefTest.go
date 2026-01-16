package No2

import "fmt"

type BigDataArray struct {
	Data [1024]byte
}
type BigDataSlice struct {
	Data []byte
}

func ValueRefTest() {
	da := BigDataArray{Data: [1024]byte{99, 100, 101}}
	ds := BigDataSlice{Data: []byte{199, 200, 201}}
	ChangeArrayData(da)
	ChangeSliceData(ds)
	fmt.Println("值类型不改变", da.Data[0])
	fmt.Println("引用类型改变", ds.Data[0])
	ChangeArrayDataPtr(&da)
	fmt.Println("通过引用传递的值类型会改变", da.Data[0])

}
func ChangeArrayData(da BigDataArray) {
	da.Data[0] = 1
}
func ChangeSliceData(da BigDataSlice) {
	da.Data[0] = 2
}
func ChangeArrayDataPtr(da *BigDataArray) {
	da.Data[0] = 3
}
