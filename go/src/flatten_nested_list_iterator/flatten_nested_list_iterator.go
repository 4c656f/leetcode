/**
 * // This is the interface that allows for creating nested lists.
 * // You should not implement it, or speculate about its implementation
 * type NestedInteger struct {
 * }
 *
 * // Return true if this NestedInteger holds a single integer, rather than a nested list.
 * func (this NestedInteger) IsInteger() bool {}
 *
 * // Return the single integer that this NestedInteger holds, if it holds a single integer
 * // The result is undefined if this NestedInteger holds a nested list
 * // So before calling this method, you should have a check
 * func (this NestedInteger) GetInteger() int {}
 *
 * // Set this NestedInteger to hold a single integer.
 * func (n *NestedInteger) SetInteger(value int) {}
 *
 * // Set this NestedInteger to hold a nested list and adds a nested integer to it.
 * func (this *NestedInteger) Add(elem NestedInteger) {}
 *
 * // Return the nested list that this NestedInteger holds, if it holds a nested list
 * // The list length is zero if this NestedInteger holds a single integer
 * // You can access NestedInteger's List element directly if you want to modify it
 * func (this NestedInteger) GetList() []*NestedInteger {}
 */
package flattern_nested_list_iterator
type NestedInteger interface{
    IsInteger() bool
    GetInteger() int
    SetInteger(value int)
    Add(elem NestedInteger)
    GetList() []*NestedInteger
}
type NestedIterator struct {
    list []int
    idx int
}
func Dfs(list *[]int, nested []*NestedInteger) {
    for _, n := range nested{
        if (n.IsInteger()){
           *list = append(*list, n.GetInteger())
           continue
        }
        Dfs(list, n.GetList())
    }
}

func Constructor(nestedList []*NestedInteger) *NestedIterator {
    list := []int{}
    Dfs(&list, nestedList)
    return &NestedIterator{list:list, idx: 0}
}

func (this *NestedIterator) Next() int {
    val := this.list[this.idx]
    this.idx+=1;
    return val
}

func (this *NestedIterator) HasNext() bool {

    return this.idx < len(this.list)
}
