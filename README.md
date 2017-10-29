# golang-parallel-io

An example of how to write fast parallel file-system programs in Golang.

This includes two programs, `recursive.go`, a naive example of a `find`-style utility, and `parallel.go`, a paralellized version that approaches the performance of `find`.

It also includes a benchmark script for you to see for yourself. Here is the output of `benchmark.sh` on my home folder:

```
$ ./benchmark.sh
parellelized recursive (goroutines)
  352176

real	0m9.484s
user	0m6.259s
sys	0m32.707s
find command
find: /Users/alexkreidler/alpine: Permission denied
  352173

real	0m5.013s
user	0m0.324s
sys	0m2.654s
naive recursive
  352175

real	0m24.520s
user	0m1.625s
sys	0m19.701s

```
