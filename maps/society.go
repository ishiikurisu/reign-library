package maps

import "math/rand"

type Institution struct {
    What string
    Where []int
    Memory string
    Script string
}

// Updates the society `s` living on the world `m`
func Tick(world [][]Block, institutions []Institution) []Institution {
    rand.Shuffle(len(institutions), func(i, j int) {
        institutions[i], institutions[j] = institutions[j], institutions[i]
    })

    for i, institution := range institutions {
        institutions
    }

    return s
}
