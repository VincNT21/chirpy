# Chirpy API Endpoints <!-- omit from toc -->

TABLE OF CONTENTS :
- [Public API Endpoints](#public-api-endpoints)
  - [Users related Endpoints](#users-related-endpoints)
    - [POST /api/users -- User creation](#post-apiusers----user-creation)
    - [PUT /api/users -- User info update](#put-apiusers----user-info-update)
    - [POST /api/login -- Authentification](#post-apilogin----authentification)
    - [POST /api/refresh -- Refresh authentification token](#post-apirefresh----refresh-authentification-token)
    - [POST /api/revoke -- Revoke authentification token](#post-apirevoke----revoke-authentification-token)
  - [Chirps related endpoints](#chirps-related-endpoints)
    - [POST /api/chirps -- Create a chirp](#post-apichirps----create-a-chirp)
    - [GET /api/chirps?author\_id={user\_id} -- List chirps created by a user](#get-apichirpsauthor_iduser_id----list-chirps-created-by-a-user)
    - [GET /api/chirps/{chirpID} -- Get a specific chirp](#get-apichirpschirpid----get-a-specific-chirp)
    - [DELETE /api/chirps/{chirpID} -- Delete a specific chirp](#delete-apichirpschirpid----delete-a-specific-chirp)
  - [Webhooks Endpoints](#webhooks-endpoints)
    - [POST /api/polka/webhooks -- Webhooks endpoint for Polka integration](#post-apipolkawebhooks----webhooks-endpoint-for-polka-integration)
- [Administrative Endpoints](#administrative-endpoints)
    - [GET /api/healthz -- Health check endpoint](#get-apihealthz----health-check-endpoint)
    - [GET /admin/metrics -- View API usage metrics](#get-adminmetrics----view-api-usage-metrics)
    - [GET /admin/reset -- Reset metric counter and users database](#get-adminreset----reset-metric-counter-and-users-database)

## Public API Endpoints

### Users related Endpoints

#### POST /api/users -- User creation
-> *Description* : Create a new user in server database

-> *Request need* : an email (string) and a password (string)

-> *Request body example* :
```json
{
    "email": "vinc@example.com",
    "password": "12345abcde"
}

```

-> *Response* : Status code 201-Created  
Response body example
```json
{
    "id": "d8b5ad72-1a8d-4990-bb83-44bd4daa32dc",
    "email": "vinc@example.com",
    "created_at": "2025-03-10 14:23:07.051327",
    "updated_at": "2025-03-10 14:23:07.051327",
    "is_chirpy_red": "false"
}
```

#### PUT /api/users -- User info update
-> *Description* : Update email or password for a specified user (defined by access token)

-> *Request header* : a valid access token  
Example:
```json
{
  "Authorization": "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJjaGlycHktYWNjZXNzIiwic3ViIjoiZjJjMmVlMWQtYWExZS00YzBiLTliNmEtODcyMmY5OWE1ZWQwIiwiZXhwIjoxNzQxNjIxODYyLCJpYXQiOjE3NDE2MTgyNjJ9._9-QuSMwwy8zEAgWyq7gcayyRUzN-DDXolWz8VmXIMc"
}
```
/!\ **If error-response 401-Unauthorized, client should call POST /api/refresh to get a new access token** /!\

-> *Request body* : a new email (string) and/or a new password (string)  
Example:
```json
{
    "email": "vinc@example.com",
    "password": "12345abcde"
}

```

-> *Response* : Status code 200-Ok 
Response body example
```json
{
    "id": "d8b5ad72-1a8d-4990-bb83-44bd4daa32dc",
    "email": "vinc@example.com",
    "created_at": "2025-03-10 14:23:07.051327",
    "updated_at": "2025-03-10 14:23:07.051327",
    "is_chirpy_red": "false"
}
```

#### POST /api/login -- Authentification
-> *Description* : Login user by verying given email/password, create a Refresh Token (valid for 60 days) stored in database and a Access Token (valid for 1 hour) not stored

-> *Request body* : a email (string) and a password (string)  
Example :
```json
{
    "email": "vinc@example.com",
    "password": "12345abcde"
}

```

-> *Response* : Status code 200-Ok  
Response body example
```json
{
    "id": "d8b5ad72-1a8d-4990-bb83-44bd4daa32dc",
    "email": "vinc@example.com",
    "created_at": "2025-03-10 14:23:07.051327",
    "updated_at": "2025-03-10 14:23:07.051327",
    "is_chirpy_red": "false",
    "refresh_token": "6fd580fcb0b1b2308dfa997d9952514c3ba60d1543c8475605fb2dccdfd23776",
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJjaGlycHktYWNjZXNzIiwic3ViIjoiNDYyODcwZDItYzdkYi00YmNkLWEyYTgtMGMzYTk4NzFkY2Y0IiwiZXhwIjoxNzQxNjIxODYyLCJpYXQiOjE3NDE2MTgyNjJ9.4a__wHtAd9mLGlDOFCu1wzL4NOd--R1TxTuioqnHPug"
}
```

#### POST /api/refresh -- Refresh authentification token
-> *Description* : Give a new access token (valid for 1 hour) if refresh token is still valid

-> *Request header* : a valid refresh token  
Example:
```json
{
  "Authorization": "Bearer 6fd580fcb0b1b2308dfa997d9952514c3ba60d1543c8475605fb2dccdfd23776"
}
```
/!\ **If error-response 401-Unauthorized, client should call POST /api/login to get new refresh and access tokens** /!\

-> *Request body* : none

-> *Response* : Status code 200-Ok  
Response body example
```json
{
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJjaGlycHktYWNjZXNzIiwic3ViIjoiNDYyODcwZDItYzdkYi00YmNkLWEyYTgtMGMzYTk4NzFkY2Y0IiwiZXhwIjoxNzQxNjIxODYyLCJpYXQiOjE3NDE2MTgyNjJ9.4a__wHtAd9mLGlDOFCu1wzL4NOd--R1TxTuioqnHPug"
}
```

#### POST /api/revoke -- Revoke authentification token
-> *Description* : Revoke a refresh token in server database

-> *Request header* : a valid refresh token  
Example:
```json
{
  "Authorization": "Bearer 6fd580fcb0b1b2308dfa997d9952514c3ba60d1543c8475605fb2dccdfd23776"
}
```

-> *Request body* : none

-> *Response* : http Status 203-No content

### Chirps related endpoints

#### POST /api/chirps -- Create a chirp
-> *Description* : Create a new chirp associated to logged in user and store it in server database

-> *Request header* : a valid access token  
Example:
```json
{
  "Authorization": "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJjaGlycHktYWNjZXNzIiwic3ViIjoiZjJjMmVlMWQtYWExZS00YzBiLTliNmEtODcyMmY5OWE1ZWQwIiwiZXhwIjoxNzQxNjIxODYyLCJpYXQiOjE3NDE2MTgyNjJ9._9-QuSMwwy8zEAgWyq7gcayyRUzN-DDXolWz8VmXIMc"
}
```
/!\ **If error-response 401-Unauthorized, client should call POST /api/refresh to get a new access token** /!\

-> *Request body* : new chirp's content (**less than 140 characters**)
Example:
```json
{
    "body": "Hello, world!"
}

```

-> *Response* : Status code 201-Created  
Response body example:
```json
{
    "id": "58d9cf84-44ee-4a41-baee-e37ad57e9886",
    "created_at": "2025-03-10 14:23:07.051327",
    "updated_at": "2025-03-10 14:23:07.051327",
    "body": "Hello, world!",
    "user_id": "d8b5ad72-1a8d-4990-bb83-44bd4daa32dc"
}
```

#### GET /api/chirps?author_id={user_id} -- List chirps created by a user
-> *Description* : Display a list of all chirps created by a user sorted by creation time, by default in ascending order. Can be sorted in descending order if "desc" is provided in query parameters.

-> *Request query parameter* : a user id and optionnaly a sort parameter ("asc" or "desc")
Example:
```
  GET /api/chirps?author_id=d8b5ad72-1a8d-4990-bb83-44bd4daa32dc&sort=desc
```

-> *Request body* : none

-> *Response* : Status code 200-Ok  
Response body example:
```json
[
    {
        "id": "441976a0-fe50-4ca6-89b6-9ffcdfdb4f87",
        "created_at": "2025-03-10T16:12:17.131285Z",
        "updated_at": "2025-03-10T16:12:17.131285Z",
        "body": "Mr President....",
        "user_id": "0402d449-79ab-4b57-8c1e-67bfc3912489"
    },
    {
        "id": "9a595310-4207-4653-ab51-f8972de4a262",
        "created_at": "2025-03-10T16:12:17.038333Z",
        "updated_at": "2025-03-10T16:12:17.038333Z",
        "body": "Gale!",
        "user_id": "f3bcc1d4-3274-4861-b9a1-50e3614210a2"
    }
]
```

#### GET /api/chirps/{chirpID} -- Get a specific chirp
-> *Description* : Display a specific chirp by given ID

-> *Request path needs to have* : a chirp id
Example:
```
  GET /api/chirps/441976a0-fe50-4ca6-89b6-9ffcdfdb4f87
```

-> *Request body* : none

-> *Response* : Status code 200-Ok  
Response body example:
```json
{
    "id": "441976a0-fe50-4ca6-89b6-9ffcdfdb4f87",
    "created_at": "2025-03-10T16:12:17.131285Z",
    "updated_at": "2025-03-10T16:12:17.131285Z",
    "body": "Mr President....",
    "user_id": "0402d449-79ab-4b57-8c1e-67bfc3912489"
}
```

#### DELETE /api/chirps/{chirpID} -- Delete a specific chirp
-> *Description* : Delete a specific chirp from database by given ID

-> *Request path needs to have* : a chirp id
Example:
```
  DELETE /api/chirps/441976a0-fe50-4ca6-89b6-9ffcdfdb4f87
```

-> *Request header* : a valid access token  
Example:
```json
{
  "Authorization": "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJjaGlycHktYWNjZXNzIiwic3ViIjoiZjJjMmVlMWQtYWExZS00YzBiLTliNmEtODcyMmY5OWE1ZWQwIiwiZXhwIjoxNzQxNjIxODYyLCJpYXQiOjE3NDE2MTgyNjJ9._9-QuSMwwy8zEAgWyq7gcayyRUzN-DDXolWz8VmXIMc"
}
```
/!\ **If error-response 401-Unauthorized, client should call POST /api/refresh to get a new access token** /!\

-> *Request body* : none

-> *Response* : Status code 203-No content 
/!\ **If error-response 403-Forbidden, it means that access token doesn't match chirp's user_id** /!\

### Webhooks Endpoints

#### POST /api/polka/webhooks -- Webhooks endpoint for Polka integration
-> *Description* : Confirm a user's paid subscription by Polka 3rd party services

-> *Request header* : the API key that matches the one in server config
Example:
```json
{
  "Authorization": "ApiKey f271c81ff7084ee5b99a5091b42d486e"
}
```
/!\ **If error-response 401-Unauthorized, it means that API key is missing or doesn't match the one in server config** /!\

-> *Request body* : a user's ID and an event field set to "user.upgraded"
Example:
```json
{
    "data": {
        "user_id": "0402d449-79ab-4b57-8c1e-67bfc3912489"
    },
    "event": "user.upgraded",
}
```
/!\ **If "event" is not set to "user.upgraded", it will result also in a status code 204-No content response** /!\

-> *Response* : Status code 204-No content

## Administrative Endpoints

#### GET /api/healthz -- Health check endpoint
-> *Description* : Returns 200 when service is healthy

-> *Request header* : none

-> *Request body* : none

-> *Response* : Status code 200-Ok

#### GET /admin/metrics -- View API usage metrics
-> *Description* : Display an html page with hits count of server API **(admin-only)**

-> *Request header* : none

-> *Request body* : none

-> *Response* : Status code 200-Ok

#### GET /admin/reset -- Reset metric counter and users database
-> *Description* : Reset metric counter to 0 and delete all users from database **(admin-only)**

-> *Request header* : none

-> *Request body* : none

-> *Response* : Status code 200-Ok



