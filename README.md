# boostly

[![Go Reference](https://pkg.go.dev/badge/github.com/masih/boostly.svg)](https://pkg.go.dev/github.com/masih/boostly)
[![Go Test](https://github.com/masih/boostly/actions/workflows/go-test.yml/badge.svg)](https://github.com/masih/boostly/actions/workflows/go-test.yml)

> The missing FileCoin Boost Client Library in pure go.

This repository providers types and client libraries to interact with a FileCoin storage provider running Boost. It covers:

* Deal making
* Deal status check
* Retrieval transports query
 
Due to the way Boost code repository is structured, it is not possible to directly depend on the Boost repo.
This project aims to reduce barriers for programmatic interaction with Boost APIs.