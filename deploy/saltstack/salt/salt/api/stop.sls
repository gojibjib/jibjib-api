stop_api_compose:
  cmd.run:
    - name: docker-compose down
    - cwd: {{ salt['pillar.get']('jibjib:lookup:api:dir') }}
