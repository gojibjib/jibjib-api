stop_serving:
  cmd.run:
    - name: docker-compose down serving
    - cwd: {{ salt['pillar.get']('jibjib:lookup:query:dir') }}
