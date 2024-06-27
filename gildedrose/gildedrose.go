package gildedrose

type Item struct {
	Name    string
	SellIn  int
	Quality int
}

// update Quality
func UpdateQuality(items []*Item) {
	for _, v := range items {
		switch v.Name {
		case "Aged Brie":
			UpdateAgedBrie(v)
		case "Sulfuras, Hand of Ragnaros":
			UpdateSulfuras(v)
		case "Backstage passes to a TAFKAL80ETC concert":
			UpdateBackstage(v)
		case "Conjured Mana Cake":
			UpdateConjured(v)
		default:
			UpdateNormal(v)
		}
	}
}

// update Aged Brie
func UpdateAgedBrie(item *Item) {
	// The Quality of an item is never more than 50
	if item.Quality < 50 {
		// "Aged Brie" actually increases in Quality the older it gets
		item.Quality = item.Quality + 1
	}
	item.SellIn = item.SellIn - 1
	// Once the sell by date has passed, Quality degrades twice as fast
	if item.SellIn < 0 {
		if item.Quality < 50 {
			// "Aged Brie" actually increases in Quality the older it gets
			item.Quality = item.Quality + 1
		}
	}
}

// update Sulfuras
func UpdateSulfuras(item *Item) {
	// "Sulfuras", being a legendary item, never has to be sold or decreases in Quality
}

// update Backstage
// "Backstage passes", comme le "Aged Brie", augmente sa qualité (quality) plus le temps passe (sellIn) ; La qualité augmente de 2 quand il reste 10 jours ou moins et de 3 quand il reste 5 jours ou moins, mais la qualité tombe à 0 après le concert.
func UpdateBackstage(item *Item) {
	// The Quality of an item is never more than 50
	if item.Quality < 50 {
		// "Backstage passes", like aged brie, increases in Quality as its SellIn value approaches;
		item.Quality = item.Quality + 1

		// Quality increases by 2 when there are 10 days or less
		if item.SellIn < 11 {
			if item.Quality < 50 {
				item.Quality = item.Quality + 1
			}
		}
		// and by 3 when there are 5 days or less but
		if item.SellIn < 6 {
			if item.Quality < 50 {
				item.Quality = item.Quality + 1
			}
		}
	}

	// Quality drops to 0 after the concert
	item.SellIn = item.SellIn - 1
	if item.SellIn < 0 {
		item.Quality = 0
	}
}

// update Conjured
func UpdateConjured(item *Item) {
	// The Quality of an item is never negative
	if item.Quality > 0 {
		// "Conjured" items degrade in Quality twice as fast as normal items
		item.Quality = item.Quality - 2
	}
	item.SellIn = item.SellIn - 1

	if item.SellIn < 0 {
		if item.Quality > 0 {
			item.Quality = item.Quality - 2
		}
	}
}

// update Normal
func UpdateNormal(item *Item) {
	// The Quality of an item is never negative
	if item.Quality > 0 {
		item.Quality = item.Quality - 1
	}

	item.SellIn = item.SellIn - 1

	// Once the sell by date has passed, Quality degrades twice as fast
	if item.SellIn < 0 {
		if item.Quality > 0 {
			item.Quality = item.Quality - 1
		}
	}
}
