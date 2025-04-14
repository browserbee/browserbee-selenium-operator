# Selenium Node

A Selenium Node is a component of Selenium Grid that executes tests. It connects to the Selenium Hub and receives test requests, running them in the specified browser environment. The BrowserBee Selenium Operator simplifies the deployment and management of Selenium Nodes in Kubernetes environments.

## Features of Selenium Node

- **Browser Execution**: Runs tests in browsers like Chrome, Firefox, and Edge.
- **Parallel Execution**: Supports running multiple tests simultaneously.
- **Dynamic Scaling**: Easily scale the number of nodes to match testing requirements.
- **Integration with Hub**: Seamlessly integrates with the Selenium Hub for test routing.

## Deploying a Selenium Node

Follow these steps to deploy a Selenium Node using the BrowserBee Selenium Operator:

1. **Apply the Selenium Node Custom Resource**:

```bash
kubectl apply -f https://raw.githubusercontent.com/browserbee/browserbee-selenium-operator/main/config/samples/selenium-node_v1_seleniumnode.yaml
```

2. **Verify the Deployment**:
   Check the status of the Selenium Node:

```bash
kubectl get seleniumnodes
```

3. **Check Node Connectivity**:
   Ensure the Selenium Node is connected to the Selenium Hub:

```bash
kubectl logs <node-pod-name>
```

Look for logs indicating successful connection to the hub.

## Customizing the Selenium Node

You can customize the Selenium Node by modifying the Custom Resource Definition (CRD). For example, you can specify the browser type, number of replicas, and other configurations in the `spec` section of the CRD.

Example:

```yaml
apiVersion: selenium-node.browserbee.io/v1
kind: SeleniumNode
metadata:
  name: custom-node
spec:
  hubRef:
    name: selenium-hub
    namespace: default
  browser: chrome
  replicas: 3
```

Apply the updated configuration:

```bash
kubectl apply -f custom-node.yaml
```

## Monitoring and Scaling

- **Monitoring**: Use Prometheus and Grafana to monitor the performance and health of your Selenium Nodes.
- **Scaling**: Adjust the number of replicas to scale your nodes up or down based on demand.

## Best Practices

- Use a dedicated namespace for your Selenium Node resources.
- Monitor resource usage to ensure optimal performance.
- Regularly update your Selenium Node image to include the latest features and security patches.
