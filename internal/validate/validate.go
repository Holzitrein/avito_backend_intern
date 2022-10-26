package validate

import (
	"avito_balance/internal/model"
	"strconv"
	"strings"
)

func TransferBalanceValidate(data model.TransferBalanceUser) bool {
	if data.Money > 0 {
		return true
	} else {
		return false
	}
}

func AddBalanceValidate(data model.BalanceAdd) bool {
	if data.Money > 0 {
		return true
	} else {
		return false
	}
}

func ReserveBalanceValidate(data model.ReserveCreate) bool {
	if data.Price >= 0 {
		return true
	} else {
		return false
	}
}

func ConformBalnceValidate(data model.ReserveConfirm) bool {
	if data.Command == "approved" || data.Command == "cancel" {
		return true
	} else {
		return false
	}
}

func ReportServiceValidate(data model.ReportServiceStruct) bool {
	year := strings.Split(data.Date, "-")[0]
	year_int, err := strconv.Atoi(year)
	if err != nil {
		return false
	}
	if year_int < 0 {
		return false
	}
	mounth := strings.Split(data.Date, "-")[1]
	mounth_int, err := strconv.Atoi(mounth)
	if err != nil {
		return false
	}
	if mounth_int < 0 || mounth_int > 12 {
		return false
	}
	return true
}

func ReportOperationValidate(data model.ReportOperationRequest) bool {
	if data.Page < 1 || data.Rows < 1 {
		return false
	}
	if data.Sort != "created_up" && data.Sort != "price_up" && data.Sort != "created_down" && data.Sort != "price_down" {
		return false
	}
	return true
}
