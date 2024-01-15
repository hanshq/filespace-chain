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

WORKDIR /app

# Copy over binaries from the build-env
COPY ./build/ /usr/bin/

COPY ./scripts/ ./scripts/

RUN ls -la ./scripts/

ENV CHAIN_ID=local-testnet
ENV KEY_NAME=local-user
ENV MONIKER_NAME=local

RUN chmod +x ./scripts/*.sh

EXPOSE 26656 26657 1317 9090
# Run simd by default, omit entrypoint to ease using container with simcli
ENTRYPOINT [ "./scripts/start_docker_node.sh" ]

# NOTE: to run this image, docker run -d -p 26657:26657 ./start_single_node.sh {{chain_id}} {{genesis_account}}

#do this :) :

#sudo docker run -p 26657:26657 --entrypoint ./start_single_node.sh b7fbf87ed5ab "filespace-01" "space1g2mljfe788k2q0fsyqkyvcwr05rmgu82cq79t2"
#sudo docker run -p 26657:26657 --entrypoint ./start_single_node.sh hanshq/filespace-v01:latest "filespace-01" "space1g2mljfe788k2q0fsyqkyvcwr05rmgu82cq79t2"