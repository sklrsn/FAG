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

helm repo add grafana https://grafana.github.io/helm-charts
helm repo update

helm install loki grafana/loki-stack --namespace loki --create-namespace --set grafana.enabled=true
kubectl get pods -n loki

kubectl get secret -n loki loki-grafana -o jsonpath="{.data.admin-password}" | base64 --decode; echo
kubectl port-forward -n loki service/loki-grafana 3000:80