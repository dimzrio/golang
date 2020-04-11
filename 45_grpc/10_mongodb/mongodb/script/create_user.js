db = db.getSiblingDB("admin")
db.createUser(
    {
      user: USER,
      pwd: USER_PASSWD,
      roles: [
         { role: "readWrite", db: "admin" }
      ]
    }
);