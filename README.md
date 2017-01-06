# bwdrivers

[![Build Status](https://travis-ci.org/xllora/bwdrivers.svg?branch=master)](https://travis-ci.org/xllora/bwdrivers) [![GoDoc](https://godoc.org/github.com/xllora/bwdrivers?status.svg)](https://godoc.org/github.com/xllora/bwdrivers)

This repository contains [BadWolf](http://google.github.io/badwolf/) storage
driver implementations that provide persistent storage. It also provides an
updated `bw` 
[command line tool](https://github.com/google/badwolf/blob/master/docs/command_line_tool.md)
that registers all the drivers contained in this repository.

Currently available driver implementations:

* `bwbolt` stores and indexes all triples using the key value store
  [BoltDB](https://github.com/boltdb/bolt).

## Bulding the command line tool

Assumming you have a fully working installation of GO, you just need to get 
the required packages and build the tool. You can achieve this by typing the 
following commands:

```
$ cd /to/some/new/folder
$ export GOPATH=$PWD
$ go get golang.org/x/net/context
$ go get github.com/peterh/liner
$ go get github.com/google/badwolf/...
$ go get github.com/xllora/bwdrivers/...
```

These commands should not output anything unless something fails. You can 
check that the new `bw` tools contains the new drivers by checking the 
`--driver` help flag information. You should see drivers other than 
`VOLATILE`

```
 $ ./bin/bw -h
Usage of ./bin/bw:
  -bolt_db_path string
    	The path to the Bolt database to use.
  -bql_channel_size int
    	Internal channel size to use on BQL queries.
  -driver string
    	The storage driver to use {VOLATILE|BWBOLT}. (default "VOLATILE")
```

## Testing the command line tool

You may always want to run all the test for both repos `badwolf` and 
`bwdrivers` to make sure eveything is A-Ok. If the tests fail, you should
consider not using the build tool since it may be tainted.

```
$ go test -race github.com/google/badwolf/... github.com/xllora/bwdrivers/...
ok  	github.com/google/badwolf/bql/grammar	1.175s
ok  	github.com/google/badwolf/bql/lexer	1.045s
ok  	github.com/google/badwolf/bql/planner	1.531s
ok  	github.com/google/badwolf/bql/semantic	1.081s
ok  	github.com/google/badwolf/bql/table	1.076s
?   	github.com/google/badwolf/bql/version	[no test files]
ok  	github.com/google/badwolf/io	1.054s
?   	github.com/google/badwolf/storage	[no test files]
ok  	github.com/google/badwolf/storage/memory	1.071s
ok  	github.com/google/badwolf/tools/compliance	1.170s
?   	github.com/google/badwolf/tools/vcli/bw	[no test files]
?   	github.com/google/badwolf/tools/vcli/bw/assert	[no test files]
?   	github.com/google/badwolf/tools/vcli/bw/command	[no test files]
?   	github.com/google/badwolf/tools/vcli/bw/common	[no test files]
?   	github.com/google/badwolf/tools/vcli/bw/io	[no test files]
?   	github.com/google/badwolf/tools/vcli/bw/run	[no test files]
?   	github.com/google/badwolf/tools/vcli/bw/version	[no test files]
ok  	github.com/google/badwolf/triple	1.053s
ok  	github.com/google/badwolf/triple/literal	1.044s
ok  	github.com/google/badwolf/triple/node	1.047s
ok  	github.com/google/badwolf/triple/predicate	1.020s
?   	github.com/xllora/bwdrivers/bw	[no test files]
ok  	github.com/xllora/bwdrivers/bwbolt	1.277s
```
