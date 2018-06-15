{% set model = "/tmp/model.tgz" %}

get_model_protobuf:
  archive.extracted:
    - source: {{ salt['pillar.get']('jibjib:lookup:files:model_proto') }}
    - name: {{ salt['pillar.get']('jibjib:lookup:serving:dir') }}
    - enforce_toplevel: False
    - makedirs: True
    - skip_verify: True

get_pickle_files:
  archive.extracted:
    - source: {{ salt['pillar.get']('jibjib:lookup:files:mappings') }}
    - name: {{ salt['pillar.get']('jibjib:lookup:query:dir') }}/input/pickle
    - enforce_toplevel: False
    - makdirs: True
    - skip_verify: True