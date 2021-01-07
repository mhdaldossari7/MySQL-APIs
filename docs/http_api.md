# HTTP API

## Store User ID

POST : `/users/{user_id}`

*`user_id` : The user id to be stored in database.*

This function stores the user id in the database. It doesn't require any authorization in the header.

**Success Response:**
```json
{
    "success": true,
    "data": {
        "user_id": 1
    },
    "message": "Added"
}
```

**Error Response:**
- Internal Server Error

```json
{
    "success": false,
    "err": "ERR_INTERNAL_ERROR",
    "message": "unable to execute query due : Error 1045: Access denied for user ''@'localhost'"
}
```
If the client receives this error, it means the provided credentials for MySQL are wrong.

- Bad Request

```json
{
    "success": false,
    "err": "ERR_BAD_REQUEST",
    "message": "You must pass integers"
}
```
If the client receives this error, It means the user id is not integer. The user id must be integer.

[Back to Top](#http-api)

---

## Delete User ID

Delete : `/users/{user_id}`

*`user_id` : The user id to be deleted from database.*

This function deleted the user id from the database. It doesn't require any authorization in the header.

**Success Response:**
```json
{
    "success": true,
    "data": {
        "user_id": 1
    },
    "message": "Deleted"
}
```

**Error Response:**
- Internal Server Error

```json
{
    "success": false,
    "err": "ERR_INTERNAL_ERROR",
    "message": "unable to execute query due : Error 1045: Access denied for user ''@'localhost'"
}
```
If the client receives this error, it means the provided credentials for MySQL are wrong.

- Bad Request

```json
{
    "success": false,
    "err": "ERR_BAD_REQUEST",
    "message": "You must pass integers"
}
```
If the client receives this error, It means the user id is not integer. The user id must be integer.

[Back to Top](#http-api)

---

## Check If User ID Exists In The Database

GET : `/users/{user_id}`

*`user_id` : The user id to be checked in database.*

This function checks if the user id is already stored in database or not. It doesn't require any authorization in the header.

**Success Response:**

- If the user id exists in the database.

    ```json
    {
        "success": true,
        "message": "Exists in DB"
    }
    ```

- If the user id doesn't exist in the database.

    ```json
    {
        "success": true,
        "message": "Doesn't Exists in DB"
    }
    ```

**Error Response:**
- Internal Server Error

```json
{
    "success": false,
    "err": "ERR_INTERNAL_ERROR",
    "message": "unable to execute query due : Error 1045: Access denied for user ''@'localhost'"
}
```
If the client receives this error, it means the provided credentials for MySQL are wrong.

- Bad Request

```json
{
    "success": false,
    "err": "ERR_BAD_REQUEST",
    "message": "You must pass integers"
}
```
If the client receives this error, It means the user id is not integer. The user id must be integer.

[Back to Top](#http-api)

---


## Get All user ids

GET : `/users?limit={limit}`

*`limit` : Number of data to be fetched from database.*

This function fetches user ids from database based on limit value, limit value must be integer. It doesn't require any authorization in the header.


**Success Response:**
```json
{
    "success": true,
    "data": [
        {
            "user_id": 1
        },
        {
            "user_id": 2
        },
        {
            "user_id": 3
        }
    ]
}
```

**Error Response:**
- Internal Server Error

```json
{
    "success": false,
    "err": "ERR_INTERNAL_ERROR",
    "message": "unable to execute query due : Error 1045: Access denied for user ''@'localhost'"
}
```
If the client receives this error, it means the provided credentials for MySQL are wrong.

- Bad Request

```json
{
    "success": false,
    "err": "ERR_BAD_REQUEST",
    "message": "You must pass integers"
}
```
If the client receives this error, It means the limit value is not integer. The user id must be integer.

[Back to Top](#http-api)