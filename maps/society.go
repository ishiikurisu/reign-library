package maps

import (
    "math/rand"
    "syscall/js"
)

type Institution struct {
    // institution's name
    What string
    // institution's position on world
    Where []int
    // javascript object holding object memory
    Memory js.Value
    // javascript function to be called to update this institution
    Script js.Value
}

// Updates the society `s` living on the world `m`
func tick(world [][]Block, institutions []Institution) []Institution {
    rand.Shuffle(len(institutions), func(i, j int) {
        institutions[i], institutions[j] = institutions[j], institutions[i]
    })

    for i, institution := range institutions {
        institution.Script.Invoke(ValueOf(world), ValueOf(institutions), i)
    }

    return institutions
}
