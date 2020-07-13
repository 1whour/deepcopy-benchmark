package benchmark

import (
	"encoding/json"
	"testing"

	"github.com/antlabs/deepcopy"
	"github.com/antlabs/fastdeepcopy"
	"github.com/jinzhu/copier"
	jsoniter "github.com/json-iterator/go"
	"github.com/petersunbag/coven"
)

type testDataDst struct {
	Int64  int64  `json:"int_64"`
	Int32  int32  `json:"int_32"`
	Int16  int8   `json:"int_16"`
	Int8   int8   `json:"int_8"`
	UInt8  int8   `json:"u_int_8"`
	UInt64 uint64 `json:"u_int_64"`
	UInt32 uint32 `json:"u_int_32"`
	UInt16 uint8  `json:"u_int_16"`
	S      string `json:"s"`
	Array  [4]int `json:"array"`
}

type testDataSrc struct {
	Int64  int64  `json:"int_64"`
	Int32  int32  `json:"int_32"`
	Int16  int8   `json:"int_16"`
	Int8   int8   `json:"int_8"`
	UInt8  int8   `json:"u_int_8"`
	UInt64 uint64 `json:"u_int_64"`
	UInt32 uint32 `json:"u_int_32"`
	UInt16 uint8  `json:"u_int_16"`
	S      string `json:"s"`
	Array  [4]int `json:"array"`
}

var td = testDataSrc{
	Int64:  64,
	Int32:  32,
	Int16:  16,
	Int8:   8,
	UInt8:  18,
	UInt64: 164,
	UInt32: 132,
	UInt16: 116,
	S:      "test deepcopy",
	Array:  [4]int{1, 2, 3},
}

func miniCopy(dst, src interface{}) error {
	bytes, err := json.Marshal(src)
	if err != nil {
		return err
	}

	return json.Unmarshal(bytes, dst)
}

func user_jsoniters(dst, src interface{}) error {
	bytes, err := jsoniter.Marshal(src)
	if err != nil {
		return err
	}
	return jsoniter.Unmarshal(bytes, dst)
}

func Benchmark_Use_reflectValue_MiniCopy(b *testing.B) {

	for i := 0; i < b.N; i++ {
		var dst testDataDst
		miniCopy(&dst, &td)
	}
}

func Benchmark_Use_reflectValue_DeepCopy(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var dst testDataDst
		deepcopy.Copy(&dst, &td).Do()
	}
}

func Benchmark_Use_reflectValue_Copier(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var dst testDataDst
		copier.Copy(&dst, &td)
	}
}

func Benchmark_Use_Ptr_jsoniter(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var dst testDataDst
		user_jsoniters(&dst, &td)
	}
}

func Benchmark_Use_Ptr_fastdeepcopy(b *testing.B) {
	fastdeepcopy.OpenCache = true
	for i := 0; i < b.N; i++ {
		var dst testDataDst
		fastdeepcopy.Copy(&dst, &td).Do()
	}
}

func Benchmark_Use_Ptr_coven(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var dst testDataDst
		c, _ := coven.NewConverter(testDataDst{}, testDataSrc{})
		c.Convert(&dst, &td)
	}
}
