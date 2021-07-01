package models

// MetricSumStat 指标统计总和数据
type MetricSumStat struct {
	Id       uint64  `field:"id"`       // ID
	ServerId uint32  `field:"serverId"` // 服务ID
	ItemId   uint64  `field:"itemId"`   // 指标
	Count    uint64  `field:"count"`    // 数量
	Total    float64 `field:"total"`    // 总和
	Time     string  `field:"time"`     // 分钟值YYYYMMDDHHII
	Version  uint32  `field:"version"`  // 版本号
}

type MetricSumStatOperator struct {
	Id       interface{} // ID
	ServerId interface{} // 服务ID
	ItemId   interface{} // 指标
	Count    interface{} // 数量
	Total    interface{} // 总和
	Time     interface{} // 分钟值YYYYMMDDHHII
	Version  interface{} // 版本号
}

func NewMetricSumStatOperator() *MetricSumStatOperator {
	return &MetricSumStatOperator{}
}
