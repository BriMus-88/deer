


POSTGRES
- window - R
- services.msc
- ENTER
- Find a service named: postgresql-x64-18
- Right-click → Start

To start in Powershell
- psql -U postgres

Too connect to the database
- psql -U postgres -d studentGolangTraining
or
-\c studentGolangTraining



database name: studentGolangTraining





CREATE TABLE students (id SERIAL PRIMARY KEY, name TEXT NOT NULL UNIQUE, course TEXT NOT NULL, age INTEGER NOT NULL, city TEXT NOT NULL);

CREATE TABLE students ( id SERIAL PRIMARY KEY, name TEXT NOT NULL UNIQUE, course TEXT NOT NULL, age INTEGER NOT NULL, city TEXT NOT NULL);







Postgres Database
Database server (listening on localhost:5432)

Default user: postgres

Admin GUI: pgAdmin

CLI tool: psql

Typical local connection details:

host: localhost
port: 5432
user: postgres
password: (the one you set during install)
database: postgres (default)