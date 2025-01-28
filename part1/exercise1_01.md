The app is in log-output directory. The image is built and pushed to Docker Hub nghson/log-output.

1. Cluster is created with `k3d cluster create log-output -a 2`
2. Deploy the image to the cluster with `kubectl create deployment log-output-dep --image=nghson/log-output`
3. `kubectl get pods`
| NAME |                             READY |  STATUS  |            RESTARTS |  AGE
 log-output-dep-74474bfcbd-t6f4w  | 0/1 |    ContainerCreating  | 0    |      9s
4. `kubectl logs -f log-output-dep-74474bfcbd-t6f4w`