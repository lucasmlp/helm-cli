# Helm CLI

This repository contains a command-line interface (CLI) tool called Helm CLI. The Helm CLI is a utility for interacting with Helm, the package manager for Kubernetes.

## Architecture
The CLI is composed of three layers. The presentation layer (package CLI), service layer (package service) and adapter layer(package adapter). The business rules are contained in the service layer. The adapter layer is responsible to conect the application with external services and databases.

The adapter packages are injected into the service packages via dependency injection. The same thing happens with the service packages and the cli package. They are injected into the CLI package via dependency injection. The whole injection process happens in main.go.

The layers are bound to each other by contracts, the interface.go files. They can be easily be injected with another implementation that satisfies those contracts.

## Installation

To install and use the Helm CLI, follow these steps:

1. Ensure that you have Go installed on your system.
2. Clone this repository to your local machine:

```shell
git clone https://github.com/lucasmlp/helm-cli.git
```

3. Change to the project directory:

```shell
cd helm-cli
```

4. Build the Helm CLI binary:

```shell
go build -o helm-cli ./cmd/main.go
```

5. Optionally, move the binary to a directory on your system's `PATH` to make it accessible globally.

## Usage

The Helm CLI provides the following commands:

- `repo-add [name] [path]`: Add a local or remote Helm repository.
- `add [chart name]`: Add a Helm chart. The application looks into the registered repositories for the chart. If it is found in the remote repository, the application downloads to a /charts folder.
- `install [chart name] [release name]`: Install a Helm chart.
- `index`: Add a Helm repository index in the current directory.
- `images`: List container images in all charts added.

To run the Helm CLI, execute the `helm-cli` command followed by the desired subcommand. For example:

```shell
helm-cli add repo my-repo https://my-repo-url
```

```shell
helm-cli add my-chart
```

```shell
helm-cli install chart my-chart-name
```

```shell
helm-cli index
```

```shell
helm-cli images
```

Note: Replace `helm-cli` with the appropriate command or binary name when running the Helm CLI.

For more information on each command and its usage, you can also run the `--help` flag or refer to the source code in the `run.go` file.

## Testing

I've provided a testing folder that contains scripts and Kubernetes manifests to make it easy to validate the application.

### Manifests
1. mongodb.yml: spins up a simple mongodb deployment and service without authentication
2. pod.yml: spins up an ubuntu pod that'll be used for testing the application. It also contains a service account, cluster role and cluster role binding. The cluster role has full access to the cluster and this is not a good thing to do. The permission depends on which resources the helm chart installation will generate.

### Shell scripts
1. copy-binary.sh: this script copies the go binary into the home folder of the pod used for testing.
2. copy-local-charts.sh: this script copies a folder containing local helm charts to the local-charts folder in the home path of the testing pod.


### Testing instructions
1. Spin up the testing pod and mongodb by applying the manifests provided in the testing folder.

```shell
kubectl apply -f ./testing/manifests/pod.yml
kubectl apply -f ./testing/manifests/mongodb.yml
```

2. Generate a binary compatible with linux platform
```shell
make build
```

3. Run the script to copy the binary to the pod:

```shell
chmod +x ./testing/scripts/copy-binary.sh
./testing/scripts/copy-binary.sh
```

4. Modify the script to copy local charts to the testing pod and run it:

```shell
chmod +x ./testing/scripts/copy-local-charts.sh
./testing/scripts/copy-local-charts.sh
```

5. Execute bash inside the pod:

```shell
kubectl exec -t -i terminal-pod -- /bin/bash
```

6. Add a local helm repository
```shell
./helm-cli repo-add local ./local-charts
```

7. Add a remote helm repository
```shell
./helm-cli repo-add bitnami https://charts.bitnami.com/bitnami
```

8. Add a local chart
```shell
./helm-cli add <name of your chart>
```

9. Add a remote chart
```shell
./helm-cli add mysql
```

10. Installing a remote chart
```shell
./helm-cli install mysql mysql-dev
```

11. Installing a local chart
```shell
./helm-cli install <chart name> <release name>
```

12. Generating an index with all charts added
```shell
./helm-cli index
```

13. Printing all container images used in all charts that were added
```shell
./helm-cli images
```

## Contributing

Contributions to the Helm CLI project are welcome! If you'd like to contribute, please follow these steps:

1. Fork the repository on GitHub.
2. Create a new branch for your feature or bug fix.
3. Make your changes, ensuring that you follow the project's coding standards.
4. Write tests for your changes, if applicable.
5. Commit your changes with clear and descriptive commit messages.
6. Push your changes to your forked repository.
7. Submit a pull request to the main repository, explaining the changes you've made.

## License

This project is licensed under the [MIT License](LICENSE).

## Acknowledgments

Special thanks to the Helm project for providing the foundation and inspiration for this Helm CLI tool.

For more information about Helm, visit the official [Helm website](https://helm.sh/).