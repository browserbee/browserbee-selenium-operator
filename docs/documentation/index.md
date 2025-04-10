# 📚 Documentation

Get the answers you need with Selenium Operator documentation.

Welcome to the Selenium Operator documentation! This site provides comprehensive guides and resources to help you deploy, manage, and automate browser testing using Selenium on Kubernetes. Whether you're just getting started or you're running complex test grids in production, this documentation has everything you need.

Selenium Operator simplifies Kubernetes-native browser automation using a set of Custom Resource Definitions (CRDs) and an Operator pattern:

- **Selenium Operator** manages the full lifecycle of Selenium Grids, including deployment, upgrades, and autoscaling.
- **SeleniumGrid CR** defines how your hub and nodes should be deployed and connected.
- **SeleniumTest CR** allows you to describe and run automated browser tests directly from Kubernetes.

Additionally, the Operator supports flexible strategies like retry logic, wait conditions, and multiple selector types (`id`, `css`, etc.) for your automation steps.

Explore the guides below, or jump directly to the **Overview** for a complete introduction to the operator and its core features.

---

## 📖 Latest Selenium Operator Documentation

Documentation for Selenium Operator **v0.2.1**, supporting Kubernetes-native test automation with CRDs.

- 🔍 [**Overview**](overview.md)  
  Introduction to the operator architecture, components, and how it fits into Kubernetes-native testing.

- 🚀 [**Get Started**](quick-start.md)  
  Follow the quick start guide for installation using Helm and deploying your first test.

- ⚙️ [**Deploying and Managing**](deployment.md)  
  Step-by-step guide to installing and managing Selenium Grids in various environments.

- 📘 [**API Reference**](reference.md)  
  Detailed schema and configuration options for the `SeleniumGrid` and `SeleniumTest` CRDs.

- 🧪 [**Example Custom Resources**](examples.md)  
  Ready-to-use manifests for different grid topologies, test patterns, and execution strategies.

- 🌉 [**Using the Grid API**](grid-api.md)  
  Interact directly with the Selenium Hub for advanced automation workflows.

---

## 🔄 View Guides in Alternative Formats

- 📄 [Helm Chart README](https://github.com/your-org/selenium-operator/blob/main/helm/README.md)
- 📦 [CRD YAML Definitions](https://github.com/your-org/selenium-operator/blob/main/config/crd)

---

## 🚧 In Development Documentation

Documentation for the upcoming version (`main` branch).

- 🧭 [Overview](overview.md)
- ⚙️ [Deploying and Managing](deployment.md)
- 📘 [API Reference](reference.md)
- 🧪 [Example Custom Resources](examples.md)
- 🌉 [Using the Grid API](grid-api.md)

---

## 🧩 Additional Resources

- 🤝 [**Contribute**](../join-us.md)  
  Learn how to contribute to the Selenium Operator and its documentation.

- 🕓 [**Documentation Archive**](archive.md)  
  Access previous versions of the Selenium Operator docs.

- 🎥 [**Demos & Presentations**](demos.md)  
  Watch recorded tutorials and walkthroughs of Selenium Operator in action.

---

### Project Name

**Selenium Operator** brings browser automation to Kubernetes by managing Selenium Grid deployments and tests through declarative APIs.

<div style="text-align: center; font-size: 0.9em; margin-top: 2rem;">
  © Selenium Operator Authors 2025 | Documentation distributed under CC-BY-4.0  
  <br/>
  Not affiliated with the official SeleniumHQ project.
</div>
