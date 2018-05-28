start_db:
  cmd.run:
    - name: docker-compose up -d
    - cwd: {{ salt['pillar.get']('jibjib:lookup:db:dir') }}