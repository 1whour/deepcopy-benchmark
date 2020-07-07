package benchmark

/*
func Test_Use_reflectvalue_deepcopy(t *testing.T) {
	dst := testDataDst{}
	deepcopy.Copy(&dst, &td).Do()
	assert.Equal(t, td, dst)

	dst.Slice[0] = "aaa"
	fmt.Println("deepcopy:", td.Slice)
}

func Test_Use_Ptr_coven(t *testing.T) {

	c, err := coven.NewConverter(testDataDst{}, testDataSrc{})
	assert.NoError(t, err)

	dst := testDataDst{}
	c.Convert(&dst, &td)
	assert.Equal(t, td, dst)
	// coven并不是一个深度拷贝库
	// 修改dst里面的数据, td里面的slice数据也跟着被改变
	dst.Slice[0] = "aaa"
	fmt.Println("coven:", td.Slice)
}

// 看下coven是如何处理指针变量的
// 结论coven只是拷贝指针地址
func Test_Use_Ptr_coven_Cycle(t *testing.T) {
	type Ring struct {
		R *Ring
	}

	c, err := coven.NewConverter(Ring{}, Ring{})
	assert.NoError(t, err)

	R := Ring{}
	R.R = &R

	r2 := Ring{}

	err = c.Convert(&r2, &R)
	fmt.Printf("%p\n", R.R)
	fmt.Printf("%p\n", r2.R)
}
*/
