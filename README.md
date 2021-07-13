# Cosmos-IE
![CreatePlan](https://img.shields.io/badge/release-v3.0.3-red)
![CreatePlan](https://img.shields.io/badge/go-1.15%2B-blue)
![CreatePlan](https://img.shields.io/badge/license-Apache--2.0-green)  
Integrated Exporter for CosmosSDK

## Introduction
This Prometheus exporter is for monitoring information which is not provided from Tendermint’s basic Prometheus exporter(localhost:26660)

## List of supported chains
Cosmos(cosmoshub-4), IRISnet(irishub-1), BandProtocol-testnet(band-laozi-testnet1), Terra(bombay)

## Install
```bash
cd $HOME
git clone https://github.com/node-a-team/Cosmos-IE.git
cd $HOME/Cosmos-IE

go build

./Cosmos-IE version
## Cosmos-IE v3.0.4
```

## Service(ex: cosmos)
- **--chain** _string_: Chain name of the monitoring node(cosmos | iris | band | terra)
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
journalctl -f -u Cosmos-IE.service
```
