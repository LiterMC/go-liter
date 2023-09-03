
package main

import (
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type Table struct {
	Heads []string
	Rows [][]string
}

func ParseTableFromSelection(s *goquery.Selection)(tb *Table){
	if !s.Is("table") {
		return nil
	}
	var heads *goquery.Selection
	if head := s.ChildrenMatcher(goquery.Single("thead")); head.Length() != 0 {
		if head.Children().Length() != 0 {
			heads = head.Children().First().Children()
		}
	}
	body := s.ChildrenMatcher(goquery.Single("tbody"))
	node := body.Children().First()
	if heads == nil {
		if node.Children().Is("th") {
			heads = node.Children()
			node = node.Next()
		}
	}
	if heads == nil {
		return nil
	}
	tb = &Table{
		Heads: make([]string, 0, heads.Length()),
	}
	heads.Each(func(_ int, s *goquery.Selection){
		t := strings.TrimSpace(s.Text())
		n, _ := strconv.Atoi(s.AttrOr("colspan", "1"))
		if n == 0 {
			n = 1
		}
		for i := 0; i < n; i++ {
			tb.Heads = append(tb.Heads, t)
		}
	})
	cols := len(tb.Heads)

	type tdT struct {
		text string
		row int
	}
	lastRow := make([]tdT, cols)
	tb.Rows = make([][]string, 0, body.Length())
	for ; node.Is("tr") && node.Length() != 0; node = node.Next() {
		row := make([]string, cols)
		D: for i, child := 0, node.Children().First(); i < cols && child.Length() != 0; child = child.Next() {
			for {
				if i >= cols {
					break D
				}
				td := lastRow[i]
				if td.row <= 0 {
					break
				}
				lastRow[i].row--
				row[i] = td.text
				i++
			}
			text := strings.TrimSpace(child.Text())
			rspan, _ := strconv.Atoi(child.AttrOr("rowspan", "1"))
			cspan, _ := strconv.Atoi(child.AttrOr("colspan", "1"))
			if cspan == 0 {
				cspan = 1
			}
			for j := i + cspan; i < j && i < cols; i++ {
				row[i] = text
				if rspan > 1 {
					lastRow[i].row = rspan - 1
					lastRow[i].text = text
				}
			}
		}
		tb.Rows = append(tb.Rows, row)
	}
	return
}

func (tb *Table)String()(string){
	var builder strings.Builder
	cols := len(tb.Heads)
	lengths := make([]int, cols)
	for _, row := range tb.Rows {
		for i, d := range row {
			if lengths[i] < len(d) {
				lengths[i] = len(d)
			}
		}
	}
	for i, d := range tb.Heads {
		if lengths[i] < len(d) {
			lengths[i] = len(d)
		}
		if i == 0 {
			builder.WriteString("|")
		}
		builder.WriteString(" ")
		builder.WriteString(d)
		for n := lengths[i]; n > len(d); n-- {
			builder.WriteByte(' ')
		}
		builder.WriteString(" |")
	}
	builder.WriteByte('\n')
	for i := 0; i < cols; i++ {
		if i == 0 {
			builder.WriteByte('|')
		}
		for n := lengths[i] + 1; n >= 0; n-- {
			builder.WriteByte('-')
		}
		builder.WriteByte('|')
	}
	builder.WriteByte('\n')
	for _, row := range tb.Rows {
		for i, d := range row {
			if i == 0 {
				builder.WriteString("|")
			}
			builder.WriteString(" ")
			builder.WriteString(d)
			for n := lengths[i]; n > len(d); n-- {
				builder.WriteByte(' ')
			}
			builder.WriteString(" |")
		}
		builder.WriteByte('\n')
	}
	return builder.String()
}

func (tb *Table)HasHeads(heads []string)(bool){
	if len(tb.Heads) != len(heads) {
		return false
	}
	for i, h := range heads {
		if strings.ToLower(h) != strings.ToLower(tb.Heads[i]) {
			return false
		}
	}
	return true
}

func tableEqualIgnoreNote(a, b *Table)(bool){
	if len(a.Heads) != len(b.Heads) || len(a.Rows) != len(b.Rows) {
		return false
	}
	if !a.HasHeads(b.Heads) {
		return false
	}
	for i, row := range a.Rows {
		row = row[:len(row) - 1] // ignore notes
		row2 := b.Rows[i][:len(row)]
		for j, v := range row {
			if strings.ToLower(v) != strings.ToLower(row2[j]) {
				return false
			}
		}
	}
	return true
}
