# go-app

Simple Go HTTP API with endpoints for hello-world messaging, static items, and random user sampling.

## Local Development

1. Install Go 1.22 or newer.
2. Copy `.env.example` to `.env` (create if missing) and set `PORT` when overriding the default `8080`.
3. Run the server:

	```bash
	go run ./...
	```

4. Visit http://localhost:8080 to verify the API.

## Container Image

Build the container locally:

```bash
docker build -t go-app:local .
docker run --rm -p 8080:8080 go-app:local
```

## Kubernetes on AWS EKS

### Prerequisites

- AWS CLI v2 configured with access to the target account.
- kubectl installed and pointed at an EKS cluster.
- An Amazon ECR repository (for example `go-app`).
- IAM role configured for GitHub Actions OIDC federation with permissions for ECR (push) and EKS (deploy).

### Manifests

Kubernetes definitions live under `k8s/`:

- `namespace.yaml` provisions the `go-app` namespace.
- `service.yaml` exposes the deployment via a load balancer on port 80.
- `deployment.yaml` manages the application pods and probes. Replace the placeholder image (`000000000000.dkr.ecr.us-east-1.amazonaws.com/go-app:latest`) with the URI for your ECR repository if deploying manually. The GitHub Actions workflow updates the image automatically during deployments.

Apply manually when needed:

```bash
kubectl apply -f k8s/namespace.yaml
kubectl apply -f k8s/service.yaml
kubectl apply -f k8s/deployment.yaml
```

### GitHub Actions Workflow

The workflow at `.github/workflows/cicd.yml` builds the Docker image, pushes it to ECR, and deploys the manifests to EKS on every push to the `main` branch or when triggered manually.

Required GitHub secrets:

- `AWS_ACCOUNT_ID` (optional if encoded in the assume-role policy but handy for documentation).
- `AWS_REGION` – target AWS region (for example `us-east-1`).
- `AWS_ROLE_TO_ASSUME` – ARN of the IAM role the workflow should assume.
- `ECR_REPOSITORY` – ECR repository name (for example `go-app`).
- `EKS_CLUSTER_NAME` – name of the target EKS cluster.

The workflow uses the commit SHA as the Docker image tag and issues `kubectl rollout status` to confirm deployment success.

### Manual ECR Push (optional fallback)

```bash
aws ecr get-login-password --region $AWS_REGION | docker login --username AWS --password-stdin "$AWS_ACCOUNT_ID.dkr.ecr.$AWS_REGION.amazonaws.com"
docker build -t "$AWS_ACCOUNT_ID.dkr.ecr.$AWS_REGION.amazonaws.com/$ECR_REPOSITORY:latest" .
docker push "$AWS_ACCOUNT_ID.dkr.ecr.$AWS_REGION.amazonaws.com/$ECR_REPOSITORY:latest"
kubectl set image deployment/go-app go-app="$AWS_ACCOUNT_ID.dkr.ecr.$AWS_REGION.amazonaws.com/$ECR_REPOSITORY:latest" --namespace go-app
```

## Observability

- Liveness/readiness probes hit `/`.
- Logs stream to stdout; aggregate them via `kubectl logs -f deployment/go-app -n go-app` or EKS logging integrations.
