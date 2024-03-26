# GO Sample App
Basic golang app with crud functionalities. 

## Installation
This project uses oracle local database to interact with the database. To install oracle check out this [link.](https://www.oracle.com/ph/database/technologies/oracle-database-software-downloads.html)

If you already have an oracle database connect to your local database and execute the following scripts to create a table.
```
CREATE TABLE CUSTOMERS (
    EMAIL VARCHAR2(50),
    FIRST_NAME VARCHAR2(50),
    LAST_NAME VARCHAR2(50),
    SEX VARCHAR2(10)
);
```

```
INSERT INTO CUSTOMERS VALUES('johndoe@gmail.com', 'John', 'Doe', 'Male')
INSERT INTO CUSTOMERS VALUES('janedoe@gmail.com', 'Jane', 'Doe', 'Female')
```

## Running the application
Go to cmd/go-sample-app/ ```cd cmd/go-sample-app/``` and execute ```go run .```

## Testing the application
Open postman and import the following curl or execute via cli 

GET
```curl localhost:8081/customers```

This should return an array of customers.

POST
```
curl --location 'localhost:8081/customer' \
--header 'Content-Type: application/json' \
--data-raw '{
    "sex": "Male",
    "firstName": "Steve",
    "lastName": "Smith",
    "email": "stevesmith@gmail.com"
}'
```

This would create a new customer.


PUT
```
curl --location --request PUT 'localhost:8081/customer/stevesmith@gmail.com' \
--header 'Content-Type: application/json' \
--data-raw '{
    "sex": "Male",
    "firstName": "Steve",
    "lastName": "Smith"
}'
```
This would update an existing customer by email.


DELETE
```
curl --location --request DELETE 'localhost:8081/customer/stevesmith@gmail.com' \
--data ''
```
This would delete an existing customer by email.

## References
- Folder structure was inspired from [here](https://www.youtube.com/watch?v=dxPakeBsgl4)
- Connecting golang to [oracle database](https://blogs.oracle.com/developers/post/connecting-a-go-application-to-oracle-database)
