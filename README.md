# QRL Ecosystem Index

A community-contributed index of projects building on QRL 2.0.

## Overview

This repository is the **data layer** for the QRL ecosystem, providing a structured, version-controlled, community-maintained index of projects built on QRL 2.0 smart contracts.

## Agent-Friendly Access

The published site at [www.qrlecosystem.com](https://www.qrlecosystem.com/) provides several machine-readable entry points:

- [`/llms.txt`](https://www.qrlecosystem.com/llms.txt) - Curated site context and links following the [llms.txt proposal](https://llmstxt.org/)
- `/index.html.md` and `/<path>/index.html.md` - Clean Markdown alternatives for public content pages
- [`/index.json`](https://www.qrlecosystem.com/index.json) - Structured project summary data
- [`/sitemap.xml`](https://www.qrlecosystem.com/sitemap.xml) - Canonical HTML page inventory
- [`/robots.txt`](https://www.qrlecosystem.com/robots.txt) - Crawl policy and sitemap location

After generating the project pages and building the site, validate these outputs from the repository root:

```sh
python3 scripts/validate_agent_outputs.py
```

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

## Local Development

Install the Go dependencies once, then generate the Hugo project pages, JSON index, local media assets, and 1200×630 social preview cards before starting the development server:

```sh
go mod download
go run scripts/generate.go
hugo server --source website --bind 127.0.0.1 --port 1313
```

Open [http://127.0.0.1:1313](http://127.0.0.1:1313) to preview the site. Run `go run scripts/generate.go` again whenever project YAML, logos, or gallery images change; the generator rebuilds each project's Open Graph card and Hugo handles template, content, and style changes while the server is running.

To run a production-style build locally:

```sh
hugo --source website --gc --minify --cleanDestinationDir --forceSyncStatic
```

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
