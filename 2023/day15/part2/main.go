package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/kkevinchou/adventofcode/utils"
)

func main() {
	startTime := time.Now()
	file := "input"
	generator := utils.RecordGenerator(file, "\r\n")

	var sequences []string
	for {
		record, done := generator()
		if done {
			break
		}

		line := record.Line

		sequences = strings.Split(line, ",")
	}

	boxes := [256]*Box{}

	for i := 0; i < 256; i++ {
		start := &Node{Dummy: true}
		end := &Node{Dummy: true}
		start.Next = end
		end.Previous = start

		boxes[i] = &Box{
			Start:  start,
			End:    end,
			Labels: map[string]*Node{},
		}
	}

	for _, seq := range sequences {
		if strings.Contains(seq, "=") {
			seqSplit := strings.Split(seq, "=")
			label := seqSplit[0]
			focalLength := utils.MustParseNum(seqSplit[1])
			key := hash(label)
			AddToBox(boxes[key], label, focalLength)
		} else {
			label := seq[0 : len(seq)-1]
			key := hash(label)
			RemoveFromBox(boxes[key], label)
		}
	}

	var total int
	for boxIndex, box := range boxes {
		node := box.Start.Next
		index := 1
		for !node.Dummy {
			score := node.Value * index * (boxIndex + 1)
			total += score
			index += 1
			node = node.Next
		}
	}

	fmt.Println(total)
	fmt.Println(time.Since(startTime))
}

type Node struct {
	Previous *Node
	Next     *Node
	Value    int

	Dummy bool
}

type Box struct {
	Labels map[string]*Node
	Start  *Node
	End    *Node
}

func RemoveFromBox(b *Box, label string) {
	if _, ok := b.Labels[label]; !ok {
		return
	}

	node := b.Labels[label]
	node.Previous.Next = node.Next
	node.Next.Previous = node.Previous

	delete(b.Labels, label)
}

func AddToBox(b *Box, label string, focalLength int) {
	if node, ok := b.Labels[label]; ok {
		node.Value = focalLength
		return
	}

	node := &Node{Value: focalLength}
	b.Labels[label] = node

	previous := b.End.Previous
	previous.Next = node
	node.Previous = previous

	node.Next = b.End
	b.End.Previous = node

	// NewNode
	// Start - NodeA - NodeB - NodeC - End
}

func hash(s string) int {
	total := 0
	for _, r := range s {

		total += int(r)
		total *= 17
		total %= 256
	}

	return total
}
