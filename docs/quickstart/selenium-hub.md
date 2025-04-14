# Selenium Hub

The Selenium Hub is the central component of a Selenium Grid. It acts as a server that manages test requests and routes them to the appropriate nodes for execution. The BrowserBee Selenium Operator simplifies the deployment and management of Selenium Hubs in Kubernetes environments.

## Features of Selenium Hub

- **Centralized Management**: Acts as the single point of contact for all test requests.
- **Dynamic Routing**: Routes test requests to the appropriate nodes based on browser and platform requirements.
- **Scalability**: Supports scaling to handle increased test loads.
- **Monitoring**: Provides a web interface to monitor the status of the grid.

## Deploying a Selenium Hub

Follow these steps to deploy a Selenium Hub using the BrowserBee Selenium Operator:

1. **Apply the Selenium Hub Custom Resource**:

    ```bash
    kubectl apply -f https://raw.githubusercontent.com/browserbee/browserbee-selenium-operator/main/config/samples/selenium-hub_v1_seleniumhub.yaml
    ```

2. **Verify the Deployment**:
   Check the status of the Selenium Hub:

   ```bash
   kubectl get seleniumhubs
   ```

3. **Access the Selenium Hub**:
    Retrieve the service URL for the Selenium Hub:

    ```bash
    kubectl get svc selenium-hub -o jsonpath='{.status.loadBalancer.ingress[0].hostname}'
    ```

    Open the URL in your browser to access the Selenium Hub console.

## Customizing the Selenium Hub

You can customize the Selenium Hub by modifying the Custom Resource Definition (CRD). For example, you can specify the number of replicas and other configurations in the `spec` section of the CRD.

Example:

```yaml
apiVersion: selenium-hub.browserbee.io/v1
kind: SeleniumHub
metadata:
  name: custom-hub
spec:
  replicas: 2
```

Apply the updated configuration:

```bash
kubectl apply -f custom-hub.yaml
```

## Monitoring and Scaling

- **Monitoring**: Use Prometheus and Grafana to monitor the performance and health of your Selenium Hub.
- **Scaling**: Adjust the number of replicas to scale your hub up or down based on demand.

## Best Practices

- Use a dedicated namespace for your Selenium Hub resources.
- Monitor resource usage to ensure optimal performance.
- Regularly update your Selenium Hub image to include the latest features and security patches.
