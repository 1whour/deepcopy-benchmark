## 深度拷贝库性能一览

## 使用refluct.Value库
* https://github.com/antlabs/deepcopy 性能最高
```
Benchmark_Use_reflectValue_MiniCopy-16    	  160227	      6761 ns/op
Benchmark_Use_reflectValue_DeepCopy-16    	  357322	      3007 ns/op
Benchmark_Use_reflectValue_Copier-16      	  115042	     10447 ns/op

```
## 使用ptr库
* https://github.com/antlabs/fastdeepcopy 性能最高，还在优化过程中
```
Benchmark_Use_Ptr_jsoniter-16             	  634638	      1687 ns/op
Benchmark_Use_Ptr_fastdeepcopy-16         	 3096130	       405 ns/op
Benchmark_Use_Ptr_coven-16                	 1304331	       921 ns/op

```
