# browserbee-selenium-operator
To provide a Kubernetes-native way to define, run, and scale Selenium Grid.

// **Use:** Install BrowserBee in your Kubernetes cluster and manage Selenium Grids via GitOps workflows.

## Description
BrowserBee Selenium Operator is a cloud-agnostic, Kubernetes-native platform for automating the lifecycle of Selenium Grid infrastructure. Designed with GitOps principles in mind, it allows developers and QA teams to declaratively define Selenium Grids for Kubernetes workflows and CI/CD pipelines. The project integrates with key tools in the CNCF landscape, including ArgoCD for GitOps workflow and Prometheus for observability. By bringing modern cloud-native principles to the world of Selenium, BrowserBee bridges the gap between infrastructure as code and end-to-end quality assurance in highly scalable, multi-cloud environments.

## Getting Started

### Prerequisites
- go version v1.21.0+
- docker version 17.03+.
- kubectl version v1.11.3+.
- Access to a Kubernetes v1.11.3+ cluster.

### To Deploy on the cluster
**Build and push your image to the location specified by `IMG`:**

```sh
make docker-build docker-push IMG=<some-registry>/browserbee-selenium-operator:tag
```

**NOTE:** This image ought to be published in the personal registry you specified.
And it is required to have access to pull the image from the working environment.
Make sure you have the proper permission to the registry if the above commands don’t work.

**Install the CRDs into the cluster:**

```sh
make install
```

**Deploy the Manager to the cluster with the image specified by `IMG`:**

```sh
make deploy IMG=<some-registry>/browserbee-selenium-operator:tag
```

> **NOTE**: If you encounter RBAC errors, you may need to grant yourself cluster-admin
privileges or be logged in as admin.

**Create instances of your solution**
You can apply the samples (examples) from the config/sample:

```sh
kubectl apply -k config/samples/
```

>**NOTE**: Ensure that the samples has default values to test it out.

### To Uninstall
**Delete the instances (CRs) from the cluster:**

```sh
kubectl delete -k config/samples/
```

**Delete the APIs(CRDs) from the cluster:**

```sh
make uninstall
```

**UnDeploy the controller from the cluster:**

```sh
make undeploy
```

## Project Distribution

Following are the steps to build the installer and distribute this project to users.

1. Build the installer for the image built and published in the registry:

```sh
make build-installer IMG=<some-registry>/browserbee-selenium-operator:tag
```

NOTE: The makefile target mentioned above generates an 'install.yaml'
file in the dist directory. This file contains all the resources built
with Kustomize, which are necessary to install this project without
its dependencies.

2. Using the installer

Users can just run kubectl apply -f <URL for YAML BUNDLE> to install the project, i.e.:

```sh
kubectl apply -f https://raw.githubusercontent.com/<org>/browserbee-selenium-operator/<tag or branch>/dist/install.yaml
```

## Contributing

We welcome contributions to the BrowserBee Selenium Operator project! Here's how you can get involved:

1. **Report Issues**: If you encounter bugs or have feature requests, please open an issue in the GitHub repository.

2. **Submit Pull Requests**: If you'd like to contribute code, fork the repository, create a new branch, and submit a pull request. Please ensure your code adheres to the project's coding standards and includes appropriate tests.

3. **Improve Documentation**: Help us improve the documentation by submitting updates or corrections.

4. **Join Discussions**: Participate in discussions on issues and pull requests to help shape the future of the project.

### Development Setup

To set up your development environment:

1. Clone the repository:
   ```sh
   git clone git@github.com:browserbee/browserbee-selenium-operator.git
   ```

2. Navigate to the project directory:
   ```sh
   cd browserbee-selenium-operator
   ```

3. Install dependencies:
   ```sh
   go mod tidy
   ```

4. Run tests to ensure everything is working:
   ```sh
   make test
   ```

### Coding Guidelines

- Follow the [Go Code Review Comments](https://github.com/golang/go/wiki/CodeReviewComments).
- Write clear and concise commit messages.
- Ensure all new code is covered by tests.

### Code of Conduct

Please note that this project is governed by a [Code of Conduct](CODE_OF_CONDUCT.md). By participating, you are expected to uphold this code.

**NOTE:** Run `make help` for more information on all potential `make` targets

More information can be found via the [Kubebuilder Documentation](https://book.kubebuilder.io/introduction.html)

## License

Copyright 2025.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.

