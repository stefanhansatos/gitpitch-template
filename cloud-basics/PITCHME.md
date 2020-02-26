### Cloud-Native Basics

- "Pets vs Cattle" 

- Container, Images, and Registries 

- Container Orchestration 

- Services

- Advantages and Features

---
### Pets vs Cattle

- Pets are unique, <br> need intense care, <br> and scale up vertically

- Cattle are almost identical, <br> replaceable, <br> and scale out horizontally

---
<img src="https://raw.githubusercontent.com/stefanhansatos/gitpitch-template/GCP_Atos_101/assets/image/enterprise-computing-approach.jpg" alt="enterprise computing approach" height="520"/>

### Pets

---
<img src="https://raw.githubusercontent.com/stefanhansatos/gitpitch-template/GCP_Atos_101/assets/image/cloud-computing-approach.jpg" alt="cloud computing approach" height="520"/>

### Cattle
---
### Container


<img src="https://raw.githubusercontent.com/stefanhansatos/gitpitch-template/GCP_Atos_101/assets/image/container-evolution.png" alt="container evolution" height="420"/>

---

### Images and Registries

- an image runs one main process <br>
in a container
- parent images: [from scratch](https://hub.docker.com/_/scratch/), <br> 
[base](https://github.com/GoogleContainerTools/base-images-docker),
and [distroless](https://github.com/GoogleContainerTools/distroless)
- registries: 
[Docker Hub](https://hub.docker.com/search?q=&type=image), <br>
[Google Container Registry](https://console.cloud.google.com/gcr/images/google-containers/GLOBAL)<br>
[Private Container Registry](https://console.cloud.google.com/gcr/images/aqueous-cargo-242610?project=aqueous-cargo-242610&authuser=0)

+++
### Docker: Run image from public repository

```bash
docker run -p 8080:80 -d gcr.io/google-containers/nginx

docker images
docker ps
docker logs -tf [container]

curl localhost:8080
# try webview

docker stop [container]
docker rm [container]
docker rmi [image]
```
+++
### Dockerfile: Multi-stage build <br> for from scratch image

```dockerfile
FROM golang:1.12 as build-env

WORKDIR /go/src/app
ADD . /go/src/app
RUN go get -d -v ./... 
RUN CGO_ENABLED=0 go build -a -ldflags '-s' -o /go/bin/app

FROM scratch
COPY --from=build-env /go/bin/app /
CMD ["/app"]
```

+++
### Dockerfile: Multi-stage build <br> for distroless image

```dockerfile
FROM golang:1.12 as build-env

WORKDIR /go/src/app
ADD . /go/src/app
RUN go get -d -v ./...
RUN go build -o /go/bin/app

FROM gcr.io/distroless/base
COPY --from=build-env /go/bin/app /
CMD ["/app"]
```
+++
### Docker: Build image and <br> push to repository

```bash
gcloud auth configure-docker

docker build \
-t gcr.io/aqueous-cargo-242610/presentation:web.distroless \
-f Dockerfile.distroless .
docker push gcr.io/aqueous-cargo-242610/presentation:web.distroless

docker build \
-t eu.gcr.io/aqueous-cargo-242610/presentation:web.scratch \
-f Dockerfile.scratch .
docker push eu.gcr.io/aqueous-cargo-242610/presentation:web.scratch
```
---

### Container Orchestration (k8s)

<img src="https://raw.githubusercontent.com/stefanhansatos/gitpitch-template/GCP_Atos_101/assets/image/k8s.png" alt="k8s" height="380"/>

---

### Pods, Container, and Volumes

- **Pods** are the core building bricks

- **Containers** run in Pods with shared <br> network namespace and storage

- **Volumes** are the storage of a Pod

---

### Deployments, ReplicaSets and Services
- **Deployments** declare <br> ReplicaSets and Pods

- **ReplicaSets** controll <br> the number of Pods running

- **Services** expose <br> pods as network services

+++ 

<img src="https://raw.githubusercontent.com/stefanhansatos/gitpitch-template/GCP_Atos_101/assets/image/pod-example.png" alt="pod-example" height="420"/>

+++

<img src="https://raw.githubusercontent.com/stefanhansatos/gitpitch-template/GCP_Atos_101/assets/image/deploy-rs-pods-node.png" alt="deploy, rs, pods, node" height="450"/>

+++
### Create GCP Network

```bash
gcloud compute networks create presentation-network \
    --subnet-mode=custom

gcloud compute networks subnets create presentation-subnet-europe-west1 \
    --network=presentation-network \
    --region=europe-west1 \
    --range=10.0.0.0/9
```
+++
### Create GCP Kubernetes Engine

```bash
gcloud container clusters create presentation-cluster \
    --zone europe-west1-b \
    --scopes cloud-platform \
    --network projects/aqueous-cargo-242610/global/networks/presentation-network \
    --subnetwork projects/aqueous-cargo-242610/regions/europe-west1/subnetworks/presentation-subnet-europe-west1

gcloud container clusters get-credentials presentation-cluster \
    --zone europe-west1-b
```

+++

### kubectl investigates cluster

```bash
kubectl cluster-info
kubectl get nodes
kubectl get namespaces

kubectl get all --namespace kube-system

```

+++
### YAML File Deployment

```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: nginx-deployment
  labels:
    app: nginx
spec:
  replicas: 3
  selector:
    matchLabels:
      app: nginx
  template:
    metadata:
      labels:
        app: nginx
    spec:
      containers:
      - name: nginx
        image: gcr.io/google-containers/nginx:1.7.9
        ports:
        - containerPort: 80
```

+++

### YAML File Service

```yaml
apiVersion: v1
kind: Service
metadata:
  name: nginx
spec:
  selector:
    app: nginx
  ports:
  - protocol: TCP
    port: 80
    targetPort: 80 
  type: LoadBalancer
```

+++

### kubectl apply YAML files

```bash
kubectl apply -f deployment.yaml
watch kubectl get all

kubectl apply -f service.yaml
watch kubectl get services

watch kubectl get po,deploy,rs
```

```bash
IP=$(kubectl get svc nginx -n foo \
    -o jsonpath="{.status.loadBalancer.ingress[*].ip}")
    
. 
```


. curl-loop.sh

---

### Services

- **Microservices** follow <br> the single concern principle (SCP) 
<br> implemented as container-based app

- **"as a Service"** provides e.g. <br> 
Infrastructure as a Service (IaaS) <br> 
Platform as a Service (PaaS) <br> 
Software as a Service (SaaS)

---

### Advantages of Cloud-Native

- Immutability 

- Declarative Configuration 

- Online Self-healing Systems 

+++
### Reconciliation Loop (k8s)

<img src="https://raw.githubusercontent.com/stefanhansatos/gitpitch-template/GCP_Atos_101/assets/image/reconciliation-loop.jpeg" alt="Reconciliation Loop (k8s)" height="420"/>


---

### Features of Cloud-Native (k8s)

- Products as a Service

- CI/CD Pipelines 

- REST APIs ([API Explorer](https://developers.google.com/apis-explorer/)) 

- Integration in k8s ([Operator Hub](https://operatorhub.io/)) 

+++ 

### Continuous Integration (CI)

<img src="https://raw.githubusercontent.com/stefanhansatos/gitpitch-template/GCP_Atos_101/assets/image/continuous-integration.png" alt="Continuous Integration (CI)" height="380"/>

+++

### Continuous Deployment (CD)

<img src="https://raw.githubusercontent.com/stefanhansatos/gitpitch-template/GCP_Atos_101/assets/image/continuous-deployment.png" alt="Continuous Deployment (CD)" height="380"/>


