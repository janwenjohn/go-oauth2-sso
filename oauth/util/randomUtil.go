package util

import (
	"bytes"
	r "crypto/rand"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

/**
*生成随机字符
**/
func RandString(length int) string {
	rand.Seed(time.Now().UnixNano())
	rs := make([]string, length)
	for start := 0; start < length; start++ {
		t := rand.Intn(3)
		if t == 0 {
			rs = append(rs, strconv.Itoa(rand.Intn(10)))
		} else if t == 1 {
			rs = append(rs, string(rand.Intn(26)+65))
		} else {
			rs = append(rs, string(rand.Intn(26)+97))
		}
	}
	return strings.Join(rs, "")
}

func Rs2(length int) []byte {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	buffer := bytes.NewBufferString("")
	for i := 0; i < length; i++ {
		isLetter := r.Intn(2)
		if isLetter > 0 {
			letter := r.Intn(52)
			if letter < 26 {
				letter += 97
			} else {
				letter += 65 - 26
			}
			buffer.WriteString(string(letter))
			//buffer.WriteString(fmt.Sprintf("%c", letter))
		} else {
			buffer.WriteString(strconv.Itoa(r.Intn(10)))
		}
	}
	return buffer.Bytes()
}

func RandomCreateBytes(n int, alphabets ...byte) []byte {
	const alphanum = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	var bytes = make([]byte, n)
	var randby bool
	if num, err := r.Read(bytes); num != n || err != nil {
		rand.Seed(time.Now().UnixNano())
		randby = true
	}
	for i, b := range bytes {
		if len(alphabets) == 0 {
			if randby {
				bytes[i] = alphanum[rand.Intn(len(alphanum))]
			} else {
				bytes[i] = alphanum[b%byte(len(alphanum))]
			}
		} else {
			if randby {
				bytes[i] = alphabets[rand.Intn(len(alphabets))]
			} else {
				bytes[i] = alphabets[b%byte(len(alphabets))]
			}
		}
	}
	return bytes
}