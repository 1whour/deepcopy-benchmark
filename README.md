## 深度拷贝库性能一览

## 使用refluct.Value库
* https://github.com/antlabs/deepcopy 性能最高
```
Benchmark_Use_reflectValue_MiniCopy-16    	  160227	      6761 ns/op
Benchmark_Use_reflectValue_DeepCopy-16    	  357322	      3007 ns/op
Benchmark_Use_reflectValue_Copier-16      	  115042	     10447 ns/op

```
## 使用ptr库
* https://github.com/petersunbag/coven 性能最高  
不过coven库，在对待slice变量并不是深度拷贝，感兴趣可以在本项目目录运行```go test```
```
Benchmark_Use_Ptr_jsoniter-16             	  343550	      3233 ns/op
Benchmark_Use_Ptr_fastdeepcopy-16         	  439089	      2877 ns/op
Benchmark_Use_Ptr_coven-16                	 1389820	       883 ns/op

```
