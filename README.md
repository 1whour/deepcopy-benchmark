## 深度拷贝库性能一览

## 使用refluct.Value库
* https://github.com/antlabs/deepcopy 性能最高
```
Benchmark_Use_reflectValue_MiniCopy-12    	  145412	      7451 ns/op
Benchmark_Use_reflectValue_DeepCopy-12    	  183664	      5587 ns/op
Benchmark_Use_reflectValue_Copier-12      	   91088	     14915 ns/op
```
## 使用ptr库
* https://github.com/petersunbag/coven 性能最高  
不过coven库，在对待slice变量并不是深度拷贝，感兴趣可以在本项目目录运行```go test```
```
Benchmark_Use_Ptr_jsoniter-12             	  811066	      1627 ns/op
Benchmark_Use_Ptr_coven-12                	 7364188	       158 ns/o
```

## deepcopy 下个版本计划
使用ptr方式重新优化一个版本，名字会和deepcopy项目区别出来，毕竟reflect是根正苗红，官方认证的做法，ptr已经属于hack范畴。
