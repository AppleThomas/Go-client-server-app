# if permission denied, use chmod u+x docker-build-run.sh

echo "Building Image:"
docker build --tag thomas3212/go-client-server .

# Run container
echo -e "Running Container:"
docker run --name go-client-server -d -p 3000:3000 thomas3212/go-client-server

# echo -e "Pushing to dockerhub:"
# docker push thomas3212/go-client-server

## Use this to remove running docker containers to avoid port issues when using k8s
# docker kill $(docker ps -q --filter ancestor=thomas3212/go-client-server)
# docker container prune