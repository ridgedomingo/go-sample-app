# GO Sample App
Basic golang app that fetches, insert data from oracle database.

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
Open postman or execute ```curl localhost:8081/customers``` this should return an array of customers.

## References
- Folder structure was inspired from [here](https://www.youtube.com/watch?v=dxPakeBsgl4)
- Connecting golang to [oracle database](https://blogs.oracle.com/developers/post/connecting-a-go-application-to-oracle-database)
