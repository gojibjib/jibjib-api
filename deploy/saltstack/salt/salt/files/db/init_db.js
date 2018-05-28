// Crate root user
db.createUser({
    "user": "{{ salt['pillar.get']('root_user:user')}}", 
    "pwd": "{{ salt['pillar.get']('root_user:pw') }}", 
    "roles": [ "root" ] 
})
db.auth("{{ salt['pillar.get']('root_user:user')}}", "{{ salt['pillar.get']('root_user:pw') }}")
// Create  read-only user
use {{ salt['pillar.get']('jibjib:lookup:db:db_name') }}
db.createUser({
    "user": "{{ salt['pillar.get']('db_user:user') }}",
    "pwd": "{{ salt['pillar.get']('db_user:pw') }}",
    "roles": ["read"]
})
