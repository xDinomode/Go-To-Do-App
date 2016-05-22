# Example of a to-do app written in Go and PostgreSQL.

1. go get github.com/lib/pq

2. sudo apt-get install postgresql postgresql-contrib

3. CREATE DATABASE notes 

4. CREATE TABLES posts(
  id SERIAL PRIMARY KEY,
  post TEXT NOT NULL);
