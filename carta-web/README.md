## go cap table

# running

```console
$ cd carta-web
$ go build carta-web
$ ./carta-web
```

# endpoints

- /certificates (lists all certs)
    - GET
    ```console
    $ curl http://localhost:8080/certificates
    ```
    - POST
    ```console
    $ curl http://localhost:8080/certificates --include --header  "Content-Type: application/json" --request "POST" --data '{"owner": "Joao", "quantity": "200000000", "issueDate": "2019-05-01T00:01:02Z"}'
    ```
- /certificates/:id (certificate details)
    - GET
    ```console
    $ curl http://localhost:8080/certificates/c2c6fb7a-84b9-4277-b695-d5c16117bf1d
    ```
- /captable (returns cap table with fully diluted ownership, optional as_of date query)
    - GET
    ```console
    $ curl http://localhost:8080/captable
    $ curl http://localhost:8080/captable?as_of=2021-01-01T00:00:00Z
    ```
