package models

type RegisterRequest struct {
	Username        string `json:"username" binding:"required"`
	Firstname_en    string `json:"firstname_en"`
	Surname_en      string `json:"surname_en"`
	Mobile_phone    string `json:"mobile_phone"`
	Company_name    string `json:"company_name"`
	Company_name_en string `json:"company_name_en"`
	Company_mobile  string `json:"company_mobile"`
	Company_alias   string `json:"company_alias"`
	Country         string `json:"country"`
	Province        string `json:"province"`
	District        string `json:"district"`
	Sub_district    string `json:"sub_district"`
	Zipcode         string `json:"zipcode"`
	Website         string `json:"website"`
	Address_no      string `json:"address_no"`
	Address1_en     string `json:"address1_en"`
	Role            string `json:"role"`
	Title           string `json:"title"`
	Job_title       string `json:"job_title"`
	Department      string `json:"department"`
	Create_location string `json:"create_location"`
	Url_logo        string `json:"url_logo"`
	Company_geolo   string `json:"company_geolo"`
}

type RegisterResponses struct { // can return whatever we want
	CompanyID string
}
