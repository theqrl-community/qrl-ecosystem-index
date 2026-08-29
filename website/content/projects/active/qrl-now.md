---
aliases:
    - /projects/active/qrl-now/
    - /projects/archived/qrl-now/
availability: live
capabilities:
    - analytics
categories:
    - finance
data_updated_at: "2026-08-23"
description: Real-time QRL market dashboard aggregating QRL/USDT prices, OHLCV history, volume, and spreads across exchanges, with volume-weighted composite views and responsive charts.
display_status: Beta
features:
    - Real-time QRL/USDT prices, bid and ask quotes, spreads, and 24-hour market statistics
    - Volume-ranked coverage of MEXC, LBank, Biconomy, XT, and Dex-Trade
    - Multi-exchange price lines, composite OHLC, market spread, and exchange OHLC comparison views
    - Historical OHLCV data across common trading timeframes with incremental older-data loading
    - Volume-weighted composite pricing, per-exchange sparklines, and resilient live updates
gallery:
    - type: image
      path: qrl-now/dashboard.png
      caption: Live dashboard with volume-ranked exchanges, 24-hour sparklines, composite market statistics, and price history controls.
id: qrl-now
keywords:
    - market-data
    - price-tracking
    - ohlcv
    - exchange-comparison
    - real-time
last_verified_at: "2026-08-23"
links:
    - type: website
      url: https://qrl.now
      primary: true
listed_at: "2026-08-23"
logos:
    - path: qrl-now/icon.png
      description: QRL Now logo
maintainer_records:
    - name: Matasx
      contact: https://github.com/Matasx
maintainers:
    - Matasx
maintenance: active
maturity: beta
platforms:
    - web
primary_category: finance
primary_link:
    type: website
    url: https://qrl.now
    primary: true
primary_url: https://qrl.now
project-types:
    - applications
project_launched_at: "2026-08-21"
project_type: application
publisher:
    name: Matasx
    url: https://github.com/Matasx
publishers:
    - Matasx
qrl_environments: []
qrl_generations:
    - 1.x
qrl_relationship: ecosystem-resource
qrl_support:
    - generation: 1.x
      environments: []
repositories: []
secondary_categories: []
source_availability: unavailable
title: QRL Now
url: /projects/qrl-now/
---

QRL Now is a browser-based market dashboard for comparing exchange-listed
QRL pairs. It combines current quotes, 24-hour statistics, OHLCV history,
volume, and spreads from multiple market sources and provides both
per-exchange and composite chart views.

The service consumes public exchange market data and does not query or
depend on a QRL node. Its current directory support metadata references QRL
1.x because that is the QRL asset represented by the tracked spot markets;
support for future QRL generations depends on their exchange listings rather
than a chain-specific application integration.
