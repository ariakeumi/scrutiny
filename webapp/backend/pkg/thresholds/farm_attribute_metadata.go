package thresholds

// FarmMetadata defines display names, thresholds, and criticality for Seagate FARM
// (Field Accessible Reliability Metrics) attributes. Uses the same metadata struct
// as ATA Device Statistics for consistency.
var FarmMetadata = map[string]AtaDeviceStatisticsMetadata{
	// Page 1: Drive Information
	"farm_poh": {
		DisplayName: "通电时长（FARM）",
		Description: "来自希捷 FARM 日志的防篡改通电时长。固件无法重置。",
		DisplayType: AtaSmartAttributeDisplayTypeRaw,
	},
	"farm_spoh": {
		DisplayName: "主轴通电时长",
		Description: "主轴电机累计通电小时数。",
		DisplayType: AtaSmartAttributeDisplayTypeRaw,
	},
	"farm_head_flight_hours": {
		DisplayName: "磁头飞行时长",
		Description: "读写磁头在盘片上方飞行的累计小时数。",
		DisplayType: AtaSmartAttributeDisplayTypeRaw,
	},
	"farm_head_load_events": {
		DisplayName: "磁头加载事件",
		Ideal:       ObservedThresholdIdealLow,
		Description: "读写磁头加载到盘片上的次数。",
		DisplayType: AtaSmartAttributeDisplayTypeRaw,
	},
	"farm_power_cycle_count": {
		DisplayName: "通断电次数（FARM）",
		Description: "来自 FARM 日志的通断电周期次数。",
		DisplayType: AtaSmartAttributeDisplayTypeRaw,
	},

	// Page 2: Workload Statistics
	"farm_total_read_commands": {
		DisplayName: "读取命令总数",
		Description: "向硬盘发出的读取命令总数。",
		DisplayType: AtaSmartAttributeDisplayTypeRaw,
	},
	"farm_total_write_commands": {
		DisplayName: "写入命令总数",
		Description: "向硬盘发出的写入命令总数。",
		DisplayType: AtaSmartAttributeDisplayTypeRaw,
	},
	"farm_logical_sectors_written": {
		DisplayName: "逻辑扇区写入数",
		Description: "写入硬盘的逻辑扇区总数。",
		DisplayType: AtaSmartAttributeDisplayTypeRaw,
	},
	"farm_logical_sectors_read": {
		DisplayName: "逻辑扇区读取数",
		Description: "从硬盘读取的逻辑扇区总数。",
		DisplayType: AtaSmartAttributeDisplayTypeRaw,
	},

	// Page 3: Error Statistics
	"farm_unrecoverable_read_errors": {
		DisplayName: "不可恢复读取错误数（FARM）",
		Ideal:       ObservedThresholdIdealLow,
		Critical:    true,
		Description: "无法被硬盘纠错机制恢复的读取错误次数。",
		DisplayType: AtaSmartAttributeDisplayTypeRaw,
	},
	"farm_unrecoverable_write_errors": {
		DisplayName: "不可恢复写入错误数（FARM）",
		Ideal:       ObservedThresholdIdealLow,
		Critical:    true,
		Description: "无法被硬盘纠错机制恢复的写入错误次数。",
		DisplayType: AtaSmartAttributeDisplayTypeRaw,
	},
	"farm_reallocated_sectors": {
		DisplayName: "重映射扇区数（FARM）",
		Ideal:       ObservedThresholdIdealLow,
		Critical:    true,
		Description: "来自 FARM 日志的、已重映射到备用扇区的扇区数量。",
		DisplayType: AtaSmartAttributeDisplayTypeRaw,
	},
	"farm_reallocation_candidates": {
		DisplayName: "待重映射扇区数",
		Ideal:       ObservedThresholdIdealLow,
		Critical:    true,
		Description: "等待重映射的扇区数量。这些是硬盘已识别为可能失效的扇区。",
		DisplayType: AtaSmartAttributeDisplayTypeRaw,
	},
	"farm_crc_errors": {
		DisplayName: "CRC 错误数（FARM）",
		Ideal:       ObservedThresholdIdealLow,
		Critical:    true,
		Description: "CRC（循环冗余校验）错误总数，通常表示接口/线缆问题。",
		DisplayType: AtaSmartAttributeDisplayTypeRaw,
	},
	"farm_command_timeouts": {
		DisplayName: "命令超时次数",
		Ideal:       ObservedThresholdIdealLow,
		Description: "超时的命令总数。",
		DisplayType: AtaSmartAttributeDisplayTypeRaw,
	},

	// Page 4: Environment Statistics
	"farm_current_temperature": {
		DisplayName: "温度（FARM）",
		Ideal:       ObservedThresholdIdealLow,
		Description: "来自 FARM 日志的硬盘当前温度（摄氏度）。",
		DisplayType: AtaSmartAttributeDisplayTypeRaw,
	},
	"farm_highest_temperature": {
		DisplayName: "最高温度",
		Ideal:       ObservedThresholdIdealLow,
		Description: "硬盘记录过的最高温度。",
		DisplayType: AtaSmartAttributeDisplayTypeRaw,
	},
	"farm_lowest_temperature": {
		DisplayName: "最低温度",
		Description: "硬盘记录过的最低温度。",
		DisplayType: AtaSmartAttributeDisplayTypeRaw,
	},
}
