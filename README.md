# QRL Ecosystem Index

A community-contributed index of projects building on QRL 2.0.

## Overview

This repository is the **data layer** for the QRL ecosystem, providing a structured, version-controlled, community-maintained index of projects built on QRL 2.0 smart contracts.

## Project Status

| Stage | Description |
|---|---|
| `development` | Projects that are still being built, tested, or prepared for release |
| `production` | Projects that are live and intended for public use |
| `archived` | Projects that are no longer active |

Note: Both `development` and `production` projects are placed in the `projects/active/` directory. Network deployment is tracked separately in each project type-specific block.

## Submitting a Project

See [CONTRIBUTING.md](CONTRIBUTING.md) for full details.

Quick steps:

1. Fork this repository
2. Copy `projects/template.yaml` to `projects/active/`
3. Name the file using your project's `id` field (e.g., `my-project.yaml`)
4. Fill out all required fields
5. Open a Pull Request

## Repository Structure

```
qrl-ecosystem/
├── README.md
├── CONTRIBUTING.md
├── schema/
│   └── project.schema.json
├── projects/
│   ├── template.yaml
│   ├── active/
│   └── archived/
└── .github/
    ├── workflows/
    └── PULL_REQUEST_TEMPLATE.md
```

## Categories

Projects are tagged with one of:

- `defi` - Decentralized finance
- `nft` - NFT-related projects
- `wallet` - Wallets and key management
- `explorer` - Block explorers
- `infrastructure` - Infrastructure services
- `tooling` - Developer tools
- `dao` - DAO governance
- `gaming` - Gaming applications
- `identity` - Identity solutions
- `oracle` - Oracle services
- `bridge` - Cross-chain bridges
- `social` - Social platforms
- `educational` - Educational resources
- `news` - News and media
