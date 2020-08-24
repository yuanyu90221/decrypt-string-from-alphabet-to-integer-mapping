# decrypt-string-from-alphabet-to-integer-mapping

## 題目解讀：

### 題目來源:
[decrypt-string-from-alphabet-to-integer-mapping](https://leetcode.com/problems/decrypt-string-from-alphabet-to-integer-mapping/)

### 原文:
Given a string s formed by digits ('0' - '9') and '#' . We want to map s to English lowercase characters as follows:

Characters ('a' to 'i') are represented by ('1' to '9') respectively.
Characters ('j' to 'z') are represented by ('10#' to '26#') respectively. 
Return the string formed after mapping.

It's guaranteed that a unique mapping will always exist.

### 解讀：

給定一個字串 s , s由 字元 '0' 到 '9' 還有 '#'組成

想要把s轉換成英文字元組成的字串

轉換規則如下:

'a' 到 'i' 對應到 '1'到'9'

'j' 到 'z' 對應到 '10#'到'26#'


## 初步解法:
### 初步觀察:

首先可以知道

因為 有部份字元對會被包含在其他字元

比如 10# 包含了 1跟 0

所以如果從左解讀會不知道如何切字元

因此需要反過來切割

因為 每個 # 都會需要搭配兩個數字字元

所以當遇到 # 就需要往右找2個字元 來切割


### 初步設計:
Given a string s

Step 0: let idx = len(s) - 1, a string result = ""

Step 1: if idx  < 0  go to step 7

Step 2: let temp = ""

Step 3: if s[idx] == '#' set temp = s[idx-2] + s[idx-1]+ s[idx] , idx = idx - 3

Step 4: if s[idx] >= '0' && s[idx] <= 9, temp = s[idx], idx = idx - 1

Step 5: find the target Character char by map rule, set result = tChar + result

Step 6： return result
## 遇到的困難
### 題目上理解的問題
因為英文不是筆者母語

所以在題意解讀上 容易被英文用詞解讀給搞模糊

### pseudo code撰寫

一開始不習慣把pseudo code寫下來

因此 不太容易把自己的code做解析

### golang table driven test不熟
對於table driven test還不太熟析

所以對於寫test還是耗費不少時間
## 我的github source code
```golang
package freq_alphabets

import "strconv"

func freqAlphabets(s string) string {
	result := ""
	lenS := len(s)
	for idx := lenS - 1; idx >= 0; {
		if s[idx] == '#' {
			temp := string(s[idx-2]) + string(s[idx-1])
			tempResult, _ := strconv.Atoi(temp)
			result = string('a'+tempResult-1) + result
			idx -= 3
		} else {
			temp := string(s[idx])
			tempResult, _ := strconv.Atoi(temp)
			result = string('a'+tempResult-1) + result
			idx -= 1
		}
	}
	return result
}

```
## 測資的撰寫
```golang
package freq_alphabets

import "testing"

func Test_freqAlphabets(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Example1",
			args: args{
				s: "10#11#12",
			},
			want: "jkab",
		},
		{
			name: "Example2",
			args: args{
				s: "1326#",
			},
			want: "acz",
		},
		{
			name: "Example3",
			args: args{
				s: "25#",
			},
			want: "y",
		},
		{
			name: "Example4",
			args: args{
				s: "12345678910#11#12#13#14#15#16#17#18#19#20#21#22#23#24#25#26#",
			},
			want: "abcdefghijklmnopqrstuvwxyz",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := freqAlphabets(tt.args.s); got != tt.want {
				t.Errorf("freqAlphabets() = %v, want %v", got, tt.want)
			}
		})
	}
}


```
## my self record
[golang leetcode 30day 27th day](https://hackmd.io/VRC5Pxh_SQu2qMyNt9P83A?view)
## 參考文章

[golang test](https://ithelp.ithome.com.tw/articles/10204692)