package main

import (
	"fmt"
)

func main() {

	sli := []int{2, 3, 4, 6, 8, 5}
	res := findPeek(sli, 0, len(sli)-1)
	fmt.Println(res)
}

func findPeek(sli []int, l, r int) int {
	if l == r {
		return l
	}
	mid := int((l + r) / 2)
	if sli[mid] < sli[mid+1] {
		return findPeek(sli, mid+1, r)
	} else {
		return findPeek(sli, l, mid)
	}

}

/*

public class Solution {
    public int findPeakElement(int[] nums) {
        return search(nums, 0, nums.length - 1);
    }
    public int search(int[] nums, int l, int r) {
        if (l == r)
            return l;
        int mid = (l + r) / 2;
        if (nums[mid] > nums[mid + 1])
            return search(nums, l, mid);
        return search(nums, mid + 1, r);
    }
}
*/
