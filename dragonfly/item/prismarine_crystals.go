package item

// DROPS
// Guardians and elder guardians have a 40% chance and 1⁄3 chance, respectively, of dropping prismarine crystals upon death.
//The maximum drop count is increased by one per level of Looting.

// Prismarine crystals are dropped by sea lanterns when not using a Silk Touch tool.
// They drop 2–3 crystals each time, which can be increased to a maximum of 5 using the Fortune enchantment.
// https://minecraft.gamepedia.com/Prismarine_Crystals

type PrismarineCrystals struct{}

// EncodeItem ...
func (PrismarineCrystals) EncodeItem() (id int32, meta int16) {
	return 422, 0
}
