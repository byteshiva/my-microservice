# Running the Go Microservice on Minikube

This guide explains how to start a Minikube cluster, deploy the microservice from Docker, and access it via Minikube's service.

## Prerequisites

Before you begin, make sure you have the following installed:

- [Minikube](https://minikube.sigs.k8s.io/docs/)
- [Kubectl](https://kubernetes.io/docs/tasks/tools/install-kubectl/)
- [Docker](https://www.docker.com/get-started)
- [Nix](https://nixos.org/nix/) (for installing Minikube using `nix-shell`)

## Step 1: Start Minikube with Docker

You can start a Minikube cluster with Docker using Nix. Run the following command:

```bash
nix-shell -I nixpkgs=https://github.com/NixOS/nixpkgs/archive/nixos-unstable.tar.gz -p minikube -v
```

This command will download and launch Minikube using Nix, ensuring you have the latest version.

Once Minikube is started, use the following command to verify that your Minikube cluster is running:

```bash
minikube start
```

### Check the Minikube status:

```bash
minikube status
```

## Step 2: Build the Docker Image for Minikube

Now, ensure that Docker is set up to use Minikube’s Docker daemon. Run the following command:

```bash
eval $(minikube docker-env)
```

This will set your Docker environment to use the Minikube Docker daemon instead of your local Docker engine.

Now, you can build your Docker image directly inside Minikube's Docker environment. Run the following command from the root of your project:

```bash
docker build -t my-microservice .
```

## Step 3: Deploy the Microservice to Minikube

Create a Kubernetes deployment using the Docker image you just built. First, create or update the `deployment.yaml` to use your newly built image:

```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: my-microservice
  labels:
    app: my-microservice
spec:
  replicas: 1
  selector:
    matchLabels:
      app: my-microservice
  template:
    metadata:
      labels:
        app: my-microservice
    spec:
      containers:
        - name: my-microservice
          image: my-microservice:latest  # This should match the image name
          ports:
            - containerPort: 8080
```

Create a Kubernetes service to expose the microservice:

```yaml
apiVersion: v1
kind: Service
metadata:
  name: my-microservice
spec:
  selector:
    app: my-microservice
  ports:
    - protocol: TCP
      port: 8080
      targetPort: 8080
  type: LoadBalancer  # This allows access from outside the cluster
```

Now, apply both `deployment.yaml` and `service.yaml`:

```bash
kubectl apply -f deployment.yaml
kubectl apply -f service.yaml
```

### Step 4: Access the Service

Minikube provides a convenient way to access services through a browser. To open your microservice in the browser, run the following:

```bash
minikube service my-microservice
```

This command will automatically open a browser window to the external URL where your microservice is accessible.

Alternatively, if you're running Minikube on a non-cloud environment, use this command to get the URL manually:

```bash
minikube service my-microservice --url
```

You can use this URL to access your microservice at `http://<external-ip>:8080`.

## Step 5: Stop Minikube

When you're done, you can stop your Minikube cluster by running:

```bash
minikube stop
```

## Acknowledgments

- [Minikube](https://minikube.sigs.k8s.io/docs/)
- [Docker](https://www.docker.com/)
- [Kubectl](https://kubernetes.io/docs/tasks/tools/install-kubectl/)
```

This `README.md` guides users on starting a Minikube cluster, building the Docker image using Minikube’s Docker environment, and deploying the microservice in Kubernetes with Minikube. It also includes steps for accessing the service and stopping Minikube.
