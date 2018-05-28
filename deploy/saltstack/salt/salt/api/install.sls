get_api_compose:
  file.managed:
    - source: salt://files/api/docker-compose.yml
    - name: {{ salt['pillar.get']('jibjib:lookup:api:dir') }}/docker-compose.yml
    - makedirs: True
    - template: jinja