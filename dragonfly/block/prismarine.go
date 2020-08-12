package block

import (
	"github.com/df-mc/dragonfly/dragonfly/block/prismarine"
	"github.com/df-mc/dragonfly/dragonfly/item"
	"github.com/df-mc/dragonfly/dragonfly/world"
)

// Prismarine is a  decorative solid type of block there are 3 types of prismarine block
// normal prismarine,dark prismarine and prismarine brick
type Prismarine struct {
	noNBT
	solid
	// prismarine type normal,dark,bricks
	prismarine prismarine.Prismarine
}

// BreakInfo ...
func (p Prismarine) BreakInfo() BreakInfo {
	return BreakInfo{
		Hardness:    1.5,
		Harvestable: pickaxeHarvestable,
		Effective:   pickaxeEffective,
		Drops:       simpleDrops(item.NewStack(p, 1)),
	}
}

// EncodeItem ...
func (p Prismarine) EncodeItem() (id int32, meta int16) {
	switch p.prismarine {
	case prismarine.Default():
		return 168, 0

	case prismarine.Dark():
		return 168, 1

	case prismarine.Bricks():
		return 168, 2
	}
	panic("invalid prismarine type")
}

// EncodeBlock ...
func (p Prismarine) EncodeBlock() (name string, properties map[string]interface{}) {
	return "minecraft:prismarine", map[string]interface{}{"prismarine_block_type": p.prismarine.String()}
}

// Hash ...
func (p Prismarine) Hash() uint64 {
	return hashPrismarine | (uint64(p.prismarine.Uint8()) << 32)
}

func allPrismarine() []world.Block {
	return []world.Block{
		Prismarine{prismarine: prismarine.Default()},
		Prismarine{prismarine: prismarine.Dark()},
		Prismarine{prismarine: prismarine.Bricks()},
	}
}
