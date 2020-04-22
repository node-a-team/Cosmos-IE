# Cosmos-IE
![CreatePlan](https://img.shields.io/badge/release-v1.0.0-red)
![CreatePlan](https://img.shields.io/badge/go-1.14%2B-blue)
![CreatePlan](https://img.shields.io/badge/license-Apache--2.0-green)  
Integrated Exporter for CosmosSDK

## Introduction
This exporter is for monitoring information which is not provided from Tendermint’s basic Prometheus exporter(localhost:26660), and other specific information monitoring purposes

## List of supported chains
Cosmos, Terra, IRISnet, Kava, IOV, E-money

## Install
```bash
mkdir Cosmos-IE && cd Cosmos-IE

wget https://github.com/node-a-team/Cosmos-IE/releases/download/v1.0.0/Cosmos-IE.tar.gz  && sha256sum Cosmos-IE.tar.gz | fgrep be26403ca8dd1dd19f95ac7f652acbe604de72fb30a01b576cf82ba44f78ba8f && tar -xvf Cosmos-IE.tar.gz || echo "Bad Binary!"
```

## Service
```bash
# Make log directory & file
sudo mkdir /var/log/userLog  &&
sudo touch /var/log/userLog/Cosmos-IE.log  &&
sudo chown ${USER}:${USER} /var/log/userLog/Cosmos-IE.log

# Path: /data/cosmos/Cosmos-IE
sudo tee /etc/systemd/system/Cosmos-IE.service > /dev/null <<EOF
[Unit]
Description=Integrated Exporter for CosmosSDK
After=network-online.target

[Service]
User=${USER}
WorkingDirectory=/data/cosmos/Cosmos-IE
ExecStart=/data/cosmos/Cosmos-IE/Cosmos-IE run --chain cosmos --oper-addr cosmosvaloper14l0fp639yudfl46zauvv8rkzjgd4u0zk2aseys
StandardOutput=file:/var/log/userLog/Cosmos-IE.log
StandardError=file:/var/log/userLog/Cosmos-IE.log
Restart=always
RestartSec=3

[Install]
WantedBy=multi-user.target
EOF

sudo systemctl enable Cosmos-IE.service
sudo systemctl start Cosmos-IE.service


## log
tail -f /var/log/userLog/Cosmos-IE.log
```
