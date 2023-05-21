package db

type Plant struct {
	ID               string   `json:"_id" bson:"_id"`
	Name             string   `json:"name" bson:"name"`
	ScientificName   string   `json:"scientific_name" bson:"scientific_name"`
	Family           string   `json:"family" bson:"family"`
	ScientificFamily string   `json:"scientific_family" bson:"scientific_family"`
	Habitats         []string `json:"habitats" bson:"habitats"`
	Regions          []any    `json:"regions" bson:"regions"`
	Soils            []string `json:"soils" bson:"soils"`
	IsGlobalRedbook  bool     `json:"is_global_redbook"`
	Sowing           struct {
		Start       string `json:"start"`
		End         string `json:"end"`
		Conditions  string `json:"conditions"`
		Description string `json:"description"`
	} `json:"sowing"`
	Harvesting struct {
		Start       string `json:"start"`
		End         string `json:"end"`
		Conditions  string `json:"conditions"`
		Description string `json:"description"`
	} `json:"harvesting"`
	ChemicalComposition []struct {
		Name  string `json:"name"`
		Value string `json:"value"`
	} `json:"chemical_composition"`
	MedicalPreparations string `json:"medical_preparations"`
	AnnualDemandTons    int    `json:"annual_demand_tons"`
}
