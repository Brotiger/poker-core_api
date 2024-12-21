db.refresh_token.createIndex(
    {
        "token": 1,
    },
    {
        "background": true
    }
)

db.refresh_token.createIndex(
    {
        "userId": 1,
    },
    {
        "background": true
    }
)


db.refresh_token.createIndex(
    {
        "updatedAt": 1
    },
    {
        "expireAfterSeconds": 10800
    }
)