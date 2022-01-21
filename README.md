# server
The server subsystem for the Prism fourth-year design project.

# Setup
## Downloading dependencies
1. Make sure that you have Go installed, this can be done via `$ brew install go`. Note that if you install Go via brew, you may also need to install `gcc` via brew as well.
2. Run `go mod download` to install all dependencies listed in the `go.mod` file.

## Setting up the test database

### OSX
1. Install Postgres [here](https://postgresapp.com/downloads.html). Then install the app and click "initialize" once the application opens.
2. Run `brew install postgresql` to install postgresql CLI.
3. (Optional) Install [Postico](https://eggerapps.at/postico/) or [Tableplus](https://tableplus.com/) as a DB GUI.
### WSL
1. Run `brew install postgresql` to install postgresql CLI.
2. Start postgres through brew's recommended method (more detailed instructions below).
3. (Optional) Install [Tableplus](https://tableplus.com/) as a DB GUI.

Brew will recommend that postgres be started through `pg_ctl`. Add the following aliases to your shell profile to easily interact with postgres:
```bash
export PGDATA='/home/linuxbrew/.linuxbrew/var/postgres'
alias pgstart='pg_ctl start'
alias pgstop='pg_ctl stop'
```

To start postgres, simply run `pgstart`.

#### Troubleshooting starting the DB

If `pgstart` fails with `could not bind IPv4 address "127.0.0.1": Permission denied`, ensure that no postgres instance is running on Windows ([read more here](https://stackoverflow.com/questions/62154886/postgres-password-authentication-failed)).

If you see the message `WARNING: could not flush dirty data: Function not implemented`, consult [this issue](https://stackoverflow.com/questions/45437824/postgresql-warning-could-not-flush-dirty-data-function-not-implemented).

### Create a database

Start the psql CLI through running `psql postgres`, then:

1. Create the user through running `CREATE ROLE prism WITH LOGIN CREATEDB PASSWORD 'prism';`
2. Verify the `prism` role has been created through running `\du`, you should see the `prism` role among others
3. Quit psql by running `\q` and log into the database using this new role through running `psql prism_development -U prism`
4. Verify you are logged in as the correct user through running `\conninfo`, you should see something along the lines of `You are connected to database "prism_development" as user "prism" via socket in "/tmp" at port "5432".`
5. Create the table through running `CREATE TABLE users(id SERIAL PRIMARY KEY, name VARCHAR(255) NOT NULL);`
6. Verify the table exists through running `\d users`, you should see that `id` and `name` fields exist.


## Verifying setup
1. Start the server through running ```bashmake run``` (this will also build the project).
2. Open a new terminal window and run `curl localhost:8080/ping`. If everything is running correctly, you should see an output of `{"message": "pong"}`
3. Run `curl -X POST -d "name=prism" localhost:8080/users` to create a new user, you should see `{"id": <an-id-here>}`
4. Run `curl localhost:8080/users/<your-id-here>` with the ID you got back from the previous command, you should see `{"name": "prism"}`