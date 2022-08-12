# GOREST

A simple RESTful API written in go using [gin framework](https://gin-gonic.com/).

## Endpoint architecture:

**/animals**

    * GET - Get a list of all animals, returned in JSON

**/animals/:id**

    * GET - Get an animal by its ID, returning the animal data as JSON

**/addanimal**


    * POST - Add a new animal from request data sent as JSON


## Run and test:

Install dependencies:
```shell
go get .
```

Run:
```shell
go run .
```

GET method:
```shell
curl http://localhost:8080/animals
```

GET method by id:
```shell
curl http://localhost:8080/animals/1
```

POST method:
```shell
curl http://localhost:8080/addanimal --include --header "Content-Type: application/json" \
	--request "POST" \
	--data '{"id": "4", "Name": "Col", "Type": "Cat", "Sex": "Male"}'
```