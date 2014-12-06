// IAL implements

package mmark

// One or more of these can be attached to block elements

type IAL struct {
	id    string            // #id
	class []string          // 0 or more .class
	attr  map[string]string // key=value pairs
}

// Parsing and thus detecting an IAL. Return a valid *IAL or nil.
// IAL can have #id, .class or key=value element seperated by spaces, that
// may be escaped
func (p *parser) isIAL(data []byte) int {
	esc := false
	quote := false
	ialB := 0
	ial := &IAL{}
	for i := 0; i < len(data); i++ {
		switch data[i] {
		case ' ':
			if quote {
				continue
			}
			chunk := data[ialB:i]
			println("IAL chunk seen", string(chunk))
			ialB = i
		case '"':
			if esc {
				esc = !esc
				continue
			}
			quote = !quote
		case '\\':
			esc = !esc
		case '}':
			// if this is mainmatter, frontmatter, or backmatter it isn't an IAL.
			s := string(data[1:i])
			switch s {
			case "frontmatter":
				fallthrough
			case "mainmatter":
				fallthrough
			case "backmatter":
				return 0
			}
			chunk := data[ialB:i]
			println("IAL 2 chunk seen", string(chunk))
			switch {
			case chunk[0] == '.':
				ial.id = string(chunk[1:])
			case chunk[0] == '#':
				ial.class = append(ial.class, string(chunk[1:]))
			default:
				// key=value
			}
			p.ial = append(p.ial, ial)
			return i + 1
		default:
			esc = false
		}
	}
	return 0
}

// renderIAL renders an IAL and returns a string that can be included in the tag:
// class="class" anchor="id" key="value"
func renderIAL(i []*IAL) string {
	if i == nil {
		return ""
	}
	s := ""
	for _, i1 := range i {
		s += " " + i1.id
	}
	return s
}
