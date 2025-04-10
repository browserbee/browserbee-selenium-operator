
# 📦 Downloads

Get the latest **Selenium Operator** releases and components to streamline browser automation on Kubernetes.

---

## 🧠 Selenium Operator

Deploy and manage Selenium Grids and test workflows using Kubernetes-native Custom Resources.

- 🔖 **Latest Release**: [v0.2.1](https://github.com/browserbee/selenium-operator/releases/latest)
- 📦 [View All Releases →](https://github.com/browserbee/selenium-operator/releases)
- 📖 [Installation Guide](../installation/index.md)

---

## 🚀 Selenium Grid Components

The operator automatically deploys the official Selenium Grid components based on your `SeleniumGrid` custom resource.

| Component      | Version | Download Link |
|----------------|---------|----------------|
| Selenium Hub   | 4.20.0  | [Docker Hub](https://hub.docker.com/r/selenium/hub) |
| Chrome Node    | 4.20.0  | [Docker Hub](https://hub.docker.com/r/selenium/node-chrome) |
| Firefox Node   | 4.20.0  | [Docker Hub](https://hub.docker.com/r/selenium/node-firefox) |
| Edge Node      | 4.20.0  | [Docker Hub](https://hub.docker.com/r/selenium/node-edge) |

---

## 🧪 Example Manifests

Get started quickly with the provided examples:

- 🧾 [Selenium Grid CR Example](https://github.com/browserbee/browserbee-selenium-operator/blob/main/examples/selenium-grid.yaml)
- 🧾 [Selenium Test CR Example](https://github.com/browserbee/browserbee-selenium-operator/blob/main/examples/selenium-test.yaml)

Apply them with:

```bash
kubectl apply -f examples/selenium-grid.yaml
kubectl apply -f examples/selenium-test.yaml
```

---

## 🔗 Related Resources

- 🧬 [GitHub Repository](https://github.com/browserbee/browserbee-selenium-operator)
- 🧭 [Changelog](https://github.com/browserbee/browserbee-selenium-operator/releases)

```
