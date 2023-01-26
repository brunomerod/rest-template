# Rest Template

This project structure was created based on the [Go Kit](https://github.com/go-kit/kit) Framework. This specific implementation is just an example for services with database connection.

**Rationale**: The idea is that you have a wrapper [ENDPOINT] for all your services and easily manage dependency injection and middleware usage, also it's very easy to maintain as it has a simple testing structure.

## What do you need to run
* Go (1.19 you can change the go.mod for your version)
* MakeFile (run all commands, you can still use direct go commands)
* Docker (to run the local database)
* [MockGen](https://github.com/golang/mock)

Clone the project and run ```go mod tidy``` so you can have all your dependencies installed.

### MakeFile Commands
```make run```   
Runs the ```cmd/main```

```make mock```  
Generates the mocks for the application.  
In a productive scenario, the mocks should be on the gitignore and the mocks should be generated on every PR for a updated fresh use by the unit testing.

```make rmmock```  
Deletes the mocks.

```make test```  
Run the tests, alias for ```go test ./...```

```make cover```  
Generates the coverage report.

## Appendix

#### The structure explained
```
/.functional_tests
- get_accounts.sh       >> curl to the get accounts service, the first time should be empty
- post_account.sh       >> curl to insert an account
- post_transaction.sh   >> curl to insert a transaction, should throw an error if the account doesn't exist

/cmd
- main.go           >> main application, where we bake everything together

/config
- config.go         >> gets all the env variables needed to run
- .env              >> the environment variables, in a normal scenario this would be on the gitignore
- /connect
-- mongodb.go       >> configure the mongodb connection
- /setup
-- logger.go        >> configure the loggrus for the application
-- setup.go         >> setup initalization


/pkg
- /service          >> the core bussisness
-- /endpoint        >> the endpoint wrapper for the application
--- service.go      >> main logic, should contain all the bussisness rules
--- entity.go       >> the models for the whole application
--- repository.go   >> the connection/database operations
--- /middleware     >> middleware functions
- /transport        >> defines the transport methods
```