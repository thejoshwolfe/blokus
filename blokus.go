package main

const BoardSize = 20

// order of players
const (
	Blue = iota
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
