package main
import(
	"fmt"
	"unsafe"
	"math"
)
//数字类型

func maink() {
	//整型
	var i8 int8
	var i16 int16
	var i32 int32
	var i64 int64
	//类型，大小，最小值，最大值
	fmt.Printf("%T %dB %v~%v\n", i8, unsafe.Sizeof(i8), math.MinInt8, math.MaxInt8)
	fmt.Printf("%T %dB %v~%v\n", i16, unsafe.Sizeof(i16), math.MinInt16, math.MaxInt16)
	fmt.Printf("%T %dB %v~%v\n", i32, unsafe.Sizeof(i32), math.MinInt32, math.MaxInt32)
	fmt.Printf("%T %dB %v~%v\n", i64, unsafe.Sizeof(i64), math.MinInt64, math.MaxInt64)
	//无符号整型
	var u8 uint8
	var u16 uint16
	var u32 uint32
	var u64 uint64
	fmt.Printf("%T %dB %v~%v\n", u8, unsafe.Sizeof(u8), 0, math.MaxUint8)
	fmt.Printf("%T %dB %v~%v\n", u16, unsafe.Sizeof(u16), 0, math.MaxUint16)
	fmt.Printf("%T %dB %v~%v\n", u32, unsafe.Sizeof(u32), 0, math.MaxUint32)
	fmt.Printf("%T %dB %v~%v\n", u64, unsafe.Sizeof(u64), 0, uint64(math.MaxUint64))
	//浮点型
	var f32 float32
	var f64 float64
	fmt.Printf("%T %dB %v~%v\n", f32, unsafe.Sizeof(f32), math.SmallestNonzeroFloat32, math.MaxFloat32)
	fmt.Printf("%T %dB %v~%v\n", f64, unsafe.Sizeof(f64), math.SmallestNonzeroFloat64, math.MaxFloat64)


	//进制
	var a int = 10
	fmt.Printf("%d\n", a)//十进制
	fmt.Printf("%b\n", a)//二进制
	fmt.Printf("%o\n", a)//八进制
	fmt.Printf("%x\n", a)//十六进制
}