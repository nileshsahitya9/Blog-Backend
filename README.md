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