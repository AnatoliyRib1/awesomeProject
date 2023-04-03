package designOrdered

type OrderedStream struct {
	stream1 []string
	idx     int
}

func Constructor(n int) OrderedStream {
	return OrderedStream{
		stream1: make([]string, n)}
}

func (this *OrderedStream) Insert(idKey int, value string) []string {

	this.stream1[idKey-1] = value
	if idKey-1 == this.idx {
		start := this.idx
		for this.idx < len(this.stream1) && this.stream1[this.idx] != "" {
			this.idx++
		}
		return this.stream1[start:this.idx]
	}
	return []string{}
}
