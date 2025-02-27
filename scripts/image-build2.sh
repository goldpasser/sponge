#!/bin/bash

# two-stage build docker image

serverName="serverNameExample_mixExample"
# image name of the service, no capital letters
SERVER_NAME="project-name-example.server-name-example"
# Dockerfile file directory
DOCKERFILE_PATH="build"
DOCKERFILE="${DOCKERFILE_PATH}/Dockerfile_build"

# image repo address, REPO_HOST="ip or domain", passed in via the first parameter
REPO_HOST=$1
if [ "X${REPO_HOST}" = "X" ];then
        echo "param 'repo host' cannot be empty, example: ./image-build.sh hub.docker.com v1.0.0"
        exit 1
fi
# the version tag, which defaults to latest if empty, is passed in via the second parameter
TAG=$2
if [ "X${TAG}" = "X" ];then
        TAG="latest"
fi
# image name and tag
IMAGE_NAME_TAG="${REPO_HOST}/${SERVER_NAME}:${TAG}"

PROJECT_FILES=$(ls)
tar zcf ${serverName}.tar.gz ${PROJECT_FILES}
mv -f ${serverName}.tar.gz ${DOCKERFILE_PATH}
echo "docker build --force-rm -f ${DOCKERFILE} -t ${IMAGE_NAME_TAG} ${DOCKERFILE_PATH}"
docker build --force-rm -f ${DOCKERFILE} -t ${IMAGE_NAME_TAG} ${DOCKERFILE_PATH}
rm -rf ${DOCKERFILE_PATH}/${serverName}.tar.gz
