
.PHONY: gover
gover:
	@tput bold setaf 4;go version;tput sgr0

.PHONY: clean
clean:
	rm -f ./log*.txt
	rm -f ./tempfile*
	rm -f ./cover.html cover.out coverage.txt

.PHONY: fmt
fmt:
# to install it:
# go install mvdan.cc/gofumpt@latest
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
