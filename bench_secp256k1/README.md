# Benchmark secp256k1 implementation in golang

There is a golang implementation of secp256k1 (github.com/btcsuite/btcd/btcec), they implemented sign/verify/recover all in golang. 

In this directory, I wrote a benchmark to compare the btcec implementation with cgo implementation.

# Run benchmark

```
$ go test -bench=.
```

# Run benchmark memory

```
$ go test -bench=. -benchmem
```

# Result

Environment: 2.2 GHz 6-Core Intel Core i7

## Benchmark
```
$ go test -bench=.
goos: darwin
goarch: amd64
pkg: bench_secp256k1
BenchmarkBTCECSignMessage-12         	   22225	     53195 ns/op
BenchmarkCGOSignMessage-12           	   28641	     42017 ns/op
BenchmarkBTCECRecoverPublicKey-12    	    5499	    213794 ns/op
BenchmarkCGORecoverPublicKey-12      	   21181	     56866 ns/op
BenchmarkBTCECVerifyMessage-12       	    6538	    181210 ns/op
BenchmarkCGOVerifyMessage-12         	   25294	     48091 ns/op
PASS
ok  	bench_secp256k1	9.341s
```

## Benchmark memory
```
$ go test -bench=. -benchmem
goos: darwin
goarch: amd64
pkg: bench_secp256k1
BenchmarkBTCECSignMessage-12         	   22222	     54137 ns/op	    4915 B/op	      73 allocs/op
BenchmarkCGOSignMessage-12           	   28312	     42111 ns/op	     164 B/op	       3 allocs/op
BenchmarkBTCECRecoverPublicKey-12    	    5464	    214905 ns/op	    3429 B/op	      69 allocs/op
BenchmarkCGORecoverPublicKey-12      	   21145	     57038 ns/op	     160 B/op	       2 allocs/op
BenchmarkBTCECVerifyMessage-12       	    6564	    181211 ns/op	    3074 B/op	      62 allocs/op
BenchmarkCGOVerifyMessage-12         	   24260	     48067 ns/op	       0 B/op	       0 allocs/op
PASS
ok  	bench_secp256k1	9.366s
```