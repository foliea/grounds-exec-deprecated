#!/bin/sh

set -e

REPOSITORY="grounds"

get_images_dirs() {
  echo $(find dockerfiles -maxdepth 1 -type d | grep dockerfiles/)
}

get_image_name() {
  echo "$REPOSITORY/$(echo $1 | cut -f2 -d "/")"
}

# Build local images
build() {
  docker build -t $(get_image_name $1) "$1/"
}

# Push images to repository
push() {
  docker push $(get_image_name $1)
}

# Pull images from repository
pull() {
  docker pull $(get_image_name $1)
}

main() {
  # If first parameter from CLI is missing or empty
  if [ -z $1 ]; then
    echo "usage: [build|push|pull]"
    return
  fi
  # For every images
  for dir in $(get_images_dirs); do
    # Launch function corresponding to first parameter from CLI
    eval $1 $dir
  done
}

main $1