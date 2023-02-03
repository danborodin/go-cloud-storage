#!/bin/bash

mongosh <<EOF
   use admin;
   db.createUser(
     {
       user: "admin",
       pwd: "test12345",
       roles: [
         { role: "userAdminAnyDatabase", db: "admin" },
         { role: "readWriteAnyDatabase", db: "admin" }
       ]
     }
   );
EOF