db.user.createIndex(
    {
        "email": 1,
    },
    {
        "unique": true,
        "background": true
    }
)

db.user.createIndex(
    {
        "username": 1,
    },
    {
        "unique": true,
        "background": true
    }
)

db.user.createIndex(
    {
        "email": 1,
        "emailConfirmed": 1,
    },
    {
        "background": true
    }
)

db.user.createIndex(
    {
        "updatedAt": 1
    },
    {
        "partialFilterExpression": {
            "emailConfirmed": false
        },
        "expireAfterSeconds": 600
    }
)