package core

/**
 * Created by @CaomaoBoy on 2021/5/1.
 *  email:<115882934@qq.com>
 */


type Memorial struct {
	Index int
	Puzzle string
	*Coordinate
}

//经纬度
type Coordinate struct {
	Longitude float64 //经度
	Latitude float64 //纬度
}
