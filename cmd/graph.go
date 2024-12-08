package cmd

import (
	"fmt"
	"github.com/muskonu/modconflict/tmpl"
	"github.com/muskonu/modconflict/utils"
	"github.com/spf13/cobra"
	"log"
	"os"
	"text/template"
)

// graphCmd represents the graph command
var graphCmd = &cobra.Command{
	Use:   "graph",
	Short: "Generate images of potentially conflicting dependency routes",
	Long:  `Generate images of potentially conflicting dependency routes.`,
	Run: func(cmd *cobra.Command, args []string) {
		s, err := utils.GetInputScanner(InputFileName)
		if err != nil {
			log.Fatal(err)
		}
		graph := newGraph()
		for s.Scan() {
			err := graph.addRecord(s.Text())
			if err != nil {
				log.Fatal(err)
			}
		}
		path := graph.findGraphConflict()
		parse, err := template.New("graphviz").Parse(tmpl.DotTmpl)
		if err != nil {
			log.Fatal(err)
		}
		f, err := os.CreateTemp("", "modgraph")
		if err != nil {
			log.Fatal(err)
		}
		defer os.Remove(f.Name())

		defer f.Close()

		tmplInfo := graph.GetTmplInfo(path)

		err = parse.Execute(f, tmplInfo)
		if err != nil {
			fmt.Println(err)
		}

		var command string
		if OutputFileName == "" {
			command = fmt.Sprintf("dot %s -T svg -o demo.svg", f.Name())
			utils.System(command)
			return
		}
		hasFormat, format := utils.GetImageFormat(OutputFileName)
		if !hasFormat {
			command = fmt.Sprintf("dot %s -T %s -o %s", f.Name(), format, OutputFileName+"."+format)
		} else {
			command = fmt.Sprintf("dot %s -T %s -o %s", f.Name(), format, OutputFileName)
		}
		utils.System(command)
	},
}

var OutputFileName string

func init() {
	rootCmd.AddCommand(graphCmd)

	graphCmd.Flags().StringVarP(&OutputFileName, "ouput", "o", "", "the name of the output file, with different formats depending on the suffix.")
}

func (g *Graph) findGraphConflict() []string {
	path := []string{}
	visited := map[*Node]bool{}

	var findPath func(n *Node)
	dep := 0
	findPath = func(n *Node) {
		dep++
		defer func() {
			dep--
		}()
		if visited[n] {
			return
		}
		visited[n] = true
		if dep != 1 {
			if len(n.Parents) != 0 {
				findPath(n.Parents[0])
				single := fmt.Sprintf("%#p -> %#p", n.Parents[0], n)
				path = append(path, single)
			}
			return
		}

		for i := 0; i < len(n.Parents); i++ {
			findPath(n.Parents[i])
			single := fmt.Sprintf("%#p -> %#p", n.Parents[i], n)
			path = append(path, single)
		}
	}

	for _, nodes := range g.PackageMap {
		if len(nodes) > 1 {
			for _, n := range nodes {
				findPath(n)
			}
		}
	}
	return path
}

type TmplInfo struct {
	NodeMap map[string]*Node
	Paths   []string
}

func (g *Graph) GetTmplInfo(paths []string) TmplInfo {
	nodeMap := map[string]*Node{}
	for _, nodes := range g.PackageMap {
		if len(nodes) > 1 {
			for _, node := range nodes {
				nodeMap[node.PackageWithVersion] = node
			}
		}
	}
	return TmplInfo{
		NodeMap: nodeMap,
		Paths:   paths,
	}
}
