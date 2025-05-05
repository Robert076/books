Run it by:

1. Cloning the repo
2. Starting minikube server:
```bash
minikube start
```
3. Running k8s files:
```bash
kubectl apply -f kubernetes/
```
5. Port forwarding from inside minikube:
```bash
kubectl port-forward service/books-service 8080:8080
```
