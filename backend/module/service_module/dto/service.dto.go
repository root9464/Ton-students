package serv_dto

type ServiceType struct {
	UserId   int64         `json:"userId" validate:"required"`
	Price    float64       `json:"price" validate:"required"`
	Infos    []InfosType   `json:"infos"`
	Tags     *[]TagsType   `json:"tags"`
	Settings *SettingsType `json:"settings"`
}

type InfosType struct {
	Title   string `json:"title" validate:"required"`
	Content string `json:"content" validate:"required"`
}

type TagsType struct {
	ServiceId string `json:"serviceId" validate:"required"`
	Name      string `json:"name" validate:"required,max=10"`
}

type SettingsType struct {
	ColorHeader string  `json:"colorHeader"`
	ButtonText  *string `json:"buttonText"`

	IsPrepayment       bool `json:"isPrepayment"`
	IsDisabled         bool `json:"isDisabled"`
	IsAdditionalButton bool `json:"isAdditionalButton"`
}

type UpdateServiceType struct {
	ID    string   `json:"id" validate:"required"`
	Price *float64 `json:"price"`

	Infos    []InfosType   `json:"infos"`
	Tags     *[]TagsType   `json:"tags"`
	Settings *SettingsType `json:"settings"`
}
