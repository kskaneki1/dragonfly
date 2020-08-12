package item

//Prismarine shards are dropped upon death of guardians or elder guardians.
//From these mobs, 0–2 shards are dropped each time.
//This can be increased to a maximum of 5 with the Looting enchantment.
// via https://minecraft.gamepedia.com/

type PrismarineShard struct{}

// EncodeItem ...
func (PrismarineShard) EncodeItem() (id int32, meta int16) {
	return 490, 0
}
