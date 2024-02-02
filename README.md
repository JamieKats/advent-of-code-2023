Run go file

```go
go run main.go
```

Create go module to organise many files

```go
go mod init github.com/JamieKats/go-helloworld
```

Run after creating module 

```go
go mod tidy
```

Compile go module.

```go
docker build --target bin --output bin/ .
```

Build image with a name

```
docker build -t gotest .
```


#### Resources
https://www.cloudwithchris.com/blog/go-dev-environment-vscode-wsl/