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

func (ab AgedBrie) qualify() Genre {
	ab.item = quality(ab.item, +1)
	return ab
}
func (ab AgedBrie) updateSellIn() Genre {
	item := ab.item
	item.sellIn = item.sellIn - 1
	return AgedBrie{item: item}
}
func (ab AgedBrie) update() Item {
	item := ab.item
	if item.sellIn < 0 {
		item = quality(item, +1)
	}
	return item
}

type Backstage struct {
	item Item
}

func (bs Backstage) qualify() Genre {
	item := bs.item
	if item.sellIn > 10 {
		item = quality(item, +1)
	}
	if item.sellIn >= 6 && item.sellIn <= 10 {
		item = quality(item, +2)
	}
	if item.sellIn >= 0 && item.sellIn <= 5 {
		item = quality(item, +3)
	}
	return Backstage{item: item}
}
func (bs Backstage) updateSellIn() Genre {
	item := bs.item
	item.sellIn = item.sellIn - 1
	return Backstage{item: item}
}
func (bs Backstage) update() Item {
	item := bs.item
	if item.sellIn < 0 {
		item.quality = item.quality - item.quality
	}
	return item
}

type Normal struct {
	item Item
}

func (ab Normal) qualify() Genre {
	ab.item = quality(ab.item, -1)
	return ab
}
func (ab Normal) updateSellIn() Genre {
	item := ab.item
	item.sellIn = item.sellIn - 1
	return Normal{item: item}
}
func (ab Normal) update() Item {
	item := ab.item
	if item.sellIn < 0 {
		item = quality(item, -1)
	}
	return item
}

type Genre interface {
	qualify() Genre
	updateSellIn() Genre
	update() Item
}

func UpdateState(ab Genre) Item {
	it := ab.qualify().updateSellIn().update()
	return it
}

func GenreFactory(item Item) Genre {
	switch item.name {
	case "Aged Brie":
		return AgedBrie{item: item}
	case "Backstage passes to a TAFKAL80ETC concert":
		return Backstage{item: item}
	default:
		return Normal{item: item}
	}
}

func UpdateQuality(items []*Item) {
	for i := 0; i < len(items); i++ {
		item := items[i]
		if item.name == "Sulfuras, Hand of Ragnaros" {
			continue
		}

		it := UpdateState(GenreFactory(*item))
		items[i] = &Item{
			name:    it.name,
			quality: it.quality,
			sellIn:  it.sellIn,
		}
	}
}
