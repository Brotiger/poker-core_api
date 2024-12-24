db.user.createIndex(
    {
        "email": 1,
    },
    {
        "unique": true,
        "background": true
    }
)