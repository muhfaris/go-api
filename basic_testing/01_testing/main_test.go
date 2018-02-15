package main

import "testing"

func TestSum(t *testing.T) {
	a := 5
	b := 10

	expected := 15

	res := sum(a, b)

	if res != expected {
		t.Errorf("Ours function doenst work, %d + %d isnt %d", a, b, res)
	}
}

/*
Example Fail Test
Running tool: /opt/devtool/go/go1.9.3/bin/go test -coverprofile=/tmp/go-code-cover -timeout 30s github.com/muhfaris/goFun/basic_testing/01_testing

--- FAIL: TestSum (0.00s)
	/media/ext-data/Project/Authsecur/GoProject/src/github.com/muhfaris/goFun/basic_testing/01_testing/main_test.go:14: Ours function doenst work, 5 + 10 isnt 15
FAIL
coverage: 20.0% of statements
FAIL	github.com/muhfaris/goFun/basic_testing/01_testing	0.001s
Error: Tests failed.

*/

/*
Example Success Test
Running tool: /opt/devtool/go/go1.9.3/bin/go test -coverprofile=/tmp/go-code-cover -timeout 30s github.com/muhfaris/goFun/basic_testing/01_testing

ok  	github.com/muhfaris/goFun/basic_testing/01_testing	0.001s	coverage: 20.0% of statements
Success: Tests passed.

*/
