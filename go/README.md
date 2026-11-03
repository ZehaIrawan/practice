### Install Go

https://go.dev/doc/install

```bash
# Download Go binary, check latest https://go.dev/dl/
wget https://go.dev/dl/go1.27.3.linux-amd64.tar.gz

# tar go binary
sudo tar -C /usr/local -xzf go1.27.3.linux-amd64.tar.gz

# add environment variable 
# nano ~/.bashrc
# source ~/.bashrc
export PATH=$PATH:/usr/local/go/bin
export PATH="$PATH:$(go env GOPATH)/bin"


# Check version
go version
```

### 

go run hello.go

go build hello.go
