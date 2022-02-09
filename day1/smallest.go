package day1

import (
	"math"
	"sort"
	"strconv"
)

func absInt64(num int64) int64 { return int64(math.Abs(float64(num))) }

type numSort struct {
	num        int64
	digitCount int
}

func newNumSort(num int64) *numSort {
	digitCount := 0
	for n := absInt64(num); n > 0; n /= 10 {
		digitCount++
	}
	return &numSort{num: num, digitCount: digitCount}
}

func (ns *numSort) Len() int           { return int(ns.digitCount) }
func (ns *numSort) Less(i, j int) bool { return ns.digitAt(i) < ns.digitAt(j) }

func (ns *numSort) digitAt(i int) int64 {
	a := int64(math.Pow10(ns.digitCount - i - 1))
	return (ns.num / a) % 10
}

func (ns *numSort) Swap(i, j int) {
	if i > j {
		i, j = j, i
	}
	a := int64(math.Pow10(ns.digitCount - i - 1))
	b := int64(math.Pow10(ns.digitCount - j - 1))
	vi, vj := (ns.num/a)%10, (ns.num/b)%10
	if vi != vj {
		num := ns.num
		ns.num = ((num/a)-vi)*a + vj*a + (num-(num%(b*10)))%a + vi*b + (num % b)
	}
}

func smallestNumberSort(num int64) int64 {
	if num == 0 {
		return num
	} else if num < 100 && num > -100 {
		a := num / 10
		b := num % 10
		if a > b && b != 0 {
			return b*10 + a
		} else {
			return a*10 + b
		}
	}

	ns := newNumSort(num)
	sort.Sort(ns)

	digit10 := int64(math.Pow10(ns.digitCount - 1))
	if ns.num%digit10 == ns.num {
		// contain 0
		i := int64(10)
		f, l := ns.num/i, ns.num%i
		for f > 10 {
			i *= 10
			f, l = ns.num/i, ns.num%i
		}
		return f*digit10 + l
	}
	return ns.num
}

func smallestNumber(num int64) int64 {
	if num == 0 {
		return num
	} else if num < 100 && num > -100 {
		a := num / 10
		b := num % 10
		if a > b && b != 0 {
			return b*10 + a
		} else {
			return a*10 + b
		}
	}
	result := int64(0)
	ndigit := int64(1)
	for n := absInt64(num); n > 0; n /= 10 {
		ndigit *= 10
	}

	am, n := ndigit/10, num
	a, b := int64(0), n/am
	first := b
repeat:
	needRepeat := false
	for c := am / 10; c > 0; c = am / 10 {
		a = (n / c) % 10
		if (first > a || b > a) && (a != 0 || am < ndigit/10) {
			result += a * am
			needRepeat = needRepeat || a != 0 || (a == 0 && b != 0 && am < ndigit/100)
		} else {
			result += b * am
			b = a
		}
		am = c
	}
	result += b
	if needRepeat {
		num, result = result, 0
		am, n = ndigit/10, num
		b = n / am
		first = b
		goto repeat
	}

	return result
}

type int64Slice []int64

func (i64s int64Slice) Len() int           { return len(i64s) }
func (i64s int64Slice) Swap(i, j int)      { i64s[i], i64s[j] = i64s[j], i64s[i] }
func (i64s int64Slice) Less(i, j int) bool { return i64s[i] < i64s[j] }

func smallestNumberArray(num int64) int64 {
	if num == 0 {
		return num
	}
	ar := make([]int64, 0)
	for num != 0 {
		ar = append(ar, num%10)
		num /= 10
	}
	sort.Sort(int64Slice(ar))
	result := int64(0)
	for ind, i := int64(1), len(ar)-1; i >= 0; ind, i = ind*10, i-1 {
		if i > 0 && ar[i-1] == 0 && ar[i] > 0 {
			ar[0], ar[i] = ar[i], ar[0]
		}
		result += ind * ar[i]
	}
	return result
}

type byteSlice []byte

func (bs byteSlice) Len() int           { return len(bs) }
func (bs byteSlice) Swap(i, j int)      { bs[i], bs[j] = bs[j], bs[i] }
func (bs byteSlice) Less(i, j int) bool { return bs[i] < bs[j] }

func smallestNumberString(num int64) int64 {
	if num == 0 {
		return num
	}
	m := int64(1)
	buf := []byte(strconv.FormatInt(absInt64(num), 10))
	if num > 0 {
		sort.Sort(byteSlice(buf))
		if buf[0] == '0' {
			for i := 0; i < len(buf); i++ {
				if buf[i] != '0' {
					buf[0], buf[i] = buf[i], buf[0]
					break
				}
			}
		}
	} else {
		sort.Slice(buf, func(i, j int) bool {
			return buf[i] > buf[j]
		})
		m = -1
	}
	result, err := strconv.ParseInt(string(buf), 10, 64)
	if err != nil {
		panic(err)
	}
	return result * m
}
