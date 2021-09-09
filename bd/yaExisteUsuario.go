package bd

import (
	"context"
	"time"

	"github.com/cardozoqsantiago/ApiGoBackend/models"
	"go.mongodb.org/mongo-driver/bson"
)

/*ChequeoYaExisteUsuario revisa si un usuario ya existe*/
func ChequeoYaExisteUsuario(email string) (models.Usuario, bool, string) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("twiteer")
	col := db.Collection("usuarios")

	condicion := bson.M{"email": email}
	var resultado models.Usuario

	err := col.FindOne(ctx, condicion).Decode(&resultado)
	ID := resultado.ID.String()
	if err != nil {
		return resultado, true, ID
	}
	return resultado, false, ID

}
