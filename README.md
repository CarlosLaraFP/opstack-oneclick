# opstack-oneclick
Go-powered CLI tool that automates the deployment of a mocked OP Stack to a local Kubernetes cluster using Helm. It simulates `op-node`, `op-geth`, and related services, allowing protocol teams to experiment with fast, repeatable, modular deployments.

> 🧠 Built as a proof-of-concept to explore how platform engineering principles — such as one-click workflows, Helm-based modularity, and Kubernetes-native automation — can accelerate OP Stack protocol development.

---

## 🧱 What is the OP Stack?

The **OP Stack** is a modular, open-source framework for building Layer 2 (L2) blockchains on Ethereum. It is developed and maintained by [OP Labs](https://www.optimism.io/) and powers chains like **Base**, **Zora**, and **Mode**.

It includes:

- **`op-node`**: Orchestrates block production and syncing.
- **`op-geth`**: Ethereum execution engine (a fork of Geth).
- **`op-batcher`**: Bundles L2 transactions and submits them to L1.
- **`op-proposer`**: Publishes L2 state roots to L1 for fault-proof validation.
- **Rollup Config**: Defines parameters such as `chain_id`, `block_time`, and genesis config.

The OP Stack's modular design enables scalability, consistency, and shared tooling across chains.

---

## 🎯 Project Goals

- **Developer Experience First**: Make it trivially easy to spin up a test OP Stack environment with sensible defaults.
- **Platform Engineering Principles**: Use Helm and Go to mirror how protocol engineers might interact with infra tooling.
- **Kubernetes-Native Workflow**: Deploy fully into KinD to remain portable, fast, and sandbox-friendly.

---

## 🛠️ Features

- 🧰 **Go CLI**: One-click deployment with `opstack deploy`
- 📦 **Helm Chart**: Parameterized mock of OP Stack components
- ⚙️ **KinD Support**: Lightweight local Kubernetes testing
- 🔧 **Custom Config**: Chain-specific values passed via `values.yaml`
- 📚 **Modular by Design**: Easy to extend with real `op-node`, `op-geth`, etc.

---

## 📄 Helm Configuration

Here are example values passed into the chart:

```yaml
chainId: 420
rpcUrl: https://rpc.testnet.optimism.io
enableFaultProofs: true
opNode:
  image: nginx  # Placeholder for real `op-node`
```
---

## Quicktsart

```bash
# Create KinD cluster
make kind

# Deploy mock OP Stack components
make deploy

# Or use Go CLI
go run main.go deploy
```
---

## Roadmap

- ✅ Replace mock container with real op-node
- ✅ Add support for full OP Stack topology: batcher, proposer, etc.
- 🔜 Chain-aware values.yaml templates (e.g., Base vs Zora)
- 🔜 Prometheus annotations + alerts for key pod signals
- 🔜 GitOps-compatible helmfile or ArgoCD example

---

## Why This Project?
OP Labs aims to support protocol teams deploying and operating their own OP Chains. This project explores how to:

- Reduce the barrier to entry for experimentation
- Standardize operational deployment patterns
- Model what "great infra" looks like for protocol developers

This PoC is intended to spark discussion, ideas, and feedback from the Optimism community and infra engineers building on the OP Stack.
