---
aliases:
    - /projects/active/quantascan/
    - /projects/archived/quantascan/
availability: live
capabilities:
    - explorer
    - analytics
categories:
    - network-operations
    - assets-tokenization
data_updated_at: "2026-08-02"
description: 'Analytics-first block explorer for QRL and QRL 2.0: aggregated time-series, cross-chain charts, rich lists, wallet classification, and validator and mining dashboards — built around network-wide data, not just single block and transaction lookups.'
display_status: Stable · Mainnet + Testnet
features:
    - Dual-generation coverage — QRL 1.x mainnet and QRL 2.0 testnet data
    - Aggregated time-series and analytics charts across the whole network
    - Rich list, wallet distribution, and wallet entity classification
    - Validator and beacon-chain dashboards for QRL 2.0
    - Mining pool dominance and reward analytics for QRL
    - Cross-chain unified account history and supply timelines
    - Free public REST API with usage-metered API keys
gallery:
    - type: image
      path: quantascan/homepage.png
      caption: Homepage with live QRL network health — block height, transactions, supply and emission at a glance.
    - type: image
      path: quantascan/charts.png
      caption: Charts hub — on-chain time-series grouped by subject, from daily transactions and blocks to fees and adoption.
    - type: image
      path: quantascan/rich-list.png
      caption: QRL rich list ranking the largest holders by on-chain balance, with signature scheme and 30-day net flow.
    - type: image
      path: quantascan/exchange-flow.png
      caption: Exchange flows — how much QRL sits on exchanges and whether it is moving in or out, per exchange.
id: quantascan
keywords:
    - explorer
    - analytics
    - dashboards
    - charts
    - rich-list
    - validators
    - mining
    - statistics
    - portfolio
links:
    - type: website
      url: https://quantascan.io
      primary: true
    - type: documentation
      url: https://quantascan.io/api-docs
listed_at: "2026-08-02"
logos:
    - path: quantascan/icon.svg
      description: Quantascan logo
maintainer_records:
    - name: 12remember
      contact: https://quantascan.io
maintainers:
    - 12remember
maturity: stable
platforms:
    - web
primary_category: network-operations
primary_link:
    type: website
    url: https://quantascan.io
    primary: true
primary_url: https://quantascan.io
project-types:
    - applications
project_type: application
publisher:
    name: 12remember
    url: https://quantascan.io
publishers:
    - 12remember
qrl_environments:
    - mainnet
    - testnet
qrl_generations:
    - 1.x
    - "2.0"
qrl_relationship: native
qrl_support:
    - generation: 1.x
      environments:
        - mainnet
        - testnet
    - generation: "2.0"
      environments:
        - testnet
repositories: []
secondary_categories:
    - assets-tokenization
source_availability: unavailable
title: Quantascan
url: /projects/quantascan/
---

Quantascan is an analytics-first platform for the QRL ecosystem, covering
both QRL (the Proof-of-Work legacy chain) and QRL 2.0 (the Proof-of-Stake,
EVM-compatible successor). Where a traditional explorer answers "what is in
this block?", Quantascan is built around network-wide questions: supply and
inflation over time, address distribution and rich lists, wallet
classification, exchange flows, mining-pool dominance, and validator
performance.

A full explorer UI (blocks, transactions, and address pages) is included for
operators who need it, but the emphasis is on aggregated data, time-series,
and cross-chain analysis. Legacy QRL data remains permanently browsable as a
read-only archive alongside the QRL 2.0 chain, with cross-chain views that
keep the two chains clearly distinguished.

All data is served through a free public REST API, and an educational
"Learn" section documents QRL concepts, post-quantum signatures, and the
QRL → QRL 2.0 migration model.
