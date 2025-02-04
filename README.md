# Pins - Decentralized Local Social Networking

[![License: MIT](https://img.shields.io/badge/License-MIT-blue.svg)](https://opensource.org/licenses/MIT)
[![Platform](https://img.shields.io/badge/Platform-iOS%20|%20Android%20|%20Web-blue)](https://flutter.dev)

A privacy-focused, peer-to-peer social platform enabling real-time local event discovery through geospatial interactions.

## 🌟 Features

- **Map-Based Pins**: Drop interactive markers to announce events
- **P2P Architecture**: Direct device-to-device communication (no central servers)
- **Real-Time Alerts**: <200ms latency for nearby event notifications
- **Offline Support**: Mesh networking via Bluetooth/Wi-Fi Direct
- **Privacy First**: End-to-end encrypted messages + IP anonymization
- **Spam Resistance**: Proof-of-Work + Blockchain-style reputation system

## 🛠️ Tech Stack

| Component       | Technology                          | Purpose                          |
|-----------------|-------------------------------------|----------------------------------|
| **Frontend**    | Flutter 3.13                        | Cross-platform UI                |
| **P2P Engine**  | Go 1.21 + libp2p                    | Network layer & cryptography     |
| **Backend**     | Supabase (PostgreSQL, Auth)         | User management & metadata        |
| **Geo-Services**| Redis 7.2 + PostGIS                 | Real-time spatial queries         |
| **Protocol**    | Protocol Buffers + WebRTC           | Efficient binary communication    |

## 📦 Installation

### Prerequisites
- Flutter 3.13+
- Go 1.21+
- Supabase account
- Redis Server 7.2+

