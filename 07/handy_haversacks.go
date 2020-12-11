package a07

import (
	"errors"
	"regexp"
	"strconv"
	"strings"
)

type Bag struct {
	name string
	contains map[string]int
}

var BagCache = make(map[string]*Bag)

func ParseBag(s string) *Bag {
	bag := &Bag{
		name: "",
		contains: make(map[string]int),
	}	
	bagArr := strings.Split(s, "bags contain")
	name := CleanBag(bagArr[0])
	contains := strings.Split(bagArr[1], ",")
	bag.name = name

	for _, child := range contains {
		name, count, err := ParseContainingBag(child)

		if err == nil {
			bag.contains[CleanBag(name)] = count
		}
	}

	BagCache[bag.name] = bag

	return bag
}

func CleanBag(s string) string {
	p := regexp.MustCompile(`\d| |\,|\.|bags|bag`)

	return p.ReplaceAllString(s, "")
}

func ParseContainingBag(s string) (name string, count int, errMessage error) {
	if (strings.Contains(s, "no other bags")) {
		errMessage = errors.New("no other bags")
		
		return 
	}

	d := regexp.MustCompile(`\d`)
	matches := d.FindAllString(s, 1)
	name = CleanBag(s)
	count, err := strconv.Atoi(matches[0])
	
	if err != nil {
		panic(err)
	}

	return
}

type Node struct {
	name string
	count int
	parnet *Node
	children []*Node
}

var	Endings = []*Node{}

func Reverse(name string, parent *Node) *Node {
	node := &Node{
		name: name,
		children: make([]*Node, 0),
		parnet: parent,
	}

	ending := true

	for _, bag := range BagCache {
		if _, exist := bag.contains[name]; exist {
			ending = false
			node.children = append(node.children, Reverse(bag.name, node))
		}
	}

	if (ending) {
		Endings = append(Endings, node)
	}

	return node
}

func CollectPath(n *Node) []string {
	current := n
	trace := make([]string, 0)

	for {
		if current == nil {
			break
		}

		trace = append(trace, current.name)
		current = current.parnet	
	}

	return trace
}

func Filter(s string, arr []string) []string {
	list := make([]string, 0)

	for _, v := range arr {
		if strings.Compare(s, v) == 0 {
			continue
		}

		list = append(list, v)
	}

	return list
}

func Unique(arr []string) [] string {
	list := make([]string, 0)
	found := make(map[string]bool)

	for _, s := range arr {
		if found[s] {
			continue
		}

		found[s] = true
		list = append(list, s)
	}

	return list
}

func Concat(a1 []string, a2 []string) []string {
	for _, item := range a2 {
		a1 = append(a1, item)	
	}

	return a1
}

func BagValue(name string) int {
	total := 0
	bag := BagCache[name]

	for childName, value := range bag.contains {
		total += value + (value * BagValue(childName))
	}

	return total
}