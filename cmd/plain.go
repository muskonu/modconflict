package cmd

import (
	"fmt"
	"github.com/logrusorgru/aurora/v4"
	"github.com/muskonu/modconflict/utils"
	"github.com/spf13/cobra"
	"log"
)

// plainCmd represents the plain command
var plainCmd = &cobra.Command{
	Use:   "plain",
	Short: "Print out the chain of possible dependency conflicts.",
	Long:  `Print out the chain of possible dependency conflicts.`,
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
		graph.findPlainConflict()
	},
}

func init() {
	rootCmd.AddCommand(plainCmd)
}

func (g *Graph) findPlainConflict() {
	visited := map[*Node]bool{}
	counter := map[*Node][]string{}
	var findPath func(n *Node) []string
	dep := 0
	findPath = func(n *Node) []string {
		dep++
		defer func() {
			dep--
		}()
		//防止循环依赖
		if visited[n] {
			panic("Discover cyclic dependencies")
		}
		visited[n] = true
		defer func() { visited[n] = false }()

		//记忆化
		if arr, ok := counter[n]; ok {
			return arr
		}

		if len(n.Parents) == 0 {
			return []string{n.PackageWithVersion}
		}

		self := []string{}

		if dep != 1 {
			paths := findPath(n.Parents[0])
			for _, path := range paths {
				path = fmt.Sprintf("%s -> %s", path, n.PackageWithVersion)
				self = append(self, path)
			}
			counter[n] = self
			return self
		}

		for i := 0; i < len(n.Parents); i++ {
			paths := findPath(n.Parents[i])
			for _, path := range paths {
				path = fmt.Sprintf("%s -> %s", path, n.PackageWithVersion)
				self = append(self, path)
			}
		}

		counter[n] = self

		return self
	}

	for packageName, nodes := range g.PackageMap {
		if len(nodes) > 1 {
			fmt.Println(aurora.Red(fmt.Sprintf("find Confict in package %s:", packageName)))
			for _, n := range nodes {
				arr := findPath(n)
				for _, s := range arr {
					fmt.Println(s)
				}
			}
		}
	}
	return
}
