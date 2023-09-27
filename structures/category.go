package structures

type InsertCategoryRequest struct {
	Name        string `bson:"name" json:"name"`
	Description string `bson:"description" json:"description"`
	Primary     string `bson:"primary" json:"primary"`
	Secondary   string `bson:"secondary" json:"secondary"`
	TableId     string `bson:"table_id" json:"table_id"`
}

type UpdateCategoryRequest struct {
	Name        string `bson:"name" json:"name"`
	Description string `bson:"description" json:"description"`
	Primary     string `bson:"primary" json:"primary"`
	Secondary   string `bson:"secondary" json:"secondary"`
}

type InsertTaskRequest struct {
	Name        string `bson:"name" json:"name"`
	Description string `bson:"description" json:"description"`
	Priority    string `bson:"priority" json:"priority"`
}

type UpdateTaskRequest struct {
	Name        string `bson:"name" json:"name"`
	Description string `bson:"description" json:"description"`
	Priority    string `bson:"priority" json:"priority"`
	Active      bool   `bson:"active" json:"active"`
}
