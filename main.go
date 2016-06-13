package main


import "fmt"
import "github.com/golang/glog"


//
//
//
//       y
//       ^
//       |
//       |
//       |
//       |
//       |
//       |         width
//  (x,y)+-----------+-----------+(x+width,y)
//       |           |           |
//       |   leftup  | rightup   |
//       |           |           |
//       |-----------+-----------| height
//       |           |           |
//       | leftbottom|rightbootom|
//       |           |           |
//-------+-----------+-----------+------------------------------> x
//       |(x,y-height)            (x+width,y-height)
//

type Rect struct {
    Xpoint int
    Ypoint int
    Width  int
    Height int
}

type Point struct{
    Xpos int
    Ypos int
}

type Quadtree struct {
    *Rect
    Leaf bool
    Depth int
    ChildQuadtree []*Quadtree
    Total int
    AverangePoint *Point
    Objectlist []*Point
    ParentNode *Quadtree
}


func NewQuadtree(x ,y,w ,h int) *Quadtree{
    return &Quadtree{
        Rect:&Rect{x,y,w,h},
        Leaf : false,
        Depth : 0,
        ChildQuadtree : nil,
        Total : 0,
        AverangePoint : nil,
        Objectlist : nil,
    }
}

func (q *Quadtree)ShowObject()  {
    glog.Infoln(*q)
    for _,value := range q.Objectlist {
        glog.Infoln(*value)
    }
    if q.Leaf==false{
        for _,value := range q.ChildQuadtree {
            value.ShowObject()
        }
    }else{
        glog.Infoln("leaf node")
        return
    }
}

func (q *Quadtree)ShowSubrect()  {

    if q.Leaf == false{
        glog.Infoln(*q)
        for index,value := range q.ChildQuadtree {
            glog.Info(index)
            value.ShowSubrect()
        }
        return
    }else{
        glog.Infoln(*q)
        return
    }
}

func (q *Quadtree)SplitAvg(avgpoint * Point)bool  {
    if avgpoint == nil{
        return false
    }
    if q.Leaf == false{
        return false
    }
    leftup := NewQuadtree(q.Xpoint,q.Ypoint,avgpoint.Xpos - q.Xpoint,q.Ypoint-avgpoint.Ypos)
    righttup := NewQuadtree(avgpoint.Xpos,q.Ypoint,q.Width+q.Xpoint-avgpoint.Xpos,q.Ypoint-avgpoint.Ypos)
    leftbottom := NewQuadtree(q.Xpoint,avgpoint.Ypos,avgpoint.Xpos - q.Xpoint,q.Height-q.Ypoint+avgpoint.Ypos)
    rightbootom := NewQuadtree(avgpoint.Xpos,avgpoint.Ypos,q.Width+q.Xpoint-avgpoint.Xpos,q.Height-q.Ypoint+avgpoint.Ypos)
    if leftup == nil|| righttup==nil|| leftbottom==nil||rightbootom ==nil {
        return false
    }

    leftup.Leaf = true
    leftup.Depth=q.Depth + 1
    leftup.ChildQuadtree= nil
    leftup.ParentNode = q
    righttup.Leaf = true
    righttup.Depth=q.Depth + 1
    righttup.ChildQuadtree= nil
    righttup.ParentNode = q
    leftbottom.Leaf = true
    leftbottom.Depth = q.Depth + 1
    leftbottom.ChildQuadtree= nil
    leftbottom.ParentNode = q
    rightbootom.Leaf =true
    rightbootom.Depth = q.Depth + 1
    rightbootom.ChildQuadtree= nil
    rightbootom.ParentNode = q
    q.Leaf=false
    q.Total = 4
    q.ChildQuadtree = make([]*Quadtree,4)
    q.ChildQuadtree[0] = leftup
    q.ChildQuadtree[1] = righttup
    q.ChildQuadtree[2] = leftbottom
    q.ChildQuadtree[4] = rightbootom
    return true
}

func (q *Quadtree)Split() bool {
    if q.Leaf == false{
        return false
    }
    leftup := NewQuadtree(q.Xpoint,q.Ypoint,q.Width/2,q.Height/2)
    righttup := NewQuadtree(q.Xpoint+q.Width/2,q.Ypoint,q.Width/2,q.Height/2)
    leftbottom := NewQuadtree(q.Xpoint,q.Ypoint+q.Height/2,q.Width/2,q.Height/2)
    rightbootom := NewQuadtree(q.Xpoint+q.Width/2,q.Ypoint+q.Height/2,q.Width/2,q.Height/2)
    if leftup == nil|| righttup==nil|| leftbottom==nil||rightbootom ==nil {
        return false
    }

    leftup.Leaf = true
    leftup.Depth=q.Depth + 1
    leftup.ChildQuadtree= nil
    leftup.ParentNode = q
    righttup.Leaf = true
    righttup.Depth=q.Depth + 1
    righttup.ChildQuadtree= nil
    righttup.ParentNode = q
    leftbottom.Leaf = true
    leftbottom.Depth = q.Depth + 1
    leftbottom.ChildQuadtree= nil
    leftbottom.ParentNode = q
    rightbootom.Leaf =true
    rightbootom.Depth = q.Depth + 1
    rightbootom.ChildQuadtree= nil
    rightbootom.ParentNode = q
    q.Leaf=false
    q.Total = 4
    q.ChildQuadtree = make([]*Quadtree,4)
    q.ChildQuadtree[0] = leftup
    q.ChildQuadtree[1] = righttup
    q.ChildQuadtree[2] = leftbottom
    q.ChildQuadtree[4] = rightbootom
    return true
}

func CheckRectMini(rect Rect)bool  {
    if rect.Height<=2 || rect.Width <= 2{
        return false
    }
    return true
}

func (q * Quadtree)GetRectBelongRect(rect *Rect)(* Quadtree){
    if rect == nil{
        return nil
    }
    if q.Leaf == true{
        if rect.Xpoint >= q.Xpoint && rect.Xpoint+rect.Width < q.Xpoint+q.Width && rect.Ypoint >= q.Ypoint && rect.Ypoint+rect.Height<q.Ypoint+q.Height{
            return q
        }else{
            return nil
        }
    }else{
        if rect.Xpoint >= q.Xpoint && rect.Xpoint+rect.Width < q.Xpoint+q.Width && rect.Ypoint >= q.Ypoint && rect.Ypoint+rect.Height<q.Ypoint+q.Height{
            for _,value := range q.ChildQuadtree {
                if re:=value.GetRectBelongRect(rect);re!=nil{
                    return value
                }else{
                    continue
                }
            }
            return q
        }else{
            return nil
        }
    }
    return nil
}

func (q * Quadtree)GetPointBelongRect(point *Point)(* Quadtree){
    if q.Leaf == true{
        if point.Xpos>=q.Xpoint && point.Xpos<q.Xpoint+q.Width && point.Ypos>=q.Ypoint && point.Ypos<q.Ypoint+q.Height{
            return q
        }else{
            return nil
        }
    }else{
        if point.Xpos>=q.Xpoint && point.Xpos<q.Xpoint+q.Width && point.Ypos>=q.Ypoint && point.Ypos<q.Ypoint+q.Height{
            for _,value := range q.ChildQuadtree {
                if re:=value.GetPointBelongRect(point);re!=nil{
                    return value
                }else{
                    continue
                }
            }
        }else{
            return nil
        }
        return q
    }
    return nil
}



func (q * Quadtree)Retrieve(rect *Rect)[]*Rect{
    return nil
}

func (q * Quadtree)ClearRect(rect *Rect){
    if q.Leaf == true{
        if rect.Xpoint <=q.Xpoint && rect.Ypoint <=q.Ypoint && rect.Width >=q.Width && rect.Height >= q.Height{

        }
    }else{
        for _,value := range q.ChildQuadtree {
            value.ClearRect(rect)
        }
    }

}



func main()  {
    fmt.Println("quad tree implement")
}
