# RestfulAPI

This is a sample webservice that I made with Golang. It has nested structs as well.
A Waitgroup is also initialized in this. As multiple goroutines are added to a go module, a waitgroup becomes useful as it can block specific code so that a set of goroutines may complete execution.

Running the following in the parent directory will create a module.
go mod init example/web-service-gin
Make sure gin package exists:
`go get -u github.com/gin-gonic/gin`

To run the goroutines, execute go run webservice_exhibit.go from the parent directory of the go file on a command line.
go run webservice_exhibit.go

To post:
The following example curl command can be run directly from the parent directory on Terminal

curl -v http://localhost:8080/customer-info-service --include --header "Content-Type: application/json" --request "POST" --data "{\"Name\": \"Henry Ford\", \"Phone\": 3139826001, \"Address\": {\"StreetInfo\": \"20900 Oakwood Blvd\", \"City\" : \"Dearborn\", \"State\": \"MI\", \"ZipCode\": 48124 } }"


Otherwise, saving the command as a .bat file on Windows (and launching it from Command Prompt) does the trick as well.

A successful GET and POST will be indicated by HTTP 200 OK status. Anything otherwise will result in a 500 Internal Server Error.
The curl command can be altered to include more data or different data, but the syntax must match or else a 500 error will result. 
