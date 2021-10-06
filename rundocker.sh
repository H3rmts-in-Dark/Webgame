docker build -t webgame .
docker run --rm --name WebGame -p 9090:18265 webgame