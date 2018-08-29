#!/bin/bash
#
# echo removing build folder
# rm -rf build
#
# echo preparing build directory
# mkdir build
# cat << 'EOF' > build/Dockerfile
# FROM alpine:3.6
#
# ADD hello /go/bin/hello
#
# EXPOSE 80
#
# RUN apk add --no-cache ca-certificates
#
# ENTRYPOINT /go/bin/hello
# EOF
#
# echo building go executable
# GOOS=linux GOARCH=amd64 go build -o build/hello
#
# echo building docker image
# cd build
# docker build --tag hyper-go .
#
# echo executing docker image
# docker run --env BEAUTIFUL=huhu hyper-go

echo "please provide your docker-hub cerdentials"
read -s -p "Username: " dockerhub_username
echo 
read -s -p "Password: " dockerhub_password
echo 

echo login to docker-hub
docker login --username ${dockerhub_username} --password ${dockerhub_password}

