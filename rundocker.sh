# only one time
# docker build -t preparedimage -f PrepareDockerimage .
docker build -t webgame .

docker run --rm --name WebGame -p 18265 webgame