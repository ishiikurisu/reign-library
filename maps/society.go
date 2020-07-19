package maps

import (
    "fmt"
    "strconv"
    "strings"
)

type Institution struct {
    // institution's name
    What string
    // institution's position on world
    Where []int
    // javascript object holding institution's memory
    Memory string
}

// Updates the society `s` living on the world `m`
func Tick(world [][]Block, institutions []Institution) []Institution {
    functions := map[string]func([][]Block, []Institution, int)[]Institution {
        "farm": updateFarm,
        "house": updateHouse,
    }

    for i := 0; i < len(institutions); i++ {
        institution := institutions[i]
        institutions = functions[institution.What](world, institutions, i)
    }

    return institutions
}

func updateFarm(world [][]Block, society[]Institution, index int) []Institution {
    farm := society[index]
    cropString := strings.Split(strings.Trim(farm.Memory, "\n"), ":")[1]
    crop, oops := strconv.Atoi(cropString)

    if oops == nil {
        farm.Memory = fmt.Sprintf("crop:%d\n", crop + 1)
    }

    society[index] = farm
    return society
}

func updateHouse(world [][]Block, society[]Institution, house int) []Institution {
    return society
}
