package models

type (
	Sms struct {
		VerifyUUId  *string `valid:"required" json:"verify_uuid,omitempty"`
		CountryCode *string `valid:"required" json:"country_code,omitempty"`
		Phone       *string `valid:"required" json:"phone,omitempty"`
	}
)

func (this *Sms) Send() {

}
