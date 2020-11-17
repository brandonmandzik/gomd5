# gomd5
__Author__ - _Brandon Mandzik, 2020_ <br>
_Gomd5_ creates a md5 hash, which is cryptographically broken, and outputs it to the standard-ouput. As input you can choose between two flags:
 - file // specifies a file.
 - url // use an url-string.

__Requieres__: go 1.13 AND darwin system

__Installation__:
cd into the directory. use $ go install and execute in the bin directory with $ go run gocat or call it globally by exporting into your PATH variable.

__Testing__:
For a simple test run $ go test -v inside of the directory.

__Piping__:
You can pipe the gocat output to the gomd5 cli. e.g:
```go
gocat ~/go/src/gomd5/main.go | gomd5
```
