game = {
    "_id": "ObjectId",
	"status": "enum:[waiting,started]",
	"name": "string",
	"password": "string",
	"ownerId": "ObjectId",
	"players": "[]ObjectId", // -> player.js
	"maxPlayers": "int",
	"updatedAt": "date",
	"createdAt": "date",
}