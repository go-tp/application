package extend

// 分页
func Page(page,size int) (start,end int){
	// page 页数 size 条数
	start = (page-1)*size
	end = size
	return start,end
}