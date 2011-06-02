package main

const BoardSize = 20

// order of players
const (
	Blue = 0
	Yellow
	Red
	Green
)

var Tiles = [][][]bool {
	[][]bool{
        []bool{true},
    },
	[][]bool{
        []bool{true, true},
    },

	[][]bool{
        []bool{true, true, true},
    },
	[][]bool{
		[]bool{true, true},
		[]bool{true, false},
	},
	[][]bool{
        []bool{true, true, true, true},
    },
	[][]bool{
		[]bool{true, true, true},
		[]bool{true, false, false},
	},
	[][]bool{
		[]bool{true, true, true},
		[]bool{false, true, false},
	},
	[][]bool{
		[]bool{true, true, false},
		[]bool{false, true, true},
	},
	[][]bool{
		[]bool{true, true},
		[]bool{true, true},
	},

	[][]bool{
        []bool{true, true, true, true, true},
    },
	[][]bool{
		[]bool{true, true, true, true},
		[]bool{true, false, false, false},
	},
	[][]bool{
		[]bool{true, true, true, true},
		[]bool{false, true, false, false},
	},
	[][]bool{
		[]bool{true, true, true, false},
		[]bool{false, false, true, true},
	},
	[][]bool{
		[]bool{true, true, true},
		[]bool{true, true, false},
	},
	[][]bool{
		[]bool{true, true, true},
		[]bool{true, false, true},
	},
	[][]bool{
		[]bool{true, true, true},
		[]bool{true, false, false},
		[]bool{true, false, false},
	},
	[][]bool{
		[]bool{true, true, true},
		[]bool{false, true, false},
		[]bool{false, true, false},
	},
	[][]bool{
		[]bool{true, true, false},
		[]bool{false, true, true},
		[]bool{false, true, false},
	},
	[][]bool{
		[]bool{true, true, false},
		[]bool{false, true, true},
		[]bool{false, false, true},
	},
	[][]bool{
		[]bool{true, true, false},
		[]bool{false, true, false},
		[]bool{false, true, true},
	},
	[][]bool{
		[]bool{false, true, false},
		[]bool{true, true, true},
		[]bool{false, true, false},
	},
}
