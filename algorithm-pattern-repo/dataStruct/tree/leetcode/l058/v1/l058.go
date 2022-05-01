package v1

//leetcode: https://leetcode-cn.com/problems/fi9suh/
//思路：
//二叉搜索树的形式，存储每次book的区间
//插入新区间的时候，搜索查看是否有区间和新区间有重叠
type MyCalendar struct {
	root *bookNode
}

func Constructor() MyCalendar {
	return MyCalendar{root: nil}
}

func (this *MyCalendar) Book(start int, end int) bool {
	res, tmp := this.search(this.root, start, end)
	if this.root == nil {
		this.root = tmp
	}
	return res
}

func (this *MyCalendar) search(root *bookNode, start, end int) (bool, *bookNode) {
	if nil == root {
		//如果没找到重叠的元素，那么说明可以写入
		return true, &bookNode{
			start: start,
			end:   end,
			left:  nil,
			right: nil,
		}
	}
	if root.isOverLap(start, end) { //存在重叠，则不允许记录
		return false, root
	}
	leftOrRight := root.isLeft(start, end)
	switch leftOrRight {
	case 1: //右子树
		res, tmp := this.search(root.right, start, end)
		if res {
			root.right = tmp
		}
		return res, root
	case 2: //左子树
		res, tmp := this.search(root.left, start, end)
		if res {
			root.left = tmp
		}
		return res, root
	}
	return false, root
}

type bookNode struct {
	start int
	end   int
	left  *bookNode
	right *bookNode
}

func (this *bookNode) isOverLap(start, end int) bool {
	if (start > this.start && start < this.end) ||
		(end > this.start && end < this.end) {
		return true
	}
	return false
}

//isLeft 是否要到左子树查找
//
func (this *bookNode) isLeft(start, end int) uint8 {
	if start >= this.end {
		return 1 //右子树
	} else if end <= this.start {
		return 2 //左子树
	}
	return 0
}
