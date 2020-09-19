package ex3

import (
	"log"
	"os"
	text "text/template"

	"../entities"
)

// Test - Executes test
func Test(tpl *text.Template) {

	restaurantMeals := entities.Restaurant{
		Breakfast: entities.Menu{
			Name: "Breakfast",
			Foods: []entities.Food{
				entities.Food{Name: "Pancaques", Description: "Some really good and daily made pancaques with honey or jam"},
				entities.Food{Name: "Starting default", Description: "Some nice eggs with bacon :)"},
				entities.Food{Name: "A healthy begin", Description: "A mix of fresh fruits"},
			},
			Drinks: []entities.Drink{
				entities.Drink{Name: "Youghourt", Description: "A home made youghourt", Hot: false},
				entities.Drink{Name: "Coffee", Description: "Just a nice cup of good coffee to warm your day", Hot: true},
				entities.Drink{Name: "Hot Chocolatte", Description: "The best hot chocollate you will taste, well maybe not the best but a good one :p", Hot: true},
				entities.Drink{Name: "Juices", Description: "Made from real fruit (options: Orange, Lemon, Grape, Blueberry)", Hot: false},
			},
		},
		Lunch: entities.Menu{
			Name: "Lunch",
			Foods: []entities.Food{
				entities.Food{Name: "Spaghetti", Description: "A home pasta with a portion of protein and sauce of your choice"},
				entities.Food{Name: "Salad", Description: "Lettuce, tomato, onions, carrots, corn, cheese"},
				entities.Food{Name: "Hamburger", Description: "The best choice of everyday"},
				entities.Food{Name: "Pizza", Description: "A very good, homemade pizza (Marguerita, 4 cheeses, Peperone)"},
				entities.Food{Name: "Ice cream", Description: "Who doens`t like to taste a good ice cream huh?"},
				entities.Food{Name: "Petit Gateau", Description: "A good option to finish your meal"},
			},
			Drinks: []entities.Drink{
				entities.Drink{Name: "Beer", Description: "Artesan beer or common beer", Hot: false},
				entities.Drink{Name: "Wine", Description: "Can choose a glass (x coins) or a bottle (y coins)", Hot: false},
				entities.Drink{Name: "Coffee", Description: "Just a nice cup of good coffee to accompany your meal", Hot: true},
				entities.Drink{Name: "Tea", Description: "Just a nice cup of good tea (Woo long, Camomile, Green tea)", Hot: true},
				entities.Drink{Name: "Juices", Description: "Made from real fruit (options: Orange, Lemon, Grape, Blueberry)", Hot: false},
			},
		},
		Dinner: entities.Menu{
			Name: "Dinner",
			Foods: []entities.Food{
				entities.Food{Name: "Spaghetti", Description: "A home pasta with a portion of protein and sauce of your choice"},
				entities.Food{Name: "Salad", Description: "Lettuce, tomato, onions, carrots, corn, cheese"},
				entities.Food{Name: "Hamburger", Description: "The best choice of everyday"},
				entities.Food{Name: "Pizza", Description: "A very good, homemade pizza (Marguerita, 4 cheeses, Peperone)"},
				entities.Food{Name: "Soup", Description: "End your day with a light meal"},
				entities.Food{Name: "Ice cream", Description: "Who doens`t like to taste a good ice cream huh?"},
				entities.Food{Name: "Petit Gateau", Description: "A good option to finish your meal"},
			},
			Drinks: []entities.Drink{
				entities.Drink{Name: "Beer", Description: "Artesan beer or common beer", Hot: false},
				entities.Drink{Name: "Wine", Description: "Can choose a glass (x coins) or a bottle (y coins)", Hot: false},
				entities.Drink{Name: "Coffee", Description: "Just a nice cup of good coffee to accompany your meal", Hot: true},
				entities.Drink{Name: "Tea", Description: "Just a nice cup of good tea (Woo long, Camomile, Green tea)", Hot: true},
				entities.Drink{Name: "Juices", Description: "Made from real fruit (options: Orange, Lemon, Grape, Blueberry)", Hot: false},
			},
		},
	}

	err := tpl.Execute(os.Stdout, restaurantMeals)
	if err != nil {
		log.Println("ERROR:", err.Error())
		return
	}

}
