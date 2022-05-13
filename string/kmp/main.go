package main

import "fmt"

func main() {
	fmt.Println(strStr("mississippi", "issip"))
}

func strStr(haystack string, needle string) int {
	// next[i]表示字符串needle[:i+1]中前后缀的最长相同长度（前缀不包括最后一个字符，后缀不包括第一个字符）
	// 在比对的过程中，若needle[j]不匹配（j>0），由于needle[j-1]肯定匹配，所以只需要回退到next[j-1]的字符继续匹配，而不需要回退到开头（next[j-1]对应的字符是最长相同前缀的后一个字符）
    next := make([]int, len(needle))
    j := 0
    for i := 1; i < len(needle); i++ {
        for j > 0 && needle[j] != needle[i] {
            j = next[j-1]
        }
        if needle[i] == needle[j] {
            j ++
        }

        next[i] = j
    }
	fmt.Println(next)

    j = 0
    for i := 0; i < len(haystack); i++ {
        for j > 0 && haystack[i] != needle[j] {
            j = next[j-1]
        }
        if haystack[i] == needle[j] {
            j ++
        }

        if j == len(needle) {
            return i - j + 1
        }
    }
    return -1
}
