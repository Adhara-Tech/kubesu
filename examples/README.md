# Kubesu example

1. Spin up a local cluster:
   ```
   kind create cluster
   ```

2. Apply example manifests
   ```
   kubectl apply -f .
   ```

3. Force a Job execution
   ```
   kubectl create job --from=cronjob/connect-nodes connect-nodes-manual01
   ```
