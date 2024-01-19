/* For license and copyright information please see the LEGAL file in the code repository */

package ms

import "strconv"

const kvNumber = 16
const notExistKey = "NotExistKey"

type smallMap map[string]string

func (sm *smallMap) Init() {
	*sm = make(smallMap, kvNumber)
}
func (sm smallMap) Fill() {
	for i := 0; i < kvNumber; i++ {
		var index = strconv.Itoa(i)
		sm[index] = index
	}
}
func (ss smallMap) GetMiddle() string {
	var middle = kvNumber / 2
	var middleIndex = strconv.Itoa(middle)
	return ss[middleIndex]
}
func (ss smallMap) GetNotExist() string {
	return ss[notExistKey]
}

type smallSlice []smallSlice_KV

type smallSlice_KV struct {
	key   string
	value string
}

func (ss *smallSlice) Init() {
	*ss = make(smallSlice, kvNumber)
}
func (ss smallSlice) Fill() {
	for i := 0; i < kvNumber; i++ {
		var index = strconv.Itoa(i)
		ss[i] = smallSlice_KV{index, index}
	}
}
func (ss smallSlice) Find(key string) string {
	for i := 0; i < kvNumber; i++ {
		if ss[i].key == key {
			return ss[i].value
		}
	}
	return ""
}
func (ss smallSlice) GetMiddle() string {
	var middle = kvNumber / 2
	var middleIndex = strconv.Itoa(middle)
	return ss.Find(middleIndex)
}
func (ss smallSlice) GetNotExist() string {
	return ss.Find(notExistKey)
}
