package blokus

const BoardSize = 20

// order of players
const (
	Empty = iota
	Blue
	Yellow
	Red
	Green
)

var Tiles = [][][]bool{
	{{true}},

	{{true, true}},

	{{true, true, true}},
	{
		{true, true},
		{true, false},
	},

	{{true, true, true, true}},
	{
		{true, true, true},
		{true, false, false},
	},
	{
		{true, true, true},
		{false, true, false},
	},
	{
		{true, true, false},
		{false, true, true},
	},
	{
		{true, true},
		{true, true},
	},

	{{true, true, true, true, true}},
	{
		{true, true, true, true},
		{true, false, false, false},
	},
	{
		{true, true, true, true},
		{false, true, false, false},
	},
	{
		{true, true, true, false},
		{false, false, true, true},
	},
	{
		{true, true, true},
		{true, true, false},
	},
	{
		{true, true, true},
		{true, false, true},
	},
	{
		{true, true, true},
		{true, false, false},
		{true, false, false},
	},
	{
		{true, true, true},
		{false, true, false},
		{false, true, false},
	},
	{
		{true, true, false},
		{false, true, true},
		{false, true, false},
	},
	{
		{true, true, false},
		{false, true, true},
		{false, false, true},
	},
	{
		{true, true, false},
		{false, true, false},
		{false, true, true},
	},
	{
		{false, true, false},
		{true, true, true},
		{false, true, false},
	},
}

var Board [BoardSize][BoardSize]int
