# Selenium Workflow

The **Selenium Workflow** defines how to deploy, test, and scale Selenium infrastructure using the **BrowserBee Selenium Operator**. It empowers teams to run browser automation at scale — declaratively, reproducibly, and Kubernetes-native.

---

## Workflow Overview

The Selenium lifecycle consists of 4 key stages:

1. **Define**: Use Kubernetes Custom Resources to describe your desired Selenium infrastructure
2. **Deploy**: Apply CRs using `kubectl`, and the operator will provision everything for you
3. **Execute**: Point your test framework to the Grid or Standalone endpoint and run tests
4. **Monitor & Scale**: Observe usage and scale as needed

---

## Step-by-Step Workflow

### 1. Define Custom Resources

Start by defining the required Selenium components — Grid, Hub, Nodes, or Standalone — in YAML.

#### Example: Selenium Grid

```yaml
apiVersion: selenium-grid.browserbee.io/v1
kind: SeleniumGrid
metadata:
  name: example-grid
spec:
  hub:
    replicas: 1
  nodes:
    - browser: chrome
      replicas: 2
    - browser: firefox
      replicas: 2
```

> The operator reads this resource and provisions all necessary pods, services, and stateful sets.

---

### 2. Deploy Resources to Kubernetes

Use `kubectl` to apply your custom resource to the cluster:

```bash
kubectl apply -f selenium-grid.yaml
```

Then verify it's running:

```bash
kubectl get seleniumgrids
```

If you’re using Standalone:

```bash
kubectl get seleniumstandalones
```

---

### 3. Run Selenium Tests

Once your Selenium infrastructure is live, connect your tests to the exposed WebDriver endpoint.

#### Python Example (Remote WebDriver)

```python
from selenium import webdriver

options = webdriver.ChromeOptions()
driver = webdriver.Remote(
    command_executor='http://<selenium-hub-url>:4444/wd/hub',
    options=options
)

driver.get("https://example.com")
print(driver.title)
driver.quit()
```

> 🧪 You can point to either the Selenium Hub (for Grids) or a Standalone service.

Use `kubectl get svc` to fetch the external or internal URL of your Hub or Standalone:

```bash
kubectl get svc selenium-hub
kubectl get svc selenium-standalone
```

---

### 4. Monitor & Scale

#### Monitor with Prometheus + Grafana

Track:

- Pod health
- Test session concurrency
- Resource utilization
- Node availability

Use native Kubernetes metrics or expose `/metrics` with sidecars.

---

#### Scale with a Simple Edit

Update your CR or use `kubectl scale`:

```yaml
spec:
  hub:
    replicas: 3
```

Apply your changes:

```bash
kubectl apply -f selenium-grid.yaml
```

Or scale manually:

```bash
kubectl scale deployment selenium-hub --replicas=3
```

---

## Best Practices

| Recommendation | Benefit |
|----------------|---------|
| **Use namespaces** | Isolate test environments and manage RBAC |
| **Pin image tags** | Avoid unexpected behavior from image updates |
| **Version your CRs in Git** | Enable GitOps workflows and review changes |
| **Integrate with CI/CD** | Automate test environment provisioning on each commit |
| **Monitor consistently** | Prevent overloading pods or stale sessions |

---

## Compatible Components

You can use any of the following components in your workflow:

| Component         | Use Case                             |
|------------------|--------------------------------------|
| `SeleniumGrid`    | Full grid setup with Hub and Nodes  |
| `SeleniumHub`     | Central hub if using Nodes manually |
| `SeleniumNode`    | Browser-specific workers             |
| `SeleniumStandalone` | Quick testing and CI runs          |
| `SeleniumTest`    | Declarative test execution (YAML-driven) |

---

## Related Resources

- [Declarative Testing](../concepts/declarative-testing.md)
- [Installation Guide](../quickstart/installation.md)
- [SeleniumGrid CRD Reference](../reference/seleniumgrid.md)
- [Create and Run a Test](../quickstart/test-cr.md)

---

> With BrowserBee, your entire Selenium workflow becomes declarative, version-controlled, and seamlessly integrated into Kubernetes and CI/CD.
