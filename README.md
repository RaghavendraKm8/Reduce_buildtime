A) Local Environmentgo

version            # should be 1.22+
go env GOPROXY        # should be https://proxy.golang.org,direct

B) Jenkins

Install Go 1.22 on your Jenkins agent (go version should work).

Create a Pipeline job pointing to your repo (with Jenkinsfile at root).

Builds will reuse .gomodcache and .gocache in the workspace.

C) Docker

Enable BuildKit:

export DOCKER_BUILDKIT=1


Then build:

docker build -t myapp .


Note:- This setup ensures:

Fast incremental builds (Go cache + Jenkins cache + Docker layer cache).

Simple project structure (Go standard cmd/ + internal/).

Minimal binary (CGO_ENABLED=0, -s -w).