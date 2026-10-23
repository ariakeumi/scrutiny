package thresholds

type ScsiAttributeMetadata struct {
	ID          string `json:"-"`
	DisplayName string `json:"display_name"`
	Ideal       string `json:"ideal"`
	Critical    bool   `json:"critical"`
	Description string `json:"description"`

	Transform          func(int64, int64, string) int64 `json:"-"` //this should be a method to extract/tranform the normalized or raw data to a chartable format. Str
	TransformValueUnit string                           `json:"transform_value_unit,omitempty"`
	DisplayType        string                           `json:"display_type"` //"raw" "normalized" or "transformed"
}

var ScsiMetadata = map[string]ScsiAttributeMetadata{
	"scsi_grown_defect_list": {
		ID:          "scsi_grown_defect_list",
		DisplayName: "增长缺陷列表",
		DisplayType: "",
		Ideal:       "low",
		Critical:    true,
		Description: "增长缺陷计数表示硬盘出厂以来被替换（缺陷）块的数量。每新增一个缺陷块，计数加一。 Each additional defective block increases the count by one.",
	},
	"read_errors_corrected_by_eccfast": {
		ID:          "read_errors_corrected_by_eccfast",
		DisplayName: "ECC 快速纠正的读取错误",
		DisplayType: "",
		Ideal:       "",
		Critical:    false,
		Description: "为获得正确数据而执行的错误纠正（即“即时 ECC”）。“无实质延迟”指纠正未推迟后续扇区的读取（例如没有损失一圈旋转）。每需要一个纠正的逻辑块计数加一；同一条命令中纠正的两个不同块计为两次事件。",
	},
	"read_errors_corrected_by_eccdelayed": {
		ID:          "read_errors_corrected_by_eccdelayed",
		DisplayName: "ECC 延迟纠正的读取错误",
		DisplayType: "",
		Ideal:       "",
		Critical:    false,
		Description: "为获得正确数据而施加错误码或算法（如 ECC、校验和），且伴随实质延迟。“可能有延迟”指纠正耗时超过一个扇区时间，导致后续扇区的读写被推迟（例如损失一圈旋转）。每需要一个纠正的逻辑块计数加一；一个可纠正的双 bit 错误块计为一次事件，同一条命令中纠正的两个不同块计为两次事件。",
	},
	"read_errors_corrected_by_rereads_rewrites": {
		ID:          "read_errors_corrected_by_rereads_rewrites",
		DisplayName: "重读/重写纠正的读取错误",
		DisplayType: "",
		Ideal:       "low",
		Critical:    true,
		Description: "该参数码指定的计数器统计通过重试纠正的错误数量。它统计的是恢复的错误数，而不是重试次数：如果恢复一个数据块需要五次重试，计数器加一而不是加五。每个通过重试恢复的逻辑块计数加一。如果重试无法恢复而最终由 ECC 恢复，则不计入本计数器，而计入参数码 01h（可能有延迟的纠错）对应的计数器。",
	},
	"read_total_errors_corrected": {
		ID:          "read_total_errors_corrected",
		DisplayName: "读取错误纠正总数",
		DisplayType: "",
		Ideal:       "",
		Critical:    false,
		Description: "该计数器统计参数码 00h、01h 和 02h 错误的总数（即 ECC 纠正的错误：快速与延迟，加上重读/重写纠正的错误）。这三个计数器之间不会“重复计数”。所有可纠正错误之和应通过参数码 01h 与 02h 相加得出，而不是使用本总数。",
	},
	"read_correction_algorithm_invocations": {
		ID:          "read_correction_algorithm_invocations",
		DisplayName: "读取纠错算法调用次数",
		DisplayType: "",
		Ideal:       "",
		Critical:    false,
		Description: "该参数码指定的计数器统计重试总数，即“纠错算法被调用的次数”。如果计数器 02h 类错误经五次尝试后恢复，则本计数器加五。如果纠正计数器 01h 类错误前需要三次重试以获得稳定的 ECC 校验子，这三次重试也计入本计数器。未能成功恢复错误（计数器 06h 类错误）所进行的重试次数同样计入本计数器。",
	},
	"read_total_uncorrected_errors": {
		ID:          "read_total_uncorrected_errors",
		DisplayName: "读取不可纠正错误总数",
		DisplayType: "",
		Ideal:       "low",
		Critical:    true,
		Description: "该参数码指定的计数器包含发生未纠正数据错误的逻辑块总数。",
	},
	"write_errors_corrected_by_eccfast": {
		ID:          "write_errors_corrected_by_eccfast",
		DisplayName: "ECC 快速纠正的写入错误",
		DisplayType: "",
		Ideal:       "",
		Critical:    false,
		Description: "为获得正确数据而执行的错误纠正（即“即时 ECC”）。“无实质延迟”指纠正未推迟后续扇区的读取（例如没有损失一圈旋转）。每需要一个纠正的逻辑块计数加一；同一条命令中纠正的两个不同块计为两次事件。 ",
	},
	"write_errors_corrected_by_eccdelayed": {
		ID:          "write_errors_corrected_by_eccdelayed",
		DisplayName: "ECC 延迟纠正的写入错误",
		DisplayType: "",
		Ideal:       "",
		Critical:    false,
		Description: "为获得正确数据而施加错误码或算法（如 ECC、校验和），且伴随实质延迟。“可能有延迟”指纠正耗时超过一个扇区时间，导致后续扇区的读写被推迟（例如损失一圈旋转）。每需要一个纠正的逻辑块计数加一；一个可纠正的双 bit 错误块计为一次事件，同一条命令中纠正的两个不同块计为两次事件。",
	},
	"write_errors_corrected_by_rereads_rewrites": {
		ID:          "write_errors_corrected_by_rereads_rewrites",
		DisplayName: "重读/重写纠正的写入错误",
		DisplayType: "",
		Ideal:       "low",
		Critical:    true,
		Description: "该参数码指定的计数器统计通过重试纠正的错误数量。它统计的是恢复的错误数，而不是重试次数：如果恢复一个数据块需要五次重试，计数器加一而不是加五。每个通过重试恢复的逻辑块计数加一。如果重试无法恢复而最终由 ECC 恢复，则不计入本计数器，而计入参数码 01h（可能有延迟的纠错）对应的计数器。",
	},
	"write_total_errors_corrected": {
		ID:          "write_total_errors_corrected",
		DisplayName: "写入错误纠正总数",
		DisplayType: "",
		Ideal:       "",
		Critical:    false,
		Description: "该计数器统计参数码 00h、01h 和 02h 错误的总数（即 ECC 纠正的错误：快速与延迟，加上重读/重写纠正的错误）。这三个计数器之间不会“重复计数”。所有可纠正错误之和应通过参数码 01h 与 02h 相加得出，而不是使用本总数。",
	},
	"write_correction_algorithm_invocations": {
		ID:          "write_correction_algorithm_invocations",
		DisplayName: "写入纠错算法调用次数",
		DisplayType: "",
		Ideal:       "",
		Critical:    false,
		Description: "该参数码指定的计数器统计重试总数，即“纠错算法被调用的次数”。如果计数器 02h 类错误经五次尝试后恢复，则本计数器加五。如果纠正计数器 01h 类错误前需要三次重试以获得稳定的 ECC 校验子，这三次重试也计入本计数器。未能成功恢复错误（计数器 06h 类错误）所进行的重试次数同样计入本计数器。",
	},
	"write_total_uncorrected_errors": {
		ID:          "write_total_uncorrected_errors",
		DisplayName: "写入不可纠正错误总数",
		DisplayType: "",
		Ideal:       "low",
		Critical:    true,
		Description: "该参数码指定的计数器包含发生未纠正数据错误的逻辑块总数。",
	},
	"temperature": {
		ID:          "temperature",
		DisplayName: "温度",
		DisplayType: "",
		Ideal:       "",
		Critical:    false,
		Description: "硬盘当前温度（摄氏度）。",
	},
	"read_gigabytes_processed": {
		ID:          "read_gigabytes_processed",
		DisplayName: "累计读取数据量",
		DisplayType: "",
		Ideal:       "",
		Critical:    false,
		Description: "读取命令累计处理的数据量，取自 SCSI 读取错误计数日志页的“gigabytes processed”字段。相当于 ATA 的 LBA 读取总数或 NVMe 的读取数据单元数，是 SAS 硬盘工作负载读取速率估算的依据。注意：统计的是全部处理过的数据（含重试），不只是主机的唯一数据，且以约 1 GB 粒度报告。",
	},
	"write_gigabytes_processed": {
		ID:          "write_gigabytes_processed",
		DisplayName: "累计写入数据量",
		DisplayType: "",
		Ideal:       "",
		Critical:    false,
		Description: "写入命令累计处理的数据量，取自 SCSI 写入错误计数日志页的“gigabytes processed”字段。相当于 ATA 的 LBA 写入总数或 NVMe 的写入数据单元数，是 SAS 硬盘工作负载写入速率与耐久度（TBW）估算的依据。注意：统计的是全部处理过的数据（含重试），不只是主机的唯一数据，且以约 1 GB 粒度报告。",
	},
	"percentage_used": {
		ID:          "percentage_used",
		DisplayName: "已使用百分比",
		DisplayType: "",
		Ideal:       "low",
		Critical:    true,
		Description: "来自固态介质日志页“Percentage used endurance indicator”的 SAS SSD 磨损指标。达到厂商额定耐久度上限时为 100%，但硬盘在该点之后仍可继续工作。仅固态 SAS 硬盘报告。",
	},
}
