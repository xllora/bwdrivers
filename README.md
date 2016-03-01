# bwdrivers

[![Build Status](https://travis-ci.org/xllora/bwdrivers.svg?branch=master)](https://travis-ci.org/xllora/bwdrivers) [![GoDoc](https://godoc.org/github.com/xllora/bwdrivers?status.svg)](https://godoc.org/github.com/xllora/bwdrivers)

This repository contains [BadWolf](http://google.github.io/badwolf/) storage
driver implementations that provide persistent storage.

Currently available driver implementations:

* `bwbolt` stores and indexes all triples using the key value store
  [BoltDB](https://github.com/boltdb/bolt).
