# Proxy Server Documentation

## Routers

### Health Check
- `GET /health-check`
    - **Response:** `OK`

### User Registration
- `POST /register`
    - **Body:** Use `multipart/form-data` to send the following fields:
        - `login` (text): your username
        - `password` (text): your password

### User Login
- `GET /login`
    - **Description:** Returns a JWT token.
    - **Body:** Use `multipart/form-data` to send the following fields:
        - `login` (text): your username
        - `password` (text): your password
- **Response:**
```json lines
{
  "token": "your_jwt_token",
  "refreshToken": "your_refresh_token"
}
```

## Models structure

```sql
Request{
    method  string
    url     string
    headers json
} 

Response {
  id INTEGER
  status INTEGER
  headers JSON
  length INTEGER
}
```

## Run project

**Start project first time or after changes** ```make build```

**otherwise** ```make up```

**Stop project** ```make down```