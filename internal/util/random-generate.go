package util

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

//
//返回一串位数为digits的随机数字串
//
func GetRamdomCode(digits int) string {
	rand.Seed(time.Now().Unix())
	randomId := rand.Intn(int(math.Pow(10, float64(digits))))
	return fmt.Sprintf("%d", randomId)
}
