package utils

import (
	"math/rand"
	"strings"
	"time"
	"unicode"
)

func LeftOverlayString(in string, overlay string, end int) (overlayed string) {
	return OverlayString(in, overlay, 0, end)
}

func RightOverlayString(in string, overlay string, start int) (overlayed string) {
	return OverlayString(in, overlay, start, len(in))
}

func OverlayString(in string, overlay string, start int, end int) (overlayed string) {
	r := []rune(in)
	l := len([]rune(r))

	if l == 0 {
		return ""
	}

	if start < 0 {
		start = 0
	}
	if start > l {
		start = l
	}
	if end < 0 {
		end = 0
	}
	if end > l {
		end = l
	}

	if start > end {
		start, end = end, start
	}

	overlayed = ""
	overlayed += string(r[:start])
	overlayed += overlay
	overlayed += string(r[end:])
	return overlayed
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func CharMapping(c rune, replace rune, fixed bool) rune {
	if unicode.IsDigit(c) {
		return rune('0' + rand.Intn(10))
	} else if unicode.IsLetter(c) {
		if fixed {
			return replace
		}
		if c <= 90 {
			return rune('A' + rand.Intn(26))
		} else {
			return rune('a' + rand.Intn(26))
		}
	}
	return replace
}

// MaskString 遮盖字符串，优先保留左边 left 位，剩余长度不足 right 时忽略
func MaskString(s string, left, right int, maskChar string) string {
	if left < 0 {
		left = 0
	}
	if right < 0 {
		right = 0
	}
	runes := []rune(s)
	length := len(runes)

	// 如果字符串为空或长度不足，直接返回
	if length == 0 {
		return s
	}

	// 优先保留左边 left 位
	if left >= length {
		return s // 不需要遮盖
	}

	// 计算实际能保留的右边位数（剩余长度是否足够）
	remaining := length - left - 1
	actualRight := right
	if remaining < actualRight {
		actualRight = 0 // 不足时忽略 right
	}

	// 左边部分
	leftPart := string(runes[:left])

	// 右边部分（如果有）
	var rightPart string
	if actualRight > 0 {
		rightPart = string(runes[length-actualRight:])
	}

	// 中间遮盖部分
	maskLength := length - left - actualRight
	maskPart := strings.Repeat(maskChar, maskLength)

	// 拼接结果
	return leftPart + maskPart + rightPart
}
