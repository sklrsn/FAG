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

kubectl apply -f https://raw.githubusercontent.com/kubernetes/autoscaler/vpa-release-1.0/vertical-pod-autoscaler/deploy/vpa-v1-crd-gen.yaml
kubectl apply -f https://raw.githubusercontent.com/kubernetes/autoscaler/vpa-release-1.0/vertical-pod-autoscaler/deploy/vpa-rbac.yaml
helm repo add fairwinds-stable https://charts.fairwinds.com/stable
helm install --name goldilocks --namespace goldilocks --set 
installVPA=true fairwinds-stable/goldilocks
kubectl -n goldilocks port-forward svc/goldilocks-dashboard 8080:80

#Istio
https://istio.io/latest/docs/setup/install/helm/

helm repo add istio https://istio-release.storage.googleapis.com/charts
helm repo update
helm install istio-base istio/base -n istio-system --set defaultRevision=default --create-namespace
helm ls -n istio-system
helm install istiod istio/istiod -n istio-system --wait
helm ls -n istio-system

kubectl create namespace istio-ingress
helm install istio-ingress istio/gateway -n istio-ingress --wait

kubectl apply -f https://github.com/kubernetes-sigs/gateway-api/releases/latest/download/standard-install.yaml
