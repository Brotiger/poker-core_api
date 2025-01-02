game = {
    "_id": "ObjectId",
	"status": "enum:[waiting,started]",
	"name": "string",
	"password": "string",
	"ownerId": "ObjectId",
	"socketIds": [
        "ObjectId",
    ],
	"maxPlayers": "int",
	"updatedAt": "date",
	"createdAt": "date",
}