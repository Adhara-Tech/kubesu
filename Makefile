SHELL=/bin/bash -o pipefail
# Prechecks of needed packages, some can be autoinstalled later on
HAS_GO := $(shell command -v go;)
HAS_GIT := $(shell command -v git;)
HAS_GOLANGCI_LINT := $(shell command -v golangci-lint;)
HAS_GO_JUNIT_REPORT := $(shell command -v go-junit-report;)
DOCKER_REPO ?= "local"
DOCKER_IMAGE_TAG ?= "SNAPSHOT"
ARTIFACT_NAME := "kubesu"
DOCKER_IMAGE_NAME := ${DOCKER_REPO}/${ARTIFACT_NAME}

UNAME_S := $(shell uname -s)

BINARIES_DIR ?= bin

GOPROXY ?= ""

ARTIFACTS_DIR ?= _artifacts

ifeq ($(CI),true)
CI_TAG := "-ci"
endif

ifeq ($(UNAME_S),Linux)
	UUID := $(shell cat /proc/sys/kernel/random/uuid)
else ifeq ($(UNAME_S),Darwin)
	UUID := $(shell uuidgen|tr '[:upper:]' '[:lower:]')
else
	$(error Platform $(UNAME_S) not supported!)
endif


check_env:
ifndef HAS_GO
	$(error You MUST install go)
endif
ifndef HAS_GIT
	$(error You MUST install git)
endif
ifndef HAS_GOLANGCI_LINT
	$(error You MUST install golangci-lint)
endif

check_docker_env:
ifndef DOCKER_REPO
	$(error DOCKER_REPO variable is mandatory)
endif
ifndef DOCKER_IMAGE_TAG
	$(error DOCKER_IMAGE_TAG variable is mandatory)
endif


.PHONY: validate
validate: check_env

.PHONY: test
test: validate
	@echo "Placeholder Makefile rule for tests"

.PHONY: build_in_docker
build_in_docker: check_docker_env
	@echo "Building in docker ${DOCKER_IMAGE_NAME}:${DOCKER_IMAGE_TAG}"
	docker build -t ${DOCKER_IMAGE_NAME}:${DOCKER_IMAGE_TAG} -t ${DOCKER_IMAGE_NAME}:${DOCKER_IMAGE_TAG} .

.PHONY: publish
publish: build_in_docker
	docker push ${DOCKER_IMAGE_NAME}:${DOCKER_IMAGE_TAG}
