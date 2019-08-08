package leboncoin

// Categories is the list of categories.
var Categories = map[int]string{
	1:  "vehicules",
	2:  "voitures",
	3:  "motos",
	4:  "caravaning",
	5:  "utilitaires",
	6:  "equipement_auto",
	7:  "nautisme",
	8:  "immobilier",
	9:  "ventes_immobilieres",
	10: "locations",
	11: "colocations",
	12: "locations_gites",
	13: "bureaux_commerces",
}

// categoryExists returns `true` if the category exists and `false` if not.
func categoryExists(c int) bool {
	_, ok := Categories[c]
	return ok
}
