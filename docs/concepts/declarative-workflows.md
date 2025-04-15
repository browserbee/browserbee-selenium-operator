# Declarative Workflows

Declarative testing with the **BrowserBee Selenium Operator** lets you define browser workflows using Kubernetes-native YAML — based on the [Selenium Workflow Specification](https://browserbee.io/selenium-workflow-specification.json). This approach treats browser tests as infrastructure, making them versioned, composable, and automation-ready.

---

## What Is Declarative Selenium Workflow?

Declarative workflows is the practice of defining **what** to test (e.g., click this, assert that) in a structured way, instead of writing imperative test code in Python or Java. You express the desired workflow in a `SeleniumWorkflow` YAML Custom Resource, and the operator handles browser automation behind the scenes.

---

## Why Use Declarative Selenium Tests?

- **No test scripts required** — YAML does the work.
- **CI/CD and GitOps friendly** — version, diff, and review tests in pull requests.
- **Portable** — define tests once, run them on any Kubernetes cluster.
- **Reusable patterns** — easily compose common UI flows.

---

## 🔧 Specification Overview

A `SeleniumWorkflow` is defined using the following fields:

### `webDriver` (required)

| Field        | Type     | Description                                 |
|--------------|----------|---------------------------------------------|
| `browser`    | string   | One of: `chrome`, `firefox`, `edge`, `safari` |
| `headless`   | boolean  | Run in headless mode                         |
| `implicitWait` | int    | Wait time (seconds) for finding elements     |
| `remoteURL`  | string   | Full WebDriver URL (e.g. from Selenium Grid) |
| `windowSize` | object   | Pixel width and height of browser window     |
| `capabilities` | object | Custom capabilities passed to the driver     |

---

### `actions` (required)

A list of test steps to execute in order.

| Field        | Required | Description |
|--------------|----------|-------------|
| `name`       | ✅        | Unique name of the action |
| `action`     | ✅        | One of: `navigate`, `input`, `click`, `assert`, `screenshot` |
| `target`     | 🚫 (only for `navigate`) | URL to open |
| `selector`   | 🚫 (for `click`, `input`, `assert`) | How to locate elements |
| `value`      | 🚫 (for `input`, `assert`) | Value to send or verify |
| `expected`   | 🚫 (for `assert`) | Expected value |
| `condition`  | 🚫 (for `assert`) | `equals`, `contains`, `notEquals`, `exists` |

---

## ✅ Full YAML Example

```yaml
apiVersion: browserbee.io/v1
kind: SeleniumWorkflow
metadata:
  name: login-flow
spec:
  description: "Test login and dashboard redirect"
  webDriver:
    browser: chrome
    headless: true
    implicitWait: 5
    remoteURL: http://selenium-hub.selenium:4444/wd/hub
    windowSize:
      width: 1280
      height: 800
    capabilities:
      timezone: "UTC"
  actions:
    - name: open-login
      action: navigate
      target: "https://example.com/login"

    - name: input-username
      action: input
      selector:
        strategy: id
        value: username
      value: "testuser"

    - name: input-password
      action: input
      selector:
        strategy: id
        value: password
      value: "secure123"

    - name: click-submit
      action: click
      selector:
        strategy: css
        value: "#submit-login"

    - name: assert-dashboard
      action: assert
      selector:
        strategy: css
        value: "h1.dashboard"
      expected: "Welcome"
      condition: contains

    - name: take-screenshot
      action: screenshot
