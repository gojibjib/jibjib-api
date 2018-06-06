stop_and_remove_all_containers:
  cmd.run:
    - name: docker container rm -f $(docker ps -a -q)

image_prune:
  cmd.run:
    - name: docker image prune -f