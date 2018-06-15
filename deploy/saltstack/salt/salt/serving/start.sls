get_latest_serving_image:
  cmd.run:
    - name: docker pull {{ salt['pillar.get']('jibjib:lookup:serving:image') }}

start_jibjib_serving:
  cmd.run:
    - name: docker-compose up -d serving
    - cwd: {{ salt['pillar.get']('jibjib:lookup:serving:dir') }}