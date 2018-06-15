get_latest_query_image:
  cmd.run:
    - name: docker pull {{ salt['pillar.get']('jibjib:lookup:query:image') }}

get_latest_serving_image:
  cmd.run:
    - name: docker pull {{ salt['pillar.get']('jibjib:lookup:serving:image') }}

start_jibjib_query:
  cmd.run:
    - name: docker-compose up -d
    - cwd: {{ salt['pillar.get']('jibjib:lookup:query:dir') }}