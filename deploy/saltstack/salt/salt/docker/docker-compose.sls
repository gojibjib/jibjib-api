download_compose:
  file.managed:
    - name: /usr/local/bin/docker-compose
    - source: https://github.com/docker/compose/releases/download/{{ salt['pillar.get']('docker-compose:lookup:version') }}/docker-compose-Linux-x86_64
    - user: root
    - group: root
    - mode: 755
    - skip_verify: True

install_bash_completion:
  pkg.installed:
    - name: bash-completion

download_compose_completion:
  file.managed:
    - name: /usr/share/bash-completion/completions/docker-compose
    - source: https://raw.githubusercontent.com/docker/compose/{{ salt['pillar.get']('docker-compose:lookup:version') }}/contrib/completion/bash/docker-compose
    - user: root
    - group: root
    - mode: 755
    - skip_verify: True
    - require:
      - install_bash_completion