package internal

import "time"

// Assume we are "and-ing" all of the search parameters
type SearchParams struct {
	RecipeName   string
	AuthorName   string
	UploaderName string
	PrepTime     time.Duration
	TotalTime    time.Duration
	Yeild        string
	Ingredients  []string
}

type SearchResult struct {
	RecipeName string
	RecipeId   ID
}

type RecipeSearch interface {
	Search(params SearchParams) []SearchResult
}

type TestSearch struct{}

func (t TestSearch) Search(params SearchParams) []SearchResult {
	return []SearchResult{{
		RecipeName: "Test Recipe 1",
		RecipeId:   1,
	}, {
		RecipeName: "Test Recipe 2",
		RecipeId:   2,
	}, {
		RecipeName: "Test Recipe 3",
		RecipeId:   3,
	}, {
		RecipeName: "Test Recipe 4",
		RecipeId:   4,
	}, {
		RecipeName: "Test Recipe 5",
		RecipeId:   5,
	}, {
		RecipeName: "Test Recipe 6",
		RecipeId:   6,
	}, {
		RecipeName: "Test Recipe 7",
		RecipeId:   7,
	}, {
		RecipeName: "Test Recipe 8",
		RecipeId:   8,
	}, {
		RecipeName: "Test Recipe 9",
		RecipeId:   9,
	}, {
		RecipeName: "Test Recipe 10",
		RecipeId:   10,
	}, {
		RecipeName: "Test Recipe 11",
		RecipeId:   11,
	}, {
		RecipeName: "Test Recipe 12",
		RecipeId:   12,
	}, {
		RecipeName: "Test Recipe 13",
		RecipeId:   13,
	}, {
		RecipeName: "Test Recipe 14",
		RecipeId:   14,
	}, {
		RecipeName: "Test Recipe 15",
		RecipeId:   15,
	}, {
		RecipeName: "Test Recipe 16",
		RecipeId:   16,
	}, {
		RecipeName: "Test Recipe 17",
		RecipeId:   17,
	}, {
		RecipeName: "Test Recipe 18",
		RecipeId:   18,
	}, {
		RecipeName: "Test Recipe 19",
		RecipeId:   19,
	}, {
		RecipeName: "Test Recipe 20",
		RecipeId:   20,
	}, {
		RecipeName: "Test Recipe 21",
		RecipeId:   21,
	}, {
		RecipeName: "Test Recipe 22",
		RecipeId:   22,
	}, {
		RecipeName: "Test Recipe 23",
		RecipeId:   23,
	}, {
		RecipeName: "Test Recipe 24",
		RecipeId:   24,
	}, {
		RecipeName: "Test Recipe 25",
		RecipeId:   25,
	}, {
		RecipeName: "Test Recipe 26",
		RecipeId:   26,
	}, {
		RecipeName: "Test Recipe 27",
		RecipeId:   27,
	}, {
		RecipeName: "Test Recipe 28",
		RecipeId:   28,
	}, {
		RecipeName: "Test Recipe 29",
		RecipeId:   29,
	}, {
		RecipeName: "Test Recipe 30",
		RecipeId:   30,
	}, {
		RecipeName: "Test Recipe 31",
		RecipeId:   31,
	}, {
		RecipeName: "Test Recipe 32",
		RecipeId:   32,
	}, {
		RecipeName: "Test Recipe 33",
		RecipeId:   33,
	}, {
		RecipeName: "Test Recipe 34",
		RecipeId:   34,
	}, {
		RecipeName: "Test Recipe 35",
		RecipeId:   35,
	}, {
		RecipeName: "Test Recipe 36",
		RecipeId:   36,
	}, {
		RecipeName: "Test Recipe 37",
		RecipeId:   37,
	}, {
		RecipeName: "Test Recipe 38",
		RecipeId:   38,
	}, {
		RecipeName: "Test Recipe 39",
		RecipeId:   39,
	}, {
		RecipeName: "Test Recipe 40",
		RecipeId:   40,
	}, {
		RecipeName: "Test Recipe 41",
		RecipeId:   41,
	}, {
		RecipeName: "Test Recipe 42",
		RecipeId:   42,
	}, {
		RecipeName: "Test Recipe 43",
		RecipeId:   43,
	}, {
		RecipeName: "Test Recipe 44",
		RecipeId:   44,
	}, {
		RecipeName: "Test Recipe 45",
		RecipeId:   45,
	}, {
		RecipeName: "Test Recipe 46",
		RecipeId:   46,
	}, {
		RecipeName: "Test Recipe 47",
		RecipeId:   47,
	}, {
		RecipeName: "Test Recipe 48",
		RecipeId:   48,
	}, {
		RecipeName: "Test Recipe 49",
		RecipeId:   49,
	}, {
		RecipeName: "Test Recipe 50",
		RecipeId:   50,
	}, {
		RecipeName: "Test Recipe 51",
		RecipeId:   51,
	}, {
		RecipeName: "Test Recipe 52",
		RecipeId:   52,
	}, {
		RecipeName: "Test Recipe 53",
		RecipeId:   53,
	}, {
		RecipeName: "Test Recipe 54",
		RecipeId:   54,
	}, {
		RecipeName: "Test Recipe 55",
		RecipeId:   55,
	}, {
		RecipeName: "Test Recipe 56",
		RecipeId:   56,
	}, {
		RecipeName: "Test Recipe 57",
		RecipeId:   57,
	}, {
		RecipeName: "Test Recipe 58",
		RecipeId:   58,
	}, {
		RecipeName: "Test Recipe 59",
		RecipeId:   59,
	}, {
		RecipeName: "Test Recipe 60",
		RecipeId:   60,
	}, {
		RecipeName: "Test Recipe 61",
		RecipeId:   61,
	}, {
		RecipeName: "Test Recipe 62",
		RecipeId:   62,
	}, {
		RecipeName: "Test Recipe 63",
		RecipeId:   63,
	}, {
		RecipeName: "Test Recipe 64",
		RecipeId:   64,
	}, {
		RecipeName: "Test Recipe 65",
		RecipeId:   65,
	}, {
		RecipeName: "Test Recipe 66",
		RecipeId:   66,
	}, {
		RecipeName: "Test Recipe 67",
		RecipeId:   67,
	}, {
		RecipeName: "Test Recipe 68",
		RecipeId:   68,
	}, {
		RecipeName: "Test Recipe 69",
		RecipeId:   69,
	}, {
		RecipeName: "Test Recipe 70",
		RecipeId:   70,
	}, {
		RecipeName: "Test Recipe 71",
		RecipeId:   71,
	}, {
		RecipeName: "Test Recipe 72",
		RecipeId:   72,
	}, {
		RecipeName: "Test Recipe 73",
		RecipeId:   73,
	}, {
		RecipeName: "Test Recipe 74",
		RecipeId:   74,
	}, {
		RecipeName: "Test Recipe 75",
		RecipeId:   75,
	}, {
		RecipeName: "Test Recipe 76",
		RecipeId:   76,
	}, {
		RecipeName: "Test Recipe 77",
		RecipeId:   77,
	}, {
		RecipeName: "Test Recipe 78",
		RecipeId:   78,
	}, {
		RecipeName: "Test Recipe 79",
		RecipeId:   79,
	}, {
		RecipeName: "Test Recipe 80",
		RecipeId:   80,
	}, {
		RecipeName: "Test Recipe 81",
		RecipeId:   81,
	}, {
		RecipeName: "Test Recipe 82",
		RecipeId:   82,
	}, {
		RecipeName: "Test Recipe 83",
		RecipeId:   83,
	}, {
		RecipeName: "Test Recipe 84",
		RecipeId:   84,
	}, {
		RecipeName: "Test Recipe 85",
		RecipeId:   85,
	}, {
		RecipeName: "Test Recipe 86",
		RecipeId:   86,
	}, {
		RecipeName: "Test Recipe 87",
		RecipeId:   87,
	}, {
		RecipeName: "Test Recipe 88",
		RecipeId:   88,
	}, {
		RecipeName: "Test Recipe 89",
		RecipeId:   89,
	}, {
		RecipeName: "Test Recipe 90",
		RecipeId:   90,
	}, {
		RecipeName: "Test Recipe 91",
		RecipeId:   91,
	}, {
		RecipeName: "Test Recipe 92",
		RecipeId:   92,
	}, {
		RecipeName: "Test Recipe 93",
		RecipeId:   93,
	}, {
		RecipeName: "Test Recipe 94",
		RecipeId:   94,
	}, {
		RecipeName: "Test Recipe 95",
		RecipeId:   95,
	}, {
		RecipeName: "Test Recipe 96",
		RecipeId:   96,
	}, {
		RecipeName: "Test Recipe 97",
		RecipeId:   97,
	}, {
		RecipeName: "Test Recipe 98",
		RecipeId:   98,
	}, {
		RecipeName: "Test Recipe 99",
		RecipeId:   99,
	}, {
		RecipeName: "Test Recipe 100",
		RecipeId:   100,
	}, {
		RecipeName: "Test Recipe 101",
		RecipeId:   101,
	}, {
		RecipeName: "Test Recipe 102",
		RecipeId:   102,
	}, {
		RecipeName: "Test Recipe 103",
		RecipeId:   103,
	}, {
		RecipeName: "Test Recipe 104",
		RecipeId:   104,
	}, {
		RecipeName: "Test Recipe 105",
		RecipeId:   105,
	}, {
		RecipeName: "Test Recipe 106",
		RecipeId:   106,
	}, {
		RecipeName: "Test Recipe 107",
		RecipeId:   107,
	}, {
		RecipeName: "Test Recipe 108",
		RecipeId:   108,
	}, {
		RecipeName: "Test Recipe 109",
		RecipeId:   109,
	}, {
		RecipeName: "Test Recipe 110",
		RecipeId:   110,
	}, {
		RecipeName: "Test Recipe 111",
		RecipeId:   111,
	}, {
		RecipeName: "Test Recipe 112",
		RecipeId:   112,
	}, {
		RecipeName: "Test Recipe 113",
		RecipeId:   113,
	}, {
		RecipeName: "Test Recipe 114",
		RecipeId:   114,
	}, {
		RecipeName: "Test Recipe 115",
		RecipeId:   115,
	}, {
		RecipeName: "Test Recipe 116",
		RecipeId:   116,
	}, {
		RecipeName: "Test Recipe 117",
		RecipeId:   117,
	}, {
		RecipeName: "Test Recipe 118",
		RecipeId:   118,
	}, {
		RecipeName: "Test Recipe 119",
		RecipeId:   119,
	}, {
		RecipeName: "Test Recipe 120",
		RecipeId:   120,
	}, {
		RecipeName: "Test Recipe 121",
		RecipeId:   121,
	}, {
		RecipeName: "Test Recipe 122",
		RecipeId:   122,
	}, {
		RecipeName: "Test Recipe 123",
		RecipeId:   123,
	}, {
		RecipeName: "Test Recipe 124",
		RecipeId:   124,
	}}
}
