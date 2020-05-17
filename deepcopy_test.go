package benchmark

import (
	"fmt"
	"testing"

	"github.com/antlabs/deepcopy"
	"github.com/petersunbag/coven"
	"github.com/stretchr/testify/assert"
)

func Test_Use_reflectvalue_deepcopy(t *testing.T) {
	dst := testData{}
	deepcopy.Copy(&dst, &td).Do()
	assert.Equal(t, td, dst)

	dst.Slice[0] = "aaa"
	fmt.Println("deepcopy:", td.Slice)
}

func Test_Use_Ptr_coven(t *testing.T) {

	c, err := coven.NewConverter(testData{}, testData{})
	assert.NoError(t, err)

	dst := testData{}
	c.Convert(&dst, &td)
	assert.Equal(t, td, dst)
	// coven并不是一个深度拷贝库
	// 修改dst里面的数据, td里面的slice数据也跟着被改变
	dst.Slice[0] = "aaa"
	fmt.Println("coven:", td.Slice)
}
