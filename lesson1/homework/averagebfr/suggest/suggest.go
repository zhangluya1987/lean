package suggest

func Suggest(Sex string,Age float64,bfrCalcResult float64) string {
	thinnish := "偏瘦"
	standard := "标准"
	weighting := "偏重"
	fat := "肥胖"
	severeObesity := "严重肥胖"
	superObese := "超级肥胖"
	inputError := "输入错误"
	minors := "小于18岁"

	switch Sex {
	case "男":
		if Age >= 18 && Age <= 39 {
			if bfrCalcResult >= 0 && bfrCalcResult <= 0.1 {
				return thinnish
			} else if bfrCalcResult >= 0.1 && bfrCalcResult <= 0.16 {
				return standard
			} else if bfrCalcResult >= 0.17 && bfrCalcResult <= 0.21 {
				return weighting
			} else if bfrCalcResult >= 0.22 && bfrCalcResult <= 0.26 {
				return fat
			} else if bfrCalcResult >= 0.27 && bfrCalcResult <= 0.45 {
				return severeObesity
			} else if bfrCalcResult >= 0.46 {
				return superObese
			} else {
				return inputError
			}
		} else if Age >= 40 && Age <= 59 {
			if bfrCalcResult >= 0 && bfrCalcResult <= 0.11 {
				return thinnish
			} else if bfrCalcResult >= 0.12 && bfrCalcResult <= 0.17 {
				return  standard
			} else if bfrCalcResult >= 0.18 && bfrCalcResult <= 0.22 {
				return weighting
			} else if bfrCalcResult >= 0.23 && bfrCalcResult <= 0.27 {
				return fat
			} else if bfrCalcResult >= 0.28 && bfrCalcResult <= 0.45 {
				return severeObesity
			} else if bfrCalcResult >= 0.46 {
				return superObese
			} else {
				return inputError
			}
		} else if Age >= 60 {
			if bfrCalcResult >= 0 && bfrCalcResult <= 0.13 {
				return thinnish
			} else if bfrCalcResult >= 0.14 && bfrCalcResult <= 0.19 {
				return  standard
			} else if bfrCalcResult >= 0.20 && bfrCalcResult <= 0.24 {
				return weighting
			} else if bfrCalcResult >= 0.25 && bfrCalcResult <= 0.29 {
				return fat
			} else if bfrCalcResult >= 0.30 && bfrCalcResult <= 0.45 {
				return severeObesity
			} else if bfrCalcResult >= 0.46 {
				return superObese
			} else {
				return inputError
			}
		} else {
			return minors
		}
	case "女":
		if Age >= 18 && Age <= 39 {
			if bfrCalcResult >= 0 && bfrCalcResult <= 0.20 {
				return thinnish
			} else if bfrCalcResult >= 0.21 && bfrCalcResult <= 0.27 {
				return  standard
			} else if bfrCalcResult >= 0.28 && bfrCalcResult <= 0.34 {
				return weighting
			} else if bfrCalcResult >= 0.35 && bfrCalcResult <= 0.39 {
				return fat
			} else if bfrCalcResult >= 0.40 && bfrCalcResult <= 0.45 {
				return severeObesity
			} else if bfrCalcResult >= 0.46 {
				return superObese
			} else {
				return inputError
			}
		} else if Age >= 40 && Age <= 59 {
			if bfrCalcResult >= 0 && bfrCalcResult <= 0.21 {
				return thinnish
			} else if bfrCalcResult >= 0.22 && bfrCalcResult <= 0.28 {
				return  standard
			} else if bfrCalcResult >= 0.29 && bfrCalcResult <= 0.35 {
				return weighting
			} else if bfrCalcResult >= 0.36 && bfrCalcResult <= 0.40 {
				return fat
			} else if bfrCalcResult >= 0.40 && bfrCalcResult <= 0.45 {
				return severeObesity
			} else if bfrCalcResult >= 0.46 {
				return superObese
			} else {
				return inputError
			}
		} else if Age >= 60 {
			if bfrCalcResult >= 0 && bfrCalcResult <= 0.22 {
				return thinnish
			} else if bfrCalcResult >= 0.23 && bfrCalcResult <= 0.29 {
				return  standard
			} else if bfrCalcResult >= 0.30 && bfrCalcResult <= 0.36 {
				return weighting
			} else if bfrCalcResult >= 0.37 && bfrCalcResult <= 0.41 {
				return fat
			} else if bfrCalcResult >= 0.42 && bfrCalcResult <= 0.45 {
				return severeObesity
			} else if bfrCalcResult >= 0.46 {
				return superObese
			} else {
				return inputError
			}
		} else {
			return minors
		}
	default:
		return inputError
	}
}