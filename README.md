# golang-parallel-io

A benchmark of filesystem IO and an example of parallel programming in Go.

This includes three programs, `recursive.go`, a naive example of a `find`-style utility, `parallel.go`, a paralellized version that approaches the performance of `find`, and `walk.go`, which uses Go's `filepath.Walk` function.

It also includes a benchmark script for you to see for yourself. Here is the output of `benchmark.sh` on my home folder:

```
$ ./benchmark.sh
CPUs/Cores: 2
GOMAXPROCS: 2
find /Users/alexkreidler
  274165

real	0m2.046s
user	0m0.416s
sys	0m1.640s
./recursive /Users/alexkreidler
  274165

real	0m13.127s
user	0m1.751s
sys	0m10.294s
./parallel /Users/alexkreidler
  274165

real	0m9.120s
user	0m4.781s
sys	0m10.676s
./walk /Users/alexkreidler
  274165

real	0m13.287s
user	0m2.033s
sys	0m10.863s
```

This is on an older MacBook Pro with an SSD running MacOS Sierra (10.11) As you can see above, it only has 2 cores, so the parallel version doesn't have as much of an improvement as on newer computers with more cores. 
