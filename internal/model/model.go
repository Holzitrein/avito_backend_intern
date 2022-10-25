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
