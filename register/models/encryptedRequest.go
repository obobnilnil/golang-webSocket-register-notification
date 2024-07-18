package models

type EncryptedRegisterRequest struct {
	// CipherUsername        string `json:"username" binding:"required"` // in case need binding
	CipherUsername        string
	CipherFirstname_en    string
	CipherSurname_en      string
	CipherMobile_phone    string
	CipherCompany_name    string
	CipherCompany_name_en string
	CipherCompany_mobile  string
	CipherCompany_domain  string
	CipherCompany_alias   string
	// CipherCountry         string
	CipherProvince     string
	CipherDistrict     string
	CipherSub_district string
	CipherZipcode      string
	CipherWebsite      string
	CipherAddress_no   string
	CipherAddress1_en  string
	// CipherRole            string
	CipherTitle           string
	CipherJob_title       string
	CipherCreate_Location string
	CipherUrl_logo        string
	CipherCompany_geolo   string
	HashPassword          string
}
