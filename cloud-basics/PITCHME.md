### Cloud Basics

- Virtualisation 
- Container and Images 
- Container Orchestration 
- "as a Service" (IaaS, PaaS, SaaS ...) 
- Microservices and Design Pattern 
- Pets versus Cattles

---

### Virtualization


<img src="https://raw.githubusercontent.com/stefanhansatos/gitpitch-template/GCP_Atos_101/assets/image/containers.png" alt="Container" height="380"/>

---

### Container and Images

[Docker Hub](https://hub.docker.com/search?q=&type=image)

[Google Container Registry](https://console.cloud.google.com/gcr/images/google-containers/GLOBAL)

[Private Container Registry](https://console.cloud.google.com/gcr/images/aqueous-cargo-242610?project=aqueous-cargo-242610&authuser=0)



[from scratch](https://hub.docker.com/_/scratch/), 
[base](https://github.com/GoogleContainerTools/base-images-docker),
and [distroless](https://github.com/GoogleContainerTools/distroless) images

+++

###

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

### Dockerfile for distroless image

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

---

### Container Orchestration Standard (k8s)

<img src="https://raw.githubusercontent.com/stefanhansatos/gitpitch-template/GCP_Atos_101/assets/image/k8s.png" alt="k8s" height="380"/>

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
    app: MyApp
  ports:
    - protocol: TCP
      port: 80
      targetPort: 9376
```
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


