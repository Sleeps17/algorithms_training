package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
)

type Segment struct {
	left  int
	right int
}

func mainC() {
	rd := bufio.NewReader(os.Stdin)

	var M int
	fmt.Fscan(rd, &M)

	segments := make([]Segment, 0, 8)

	for {
		var l, r int
		fmt.Fscan(rd, &l, &r)

		if l == r && l == 0 {
			break
		}

		if r < 0 {
			continue
		}

		segments = append(segments, Segment{l, r})
	}

	sort.Slice(segments, func(i, j int) bool {
		if segments[i].left == segments[j].left {
			return segments[i].right < segments[j].right
		}

		return segments[i].left < segments[j].left
	})

	nowRight := 0
	maxRight := math.MinInt
	var curSeg Segment

	ans := make([]Segment, 0, len(segments))

	for _, seg := range segments {

		if seg.left > nowRight {
			nowRight = maxRight
			ans = append(ans, curSeg)
			if nowRight >= M {
				break
			}
		}

		if seg.right > maxRight && seg.left <= nowRight {
			maxRight = seg.right
			curSeg = seg
		}
	}

	if nowRight < M {
		nowRight = maxRight
		ans = append(ans, curSeg)
	}

	if nowRight < M {
		fmt.Println("No solution")
	} else {
		fmt.Println(len(ans))
		for _, seg := range ans {
			fmt.Println(seg.left, seg.right)
		}
	}
}
