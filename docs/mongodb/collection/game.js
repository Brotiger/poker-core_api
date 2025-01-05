game = {
    "_id": "ObjectId",
	"status": "enum:[waiting,started]",
	"name": "string",
	"password": "string",
	"ownerId": "ObjectId",
	"maxPlayers": "int",
	"countPlayers": "int",
	"updatedAt": "date",
	"createdAt": "date",
}