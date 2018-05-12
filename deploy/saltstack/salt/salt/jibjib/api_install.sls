get_jibjib_api_compose:
  file.managed:
    - source: https://raw.githubusercontent.com/gojibjib/jibjib-api/master/docker-compose.yml
    - name: /home/ubuntu/docker-compose.yml
    - skip_verify: True