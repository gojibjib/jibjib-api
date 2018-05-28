stop_db_compose:
  cmd.run:
    - name: docker-compose down
    - cwd: {{ salt['pillar.get']('jibjib:lookup:db:dir') }}