SHELL=PATH='$(PATH)' /bin/sh

.PHONY: gen
.PHONY: abi

all: abi gen

gen:
	abigen --abi Stamp.abi --pkg contract --type Stamp --out Stamp.go

abi:
	solc --abi --overwrite stamp.sol -o ./

clean:
	rm -f *.go
