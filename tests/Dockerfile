# syntax=docker/dockerfile:1.5
# this is here so we can grab the latest version of kind and have dependabot keep it up to date
FROM kindest/node:v1.30.0@sha256:047357ac0cfea04663786a612ba1eaba9702bef25227a794b52890dd8bcd692e

# this is here so we can grab the latest version of skopeo and have dependabot keep it up to date
FROM quay.io/skopeo/stable:v1.15.1

FROM python:3.12@sha256:4584ea46d313a10e849eb7c5ef36be14773418233516ceaa9e52a8ff7d5e35a5

RUN apt-get update \
	&& apt-get install -y curl git \
	&& rm -rf /var/lib/apt/lists/*

WORKDIR /workspace/tests

COPY --link tests/requirements.txt /workspace/tests/
RUN pip install --require-hashes -r requirements.txt --no-deps

COPY --link deployments /workspace/deployments
COPY --link config /workspace/config

RUN curl -LO https://storage.googleapis.com/kubernetes-release/release/$(curl -s https://storage.googleapis.com/kubernetes-release/release/stable.txt)/bin/linux/amd64/kubectl \
	&& install -o root -g root -m 0755 kubectl /usr/local/bin/kubectl \
	&& apt-get update && apt-get install -y apache2-utils

COPY --link tests /workspace/tests

COPY --link pyproject.toml /workspace/

ENTRYPOINT ["python3", "-m", "pytest"]
