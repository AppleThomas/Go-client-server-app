appPod=$(kubectl get pods -n default -l app=go-client-server --output=jsonpath={.items..metadata.name})

kubectl port-forward $appPod 45287:3000 &

sleep 10

curl -v localhost:45287

echo -e "\ndone"