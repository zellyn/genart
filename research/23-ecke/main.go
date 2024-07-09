package main

import (
	"cmp"
	"fmt"
	"os"
	"slices"
	"strings"
)

const (
	JS1 = 1306859721
	JS2 = 1485627309
	JS3 = 1649173265
	JS4 = 1805297143
	JS5 = 1973195467
	JS6 = 2013911545
)

var bookLeanings = strings.ReplaceAll(strings.TrimSpace(`
............
..r.....bbrr
.rr.r.r..bbb
rr..bbrrrrr.
rrb..brbbb.r
.rrbrbr.bbbb
br..b.bbbrbb
`), "\n", "")

var plotLeanings = strings.ReplaceAll(strings.TrimSpace(`
............
....rr.rb.r.
brr..r.rrb.r
brr.b..bbbrr
bbbbr..r...b
rbr..bbbbbbb
b.bbr..rbrbb
bbbbr.rr.rrr
`), "\n", "")

var plotShifts = strings.ReplaceAll(strings.TrimSpace(`
............
............
............
r....r.....b
..rr.....r.b
rb....brb.r.
br.bbr.brb.b
bbr.b.r.b.rb
b....rrb.brb
rrrbbbbrrbrr
rrrrrbbrrrb.
`), "\n", "")

var bookShifts = strings.ReplaceAll(strings.TrimSpace(`
............
............
.......b..rb
rb.r..r..b..
bb.r.brrb..r
rrr.rrrb..rb
rrrrb.rbr.rr
rrr.brr.rrbb
bbrr.b..b..b
b..b....b..b
..bb......bb
`), "\n", "")

// var bookShifts = strings.ReplaceAll(strings.TrimSpace(`
// bbrbbrbbbrbb
// bbrbrrbbbrrb
// brbbbbbbrbrb
// rbbrbbrrrbbb
// bbbrrbrrbrbr
// rrrrrrrbrrrb
// rrrrbbrbrrrr
// rrrbbrrrrrbb
// bbrrrbbrbrrb
// brrbrrrrbrbb
// rrbbbrrbbrbb
// rbbbrbbrbbbr
// brrrrbrrbbbr
// brrbbbbrbbrb
// rbrrrrrbbbbr
// bbrrrrbrbbrr
// rrrrrrrrrbrb
// bbbbrbrbbbbb
// rrbrbbbbbrrb
// rbrrbrrrrbbr
// rbrbbrrbrbrr
// brbrrbbbrrrb
// `), "\n", "")

var leanings = plotLeanings
var shifts = bookShifts

type random uint64

type winner struct {
	extra uint64
	seed  uint64
}

const PI = 3.14159
const PI4T = PI / 4

var MOD uint64 = 1 << 31

// const MOD = 1 << 31

func newRandom(seed uint64) *random {
	r := random(seed)
	return &r
}

func (r *random) next(mul uint64, min, max float32) float32 {
	*r = random((uint64(*r) * mul) % MOD)
	res := float32(*r)/float32(MOD)*(max-min) + min
	return res
}

func (r *random) raw(mul uint64) uint64 {
	*r = random((uint64(*r) * mul) % MOD)
	return uint64(*r)
}

func computeFactor(mul uint64, times int) uint64 {
	res := uint64(1)
	for _ = range times {
		res = (res * mul) % MOD
	}
	return res
}

var backwards = false

func schotterShiftSearch() error {
	ITERS := (MOD + 5000000) / 10000000
	var winning []winner

	mul := uint64(5)

OUTER:
	for seed := uint64((MOD / 2) | 1); seed <= MOD; seed += 2 {
		print := seed == 1649173265 && false
		if print {
			fmt.Printf("Checking seed %d\n", seed)
		}
		errs := 0
		if (seed+1)%10_000_000 == 0 {
			fmt.Printf("mul=%d: %d/%d (%d)\n", mul, seed/10000000, ITERS, len(winning))
		}
		r := random(seed)

		i := float32(0)
		for _, want := range shifts {
			before := uint64(r)
			// r.next(mul, -1, 1)
			move_limit := 5 * i / 264
			// twist_limit := PI / 4 * i / 264

			y_offset := r.next(mul, -move_limit, move_limit)
			_ = y_offset
			x_offset := r.next(mul, -move_limit, move_limit)
			// angle := r.next(mul, PI/4-twist_limit, PI/4+twist_limit)
			after := uint64(r)
			i++
			if print {
				fmt.Printf("RANDOM: %d → %d: want: %c  val:%.16f\n", before, after, want, x_offset)
			}
			if want == '.' {
				continue
			}

			color := 'b'
			if x_offset > 0 {
				color = 'r'
			}

			if backwards {
				color ^= 16 // delicious!
			}

			if want != color {
				if print {
					fmt.Printf("wrong\n")
				}
				errs++
				if errs > 3 {
					continue OUTER
				}
			}
		}

		fmt.Printf("Winning seed: %d   (errs=%d)\n", seed, errs)
		winning = append(winning, winner{
			seed:  seed,
			extra: mul,
		})
	}

	if len(winning) > 0 {
		fmt.Println("All winning seeds:")
		for _, winner := range winning {
			fmt.Printf("mul: %d  seed: %d\n", winner.extra, winner.seed)
		}
	} else {
		fmt.Printf("No winning seeds for MOD=%d\n", MOD)
	}
	return nil
}

func reverse(moves string) string {
	return strings.Map(func(r rune) rune {
		switch r {
		case 'n':
			return 's'
		case 's':
			return 'n'
		case 'e':
			return 'w'
		case 'w':
			return 'e'
		case '?':
			return '?'
		default:
			panic(fmt.Sprintf("unexpected move character '%c'", r))
		}
	}, moves)
}

var hlmap = map[rune]rune{
	'e': 'h',
	'n': 'h',
	'w': 'l',
	's': 'l',
	'?': '?',
}

func lowHigh(s string, ns bool) string {
	rs := []rune(s)
	result := make([]rune, 0, len(rs)/2)
	start := 0
	if ns {
		start = 1
	}
	for i := start; i < len(rs); i += 2 {
		c := hlmap[rs[i]]
		if c == 0 {
			panic(fmt.Sprintf("weird input char in moves: '%c'", rs[i]))
		}
		result = append(result, c)
	}
	return string(result)
}

type trace struct {
	index int
	moves string
}

func main() {
	// if err := eckePrint(); err != nil {
	// 	fmt.Fprintf(os.Stderr, "Error: %v\n", err)
	// 	os.Exit(1)
	// }

	if err := eckeSearch(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}

func elirr(n int, r1, r2 *random) string {
	s := ""
	x0 := r1.raw(5)
	y0 := r2.raw(5)
	x, y := x0, y0

	for _ = range n {
		lastx, lasty := x, y
		x = r1.raw(5)
		y = r2.raw(5)
		if x > lastx {
			s += "e"
		} else {
			s += "w"
		}
		if y > lasty {
			s += "n"
		} else {
			s += "s"
		}
	}

	if x0 > x {
		s += "e"
	} else {
		s += "w"
	}

	if y0 > y {
		s += "n"
	} else {
		s += "s"
	}

	return s
}

func elirr1d(n int, r *random) string {
	s := ""
	x0 := r.raw(5)
	x := x0

	for _ = range n {
		lastx := x
		x = r.raw(5)
		if x > lastx {
			s += "h"
		} else {
			s += "l"
		}
	}

	if x0 > x {
		s += "h"
	} else {
		s += "l"
	}

	return s
}

func printMoves(n int, seed1 uint64, seed2 uint64, generators int, debugs []int) {
	jumpFactor := computeFactor(5, n+1)

	r1 := newRandom(seed1)
	r2 := newRandom(seed2)
	if generators == 1 {
		r2 = r1
	}

	for index := 1; index <= slices.Max(debugs); index++ {
		if generators == 1 {
			fmt.Printf("%d: seed=%d ", index, uint64(*r1))
		} else {
			fmt.Printf("%d: seed1=%d, seed2=%d ", index, uint64(*r1), uint64(*r2))
		}
		if slices.Contains(debugs, index) {
			fmt.Println(elirr(n, r1, r2))
		} else {
			r1.raw(jumpFactor)
			r2.raw(jumpFactor)
			fmt.Println()
		}
	}
	fmt.Printf("printMoves jumpFactor=%d\n", jumpFactor)
}

func printMoves1D(n int, seed1 uint64, debugs []int) {
	jumpFactor := computeFactor(5, n+1)

	r1 := newRandom(seed1)

	for index := 1; index <= slices.Max(debugs); index++ {
		fmt.Printf("%d: seed=%d ", index, uint64(*r1))
		if slices.Contains(debugs, index) {
			fmt.Println(elirr1d(n, r1))
		} else {
			r1.raw(jumpFactor)
			fmt.Println()
		}
	}
	fmt.Printf("printMoves jumpFactor=%d\n", jumpFactor)
}

func eckePrint() error {

	// printMoves(11, JS2, JS5, 1, []int{35, 71, 73, 84})
	printMoves(11, JS2, JS5, 2, []int{35, 71, 73, 84})
	printMoves1D(11, JS2, []int{35, 71, 73, 84})
	printMoves1D(11, JS5, []int{35, 71, 73, 84})

	return nil
}

// JS2
var example1gen = []trace{
	{index: 35, moves: "esenwnesenwneswnwswsenws"},
	{index: 71, moves: "eswneswneswswswsenwsenws"},
	{index: 73, moves: "wnesenwsenwswswsenenwnes"},
	{index: 84, moves: "wneswswswswsenwswneswsen"},
}

// JS2, JS5
var example2gen = []trace{
	{index: 35, moves: "esenwswnwseneswsenwswsen"},
	{index: 71, moves: "enenenwswsesenwsenwswswn"},
	{index: 73, moves: "wnwswsenwswnenwsenwnesen"},
	{index: 84, moves: "enwnesenenwswnenwnwswnws"},
}

var actual = []trace{
	{index: 35, moves: "wnwsenwswneswnwnesenesen"},
	{index: 71, moves: "eswneneswswneswnwsenenws"},
	{index: 73, moves: "wneswseneswnenwsenwsesws"},
	{index: 84, moves: "wnesenwswneswnenws?nwnes"},
}

var actualReversed = []trace{
	{index: 281 - 35, moves: reverse("wnwsenwswneswnwnesenesen")},
	{index: 281 - 71, moves: reverse("eswneneswswneswnwsenenws")},
	{index: 281 - 73, moves: reverse("wneswseneswnenwsenwsesws")},
	{index: 281 - 84, moves: reverse("wnesenwswneswnenws?nwnes")},
}

func searchSingleRandom(traces []trace) {
	n := 11
	// Factor to jump past one polygon. Double things, since we're using
	// only one random generator.
	singlePolygonFactor := computeFactor(5, 2*(n+1))
	// fmt.Printf("searchSingleRandom: singlePolygonFactor=%d\n", singlePolygonFactor)

	slices.SortFunc(traces, func(a, b trace) int {
		return cmp.Compare(a.index, b.index)
	})

	var jumps []uint64
	lastIndex := 0
	for _, tr := range traces {
		jumpCount := tr.index - lastIndex - 1
		lastIndex = tr.index
		// fmt.Printf("jumpCount=%d\n", jumpCount)
		jumps = append(jumps, computeFactor(singlePolygonFactor, jumpCount))
	}

	// fmt.Printf("%v\n", traces)
	// fmt.Printf("%v\n", jumps)

	var winning []uint64
	ITERS := (MOD + 5000000) / 10000000
OUTER:
	// for seed := uint64((MOD / 2) | 1); seed <= MOD; seed += 2 {
	for seed := uint64(1); seed < MOD; seed += 2 {
		// fmt.Printf("starting seed: %d\n", seed)
		errs := 0
		if (seed+1)%10_000_000 == 0 {
			fmt.Printf(" %d/%d (%d)\n", seed/10000000, ITERS, len(winning))
		}
		r := newRandom(seed)

		for i, tr := range traces {
			// Jump forward to the polygon we're checking.
			r.raw(jumps[i])

			// fmt.Printf(" jumped to %d\n", uint64(*r))

			got := elirr(n, r, r)
			if !eq(got, tr.moves) {
				errs++
				if errs > 2 {
					continue OUTER
				}
				continue
			}
			fmt.Printf("%d", i)
		}

		fmt.Printf(" Winning seed: %d   (errs=%d)\n", seed, errs)
		winning = append(winning, seed)
	}
	fmt.Println()

	if len(winning) > 0 {
		fmt.Println("All winning seeds:")
		for _, winner := range winning {
			fmt.Printf(" %d\n", winner)
		}
	} else {
		fmt.Printf("No winning seeds.\n")
	}
}

func eq(s1, s2 string) bool {
	for i := len(s1) - 1; i >= 0; i-- {
		c1, c2 := s1[i], s2[i]
		if c1 == c2 || c1 == '?' || c2 == '?' {
			continue
		}
		return false
	}

	return true
}

func searchOneDimensionRandom(traces []trace, ns bool) {
	n := 11
	// Factor to jump past one polygon.
	singlePolygonFactor := computeFactor(5, n+1)
	// fmt.Printf("searchOneDimensionRandom: singlePolygonFactor=%d\n", singlePolygonFactor)

	slices.SortFunc(traces, func(a, b trace) int {
		return cmp.Compare(a.index, b.index)
	})

	traces = slices.Clone(traces)
	for i, tr := range traces {
		traces[i].moves = lowHigh(tr.moves, ns)
	}

	// fmt.Printf("%v\n", traces)

	var jumps []uint64
	lastIndex := 0
	for _, tr := range traces {
		jumpCount := tr.index - lastIndex - 1
		lastIndex = tr.index
		// fmt.Printf("jumpCount=%d\n", jumpCount)
		jumps = append(jumps, computeFactor(singlePolygonFactor, jumpCount))
	}

	// fmt.Printf("%v\n", traces)
	// fmt.Printf("%v\n", jumps)

	var winning []uint64
	ITERS := (MOD + 5000000) / 10000000
OUTER:
	for seed := uint64(1); seed < MOD; seed += 2 {
		// for seed := uint64(JS2); seed <= JS5; seed += JS5 - JS2 {

		// fmt.Printf("starting seed: %d\n", seed)
		errs := 0
		if (seed+1)%10_000_000 == 0 {
			fmt.Printf(" %d/%d (%d)\n", seed/10000000, ITERS, len(winning))
		}
		r := newRandom(seed)

		for i, tr := range traces {
			// Jump forward to the polygon we're checking.
			r.raw(jumps[i])

			// fmt.Printf(" jumped to %d\n", uint64(*r))

			got := elirr1d(n, r)
			if !eq(got, tr.moves) {
				errs++
				if errs > 1 {
					continue OUTER
				}
				continue
			}
			if i >= 2 {
				fmt.Printf("%d", i)
			}
		}

		fmt.Printf(" Winning seed: %d   (errs=%d)\n", seed, errs)
		winning = append(winning, seed)
	}
	fmt.Println()

	if len(winning) > 0 {
		fmt.Println("All winning seeds:")
		for _, winner := range winning {
			fmt.Printf(" %d\n", winner)
		}
	} else {
		fmt.Printf("No winning seeds.\n")
	}
}

func eckeSearch() error {
	// searchSingleRandom(example1gen)
	// fmt.Println("actual")
	// searchSingleRandom(actual)
	// fmt.Println("\nactual, reversed")
	// searchSingleRandom(actualReversed)
	// searchOneDimensionRandom(example2gen, false)
	// searchOneDimensionRandom(example2gen, true)
	// fmt.Println("1d actual E/W")
	// searchOneDimensionRandom(actual, false)
	// fmt.Println("\n1d actual N/S")
	// searchOneDimensionRandom(actual, true)
	// fmt.Println("\n1d actual reversed E/W")
	// searchOneDimensionRandom(actualReversed, false)
	// fmt.Println("\n1d actual reversed N/S")
	// searchOneDimensionRandom(actualReversed, true)

	// fmt.Println("\ntest diagonals")
	// searchSingleRandomDiag(testDiagonals)
	for MOD > 8 {
		MOD /= 2
		fmt.Printf("---------------------------------------\n")
		fmt.Printf("mod=%d\n", MOD)
		fmt.Printf("---------------------------------------\n")
		fmt.Println("\nactual diagonals")
		searchSingleRandomDiag(diagonals)
		fmt.Println("\nreversed actual diagonals")
		searchSingleRandomDiag(reverseString(diagonals))
	}
	return nil
}

// var diagonals = [][]string{
// 	{"/nsd", "\\nsd", "\\nsu", "/n.u", "/nsu", "\\wsu", "\\ns", "", "/wsu", "\\.sd",
// 		"", "\\nsu", "/n.", "/.su", "\\ntu", "\\ns", "\\nsu", "\\n.d", "/ns", ""},
// 	{"", "", "", "\\n.", "/nt", "/wsd", "/ns", "", "", "/ns",
// 		"/nsu", "", "/ntd", "/ns", "/nsu", "", "/wsu", "/nsu", "", ""},
// 	{"/w.u", "/n.", "\\ns", "", "/..", "/nsd", "\\ws", "/.sd", "\\n.d", "/.s",
// 		"\\nsd", "", "/.s", "/.td", "/..d", "\\ns", "/n.d", "\\.s", "/.s", ""},
// 	{"/nsu", "", "/n.u", "/ns", "/.su", "/n.", "/n.", "\\n.", "\\.sd", "/nsu",
// 		"/..", "/ntd", "\\nsd", "/n.", "", "", "/..d", "...u", "/.s", "\\ns"},
// 	{"/n.d", "", "/.su", "/ns", "/nsd", "", "/ns", "", "/n.u", "/n.u",
// 		"\\..u", "\\nsd", "\\..d", "\\n.", "/n.", "", "", "", "", ""},
// 	{"/n.d", "/w.u", "", "", "", "", "/n.u", "/.s", "/ns", "\\n.",
// 		"\\nsd", "/wtd", "\\n.u", "/.s", "\\.s", "", "", "", "/nt", ""},
// 	{"/n.d", "\\nsd", "", "", "", "", "/n.", "", "", "\\.s",
// 		"\\.s", "/nsu", "\\.td", "\\n.d", "\\.s", "/..", "/n.u", "", "\\ns", ""},
// 	{"/n.d", "\\ns", "/nsd", "", "/..", "", "\\ntd", "/wtu", "", "",
// 		"\\n.", "", "", "/nt", "\\nsd", "\\.tu", "/n.", "/n.d", "\\n.d", "/nsu"},
// 	{"/.su", "", "/nsu", "", "", "\\n.", "\\.s", "", "\\nsu", "\\nsu",
// 		"/wtu", "\\.s", "\\w.", "\\ns", "/..u", "", "", "\\.s", "/n.u", "/nsd"},
// 	{"/.sd", "", "", "/nsd", "\\ns", "\\ns", "\\..", "", "/ns", "/..d",
// 		"/wt", "/ns", "", "", "/.tu", "", "\\.sd", "", "\\n.u", ""},
// 	{"\\.sd", "/wt", "/..d", "/..u", "/w.", "/n.", "\\n.", "", "", "...d",
// 		"\\.sd", "", "\\..u", "\\..d", "", "/n.", "\\n.d", "", "/n.", "\\n.u"},
// 	{"/..d", "/.s", "/n.d", "/w.d", "", "\\ns", "/.tu", "/n.u", "\\.s", "/n.u",
// 		"", "\\n.d", "/nsd", "/.sd", "/nsu", "", "/..", "/n.", "\\.s", "\\..u"},
// 	{"/n.", "/.su", "", "", "\\ns", "\\nsd", "", "", "/wtd", "/n..",
// 		"\\ns", "/.su", "", "\\.s", "/.sd", "", "/n.", "/..u", "", "/n."},
// 	{"", "/n.d", "/n.", "\\.su", "/..", "", "/.s", "", "\\nsd", "/ns.",
// 		"\\n.u", "/nsu", "/..", "\\..", "/w.u", "/..d", "\\..u", "\\nsd", "\\n.u", "/n.u"},
// }

var testDiagonals string = strings.Replace(`
/\.///\/\\\/\/
/\\\/\.//\///\
/\\/////////\\
\\//\////\\///
//\/\/\//\//\\
//\\///.\\//\/
\//\\\/////\/\
//\//////\\/\\
/////\/\\\\/\/
/\//////////\\
\//\/\//.///\/
/\\/\////./\//
/////\\/\//\//
///\/\\\//////
//\/////\\\\//
\\\\///\//////
/\././///./\//
\////./\\\/\//
//\/\\.///////
//\\//\\//\///
`, "\n", "", -1)

var diagonals string = strings.Replace(`
/.////////\//.
\./../\\..////
\.\//..//.//./
/\.//....///.\
/////../.\/.\/
\///....\\/\\.
\/\////\\\\/./
../\././.../..
/.\\//..\/.\/\
\////\\.\/.///
./\/\\\\//\.\\
\../\//.\/.\//
///\\\\.\.\/./
////\/\/\.\/\\
\//./\\\//.///
\.\.../\../../
\///..//.\\//\
\/\..../\..//\
/.//./\\/\/\.\
...\...//.\\//
`, "\n", "", -1)

func reverseString(s string) string {
	res := ""
	for _, c := range s {
		res = string(c) + res
	}
	return res
}

func searchSingleRandomDiag(diagonals string) {
	fmt.Println(len(diagonals), diagonals)
	n := 11

	// Factor to jump past one polygon. Double things, since we're using
	// only one random generator.
	singlePolygonFactor := computeFactor(5, 2*(n+1))

	// Factor to jump past most of one polygon. As before, except
	// subtract four, because we want the first and last two
	// evaluations.
	mostOfSinglePolygonFactor := computeFactor(5, 2*(n+1)-4)

	var winning []winner
	ITERS := (MOD + 5000000) / 10000000

	maxCount := 0

OUTER:
	// for seed := uint64((MOD / 2) | 1); seed <= MOD; seed += 2 {
	// for seed := uint64(JS2); seed <= JS2; seed += 2 {

	for seed := uint64(1); seed < MOD; seed += 2 {
		// fmt.Printf("starting seed: %d\n", seed)
		errs := 0
		if (seed+1)%10_000_000 == 0 {
			fmt.Printf(" %d/%d (%d) (max=%d)\n", seed/10000000, ITERS, len(winning), maxCount)
		}
		r := newRandom(seed)

		index := -1
		for _ = range 20 {
			// fmt.Printf("%d\n", uint64(*r))
			// fmt.Println()
			for _ = range 14 {
				index++
				diag := diagonals[index]
				// fmt.Printf("%c", diag)
				if diag == '.' {
					r.raw(singlePolygonFactor)
					continue
				}
				if diag != '/' && diag != '\\' {
					panic(fmt.Sprintf("Weird diagonal character '%c'", diag))
				}

				same := diag == '/'

				x0 := r.next(5, 0, 10)
				y0 := r.next(5, 0, 10)
				r.raw(mostOfSinglePolygonFactor)
				xn := r.next(5, 0, 10)
				yn := r.next(5, 0, 10)

				s1 := x0-xn > 0
				s2 := y0-yn > 0

				if (s1 == s2) != same {
					errs++
					if errs > 5 {
						if index > maxCount {
							maxCount = index
						}
						continue OUTER
					}
				}
			}
		}

		fmt.Printf("Winning seed: %d   (errs=%d)\n", seed, errs)
		winning = append(winning, winner{
			seed:  seed,
			extra: MOD,
		})
	}
	fmt.Println()

	if len(winning) > 0 {
		fmt.Println("All winning seeds:")
		for _, winner := range winning {
			fmt.Printf("mod: %d  seed: %d\n", winner.extra, winner.seed)
		}
	} else {
		fmt.Printf("No winning seeds for MOD=%d\n", MOD)
	}
}
