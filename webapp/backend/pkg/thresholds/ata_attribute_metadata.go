package thresholds

import (
	"strconv"
	"strings"
)

const AtaSmartAttributeDisplayTypeRaw = "raw"
const AtaSmartAttributeDisplayTypeNormalized = "normalized"
const AtaSmartAttributeDisplayTypeTransformed = "transformed"

type AtaAttributeMetadata struct {
	ID          int64  `json:"-"`
	DisplayName string `json:"display_name"`
	Ideal       string `json:"ideal"`
	Critical    bool   `json:"critical"`
	Description string `json:"description"`

	// UseDeltaEvaluation indicates this attribute is a cumulative counter where
	// the trend matters more than the absolute value. When true, Scrutiny will
	// suppress warnings if the value hasn't increased since the last measurement.
	// Useful for attributes like UltraDMA CRC Error Count (199) where errors
	// accumulate from transient issues (e.g., loose cables) but the count persists
	// even after the root cause is fixed.
	UseDeltaEvaluation bool `json:"use_delta_evaluation,omitempty"`

	Transform          func(int64, int64, string) int64 `json:"-"` //this should be a method to extract/tranform the normalized or raw data to a chartable format. Str
	TransformValueUnit string                           `json:"transform_value_unit,omitempty"`
	ObservedThresholds []ObservedThreshold              `json:"observed_thresholds,omitempty"` //these thresholds must match the DisplayType
	DisplayType        string                           `json:"display_type"`                  //"raw" "normalized" or "transformed"
}

const ObservedThresholdIdealLow = "low"
const ObservedThresholdIdealHigh = "high"

type ObservedThreshold struct {
	Low  int64 `json:"low"`  //threshold (row/normalized data) boundary low value
	High int64 `json:"high"` //threshold (row/normalized data) boundary high value

	AnnualFailureRate float64   `json:"annual_failure_rate"` //error rate %
	ErrorInterval     []float64 `json:"error_interval"`
}

var AtaMetadata = map[int]AtaAttributeMetadata{
	1: {
		ID:          1,
		DisplayName: "读取错误率",
		DisplayType: AtaSmartAttributeDisplayTypeNormalized,
		Ideal:       ObservedThresholdIdealLow,
		Critical:    false,
		Description: "（厂商特定原始值。）存储从盘片表面读取数据时发生的硬件读取错误率相关数据。原始值的结构因厂商而异，作为十进制数通常没有意义。",
	},
	2: {
		ID:          2,
		DisplayName: "吞吐量性能",
		DisplayType: AtaSmartAttributeDisplayTypeNormalized,
		Ideal:       ObservedThresholdIdealHigh,
		Critical:    false,
		Description: "硬盘的整体（一般）吞吐性能。如果该属性值持续下降，磁盘出问题的概率很高。",
	},
	3: {
		ID:          3,
		DisplayName: "主轴起转时间",
		DisplayType: AtaSmartAttributeDisplayTypeNormalized,
		Ideal:       ObservedThresholdIdealLow,
		Critical:    false,
		Description: "主轴起转的平均时间（从 0 RPM 到完全运转，单位毫秒）。",
		ObservedThresholds: []ObservedThreshold{
			{
				Low:               78,
				High:              96,
				AnnualFailureRate: 0.11452195377351217,
				ErrorInterval:     []float64{0.10591837762295722, 0.12363823501915781},
			},
			{
				Low:               96,
				High:              114,
				AnnualFailureRate: 0.040274562840558074,
				ErrorInterval:     []float64{0.03465055611002801, 0.046551312468303144},
			},
			{
				Low:               114,
				High:              132,
				AnnualFailureRate: 0.009100406705780476,
				ErrorInterval:     []float64{0.006530608971356785, 0.012345729280075591},
			},
			{
				Low:               132,
				High:              150,
				AnnualFailureRate: 0.008561351734020232,
				ErrorInterval:     []float64{0.004273795939256936, 0.015318623141355509},
			},
			{
				Low:               150,
				High:              168,
				AnnualFailureRate: 0.015780508262068848,
				ErrorInterval:     []float64{0.005123888078524015, 0.03682644215646287},
			},
			{
				Low:               168,
				High:              186,
				AnnualFailureRate: 0.05262688124794024,
				ErrorInterval:     []float64{0.0325768689524594, 0.08044577830285578},
			},
			{
				Low:               186,
				High:              204,
				AnnualFailureRate: 0.01957419424036038,
				ErrorInterval:     []float64{0.0023705257325185624, 0.0707087198669825},
			},
			{
				Low:               204,
				High:              222,
				AnnualFailureRate: 0.026050959960031404,
				ErrorInterval:     []float64{0.0006595532020744994, 0.1451466588889228},
			},
		},
	},
	4: {
		ID:          4,
		DisplayName: "启停次数",
		DisplayType: AtaSmartAttributeDisplayTypeRaw,
		Ideal:       "",
		Critical:    false,
		Description: "主轴启停周期的累计次数。硬盘在完全断电（切断电源）后再次上电，以及从休眠模式中唤醒时，主轴都会重新启动，计数随之增加。",
		ObservedThresholds: []ObservedThreshold{
			{
				Low:               0,
				High:              13,
				AnnualFailureRate: 0.01989335424860646,
				ErrorInterval:     []float64{0.016596548909440657, 0.023653263230617408},
			},
			{
				Low:               13,
				High:              26,
				AnnualFailureRate: 0.03776935438256488,
				ErrorInterval:     []float64{0.03310396052098642, 0.04290806173460437},
			},
			{
				Low:               26,
				High:              39,
				AnnualFailureRate: 0.11022223828187004,
				ErrorInterval:     []float64{0.09655110535164119, 0.12528657238811672},
			},
			{
				Low:               39,
				High:              52,
				AnnualFailureRate: 0.16289995457762474,
				ErrorInterval:     []float64{0.13926541653588131, 0.18939614504497515},
			},
			{
				Low:               52,
				High:              65,
				AnnualFailureRate: 0.19358212432279714,
				ErrorInterval:     []float64{0.15864522253849073, 0.23392418181765526},
			},
			{
				Low:               65,
				High:              78,
				AnnualFailureRate: 0.1157094940074447,
				ErrorInterval:     []float64{0.07861898732346269, 0.16424039052527728},
			},
			{
				Low:               78,
				High:              91,
				AnnualFailureRate: 0.12262136155304391,
				ErrorInterval:     []float64{0.0670382394080032, 0.20573780888032978},
			},
			{
				Low:               91,
				High:              104,
				AnnualFailureRate: 0,
				ErrorInterval:     []float64{0, 0},
			},
		},
	},
	5: {
		ID:          5,
		DisplayName: "重映射扇区数",
		DisplayType: AtaSmartAttributeDisplayTypeRaw,
		Ideal:       ObservedThresholdIdealLow,
		Critical:    true,
		Description: "已重映射扇区的计数。原始值表示已发现并重映射的坏扇区数量。因此属性值越高，硬盘被迫重映射的扇区越多。该值主要用于衡量硬盘的预期寿命；出现过任何重映射的硬盘，在随后几个月内发生故障的概率会显著升高。",
		ObservedThresholds: []ObservedThreshold{
			{
				Low:               0,
				High:              0,
				AnnualFailureRate: 0.025169175350572493,
				ErrorInterval:     []float64{0.022768612038746357, 0.027753988579272894},
			},
			{
				Low:               1,
				High:              4,
				AnnualFailureRate: 0.027432608477803388,
				ErrorInterval:     []float64{0.010067283827589948, 0.05970923963096652},
			},
			{
				Low:               4,
				High:              16,
				AnnualFailureRate: 0.07501976284584981,
				ErrorInterval:     []float64{0.039944864177334186, 0.12828607921150972},
			},
			{
				Low:               16,
				High:              70,
				AnnualFailureRate: 0.23589260654405794,
				ErrorInterval:     []float64{0.1643078435800227, 0.32806951196017664},
			},
			{
				Low:               70,
				High:              260,
				AnnualFailureRate: 0.36193219378600433,
				ErrorInterval:     []float64{0.2608488901774093, 0.4892271827875412},
			},
			{
				Low:               260,
				High:              1100,
				AnnualFailureRate: 0.5676621428968173,
				ErrorInterval:     []float64{0.4527895568499355, 0.702804359408436},
			},
			{
				Low:               1100,
				High:              4500,
				AnnualFailureRate: 1.5028253400346423,
				ErrorInterval:     []float64{1.2681757596263297, 1.768305221795894},
			},
			{
				Low:               4500,
				High:              17000,
				AnnualFailureRate: 2.0659987547404763,
				ErrorInterval:     []float64{1.6809790460512237, 2.512808045182302},
			},
			{
				Low:               17000,
				High:              70000,
				AnnualFailureRate: 1.7755385684503124,
				ErrorInterval:     []float64{1.2796520259849835, 2.400012341226441},
			},
		},
	},
	6: {
		ID:          6,
		DisplayName: "读取通道裕量",
		DisplayType: AtaSmartAttributeDisplayTypeNormalized,
		Ideal:       "",
		Critical:    false,
		Description: "读取数据时通道的裕量。该属性的具体功能未作规范定义。",
	},
	7: {
		ID:          7,
		DisplayName: "寻道错误率",
		DisplayType: AtaSmartAttributeDisplayTypeNormalized,
		Ideal:       "",
		Critical:    false,
		Description: "（厂商特定原始值。）磁头寻道错误率。机械定位系统发生局部故障时会产生寻道错误，原因可能是伺服损坏或硬盘热胀等多种因素。原始值的结构因厂商而异，作为十进制数通常没有意义。",
	},
	8: {
		ID:          8,
		DisplayName: "寻道时间性能",
		DisplayType: AtaSmartAttributeDisplayTypeNormalized,
		Ideal:       ObservedThresholdIdealHigh,
		Critical:    false,
		Description: "磁头寻道操作的平均性能。该属性下降是机械子系统出问题的征兆。",
	},
	9: {

		ID:          9,
		DisplayName: "通电时长",
		DisplayType: AtaSmartAttributeDisplayTypeTransformed,
		Ideal:       "",
		Critical:    false,
		Description: "通电状态的小时数。原始值显示通电状态的总小时数（或分钟数、秒数，取决于厂商）。按默认定义，状态完好的硬盘预期总寿命为 5 年（每天 24 小时不间断运行），即 24/7 模式下的 1825 天或 43800 小时。部分 2005 年以前的硬盘，该原始值可能前进不规则和/或发生“回绕”（周期性归零）。",
		Transform: func(normValue int64, rawValue int64, rawString string) int64 {
			// smartctl "h+m+s" format (e.g., "1730h+05m+02.453s")
			if strings.Contains(rawString, "h+") {
				hIdx := strings.Index(rawString, "h+")
				if hIdx > 0 {
					hours, err := strconv.ParseInt(rawString[:hIdx], 10, 64)
					if err == nil {
						return hours
					}
				}
			}

			// smartctl parsed format with parenthetical hours
			// e.g., "103800 (1730 hours)" when smartctl detects minutes encoding
			if strings.Contains(rawString, "hours)") {
				parenStart := strings.Index(rawString, "(")
				if parenStart >= 0 {
					inner := rawString[parenStart+1:]
					spaceIdx := strings.Index(inner, " ")
					if spaceIdx > 0 {
						hours, err := strconv.ParseInt(inner[:spaceIdx], 10, 64)
						if err == nil {
							return hours
						}
					}
				}
			}

			// Packed 48-bit raw value with extra data in upper bytes.
			// Some manufacturers store additional data (e.g., temperature, flags) in
			// the upper bytes. The actual power-on hours are in the lower 32 bits.
			if rawValue > 0xFFFFFFFF {
				return rawValue & 0xFFFFFFFF
			}

			// Standard: raw value is already in hours
			return rawValue
		},
		TransformValueUnit: "hours",
	},
	10: {
		ID:          10,
		DisplayName: "起转重试次数",
		DisplayType: AtaSmartAttributeDisplayTypeRaw,
		Ideal:       ObservedThresholdIdealLow,
		Critical:    true,
		Description: "起转重试次数。该属性记录主轴第一次启动未成功后，为达到完全运转速度而重试的总次数。该属性值增加是硬盘机械子系统出问题的征兆。",
		ObservedThresholds: []ObservedThreshold{
			{
				Low:               0,
				High:              0,
				AnnualFailureRate: 0.05459827163896099,
				ErrorInterval:     []float64{0.05113785787727033, 0.05823122757702782},
			},
			{ //TODO: using fake data from attribute 11. Not enough data, but critical and correlated with failure.
				Low:               0,
				High:              80,
				AnnualFailureRate: 0.5555555555555556,
				ErrorInterval:     []float64{0.014065448880161053, 3.095357439410498},
			},
		},
	},
	11: {
		ID:          11,
		DisplayName: "重新校准重试次数",
		DisplayType: AtaSmartAttributeDisplayTypeRaw,
		Ideal:       ObservedThresholdIdealLow,
		Critical:    false,
		Description: "该属性表示请求重新校准的次数（在第一次尝试未成功的情况下）。该属性值增加是硬盘机械子系统出问题的征兆。",
		ObservedThresholds: []ObservedThreshold{
			{
				Low:               0,
				High:              0,
				AnnualFailureRate: 0.04658866433672694,
				ErrorInterval:     []float64{0.03357701137320878, 0.06297433993055492},
			},
			{
				Low:               0,
				High:              80,
				AnnualFailureRate: 0.5555555555555556,
				ErrorInterval:     []float64{0.014065448880161053, 3.095357439410498},
			},
			{
				Low:               80,
				High:              160,
				AnnualFailureRate: 0,
				ErrorInterval:     []float64{0, 0},
			},
			{
				Low:               160,
				High:              240,
				AnnualFailureRate: 0,
				ErrorInterval:     []float64{0, 0},
			},
			{
				Low:               240,
				High:              320,
				AnnualFailureRate: 0,
				ErrorInterval:     []float64{0, 0},
			},
			{
				Low:               320,
				High:              400,
				AnnualFailureRate: 0,
				ErrorInterval:     []float64{0, 0},
			},
			{
				Low:               400,
				High:              480,
				AnnualFailureRate: 0,
				ErrorInterval:     []float64{0, 0},
			},
			{
				Low:               480,
				High:              560,
				AnnualFailureRate: 0,
				ErrorInterval:     []float64{0, 0},
			},
		},
	},
	12: {
		ID:          12,
		DisplayName: "通断电次数",
		DisplayType: AtaSmartAttributeDisplayTypeRaw,
		Ideal:       ObservedThresholdIdealLow,
		Critical:    false,
		Description: "该属性表示硬盘完整通断电周期的次数。",
	},
	13: {
		ID:          13,
		DisplayName: "软读取错误率",
		DisplayType: AtaSmartAttributeDisplayTypeNormalized,
		Ideal:       ObservedThresholdIdealLow,
		Critical:    false,
		Description: "上报给操作系统的未纠正读取错误。",
	},
	18: {
		ID:          18,
		DisplayName: "磁头健康自评估",
		DisplayType: AtaSmartAttributeDisplayTypeNormalized,
		Ideal:       ObservedThresholdIdealHigh,
		Critical:    false,
		Description: "部分厂商（西部数据）使用的厂商特定属性，用于报告磁头健康或磁头飞行高度状态。",
	},
	22: {
		ID:          22,
		DisplayName: "当前氦气水平",
		DisplayType: AtaSmartAttributeDisplayTypeNormalized,
		Ideal:       ObservedThresholdIdealHigh,
		Critical:    false,
		Description: "HGST He8（八碟氦气）硬盘专属。该值测量该厂商硬盘内部的氦气水平。这是一个预告故障（pre-fail）属性，一旦硬盘检测到内部环境超出规格即触发。",
	},
	23: {
		ID:          23,
		DisplayName: "氦气状态（低位）",
		DisplayType: AtaSmartAttributeDisplayTypeNormalized,
		Ideal:       ObservedThresholdIdealHigh,
		Critical:    false,
		Description: "表示氦气硬盘内部的氦气状态。用于 HGST/WD 氦气硬盘。",
	},
	24: {
		ID:          24,
		DisplayName: "氦气状态（高位）",
		DisplayType: AtaSmartAttributeDisplayTypeNormalized,
		Ideal:       ObservedThresholdIdealHigh,
		Critical:    false,
		Description: "表示氦气硬盘内部的氦气状态。用于 HGST/WD 氦气硬盘。",
	},
	170: {
		ID:          170,
		DisplayName: "可用预留空间",
		DisplayType: AtaSmartAttributeDisplayTypeNormalized,
		Ideal:       "",
		Critical:    false,
		Description: "参见属性 E8。",
	},
	171: {
		ID:          171,
		DisplayName: "SSD 编程失败计数",
		DisplayType: AtaSmartAttributeDisplayTypeNormalized,
		Ideal:       "",
		Critical:    false,
		Description: "（金士顿）硬盘投入使用以来闪存编程操作失败的总次数。与属性 181 相同。",
	},
	172: {
		ID:          172,
		DisplayName: "SSD 擦除失败计数",
		DisplayType: AtaSmartAttributeDisplayTypeNormalized,
		Ideal:       "",
		Critical:    false,
		Description: "（金士顿）统计闪存擦除失败次数。该属性返回硬盘投入使用以来闪存擦除操作失败的总次数。与属性 182 相同。",
	},
	173: {
		ID:          173,
		DisplayName: "SSD 磨损均衡计数",
		DisplayType: AtaSmartAttributeDisplayTypeNormalized,
		Ideal:       "",
		Critical:    false,
		Description: "统计所有块中最坏情况下的最大擦除次数。",
	},
	174: {
		ID:          174,
		DisplayName: "异常断电次数",
		Ideal:       "",
		Critical:    false,
		Description: "按传统 HDD 术语也称为“断电磁头缩回计数”。原始值报告 SSD 整个生命周期内累计的非正常关机次数，“非正常关机”指最后一条命令不是 STANDBY IMMEDIATE 的情况下直接断电（无论是否借助电容供电执行了 PLI 动作）。标准化值恒为 100。",
	},
	175: {
		ID:          175,
		DisplayName: "掉电保护失败",
		DisplayType: AtaSmartAttributeDisplayTypeNormalized,
		Ideal:       "",
		Critical:    false,
		Description: "最近一次测试结果，以电容放电所需微秒数表示，达到最大值后饱和。同时记录距上次测试的分钟数和终身测试次数。原始值包含以下数据：字节 0-1：最近一次测试结果（电容放电微秒数），最大值饱和。正常范围 25 ≤ 结果 ≤ 5000000，更小的值表示特定错误码。字节 2-3：距上次测试的分钟数，最大值饱和。字节 4-5：终身测试次数，不随断电周期重置，最大值饱和。测试失败时标准化值置为 1；电容曾在过温条件下测试则为 11；否则为 100。",
	},
	176: {
		ID:          176,
		DisplayName: "擦除失败计数",
		DisplayType: AtaSmartAttributeDisplayTypeNormalized,
		Ideal:       "",
		Critical:    false,
		Description: "S.M.A.R.T. 参数，表示闪存擦除命令失败的次数。",
	},
	177: {
		ID:          177,
		DisplayName: "磨损均衡计数",
		DisplayType: AtaSmartAttributeDisplayTypeNormalized,
		Ideal:       ObservedThresholdIdealHigh,
		Critical:    true,
		Description: "表示 SSD 闪存的介质编程与擦除次数（磨损指标）。标准化值越高越好（100 = 全新，数值下降 = 磨损加剧）。三星、英睿达等 SSD 厂商使用。",
		ObservedThresholds: []ObservedThreshold{
			{
				Low:               0,
				High:              0,
				AnnualFailureRate: 0.30,
				ErrorInterval:     []float64{0.20, 0.40},
			},
			{
				Low:               0,
				High:              9,
				AnnualFailureRate: 0.20,
				ErrorInterval:     []float64{0.15, 0.25},
			},
			{
				Low:               9,
				High:              24,
				AnnualFailureRate: 0.12,
				ErrorInterval:     []float64{0.09, 0.15},
			},
			{
				Low:               24,
				High:              49,
				AnnualFailureRate: 0.05,
				ErrorInterval:     []float64{0.03, 0.07},
			},
			{
				Low:               49,
				High:              74,
				AnnualFailureRate: 0.02,
				ErrorInterval:     []float64{0.01, 0.03},
			},
			{
				Low:               74,
				High:              99,
				AnnualFailureRate: 0.01,
				ErrorInterval:     []float64{0.005, 0.015},
			},
			{
				Low:               99,
				High:              255,
				AnnualFailureRate: 0.008,
				ErrorInterval:     []float64{0.004, 0.012},
			},
		},
	},
	179: {
		ID:          179,
		DisplayName: "已用预留块总数",
		DisplayType: AtaSmartAttributeDisplayTypeNormalized,
		Ideal:       "",
		Critical:    false,
		Description: "预告故障（Pre-Fail）属性，至少用于三星设备。",
	},
	180: {
		ID:          180,
		DisplayName: "未用预留块总数",
		DisplayType: AtaSmartAttributeDisplayTypeNormalized,
		Ideal:       "",
		Critical:    false,
		Description: "“预告故障”（Pre-Fail）属性，至少用于 HP 设备。",
	},
	181: {
		ID:          181,
		DisplayName: "编程失败总计数",
		DisplayType: AtaSmartAttributeDisplayTypeNormalized,
		Ideal:       "",
		Critical:    false,
		Description: "硬盘投入使用以来闪存编程操作失败的总次数。",
	},
	182: {
		ID:          182,
		DisplayName: "擦除失败计数",
		DisplayType: AtaSmartAttributeDisplayTypeNormalized,
		Ideal:       "",
		Critical:    false,
		Description: "“预告故障”（Pre-Fail）属性，至少用于三星设备。",
	},
	183: {
		ID:          183,
		DisplayName: "SATA 降速错误计数或运行时坏块",
		DisplayType: AtaSmartAttributeDisplayTypeRaw,
		Ideal:       ObservedThresholdIdealLow,
		Critical:    false,
		Description: "西部数据、三星或希捷属性：表示链路降速次数（如从 6Gbit/s 降至 3Gbit/s），或正常操作中检测到不可纠正错误的数据块总数。该参数恶化可能是硬盘老化和/或潜在机电问题的征兆，但不直接代表硬盘即将故障。",
		ObservedThresholds: []ObservedThreshold{
			{
				Low:               0,
				High:              0,
				AnnualFailureRate: 0.09084549203210031,
				ErrorInterval:     []float64{0.08344373475686712, 0.09872777224842152},
			},
			{
				Low:               1,
				High:              2,
				AnnualFailureRate: 0.05756065656498585,
				ErrorInterval:     []float64{0.04657000847949464, 0.07036491775108872},
			},
			{
				Low:               2,
				High:              4,
				AnnualFailureRate: 0.6193088626208925,
				ErrorInterval:     []float64{0.41784508895529787, 0.8841019099092139},
			},
			{
				Low:               4,
				High:              8,
				AnnualFailureRate: 0.5533447034299792,
				ErrorInterval:     []float64{0.31628430884775033, 0.8985971312402635},
			},
			{
				Low:               8,
				High:              16,
				AnnualFailureRate: 0.3882388694727245,
				ErrorInterval:     []float64{0.21225380267814295, 0.6513988534774338},
			},
			{
				Low:               16,
				High:              35,
				AnnualFailureRate: 0.37116708385481856,
				ErrorInterval:     []float64{0.19763084005134446, 0.6347070173754686},
			},
			{
				Low:               35,
				High:              70,
				AnnualFailureRate: 0.2561146752205292,
				ErrorInterval:     []float64{0.10297138269895259, 0.5276941165819332},
			},
			{
				Low:               70,
				High:              130,
				AnnualFailureRate: 0.40299684542586756,
				ErrorInterval:     []float64{0.16202563309223209, 0.8303275247667772},
			},
			{
				Low:               130,
				High:              260,
				AnnualFailureRate: 0,
				ErrorInterval:     []float64{0, 0},
			},
		},
	},
	184: {
		ID:          184,
		DisplayName: "端到端错误",
		DisplayType: AtaSmartAttributeDisplayTypeNormalized,
		Ideal:       ObservedThresholdIdealLow,
		Critical:    true,
		Description: "该属性属于惠普 SMART IV 技术，也属于其他厂商的 IO 错误检测与纠正方案，统计经由硬盘缓存 RAM 到介质的数据路径中发生的奇偶校验错误次数。",
		ObservedThresholds: []ObservedThreshold{
			{
				Low:               93,
				High:              94,
				AnnualFailureRate: 1.631212012870933,
				ErrorInterval:     []float64{1.055634407303844, 2.407990716767714},
			},
			{
				Low:               94,
				High:              95,
				AnnualFailureRate: 0,
				ErrorInterval:     []float64{0, 0},
			},
			{
				Low:               95,
				High:              96,
				AnnualFailureRate: 0,
				ErrorInterval:     []float64{0, 0},
			},
			{
				Low:               96,
				High:              97,
				AnnualFailureRate: 0,
				ErrorInterval:     []float64{0, 0},
			},
			{
				Low:               97,
				High:              97,
				AnnualFailureRate: 0,
				ErrorInterval:     []float64{0, 0},
			},
			{
				Low:               97,
				High:              98,
				AnnualFailureRate: 1.8069306930693072,
				ErrorInterval:     []float64{0.04574752432804858, 10.067573453924245},
			},
			{
				Low:               98,
				High:              99,
				AnnualFailureRate: 0.8371559633027523,
				ErrorInterval:     []float64{0.10138347095016888, 3.0240951820174824},
			},
			{
				Low:               99,
				High:              100,
				AnnualFailureRate: 0.09334816849865138,
				ErrorInterval:     []float64{0.08689499010435861, 0.10015372448181788},
			},
		},
	},
	185: {
		ID:          185,
		DisplayName: "磁头稳定性",
		DisplayType: AtaSmartAttributeDisplayTypeNormalized,
		Ideal:       "",
		Critical:    false,
		Description: "西部数据属性。",
	},
	186: {
		ID:          186,
		DisplayName: "运行时受激振动检测",
		DisplayType: AtaSmartAttributeDisplayTypeNormalized,
		Ideal:       "",
		Critical:    false,
		Description: "西部数据属性。",
	},
	187: {
		ID:          187,
		DisplayName: "上报的不可纠正错误数",
		DisplayType: AtaSmartAttributeDisplayTypeRaw,
		Ideal:       ObservedThresholdIdealLow,
		Critical:    true,
		Description: "无法通过硬件 ECC 恢复的错误计数（参见属性 195）。",
		ObservedThresholds: []ObservedThreshold{
			{
				Low:               0,
				High:              0,
				AnnualFailureRate: 0.028130798308190524,
				ErrorInterval:     []float64{0.024487830609364304, 0.032162944988161336},
			},
			{
				Low:               1,
				High:              1,
				AnnualFailureRate: 0.33877621175661743,
				ErrorInterval:     []float64{0.22325565823630591, 0.4929016016666955},
			},
			{
				Low:               1,
				High:              3,
				AnnualFailureRate: 0.24064820598237213,
				ErrorInterval:     []float64{0.14488594021076606, 0.3758019832614595},
			},
			{
				Low:               3,
				High:              6,
				AnnualFailureRate: 0.5014425058387142,
				ErrorInterval:     []float64{0.3062941096766342, 0.7744372808405151},
			},
			{
				Low:               6,
				High:              11,
				AnnualFailureRate: 0.38007108544136836,
				ErrorInterval:     []float64{0.2989500188963677, 0.4764223967570595},
			},
			{
				Low:               11,
				High:              20,
				AnnualFailureRate: 0.5346094598348444,
				ErrorInterval:     []float64{0.40595137663302483, 0.6911066985735377},
			},
			{
				Low:               20,
				High:              35,
				AnnualFailureRate: 0.8428063943161636,
				ErrorInterval:     []float64{0.6504601819243522, 1.0742259350903411},
			},
			{
				Low:               35,
				High:              65,
				AnnualFailureRate: 1.4429071005017484,
				ErrorInterval:     []float64{1.1405581860945952, 1.8008133631629157},
			},
			{
				Low:               65,
				High:              120,
				AnnualFailureRate: 1.6190935390549661,
				ErrorInterval:     []float64{1.0263664163011208, 2.4294352761068576},
			},
		},
	},
	188: {
		ID:          188,
		DisplayName: "命令超时",
		DisplayType: AtaSmartAttributeDisplayTypeTransformed,
		Ideal:       ObservedThresholdIdealLow,
		Critical:    true,
		Description: "因硬盘超时而中止的操作次数。正常情况下该属性值应为零。",
		Transform: func(normValue int64, rawValue int64, rawString string) int64 {
			// Parse Seagate command timeout values if the string contains 3 pieces
			// and each piece is less than or equal to the next (as a sanity check)
			// See https://github.com/AnalogJ/scrutiny/issues/522
			pieces := strings.Split(rawString, " ")
			if len(pieces) == 3 {
				int_pieces := make([]int, len(pieces))
				var err error
				for i, s := range pieces {
					int_pieces[i], err = strconv.Atoi(s)
					if err != nil {
						return rawValue
					}
				}
				if int_pieces[2] >= int_pieces[1] && int_pieces[1] >= int_pieces[0] {
					return int64(int_pieces[2])
				}
			}
			// Seagate drives use a packed 48-bit format with three 16-bit fields:
			// Bytes [1-0]: Total command timeout count
			// Bytes [3-2]: Commands with >5s completion time
			// Bytes [5-4]: Commands with >7.5s completion time
			// Reference: https://superuser.com/questions/1747844/interpreting-seagate-smart-value-188-command-timeout-data-format
			// Reference 2: https://t1.daumcdn.net/brunch/service/user/axm/file/zRYOdwPu3OMoKYmBOby1fEEQEbU.pdf
			if rawValue > 0xFFFF {
				// Extract total timeout count from lower 16 bits (bytes 0-1)
				// This represents all command timeouts regardless of duration
				total_timeouts := rawValue & 0xFFFF
				return total_timeouts
			}
			return rawValue
		},
		ObservedThresholds: []ObservedThreshold{
			{
				Low:               0,
				High:              0,
				AnnualFailureRate: 0.024893587674442153,
				ErrorInterval:     []float64{0.020857343769186413, 0.0294830350167543},
			},
			{
				Low: 0,
				// This is set arbitrarily to avoid notifications caused by low
				// historical numbers of command timeouts (e.g. caused by a bad cable)
				High:              100,
				AnnualFailureRate: 0.024893587674442153,
				ErrorInterval:     []float64{0.020857343769186413, 0.0294830350167543},
			},
			{
				Low:               100,
				High:              13000000000,
				AnnualFailureRate: 0.10044174089362015,
				ErrorInterval:     []float64{0.0812633664077498, 0.1227848196758574},
			},
			{
				Low:               13000000000,
				High:              26000000000,
				AnnualFailureRate: 0.334030592234279,
				ErrorInterval:     []float64{0.2523231196342665, 0.4337665082489293},
			},
			{
				Low:               26000000000,
				High:              39000000000,
				AnnualFailureRate: 0.36724705400842445,
				ErrorInterval:     []float64{0.30398009356575617, 0.4397986538328568},
			},
			{
				Low:               39000000000,
				High:              52000000000,
				AnnualFailureRate: 0.29848155926978354,
				ErrorInterval:     []float64{0.2509254838615984, 0.35242890006477073},
			},
			{
				Low:               52000000000,
				High:              65000000000,
				AnnualFailureRate: 0.2203079701535098,
				ErrorInterval:     []float64{0.18366082845676174, 0.26212468677179274},
			},
			{
				Low:               65000000000,
				High:              78000000000,
				AnnualFailureRate: 0.3018169948863018,
				ErrorInterval:     []float64{0.23779746376787655, 0.37776897542831006},
			},
			{
				Low:               78000000000,
				High:              91000000000,
				AnnualFailureRate: 0.32854928239235887,
				ErrorInterval:     []float64{0.2301118782147336, 0.4548506948185028},
			},
			{
				Low:               91000000000,
				High:              104000000000,
				AnnualFailureRate: 0.28488916640649387,
				ErrorInterval:     []float64{0.1366154288236293, 0.5239213202729072},
			},
		},
	},
	189: {
		ID:          189,
		DisplayName: "高飞写入次数",
		DisplayType: AtaSmartAttributeDisplayTypeRaw,
		Ideal:       ObservedThresholdIdealLow,
		Critical:    false,
		Description: "硬盘厂商实现了一种飞行高度传感器，通过检测记录磁头是否在正常工作范围之外飞行来为写入操作提供额外保护。当出现不安全的飞行高度时，写入过程会停止，数据将被重写或重映射到硬盘的安全区域。该属性表示硬盘整个生命周期内检测到的此类事件次数。",
		ObservedThresholds: []ObservedThreshold{
			{
				Low:               0,
				High:              0,
				AnnualFailureRate: 0.09070551401946862,
				ErrorInterval:     []float64{0.08018892683853401, 0.10221801211956287},
			},
			{
				Low:               1,
				High:              2,
				AnnualFailureRate: 0.0844336097370013,
				ErrorInterval:     []float64{0.07299813695315267, 0.09715235540340669},
			},
			{
				Low:               2,
				High:              5,
				AnnualFailureRate: 0.07943219628781906,
				ErrorInterval:     []float64{0.06552176680630226, 0.09542233189887633},
			},
			{
				Low:               5,
				High:              13,
				AnnualFailureRate: 0.09208847603893404,
				ErrorInterval:     []float64{0.07385765060838133, 0.11345557807163456},
			},
			{
				Low:               13,
				High:              30,
				AnnualFailureRate: 0.18161161650924224,
				ErrorInterval:     []float64{0.13858879602902988, 0.23377015012749933},
			},
			{
				Low:               30,
				High:              70,
				AnnualFailureRate: 0.2678117886102384,
				ErrorInterval:     []float64{0.19044036194841887, 0.36610753129699186},
			},
			{
				Low:               70,
				High:              150,
				AnnualFailureRate: 0.26126480798826107,
				ErrorInterval:     []float64{0.15958733218826962, 0.4035023060905559},
			},
			{
				Low:               150,
				High:              350,
				AnnualFailureRate: 0.11337164155924832,
				ErrorInterval:     []float64{0.030889956621649995, 0.2902764300762812},
			},
		},
	},
	190: {
		ID:          190,
		DisplayName: "温差",
		DisplayType: AtaSmartAttributeDisplayTypeNormalized,
		Ideal:       "",
		Critical:    false,
		Description: "数值等于 (100 − 温度°C)，厂商据此设定与最高温度对应的最小阈值。这也遵循“100 为最佳值、越低越不理想”的约定。不过部分老硬盘可能在此报告原始温度（与 0xC2 相同）或温度减 50。",
	},
	191: {
		ID:          191,
		DisplayName: "冲击感知错误率",
		DisplayType: AtaSmartAttributeDisplayTypeNormalized,
		Ideal:       ObservedThresholdIdealLow,
		Critical:    false,
		Description: "由外部冲击和振动导致的错误计数。",
	},
	192: {
		ID:          192,
		DisplayName: "断电磁头缩回计数",
		DisplayType: AtaSmartAttributeDisplayTypeRaw,
		Ideal:       ObservedThresholdIdealLow,
		Critical:    false,
		Description: "断电或紧急磁头缩回周期的次数。",
		ObservedThresholds: []ObservedThreshold{
			{
				Low:               1,
				High:              2,
				AnnualFailureRate: 0.02861098445412803,
				ErrorInterval:     []float64{0.022345416230915037, 0.036088863823297186},
			},
			{
				Low:               2,
				High:              6,
				AnnualFailureRate: 0.0738571777154862,
				ErrorInterval:     []float64{0.06406927746420421, 0.0847175264009771},
			},
			{
				Low:               6,
				High:              16,
				AnnualFailureRate: 0.11970378206823593,
				ErrorInterval:     []float64{0.10830059875098269, 0.13198105985656441},
			},
			{
				Low:               16,
				High:              40,
				AnnualFailureRate: 0.027266868552620425,
				ErrorInterval:     []float64{0.021131448605713823, 0.03462795920968522},
			},
			{
				Low:               40,
				High:              100,
				AnnualFailureRate: 0.011741682974559688,
				ErrorInterval:     []float64{0.00430899071133239, 0.025556700631152028},
			},
			{
				Low:               100,
				High:              250,
				AnnualFailureRate: 0.012659940134091309,
				ErrorInterval:     []float64{0.00607093338127348, 0.023282080653656938},
			},
			{
				Low:               250,
				High:              650,
				AnnualFailureRate: 0.01634692899031039,
				ErrorInterval:     []float64{0.009522688540043157, 0.026173016865409605},
			},
			{
				Low:               650,
				High:              1600,
				AnnualFailureRate: 0.005190074354440066,
				ErrorInterval:     []float64{0.0025908664180103293, 0.009286476666453648},
			},
		},
	},
	193: {
		ID:          193,
		DisplayName: "磁头加载周期计数",
		DisplayType: AtaSmartAttributeDisplayTypeNormalized,
		Ideal:       ObservedThresholdIdealLow,
		Critical:    false,
		Description: "磁头加载/卸载到着陆区的周期计数。部分硬盘改用属性 225 (0xE1) 记录加载周期计数。",
	},
	194: {
		ID:          194,
		DisplayName: "温度",
		DisplayType: AtaSmartAttributeDisplayTypeTransformed,
		Ideal:       ObservedThresholdIdealLow,
		Critical:    false,
		Description: "表示设备温度（需装有相应传感器）。原始值的最低字节为精确温度值（摄氏度）。",
		Transform: func(normValue int64, rawValue int64, rawString string) int64 {
			return rawValue & 0b11111111
		},
		TransformValueUnit: "°C",
	},
	195: {
		ID:          195,
		DisplayName: "硬件 ECC 已恢复错误",
		DisplayType: AtaSmartAttributeDisplayTypeNormalized,
		Ideal:       "",
		Critical:    false,
		Description: "（厂商特定原始值。）原始值结构因厂商而异，作为十进制数通常没有意义。",
		ObservedThresholds: []ObservedThreshold{
			{
				Low:               12,
				High:              24,
				AnnualFailureRate: 0.31472916829975706,
				ErrorInterval:     []float64{0.15711166685282174, 0.5631374192486645},
			},
			{
				Low:               24,
				High:              36,
				AnnualFailureRate: 0.15250310197260136,
				ErrorInterval:     []float64{0.10497611828070175, 0.21417105521823687},
			},
			{
				Low:               36,
				High:              48,
				AnnualFailureRate: 0.2193119102723874,
				ErrorInterval:     []float64{0.16475385681835103, 0.28615447006525274},
			},
			{
				Low:               48,
				High:              60,
				AnnualFailureRate: 0.05672658497265746,
				ErrorInterval:     []float64{0.043182904776447234, 0.07317316161437043},
			},
			{
				Low:               60,
				High:              72,
				AnnualFailureRate: 0,
				ErrorInterval:     []float64{0, 0},
			},
			{
				Low:               72,
				High:              84,
				AnnualFailureRate: 0,
				ErrorInterval:     []float64{0, 0},
			},
			{
				Low:               84,
				High:              96,
				AnnualFailureRate: 0,
				ErrorInterval:     []float64{0, 0},
			},
			{
				Low:               96,
				High:              108,
				AnnualFailureRate: 0.04074570216566197,
				ErrorInterval:     []float64{0.001031591863615295, 0.22702052218047528},
			},
		},
	},
	196: {
		ID:          196,
		DisplayName: "重映射事件计数",
		DisplayType: AtaSmartAttributeDisplayTypeRaw,
		Ideal:       ObservedThresholdIdealLow,
		Critical:    true,
		Description: "重映射操作的计数。原始值显示将数据从重映射扇区转移到备用区的尝试总次数，无论成功与否均计数。",
		ObservedThresholds: []ObservedThreshold{
			{
				Low:               0,
				High:              0,
				AnnualFailureRate: 0.007389855800729792,
				ErrorInterval:     []float64{0.005652654139732716, 0.009492578928212054},
			},
			{
				Low:               1,
				High:              1,
				AnnualFailureRate: 0.026558331312151347,
				ErrorInterval:     []float64{0.005476966404484466, 0.07761471429677293},
			},
			{
				Low:               1,
				High:              2,
				AnnualFailureRate: 0.02471894893674658,
				ErrorInterval:     []float64{0.0006258296027540169, 0.13772516847438018},
			},
			{
				Low:               2,
				High:              4,
				AnnualFailureRate: 0.03200912040691046,
				ErrorInterval:     []float64{0.0008104007642081744, 0.17834340416493005},
			},
			{
				Low:               4,
				High:              7,
				AnnualFailureRate: 0.043078012510326925,
				ErrorInterval:     []float64{0.001090640849081295, 0.24001532369794615},
			},
			{
				Low:               7,
				High:              11,
				AnnualFailureRate: 0.033843300880853036,
				ErrorInterval:     []float64{0.0008568381932559863, 0.18856280368036135},
			},
			{
				Low:               11,
				High:              17,
				AnnualFailureRate: 0.16979376647542252,
				ErrorInterval:     []float64{0.035015556653263225, 0.49620943874336304},
			},
			{
				Low:               17,
				High:              27,
				AnnualFailureRate: 0.059042381106438044,
				ErrorInterval:     []float64{0.0014948236677880642, 0.32896309247698113},
			},
			{
				Low:               27,
				High:              45,
				AnnualFailureRate: 0.24701105346266636,
				ErrorInterval:     []float64{0.050939617608142244, 0.721871118983972},
			},
		},
	},
	197: {
		ID:          197,
		DisplayName: "当前待映射扇区数",
		DisplayType: AtaSmartAttributeDisplayTypeRaw,
		Ideal:       ObservedThresholdIdealLow,
		Critical:    true,
		Description: "“不稳定”扇区（因不可恢复的读取错误而等待重映射）的计数。如果不稳定扇区随后被成功读取，则会被重映射且该值减少。扇区发生读取错误不会立即重映射（因为无法读出正确值、不知道要重映射什么内容，而且它之后可能又能读了）；硬盘固件会记住该扇区需要重映射，并在下次写入时执行。",
		ObservedThresholds: []ObservedThreshold{
			{
				Low:               0,
				High:              0,
				AnnualFailureRate: 0.025540791394761345,
				ErrorInterval:     []float64{0.023161777231213983, 0.02809784482748174},
			},
			{
				Low:               1,
				High:              2,
				AnnualFailureRate: 0.34196613799103254,
				ErrorInterval:     []float64{0.22723401523750225, 0.4942362818474496},
			},
			{
				Low:               2,
				High:              6,
				AnnualFailureRate: 0.6823772508117681,
				ErrorInterval:     []float64{0.41083568090070416, 1.0656166047061635},
			},
			{
				Low:               6,
				High:              16,
				AnnualFailureRate: 0.6108100007493069,
				ErrorInterval:     []float64{0.47336936083368364, 0.7757071095273286},
			},
			{
				Low:               16,
				High:              40,
				AnnualFailureRate: 0.9564879341127684,
				ErrorInterval:     []float64{0.7701044196378299, 1.174355230793638},
			},
			{
				Low:               40,
				High:              100,
				AnnualFailureRate: 1.6519989942167461,
				ErrorInterval:     []float64{1.328402276482456, 2.0305872327541317},
			},
			{
				Low:               100,
				High:              250,
				AnnualFailureRate: 2.5137741046831956,
				ErrorInterval:     []float64{1.9772427971560862, 3.1510376077891613},
			},
			{
				Low:               250,
				High:              650,
				AnnualFailureRate: 3.3203378817413904,
				ErrorInterval:     []float64{2.5883662702274406, 4.195047163573006},
			},
			{
				Low:               650,
				High:              1600,
				AnnualFailureRate: 3.133047210300429,
				ErrorInterval:     []float64{1.1497731080460096, 6.819324775707182},
			},
		},
	},
	198: {
		ID:          198,
		DisplayName: "（离线）不可纠正扇区数",
		DisplayType: AtaSmartAttributeDisplayTypeRaw,
		Ideal:       ObservedThresholdIdealLow,
		Critical:    true,
		Description: "读/写扇区时不可纠正错误的总数。 该属性值上升表明盘片表面存在缺陷和/或机械子系统存在问题。",
		ObservedThresholds: []ObservedThreshold{
			{
				Low:               0,
				High:              0,
				AnnualFailureRate: 0.028675322159886437,
				ErrorInterval:     []float64{0.026159385510707116, 0.03136793218577656},
			},
			{
				Low:               0,
				High:              2,
				AnnualFailureRate: 0.8135764944275583,
				ErrorInterval:     []float64{0.40613445471964466, 1.4557130815309443},
			},
			{
				Low:               2,
				High:              4,
				AnnualFailureRate: 1.1173469387755102,
				ErrorInterval:     []float64{0.5773494680315332, 1.9517802404552516},
			},
			{
				Low:               4,
				High:              6,
				AnnualFailureRate: 1.3558692421991083,
				ErrorInterval:     []float64{0.4402470522980859, 3.1641465148237544},
			},
			{
				Low:               6,
				High:              8,
				AnnualFailureRate: 0.7324414715719062,
				ErrorInterval:     []float64{0.15104704003805655, 2.140504796291604},
			},
			{
				Low:               8,
				High:              10,
				AnnualFailureRate: 0.5777213677766163,
				ErrorInterval:     []float64{0.43275294849366835, 0.7556737733062419},
			},
			{
				Low:               10,
				High:              12,
				AnnualFailureRate: 1.7464114832535886,
				ErrorInterval:     []float64{0.47583835092536914, 4.471507017371231},
			},
			{
				Low:               12,
				High:              14,
				AnnualFailureRate: 2.6449275362318843,
				ErrorInterval:     []float64{0.3203129951758959, 9.554387676519005},
			},
			{
				Low:               14,
				High:              16,
				AnnualFailureRate: 0.796943231441048,
				ErrorInterval:     []float64{0.5519063550198366, 1.113648286331181},
			},
		},
	},
	199: {
		ID:                 199,
		DisplayName:        "UltraDMA CRC 错误计数",
		DisplayType:        AtaSmartAttributeDisplayTypeRaw,
		Ideal:              ObservedThresholdIdealLow,
		Critical:           false,
		UseDeltaEvaluation: true,
		Description:        "由 ICRC（接口循环冗余校验）判定的、经由接口线缆传输数据时的错误计数。",
		ObservedThresholds: []ObservedThreshold{
			{
				Low:               0,
				High:              1,
				AnnualFailureRate: 0.04068379316116366,
				ErrorInterval:     []float64{0.037534031558106425, 0.04402730201866553},
			},
			{
				Low:               1,
				High:              2,
				AnnualFailureRate: 0.1513481259734218,
				ErrorInterval:     []float64{0.12037165605991791, 0.18786293065527596},
			},
			{
				Low:               2,
				High:              4,
				AnnualFailureRate: 0.16849758722418978,
				ErrorInterval:     []float64{0.12976367397863445, 0.2151676572000481},
			},
			{
				Low:               4,
				High:              8,
				AnnualFailureRate: 0.15385127340491614,
				ErrorInterval:     []float64{0.10887431782430312, 0.21117289306426648},
			},
			{
				Low:               8,
				High:              16,
				AnnualFailureRate: 0.14882894050104387,
				ErrorInterval:     []float64{0.09631424312463635, 0.2197008753522735},
			},
			{
				Low:               16,
				High:              35,
				AnnualFailureRate: 0.20878219917249793,
				ErrorInterval:     []float64{0.14086447304552446, 0.29804957135975},
			},
			{
				Low:               35,
				High:              70,
				AnnualFailureRate: 0.13742940270409038,
				ErrorInterval:     []float64{0.06860426267470295, 0.24589916335290812},
			},
			{
				Low:               70,
				High:              130,
				AnnualFailureRate: 0.22336578581363,
				ErrorInterval:     []float64{0.11150339549604707, 0.39966309081252904},
			},
			{
				Low:               130,
				High:              260,
				AnnualFailureRate: 0.18277416124186283,
				ErrorInterval:     []float64{0.07890890989692058, 0.3601379610272007},
			},
		},
	},
	200: {
		ID:          200,
		DisplayName: "多区错误率",
		DisplayType: AtaSmartAttributeDisplayTypeNormalized,
		Ideal:       ObservedThresholdIdealLow,
		Critical:    false,
		Description: "写入扇区时发现的错误计数。数值越高，磁盘机械状况越差。",
	},
	201: {
		ID:          201,
		DisplayName: "软读取错误率",
		DisplayType: AtaSmartAttributeDisplayTypeNormalized,
		Ideal:       ObservedThresholdIdealLow,
		Critical:    true,
		Description: "表示不可纠正的软件读取错误的次数。",
	},
	202: {
		ID:          202,
		DisplayName: "数据地址标记错误",
		DisplayType: AtaSmartAttributeDisplayTypeNormalized,
		Ideal:       ObservedThresholdIdealLow,
		Critical:    false,
		Description: "数据地址标记（DAM）错误的计数（或厂商特定含义）。",
	},
	203: {
		ID:          203,
		DisplayName: "Run Out Cancel",
		DisplayType: AtaSmartAttributeDisplayTypeNormalized,
		Ideal:       ObservedThresholdIdealLow,
		Critical:    false,
		Description: "纠错过程中由错误校验和导致的错误次数。",
	},
	204: {
		ID:          204,
		DisplayName: "软件 ECC 纠正",
		Ideal:       ObservedThresholdIdealLow,
		Critical:    false,
		Description: "由内部纠错软件纠正的错误计数。",
	},
	205: {
		ID:          205,
		DisplayName: "热划伤率",
		DisplayType: AtaSmartAttributeDisplayTypeNormalized,
		Ideal:       ObservedThresholdIdealLow,
		Critical:    false,
		Description: "由高温导致的错误计数。",
	},
	206: {
		ID:          206,
		DisplayName: "磁头飞行高度",
		DisplayType: AtaSmartAttributeDisplayTypeNormalized,
		Ideal:       "",
		Critical:    false,
		Description: "磁头与盘片表面的间距。过低更容易发生磁头碰撞，过高则更容易出现读写错误。",
	},
	207: {
		ID:          207,
		DisplayName: "起转峰值电流",
		DisplayType: AtaSmartAttributeDisplayTypeNormalized,
		Ideal:       ObservedThresholdIdealLow,
		Critical:    false,
		Description: "硬盘起转时所用的浪涌电流量。",
	},
	208: {
		ID:          208,
		DisplayName: "起转嗡鸣",
		DisplayType: AtaSmartAttributeDisplayTypeNormalized,
		Ideal:       "",
		Critical:    false,
		Description: "因供电不足而起转所需的嗡鸣（buzz）例程序数。",
	},
	209: {
		ID:          209,
		DisplayName: "离线寻道性能",
		DisplayType: AtaSmartAttributeDisplayTypeNormalized,
		Ideal:       "",
		Critical:    false,
		Description: "硬盘内部测试期间的寻道性能。",
	},
	210: {
		ID:          210,
		DisplayName: "写入时振动",
		DisplayType: AtaSmartAttributeDisplayTypeNormalized,
		Ideal:       "",
		Critical:    false,
		Description: "见于迈拓 6B200M0 200GB 和迈拓 2R015H1 15GB 硬盘。",
	},
	211: {
		ID:          211,
		DisplayName: "写入时振动",
		DisplayType: AtaSmartAttributeDisplayTypeNormalized,
		Ideal:       "",
		Critical:    false,
		Description: "对写入操作期间遇到的振动的记录。",
	},
	212: {
		ID:          212,
		DisplayName: "写入时冲击",
		DisplayType: AtaSmartAttributeDisplayTypeNormalized,
		Ideal:       "",
		Critical:    false,
		Description: "对写入操作期间遇到的冲击的记录。",
	},
	220: {
		ID:          220,
		DisplayName: "盘片偏移",
		Ideal:       ObservedThresholdIdealLow,
		DisplayType: AtaSmartAttributeDisplayTypeNormalized,
		Critical:    false,
		Description: "盘片相对于主轴的偏移距离（通常由冲击或温度引起）。测量单位未知。",
	},
	221: {
		ID:          221,
		DisplayName: "冲击感知错误率",
		Ideal:       ObservedThresholdIdealLow,
		DisplayType: AtaSmartAttributeDisplayTypeNormalized,
		Critical:    false,
		Description: "由外部冲击和振动导致的错误计数。",
	},
	222: {
		ID:          222,
		DisplayName: "受载时长",
		DisplayType: AtaSmartAttributeDisplayTypeNormalized,
		Ideal:       "",
		Critical:    false,
		Description: "在数据负载下工作的时间（磁头臂的运动）。",
	},
	223: {
		ID:          223,
		DisplayName: "加载/卸载重试次数",
		DisplayType: AtaSmartAttributeDisplayTypeNormalized,
		Ideal:       "",
		Critical:    false,
		Description: "磁头改变位置的次数。",
	},
	224: {
		ID:          224,
		DisplayName: "加载摩擦力",
		Ideal:       ObservedThresholdIdealLow,
		DisplayType: AtaSmartAttributeDisplayTypeNormalized,
		Critical:    false,
		Description: "运行时机械部件摩擦造成的阻力。",
	},
	225: {
		ID:          225,
		DisplayName: "加载/卸载周期计数",
		Ideal:       ObservedThresholdIdealLow,
		DisplayType: AtaSmartAttributeDisplayTypeNormalized,
		Critical:    false,
		Description: "加载周期总数。部分硬盘改用属性 193 (0xC1) 记录加载周期计数，其含义参见属性 193 的描述。",
	},
	226: {
		ID:          226,
		DisplayName: "磁头加载时间",
		DisplayType: AtaSmartAttributeDisplayTypeNormalized,
		Ideal:       "",
		Critical:    false,
		Description: "磁头臂架处于加载状态的总时间（不在停泊区的时间）。",
	},
	227: {
		ID:          227,
		DisplayName: "扭矩放大计数",
		Ideal:       ObservedThresholdIdealLow,
		Critical:    false,
		Description: "为补偿盘片转速波动而进行补偿的尝试次数。",
	},
	228: {
		ID:          228,
		DisplayName: "断电磁头缩回周期",
		Ideal:       ObservedThresholdIdealLow,
		Critical:    false,
		Description: "断电周期次数。每当发生“缩回事件”、磁头离开介质（例如断电、休眠或空闲）时计数。",
	},
	230: {
		ID:          230,
		DisplayName: "GMR 磁头幅度",
		DisplayType: AtaSmartAttributeDisplayTypeNormalized,
		Ideal:       "",
		Critical:    false,
		Description: "“抖动”幅度（操作之间磁头的反复移动）。",
	},
	231: {
		ID:          231,
		DisplayName: "剩余寿命",
		DisplayType: AtaSmartAttributeDisplayTypeNormalized,
		Ideal:       "",
		Critical:    false,
		Description: "以编程/擦除周期或可用预留块衡量 SSD 的近似剩余寿命。标准化值 100 表示全新硬盘，阈值 10 表示需要更换；值为 0 可能意味着硬盘正以只读模式运行以允许数据恢复。",
	},
	232: {
		ID:          232,
		DisplayName: "剩余耐久度",
		DisplayType: AtaSmartAttributeDisplayTypeNormalized,
		Ideal:       "",
		Critical:    false,
		Description: "SSD 已完成的物理擦除周期数占其设计可承受的最大物理擦除周期数的百分比。",
	},
	233: {
		ID:          233,
		DisplayName: "介质磨损指示",
		DisplayType: AtaSmartAttributeDisplayTypeNormalized,
		Ideal:       "",
		Critical:    false,
		Description: "Intel SSD 报告的标准化值从 100（全新）最低降到 1。随 NAND 擦除周期从 0 增加到额定最大周期，该值逐渐下降。",
	},
	234: {
		ID:          234,
		DisplayName: "平均擦除次数",
		DisplayType: AtaSmartAttributeDisplayTypeNormalized,
		Ideal:       "",
		Critical:    false,
		Description: "解码方式：字节 0-1-2 = 平均擦除次数（大端），字节 3-4-5 = 最大擦除次数（大端）。",
	},
	235: {
		ID:          235,
		DisplayName: "良好块计数",
		DisplayType: AtaSmartAttributeDisplayTypeNormalized,
		Ideal:       "",
		Critical:    false,
		Description: "解码方式：字节 0-1-2 = 良好块计数（大端），字节 3-4 = 系统（空闲）块计数。",
	},
	240: {
		ID:          240,
		DisplayName: "磁头飞行时长",
		DisplayType: AtaSmartAttributeDisplayTypeNormalized,
		Ideal:       "",
		Critical:    false,
		Description: "磁头定位所花费的时间。部分富士通硬盘报告的是数据传输期间链路重置的次数。",
	},
	241: {
		ID:          241,
		DisplayName: "LBA 写入总数",
		DisplayType: AtaSmartAttributeDisplayTypeNormalized,
		Ideal:       "",
		Critical:    false,
		Description: "LBA（逻辑块地址）写入总数。",
	},
	242: {
		ID:          242,
		DisplayName: "LBA 读取总数",
		DisplayType: AtaSmartAttributeDisplayTypeNormalized,
		Ideal:       "",
		Critical:    false,
		Description: "LBA 读取总数。某些 S.M.A.R.T. 工具报告的原始值为负数，因为它实际是 48 位而非 32 位。",
	},
	243: {
		ID:          243,
		DisplayName: "LBA 写入总数（扩展）",
		DisplayType: AtaSmartAttributeDisplayTypeNormalized,
		Ideal:       "",
		Critical:    false,
		Description: "写入设备的 12 字节 LBA 总数的高 5 字节。低 7 字节的值位于属性 0xF1。",
	},
	244: {
		ID:          244,
		DisplayName: "LBA 读取总数（扩展）",
		DisplayType: AtaSmartAttributeDisplayTypeNormalized,
		Ideal:       "",
		Critical:    false,
		Description: "从设备读取的 12 字节 LBA 总数的高 5 字节。低 7 字节的值位于属性 0xF2。",
	},
	246: {
		ID:          246,
		DisplayName: "主机扇区写入总数",
		DisplayType: AtaSmartAttributeDisplayTypeNormalized,
		Ideal:       "",
		Critical:    false,
		Description: "主机系统写入的扇区总数。英睿达/美光 SSD 使用。",
	},
	247: {
		ID:          247,
		DisplayName: "主机编程页数",
		DisplayType: AtaSmartAttributeDisplayTypeNormalized,
		Ideal:       "",
		Critical:    false,
		Description: "主机写入命令编程的页总数。英睿达/美光 SSD 使用。",
	},
	248: {
		ID:          248,
		DisplayName: "FTL 编程页数",
		DisplayType: AtaSmartAttributeDisplayTypeNormalized,
		Ideal:       "",
		Critical:    false,
		Description: "FTL（闪存转换层）编程的页总数。英睿达/美光 SSD 用于垃圾回收等后台操作。",
	},
	249: {
		ID:          249,
		DisplayName: "NAND 写入量 (1GiB)",
		DisplayType: AtaSmartAttributeDisplayTypeNormalized,
		Ideal:       "",
		Critical:    false,
		Description: "NAND 写入总量。原始值以 1 GB 为增量报告 NAND 写入次数。",
	},
	250: {
		ID:          250,
		DisplayName: "读取错误重试率",
		DisplayType: AtaSmartAttributeDisplayTypeNormalized,
		Ideal:       ObservedThresholdIdealLow,
		Critical:    false,
		Description: "从磁盘读取时的错误计数。",
	},
	251: {
		ID:          251,
		DisplayName: "剩余最小备用块",
		DisplayType: AtaSmartAttributeDisplayTypeNormalized,
		Ideal:       "",
		Critical:    false,
		Description: "“剩余最小备用块”属性表示剩余备用块占可用备用块总数的百分比。",
	},
	252: {
		ID:          252,
		DisplayName: "新增坏闪存块",
		DisplayType: AtaSmartAttributeDisplayTypeNormalized,
		Ideal:       "",
		Critical:    false,
		Description: "“新增坏闪存块”属性表示硬盘自出厂初始化以来检测到的坏闪存块总数。",
	},
	254: {
		ID:          254,
		DisplayName: "跌落保护",
		DisplayType: AtaSmartAttributeDisplayTypeNormalized,
		Ideal:       ObservedThresholdIdealLow,
		Critical:    false,
		Description: "检测到的“跌落事件”计数。",
	},
}

// AtaDeviceStatisticsMetadata provides metadata for ATA Device Statistics (GP Log 0x04)
// These are string-keyed using the format "devstat_<page>_<offset>"
// Key device statistics pages:
// - Page 1: General Statistics (power-on hours, logical sectors written/read)
// - Page 2: Free-fall Statistics
// - Page 3: Rotating Media Statistics
// - Page 4: General Errors Statistics
// - Page 5: Temperature Statistics
// - Page 7: SSD Statistics (Percentage Used Endurance Indicator)
type AtaDeviceStatisticsMetadata struct {
	DisplayName string `json:"display_name"`
	Ideal       string `json:"ideal"`
	Critical    bool   `json:"critical"`
	Description string `json:"description"`
	DisplayType string `json:"display_type"`
	Threshold   int64  `json:"threshold"` // Fixed threshold for this attribute. 0 means no fixed threshold (use for error counts).
}

var AtaDeviceStatsMetadata = map[string]AtaDeviceStatisticsMetadata{
	// Page 7 (Solid State Device Statistics) - most important for SSD health
	"devstat_7_8": {
		DisplayName: "耐久度已用百分比指示",
		Ideal:       ObservedThresholdIdealLow,
		Critical:    true,
		Description: "基于设备实际使用情况和厂商对寿命的预测，给出的设备寿命已用百分比的厂商估算值。值为 100 表示设备预计耐久度已耗尽，但不一定代表设备故障。",
		DisplayType: AtaSmartAttributeDisplayTypeRaw,
		Threshold:   100,
	},
	// Page 1 (General Statistics)
	"devstat_1_8": {
		DisplayName: "累计上电复位次数",
		Ideal:       "",
		Critical:    false,
		Description: "设备出厂以来上电复位事件的次数。",
		DisplayType: AtaSmartAttributeDisplayTypeRaw,
	},
	"devstat_1_16": {
		DisplayName: "通电时长（设备统计）",
		Ideal:       "",
		Critical:    false,
		Description: "来自设备统计的累计通电时间（小时）。",
		DisplayType: AtaSmartAttributeDisplayTypeRaw,
	},
	"devstat_1_24": {
		DisplayName: "逻辑扇区写入数",
		Ideal:       "",
		Critical:    false,
		Description: "写入设备的逻辑扇区总数。",
		DisplayType: AtaSmartAttributeDisplayTypeRaw,
	},
	"devstat_1_32": {
		DisplayName: "写命令数",
		Ideal:       "",
		Critical:    false,
		Description: "已处理的写命令总数。",
		DisplayType: AtaSmartAttributeDisplayTypeRaw,
	},
	"devstat_1_40": {
		DisplayName: "逻辑扇区读取数",
		Ideal:       "",
		Critical:    false,
		Description: "从设备读取的逻辑扇区总数。",
		DisplayType: AtaSmartAttributeDisplayTypeRaw,
	},
	"devstat_1_48": {
		DisplayName: "读命令数",
		Ideal:       "",
		Critical:    false,
		Description: "已处理的读命令总数。",
		DisplayType: AtaSmartAttributeDisplayTypeRaw,
	},
	// Page 5 (Temperature Statistics)
	"devstat_5_8": {
		DisplayName: "当前温度（设备统计）",
		Ideal:       "",
		Critical:    false,
		Description: "来自设备统计的当前设备温度。",
		DisplayType: AtaSmartAttributeDisplayTypeRaw,
	},
	"devstat_5_16": {
		DisplayName: "短期平均温度",
		Ideal:       "",
		Critical:    false,
		Description: "短期周期内的平均温度。",
		DisplayType: AtaSmartAttributeDisplayTypeRaw,
	},
	"devstat_5_24": {
		DisplayName: "长期平均温度",
		Ideal:       "",
		Critical:    false,
		Description: "长期周期内的平均温度。",
		DisplayType: AtaSmartAttributeDisplayTypeRaw,
	},
	"devstat_5_32": {
		DisplayName: "最高温度",
		Ideal:       "",
		Critical:    false,
		Description: "出厂以来达到的最高温度。",
		DisplayType: AtaSmartAttributeDisplayTypeRaw,
	},
	"devstat_5_40": {
		DisplayName: "最低温度",
		Ideal:       "",
		Critical:    false,
		Description: "出厂以来达到的最低温度。",
		DisplayType: AtaSmartAttributeDisplayTypeRaw,
	},
	"devstat_5_88": {
		DisplayName: "超温时长",
		Ideal:       ObservedThresholdIdealLow,
		Critical:    false,
		Description: "设备在超过建议最高温度下运行的分钟数。",
		DisplayType: AtaSmartAttributeDisplayTypeRaw,
	},
	// Page 2 (Free-Fall Statistics)
	"devstat_2_16": {
		DisplayName: "检测到的跌落事件数",
		Ideal:       ObservedThresholdIdealLow,
		Critical:    false,
		Description: "加速度计检测到的跌落事件次数。带跌落保护的硬盘使用。",
		DisplayType: AtaSmartAttributeDisplayTypeRaw,
	},
	// Page 4 (General Errors Statistics)
	"devstat_4_8": {
		DisplayName: "上报的不可纠正错误数",
		Ideal:       ObservedThresholdIdealLow,
		Critical:    true,
		Description: "上报给主机的不可纠正错误总数。",
		DisplayType: AtaSmartAttributeDisplayTypeRaw,
	},
	"devstat_4_16": {
		DisplayName: "接口 CRC 错误数",
		Ideal:       ObservedThresholdIdealLow,
		Critical:    false,
		Description: "检测到的接口 CRC 错误数（可能意味着线缆问题）。",
		DisplayType: AtaSmartAttributeDisplayTypeRaw,
	},
	// Page 6 (Transport Statistics)
	"devstat_6_8": {
		DisplayName: "硬件复位次数",
		Ideal:       ObservedThresholdIdealLow,
		Critical:    false,
		Description: "设备收到的硬件复位事件次数。",
		DisplayType: AtaSmartAttributeDisplayTypeRaw,
	},
	"devstat_6_16": {
		DisplayName: "ASR（异步信号恢复）事件数",
		Ideal:       ObservedThresholdIdealLow,
		Critical:    false,
		Description: "异步信号恢复（ASR）事件次数。",
		DisplayType: AtaSmartAttributeDisplayTypeRaw,
	},
	"devstat_6_24": {
		DisplayName: "接口 CRC 错误数（传输层）",
		Ideal:       ObservedThresholdIdealLow,
		Critical:    false,
		Description: "来自传输层统计的接口 CRC 错误数。",
		DisplayType: AtaSmartAttributeDisplayTypeRaw,
	},
	// Page 3 (Rotating Media Statistics) - for HDDs
	"devstat_3_8": {
		DisplayName: "主轴电机通电时长",
		Ideal:       "",
		Critical:    false,
		Description: "主轴电机累计通电小时数。",
		DisplayType: AtaSmartAttributeDisplayTypeRaw,
	},
	"devstat_3_16": {
		DisplayName: "磁头飞行时长",
		Ideal:       "",
		Critical:    false,
		Description: "磁头在介质上方飞行（悬浮）的累计小时数。",
		DisplayType: AtaSmartAttributeDisplayTypeRaw,
	},
	"devstat_3_24": {
		DisplayName: "磁头加载事件",
		Ideal:       "",
		Critical:    false,
		Description: "磁头加载事件次数（磁头从停泊位移动到工作位）。",
		DisplayType: AtaSmartAttributeDisplayTypeRaw,
	},
	"devstat_3_32": {
		DisplayName: "已重映射逻辑扇区数",
		Ideal:       ObservedThresholdIdealLow,
		Critical:    true,
		Description: "因缺陷而已被重映射的逻辑扇区数量。",
		DisplayType: AtaSmartAttributeDisplayTypeRaw,
	},
	"devstat_3_40": {
		DisplayName: "读取恢复尝试次数",
		Ideal:       ObservedThresholdIdealLow,
		Critical:    false,
		Description: "已执行的读取恢复尝试次数。",
		DisplayType: AtaSmartAttributeDisplayTypeRaw,
	},
	"devstat_3_48": {
		DisplayName: "机械启动失败次数",
		Ideal:       ObservedThresholdIdealLow,
		Critical:    true,
		Description: "主轴电机未能达到工作转速的次数。",
		DisplayType: AtaSmartAttributeDisplayTypeRaw,
	},
	"devstat_3_56": {
		DisplayName: "待重映射逻辑扇区数",
		Ideal:       ObservedThresholdIdealLow,
		Critical:    true,
		Description: "等待重映射的逻辑扇区数量。",
		DisplayType: AtaSmartAttributeDisplayTypeRaw,
	},
	"devstat_3_64": {
		DisplayName: "高优先级磁头卸载事件数",
		Ideal:       ObservedThresholdIdealLow,
		Critical:    false,
		Description: "因检测到冲击或错误而紧急卸载磁头的次数。",
		DisplayType: AtaSmartAttributeDisplayTypeRaw,
	},
	// Page 5 (Temperature Statistics) - additional entries
	"devstat_5_48": {
		DisplayName: "最高短期平均温度",
		Ideal:       "",
		Critical:    false,
		Description: "记录到的最高短期平均温度。",
		DisplayType: AtaSmartAttributeDisplayTypeRaw,
	},
	"devstat_5_56": {
		DisplayName: "最低短期平均温度",
		Ideal:       "",
		Critical:    false,
		Description: "记录到的最低短期平均温度。",
		DisplayType: AtaSmartAttributeDisplayTypeRaw,
	},
	"devstat_5_64": {
		DisplayName: "最高长期平均温度",
		Ideal:       "",
		Critical:    false,
		Description: "记录到的最高长期平均温度。",
		DisplayType: AtaSmartAttributeDisplayTypeRaw,
	},
	"devstat_5_72": {
		DisplayName: "最低长期平均温度",
		Ideal:       "",
		Critical:    false,
		Description: "记录到的最低长期平均温度。",
		DisplayType: AtaSmartAttributeDisplayTypeRaw,
	},
	"devstat_5_80": {
		DisplayName: "低温时长",
		Ideal:       ObservedThresholdIdealLow,
		Critical:    false,
		Description: "设备在低于建议最低温度下运行的分钟数。",
		DisplayType: AtaSmartAttributeDisplayTypeRaw,
	},
	"devstat_5_96": {
		DisplayName: "规定最高工作温度",
		Ideal:       "",
		Critical:    false,
		Description: "厂商规定的工作温度上限。",
		DisplayType: AtaSmartAttributeDisplayTypeRaw,
	},
	"devstat_5_104": {
		DisplayName: "规定最低工作温度",
		Ideal:       "",
		Critical:    false,
		Description: "厂商规定的工作温度下限。",
		DisplayType: AtaSmartAttributeDisplayTypeRaw,
	},
	// Page 4 (General Errors Statistics) - additional entry
	"devstat_4_24": {
		DisplayName: "不可恢复读取错误数",
		Ideal:       ObservedThresholdIdealLow,
		Critical:    true,
		Description: "上报的不可恢复读取错误次数。",
		DisplayType: AtaSmartAttributeDisplayTypeRaw,
	},
	// Page 255 (Vendor Specific Statistics)
	"devstat_255_8": {
		DisplayName: "厂商特定统计 1",
		Ideal:       "",
		Critical:    false,
		Description: "厂商特定的设备统计项（因厂商而异）。",
		DisplayType: AtaSmartAttributeDisplayTypeRaw,
	},
	"devstat_255_16": {
		DisplayName: "厂商特定统计 2",
		Ideal:       "",
		Critical:    false,
		Description: "厂商特定的设备统计项（因厂商而异）。",
		DisplayType: AtaSmartAttributeDisplayTypeRaw,
	},
	"devstat_255_24": {
		DisplayName: "厂商特定统计 3",
		Ideal:       "",
		Critical:    false,
		Description: "厂商特定的设备统计项（因厂商而异）。",
		DisplayType: AtaSmartAttributeDisplayTypeRaw,
	},
}
