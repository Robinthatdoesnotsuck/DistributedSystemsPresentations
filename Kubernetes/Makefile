create-cluster:
	kind create cluster --wait 5m
kubeconfig:
	export KUBECONFIG="$(kind get kubeconfig-path)"
delete-cluster:
	kind delete cluster

kube-info:
	kubectl cluster-info

kube-ver:
	kubectl version

kube-components:
	kubectl get componentstatuses

kube-nodes:
	kubectl get nodes

kube-get-node:
	kubectl describe nodes kube#k