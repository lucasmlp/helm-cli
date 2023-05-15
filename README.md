# Helm CLI

This repository contains a command-line interface (CLI) tool known as Helm CLI. The Helm CLI is a utility for interacting with Helm, the package manager for Kubernetes.

## Considerations About the Requirements

1. Storing a list of Helm charts internally in the application won't work. The reason for this is that this application is a CLI application. It has a short lifecycle and only runs when a command is received. In this way, an internal list wouldn't last for more than one lifecycle.

2. I've added a command to add the repositories to the application. This way, it's not necessary to hardcode the path of the repositories.

## Architecture

The CLI is composed of three layers: the presentation layer (package CLI), service layer (package service), and adapter layer (package adapter). The business rules are contained in the service layer and the adapter layer is responsible for connecting the application with external services and databases.

The adapter packages are injected into the service packages via dependency injection. The same thing happens with the service packages and the CLI package; they are injected into the CLI package via dependency injection. The whole injection process happens in main.go.

The layers are bound to each other by contracts, the interface.go files. They can easily be injected with another implementation that satisfies those contracts.

## Unit tests

I've written unit tests to all methods in the service package and got 100% coverage in all of them but the add repoitory. To cover 100% of that package I would have had to wrap the os and url packages in the adapter packages so as to be able to create mocks for it.
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
- `add [chart name]`: Add a Helm chart. The application looks into the registered repositories for the chart. If it is found in the remote repository, the application downloads it to a /charts folder.
- `install [chart name] [release name]`: Installs a Helm chart. If it's a local chart, unpack it and edit the valus.yml before running the install command. If it's a remote chart, the add command pulled the chart and unpacked it into the ./charts folder. Edit the values.yml before running the install command.
- `index`: Add a Helm repository index in the current directory.
- `images`: List container images in all charts added.

To run the Helm CLI, execute the `helm-cli` command followed by the desired subcommand. For example:

```shell
helm-cli repo-add my-repo https://my-repo-url
```

```shell
helm-cli add my-chart
```

```shell
helm-cli install my-chart my-release
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

I've provided a testing folder that contains scripts and Kubernetes manifests to facilitate validation of the application.

### Manifests
1. `mongodb.yml`: Spins up a simple MongoDB deployment and service without authentication.
2. `pod.yml`: Spins up an Ubuntu pod to be used for testing the application. It also contains a service account, cluster role, and cluster role binding. The cluster role has full access to the cluster, which is not recommended. The permission depends on which resources the Helm chart installation will generate.

### Shell Scripts
1. `copy-binary.sh`: This script copies the Go binary into the home folder of the pod used for testing.
2. `copy-local-charts.sh`: This script copies a folder containing local Helm charts to the local-charts folder in the home path of the testing pod.

### Testing Instructions
1. Spin up the testing pod and MongoDB by applying the manifests provided in the testing folder.

```shell
kubectl apply -f ./testing/manifests/pod.yml
kubectl apply -f ./testing/manifests/mongodb.yml
```

2. Generate a binary compatible with the Linux platform:
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

6. Add a local Helm repository:

```shell
./helm-cli repo-add local ./local-charts
```

7. Add a remote Helm repository:

```shell
./helm-cli repo-add bitnami https://charts.bitnami.com/bitnami
```

8. Add a local chart:

```shell
./helm-cli add <name of your chart>
```

9. Add a remote chart:

```shell
./helm-cli add mysql
```

10. Install a remote chart:

```shell
./helm-cli install mysql mysql-dev
```

11. Install a local chart:

```shell
./helm-cli install <chart name> <release name>
```

12. Generate an index with all charts added:

```shell
./helm-cli index
```

13. Print all container images used in all charts that were added:

```shell
./helm-cli images
```

## License

This project is licensed under the [MIT License](LICENSE).

## Acknowledgments

Special thanks to the Helm project for providing the foundation and inspiration for this Helm CLI tool.

For more information about Helm, visit the official [Helm website](https://helm.sh/).
