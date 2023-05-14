#!/bin/bash

# Set variables
POD_NAME="terminal-pod"
CONTAINER_NAME="terminal-container"
GO_BINARY_PATH="./helm-cli"  # Replace with the actual path of the Go binary file

# Copy Go binary to the pod
kubectl cp $GO_BINARY_PATH $POD_NAME:/helm-cli -c $CONTAINER_NAME
