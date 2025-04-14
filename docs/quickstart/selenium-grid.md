# Selenium Grid with BrowserBee

Selenium Grid is a core component of the **BrowserBee Selenium Operator**, enabling scalable, cross-browser test execution inside Kubernetes. It allows you to distribute Selenium tests across multiple containers (called nodes) managed by a central hub.

> With BrowserBee, Selenium Grid becomes fully declarative, GitOps-friendly, and cloud-native.

---

## Key Features

- **Parallel Testing**: Run many tests simultaneously across distributed browser nodes.
- **Cross-Browser Coverage**: Support for Chrome, Firefox, and Edge (with more coming soon).
- **Dynamic Scaling**: Add or remove nodes based on load using simple CR edits.
- **Centralized Hub**: A controller (hub) routes all test requests to matching nodes.

---

## Deploying a Selenium Grid

### 1. Apply the Sample Grid

Create a simple Grid with 1 hub and 2 Chrome nodes:

```bash
kubectl apply -f https://raw.githubusercontent.com/browserbee/browserbee-selenium-operator/main/config/samples/selenium-grid_v1_seleniumgrid.yaml
```

> This YAML is a `SeleniumGrid` CRD managed by the operator.

---

### 2. Check Grid Status

Verify the grid has been created successfully:

```bash
kubectl get seleniumgrids
```

You should see a CR with status fields indicating readiness (depending on operator version).

---

### 3. Access the Selenium Grid Console

To use the Selenium Hub:

```bash
kubectl get svc selenium-hub -o jsonpath='{.status.loadBalancer.ingress[0].hostname}'
```

Open the resulting URL in your browser to access the Selenium Console UI.

> 💡 Tip: If you're using a local cluster like Kind or Minikube, you may need to expose the hub using a `NodePort` or `Ingress`.

---

## Customize Your Grid

You can fully configure the grid by modifying the CR’s `spec` section. Below is a multi-browser setup example:

```yaml
apiVersion: selenium-grid.browserbee.io/v1
kind: SeleniumGrid
metadata:
  name: cross-browser-grid
spec:
  hub:
    replicas: 1
  nodes:
    - browser: chrome
      replicas: 3
    - browser: firefox
      replicas: 2
    - browser: edge
      replicas: 1
```

Apply your changes:

```bash
kubectl apply -f custom-grid.yaml
```

> Each node runs the appropriate WebDriver image with the right browser preinstalled.

---

## Monitoring & Scaling

### Monitoring

Integrate with Prometheus + Grafana for real-time visibility into:

- Pod health
- Browser session usage
- Response latency

BrowserBee emits standard Kubernetes metrics for all Grid components.

### Scaling the Grid

Adjust node or hub replicas by editing the `SeleniumGrid` spec:

```yaml
spec:
  nodes:
    - browser: chrome
      replicas: 10
```

Apply and the operator will handle the scaling automatically.

---

## Best Practices

| Tip | Description |
|-----|-------------|
| **Use a dedicated namespace** | Keep Grid resources isolated (e.g., `browserbee-testing`) |
| **Update browser images regularly** | Stay current with browser features and security patches |
| **Monitor usage trends** | Use Grafana dashboards to right-size node counts |
| **Pin compatible image versions** | Ensure version consistency across test runs |
| **Use selectors or labels** | Group grids by environment, test suite, or CI branch |

---

## 🔗 Related Resources

- [Installing the Operator](../quickstart/installation.md)
- [Declarative Test Workflows](../quickstart/selenium-workflow.md)

---

> With BrowserBee, launching a multi-browser Selenium Grid is as simple as editing a YAML file.
