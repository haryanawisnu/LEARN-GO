# Session 1 Introduction

Show terminal in visual code
ctrl + `

Access Modifier in Golang based on name function
- If start with lower case mean PRIVATE
- If start with upper case mean PUBLIC

For Documentation
https://godoc.org/

Standard Code Golang
https://golang.org/cmd/gofmt/

Get library from github
Repo move to $GOPATH/src/github/...repo
example : 
```bash
go get github.com/iancoleman/strcase
```

import library on file
- "github.com/iancoleman/strcase"
- a "github.com/iancoleman/strcase" (alias import)

Resource Go-Lint
https://github.com/golang/go/wiki/CodeReviewComments

func init() will to call if import
if you want to just call init file go 
import like :
```go
_ "LEARN-GO/section-1/test"
```

Type Data
- int
- int32
- int64
- float
- float32
- float64
https://www.callicoder.com/golang-basic-types-operators-type-conversion/
https://golang.org/pkg/fmt/

Another Command
```bash
go run main.go
go run *.go
go build #create binaryexport

GOOS=linux GOARCH=amd64 go build
go tool dist list
go env GOOS #view os env
go env

go install

echo $GOPATH
go help gopath

#Setting GOPATH
cd ~
mkdir -p go/bin go/src/ go/pkg
ls -l go
#By default GOPATH in `~/go`
#Change GOPATH in `~.bashrc` atau `~/.zhrc`
echo 'export GOPATH="NEWPATH"' >> ~/.zshrc
source ~/.zshrc
```