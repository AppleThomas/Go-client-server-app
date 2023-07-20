# if permission denied, use chmod u+x docker-build-run.sh

echo "Building Image:"
docker build --tag thomas3212/go-client-server .

# Run container
echo -e "Running Container:"
docker run --name go-client-server -d -p 3000:3000 thomas3212/go-client-server

# Make standard request
echo -e "Making Request:"
curl http://localhost:3000