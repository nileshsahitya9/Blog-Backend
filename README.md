go run cmd/main.go



#Local setup

1. brew install postgresql@14
2. brew services start postgresql@14
3. psql postgres
4. CREATE ROLE blog_user WITH LOGIN;
   ALTER ROLE blog_user CREATEDB;
5. psql postgres -U blog_user
6. CREATE DATABASE blog_db;
7. GRANT ALL PRIVILEGES ON DATABASE blog_db TO blog_user;


#For Migrations 

1. brew install golang-migrate
2. migrate create -ext sql -dir db/migrations/ -seq <name_of_the_table> 
3. make migration_up to apply the migrations and move them up.
   make migration_down to undo the migrations and move them down.
   make migration_fix VERSION=<version> to force a specific version of the migration.


#Postgresql commands
1. psql postgres -U blog_user -W then enter the password
2. \dt+ - to check all the tables
3. \d+ <table_name> DESCRIBE TABLE - for the table description

#APIS - will prepare swagger

1.curl --location --request POST 'http://localhost:8080/api/v1/authentication/register' \
--header 'Content-Type: application/json' \
--data-raw '{
    "email" : "nileshsahitya9@gmail.com"
}'