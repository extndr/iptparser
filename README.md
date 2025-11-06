# ipt-parser

A lightweight tool to quickly view DNAT rules from iptables in a human-readable way  
Quickly see which external ports are forwarded to internal addresses on Linux servers.

## Features

- Extracts DNAT rules from iptables `nat` table.
- Displays rules in a clear and readable format (`<DPort> → <Destination>`).
- Can be monitored in real-time using `watch`.

## Installation

Clone the repository and build the binary:

```bash
git clone https://github.com/extndr/ipt-parser.git
cd ipt-parser
go build -o ipt-parser ./cmd
```

## Usage

Run as root to see DNAT rules:

```bash
sudo ./ipt-parser
```

Example output:

```bash
80 → 192.168.1.10:80
443 → 192.168.1.20:443
```
