game = {
    "_id": "ObjectId",
	"status": "enum:[created,started]",
	"name": "string",
	"password": "string",
	"ownerId": "ObjectId",
	"users": [
        "ObjectId",
    ],
	"maxPlayers": "int",
	"updatedAt": "date",
	"createdAt": "date",
}