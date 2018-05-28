// Crate root user
db.createUser({
    "user": "root", 
    "pwd": "example", 
    "roles": [ "root" ] 
})
db.auth("root", "example")
// Create  read-only user
use birds
db.createUser({
    "user": "test-r",
    "pwd": "test-r",
    "roles": ["read"]
})
