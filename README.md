# SET UP

## Set Up the DB

```sql
CREATE DATABASE journey_family_db;

CREATE USER journey_family_user WITH 
    SUPERUSER 
    CREATEDB 
    CREATEROLE 
    LOGIN 
    PASSWORD 'journey_family_pass';

ALTER ROLE journey_family_user 
    SET client_encoding TO 'utf8';

ALTER ROLE journey_family_user 
    SET default_transaction_isolation TO 'read committed';

ALTER ROLE journey_family_user 
    SET timezone TO 'UTC';

GRANT ALL PRIVILEGES ON DATABASE journey_family_db TO journey_family_user;

CREATE DATABASE journey_family_test;

GRANT ALL PRIVILEGES ON DATABASE journey_family_test TO journey_family_user;

\q
```

### Set test database as a template after running all migrations

```sql
ALTER DATABASE "journey_family_test" WITH IS_TEMPLATE true;
```

## Setting up the project

Clone the repository and if you have go set up on your machine run


```bash
go mod download

go mod tidy
```

```bash
docker run --name journey_family_db -e POSTGRES_DB=journey_family -e POSTGRES_USER=journey_family -e POSTGRES_PASSWORD=journey_family -p 5434:5432 -d postgis/postgis
docker run --name journey_family_test -e POSTGRES_DB=journey_family_test -e POSTGRES_USER=journey_family -e POSTGRES_PASSWORD=journey_family -p 5435:5432 -d postgis/postgis
```


## journey-family-membership-backend sample data

# Using psql
sudo -u postgres psql < seed-database.sql

# Or connect and run interactively
sudo -u postgres psql 
# Then type: \i seed-database.sql
