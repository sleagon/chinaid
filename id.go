package chinaid

import (
	"fmt"
	"strconv"
	"time"
)

const (
	// Female 女
	Female = iota
	// Male 男
	Male
)

const (
	// IDCardLength 身份证长度
	IDCardLength = 18
	// IDCardModBase 取模运算的基
	IDCardModBase = 11
)

// errors
var (
	// ErrInvalidIDCardNo 身份证号不正确
	ErrInvalidIDCardNo = fmt.Errorf("身份证号不合法")
)

var (
	idCardWeight   = []int{7, 9, 10, 5, 8, 4, 2, 1, 6, 3, 7, 9, 10, 5, 8, 4, 2}
	idCardCheckMap = []byte{'1', '0', 'X', '9', '8', '7', '6', '5', '4', '3', '2'}
)

// IDCard 身份证号信息
type IDCard string

// IDCardDetail 解析出来的具体信息
type IDCardDetail struct {
	Sex      int       `json:"sex"`
	Birthday time.Time `json:"birthday"`
	Addr
}

// Valid 判断身份证号是否合法
func (card IDCard) Valid() bool {
	id := string(card)
	// basic check
	switch {
	case len(id) != 18:
		return false
	default:
	}
	// char check
	x := id[IDCardLength-1]
	switch x {
	case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9', 'x', 'X':
	default:
		return false
	}
	for _, v := range id[:IDCardLength-1] {
		switch {
		case v < '0', v > '9':
			return false
		default:
		}
	}

	// sum check
	sum := 0
	for k, v := range idCardWeight {
		sum += v * int(id[k]-'0')
	}
	mod := idCardCheckMap[sum%IDCardModBase]
	if mod != x {
		return false
	}
	return true
}

// Decode 根据身份证号解析出具体信息
func (card IDCard) Decode() (IDCardDetail, error) {
	icd := IDCardDetail{}
	// check
	if !card.Valid() {
		return icd, ErrInvalidIDCardNo
	}
	id := string(card)

	// parse address code
	ac, err := strconv.Atoi(id[:6])
	if err != nil {
		return icd, ErrInvalidIDCardNo
	}
	icd.InitAddr(ac)

	// birthday
	birth, err := time.Parse("20060102", id[6:14])
	if err != nil {
		return icd, ErrInvalidIDCardNo
	}
	icd.Birthday = birth

	// sex
	order, err := strconv.Atoi(id[14:17])
	if err != nil {
		return icd, ErrInvalidIDCardNo
	}
	switch order % 2 {
	case Female:
		icd.Sex = Female
	default:
		icd.Sex = Male
	}
	return icd, nil
}

// Addr 地址信息
// 修复问题: https://github.com/sleagon/chinaid/issues/1 改成3级划分了
type Addr struct {
	Code     int    `json:"code"` // 这里之前是CityCode，改成三级以后，直接叫Code了
	Province string `json:"province"`
	City     string `json:"city"`
	District string `json:"district"`
}

// InitAddr 初始化地址信息
func (addr *Addr) InitAddr(code int) {
	addr.Code = code
	addr.District = "未知"
	addr.City = "未知"
	addr.Province = "未知"
	if district, ex := cityMap[code]; ex {
		addr.District = district
	}
	if city, ex := cityMap[code-code%100]; ex {
		addr.City = city
	}
	if prov, ex := cityMap[code-code%10000]; ex {
		addr.Province = prov
	}
}
