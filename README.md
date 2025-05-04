# opstack-oneclick
Go-powered CLI + Helm scaffolding tool for deploying a mock OP Stack to Kubernetes — simulating how protocol developers would spin up their own L2 chain on the Superchain.

> Built as a proof-of-concept to explore how platform engineering principles — such as one-click workflows, Helm-based modularity, and Kubernetes-native automation — can accelerate OP Stack protocol development.

---

## Purpose

This project demonstrates what it would feel like for a protocol team to deploy an OP Stack-based L2 chain with 1 command, using Kubernetes-native tooling under the hood. While the services are mocked (or partially functional), the deployment flow mirrors how Base, Zora, or Mode might bootstrap their own OP Chain. In the real Superchain ecosystem, protocol teams are responsible for deploying and managing their own OP Chains using the OP Stack. This project:

- Abstracts away infra complexity with a single CLI command
- Demonstrates modular, reproducible configuration of OP Stack parameters
- Simulates onboarding DX for external protocol teams
- Models decentralized Superchain participation

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

- Go CLI powered by Cobra
- deploy, status, and destroy commands
- Configurable chainId, rpcUrl, and namespace
- Real OP Stack container image (op-node)
- Helm chart with injected values and ConfigMap-mounted rollup.json
- Health checks (readiness/liveness probes) built in
- Local testing with KinD
- Namespace isolation for multi-chain simulation

---

## Technologies
| Layer              | Tools Used                                        |
| ------------------ | ------------------------------------------------- |
| CLI UX             | Go + Cobra                                        |
| Infra Backend      | Helm + Kubernetes                                 |
| Local Runtime      | KinD (Kubernetes in Docker)                       |
| OP Stack Component | `ghcr.io/ethereum-optimism/op-node:latest`        |
| Configurable       | Chain ID, RPC URL, Namespace, Rollup config       |
| Observability      | Readiness/liveness probes, Prometheus annotations |

---

## Quickstart

```bash
make kind
make build
make deploy
make status
make destroy
```

---

## Roadmap

- Add support for full OP Stack topology: batcher, proposer, etc.
- Replace mock RPC URLs with actual op-geth deployment
- Chain-aware values.yaml templates (e.g., Base vs Zora)
- Add metrics aggregation via Prometheus + Grafana
- Implement chain upgrade workflow (upgrade command)

---

## Why This Project?
OP Labs aims to support protocol teams deploying and operating their own OP Chains. This project explores how to:

- Reduce the barrier to entry for experimentation
- Standardize operational deployment patterns
- Model what "great infra" looks like for protocol developers

This PoC is intended to spark discussion, ideas, and feedback from the Optimism community and infra engineers building on the OP Stack.
