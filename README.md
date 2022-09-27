# go-gin-gee

## Install

```
go get

# or
go get github.com/bitfield/script
```

If `i/o timeout`, run the command to replace the proxy: 

```
go env -w GOPROXY=https://goproxy.cn
```

## Run

```
go run scripts/init/main.go

go run cmd/api/main.go

# or
go run scripts/ChangeGitUser.go
```

Visit: `http://127.0.0.1:3000/api/ping`

## Build

### Default

```
go build cmd/api/main.go
```

### Linux

```
GOOS=linux GOARCH=amd64 go build cmd/api/main.go
```

### Mac

```
GOOS=darwin GOARCH=amd64 go build cmd/api/main.go

# or
GOOS=darwin GOARCH=amd64 go build -o dist/change-git-user-mac scripts/change-git-user/main.go

GOOS=darwin GOARCH=amd64 go build -o dist/init scripts/init/main.go
```

### Windows

```
GOOS=windows GOARCH=amd64 go build cmd/api/main.go
```
