
.PHONY: gover
gover:
	@tput bold setaf 4;go version;tput sgr0

.PHONY: clean
clean:
	rm -f ./log*.txt
	rm -f ./tempfile*
	rm -f ./cover.html cover.out coverage.txt

.PHONY: lnt
lnt:
# to install it:
# go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	@tput bold setaf 1;golangci-lint version;tput sgr0
# golangci-lint run -v --enable-all --disable gochecknoglobals --disable paralleltest --disable exhaustivestruct --disable depguard --disable wsl
	golangci-lint run -v --config=golangci.yml


.PHONY: fmt
fmt:
# to install it:
# go install mvdan.cc/gofumpt@latest
	@tput bold setaf 1;echo "gofumpt version:"; gofumpt --version;tput sgr0
	gofumpt -l -w .
	@tput bold setaf 4;echo "gofumpt done";tput sgr0

.PHONY: gci
gci:
# to install it:
# go install github.com/daixiang0/gci@latest
# after that add a place of binaries to $PATH
# export PATH=$PATH:/your path/go/bin
	gci write --skip-generated -s default .
	@tput bold setaf 4;echo "gci done";tput sgr0

.PHONY: gofmt
gofmt:
	gofmt -s -w .
	@tput bold setaf 4;echo "gofmt done";tput sgr0


.PHONY: fix
fix: gover gofmt gci fmt
