ascii-art
=========


To run the tests:
```bash
# install gopherjs and atom bindings
go get -u github.com/gopherjs/gopherjs
go get -u github.com/JohannWeging/atom

# from lib dir
# install the godeps of atom-testing.go
go get "honnef.co/go/js/console"
gopherjs build atom-testing.go
apm link
# reload atom and run the test package
```
At the moment there is no way to run all tests. Just call the tests from the
test function at the bottom.  
