{% set image = salt['pillar.get']('jibjib:lookup:db:image') %}
{% set container_name = salt['pillar.get']('jibjib:lookup:db:container_name') %}
{% set dir = salt['pillar.get']('jibjib:lookup:db:dir') %}

# Prepare folder / Clean old install
ensure_db_exists_clean:
  file.directory:
    - name: {{ dir }}
    - makedirs: True
    - clean: True

# Copy files from Salt Fileserver
copy_db_compose_file:
  file.managed:
    - source: salt://files/db/docker-compose.yml
    - name: {{ dir }}/docker-compose.yml
    - makedirs: True
    - template: jinja

copy_db_conf_file:
  file.managed:
    - source: salt://files/db/mongod.conf
    - name: {{ dir }}/conf/mongod.conf
    - makedirs: True

copy_db_init_js:
  file.managed:
    - source: salt://files/db/init_db.js
    - name: {{ dir }}/initdb/init_db.js
    - makedirs: True
    - template: jinja

copy_db_setup_sh:
  file.managed:
    - source: salt://files/db/setup.sh
    - name: {{ dir }}/initdb/setup.sh
    - makedirs: True
    - template: jinja

# Grab import file
download_birds_json:
  file.managed:
    - source: {{ salt['pillar.get']('jibjib:lookup:files:birds_json') }}
    - name: {{ dir }}/import/birds.json
    - makedirs: True
    - skip_verify: True

# Run init routine
start_init_db_container:
  cmd.run:
    - name: docker run --rm --name {{ container_name }} -d -v $(pwd)/data:/data/db -v $(pwd)/conf:/etc/mongo -v $(pwd)/import:/import -v $(pwd)/initdb:/initdb {{ image }} --config=/etc/mongo/mongod.conf
    - cwd: {{ dir }}

sleep_some_time:
  cmd.run:
    - name: sleep 4

run_init_script:
  cmd.run:
    - name: docker exec {{ container_name }} bash /initdb/setup.sh
    - cwd: {{ dir }}

sleep_some_more:
  cmd.run:
    - name: sleep 4

stop_init_db_container:
  cmd.run:
    - name: docker container rm -f {{ container_name }}