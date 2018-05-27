// Create read-only user
db.createUser({
    "user": "prod-r",
    "pwd": "prod-r",
    "roles": ["read"]
})