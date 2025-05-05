## 🚀 Running the Project on macOS with Minikube

Follow these steps to get the app running locally using Minikube:

1. 📦 Clone the repository:  
   `git clone https://github.com/Robert076/books && cd books`

2. 🚜 Start the Minikube cluster:  
   `minikube start`

3. ⚙️ Apply Kubernetes configuration files:  
   `kubectl apply -f kubernetes/`

4. 🌐 Expose the service via port forwarding:  
   `kubectl port-forward service/books-service 8080:8080`

✅ Now you can access the service at: [http://localhost:8080](http://localhost:8080)
