ascii-art
=========

This is a port of the https://github.com/atom/ascii-art
to the Go programming language with the help of gopherjs.

To run the example:
```bash
# install gopherjs and atom bindings
go get -u github.com/gopherjs/gopherjs
go get -u github.com/JohannWeging/atom

cd lib
# install the godeps of ascii-art.go
go get github.com/CrowdSurge/banner
gopherjs build ascii-art.go
apm link
# reload atom and run the Ascii Art: Convert command on slected text
```
