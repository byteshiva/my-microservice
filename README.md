
# Go Microservice with Gin, Docker, GHCR, and Kubernetes

This is a scalable microservice built with Go and the Gin framework. The project demonstrates how to:
- Build a RESTful API with Gin.
- Dockerize the application with a multi-stage build.
- Push the Docker image to GitHub Container Registry (GHCR).
- Deploy the microservice to Kubernetes with deployment and service configurations.

## Features

- **Health Check**: A simple `/health` endpoint to check if the service is up.
- **Greeting**: A `/greet/:name` endpoint that returns a personalized greeting message.
- **Dockerized**: The app is containerized using a multi-stage Docker build.
- **Kubernetes Ready**: Configurations for deploying the service in a Kubernetes cluster.

## Prerequisites

Before you begin, make sure you have the following installed:

- [Go](https://golang.org/dl/) (version 1.18 or later)
- [Docker](https://www.docker.com/get-started)
- [Kubernetes](https://kubernetes.io/docs/setup/)
- [kubectl](https://kubernetes.io/docs/tasks/tools/)
- [GitHub Personal Access Token](https://github.com/settings/tokens) (for pushing to GHCR)

## Project Setup

### 1. Clone the repository

Clone the repository to your local machine:

```bash
git clone https://github.com/yourusername/my-microservice.git
cd my-microservice
```

### 2. Install Go dependencies

Run the following command to install the necessary dependencies:

```bash
go mod tidy
```

### 3. Running the Application Locally

To run the application locally, execute the following command:

```bash
go run cmd/myservice/main.go
```

The service will be available at `http://localhost:8080`. You can test the following endpoints:

- **Health Check**: `curl http://localhost:8080/health`
- **Greeting**: `curl http://localhost:8080/greet/Go`

### 4. Dockerizing the Application

To build the Docker image, use the following command:

```bash
docker build -t my-microservice .
```

To run the container:

```bash
docker run -p 8080:8080 my-microservice
```

The application will be available at `http://localhost:8080`.

### 5. Pushing Docker Image to GitHub Container Registry (GHCR)

To push your Docker image to GitHub's container registry:

#### Step 1: Tag the Docker image

```bash
docker tag my-microservice ghcr.io/yourusername/my-microservice:latest
```

#### Step 2: Log in to GHCR

Log in to GitHub's container registry with your Personal Access Token (PAT):

```bash
echo "<your_personal_access_token>" | docker login ghcr.io -u <your_github_username> --password-stdin
```

#### Step 3: Push the image

Push the Docker image to GHCR:

```bash
docker push ghcr.io/yourusername/my-microservice:latest
```

### 6. Deploying to Kubernetes

To deploy the application to Kubernetes, follow these steps:

#### Step 1: Update `deployment.yaml`

In the `deployment.yaml` file, update the `image` field with the correct GHCR path:

```yaml
image: ghcr.io/yourusername/my-microservice:latest
```

#### Step 2: Create a Kubernetes Secret for Private GHCR Access

If your Docker image is private, you need to create a Kubernetes secret to allow your Kubernetes cluster to pull the image. Follow these steps:

1. **Create a Personal Access Token (PAT)** on GitHub with `read:packages` and `repo` scopes.
2. **Create the Docker registry secret** using the following command:

   ```bash
   kubectl create secret docker-registry ghcr-secret \
     --docker-server=ghcr.io \
     --docker-username=<your_github_username> \
     --docker-password=<your_personal_access_token> \
     --docker-email=<your_email>
   ```

   Replace `<your_github_username>`, `<your_personal_access_token>`, and `<your_email>` with your GitHub username, the generated PAT, and your email address.

#### Step 3: Reference the Secret in `deployment.yaml`

Add the `imagePullSecrets` field to your `deployment.yaml` so Kubernetes can use the created secret to access the private image:

```yaml
spec:
  containers:
    - name: my-microservice
      image: ghcr.io/yourusername/my-microservice:latest
      ports:
        - containerPort: 8080
  imagePullSecrets:
    - name: ghcr-secret
```

#### Step 4: Apply the Deployment and Service

Run the following commands to apply the deployment and service configurations:

```bash
kubectl apply -f deployment.yaml
kubectl apply -f service.yaml
```

#### Step 5: Verify the Deployment

Check the status of your deployment:

```bash
kubectl get deployments
kubectl get services
```

Once the service has an external IP (if using a cloud provider), you can access it at `http://<external-ip>/health`.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Acknowledgments

- [Gin Framework](https://gin-gonic.com/)
- [Docker](https://www.docker.com/)
- [Kubernetes](https://kubernetes.io/)
- [GitHub Container Registry](https://docs.github.com/en/packages/working-with-a-github-packages-registry)
```

This update adds the necessary steps for creating and using Kubernetes secrets for private GitHub Container Registry images. Make sure to replace placeholders like `yourusername`, `your_personal_access_token`, and `your_email` with your actual GitHub credentials.
