package main

type Item struct {
	name            string
	sellIn, quality int
}

func increaseQ(item Item) Item {
	if item.quality < 50 {
		item.quality = item.quality + 1
	}
	return item
}

func decreaseQ(item Item) Item {
	if item.quality > 0 {
		item.quality = item.quality - 1
	}
	return item
}

func updateQuality(item Item) Item {
	if item.name == "Aged Brie" {
		item = increaseQ(item)
	} else if item.name == "Backstage passes to a TAFKAL80ETC concert" {
		if item.quality < 50 {
			item.quality = item.quality + 1
		}
		if item.sellIn < 11 {
			item = increaseQ(item)
		}
		if item.sellIn < 6 {
			item = increaseQ(item)
		}
	} else {
		item = decreaseQ(item)
	}

	item.sellIn = item.sellIn - 1

	if item.sellIn < 0 {
		if item.name != "Aged Brie" {
			if item.name == "Backstage passes to a TAFKAL80ETC concert" {
				item.quality = item.quality - item.quality
			} else {
				item = decreaseQ(item)
			}
		} else {
			item = increaseQ(item)
		}
	}
	return item
}

func UpdateQuality(items []*Item) {
	for i := 0; i < len(items); i++ {
		item := items[i]

		if item.name == "Sulfuras, Hand of Ragnaros" {
			continue
		}
		it := updateQuality(*item)
		items[i] = &Item{
			name:    it.name,
			quality: it.quality,
			sellIn:  it.sellIn,
		}
	}

}
