## 深度拷贝库性能一览

```
goos: darwin
goarch: arm64
pkg: benchmark
Benchmark_Use_reflectValue_MiniCopy-8   	  334728	      3575 ns/op
Benchmark_Use_reflectValue_DeepCopy-8   	  595302	      1956 ns/op
Benchmark_Use_reflectValue_Copier-8     	  203574	      5860 ns/op
Benchmark_Use_Ptr_jsoniter-8            	  821113	      1477 ns/op
Benchmark_Use_Ptr_pcopy-8               	 3390382	       354.0 ns/op
Benchmark_Use_Ptr_coven-8               	 1414197	       848.7 ns/op
PASS
ok  	benchmark	9.771s
```
