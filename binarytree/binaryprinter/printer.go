package binaryprinter

import (
	"fmt"
	"log"
	"math"
	"strings"

	"github.com/dragon1672/go-collections/binarytree"
	"github.com/dragon1672/go-collections/queue"
	"github.com/dragon1672/go-collections/vector"
)

const (
	leftDown  = '/'
	rightDown = '\\'
	sideWays  = '-'
	spacer    = ' '
)

type trimmableChar struct {
	final     rune
	trashable bool
}

type TrimmableBoard[T any] struct {
	width, height int
	data          map[vector.IntVec2]trimmableChar

	NodeSpacing int // defaults to 1
	NoCollapse  bool
}

func (t *TrimmableBoard[T]) String() string {
	sb := strings.Builder{}
	sb.Grow(t.height * (t.width + 1))
	for y := 0; y < t.height; y++ {
		for x := 0; x < t.width; x++ {
			if r, ok := t.data[vector.Of(x, y)]; ok {
				sb.WriteRune(r.final)
			} else {
				sb.WriteRune(spacer)
			}
		}
		if y < t.height-1 {
			sb.WriteRune('\n')
		}
	}
	return sb.String()
}

func (t *TrimmableBoard[T]) WriteBinaryTree(node *binarytree.Node[T]) error {
	h, err := binarytree.Height(node)
	if err != nil {
		return err
	}
	dataLength, err := maxNodeContents(node)
	if err != nil {
		return err
	}
	if dataLength%2 == 0 {
		dataLength++ // odd numbered dataLength centers better
	}
	boardWidth := t.getNodeWidth(h, dataLength)
	t.writeBinaryTree(vector.Of(boardWidth/2, 0), node, h, dataLength)
	if !t.NoCollapse {
		t.trim()
	}
	return nil
}

func (t *TrimmableBoard[T]) writeBinaryTree(pos vector.IntVec2, node *binarytree.Node[T], height int, dataLength int) {
	if node == nil {
		return
	}
	data := fmt.Sprint(node.Data)
	t.writeStringCenterAligned(pos, data)
	if height > 0 {
		lastWidth := t.getNodeWidth(height-1, dataLength)
		numOfDashes := (lastWidth-1)/2 - ((dataLength-1)/2 - 1)
		offset := (dataLength - 1) / 2
		for i := 1; i <= numOfDashes; i++ {
			atTheEnd := i == numOfDashes
			offset++
			if node.Left != nil {
				if atTheEnd {
					t.WriteRune(pos.Add(vector.Of(-offset, 0)), leftDown)
				} else {
					t.WriteTrashableRune(pos.Add(vector.Of(-offset, 0)), sideWays)
				}
			}
			if node.Right != nil {
				if atTheEnd {
					t.WriteRune(pos.Add(vector.Of(offset, 0)), rightDown)
				} else {
					t.WriteTrashableRune(pos.Add(vector.Of(offset, 0)), sideWays)
				}
			}
		}
		t.writeBinaryTree(pos.Add(vector.Of(-offset, 1)), node.Left, height-1, dataLength)
		t.writeBinaryTree(pos.Add(vector.Of(offset, 1)), node.Right, height-1, dataLength)
	}
}

func (t *TrimmableBoard[T]) writeStringCenterAligned(startingPos vector.IntVec2, str string) {
	pos := vector.Of(startingPos.X-len(str)/2, startingPos.Y)
	t.writeString(pos, str)
}

func (t *TrimmableBoard[T]) writeString(startingPos vector.IntVec2, str string) {
	for i, r := range str {
		pos := startingPos.Add(vector.Of(i, 0))
		t.WriteRune(pos, r)
	}
}

func (t *TrimmableBoard[T]) WriteRune(pos vector.IntVec2, r rune) {
	if t.data == nil {
		t.data = make(map[vector.IntVec2]trimmableChar)
	}
	t.data[pos] = trimmableChar{final: r}
	if t.width <= pos.X {
		t.width = pos.X + 1
	}
	if t.height <= pos.Y {
		t.height = pos.Y + 1
	}
}

func (t *TrimmableBoard[T]) WriteTrashableRune(pos vector.IntVec2, r rune) {
	if t.data == nil {
		t.data = make(map[vector.IntVec2]trimmableChar)
	}
	t.data[pos] = trimmableChar{final: r, trashable: true}
	if t.width <= pos.X {
		t.width = pos.X + 1
	}
	if t.height <= pos.Y {
		t.height = pos.Y + 1
	}
}

//returns the required width to print given tree
func (t *TrimmableBoard[T]) getNodeWidth(height, nodePrintLength int) int {
	if t.NodeSpacing < 1 {
		t.NodeSpacing = 1
	}
	return (t.NodeSpacing+nodePrintLength)*int(math.Pow(float64(2), float64(height-1))) - 1 // trust my math
}

func (t *TrimmableBoard[T]) trim() {
	// determine what columns can be removed
	emptyCols := t.getEmptyColumns()
	// determine how far each column needs to be shifted
	curColShiftAmt := 0
	colShifts := make(map[int]vector.IntVec2)

	for x := 0; x < t.width; x++ {
		if _, ok := emptyCols[x]; ok {
			// this column should be trashed
			curColShiftAmt--
		} else {
			colShifts[x] = vector.Of(curColShiftAmt, 0)
		}
	}

	// update values and drop excess entries
	newData := make(map[vector.IntVec2]trimmableChar)
	newWidth := 0
	for k, v := range t.data {
		shift, ok := colShifts[k.X]
		if !ok {
			if !v.trashable {
				log.Fatalf("calculated to delete a non trashable value %v at position %v", v.final, k)
			}
			continue // missing columns are the ones to be deleted
		}
		newPos := k.Add(shift)
		newData[newPos] = v
		if newWidth <= newPos.X {
			newWidth = newPos.X + 1
		}
	}
	t.data = newData
	t.width = newWidth
}

func (t *TrimmableBoard[T]) getEmptyColumns() map[int]bool {
	ret := make(map[int]bool)
	for x := 0; x < t.width; x++ {
		ret[x] = true
	}
	for k, v := range t.data {
		if !v.trashable {
			delete(ret, k.X)
		}
	}
	return ret
}

func maxNodeContents[T any](n *binarytree.Node[T]) (int, error) {
	traverse := queue.MakeNew(n)
	seen := make(map[*binarytree.Node[T]]bool)
	maxLen := 0
	for traverse.Size() > 0 {
		elem := traverse.Pop()
		if elem == nil {
			continue
		}
		if _, exists := seen[elem]; exists {
			return -1, fmt.Errorf("cycle detected, tree is not valid binary tree")
		}
		seen[elem] = true
		l := len(fmt.Sprint(elem.Data))
		if maxLen < l {
			maxLen = l
		}
		traverse.Push(elem.Right)
		traverse.Push(elem.Left)
	}
	return maxLen, nil
}
