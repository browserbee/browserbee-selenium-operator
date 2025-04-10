# Deployment

This section describes how to deploy the operator into a Kubernetes cluster using the packaged release artifacts.

## Prerequisites

Before deploying, ensure the following:

- You have access to a Kubernetes cluster (v1.22+ recommended).
- `kubectl` is configured to point to your target cluster.
- Custom Resource Definitions (CRDs) must be installed prior to applying the operator.
- The container image is available on Quay at `quay.io/<your-org>/<your-operator>:<version>`.

## Quickstart

1. **Download the latest release**

You can find the release YAMLs under [GitHub Releases](https://github.com/<your-org>/<your-repo>/releases). Each release includes:

- `operator.yaml` – All required manifests bundled together
- `crds/` – Individual CRDs, if you want to install them separately

2. **Apply the manifests**

```bash
kubectl apply -f https://github.com/<your-org>/<your-repo>/releases/download/vX.Y.Z/operator.yaml
```

Or, if installing CRDs separately:

```bash
kubectl apply -f https://github.com/<your-org>/<your-repo>/releases/download/vX.Y.Z/crds/
kubectl apply -f https://github.com/<your-org>/<your-repo>/releases/download/vX.Y.Z/operator.yaml
```

3. **Verify the deployment**

```bash
kubectl get pods -n <operator-namespace>
kubectl get deployments -n <operator-namespace>
```

4. **Apply a Custom Resource**

After deployment, you can apply your custom resource to start using the operator:

```bash
kubectl apply -f config/samples/<sample-cr>.yaml
```

## Image Configuration

The operator image is hosted on Quay:

```
quay.io/<your-org>/<your-operator>:vX.Y.Z
```

You can customize the image tag or repository in the deployment YAML or Helm values.
f you’d like a Helm-flavored version or Kustomize-based instructions instead. Want me to auto-generate `operator.yaml` packaging as part of a GitHub Action too?