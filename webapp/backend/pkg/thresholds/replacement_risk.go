package thresholds

// ReplacementRiskWeight defines how much a single SMART attribute can contribute
// to the overall replacement risk score. The weight is a fraction of 100; the
// scorer multiplies it by the attribute's normalized severity (0.0-1.0) to get
// the points added to the total score.
type ReplacementRiskWeight struct {
	// AttributeID is the protocol-specific attribute identifier (int key for ATA,
	// string key for NVMe and SCSI).
	AttributeID string

	// DisplayName is the human-readable attribute name.
	DisplayName string

	// Weight is the maximum number of points (out of 100) this attribute can
	// contribute. All weights across a drive type should sum to 100.
	Weight float64

	// TrendMultiplier scales the additional score added when the attribute value
	// is worsening. A value of 1.0 means the trend can double the attribute's
	// contribution at maximum deterioration rate. Set to 0 for attributes where
	// trend analysis is not meaningful.
	TrendMultiplier float64
}

// AtaReplacementRiskWeights defines attribute weights for ATA/SATA hard drives
// and SSDs. Weights are informed by Backblaze failure correlation research
// (https://www.backblaze.com/cloud-storage/resources/hard-drive-test-data).
//
// Attributes with direct failure correlation receive higher weights. Attributes
// that are noisy or vendor-specific (Raw Read Error Rate, Seek Error Rate)
// receive lower weights and rely more on trend analysis.
var AtaReplacementRiskWeights = []ReplacementRiskWeight{
	{
		AttributeID:     "5",
		DisplayName:     "重映射扇区数",
		Weight:          25,
		TrendMultiplier: 1.5,
	},
	{
		AttributeID:     "197",
		DisplayName:     "当前待映射扇区数",
		Weight:          20,
		TrendMultiplier: 1.5,
	},
	{
		AttributeID:     "198",
		DisplayName:     "离线不可纠正扇区",
		Weight:          20,
		TrendMultiplier: 1.2,
	},
	{
		AttributeID:     "196",
		DisplayName:     "重映射事件计数",
		Weight:          10,
		TrendMultiplier: 1.0,
	},
	{
		AttributeID:     "10",
		DisplayName:     "起转重试次数",
		Weight:          10,
		TrendMultiplier: 1.0,
	},
	{
		AttributeID:     "199",
		DisplayName:     "UDMA CRC 错误计数",
		Weight:          5,
		TrendMultiplier: 0.5,
	},
	{
		AttributeID:     "1",
		DisplayName:     "原始读取错误率",
		Weight:          5,
		TrendMultiplier: 0.8,
	},
	{
		AttributeID:     "7",
		DisplayName:     "寻道错误率",
		Weight:          3,
		TrendMultiplier: 0.8,
	},
	{
		AttributeID:     "194",
		DisplayName:     "温度",
		Weight:          2,
		TrendMultiplier: 0.3,
	},
}

// NvmeReplacementRiskWeights defines attribute weights for NVMe drives.
// NVMe drives expose direct wear indicators (Percentage Used, Available Spare)
// making prediction more straightforward than ATA.
var NvmeReplacementRiskWeights = []ReplacementRiskWeight{
	{
		AttributeID:     "percentage_used",
		DisplayName:     "已使用百分比",
		Weight:          40,
		TrendMultiplier: 1.0,
	},
	{
		AttributeID:     "available_spare",
		DisplayName:     "可用备用容量",
		Weight:          30,
		TrendMultiplier: 1.2,
	},
	{
		AttributeID:     "media_errors",
		DisplayName:     "介质与数据完整性错误",
		Weight:          20,
		TrendMultiplier: 1.5,
	},
	{
		AttributeID:     "critical_warning",
		DisplayName:     "严重警告",
		Weight:          7,
		TrendMultiplier: 0,
	},
	{
		AttributeID:     "unsafe_shutdowns",
		DisplayName:     "不安全关机次数",
		Weight:          3,
		TrendMultiplier: 0.5,
	},
}

// ScsiReplacementRiskWeights defines attribute weights for SCSI/SAS drives.
var ScsiReplacementRiskWeights = []ReplacementRiskWeight{
	{
		AttributeID:     "scsi_grown_defect_list",
		DisplayName:     "增长缺陷列表",
		Weight:          40,
		TrendMultiplier: 1.5,
	},
	{
		AttributeID:     "read_total_uncorrected_errors",
		DisplayName:     "读取不可纠正错误总数",
		Weight:          25,
		TrendMultiplier: 1.5,
	},
	{
		AttributeID:     "write_total_uncorrected_errors",
		DisplayName:     "写入不可纠正错误总数",
		Weight:          25,
		TrendMultiplier: 1.5,
	},
	{
		AttributeID:     "read_errors_corrected_by_eccdelayed",
		DisplayName:     "ECC 延迟纠正的读取错误",
		Weight:          5,
		TrendMultiplier: 1.0,
	},
	{
		AttributeID:     "write_errors_corrected_by_eccdelayed",
		DisplayName:     "ECC 延迟纠正的写入错误",
		Weight:          5,
		TrendMultiplier: 1.0,
	},
}

// ReplacementRiskWeightsForProtocol returns the correct weight table for the
// given device protocol string ("ATA", "NVMe", "SCSI"). Returns nil for
// unrecognized protocols.
func ReplacementRiskWeightsForProtocol(protocol string) []ReplacementRiskWeight {
	switch protocol {
	case "ATA":
		return AtaReplacementRiskWeights
	case "NVMe":
		return NvmeReplacementRiskWeights
	case "SCSI":
		return ScsiReplacementRiskWeights
	default:
		return nil
	}
}
