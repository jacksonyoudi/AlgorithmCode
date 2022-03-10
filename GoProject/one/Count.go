package one
//
//import "time"
//
//type apiData struct {
//	Date     string // 到分钟
//	endPoint string
//}
//
////func init() {
////	var  apiChan = make(chan *apiData, 10000)
////}
//
//type ApiCount struct {
//	dataChan   chan *apiData
//	metircData map[string]int
//	metircData5min map[string]int
//}
//
//func newApiCount(dataChan chan *apiData, metircData map[string]int) *ApiCount {
//	return &ApiCount{
//		dataChan:   dataChan,
//		metircData: metircData,
//	}
//}
//
//
//// 打包
//
//func (count *ApiCount) run() {
//	// 定时器
//	go count.ontimer()
//	// 处理数据
//
//
//	for {
//		select {
//			case data := <- count.dataChan:
//				key := data.endPoint + "-"+ data.Date
//				if ok := count.metircData[key]; ok {
//					count.metircData[key] = count.metircData[key] + 1
//				} else {
//					count.metircData[key] = 1
//				}
//		}
//	}
//}
//
//// 接收数据
//func (count *ApiCount) send(apidata *apiData) {
//	count.dataChan <- apidata
//}
//
//func (count *ApiCount) ontimer() {
//	// 1. 1分钟进行聚合一次数据
//
//	// 2. 清理过期的数据
//
//	timer := time.NewTimer(60)
//
//	for {
//
//		select {
//		case <- timer.C:
//			// 获取到对应 5分钟的数据
//			total := 0
//
//			// 统计5分钟
//			for k, v := range count.metircData {
//				// 切割 endPoint 和date
//				endPoint, date := "", ""
//				// 判断 时间是否在 5分钟之内 如果在 就++
//				total +=1
//				key5 := endPoint + "5key"
//				count.metircData5min[key5] = 1
//
//			}
//			// 清理过期时间
//
//		}
//	}
//
//
//}
