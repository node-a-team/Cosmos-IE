# Cosmos-IE
![CreatePlan](https://img.shields.io/badge/relase-v1.0.0-red)
![CreatePlan](https://img.shields.io/badge/go-1.14%2B-blue)
![CreatePlan](https://img.shields.io/badge/license-Apache--2.0-green)  
Integrated Exporter for CosmosSDK

## Introduction
This exporter is for monitoring information which is not provided from Tendermint’s basic Prometheus exporter(localhost:26660), and other specific information monitoring purposes

## List of supported chains
Cosmos, Terra, IRISnet, Kava, IOV, E-money

## Install
```bash
mkdir exporter && cd exporter

wget https://github.com/node-a-team/Cosmos-IE/releases/download/v1.0.0/Cosmos-IE.tar.gz  && sha256sum Cosmos-IE.tar.gz | fgrep f010d7f8824c6a0e8573144ef8575f0596ca4c79c49587050e67b617fae97dc2 && tar -xvf Cosmos-IE.tar.gz || echo "Bad Binary!"
```
