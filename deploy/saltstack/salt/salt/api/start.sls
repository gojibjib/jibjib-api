get_latest_api_image:
  cmd.run:
    - name: docker pull {{ salt['pillar.get']('jibjib:lookup:api:image') }}

start_jibjib_api:
  cmd.run:
    - name: docker-compose up -d
    - cwd: {{ salt['pillar.get']('jibjib:lookup:api:dir') }}