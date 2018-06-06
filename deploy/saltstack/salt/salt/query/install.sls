get_query_compose:
  file.managed:
    - source: salt://files/query/docker-compose.yml
    - name: {{ salt['pillar.get']('jibjib:lookup:query:dir') }}/docker-compose.yml
    - makedirs: True
    - template: jinja