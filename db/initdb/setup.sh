#!/usr/bin/env bash
# For database initiliaziation, start your container first with
# docker run --rm --name mongo -d \
#   -v $(pwd)/data:/data/db \
#   -v $(pwd)/conf:/etc/mongo \
#   -v $(pwd)/import:/import \
#   -v $(pwd)/initdb:/initdb \
#   mongo:3.6.5 --config=/etc/mongo/mongod.conf
#  
# Make sure this script setup.sh is present in ./initdb
# 
# Then run
# docker exec mongo bash /initdb/setup.sh
# 
# Then shutdown the container and start docker compose

cat /initdb/init_db.js | mongo admin
mongoimport --db birds --collection birds --file /import/birds.json -u root -p example --authenticationDatabase admin --jsonArray
