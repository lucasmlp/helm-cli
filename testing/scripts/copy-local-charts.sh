
POD_NAME="terminal-pod"
CONTAINER_NAME="terminal-container"
LOCAL_REPO_PATH=/home/lucasmachado/development/helm-charts

kubectl cp $LOCAL_REPO_PATH $POD_NAME:/local-charts -c $CONTAINER_NAME
