# Cosmos-IE
![CreatePlan](https://img.shields.io/badge/release-v1.0.0-red)
![CreatePlan](https://img.shields.io/badge/go-1.14%2B-blue)
![CreatePlan](https://img.shields.io/badge/license-Apache--2.0-green)  
Integrated Exporter for CosmosSDK

## Introduction
This Prometheus exporter is for monitoring information which is not provided from Tendermint’s basic Prometheus exporter(localhost:26660)

## List of supported chains
Cosmos, Terra, IRISnet, Kava, E-money, BandProtocol, IOV

## Docs
[Cosmos-IE(Integrated Exporter for CosmosSDK)](https://www.notion.so/wlsaud619/Cosmos-IE-Integrated-Exporter-for-CosmosSDK-1e9c6cf1bdb0483180829676b533565b)

## Install
```bash
mkdir Cosmos-IE && cd Cosmos-IE

wget https://github.com/node-a-team/Cosmos-IE/releases/download/v1.1.0/Cosmos-IE.tar.gz  && sha256sum Cosmos-IE.tar.gz | fgrep 78911047f3fab4c862589f995d0790b99ad32ecd5a19125a6012e6b85b2b8378 && tar -xvf Cosmos-IE.tar.gz || echo "Bad Binary!"
```

## Service(ex: cosmos)
```bash
## Create a systemd service
sudo tee /etc/systemd/system/Cosmos-IE.service > /dev/null <<EOF
[Unit]
Description=Integrated Exporter for CosmosSDK
After=network-online.target

[Service]
User=${USER}
ExecStart=$HOME/Cosmos-IE/Cosmos-IE run \
  --chain "cosmos" \
  --oper-addr "cosmosvaloper14l0fp639yudfl46zauvv8rkzjgd4u0zk2aseys"
Restart=always
RestartSec=3
StandardOutput=syslog
StandardError=syslog
SyslogIdentifier=Cosmos-IE

[Install]
WantedBy=multi-user.target
EOF

## Start service
sudo systemctl enable Cosmos-IE
sudo systemctl start Cosmos-IE

## log
journalctl -f | grep Cosmos-IE
```
