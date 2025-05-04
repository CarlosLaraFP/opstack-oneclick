test:
	go test ./... -v

kind-install:
	curl -Lo kind https://kind.sigs.k8s.io/dl/v0.20.0/kind-linux-amd64
	chmod +x kind
	sudo mv kind /usr/local/bin/kind

kind: kind-install
	kind create cluster --name opstack

docker:
	docker pull us-docker.pkg.dev/oplabs-tools-artifacts/images/op-node:v1.13.2
	docker tag us-docker.pkg.dev/oplabs-tools-artifacts/images/op-node:v1.13.2 op-node:local

kind-load: docker
	kind load docker-image op-node:local --name opstack

build:
	go build -o opstack main.go

deploy: build
	./opstack deploy --namespace opstack --chain-id 999 --rpc-url http://localhost:8545

status:
	./opstack status --namespace opstack

destroy:
	./opstack destroy --namespace opstack

clean: destroy
	kind delete cluster
