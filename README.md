# Hiro

Pre-alpha 0.01 not working yet.

Plan for Hiro: [https://www.hirolang.org](https://www.hirolang.org)

Build Hiro:

`git build`

should create a `hiro` executable.

`mkdir target`

creates a target for transpiled Go code.

`./hiro -s addexample/ -t target/`

transpiles the `addexample` to target.

`(cd target; go get -t HiroTesting; go mod vendor; go test ./...)`

runs the tests for the generated Go code.