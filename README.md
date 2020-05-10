# workshop-vrn

# RUN:
## 1. Via docker:
docker build -t workshop .
docker run --rm -p 8080:8080 workshop

## 2. Via DockerCompose:
docker-compose up --build
docker-compose down