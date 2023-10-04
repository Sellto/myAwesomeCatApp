#bin/bash

docker build --platform $2 -t myawesomecatapp:$1 -f ./deployements/docker/Dockerfile .
docker tag myawesomecatapp:$1 selltom/myawesomecatapp:$1
docker push selltom/myawesomecatapp:$1