db.users.createIndex(
    {
        "token": 1,
    },
    {
        "background": true
    }
)

db.key.createIndex(
    {
        "updatedAt": 1
    },
    {
        "expireAfterSeconds": 10800
    }
)