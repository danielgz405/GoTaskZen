package structures

type InsertBoardRequest struct {
	Name                string `bson:"name" json:"name"`
	Description         string `bson:"description" json:"description"`
	Image               string `bson:"image" json:"image"`
	Background          string `bson:"background" json:"background"`
	Primary             string `bson:"primary" json:"primary"`
	Secondary           string `bson:"secondary" json:"secondary"`
	DesertRef           string `bson:"desert_ref" json:"desert_ref"`
	DesertRefBackground string `bson:"desert_ref_background" json:"desert_ref_background"`
}

type UpdateBoardRequest struct {
	Name                string `bson:"name" json:"name"`
	Description         string `bson:"description" json:"description"`
	Saved               bool   `bson:"saved" json:"saved"`
	Image               string `bson:"image" json:"image"`
	Background          string `bson:"background" json:"background"`
	Primary             string `bson:"primary" json:"primary"`
	Secondary           string `bson:"secondary" json:"secondary"`
	DesertRef           string `bson:"desert_ref" json:"desert_ref"`
	DesertRefBackground string `bson:"desert_ref_background" json:"desert_ref_background"`
}
