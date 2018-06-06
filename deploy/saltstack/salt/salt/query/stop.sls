stop_query_compose:
  cmd.run:
    - name: docker-compose down
    - cwd: {{ salt['pillar.get']('jibjib:lookup:query:dir') }}
