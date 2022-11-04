

#enter to pods
kubectl exec -it task-pv-pod -- /bin/bash

#scale
kubectl scale --replicas=3 deployment/myotherwebserver

#view events
kubectl get events --sort-by='.metadata.creationTimestamp' -A

#view configs
kubectl config view

#clean up
kubectl delete pod task-pv-pod
kubectl delete pvc task-pv-claim
kubectl delete pv task-pv-volume

#delete a pv or pvc stuck
kubectl patch pv jhooq-pv -p '{"metadata":{"finalizers":null}}'
kubectl patch pvc jhooq-pv-claim -p '{"metadata":{"finalizers":null}}'

#nginx
https://docs.nginx.com/nginx-ingress-controller/installation/installation-with-manifests/

#keycloak
kubectl create -f https://raw.githubusercontent.com/keycloak/keycloak-quickstarts/latest/kubernetes-examples/keycloak.yaml