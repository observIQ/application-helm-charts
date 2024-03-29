# ObservIQ Couchbase CE chart

This chart deploys a StatefulSet with 3 replicas that are running Couchbase Community Edition.

> NOTE: To run this, we recommend starting Docker/k8s with 4 cores and 8 gigabytes.

## Creating a cluster

1. Forward port `8091` on the first pod in the statefulset by running the following command:
```
kubectl -n <YOUR_NAMESPACE> port-forward pod/couchbase-community-0 8091
```

2. Run the following command to get the IP addresses of the individual pods. Note the IP addresses for the second and third pods. 
```
kubectl -n <YOUR_NAMESPACE> get pod -o wide
```

3. Using your browser of choice, visit `localhost:8091` and you will see the login page for the node. The username and password will be `Administrator:password` for all nodes.

4. After a successful login, go to the **Servers** page
5. Click on **Add Server**
6. Using the information collected in **step 2**, put in the IP address of one of the nodes we want to add. Use the `Administrator:password` credentials, leave the remaining defaults. Click **Add server**
7. Click on **Rebalance**
> NOTE: You can add both nodes and rebalance a single time.
8. Repeat **steps 5-7** for the remaining node.
