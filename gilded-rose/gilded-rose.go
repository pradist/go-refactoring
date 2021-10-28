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

type AgedBrie struct {
	item Item
}

func (ab AgedBrie) updateQuality() Item {
	return quality(ab.item, +1)
}
func (ab AgedBrie) updateSellIn(item Item) Item {
	item.sellIn = item.sellIn - 1
	return item
}
func (ab AgedBrie) overDeal(item Item) Item {
	if item.sellIn < 0 {
		item = quality(item, +1)
	}
	return item
}

func upQ(ab AgedBrie) Item {
	it := ab.updateQuality()
	it = ab.updateSellIn(it)
	it = ab.overDeal(it)
	return it
}

//ab.updateQuelity()

func updateQuality(item Item) Item {
	// Update quality
	switch item.name {
	case "Backstage passes to a TAFKAL80ETC concert":
		if item.sellIn > 10 {
			item = quality(item, +1)
		}
		if item.sellIn >= 6 && item.sellIn <= 10 {
			item = quality(item, +2)
		}
		if item.sellIn >= 0 && item.sellIn <= 5 {
			item = quality(item, +3)
		}
	default:
		item = quality(item, -1)
	}

	// Update sellIn
	item.sellIn = item.sellIn - 1

	// Over deal
	if item.sellIn < 0 {
		switch item.name {
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

		switch item.name {
		case "Aged Brie":
			it := upQ(AgedBrie{item: *item})
			items[i] = &Item{
				name:    it.name,
				quality: it.quality,
				sellIn:  it.sellIn,
			}
		default:
			it := updateQuality(*item)
			items[i] = &Item{
				name:    it.name,
				quality: it.quality,
				sellIn:  it.sellIn,
			}
		}
	}
}
