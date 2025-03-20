$ helm repo add ingress-nginx https://kubernetes.github.io/ingress-nginx
$ helm update
$ helm install ingress-controller ingress-nginx/ingress-nginx

openssl req -x509 -newkey rsa:4096 -keyout certs/nginx.key -out certs/nginx.crt -days 365 -nodes -subj "/CN=www.sklrsn.in" -addext "subjectAltName=DNS:sklrsn.in,DNS:www.sklrsn.in"
kubectl create secret tls nginx-tls-secret --cert=certs/nginx.crt --key=certs/nginx.key -n fag

Argo CD

kubectl create namespace argocd
kubectl apply -n argocd -f https://raw.githubusercontent.com/argoproj/argo-cd/stable/manifests/install.yaml

kubectl get secret argocd-initial-admin-secret -n argocd -o jsonpath="{.data.password}" | base64 -d

kubectl port-forward service/argocd-server -n argocd 8443:443