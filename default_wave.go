package godglab

var DefaultWaves map[string]WavePatten

func init() {
	DefaultWaves = make(map[string]WavePatten)
	DefaultWaves["呼吸"] = WavePatten{Name: "呼吸", Data: []Pulse{
		{10, 10, 10, 10, 0, 0, 0, 0},
		{10, 10, 10, 10, 0, 5, 10, 20},
		{10, 10, 10, 10, 20, 25, 30, 40},
		{10, 10, 10, 10, 40, 45, 50, 60},
		{10, 10, 10, 10, 60, 65, 70, 80},
		{10, 10, 10, 10, 100, 100, 100, 100},
		{10, 10, 10, 10, 100, 100, 100, 100},
		{10, 10, 10, 10, 100, 100, 100, 100},
		{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0}}}
	DefaultWaves["潮汐"] = WavePatten{Name: "潮汐", Data: []Pulse{
		{10, 10, 10, 10, 0, 0, 0, 0},
		{10, 10, 10, 10, 0, 4, 8, 17},
		{10, 10, 10, 10, 17, 21, 25, 33},
		{10, 10, 10, 10, 50, 50, 50, 50},
		{10, 10, 10, 10, 50, 54, 58, 67},
		{10, 10, 10, 10, 67, 71, 75, 83},
		{10, 10, 10, 10, 100, 100, 100, 100},
		{10, 10, 10, 10, 100, 98, 96, 92},
		{10, 10, 10, 10, 92, 90, 88, 84},
		{10, 10, 10, 10, 84, 82, 80, 76},
		{10, 10, 10, 10, 68, 68, 68, 68}}}
	DefaultWaves["连击"] = WavePatten{Name: "连击", Data: []Pulse{
		{10, 10, 10, 10, 100, 100, 100, 100},
		{10, 10, 10, 10, 0, 0, 0, 0},
		{10, 10, 10, 10, 100, 100, 100, 100},
		{10, 10, 10, 10, 100, 92, 84, 67},
		{10, 10, 10, 10, 67, 58, 50, 33},
		{10, 10, 10, 10, 0, 0, 0, 0},
		{10, 10, 10, 10, 0, 0, 0, 1},
		{10, 10, 10, 10, 2, 2, 2, 2}}}
	DefaultWaves["快速按捏"] = WavePatten{Name: "快速按捏", Data: []Pulse{
		{10, 10, 10, 10, 0, 0, 0, 0},
		{10, 10, 10, 10, 100, 100, 100, 100},
		{0, 0, 0, 0, 0, 0, 0, 0}}}
	DefaultWaves["按捏渐强"] = WavePatten{Name: "按捏渐强", Data: []Pulse{
		{10, 10, 10, 10, 0, 0, 0, 0},
		{10, 10, 10, 10, 29, 29, 29, 29},
		{10, 10, 10, 10, 0, 0, 0, 0},
		{10, 10, 10, 10, 52, 52, 52, 52},
		{10, 10, 10, 10, 2, 2, 2, 2},
		{10, 10, 10, 10, 73, 73, 73, 73},
		{10, 10, 10, 10, 0, 0, 0, 0},
		{10, 10, 10, 10, 87, 87, 87, 87},
		{10, 10, 10, 10, 0, 0, 0, 0},
		{10, 10, 10, 10, 100, 100, 100, 100},
		{10, 10, 10, 10, 0, 0, 0, 0}}}
	DefaultWaves["心跳节奏"] = WavePatten{Name: "心跳节奏", Data: []Pulse{
		{110, 110, 110, 110, 100, 100, 100, 100},
		{110, 110, 110, 110, 100, 100, 100, 100},
		{10, 10, 10, 10, 0, 0, 0, 0},
		{10, 10, 10, 10, 0, 0, 0, 0},
		{10, 10, 10, 10, 0, 0, 0, 0},
		{10, 10, 10, 10, 0, 0, 0, 0},
		{10, 10, 10, 10, 0, 0, 0, 0},
		{10, 10, 10, 10, 75, 75, 75, 75},
		{10, 10, 10, 10, 75, 77, 79, 83},
		{10, 10, 10, 10, 83, 85, 88, 92},
		{10, 10, 10, 10, 100, 100, 100, 100},
		{10, 10, 10, 10, 0, 0, 0, 0},
		{10, 10, 10, 10, 0, 0, 0, 0},
		{10, 10, 10, 10, 0, 0, 0, 0},
		{10, 10, 10, 10, 0, 0, 0, 0},
		{10, 10, 10, 10, 0, 0, 0, 0}}}
	DefaultWaves["压缩"] = WavePatten{Name: "压缩", Data: []Pulse{
		{25, 25, 24, 24, 100, 100, 100, 100},
		{24, 23, 23, 23, 100, 100, 100, 100},
		{22, 22, 22, 21, 100, 100, 100, 100},
		{21, 21, 20, 20, 100, 100, 100, 100},
		{20, 19, 19, 19, 100, 100, 100, 100},
		{18, 18, 18, 17, 100, 100, 100, 100},
		{17, 16, 16, 16, 100, 100, 100, 100},
		{15, 15, 15, 14, 100, 100, 100, 100},
		{14, 14, 13, 13, 100, 100, 100, 100},
		{13, 12, 12, 12, 100, 100, 100, 100},
		{11, 11, 11, 10, 100, 100, 100, 100},
		{10, 10, 10, 10, 100, 100, 100, 100},
		{10, 10, 10, 10, 100, 100, 100, 100},
		{10, 10, 10, 10, 100, 100, 100, 100},
		{10, 10, 10, 10, 100, 100, 100, 100},
		{10, 10, 10, 10, 100, 100, 100, 100},
		{10, 10, 10, 10, 100, 100, 100, 100},
		{10, 10, 10, 10, 100, 100, 100, 100},
		{10, 10, 10, 10, 100, 100, 100, 100},
		{10, 10, 10, 10, 100, 100, 100, 100},
		{10, 10, 10, 10, 100, 100, 100, 100}}}
	DefaultWaves["节奏步伐"] = WavePatten{Name: "节奏步伐", Data: []Pulse{
		{10, 10, 10, 10, 0, 0, 0, 0},
		{10, 10, 10, 10, 0, 5, 10, 20},
		{10, 10, 10, 10, 20, 25, 30, 40},
		{10, 10, 10, 10, 40, 45, 50, 60},
		{10, 10, 10, 10, 60, 65, 70, 80},
		{10, 10, 10, 10, 100, 100, 100, 100},
		{10, 10, 10, 10, 0, 0, 0, 0},
		{10, 10, 10, 10, 0, 6, 12, 25},
		{10, 10, 10, 10, 25, 31, 38, 50},
		{10, 10, 10, 10, 50, 56, 62, 75},
		{10, 10, 10, 10, 100, 100, 100, 100},
		{10, 10, 10, 10, 0, 0, 0, 0},
		{10, 10, 10, 10, 0, 8, 16, 33},
		{10, 10, 10, 10, 33, 42, 50, 67},
		{10, 10, 10, 10, 100, 100, 100, 100},
		{10, 10, 10, 10, 0, 0, 0, 0},
		{10, 10, 10, 10, 0, 12, 25, 50},
		{10, 10, 10, 10, 100, 100, 100, 100},
		{10, 10, 10, 10, 0, 0, 0, 0},
		{10, 10, 10, 10, 100, 100, 100, 100},
		{10, 10, 10, 10, 0, 0, 0, 0},
		{10, 10, 10, 10, 100, 100, 100, 100},
		{10, 10, 10, 10, 0, 0, 0, 0},
		{10, 10, 10, 10, 100, 100, 100, 100},
		{10, 10, 10, 10, 0, 0, 0, 0},
		{10, 10, 10, 10, 100, 100, 100, 100}}}
	DefaultWaves["颗粒摩擦"] = WavePatten{Name: "颗粒摩擦", Data: []Pulse{
		{10, 10, 10, 10, 100, 100, 100, 100},
		{10, 10, 10, 10, 100, 100, 100, 100},
		{10, 10, 10, 10, 100, 100, 100, 100},
		{10, 10, 10, 10, 0, 0, 0, 0}}}
	DefaultWaves["渐变弹跳"] = WavePatten{Name: "渐变弹跳", Data: []Pulse{
		{10, 10, 10, 10, 1, 1, 1, 1},
		{10, 10, 10, 10, 1, 9, 18, 34},
		{10, 10, 10, 10, 34, 42, 50, 67},
		{10, 10, 10, 10, 100, 100, 100, 100},
		{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0}}}
	DefaultWaves["波浪涟漪"] = WavePatten{Name: "波浪涟漪", Data: []Pulse{
		{10, 10, 10, 10, 0, 0, 0, 0},
		{10, 10, 10, 10, 0, 12, 25, 50},
		{10, 10, 10, 10, 100, 100, 100, 100},
		{10, 10, 10, 10, 73, 73, 73, 73}}}
	DefaultWaves["雨水冲刷"] = WavePatten{Name: "雨水冲刷", Data: []Pulse{
		{10, 10, 10, 10, 34, 34, 34, 34},
		{10, 10, 10, 10, 34, 42, 50, 67},
		{10, 10, 10, 10, 100, 100, 100, 100},
		{10, 10, 10, 10, 100, 100, 100, 100},
		{10, 10, 10, 10, 100, 100, 100, 100},
		{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0}}}
	DefaultWaves["变速敲击"] = WavePatten{Name: "变速敲击", Data: []Pulse{
		{10, 10, 10, 10, 100, 100, 100, 100},
		{10, 10, 10, 10, 100, 100, 100, 100},
		{10, 10, 10, 10, 100, 100, 100, 100},
		{10, 10, 10, 10, 0, 0, 0, 0},
		{10, 10, 10, 10, 0, 0, 0, 0},
		{10, 10, 10, 10, 0, 0, 0, 0},
		{10, 10, 10, 10, 0, 0, 0, 0},
		{110, 110, 110, 110, 100, 100, 100, 100},
		{110, 110, 110, 110, 100, 100, 100, 100},
		{110, 110, 110, 110, 100, 100, 100, 100},
		{110, 110, 110, 110, 100, 100, 100, 100},
		{0, 0, 0, 0, 0, 0, 0, 0}}}
	DefaultWaves["信号灯"] = WavePatten{Name: "信号灯", Data: []Pulse{
		{197, 197, 197, 197, 100, 100, 100, 100},
		{197, 197, 197, 197, 100, 100, 100, 100},
		{197, 197, 197, 197, 100, 100, 100, 100},
		{197, 197, 197, 197, 100, 100, 100, 100},
		{10, 10, 10, 10, 0, 0, 0, 0},
		{10, 10, 10, 10, 0, 8, 16, 33},
		{10, 10, 10, 10, 33, 42, 50, 67},
		{10, 10, 10, 10, 100, 100, 100, 100}}}
	DefaultWaves["挑逗1"] = WavePatten{Name: "挑逗1", Data: []Pulse{
		{10, 10, 10, 10, 0, 0, 0, 0},
		{10, 10, 10, 10, 0, 6, 12, 25},
		{10, 10, 10, 10, 25, 31, 38, 50},
		{10, 10, 10, 10, 50, 56, 62, 75},
		{10, 10, 10, 10, 100, 100, 100, 100},
		{10, 10, 10, 10, 100, 100, 100, 100},
		{10, 10, 10, 10, 100, 100, 100, 100},
		{10, 10, 10, 10, 0, 0, 0, 0},
		{10, 10, 10, 10, 0, 0, 0, 0},
		{10, 10, 10, 10, 0, 0, 0, 0},
		{10, 10, 10, 10, 0, 0, 0, 0},
		{10, 10, 10, 10, 100, 100, 100, 100}}}
	DefaultWaves["挑逗2"] = WavePatten{Name: "挑逗2", Data: []Pulse{
		{10, 10, 10, 10, 1, 1, 1, 1},
		{10, 10, 10, 10, 1, 4, 6, 12},
		{10, 10, 10, 10, 12, 15, 18, 23},
		{10, 10, 10, 10, 23, 26, 28, 34},
		{10, 10, 10, 10, 34, 37, 40, 45},
		{10, 10, 10, 10, 45, 48, 50, 56},
		{10, 10, 10, 10, 56, 59, 62, 67},
		{10, 10, 10, 10, 67, 70, 72, 78},
		{10, 10, 10, 10, 78, 81, 84, 89},
		{10, 10, 10, 10, 100, 100, 100, 100},
		{10, 10, 10, 10, 100, 100, 100, 100},
		{10, 10, 10, 10, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0}}}
}

// var (
// 	// 向上平移
// 	Sin1 = []Pulse{
// 		{10, 10, 10, 10, 0, 0, 2, 5},
// 		{10, 10, 10, 10, 9, 14, 20, 27},
// 		{10, 10, 10, 10, 34, 42, 49, 57},
// 		{10, 10, 10, 10, 65, 72, 79, 85},
// 		{10, 10, 10, 10, 90, 94, 97, 99},
// 		{10, 10, 10, 10, 100, 99, 97, 94},
// 		{10, 10, 10, 10, 90, 85, 79, 72},
// 		{10, 10, 10, 10, 65, 57, 50, 42},
// 		{10, 10, 10, 10, 34, 27, 20, 14},
// 		{10, 10, 10, 10, 9, 5, 2, 0},
// 	}
// 	// 整流
// 	Sin2 = []Pulse{
// 		{10, 10, 10, 10, 0, 15, 30, 45},
// 		{10, 10, 10, 10, 58, 70, 80, 89},
// 		{10, 10, 10, 10, 95, 98, 100, 98},
// 		{10, 10, 10, 10, 95, 89, 80, 70},
// 		{10, 10, 10, 10, 58, 45, 30, 15},
// 		{10, 10, 10, 10, 0, 15, 30, 45},
// 		{10, 10, 10, 10, 58, 70, 80, 89},
// 		{10, 10, 10, 10, 95, 98, 100, 98},
// 		{10, 10, 10, 10, 95, 89, 80, 70},
// 		{10, 10, 10, 10, 58, 45, 30, 15},
// 	}
// 	Square = []Pulse{
// 		{10, 10, 10, 10, 100, 100, 100, 100},
// 		{10, 10, 10, 10, 100, 100, 100, 100},
// 		{10, 10, 10, 10, 100, 100, 100, 100},
// 		{10, 10, 10, 10, 100, 100, 100, 100},
// 		{10, 10, 10, 10, 100, 100, 100, 100},
// 		{10, 10, 10, 10, 0, 0, 0, 0},
// 		{10, 10, 10, 10, 0, 0, 0, 0},
// 		{10, 10, 10, 10, 0, 0, 0, 0},
// 		{10, 10, 10, 10, 0, 0, 0, 0},
// 		{10, 10, 10, 10, 0, 0, 0, 0},
// 	}
// 	Square1d25 = []Pulse{
// 		{10, 10, 10, 10, 100, 100, 100, 100},
// 		{10, 10, 10, 10, 100, 100, 100, 100},
// 		{10, 10, 10, 10, 100, 100, 100, 100},
// 		{10, 10, 10, 10, 100, 100, 100, 100},
// 		{10, 10, 10, 10, 0, 0, 0, 0},
// 		{10, 10, 10, 10, 0, 0, 0, 0},
// 		{10, 10, 10, 10, 0, 0, 0, 0},
// 		{10, 10, 10, 10, 0, 0, 0, 0},
// 	}
// 	Square1d67 = []Pulse{
// 		{10, 10, 10, 10, 100, 100, 100, 100},
// 		{10, 10, 10, 10, 100, 100, 100, 100},
// 		{10, 10, 10, 10, 100, 100, 100, 100},
// 		{10, 10, 10, 10, 0, 0, 0, 0},
// 		{10, 10, 10, 10, 0, 0, 0, 0},
// 		{10, 10, 10, 10, 0, 0, 0, 0},
// 	}
// 	Square2d5 = []Pulse{
// 		{10, 10, 10, 10, 100, 100, 100, 100},
// 		{10, 10, 10, 10, 100, 100, 100, 100},
// 		{10, 10, 10, 10, 0, 0, 0, 0},
// 		{10, 10, 10, 10, 0, 0, 0, 0},
// 	}
// 	Square5 = []Pulse{
// 		{10, 10, 10, 10, 100, 100, 100, 100},
// 		{10, 10, 10, 10, 0, 0, 0, 0},
// 	}
// 	Sawtooth = []Pulse{
// 		{10, 10, 10, 10, 2, 5, 7, 10},
// 		{10, 10, 10, 10, 12, 15, 17, 20},
// 		{10, 10, 10, 10, 22, 25, 27, 30},
// 		{10, 10, 10, 10, 32, 35, 37, 40},
// 		{10, 10, 10, 10, 42, 45, 47, 50},
// 		{10, 10, 10, 10, 52, 55, 57, 60},
// 		{10, 10, 10, 10, 62, 65, 67, 70},
// 		{10, 10, 10, 10, 72, 75, 77, 80},
// 		{10, 10, 10, 10, 82, 85, 87, 90},
// 		{10, 10, 10, 10, 92, 95, 97, 100},
// 	}
// 	// 1.25倍速
// 	Sawtooth1d25 = []Pulse{
// 		{10, 10, 10, 10, 3, 6, 9, 12},
// 		{10, 10, 10, 10, 15, 18, 21, 25},
// 		{10, 10, 10, 10, 28, 31, 34, 37},
// 		{10, 10, 10, 10, 40, 43, 46, 50},
// 		{10, 10, 10, 10, 53, 56, 59, 62},
// 		{10, 10, 10, 10, 65, 68, 71, 75},
// 		{10, 10, 10, 10, 78, 81, 84, 87},
// 		{10, 10, 10, 10, 90, 93, 96, 100},
// 	}
// 	// 1.67倍速
// 	Sawtooth1d67 = []Pulse{
// 		{10, 10, 10, 10, 4, 8, 12, 16},
// 		{10, 10, 10, 10, 20, 25, 29, 33},
// 		{10, 10, 10, 10, 37, 41, 45, 50},
// 		{10, 10, 10, 10, 54, 58, 62, 66},
// 		{10, 10, 10, 10, 70, 75, 79, 83},
// 		{10, 10, 10, 10, 87, 91, 95, 100},
// 	}
// 	// 2倍速
// 	Sawtooth2 = []Pulse{
// 		{10, 10, 10, 10, 5, 10, 15, 20},
// 		{10, 10, 10, 10, 25, 30, 35, 40},
// 		{10, 10, 10, 10, 45, 50, 55, 60},
// 		{10, 10, 10, 10, 65, 70, 75, 80},
// 		{10, 10, 10, 10, 85, 90, 95, 100},
// 	}
// 	// 3倍速
// 	Sawtooth3 = []Pulse{
// 		{10, 10, 10, 10, 7, 14, 22, 29},
// 		{10, 10, 10, 10, 37, 44, 52, 59},
// 		{10, 10, 10, 10, 67, 74, 82, 89},
// 	}
// 	// 4倍速
// 	Sawtooth4 = []Pulse{
// 		{10, 10, 10, 10, 10, 20, 30, 40},
// 		{10, 10, 10, 10, 50, 60, 70, 80},
// 	}
// )
