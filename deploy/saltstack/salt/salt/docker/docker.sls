install_docker_core_package_dependencies:
  pkg.installed:
    - pkgs: {{ salt['pillar.get']('system:lookup:pkg:required')}}

{% if salt['grains.get']('os') == 'Debian' %}
install_docker_distro_specific_package_dependencies:
  pkg.installed:
    - pkgs:
      - gnupg2
{% endif %}

add_docker_package_repo:
  pkgrepo.managed:
    - name: deb [arch={{ salt['grains.get']('osarch') }}] https://download.docker.com/linux/{{ salt['grains.get']('os')|lower }} {{ salt['grains.get']('oscodename') }} stable
    - key_url: {{ salt['pillar.get']('docker:lookup:pkg:key_url') }}

install_docker_package:
  pkg.installed:
    - name: {{ salt['pillar.get']('docker:lookup:pkg:name') }}
    - require:
      - add_docker_package_repo

add_docker_group:
  group.present:
    - name: docker

{% for usr in salt['pillar.get']('docker:lookup:users') %}
add_{{usr}}_to_docker_group:
  user.present:
    - name: {{ usr }}
    - groups:
      - docker
    - require:
      - add_docker_group
{% endfor %}