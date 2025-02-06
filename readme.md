# JEC Live Code

1. Clone the project

`git clone https://github.com/26kito/jec-live-code.git`

2. Start Notification Service:

`cd notification-service/cmd && go run main.go`

3. Start gateway in other terminal:

`cd gateway && go run main.go`

Features:
1. Create notification:

Http Method POST

- Endpoint:
http://localhost:8080/notifications

- Request payload:
```
{
    "email": "kito@example.com",
    "message": "Hi there",
    "type": "email"
}
```

2. Get unsend notification:

Http Method GET

- Endpoint:
http://localhost:8080/unsend-notifications

3. Update notification status:

Http Method PUT

-Endpoint:
http://localhost:8080/notifications/{notification_id}