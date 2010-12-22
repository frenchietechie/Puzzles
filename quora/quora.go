package main

import (
    "fmt"
)

var width,height int
var zerosCount , solutionsCount int
var datacenter [][]*Point
var start *Point

type Point struct {
	visited bool
	i , j int
	reachCount int
	value int
	neighbours []*Point
}

func NewPoint(i int , j int , value int) *Point {
    return &Point{i: i, j: j , value : value}
}

func (p *Point) Initialize() {
		p.visited = false
		if p.value == 2 {
			start = p
		}
		if p.value == 0 {
			zerosCount++
		}
		
		p.neighbours = make([]*Point, 0 , 4)

		if p.i + 1 < width && datacenter[p.j][p.i + 1].value != 1 {
			p.neighbours = append(p.neighbours , datacenter[p.j][p.i + 1])
			p.reachCount++
		}
		if p.i - 1 >= 0 && datacenter[p.j][p.i - 1].value != 1 {
			p.neighbours = append(p.neighbours , datacenter[p.j][p.i - 1])
			p.reachCount++
		}
		if p.j + 1 < height && datacenter[p.j + 1][p.i].value != 1 {
				p.neighbours = append(p.neighbours , datacenter[p.j + 1][p.i])
			p.reachCount++
		}
		if p.j - 1 >= 0 && datacenter[p.j - 1][p.i].value != 1 {
				p.neighbours = append(p.neighbours , datacenter[p.j - 1][p.i])
			p.reachCount++
		}
}

func (p *Point) Visit(visitedCount int) {
		if p.value == 3 {
			if visitedCount == zerosCount {
				solutionsCount++
			}
			return
		} else { 
	p.visited = true
	for i:= 0 ; i < len(p.neighbours) ; i++ {
				p.neighbours[i].reachCount--
	}
	hadTo := false
	
	for i:= 0 ; i < len(p.neighbours) ; i++ {
			if (!p.neighbours[i].visited) && (p.neighbours[i].reachCount <= 1) && (p.neighbours[i].value != 3) {
				hadTo = true
				p.neighbours[i].Visit(visitedCount + 1)
				break
			}
	}
	if !hadTo {
			for i:= 0 ; i < len(p.neighbours) ; i++ {
				if !p.neighbours[i].visited {
					p.neighbours[i].Visit(visitedCount + 1)
				}
			}
	}
	//remake state as it was
	p.visited = false
	for i:= 0 ; i < len(p.neighbours) ; i++ {
		p.neighbours[i].reachCount++
	}
	}
}


func ParseInput() {
	fmt.Scanf("%d", &width)
	fmt.Scanf("%d", &height)
	
	datacenter = make([][]*Point, height, height)
	
	for j := 0 ; j < height ; j++ {
   	datacenter[j] = make([]*Point,width,width)
   	for i := 0 ; i < width ; i++ {
   			var data int
   			fmt.Scanf("%d", &data)
   			datacenter[j][i] = NewPoint(i,j,data)
   	}
	}
	for j := 0 ; j < height ; j++ {
   		for i := 0 ; i < width ; i++ {
   			datacenter[j][i].Initialize()
   		}
	}
}

func OutputDatacenter() {
		for j := 0 ; j < height ; j++ {
   		for i := 0 ; i < width ; i++ {
   			fmt.Printf("Point of value : %d\n" , datacenter[j][i].value)
   		}
	}
}


func main() {   
	ParseInput()
	start.Visit(-1)
	fmt.Printf("%d\n",solutionsCount)
}


