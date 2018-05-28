stop_all_containers:
  cmd.run:
    - name: docker container stop $(docker ps -a -q)

remove_all_containers:
  cmd.run:
    - name: docker container rm -f $(docker ps -a -q)