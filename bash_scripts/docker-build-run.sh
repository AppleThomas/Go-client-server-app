# if permission denied, use chmod u+x docker-build-run.sh

echo "Building Image:"
docker build --tag go-client-server:latest .

# Run container
echo -e "Running Container:"
docker run --name hello-world -d -p 8080:8080 go-client-server

# Make standard request
echo -e "Making Request:"
curl http://localhost:8080/