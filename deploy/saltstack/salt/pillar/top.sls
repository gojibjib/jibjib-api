base:
  '*':
    - main
  
  'jibjib_api':
    - db_api_auth
    - db_ip

  'jibjib_db':
    - db_api_auth
    - db_root_auth