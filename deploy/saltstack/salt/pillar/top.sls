base:
  '*':
    - main
  
  'api':
    - db_api_auth
    - db_ip

  'db':
    - db_api_auth
    - db_root_auth