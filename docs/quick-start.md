# Quick Start

## Requirements

* Installed kubectl command-line tool.
* Installed kind command-line tool.

## 1. Create a Kubernetes Cluster

```shell
kind create cluster --name selenium-cluster
```

## 2. Install BrowserBee Selenium Operator

```shell
kubectl apply -f https://raw.githubusercontent.com/browserbee/browserbee-selenium-operator/main/manifests/install.yaml
```

This will create new namespace, `browserbee-selenium-operator-system`, where BrowserBee services and application resources will live.

## 3. Deploy a Selenium Grid

```shell
kubectl apply -f https://raw.githubusercontent.com/browserbee/browserbee-selenium-operator/main/manifests/selenium-grid/install.yaml
```