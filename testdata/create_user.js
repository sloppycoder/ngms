db.createUser(
  {
    user: "admin",
    pwd: passwordPrompt(), // or cleartext password
    roles: [ { role: "userAdminAnyDatabase", db: "admin" }, "readWriteAnyDatabase" ]
  }
)

use dev
db.createUser(
  {
    user: "dev",
    pwd:  "dev",
    roles: [ { role: "readWrite", db: "dev" } ]
  }
)

