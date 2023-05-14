# Helm CLI

This repository contains a command-line interface (CLI) tool called Helm CLI. The Helm CLI is a utility for interacting with Helm, the package manager for Kubernetes.

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
go build ./cmd/main -o helm-cli
```

5. Optionally, move the binary to a directory on your system's `PATH` to make it accessible globally.

## Usage

The Helm CLI provides the following commands:

- `add repo [name] [path]`: Add a local or remote Helm repository.
- `add [chart name]`: Add a Helm chart.
- `install chart [chart name]`: Install a Helm chart.
- `index`: Add a Helm repository index.
- `images`: List container images.

To run the Helm CLI, execute the `cli-app` command followed by the desired subcommand. For example:

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