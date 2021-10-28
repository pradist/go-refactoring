package main

import (
	"reflect"
	"testing"

	"github.com/kr/pretty"
)

func TestSulfuras(t *testing.T) {
	cases := []struct {
		name    string
		inputs  []*Item
		outputs []*Item
	}{
		{name: "#4", inputs: []*Item{{"Sulfuras, Hand of Ragnaros", 0, 80}}, outputs: []*Item{{"Sulfuras, Hand of Ragnaros", 0, 80}}},
		{name: "#5", inputs: []*Item{{"Sulfuras, Hand of Ragnaros", -1, 80}}, outputs: []*Item{{"Sulfuras, Hand of Ragnaros", -1, 80}}},
		//{name: "#9",inputs: []*Item{{"Conjured Mana Cake", 3, 6}},  outputs: []*Item{{"Conjured Mana Cake", 3-1, 6-2}}}, // TODO: implement new requirement
	}

	for _, c := range cases {
		UpdateQuality(c.inputs)

		if !reflect.DeepEqual(c.inputs, c.outputs) {
			t.Errorf("%s not match : \n%#v", c.name, pretty.Diff(c.inputs, c.outputs))
		}
	}
}

func TestBackstage(t *testing.T) {
	cases := []struct {
		name    string
		inputs  []*Item
		outputs []*Item
	}{
		{name: "#1", inputs: []*Item{{"Backstage passes to a TAFKAL80ETC concert", 15, 1}}, outputs: []*Item{{"Backstage passes to a TAFKAL80ETC concert", 14, 2}}},
		{name: "#2", inputs: []*Item{{"Backstage passes to a TAFKAL80ETC concert", 15, 49}}, outputs: []*Item{{"Backstage passes to a TAFKAL80ETC concert", 14, 50}}},
		{name: "#3", inputs: []*Item{{"Backstage passes to a TAFKAL80ETC concert", 15, 50}}, outputs: []*Item{{"Backstage passes to a TAFKAL80ETC concert", 14, 50}}},
		{name: "#4", inputs: []*Item{{"Backstage passes to a TAFKAL80ETC concert", 11, 20}}, outputs: []*Item{{"Backstage passes to a TAFKAL80ETC concert", 10, 21}}},

		{name: "#5", inputs: []*Item{{"Backstage passes to a TAFKAL80ETC concert", 10, 48}}, outputs: []*Item{{"Backstage passes to a TAFKAL80ETC concert", 9, 50}}},
		{name: "#6", inputs: []*Item{{"Backstage passes to a TAFKAL80ETC concert", 10, 49}}, outputs: []*Item{{"Backstage passes to a TAFKAL80ETC concert", 9, 50}}},
		{name: "#7", inputs: []*Item{{"Backstage passes to a TAFKAL80ETC concert", 10, 20}}, outputs: []*Item{{"Backstage passes to a TAFKAL80ETC concert", 9, 22}}},
		{name: "#8", inputs: []*Item{{"Backstage passes to a TAFKAL80ETC concert", 9, 20}}, outputs: []*Item{{"Backstage passes to a TAFKAL80ETC concert", 8, 22}}},
		{name: "#9", inputs: []*Item{{"Backstage passes to a TAFKAL80ETC concert", 6, 20}}, outputs: []*Item{{"Backstage passes to a TAFKAL80ETC concert", 5, 22}}},

		{name: "#10", inputs: []*Item{{"Backstage passes to a TAFKAL80ETC concert", 5, 47}}, outputs: []*Item{{"Backstage passes to a TAFKAL80ETC concert", 4, 50}}},
		{name: "#11", inputs: []*Item{{"Backstage passes to a TAFKAL80ETC concert", 5, 48}}, outputs: []*Item{{"Backstage passes to a TAFKAL80ETC concert", 4, 50}}},
		{name: "#12", inputs: []*Item{{"Backstage passes to a TAFKAL80ETC concert", 5, 49}}, outputs: []*Item{{"Backstage passes to a TAFKAL80ETC concert", 4, 50}}},
		{name: "#13", inputs: []*Item{{"Backstage passes to a TAFKAL80ETC concert", 4, 47}}, outputs: []*Item{{"Backstage passes to a TAFKAL80ETC concert", 3, 50}}},
		{name: "#14", inputs: []*Item{{"Backstage passes to a TAFKAL80ETC concert", 1, 47}}, outputs: []*Item{{"Backstage passes to a TAFKAL80ETC concert", 0, 50}}},

		{name: "#15", inputs: []*Item{{"Backstage passes to a TAFKAL80ETC concert", 0, 47}}, outputs: []*Item{{"Backstage passes to a TAFKAL80ETC concert", -1, 0}}},
		{name: "#16", inputs: []*Item{{"Backstage passes to a TAFKAL80ETC concert", -1, 47}}, outputs: []*Item{{"Backstage passes to a TAFKAL80ETC concert", -2, 0}}},
	}

	for _, c := range cases {
		UpdateQuality(c.inputs)

		if !reflect.DeepEqual(c.inputs, c.outputs) {
			t.Errorf("%s not match : \n%#v", c.name, pretty.Diff(c.inputs, c.outputs))
		}
	}
}

func TestAgedBrie(t *testing.T) {
	cases := []struct {
		name    string
		inputs  []*Item
		outputs []*Item
	}{
		{name: "#1", inputs: []*Item{{"Aged Brie", 2, 50}}, outputs: []*Item{{"Aged Brie", 1, 50}}},
		{name: "#2", inputs: []*Item{{"Aged Brie", 2, 49}}, outputs: []*Item{{"Aged Brie", 1, 50}}},
		{name: "#3", inputs: []*Item{{"Aged Brie", 2, 48}}, outputs: []*Item{{"Aged Brie", 1, 49}}},
		{name: "#4", inputs: []*Item{{"Aged Brie", 2, 20}}, outputs: []*Item{{"Aged Brie", 1, 21}}},
		{name: "#5", inputs: []*Item{{"Aged Brie", 1, 20}}, outputs: []*Item{{"Aged Brie", 0, 21}}},
		{name: "#6", inputs: []*Item{{"Aged Brie", 0, 20}}, outputs: []*Item{{"Aged Brie", -1, 22}}},
		{name: "#7", inputs: []*Item{{"Aged Brie", -1, 20}}, outputs: []*Item{{"Aged Brie", -2, 22}}},
		{name: "#8", inputs: []*Item{{"Aged Brie", -2, 48}}, outputs: []*Item{{"Aged Brie", -3, 50}}},
		{name: "#9", inputs: []*Item{{"Aged Brie", -2, 49}}, outputs: []*Item{{"Aged Brie", -3, 50}}},
		{name: "#10", inputs: []*Item{{"Aged Brie", -2, 50}}, outputs: []*Item{{"Aged Brie", -3, 50}}},
	}

	for _, c := range cases {
		UpdateQuality(c.inputs)

		if !reflect.DeepEqual(c.inputs, c.outputs) {
			t.Errorf("%s not match : \n%#v", c.name, pretty.Diff(c.inputs, c.outputs))
		}
	}
}

func TestNormalItem(t *testing.T) {
	cases := []struct {
		name    string
		inputs  []*Item
		outputs []*Item
	}{
		{name: "#1", inputs: []*Item{{"+5 Dexterity Vest", 10, 20}}, outputs: []*Item{{"+5 Dexterity Vest", 9, 19}}},
		{name: "#2", inputs: []*Item{{"+5 Dexterity Vest", 1, 20}}, outputs: []*Item{{"+5 Dexterity Vest", 0, 19}}},
		{name: "#3", inputs: []*Item{{"+5 Dexterity Vest", 0, 20}}, outputs: []*Item{{"+5 Dexterity Vest", -1, 18}}},
		{name: "#3", inputs: []*Item{{"+5 Dexterity Vest", -1, 20}}, outputs: []*Item{{"+5 Dexterity Vest", -2, 18}}},
	}

	for _, c := range cases {
		UpdateQuality(c.inputs)

		if !reflect.DeepEqual(c.inputs, c.outputs) {
			t.Errorf("%s not match : \n%#v", c.name, pretty.Diff(c.inputs, c.outputs))
		}
	}
}
