package cmd

import (
	"errors"
	"github.com/muskonu/modconflict/utils"
	"github.com/spf13/cobra"
	"os"
	"strings"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "modconflict",
	Short: "Tool to quickly locate paths with conflicting dependencies.",
	Long:  `Tool to quickly locate paths with conflicting dependencies.`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

var InputFileName string

func init() {
	rootCmd.PersistentFlags().StringVarP(&InputFileName, "file", "f", "", "select the file where the results of the go mod graph are saved as input")
}

type Graph struct {
	PackageMap map[string][]*Node
	NodeMap    map[string]*Node
}

type Node struct {
	PackageWithVersion string
	Parents            []*Node
}

func newGraph() Graph {
	return Graph{
		PackageMap: map[string][]*Node{},
		NodeMap:    map[string]*Node{},
	}
}

func (g *Graph) addNode(s string) {
	n := &Node{
		PackageWithVersion: s,
		Parents:            nil,
	}
	g.NodeMap[s] = n
	pkg, _ := utils.SplitPkgVersion(s)
	pkgNodes := g.PackageMap[pkg]
	pkgNodes = append(pkgNodes, n)
	g.PackageMap[pkg] = pkgNodes
}

func (g *Graph) addRecord(s string) error {
	splits := strings.Split(s, " ")
	if len(splits) != 2 {
		return errors.New("input not valid")
	}
	p, c := splits[0], splits[1]
	if _, ok := g.NodeMap[p]; !ok {
		g.addNode(p)
	}
	if _, ok := g.NodeMap[c]; !ok {
		g.addNode(c)
	}
	cnode := g.NodeMap[c]
	pnode := g.NodeMap[p]
	cnode.Parents = append(cnode.Parents, pnode)
	return nil
}
