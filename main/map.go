package main

import "fmt"

// 自定義之資料型態
type ProductOption struct {
	code      string
	localName string
}

func (d *ProductOption) getCode() string {
	return d.code
}

func (d *ProductOption) getLocalName() string {
	return d.localName
}

func initMap() {

	m := make(map[string]map[string]ProductOption)

	content := make(map[string]ProductOption)
	content["Personal Accident"] = ProductOption{"IPA", "Personal Accident"}
	//content["localName"] = "Personal Accident"
	m["IPA"] = content

	fmt.Println(m)

}

func dataSupplier() []ProductOption {
	options := []ProductOption{
		ProductOption{"IPA", "1"},
		ProductOption{"GPA", "2"},
		ProductOption{"IPA", "3"},
		ProductOption{"CTA", "4"},
		ProductOption{"IPA", "5"},
		ProductOption{"CTA", "6"},
		ProductOption{"GPA", "7"},
		ProductOption{"CTA", "8"},
		ProductOption{"IPA", "9"},
	}
	return options
}

/*
 * options 原始資料
 * getter 相當於java中的classifier
 * return 回傳Map型態資料
 */
func groupBy(options []ProductOption, getter func(d *ProductOption) string) map[string][]ProductOption {

	productOptions := make(map[string][]ProductOption, len(options))
	// 先取出所有code
	allKeys := MapProductOption(options, func(s ProductOption) string {
		return getter(&s)
	})
	// 針對每一種code過濾出Array資料
	// 依指定的key欄位, 塞回回傳Map物件
	for _, v := range allKeys {
		filterd := FilterProductOption(options, func(s ProductOption) bool {
			// ＊＊＊＊＊ 對該物件使用Getter ＊＊＊＊＊
			return getter(&s) == v
		})
		productOptions[v] = filterd
	}
	return productOptions
}

func Index(vs []string, t string) int {
	for i, v := range vs {
		if v == t {
			return i
		}
	}
	return -1
}

func Include(vs []string, t string) bool {
	return Index(vs, t) >= 0
}

func Any(vs []string, f func(string) bool) bool {
	for _, v := range vs {
		if f(v) {
			return true
		}
	}
	return false
}

func All(vs []string, f func(string) bool) bool {
	for _, v := range vs {
		if !f(v) {
			return false
		}
	}
	return true
}

func Filter(vs []string, f func(string) bool) []string {
	vsf := make([]string, 0)
	for _, v := range vs {
		if f(v) {
			vsf = append(vsf, v)
		}
	}
	return vsf
}

func Map(vs []string, mapper func(string) string) []string {
	vsm := make([]string, len(vs))
	for k, v := range vs {
		vsm[k] = mapper(v)
	}
	return vsm
}

func MapProductOption(vs []ProductOption, mapper func(ProductOption) string) []string {
	vsm := make([]string, len(vs))
	for k, v := range vs {
		vsm[k] = mapper(v)
	}
	return vsm
}

func FilterProductOption(vs []ProductOption, f func(ProductOption) bool) []ProductOption {
	vsf := make([]ProductOption, 0)
	for _, v := range vs {
		if f(v) {
			vsf = append(vsf, v)
		}
	}
	return vsf
}
