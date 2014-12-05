//
// Functions to parse quote-like elements.
//

package mmark

import "bytes"

// returns tablequote prefix length
func (p *parser) tablePrefix(data []byte) int {
	i := 0
	for i < 3 && data[i] == ' ' {
		i++
	}
	if data[i] == 'T' && data[i+1] == '>' {
		if data[i+2] == ' ' {
			return i + 3
		}
		return i + 2
	}
	return 0
}

// parse an figure table
func (p *parser) tables(out *bytes.Buffer, data []byte) int {
	var raw bytes.Buffer
	beg, end := 0, 0
	for beg < len(data) {
		end = beg
		for data[end] != '\n' {
			end++
		}
		end++

		if pre := p.tablePrefix(data[beg:]); pre > 0 {
			// skip the prefix
			beg += pre
		} else if p.isEmpty(data[beg:]) > 0 &&
			(end >= len(data) ||
				(p.tablePrefix(data[end:]) == 0 && p.isEmpty(data[end:]) == 0)) {
			break
		}
		raw.Write(data[beg:end])
		beg = end
	}

	var cooked bytes.Buffer
	p.block(&cooked, raw.Bytes())
	p.r.Tables(out, cooked.Bytes())
	return end
	return 0
}

// returns figurequote prefix length
func (p *parser) figurePrefix(data []byte) int {
	i := 0
	for i < 3 && data[i] == ' ' {
		i++
	}
	if data[i] == 'F' && data[i+1] == '>' {
		if data[i+2] == ' ' {
			return i + 3
		}
		return i + 2
	}
	return 0
}

// parse an figure fragment
func (p *parser) figure(out *bytes.Buffer, data []byte) int {
	var raw bytes.Buffer
	beg, end := 0, 0
	for beg < len(data) {
		end = beg
		for data[end] != '\n' {
			end++
		}
		end++

		if pre := p.figurePrefix(data[beg:]); pre > 0 {
			// skip the prefix
			beg += pre
		} else if p.isEmpty(data[beg:]) > 0 &&
			(end >= len(data) ||
				(p.figurePrefix(data[end:]) == 0 && p.isEmpty(data[end:]) == 0)) {
			// figure ends with at least one blank line
			// followed by something without a figure prefix
			break
		}
		raw.Write(data[beg:end])
		beg = end
	}

	var cooked bytes.Buffer
	p.block(&cooked, raw.Bytes())
	p.r.Figure(out, cooked.Bytes())
	return end
	return 0
}

// returns notequote prefix length
func (p *parser) notePrefix(data []byte) int {
	i := 0
	for i < 3 && data[i] == ' ' {
		i++
	}
	if data[i] == 'N' && data[i+1] == '>' {
		if data[i+2] == ' ' {
			return i + 3
		}
		return i + 2
	}
	return 0
}

// parse an note fragment
func (p *parser) note(out *bytes.Buffer, data []byte) int {
	var raw bytes.Buffer
	beg, end := 0, 0
	for beg < len(data) {
		end = beg
		for data[end] != '\n' {
			end++
		}
		end++

		if pre := p.notePrefix(data[beg:]); pre > 0 {
			// skip the prefix
			beg += pre
		} else if p.isEmpty(data[beg:]) > 0 &&
			(end >= len(data) ||
				(p.notePrefix(data[end:]) == 0 && p.isEmpty(data[end:]) == 0)) {
			break
		}
		raw.Write(data[beg:end])
		beg = end
	}

	var cooked bytes.Buffer
	p.block(&cooked, raw.Bytes())
	p.r.Note(out, cooked.Bytes())
	return end
}

// returns asidequote prefix length
func (p *parser) asidePrefix(data []byte) int {
	i := 0
	for i < 3 && data[i] == ' ' {
		i++
	}
	if data[i] == 'A' && data[i+1] == '>' {
		if data[i+2] == ' ' {
			return i + 3
		}
		return i + 2
	}
	return 0
}

// parse an aside fragment
func (p *parser) aside(out *bytes.Buffer, data []byte) int {
	var raw bytes.Buffer
	beg, end := 0, 0
	for beg < len(data) {
		end = beg
		for data[end] != '\n' {
			end++
		}
		end++

		if pre := p.asidePrefix(data[beg:]); pre > 0 {
			// skip the prefix
			beg += pre
		} else if p.isEmpty(data[beg:]) > 0 &&
			(end >= len(data) ||
				(p.asidePrefix(data[end:]) == 0 && p.isEmpty(data[end:]) == 0)) {
			break
		}
		raw.Write(data[beg:end])
		beg = end
	}

	var cooked bytes.Buffer
	p.block(&cooked, raw.Bytes())
	p.r.Aside(out, cooked.Bytes())
	return end
}

// returns abstractquote prefix length
func (p *parser) abstractPrefix(data []byte) int {
	i := 0
	for i < 3 && data[i] == ' ' {
		i++
	}
	if data[i] == 'A' && data[i+1] == 'B' && data[i+2] == '>' {
		if data[i+3] == ' ' {
			return i + 4
		}
		return i + 3
	}
	return 0
}

// parse an abstract fragment
func (p *parser) abstract(out *bytes.Buffer, data []byte) int {
	var raw bytes.Buffer
	beg, end := 0, 0
	for beg < len(data) {
		end = beg
		for data[end] != '\n' {
			end++
		}
		end++

		if pre := p.abstractPrefix(data[beg:]); pre > 0 {
			// skip the prefix
			beg += pre
		} else if p.isEmpty(data[beg:]) > 0 &&
			(end >= len(data) ||
				(p.abstractPrefix(data[end:]) == 0 && p.isEmpty(data[end:]) == 0)) {
			// abstract ends with at least one blank line
			// followed by something without a abstract prefix
			break
		}

		// this line is part of the abstract
		raw.Write(data[beg:end])
		beg = end
	}

	var cooked bytes.Buffer
	p.block(&cooked, raw.Bytes())
	p.r.Abstract(out, cooked.Bytes())
	return end
}

// returns blockquote prefix length
func (p *parser) quotePrefix(data []byte) int {
	i := 0
	for i < 3 && data[i] == ' ' {
		i++
	}
	if data[i] == '>' {
		if data[i+1] == ' ' {
			return i + 2
		}
		return i + 1
	}
	return 0
}

// parse a blockquote fragment
func (p *parser) quote(out *bytes.Buffer, data []byte) int {
	var raw bytes.Buffer
	beg, end := 0, 0
	for beg < len(data) {
		end = beg
		for data[end] != '\n' {
			end++
		}
		end++

		if pre := p.quotePrefix(data[beg:]); pre > 0 {
			// skip the prefix
			beg += pre
		} else if p.isEmpty(data[beg:]) > 0 &&
			(end >= len(data) ||
				(p.quotePrefix(data[end:]) == 0 && p.isEmpty(data[end:]) == 0)) {
			// blockquote ends with at least one blank line
			// followed by something without a blockquote prefix
			break
		}

		// this line is part of the blockquote
		raw.Write(data[beg:end])
		beg = end
	}

	var cooked bytes.Buffer
	p.block(&cooked, raw.Bytes())
	p.r.SetIAL(p.ial)
	p.ial = nil
	p.r.BlockQuote(out, cooked.Bytes())
	return end
}