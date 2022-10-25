package model

type BalanceGet struct {
	UserId int `json:"userId"`
}

type BalanceGetReturn struct {
	Balance float32 `json:"balance"`
}

type BalanceAdd struct {
	UserId int     `json:"userId"`
	Money  float32 `json:"money"`
	Note   string  `json:"note"`
}

type BalanceAddReturn struct {
	Status string `json:"status"`
}

type ReserveCreate struct {
	UserId    int     `json:"userId"`
	IdService int     `json:"idService"`
	IdOrder   int     `json:"idOrder"`
	Price     float32 `json:"price"`
}

type ReserveConfirm struct {
	UserId    int     `json:"userId"`
	IdService int     `json:"idService"`
	IdOrder   int     `json:"idOrder"`
	Price     float32 `json:"price"`
	Command   string  `json:"cmd"`
}

type ReserveCreateReturn struct {
	Status string `json:"status"`
}

type TransferBalanceUser struct {
	UserId  int     `json:"userId"`
	UserId2 int     `json:"userId2"`
	Money   float32 `json:"money"`
}

type TransferBalanceUserReturn struct {
	Status string `json:"status"`
}

type ReportServiceStruct struct {
	Date string `json:"date"`
}

type ReportServiceStructReturn struct {
	Url string `json:"url"`
}

type ServiceName struct {
	Id   int
	Name string
}
type ReportOperationRequest struct {
	Iduser int    `json:"iduser"`
	Page   int    `json:"page"`
	Sort   string `json:"sort"`
}
