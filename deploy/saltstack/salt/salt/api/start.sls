start_jibjib_api:
  cmd.run:
    - name: docker-compose up -d
    - cwd: {{ salt['pillar.get']('jibjib:lookup:api:dir') }}