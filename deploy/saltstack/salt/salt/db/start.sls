get_latest_db_image:
  cmd.run:
    - name: docker pull {{ salt['pillar.get']('jibjib:lookup:db:image') }}

start_db:
  cmd.run:
    - name: docker-compose up -d
    - cwd: {{ salt['pillar.get']('jibjib:lookup:db:dir') }}