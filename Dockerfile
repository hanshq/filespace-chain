# Simple usage with a mounted data directory:
# > docker build -t simapp .
#
# Server:
# > docker run -it -p 26657:26657 -p 26656:26656 -v ~/.simapp:/root/.simapp simapp simd init test-chain
# TODO: need to set validator in genesis so start runs
# > docker run -it -p 26657:26657 -p 26656:26656 -v ~/.simapp:/root/.simapp simapp simd start
#
# Client: (Note the simapp binary always looks at ~/.simapp we can bind to different local storage)
# > docker run -it -p 26657:26657 -p 26656:26656 -v ~/.simappcli:/root/.simapp simapp simd keys add foo
# > docker run -it -p 26657:26657 -p 26656:26656 -v ~/.simappcli:/root/.simapp simapp simd keys list

# TODO: cp $GOPATH/bin/filespace-chaind ./build

# Final image
FROM frolvlad/alpine-glibc

# Install ca-certificates
RUN apk add bash jq --update ca-certificates

WORKDIR /persistent

# Copy over binaries from the build-env
COPY ./build/ /usr/bin/

COPY ./scripts/ /tmp/scripts/
COPY ./data/genesis/ /tmp/data/genesis/

RUN ls -la /tmp/data/genesis/
RUN ls -la /tmp/scripts/

ENV CHAIN_ID=filespacechain
#ENV KEY_NAME=node1
#ENV MONIKER_NAME=node1

RUN chmod +x /tmp/scripts/*.sh

EXPOSE 26656 26657 1317 9090 9091
# Run simd by default, omit entrypoint to ease using container with simcli
ENTRYPOINT [ "/tmp/scripts/docker_entry.sh" ]

# NOTE: to run this image, docker run -d -p 26657:26657 ./start_local_node.sh {{chain_id}} {{genesis_account}}

#do this :) :
# ./scripts/docker_push.sh 12
