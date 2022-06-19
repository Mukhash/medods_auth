To start in docker-compose:
    make start

To stop and/or delete containers:
    make stop

To start locally:
    Change DB URL in config/config.yml to
        URL: "mongodb://localhost:27017"

    docker-compose up mongodb-medods
    go run cmd/app/main.go


Example of request:
    for /api/auth  
    {
        "guid": "123456789"
    }

    for /api/refresh
    {
        "refresh_token": "eyJhbGciOiJIUzUxMiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NTgyNDg4NTIsImlhdCI6MTY1NTY1Njg1MiwiVVVJRCI6IjEyMzQ1In0.HwHKoHMtwNTOzbZ-ps8Jhw7-8srjevEw_Oh6TKXVm97v2i65DvM-yKhmb9owmMfedkke6oMrd4EkI_PaG-stfw"
    }
