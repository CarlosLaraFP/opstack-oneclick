.PHONY: kind deploy

kind:
	kind create cluster --name opstack

deploy:
	helm upgrade --install opstack ./helm/opstack --namespace opstack --create-namespace
