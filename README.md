Run it by: (macOS, minikube)

1. Cloning the repo
```bash
git clone https://github.com/Robert076/books/blob/main/README.md
```
2. Starting minikube server:
```bash
minikube start
```
3. Running k8s files:
```bash
kubectl apply -f kubernetes/
```
4. Port forwarding from inside minikube:
```bash
kubectl port-forward service/books-service 8080:8080
```
