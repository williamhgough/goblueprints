# Base this docker container off the official golang docker image.
# Docker containers inherit everything from their base.
FROM golang:1.10.3
# Create a directory inside the container to store all our application and then make it the working directory.
RUN mkdir -p /go/src/app
WORKDIR /go/src/app
# Copy the app directory (where the Dockerfile lives) into the container.
COPY . /go/src/app
# Download and install any required third party dependencies into the container.
RUN go get github.com/pilu/fresh
# install dep
RUN go get github.com/golang/dep/cmd/dep
# add Gopkg.toml and Gopkg.lock
ADD Gopkg.toml Gopkg.toml
ADD Gopkg.lock Gopkg.lock
# install packages
# --vendor-only is used to restrict dep from scanning source code
# and finding dependencies
RUN dep ensure --vendor-only

WORKDIR /go/src/app/cmd
# Now tell Docker what command to run when the container starts
CMD ["fresh", "main.go"]