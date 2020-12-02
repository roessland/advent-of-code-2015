package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"math"
	"os"
	"strings"
)

const wall rune = '#'
const empty rune = '.'

func ReadInput(filename string) [][]rune {
	f, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	reader := bufio.NewReader(f)
	out := make([][]rune, 0)
	line := make([]rune, 0)
	for {
		r, _, err := reader.ReadRune()
		if r == '\r' {
			continue
		}
		if err == io.EOF {
			if len(line) > 0 {
				out = append(out, line)
			}
			break
		} else if err != nil {
			log.Fatal(err)
		} else if r == '\n' {
			out = append(out, line)
			line = nil
		} else {
			line = append(line, r)
		}
	}
	return out
}

func PrintMap(tiles [][]rune) {
	for _, line := range tiles {
		fmt.Printf("%s\n", string(line))
	}
}

type Pos struct {
	I, J int
}

type WorldEdge struct {
	To *WorldNode
	Weight int
}

type WorldNode struct {
	Edges []*WorldEdge
	Pos   Pos
	Tile  rune
}

func (n *WorldNode) DotNodeName() string {
	return fmt.Sprintf(`%c\n(%d,%d)`, n.Tile, n.Pos.I, n.Pos.J)
}

func CreateWorld(input [][]rune) map[Pos]*WorldNode {
	nodeAtPos := map[Pos]*WorldNode{}
	height := len(input)
	width := len(input[0])
	for i := 1; i < height-1; i++ {
		for j := 1; j < width-1; j++ {
			pos := Pos{I: i, J: j}
			// make da node
			n := WorldNode{
				Edges: make([]*WorldEdge, 0),
				Pos: pos,
				Tile: input[i][j],
			}
			if n.Tile != wall {
				nodeAtPos[pos] = &n
			}
		}
	}
	for i := 1; i < height-1; i++ {
		for j := 1; j < width-1; j++ {
			this := nodeAtPos[Pos{I: i, J: j}]
			if this == nil {
				continue
			}
			east := nodeAtPos[Pos{I: i, J: j+1}]
			west := nodeAtPos[Pos{I: i, J: j-1}]
			south := nodeAtPos[Pos{I: i+1, J: j}]
			north := nodeAtPos[Pos{I: i-1, J: j}]
			for _, dir := range []*WorldNode{east, west, south, north} {
				if dir == nil {
					continue
				}
				if dir.Tile != wall {
					this.Edges = append(this.Edges, &WorldEdge{dir, 1})
				}
			}
		}
	}
	return nodeAtPos
}

func WorldToDot(world map[Pos]*WorldNode, filename string) {
	var lines []string
	lines = append(lines, "digraph G {")
	for _, node := range world {
		lines = append(lines, fmt.Sprintf(`"%s" [`, node.DotNodeName()))
		lines = append(lines, fmt.Sprintf(`    pos = "%d,%d!"`, node.Pos.I, node.Pos.J))
		lines = append(lines, "]")
		for _, edge := range node.Edges {
			line := fmt.Sprintf(`"%s" -> "%s"`, node.DotNodeName(), edge.To.DotNodeName())
			lines = append(lines, line)
		}
	}

	lines = append(lines, "}")
	err := ioutil.WriteFile(filename, []byte(strings.Join(lines, "\n")), 0644)
	if err != nil {
		log.Fatal(err)
	}
}
func FindReachableKeysAndDoors(world map[Pos]*WorldNode, from *WorldNode) []*WorldEdge {
	var visited map[*WorldNode]bool
	distanceTo := map[*WorldNode]int{}
	for _, node := range world {
		distanceTo[node] = math.MaxInt32
	}
	var dfs func(n *WorldNode)
	dfs = func(n *WorldNode) {
		if visited[n] {
			return
		}
		visited[n] = true
		for _, edge := range n.Edges {
			dfs(edge.To)
		}
	}
	dfs(from)
}

func main() {
	input := ReadInput("input.txt")
	PrintMap(input)
	w := CreateWorld(input)
	fmt.Println("Simpifying")

	WorldToDot(w, "world.dot")
}