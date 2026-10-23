package thresholds

// https://media.kingston.com/support/downloads/MKP_521.6_SMART-DCP1000_attribute.pdf
// https://www.percona.com/blog/2017/02/09/using-nvme-command-line-tools-to-check-nvme-flash-health/
// https://nvmexpress.org/resources/nvm-express-technology-features/nvme-features-for-error-reporting-smart-log-pages-failures-and-management-capabilities-in-nvme-architectures/
// https://www.micromat.com/product_manuals/drive_scope_manual_01.pdf
type NvmeAttributeMetadata struct {
	ID          string `json:"-"`
	DisplayName string `json:"display_name"`
	Ideal       string `json:"ideal"`
	Critical    bool   `json:"critical"`
	Description string `json:"description"`

	Transform          func(int64, int64, string) int64 `json:"-"` //this should be a method to extract/tranform the normalized or raw data to a chartable format. Str
	TransformValueUnit string                           `json:"transform_value_unit,omitempty"`
	DisplayType        string                           `json:"display_type"` //"raw" "normalized" or "transformed"
}

var NmveMetadata = map[string]NvmeAttributeMetadata{
	"critical_warning": {
		ID:          "critical_warning",
		DisplayName: "严重警告",
		DisplayType: "",
		Ideal:       "low",
		Critical:    true,
		Description: "该字段指示控制器状态的严重警告。每一位对应一种严重警告类型，可同时置多位。某位清零（‘0’）表示该严重警告不适用。严重警告可能触发向主机的异步事件通知。本字段中的位表示当前的关联状态，不会持久保存。",
	},
	"temperature": {
		ID:          "temperature",
		DisplayName: "温度",
		DisplayType: "",
		Ideal:       "",
		Critical:    false,
		Description: "",
	},
	"available_spare": {
		ID:          "available_spare",
		DisplayName: "可用备用容量",
		DisplayType: "",
		Ideal:       "high",
		Critical:    true,
		Description: "包含剩余可用备用容量的标准化百分比（0 到 100%）。",
	},
	"percentage_used": {
		ID:          "percentage_used",
		DisplayName: "已使用百分比",
		DisplayType: "",
		Ideal:       "low",
		Critical:    true,
		Description: "基于实际使用情况和厂商对 NVM 寿命的预测，给出的 NVM 子系统寿命已用百分比的厂商估算值。值为 100 表示 NVM 子系统的预计耐久度已耗尽，但不一定代表 NVM 子系统故障。该值允许超过 100，大于 254 的百分比应表示为 255。该值每小时上电更新一次（控制器不处于睡眠状态时）。",
	},
	"data_units_read": {
		ID:          "data_units_read",
		DisplayName: "读取数据单元数",
		DisplayType: "",
		Ideal:       "",
		Critical:    false,
		Description: "主机从控制器读取的 512 字节数据单元数量，不含元数据。该值以千为单位报告（即值为 1 对应读取 1000 个 512 字节单元），向上取整。当 LBA 大小不是 512 字节时，控制器应将读取数据量换算为 512 字节单元。",
	},
	"data_units_written": {
		ID:          "data_units_written",
		DisplayName: "写入数据单元数",
		DisplayType: "",
		Ideal:       "",
		Critical:    false,
		Description: "主机写入控制器的 512 字节数据单元数量，不含元数据。该值以千为单位报告（即值为 1 对应写入 1000 个 512 字节单元），向上取整。当 LBA 大小不是 512 字节时，控制器应将写入数据量换算为 512 字节单元。",
	},
	"host_reads": {
		ID:          "host_reads",
		DisplayName: "主机读取次数",
		DisplayType: "",
		Ideal:       "",
		Critical:    false,
		Description: "控制器已完成的读命令数量",
	},
	"host_writes": {
		ID:          "host_writes",
		DisplayName: "主机写入次数",
		DisplayType: "",
		Ideal:       "",
		Critical:    false,
		Description: "控制器已完成的写命令数量",
	},
	"controller_busy_time": {
		ID:          "controller_busy_time",
		DisplayName: "控制器繁忙时间",
		DisplayType: "",
		Ideal:       "",
		Critical:    false,
		Description: "控制器处理 I/O 命令的繁忙时间。当 I/O 队列存在未完成命令时（即命令已通过 I/O 提交队列尾门铃写入发出，但对应的完成队列条目尚未发布到关联的 I/O 完成队列），控制器处于繁忙状态。该值以分钟为单位报告。",
	},
	"power_cycles": {
		ID:          "power_cycles",
		DisplayName: "通断电次数",
		DisplayType: "",
		Ideal:       "",
		Critical:    false,
		Description: "通断电（上电周期）次数。",
	},
	"power_on_hours": {
		ID:          "power_on_hours",
		DisplayName: "通电时长",
		DisplayType: "",
		Ideal:       "",
		Critical:    false,
		Description: "通电小时数。通电时长始终在记录，即使处于低功耗模式。",
	},
	"unsafe_shutdowns": {
		ID:          "unsafe_shutdowns",
		DisplayName: "不安全关机次数",
		DisplayType: "",
		Ideal:       "",
		Critical:    false,
		Description: "不安全关机次数。在断电前未收到关机通知（CC.SHN）时该计数递增。",
	},
	"media_errors": {
		ID:          "media_errors",
		DisplayName: "介质错误数",
		DisplayType: "",
		Ideal:       "low",
		Critical:    true,
		Description: "控制器检测到未恢复的数据完整性错误的次数。不可纠正的 ECC 错误、CRC 校验失败或 LBA 标签不匹配等错误都计入本字段。",
	},
	"num_err_log_entries": {
		ID:          "num_err_log_entries",
		DisplayName: "错误日志条目数",
		DisplayType: "",
		Ideal:       "low",
		Critical:    true,
		Description: "控制器生命周期内错误信息日志条目的数量。",
	},
	"warning_temp_time": {
		ID:          "warning_temp_time",
		DisplayName: "警告温度持续时长",
		DisplayType: "",
		Ideal:       "",
		Critical:    false,
		Description: "控制器处于工作状态、且综合温度大于等于警告综合温度阈值（WCTEMP）并小于严重综合温度阈值（CCTEMP）（Identify Controller 数据结构中的字段）的累计分钟数。",
	},
	"critical_comp_time": {
		ID:          "critical_comp_time",
		DisplayName: "严重温度持续时长",
		DisplayType: "",
		Ideal:       "",
		Critical:    false,
		Description: "控制器处于工作状态、且综合温度大于 Identify Controller 数据结构中严重综合温度阈值（CCTEMP）字段的累计分钟数。",
	},
}
