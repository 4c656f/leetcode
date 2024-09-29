package leetcode
type Node struct {
    prev *Node
    next *Node
    str map[string]int
    stringCount int
    val int
}



type AllOne struct {
    start *Node
    end *Node
    count map[string]*Node    
}


func Constructor() AllOne {
    return AllOne{
        nil,
        nil,
        map[string]*Node{},
    }
}

func (this *AllOne)Remove(n *Node){
    if n.prev != nil{
        n.prev.next = n.next
    }
    if n.next != nil{
        n.next.prev = n.prev
    }
    if this.start == n{
        this.start = n.next
    }
    if this.end == n{
        this.end = n.prev
    }
    n.next = nil
    n.prev = nil
}

func (this *AllOne)ShiftToRight(n*Node, c int, key string)*Node{
    n.stringCount--
    delete(n.str, key)
    if n.next != nil && n.next.val == c{
        n.next.stringCount++
        n.next.str[key]=1
        return n.next
    }
    newNode := &Node{
        prev: n,
        next: n.next,
        stringCount: 1,
        val: c,
        str: map[string]int{
            key: 1,
        },
    }
    if n.next != nil {
        n.next.prev = newNode
    }
    n.next = newNode
    if n == this.end{
        this.end = newNode
    }
    return newNode
}

func (this *AllOne)ShiftToLeft(n*Node, c int, key string)*Node{
    n.stringCount--
    delete(n.str, key)
    if n.prev !=nil && n.prev.val == c{
        n.prev.stringCount++
        n.prev.str[key]=1
        return n.prev
    }
    newNode := &Node{
        prev: n.prev,
        next: n,
        stringCount: 1,
        val: c,
        str: map[string]int{
            key: 1,
        },
    }
    if n.prev != nil{
        n.prev.next = newNode
    }
    n.prev = newNode
    
    if n == this.start{
        this.start = newNode
    }
    return newNode
}

func (this *AllOne) Inc(key string)  {
    n, ok := this.count[key]

    if !ok{
        if this.start == nil{
            newNode := &Node{
                prev: nil,
                next: nil,
                val: 1,
                stringCount: 1,
                str: map[string]int{key: 1,},
            }
            this.start = newNode
            this.end = newNode
            this.count[key] = newNode
            return
        }
        if this.start.val == 1{
            this.start.stringCount++
            this.start.str[key] = 1
            this.count[key] = this.start
            return
        }
        this.start.stringCount++
        newN := this.ShiftToLeft(this.start, 1, key)
        this.count[key] = newN
        return
    }
    newN := this.ShiftToRight(n, n.val+1, key)
    if n.stringCount == 0{
        this.Remove(n)
    }
    this.count[key] = newN
}


func (this *AllOne) Dec(key string)  {
    n := this.count[key]
    if n.val == 1{
        delete(this.count, key)
        n.stringCount--
        delete(n.str, key)
        if n.stringCount == 0{
            this.Remove(n)
        }
        return
    }
    newN := this.ShiftToLeft(n, n.val -1, key)
    this.count[key] = newN
    if n.stringCount == 0{
            this.Remove(n)
    }
}


func (this *AllOne) GetMaxKey() string {
    if this.end == nil{
        return ""
    }
    for k, _ := range this.end.str{
        return k
    }
    return ""
}


func (this *AllOne) GetMinKey() string {
    if this.start == nil{
        return ""
    }
    for k, _ := range this.start.str{
        return k
    }
    return ""
}