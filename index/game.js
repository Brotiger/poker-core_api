db.game.createIndex(
    {
        "ownerId": 1,
    },
    {
        "background": true
    }
)

db.game.createIndex(
    {
        "createdAt": 1,
    },
    {
        "background": true
    }
)

db.game.createIndex(
    {
        "name": 1,
    },
    {
        "background": true
    }
)

db.game.createIndex(
    {
        "name": 1,
        "createdAt": 1,
    },
    {
        "background": true
    }
)

db.game.createIndex(
    {
        "name": 1,
    },
    {
        "background": true
    }
)