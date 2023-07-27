package main

type set struct {
	setcount int
	setsize  int
	data     [][]int
}

func NewSet(size int) set {
	st := set{setcount: 0, setsize: size, data: make([][]int, size)}

	for i := 0; i < size; i++ {
		st.data[i] = make([]int, 0)
	}

	return st
}

func (st *set) add(x int) {
	if st.find(x) == true {
		return
	}
	st.data[x%st.setsize] = append(st.data[x%st.setsize], x)
	st.setcount++
}

func (st *set) find(x int) bool {
	for _, elem := range st.data[x%st.setsize] {
		if elem == x {
			return true
		}
	}
	return false
}

func (st *set) delete(x int) {
	ind := -1

	for pos, elem := range st.data[x%st.setsize] {
		if elem == x {
			ind = pos
		}
	}

	if ind == -1 {
		return
	}

	st.data[x%st.setsize][ind] = st.data[x%st.setsize][len(st.data[x%st.setsize])-1]
	st.data[x%st.setsize] = st.data[x%st.setsize][:len(st.data[x%st.setsize])-1]
	st.setcount--
}
