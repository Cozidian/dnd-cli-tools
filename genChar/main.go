package main

import (
	"fmt"
	"math/rand"
	"strings"

	"github.com/spf13/pflag"
)

type Character struct {
	Name string
	Race string
	Class string
}

func findLongest(a, b, c int) int {
	if a > b && a > c {
		return a
	} else if b > c {
		return b
	} else {
		return c
	}
}

func printHorizontalLine(l int) {
	for i := 0; i < l; i++ {
		fmt.Printf("-")
	}
	fmt.Printf("\n")
}

func printLine(name string, displayName string, longest int) {
	if len(name) < longest {
		fmt.Printf("| %s | %s", displayName, name)
		fmt.Printf("%s |\n", strings.Repeat(" ", longest-len(name)))
	} else {
		fmt.Printf("| %s | %s |\n", displayName, name)
	}
}

func (c *Character) Print() {
	l1 := len(c.Name)
	l2 := len(c.Race)
	l3 := len(c.Class)
	longest := findLongest(l1, l2, l3)
	printHorizontalLine(longest + 12)
	printLine(c.Name, "Name ", longest)
	printLine(c.Race, "Race ", longest)
	printLine(c.Class, "Class", longest)
	printHorizontalLine(longest + 12)
}

// keys returns the keys of a map as a slice of strings since maps are not ordered and indexed
func keys(m map[string]struct{}) []string {
	s := make([]string, 0, len(m))
	for k := range m {
		s = append(s, k)
	}
	return s
}

func namePrefix(race string) string {
	elf := []string{"El", "Ae", "Ly", "Fi", "Ca", "Si", "Ra", "Fa", "Na", "Li", "My", "Ta", "Lae", "Va", "Zy"}
	dwarf := []string{"Thor", "Dur", "Bal", "Gar", "Gim", "Bro", "Thra", "Kaz", "Brim", "Dol", "Hra", "Kre", "Nor", "Zul", "Vor"}
	orc := []string{"Gro", "Dra", "Thok", "Mar", "Kor", "Bra", "Ur", "Gra", "Krog", "Vul", "Rok", "Zag", "Thar", "Krul", "Bor"}
	human := []string{"Al", "Be", "Da", "Ev", "Ma", "Pe", "Re", "Jo", "An", "Mi", "Sa", "Te", "Le", "Vi", "Mo"}
	halfling := []string{"Bil", "Fil", "Tol", "Mar", "Per", "Pip", "San", "Ric", "Bar", "Hal", "Kit", "Lun", "Mel", "Pod", "Tam"}
	gnome := []string{"Gim", "Fin", "Bin", "Zil", "Mer", "Wren", "Ban", "Nim", "Quin", "Wil", "Fen", "Lor", "Zin", "Tink", "Zan"}
	dragonborn := []string{"Arj", "Bal", "Tor", "Nor", "Vex", "Xan", "Kor", "Mir", "Rai", "Zev", "Jar", "Kriv", "Zor", "Har", "Dar"}
	thiefling := []string{"Az", "Bel", "Dav", "Fen", "Gol", "Hek", "Kis", "Lor", "Mir", "Nex", "Riz", "Sar", "Tor", "Vul", "Xan"}


	// select the prefix based on race and retun a random one
	switch race {
	case "elf":
		return elf[rand.Intn(len(elf))]
	case "dwarf":
		return dwarf[rand.Intn(len(dwarf))]
	case "orc":
		return orc[rand.Intn(len(orc))]
	case "human":
		return human[rand.Intn(len(human))]
	case "halfling":
		return halfling[rand.Intn(len(halfling))]
	case "gnome":
		return gnome[rand.Intn(len(gnome))]
	case "dragonborn":
		return dragonborn[rand.Intn(len(dragonborn))]
	case "thiefling":
		return thiefling[rand.Intn(len(thiefling))]
	default:
		return ""
	}
}

func nameRoot(root string) string{
	nature := []string{"thorn", "leaf", "river", "star", "moon", "sun", "stone", "tree", "bark", "stream", "blossom", "sky", "breeze", "moss", "vine"}

	mystical := []string{ "shadow", "light", "spell", "magic", "crystal", "rune", "spirit", "phantom", "wisp", "ether", "glyph", "charm", "essence", "portal", "aura"}

	warrior := []string{ "blade", "shield", "spear", "battle", "axe", "hammer", "warrior", "knight", "sword", "armor", "fight", "strike", "rage", "clash", "valor"}

	noble := []string{ "royal", "crown", "noble", "lord", "king", "queen", "prince", "duke", "regal", "manor", "scepter", "throne", "herald", "rule", "sovereign", "heir"}

	switch root {
	case "nature":
		return nature[rand.Intn(len(nature))]
	case "mystical":
		return mystical[rand.Intn(len(mystical))]
	case "warrior":
		return warrior[rand.Intn(len(warrior))]
	case "noble":
		return noble[rand.Intn(len(noble))]
	default:
		return ""
	}
}

func nameSulfix(race string) string {
	elf := []string{"ar", "ia", "el", "ian", "ae", "ien", "ir", "or", "ith", "eal", "wyn", "ala", "lys", "ial", "en"}
	dwarf := []string{"in", "or", "ur", "gim", "son", "dar", "ok", "th", "grum", "ik", "ord", "dur", "mir", "ol", "lin"}
	orc := []string{"gash", "thak", "nar", "guk", "ra", "ul", "zag", "or", "rok", "tok", "var", "zug", "ak", "bor", "dra"}
	human := []string{"er", "an", "ia", "or", "el", "is", "os", "us", "en", "al", "on", "as", "ir", "ar", "y"}
	halfling := []string{"en", "y", "kin", "bel", "lin", "o", "i", "le", "bie", "do", "da", "fa", "mo", "ta", "li"}
	gnome := []string{"wick", "waddle", "sprocket", "knackle", "pock", "fizz", "glim", "gax", "nik", "bum", "tan", "dap", "zop", "mik", "wip"}
	dragonborn := []string{"ax", "ion", "kan", "ir", "us", "ar", "ek", "ur", "ith", "om", "lor", "zen", "rox", "kir", "ton"}
	thiefling := []string{"zor", "ith", "dan", "ro", "us", "ax", "nar", "ir", "on", "tor", "is", "ak", "ex", "um", "or"}

	switch race {
	case "elf":
		return elf[rand.Intn(len(elf))]
	case "dwarf":
		return dwarf[rand.Intn(len(dwarf))]
	case "orc":
		return orc[rand.Intn(len(orc))]
	case "human":
		return human[rand.Intn(len(human))]
	case "halfling":
		return halfling[rand.Intn(len(halfling))]
	case "gnome":
		return gnome[rand.Intn(len(gnome))]
	case "dragonborn":
		return dragonborn[rand.Intn(len(dragonborn))]
	case "thiefling":
		return thiefling[rand.Intn(len(thiefling))]
	default:
		return ""
	}
}

func nameGenerator(race string) string {
	root := []string{"nature", "mystical", "warrior", "noble"}
	prefix := namePrefix(race)
	nameRoot := nameRoot(root[rand.Intn(len(root))] )
	sulfix := nameSulfix(race)
	return prefix + nameRoot + sulfix
}

func main() {
	dndClasses := map[string]struct{}{
		"fighter": {},
		"wizard": {},
		"cleric": {},
		"rogue": {},
		"ranger": {},
		"paladin": {},
		"barbarian": {},
		"monk": {},
		"druid": {},
		"sorcerer": {},
		"warlock": {},
		"bard": {},
	}
	dndRaces := map[string]struct{}{
		"elf": {},
		"dwarf": {},
		"orc": {},
		"human": {},
		"halfling": {},
		"gnome": {},
		"dragonborn": {},
		"thiefling": {},
	}

	char := Character{}

	name := pflag.StringP("name", "n", "", "character name")
	class := pflag.StringP("class", "c", "", "character class")
	race := pflag.StringP("race", "r", "", "character race")
	pflag.Parse()

	if _, ok := dndClasses[*class]; !ok {
		classKey := keys(dndClasses)
		*class = classKey[rand.Intn(len(dndClasses))]
	}

	if _, ok := dndRaces[*race]; !ok {
		classKey := keys(dndRaces)
		*race= classKey[rand.Intn(len(dndRaces))]
	}

	if *name == "" {
		*name = nameGenerator(*race)
	}

	char.Name = *name
	char.Race = *race
	char.Class = *class

	char.Print()
}
