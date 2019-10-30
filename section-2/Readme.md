Section 2

Go Module
- can use resolve import out from GOPATH

```bash
go mod help init
go mod init github.com/user/mypackage
```

Internal
- give access previllage in package, cant to give access previllage to out folder internal.

Pointer
- for save type data on memory

```go
var ptr *int
```
Operator :
- `*` to get value pointer
- `&` to get address pointer

slice = pointer of array

Foreach
```go
for index, item := range slice {
    fmt.Println(index, item)
}
```

struct  = structure data
class = struct + method
interface = method