// Generated at 2023-09-05 22:06:22.506 -06:00

package liter

import (
	"strconv"
)

const (
	V_UNSET   = -1
	V_HIGHEST = V1_20
	V1_20     = 763 // https://wiki.vg/index.php?title=Protocol&oldid=18256
	V1_19_4   = 762 // https://wiki.vg/index.php?title=Protocol&oldid=18242
	V1_19_3   = 761 // https://wiki.vg/index.php?title=Protocol&oldid=18067
	V1_19_2   = 760 // https://wiki.vg/index.php?title=Protocol&oldid=17873
	V1_19     = 759 // https://wiki.vg/index.php?title=Protocol&oldid=17753
	V1_18_2   = 758 // https://wiki.vg/index.php?title=Protocol&oldid=17499
	V1_18_1   = 757 // https://wiki.vg/index.php?title=Protocol&oldid=17341
	V1_17_1   = 756 // https://wiki.vg/index.php?title=Protocol&oldid=16918
	V1_17     = 755 // https://wiki.vg/index.php?title=Protocol&oldid=16866
	V1_16_5   = 754 // https://wiki.vg/index.php?title=Protocol&oldid=16681
	V1_16_3   = 753 // https://wiki.vg/index.php?title=Protocol&oldid=16091
	V1_15_2   = 578 // https://wiki.vg/index.php?title=Protocol&oldid=16067
	V1_14_4   = 498 // https://wiki.vg/index.php?title=Protocol&oldid=15346
	V1_13_2   = 404 // https://wiki.vg/index.php?title=Protocol&oldid=14889
	V1_12_2   = 340 // https://wiki.vg/index.php?title=Protocol&oldid=14204
	V1_12_1   = 338 // https://wiki.vg/index.php?title=Protocol&oldid=13339
	V1_12     = 335 // https://wiki.vg/index.php?title=Protocol&oldid=13223
	V1_11_2   = 316 // https://wiki.vg/index.php?title=Protocol&oldid=8543
	V1_11     = 315 // https://wiki.vg/index.php?title=Protocol&oldid=8405
	V1_10_2   = 210 // https://wiki.vg/index.php?title=Protocol&oldid=8235
	V1_9_4    = 110 // https://wiki.vg/index.php?title=Protocol&oldid=7959
	V1_8_9    = 47  // https://wiki.vg/index.php?title=Protocol&oldid=7368
	V1_7_5    = 4   // https://wiki.vg/index.php?title=Protocol&oldid=5486
	V_LOWEST  = V1_7_5
)

var avaliableProtocols = []int{
	V1_20,
	V1_19_4,
	V1_19_3,
	V1_19_2,
	V1_19,
	V1_18_2,
	V1_18_1,
	V1_17_1,
	V1_17,
	V1_16_5,
	V1_16_3,
	V1_15_2,
	V1_14_4,
	V1_13_2,
	V1_12_2,
	V1_12_1,
	V1_12,
	V1_11_2,
	V1_11,
	V1_10_2,
	V1_9_4,
	V1_8_9,
	V1_7_5,
}

func GetProtocolName(p int) string {
	switch p {
	case V_UNSET:
		return "unset"
	case V1_20:
		return "1.20"
	case V1_19_4:
		return "1.19.4"
	case V1_19_3:
		return "1.19.3"
	case V1_19_2:
		return "1.19.2"
	case V1_19:
		return "1.19"
	case V1_18_2:
		return "1.18.2"
	case V1_18_1:
		return "1.18.1"
	case V1_17_1:
		return "1.17.1"
	case V1_17:
		return "1.17"
	case V1_16_5:
		return "1.16.5"
	case V1_16_3:
		return "1.16.3"
	case V1_15_2:
		return "1.15.2"
	case V1_14_4:
		return "1.14.4"
	case V1_13_2:
		return "1.13.2"
	case V1_12_2:
		return "1.12.2"
	case V1_12_1:
		return "1.12.1"
	case V1_12:
		return "1.12"
	case V1_11_2:
		return "1.11.2"
	case V1_11:
		return "1.11"
	case V1_10_2:
		return "1.10.2"
	case V1_9_4:
		return "1.9.4"
	case V1_8_9:
		return "1.8.9"
	case V1_7_5:
		return "1.7.5"
	}
	return "Unknown_" + strconv.Itoa(p)
}

func IsProtocolSupport(p int) bool {
	return V_LOWEST <= p && p <= V_HIGHEST
}

func ToNearestProtocol(p int) int {
	if !IsProtocolSupport(p) {
		return p
	}
	for _, v := range avaliableProtocols {
		if p >= v {
			return v
		}
	}
	panic("unreachable")
}
