# shortcuts

## scripts

### batch-git-pull

GOOS=darwin GOARCH=amd64 go build -o dist/batch-git-pull-mac-darwin-amd64 scripts/batch-git-pull/main.go
GOOS=linux GOARCH=amd64 go build -o dist/batch-git-pull-linux-amd64 scripts/batch-git-pull/main.go

### change-git-user

GOOS=darwin GOARCH=amd64 go build -o dist/change-git-user-mac-darwin-amd64 scripts/change-git-user/main.go
GOOS=linux GOARCH=amd64 go build -o dist/change-git-user-linux-amd64 scripts/change-git-user/main.go
GOOS=windows GOARCH=amd64 go build -o dist/change-git-user-windows-amd64 scripts/change-git-user/main.go

### transfer-files-to-json

GOOS=darwin GOARCH=amd64 go build -o dist/transfer-files-to-json-mac-darwin-amd64 scripts/transfer-files-to-json/main.go
GOOS=linux GOARCH=amd64 go build -o dist/transfer-files-to-json-linux-amd64 scripts/transfer-files-to-json/main.go

## cmd

### startup

GOOS=darwin GOARCH=amd64 go build -o dist/startup-mac-darwin-amd64 cmd/startup/main.go

### startupnode

GOOS=linux GOARCH=amd64 go build -o dist/startupnode-linux-amd64 cmd/startupnode/main.go

## Check Memory Usage 

# https://linuxhint.com/check_memory_usage_process_linux/
ps -o pid,user,%mem,command ax | sort -b -k3 -r

## Supervisor

brew install supervisor

"""
To restart supervisor after an upgrade:
  brew services restart supervisor
Or, if you don't want/need a background service you can just run:
  /usr/local/opt/supervisor/bin/supervisord -c /usr/local/etc/supervisord.conf --nodaemon

/usr/local/opt/supervisor/bin/supervisord -v

cat /usr/local/etc/supervisord.conf

[include]
files = /usr/local/etc/supervisor.d/*.ini

/Users/mazey/Desktop/supervisor/b-startup.ini
"""

## Nginx

"""
where nginx
brew services restart supervisor
brew services list
nginx -t
"""

## Shadowsocks

# https://github.com/shadowsocks/go-shadowsocks2/releases
./dist/shadowsocks2-macos-amd64 -s "ss://CHACHA20-IETF-POLY1305:your-password@:8488" -verbose

./dist/shadowsocks2-macos-amd64 -s 'ss://CHACHA20-IETF-POLY1305:your-password@:8488' -verbose

# verbose
./dist/shadowsocks2-macos-amd64 -s "ss://CHACHA20-IETF-POLY1305:server.password@:1129" -verbose

# clear
./dist/shadowsocks2-macos-amd64 -s "ss://CHACHA20-IETF-POLY1305:server.password@:1129"