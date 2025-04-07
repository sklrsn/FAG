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

Loki/Grafana
helm repo add grafana https://grafana.github.io/helm-charts
helm repo update
helm install loki grafana/loki-stack --namespace loki --create-namespace --set grafana.enabled=true
kubectl get pods -n loki
kubectl get secret -n loki loki-grafana -o jsonpath="{.data.admin-password}" | base64 --decode; echo
kubectl port-forward -n loki service/loki-grafana 3000:80

VPA (goldilocks)
kubectl apply -f https://raw.githubusercontent.com/kubernetes/autoscaler/vpa-release-1.0/vertical-pod-autoscaler/deploy/vpa-v1-crd-gen.yaml
kubectl apply -f https://raw.githubusercontent.com/kubernetes/autoscaler/vpa-release-1.0/vertical-pod-autoscaler/deploy/vpa-rbac.yaml
helm repo add fairwinds-stable https://charts.fairwinds.com/stable
helm install --name goldilocks --namespace goldilocks --set 
installVPA=true fairwinds-stable/goldilocks
kubectl -n goldilocks port-forward svc/goldilocks-dashboard 8080:80

#Istio
# Download Istio
brew install istioctl

# Install Istio with the demo profile and Gateway API enabled
istioctl install --set profile=demo --set "components.pilot.k8s.overlays[0].name=gateway-api-support" -y

# Install Gateway API CRDs
kubectl apply -f https://github.com/kubernetes-sigs/gateway-api/releases/download/v1.2.1/standard-install.yaml

# Verify installation
kubectl get pods -n istio-system

kubectl create namespace dev
kubectl label namespace dev istio-injection=enabled

---
#Kiali
helm repo add kiali https://kiali.org/helm-charts
helm install kiali-server kiali/kiali-server -n istio-system --create-namespace --atomic
kubectl apply -f https://raw.githubusercontent.com/istio/istio/release-1.18/samples/addons/prometheus.yaml
kubectl apply -f https://raw.githubusercontent.com/istio/istio/release-1.18/samples/addons/grafana.yaml

```bash
    auth:
      openid: {}
      openshift:
        client_id_prefix: kiali
      strategy: anonymous
```

# HPA/VPA

HPA
kubectl apply -f https://github.com/kubernetes-sigs/metrics-server/releases/latest/download/components.yaml

VPA
kubectl apply -f https://raw.githubusercontent.com/kubernetes/autoscaler/vpa-release-1.0/vertical-pod-autoscaler/deploy/vpa-v1-crd-gen.yaml
kubectl apply -f https://raw.githubusercontent.com/kubernetes/autoscaler/vpa-release-1.0/vertical-pod-autoscaler/deploy/vpa-rbac.yaml