# Example of a to-do app written in Go and PostgreSQL.

Using the [pq driver](https://github.com/lib/pq)

>  go get github.com/lib/pq

You should also have Postgresql installed

uses a database called notes 

> CREATE DATABASE notes 

and a table called posts

  CREATE TABLES posts(
  id SERIAL PRIMARY KEY,
  post TEXT NOT NULL);
