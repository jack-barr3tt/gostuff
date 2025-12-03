package types

type Range struct {
	Start int
	End   int
}

type WRange struct {
	Start int
	Width int
}

func (r Range) Contains(num int) bool {
	return num >= r.Start && num <= r.End
}

func (r Range) ContainsRange(other Range) bool {
	return r.Contains(other.Start) && r.Contains(other.End)
}

func (a Range) SplitAfter(num int) (Range, Range) {
	if !a.Contains(num) {
		panic("Range does not contain number")
	}

	return Range{Start: a.Start, End: num}, Range{Start: num + 1, End: a.End}
}

func (a Range) SubtractRange(b Range) []Range {
	// Case 1: a and b are equal - we return empty
	if a.Start == b.Start && a.End == b.End {
		return []Range{}
	}

	// Case 2: b is contained in a and is smaller - we return two ranges, one on either side of b
	if b.Start > a.Start && b.End < a.End {
		c, _ := a.SplitAfter(b.Start - 1)
		_, d := a.SplitAfter(b.End)
		return []Range{c, d}
	}

	// Case 3: a is contained by b - we return empty
	if b.ContainsRange(a) {
		return []Range{}
	}

	// Case 4: a and b are disjoint - we return a
	if a.End < b.Start || a.Start > b.End {
		return []Range{a}
	}

	// Case 5: a overlaps b on the left - we return a but up to but not including b.Start
	if a.Start < b.Start && a.End >= b.Start && a.End <= b.End {
		c, _ := a.SplitAfter(b.Start - 1)
		return []Range{c}
	}

	// Case 6: a overlaps b on the right - we return a but starting after b.End
	if a.Start >= b.Start && a.Start <= b.End && a.End > b.End {
		_, d := a.SplitAfter(b.End)
		return []Range{d}
	}

	// 🤯
	panic("Unhandled case in Range.SubtractRange")
}

func (a Range) AddRange(b Range) *Range {
	if a.End < b.Start-1 || b.End < a.Start-1 {
		return nil
	}

	start := a.Start
	if b.Start < start {
		start = b.Start
	}

	end := a.End
	if b.End > end {
		end = b.End
	}

	result := Range{Start: start, End: end}
	return &result
}
