db.connect_token.createIndex(
    {
        "updatedAt": 1
    },
    {
        "expireAfterSeconds": 10
    }
)

db.connect_token.createIndex(
    {
        "token": 1
    }
)