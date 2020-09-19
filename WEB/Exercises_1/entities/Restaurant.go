package entities

// Drink - Represents a drink
// Name        string 	=> drinks name
// Description string 	=> drinks description
// Hot         bool 	=> indicates if drink is hot or cold
type Drink struct {
	Name        string
	Description string
	Hot         bool
}

// Food - Represents a food
// Name        string	=> food name
// Description string	=> food description
type Food struct {
	Name, Description string
}

// Menu - Represents a menu
// Name   string	=> menu name
// Foods  []Food	=> list of food that exists on the menu
// Drinks []Drink	=> list of drink that exists on the menu
type Menu struct {
	Name   string
	Foods  []Food
	Drinks []Drink
}

// Restaurant - Represents all menus that the restaurant provides
// Options -> Breakfast, Lunch, Dinner
type Restaurant struct {
	Name      string
	Breakfast Menu
	Lunch     Menu
	Dinner    Menu
}

// Restaurants - List of Restaurants
type Restaurants struct {
	Restaurants []Restaurant
}
