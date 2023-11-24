package top75

import (
	"math"
	"sort"
	"strings"
)

// GcdOfStrings 1071. 字符串的最大公因子
// 对于字符串s 和t，只有在s = t + ... + t（t 自身连接 1 次或多次）时，我们才认定“t 能除尽 s”。
// 给定两个字符串str1和str2。返回 最长字符串x，要求满足x 能除尽 str1 且 x 能除尽 str2 。
// str1 = "ABABAB", str2 = "ABAB" => "AB"
func GcdOfStrings(str1 string, str2 string) string {
	var short, long string
	if len(str1) > len(str2) {
		short = str2
		long = str1
	} else {
		short = str1
		long = str2
	}
	var result = ""
	for i := 0; i < len(short); i++ {
		gcd := findGcd(short, long, i+1)
		if len(gcd) > len(result) {
			result = gcd
		}
	}
	return result
}

func findGcd(short string, long string, index int) string {
	if len(short)%index != 0 || len(long)%index != 0 {
		return ""
	}
	shortCount, longCount := len(short)/index, len(long)/index
	var shortAppend, longAppend string
	indexStr := short[0:index]
	for i := 0; i < shortCount; i++ {
		shortAppend = shortAppend + indexStr
	}
	for i := 0; i < longCount; i++ {
		longAppend = longAppend + indexStr
	}
	if longAppend == long && shortAppend == short {
		return indexStr
	}
	return ""
}

// kidsWithCandies 1431. 拥有最多糖果的孩子
//给你一个数组candies和一个整数extraCandies，其中candies[i]代表第 i 个孩子拥有的糖果数目。
//
//对每一个孩子，检查是否存在一种方案，将额外的extraCandies个糖果分配给孩子们之后，此孩子有 最多的糖果。
//注意，允许有多个孩子同时拥有 最多的糖果数目。
//输入：candies = [2,3,5,1,3], extraCandies = 3
//输出：[true,true,true,false,true]
//解释：
//孩子 1 有 2 个糖果，如果他得到所有额外的糖果（3个），那么他总共有 5 个糖果，他将成为拥有最多糖果的孩子。
//孩子 2 有 3 个糖果，如果他得到至少 2 个额外糖果，那么他将成为拥有最多糖果的孩子。
//孩子 3 有 5 个糖果，他已经是拥有最多糖果的孩子。
//孩子 4 有 1 个糖果，即使他得到所有额外的糖果，他也只有 4 个糖果，无法成为拥有糖果最多的孩子。
//孩子 5 有 3 个糖果，如果他得到至少 2 个额外糖果，那么他将成为拥有最多糖果的孩子。
func kidsWithCandies(candies []int, extraCandies int) []bool {
	result := make([]bool, len(candies))
	max := 0
	for i := 0; i < len(candies); i++ {
		if max < candies[i] {
			max = candies[i]
		}
	}
	for i := 0; i < len(candies); i++ {
		if max <= extraCandies+candies[i] {
			result[i] = true
		}
	}
	return result
}

// canPlaceFlowers 605. 种花问题
// 假设有一个很长的花坛，一部分地块种植了花，另一部分却没有。可是，花不能种植在相邻的地块上，它们会争夺水源，两者都会死去。
//
//给你一个整数数组flowerbed 表示花坛，由若干 0 和 1 组成，其中 0 表示没种植花，1 表示种植了花。
//另有一个数n ，能否在不打破种植规则的情况下种入n朵花？能则返回 true ，不能则返回 false。
// flowerbed = [1,0,0,0,1], n = 1 -> true
func canPlaceFlowers(flowerbed []int, n int) bool {
	count := 0
	if len(flowerbed) == 1 {
		return 1-flowerbed[0] >= n
	}
	for i := 0; i < len(flowerbed); i++ {
		if flowerbed[i] == 1 {
			continue
		}
		if i == 0 {
			if flowerbed[i+1] == 0 {
				flowerbed[i] = 1
				count++
			}
		} else if flowerbed[i-1] == 0 {
			if i == len(flowerbed)-1 {
				flowerbed[i] = 1
				count++
			} else if flowerbed[i+1] == 0 {
				flowerbed[i] = 1
				count++
			}
		}
	}
	return count >= n
}

// reverseVowels 345. 反转字符串中的元音字母
// 给你一个字符串 s ，仅反转字符串中的所有元音字母，并返回结果字符串。
// 元音字母包括 'a'、'e'、'i'、'o'、'u'，且可能以大小写两种形式出现不止一次。
func reverseVowels(s string) string {
	if len(s) == 1 {
		return s
	}
	left, right := 0, len(s)-1
	result := []byte(s)
	for left < right {
		if left < len(s) && !strings.Contains("aeiouAEIOU", string(s[left])) {
			left++
		}
		if right > 0 && !strings.Contains("aeiouAEIOU", string(s[right])) {
			right--
		}
		if left < right {
			result[left], result[right] = result[right], result[left]
			left++
			right--
		}
	}
	return string(result)
}

// increasingTriplet 334. 递增的三元子序列
func increasingTriplet(nums []int) bool {
	if len(nums) < 3 {
		return false
	}
	first, sceond := math.MaxInt, math.MaxInt
	for _, item := range nums {
		if item <= first {
			first = item
		} else if item <= sceond {
			sceond = item
		} else {
			return true
		}
	}
	return false
}

// threeSum 15. 三数之和
//给你一个整数数组 nums ，判断是否存在三元组 [nums[i], nums[j], nums[k]]
//满足 i != j、i != k 且 j != k ，同时还满足 nums[i] + nums[j] + nums[k] == 0 。请
//你返回所有和为 0 且不重复的三元组。
//输入：nums = [-1,0,1,2,-1,-4]
//输出：[[-1,-1,2],[-1,0,1]]

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	if len(nums) == 0 || nums[0] > 0 {
		return nil
	}
	result := make([][]int, 0)
	length := len(nums)
	for i := 0; i < length; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		left, right := i+1, length-1
		for left < right {
			sum := nums[i] + nums[left] + nums[right]
			if sum > 0 {
				right--
			} else if sum < 0 {
				left++
			} else {
				result = append(result, []int{nums[i], nums[left], nums[right]})
				for left < right && nums[left] == nums[left+1] {
					left++
				}
				for left < right && nums[right] == nums[right-1] {
					right--
				}
				left++
				right--
			}
		}
	}
	return result
}

func threeSum2(nums []int) [][]int {
	sort.Ints(nums)
	result := make([][]int, 0)
	length := len(nums)
	for i := 0; i < length; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		target, left, right := 0-nums[i], i+1, length-1
		for left < right {
			if left+right < target {
				right++
			} else if left+right > target {
				right--
			} else {
				result = append(result, []int{nums[i], nums[left], nums[right]})
				for left < right && nums[right] == nums[right-1] {
					right--
				}
				for left < right && nums[left] == nums[left+1] {
					left++
				}
				left++
				right--
			}
		}

	}
	return result
}
