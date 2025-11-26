.PHONY: help cluster-up cluster-down cluster-status

CLUSTER_NAME ?= kind-cluster
KUBECONFIG ?= $(HOME)/.kube/config
CLUSTER_CONFIG ?= ./config.yaml
CLUSTER_NODES_CONFIG ?= ./nodes-config.yaml

help:
	@echo "Kind cluster management"
	@echo "  make cluster-up      - Create and start a kind cluster"
	@echo "  make cluster-down    - Delete the kind cluster"
	@echo "  make cluster-status  - Show cluster status"

cluster-up:
	@echo "Creating kind cluster: $(CLUSTER_NAME)"
	kind create cluster --name $(CLUSTER_NAME) --config $(CLUSTER_CONFIG)
	@echo "Cluster created successfully"
	@kubectl cluster-info --context kind-$(CLUSTER_NAME)

cluster-down:
	@echo "Deleting kind cluster: $(CLUSTER_NAME)"
	kind delete cluster --name $(CLUSTER_NAME)
	@echo "Cluster deleted successfully"

cluster-status:
	@echo "Cluster status for: $(CLUSTER_NAME)"
	@kubectl cluster-info --context kind-$(CLUSTER_NAME) || echo "Cluster not found"
	@kubectl get nodes --context kind-$(CLUSTER_NAME) 2>/dev/null || echo "No nodes available"