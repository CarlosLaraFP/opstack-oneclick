kind:
	kind create cluster --name opstack

build:
	go build -o opstack main.go

deploy:
	./opstack deploy --namespace opstack --chain-id 999 --rpc-url http://localhost:8545

status:
	./opstack status --namespace opstack

clean:
	kind delete cluster
