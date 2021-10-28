package main

type Item struct {
	name            string
	sellIn, quality int
}

func quality(item Item, q int) Item {
	item.quality = item.quality + q
	if item.quality > 50 {
		item.quality = 50
		return item
	}
	if item.quality < 0 {
		item.quality = 0
		return item
	}
	return item
}

func updateQuality(item Item) Item {
	// Update quality
	switch item.name {
	case "Aged Brie":
		item = quality(item, +1)
	case "Backstage passes to a TAFKAL80ETC concert":
		if item.quality < 50 {
			item.quality = item.quality + 1
		}
		if item.sellIn < 11 {
			item = quality(item, +1)
		}
		if item.sellIn < 6 {
			item = quality(item, +1)
		}
	default:
		item = quality(item, -1)
	}

	// Update sellIn
	item.sellIn = item.sellIn - 1

	// Over deal
	if item.sellIn < 0 {
		switch item.name {
		case "Aged Brie":
			item = quality(item, +1)
		case "Backstage passes to a TAFKAL80ETC concert":
			item.quality = item.quality - item.quality
		default:
			item = quality(item, -1)
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
