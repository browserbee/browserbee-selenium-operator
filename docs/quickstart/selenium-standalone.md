# 🧪 Selenium Standalone

**Selenium Standalone** is a self-contained Selenium instance that runs both the Hub and Node in a single container. It’s perfect for quick testing, local development, or small-scale automation workflows.

> ⚙️ Powered by the **BrowserBee Selenium Operator**, Standalone deployments are Kubernetes-native, declarative, and scalable.

---

## 🎯 Why Use Standalone Mode?

| Scenario                       | Benefit                                     |
|-------------------------------|---------------------------------------------|
| 🔬 Local testing or CI jobs    | No need to set up a full Selenium Grid      |
| 🧪 Smoke or regression tests   | Fast spin-up and teardown                   |
| 🧱 Simple workflows            | Reduced complexity, especially for demos    |

---

## 🌟 Key Features

- **🎯 All-in-One**: Hub and Node combined into one pod
- **🛠 Easy to Configure**: Declarative setup using CRDs
- **🌍 Cross-Browser Ready**: Works with Chrome, Firefox, and more
- **📈 Scalable**: Increase replicas to run tests in parallel

---

## 🚀 Deploying Selenium Standalone

### 1. Apply a Sample Standalone Resource

```bash
kubectl apply -f https://raw.githubusercontent.com/browserbee/browserbee-selenium-operator/main/config/samples/selenium-standalone_v1_seleniumstandalone.yaml
```

This will create a minimal `SeleniumStandalone` resource and launch a pod with the default browser image.

---

### 2. Verify the Deployment

Check that the instance is running:

```bash
kubectl get seleniumstandalones
```

You should see a resource named `selenium-standalone` or similar.

---

### 3. Access the Console UI

Retrieve the exposed service URL:

```bash
kubectl get svc selenium-standalone -o jsonpath='{.status.loadBalancer.ingress[0].hostname}'
```

Paste the URL in your browser to access the Selenium WebDriver Console.

> 🧪 Using Minikube or Kind? Try `kubectl port-forward` or set up an `Ingress`.

---

## 🛠 Customizing the Standalone Instance

You can customize everything from browser type to replicas and environment variables via the CRD spec.

### Example: Chrome Standalone with 3 Replicas

```yaml
apiVersion: selenium-standalone.browserbee.io/v1
kind: SeleniumStandalone
metadata:
  name: custom-standalone
spec:
  image: selenium/standalone-chrome
  replicas: 3
```

Apply your custom config:

```bash
kubectl apply -f custom-standalone.yaml
```

> 💡 Use the `image` field to swap in Firefox, Edge, or custom driver builds.

---

## 📊 Monitoring & 🔁 Scaling

### 📈 Monitoring

You can integrate Prometheus and Grafana to track:

- Pod health and readiness
- Browser session usage
- Resource consumption per test run

Use Kubernetes-native metrics or expose `/metrics` via sidecars.

---

### 🔼 Scaling the Standalone Instance

Update the number of replicas in your CR:

```yaml
spec:
  replicas: 5
```

The operator will automatically handle scaling the underlying Deployment.

---

## ✅ Best Practices

| Practice | Description |
|----------|-------------|
| **Use a separate namespace** | e.g., `selenium-dev` or `standalone-ci` |
| **Pin versions** | Use a fixed image tag to avoid breaking changes |
| **Update images regularly** | Get the latest WebDriver, browser patches, and fixes |
| **Use limits/requests** | Prevent Standalone pods from consuming too many cluster resources |
| **Expose UI securely** | Add basic auth or use internal-only access for the console |

---

## 🔗 Related Pages

- 🧩 [Selenium Grid Overview](../quickstart/selenium-grid.md)
- ⚙️ [Installation Guide](../quickstart/installation.md)
- 📄 [Standalone CRD Reference](../reference/seleniumstandalone.md)
- 🧪 [Writing Selenium Tests](../quickstart/test-cr.md)

---

> With BrowserBee, setting up a Selenium Standalone is as simple as applying a YAML. Perfect for quick tests, CI runs, or low-overhead test automation.
