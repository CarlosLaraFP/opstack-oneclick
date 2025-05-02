# opstack-oneclick
A Go-powered CLI tool that installs a mock OP Stack (simulating op-node, op-geth, and others) onto KinD using Helm with chain-specific config passed via values.yaml — designed for protocol teams who want fast, repeatable deployments.

---

The OP Stack is a modular L2 blockchain framework built by Optimism. It includes:

- op-node: Central service coordinating block production, sync, and communication with L1
- op-geth: Ethereum execution engine (a modified Geth client)
- op-batcher: Bundles L2 transactions and submits them to L1
- op-proposer: Posts L2 state roots to L1 for fraud-proof validation
- fault proofs: Coming soon; enables permissionless fraud detection
- Rollup config: JSON file defining chain params like genesis block, gas limit, etc.

These are all containerized services that can be run in Kubernetes — many OP Stack builders run them this way today.

---

## Highlights

Purpose: “Scaffold for automated OP Stack protocol deployments”

Why this matters: “Protocol velocity, modular config, 1-click onboarding”

OP Stack explainer: “L2 chain composed of op-node, op-geth, batcher…”

Helm value examples: chainId, rpcUrl, enableFaultProofs

Future roadmap: swap mock images with real ones (link to https://github.com/ethereum-optimism/optimism)