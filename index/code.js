db.code.createIndex(
    {
        "updatedAt": 1
    },
    {
        "expireAfterSeconds": 900
    }
)