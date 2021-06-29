package main

import (
	"fmt"
	"sort"
)
type Point struct{
				x,y int
			}
func sum(a,b int)int  {
	return a + b
}
func (s Points) Len()int {return 0}
func (s Points) Swap(i,j int){}
func (s Points) Less(i,j)bool{ return false}
	
}
func main() {
	var newx,newy int

	
	s:= make([]Point,0)
	for{
		
		_, err:=fmt.Scanf("%d %d", &newx,&newy)
		if(err != nil){
			break
		}
		_ = 5
		_ = "heloo"
		//x := sum(newx,newy)
		var newpo Point
		newpo.x = newx
		newpo.y = newy
		s = append(s,newpo)
		
	}

	x := s[0]

	_=x

	sort.Slice(
		s, 
		func (i,j int)bool  {
			return s[i].x*s[i].x+s[i].y*s[i].y < s[j].x*s[j].x+s[j].y*s[j].y
		},
	)

fmt.Print(s)
}
func dist (s Point) float64 {
	zz := s[.x

	
	return s.x*s.x+s.y*s.y
}
