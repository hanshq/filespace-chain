if [ -z "$1" ]; then
  echo "Need to input tag name..."
  exit 1
fi

sudo docker login

TAG=$1

echo "Build container? (Y/n) "
read -r answer
if [ ! "$answer" != "${answer#[Nn]}" ] ;then
  sudo docker build -t hanshq/filespace-chain:$TAG ./
fi

echo "Run container? (Y/n) "
read -r answer
if [ ! "$answer" != "${answer#[Nn]}" ] ;then
  sudo docker rm filespace-$TAG
  sudo docker run -p 26656:26656 -p 26657:26657 -p 20080:1317 -p 29090:9090 -p 29091:9091 --name filespace-$TAG hanshq/filespace-chain:$TAG
fi

echo "Push container to repo? (Y/n) "
read -r answer
if [ ! "$answer" != "${answer#[Nn]}" ] ;then
  sudo docker push hanshq/filespace-chain:$TAG
fi

