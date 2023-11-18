package piscine

import "github.com/01-edu/z01"

func QuadA(x,y int){
    if x>0 &&y>0{
        //top line
        z01.PrintRune('o')
        for i:= 0; i < x-2; i++{
            z01.PrintRune('-')
        }
        if x > 1{
            z01.PrintRune('o')
        }
        z01.PrintRune('\n')
    //mid
        for i:=0;i<y-2;i++{
            z01.PrintRune('|')
            for j:=0;j<x-2;j++{
                z01.PrintRune(' ')
            }
            if x>1{
                z01.PrintRune('|')
            }
            z01.PrintRune('\n')
        }   
        if y>1{
            z01.PrintRune('o')
            for i:=0;i<x-2;i++{
                z01.PrintRune('-')
            }
            if x>1{
                z01.PrintRune('o')
            }
        }
        z01.PrintRune('\n')
    }
}
//end

