---
aliases:
    - /projects/active/zondscan/
    - /projects/archived/zondscan/
availability: live
capabilities:
    - explorer
    - analytics
    - faucet
categories:
    - network-operations
data_updated_at: "2026-08-25"
description: A QRL Zond network explorer for browsing blocks, transactions, smart contracts, validators, addresses, and chain activity.
display_status: Beta · Testnet
features:
    - Real-time QRL Zond blockchain data exploration
    - Transaction, block, smart contract, and validator views
    - Address details, balance checking, richlist, and unit conversion tools
    - Next.js frontend with a Go-powered backend API
    - Responsive web interface for desktop and mobile devices
gallery:
    - type: image
      path: zondscan/network-overview.jpg
      caption: Live Zond network overview with chain statistics, recent blocks, and transactions.
    - type: image
      path: zondscan/latest-blocks.jpg
      caption: Latest synced Zond blocks with hashes, transaction counts, timestamps, and gas usage.
    - type: image
      path: zondscan/validators.jpg
      caption: Validator dashboard with epoch progress, staking totals, status distribution, and network trends.
id: zondscan
keywords:
    - explorer
    - zond
    - validators
links:
    - type: website
      url: https://zondscan.com/
      primary: true
listed_at: "2026-06-09"
logos:
    - path: zondscan/icon.png
      description: ZondScan logo
maintainer_records:
    - name: DigitalGuards
      contact: https://github.com/DigitalGuards/zondscan
maintainers:
    - DigitalGuards
maturity: beta
platforms:
    - web
primary_category: network-operations
primary_link:
    type: website
    url: https://zondscan.com/
    primary: true
primary_url: https://zondscan.com/
project-types:
    - applications
project_type: application
publisher:
    name: DigitalGuards
    url: https://github.com/DigitalGuards/zondscan
publishers:
    - DigitalGuards
qrl_environments:
    - testnet
qrl_generations:
    - "2.0"
qrl_relationship: native
qrl_support:
    - generation: "2.0"
      environments:
        - testnet
repositories:
    - id: main
      role: client
      url: https://github.com/DigitalGuards/zondscan
      license: MIT
secondary_categories: []
source_availability: full
title: ZondScan
url: /projects/zondscan/
---

ZondScan is a DigitalGuards block explorer for the QRL Zond network. It
provides a web interface for exploring blocks, transactions, smart contracts,
validators, addresses, and network activity.

The project is open source and built with a Next.js frontend, a Go backend
API, and a blockchain synchronizer that indexes QRL Zond data into MongoDB
for fast lookup and analysis.
