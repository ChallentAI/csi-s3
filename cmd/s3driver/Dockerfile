FROM ghcr.io/challentai/docker-build-s3fs:0.0.3 as s3fsbuild
FROM golang:1.16-alpine as gobuild

WORKDIR /build
ADD go.mod go.sum /build/

RUN go get -d -v ./...
RUN go mod download

ADD . /build

RUN --mount=type=cache,target=/root/.cache/go-build CGO_ENABLED=0 GOOS=linux go build -a -ldflags '-extldflags "-static"' -o ./s3driver ./cmd/s3driver

FROM debian:buster-slim
LABEL maintainers="Cyrill Troxler <cyrilltroxler@gmail.com>"
LABEL description="csi-s3 slim image"

# s3fs and some other dependencies
RUN apt-get update && \
  apt-get install -y \
  s3fs curl unzip && \
  rm -rf /var/lib/apt/lists/*

# install rclone
ARG RCLONE_VERSION=v1.54.1
RUN cd /tmp \
  && curl -O https://downloads.rclone.org/${RCLONE_VERSION}/rclone-${RCLONE_VERSION}-linux-amd64.zip \
  && unzip /tmp/rclone-${RCLONE_VERSION}-linux-amd64.zip \
  && mv /tmp/rclone-*-linux-amd64/rclone /usr/bin \
  && rm -r /tmp/rclone*

# Build newer s3fs-fuse
WORKDIR /s3fs
COPY --from=s3fsbuild /s3fs-debian-package.tar s3fs-fuse.tar
RUN tar -xvf s3fs-fuse.tar
RUN ls ./
RUN apt-get install -y fuse mime-support libxml2 libcurl4 libssl1.1
RUN dpkg -i s3fs_1.91+git-c17c815-3_amd64.deb
# RUN apt update && \
#   apt install -y \
#   git automake g++
# RUN cd /tmp \
#   && git clone https://github.com/s3fs-fuse/s3fs-fuse.git

# RUN apt update && \
#   apt install -y \
#   autoconf libcurl4-openssl-dev libssl-dev uuid-dev zlib1g-dev libfuse-dev

# RUN cd /tmp \
#   && cd s3fs-fuse \
#   && ./autogen.sh \
#   && ./configure --prefix=/usr --with-openssl\
#   && make \
#   && make install

COPY --from=gobuild /build/s3driver /s3driver

# KINDA hacky, should probably be built here
ADD libs3fsawscred.so /libs3fsawscred.so
ADD libaws-cpp-sdk-identity-management.so /usr/local/lib/libaws-cpp-sdk-identity-management.so
ADD libaws-cpp-sdk-core.so /usr/local/lib/libaws-cpp-sdk-core.so
ADD libaws-cpp-sdk-sts.so /usr/local/lib/libaws-cpp-sdk-sts.so
ADD libaws-cpp-sdk-cognito-identity.so /usr/local/lib/libaws-cpp-sdk-cognito-identity.so

RUN ldd /libs3fsawscred.so

ENTRYPOINT ["/s3driver"]
