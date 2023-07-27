#!/usr/bin/env bash
up() {

  echo "Starting Portainer..."
  docker run -d -p 8000:8000 -p 9443:9443 --name info-manager \
  --restart=always \
  -v /var/run/docker.sock:/var/run/docker.sock \
  -v portainer_data:/data portainer/portainer-ce:latest &

}

down() {
  
}

case $1 in
  up)
    up
    ;;
  down)
    down
    ;;
  *)
    echo "Usage: $0 {up|down}"
    ;;
esac
