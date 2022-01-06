package currency

type Currency struct {
	Id           int    `json:"id" bson:"id"`
	Name         string `json:"picture" bson:"name"`
	Abbreviation string `json:"abbreviation" bson:"abbreviation"`
	Likes        string `json:"likes" bson:"likes"`
	Dislike      string `json:"dislikes" bson:"dislikes"`
}
