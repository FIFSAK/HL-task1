# Proxy Server Documentation

## Routers

### Health Check
- **Endpoint:** `GET /health-check`
  - **Response:** `OK`

### Make Request
- **Endpoint:** `POST /`
  - **Body:**
    ```json
    {
      "method": "GET",
      "url": "http://google.com",
      "headers": {
        "Authentication": "Basic bG9naW46cGFzc3dvcmQ=",

      }
    }
    ```

### Get Response
- **Endpoint:** `GET /?id={response_id}`
  - **Response:**
    ```json
    {
      "id": "response_id",
      "status": "HTTP-статус ответа стороннего сервиса",
      "headers": { "массив заголовков из ответа стороннего сервиса" },
      "length": "длина содержимого ответа"
    }
    ```

## Models Structure

```sql
### Request
Request {
  method  string,
  url     string,
  headers json
}

### Response
Response {
  id      INTEGER,
  status  INTEGER,
  headers JSON,
  length  INTEGER
}
```

## Run project

- **Start project first time or after changes:**
```bash
make build
```
- **Otherwise** 
```bash
make up
```
- **Stop project**
```bash
make down
```
