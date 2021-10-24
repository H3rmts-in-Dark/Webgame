# only one time
# docker build -t preparedImage -f PrepareDockerimage .
docker build -t webgame .

docker run --rm --name WebGame -p 18265 webgame