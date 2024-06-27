package gildedrose_test

import (
	"idonia/gildedrose"
	"testing"
)

func TestUpdateQuality(t *testing.T) {
	items := []*gildedrose.Item{
		{"Aged Brie", 2, 0},
		{"Elixir of the Mongoose", 5, 7},
		{"Sulfuras, Hand of Ragnaros", 0, 80},
		{"Sulfuras, Hand of Ragnaros", -1, 80},
		{"Backstage passes to a TAFKAL80ETC concert", 15, 20},
		{"Backstage passes to a TAFKAL80ETC concert", 10, 49},
		{"Backstage passes to a TAFKAL80ETC concert", 5, 49},
		{"Conjured Mana Cake", 3, 6},
	}

	expectedQualities := []int{
		1,  // Aged Brie
		6,  // Elixir of the Mongoose
		80, // Sulfuras, Hand of Ragnaros
		80, // Sulfuras, Hand of Ragnaros
		21, // Backstage passes to a TAFKAL80ETC concert
		50, // Backstage passes to a TAFKAL80ETC concert
		50, // Backstage passes to a TAFKAL80ETC concert
		4,  // Conjured Mana Cake
	}

	gildedrose.UpdateQuality(items)

	for i, item := range items {
		if item.Quality != expectedQualities[i] {
			t.Errorf("Expected quality of %s to be %d, but got %d", item.Name, expectedQualities[i], item.Quality)
		}
	}
}

func TestUpdateQualityAfterSellByDate(t *testing.T) {
	items := []*gildedrose.Item{
		{"Aged Brie", -1, 0},
		{"Elixir of the Mongoose", -1, 7},
		{"Backstage passes to a TAFKAL80ETC concert", -1, 20},
		{"Conjured Mana Cake", -1, 6},
	}

	expectedQualities := []int{
		2, // Aged Brie
		5, // Elixir of the Mongoose
		0, // Backstage passes to a TAFKAL80ETC concert
		2, // Conjured Mana Cake
	}

	gildedrose.UpdateQuality(items)

	for i, item := range items {
		if item.Quality != expectedQualities[i] {
			t.Errorf("Expected quality of %s to be %d, but got %d", item.Name, expectedQualities[i], item.Quality)
		}
	}
}
