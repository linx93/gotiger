package tools

import (
	gouuid "github.com/satori/go.uuid"
	"github.com/yitter/idgenerator-go/idgen"
	"strconv"
)

const uniqueWorkerId = 0

func init() {
	initSnowFlake()
}

// 关于雪花算法的初始化
func initSnowFlake() {
	// 创建 IdGeneratorOptions 对象，可在构造函数中输入 WorkerId：
	var options = idgen.NewIdGeneratorOptions(uniqueWorkerId)
	// options.WorkerIdBitLength = 10  // 默认值6，限定 WorkerId 最大值为2^6-1，即默认最多支持64个节点。
	// options.SeqBitLength = 6; // 默认值6，限制每毫秒生成的ID个数。若生成速度超过5万个/秒，建议加大 SeqBitLength 到 10。
	// options.BaseTime = Your_Base_Time // 如果要兼容老系统的雪花算法，此处应设置为老系统的BaseTime。
	// ...... 其它参数参考 IdGeneratorOptions 定义。

	// 保存参数（务必调用，否则参数设置不生效）：
	idgen.SetIdGenerator(options)
}

// IdGenerator id生成器
type IdGenerator interface {
	GetStrId() string
	GetNumberId() int64
}

type UUID struct{}

func (uu *UUID) GetStrId() string {
	return gouuid.NewV4().String()
}

// GetNumberId 不支持数字uuid
func (uu *UUID) GetNumberId() int64 {
	panic("不支持数字uuid")
	return -1
}

type SnowFlake struct{}

func (s *SnowFlake) GetStrId() string {
	return strconv.Itoa(int(idgen.NextId()))
}

func (s *SnowFlake) GetNumberId() int64 {
	return idgen.NextId()
}
