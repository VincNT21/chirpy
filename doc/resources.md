# Chirpy API Resources

## User resource

### Structure
- `id`: string - The unique identifier for the user
- `email`: string - The user's email adress
- `created_at`: string (ISO 8601 datetime) - When the user was created
- `updated_at`: string (ISO 8601 datetime) - Last time the user info was updated
- `is_chirpy_red`: boolean - If user suscribed a paid Red membership or not

### Example
```json
{
    "id": "d8b5ad72-1a8d-4990-bb83-44bd4daa32dc",
    "email": "vinc@example.com",
    "created_at": "2025-03-10 14:23:07.051327",
    "updated_at": "2025-03-10 14:23:07.051327",
    "is_chirpy_red": "false"
}
```

## Chirp resource

### Structure
- `id`: string - The unique identifier for the chirp
- `body`: string - The chirp's body
- `user_id`: string - ID of the user that creates the chirp
- `created_at`: string (ISO 8601 datetime) - When the chirp was created
- `updated_at`: string (ISO 8601 datetime) - Last time the chirp info was updated

### Example
```json
{
    "id": "d8b5ad72-1a8d-4990-bb83-44bd4daa32dc",
    "body": "I'm the one who knocks!",
    "user_id": "d8b5ad72-1a8d-4990-bb83-44bd4daa32dc",
    "created_at": "2025-03-10 14:23:07.100534",
    "updated_at": "2025-03-10 14:23:07.100534"
}
```

