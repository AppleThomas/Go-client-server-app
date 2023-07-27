# docker push thomas3212/go-client-server-web
# minikube start if port issues

# deletes any old deployment and services first 
kubectl delete deployment go-client-server
kubectl delete service go-client-server


kubectl apply -f k8s/

