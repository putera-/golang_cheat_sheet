# GOLANG

### Check Version

```sh
% go version
go version go1.22.3 darwin/arm64
```

### Create Module

perintah ini akan menghasilkan file go.mod

```sh
% go mod ini [nama-module]
# example
go mod init user-module
```

### Build / Compile

akan menghasilan file dengan nama file sesuai dengan nama module (nama package di file go.mod)

```sh
% go build
```

### Run

```sh
% ./helloworld
Hello World
```

### Run without build / compile

```sh
% go run helloworld.go
Hello World
```
