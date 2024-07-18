package services

import (
	"errors"
	"fmt"
	"log"
	"strings"
	"webSocket_git/auth"
	"webSocket_git/register/models"
	"webSocket_git/register/repositories"
	"webSocket_git/register/transactions"
	"webSocket_git/utilts/encrypt"
	"webSocket_git/utilts/generate"
	"webSocket_git/utilts/sendEmailFunctions"

	"golang.org/x/crypto/bcrypt"
)

type ServicePort interface {
	RegisterChicCRMServices(loginData models.RegisterRequest) (models.RegisterResponses, error)
}

type serviceAdapter struct {
	r repositories.RepositoryPort
	t transactions.ITransaction // *** for log(mongoDB)
}

func NewServiceAdapter(r repositories.RepositoryPort, t transactions.ITransaction) ServicePort {
	// return &serviceAdapter{r: r}
	return &serviceAdapter{r: r, t: t}
}

func (s *serviceAdapter) RegisterChicCRMServices(loginData models.RegisterRequest) (models.RegisterResponses, error) {
	var err error
	var responses models.RegisterResponses
	const (
		keyUsername = "your_credential"
		keyPassword = "your_credential"
	)

	if len(loginData.Mobile_phone) != 10 {
		log.Printf(`{"status": "Error", "message": "Mobile phone must be 10 digits 089-XXX-XXXX"}`)
		return models.RegisterResponses{}, errors.New("mobile phone must be 10 digits 089-XXX-XXXX")
	}

	generatedPassword, err := generate.GenerateRandomPassword(8)
	if err != nil {
		log.Printf("Failed to generate password. Error:%v\n", err)
		return models.RegisterResponses{}, err
	}
	fmt.Println(generatedPassword)
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(generatedPassword), bcrypt.DefaultCost)
	if err != nil {
		log.Printf("Failed to hash password. Error:%v\n", err)
		return models.RegisterResponses{}, err
	}
	hashedPasswordString := string(hashedPassword)

	if !strings.Contains(loginData.Username, "@") || strings.HasPrefix(loginData.Username, "@") || strings.HasSuffix(loginData.Username, "@") || strings.Count(loginData.Username, "@") != 1 {
		return models.RegisterResponses{}, errors.New("username must be a valid email address")
	}
	emailDomain := loginData.Username
	splitEmail := strings.Split(emailDomain, "@")
	domain := "@" + splitEmail[1]
	// fmt.Println(splitEmail, splitEmail[1])

	token, err := auth.CreateTokenI(loginData.Username)
	if err != nil {
		log.Printf("Failed to create JWT for send email")
		return models.RegisterResponses{}, err
	}
	registrationComplete := make(chan bool)
	// defer close(registrationComplete) // do not use

	var (
		cipherProvince, cipherDistrict, cipherSubdistrict, cipherZipcode, cipherCreateLocation, cipherUrlLogo, cipherCompanyNameEN, cipherCompanyDomain, cipherCompanyMobile, cipherCompanyAlias, cipherTitle, cipherUsername, cipherFirstnameEn, cipherSurnamEn, cipherMobilePhone, cipherAddressNo, cipherAddress1En, cipherCompanyGeolo, cipherJobTitle, cipherWebsite string
		errChan                                                                                                                                                                                                                                                                                                                                                           = make(chan error, 20)
	)
	go func() {
		<-registrationComplete
		to := loginData.Username
		subject := "Welcome! You have successfully registered."
		body := "Please click the link provided below to Login<br>" +
			"Email: " + loginData.Username + "<br>" +
			"<a href='your_domain_name" + token + "'>Confirm Link</a><br>"
		if err := sendEmailFunctions.SendEmailRegister(to, subject, body); err != nil {
			log.Printf("เกิดข้อผิดพลาดในการส่งอีเมล: %s", err.Error())
			// notifyAdmin("เกิดข้อผิดพลาดในการส่งอีเมล: " + err.Error()) // ฟังก์ชันสมมุติเพื่อส่งการแจ้งเตือน
			return
		}
	}()

	fmt.Println(token)
	go func() {
		var err error
		cipherProvince, err = encrypt.SendToFortanixSDKMSTokenization(loginData.Province, keyUsername, keyPassword)
		errChan <- err
	}()
	go func() {
		var err error
		cipherDistrict, err = encrypt.SendToFortanixSDKMSTokenization(loginData.District, keyUsername, keyPassword)
		errChan <- err
	}()
	go func() {
		var err error
		cipherSubdistrict, err = encrypt.SendToFortanixSDKMSTokenization(loginData.Sub_district, keyUsername, keyPassword)
		errChan <- err
	}()
	go func() {
		var err error
		cipherZipcode, err = encrypt.SendToFortanixSDKMSTokenization(loginData.Zipcode, keyUsername, keyPassword)
		errChan <- err
	}()
	go func() {
		var err error
		cipherCreateLocation, err = encrypt.SendToFortanixSDKMSTokenization(loginData.Create_location, keyUsername, keyPassword)
		errChan <- err
	}()
	go func() {
		var err error
		cipherUrlLogo, err = encrypt.SendToFortanixSDKMSTokenization(loginData.Url_logo, keyUsername, keyPassword)
		errChan <- err
	}()
	go func() {
		var err error
		cipherCompanyNameEN, err = encrypt.SendToFortanixSDKMSTokenization(loginData.Company_name_en, keyUsername, keyPassword)
		errChan <- err
	}()
	go func() {
		var err error
		cipherCompanyDomain, err = encrypt.SendToFortanixSDKMSTokenization(domain, keyUsername, keyPassword)
		errChan <- err
	}()
	go func() {
		var err error
		cipherCompanyMobile, err = encrypt.SendToFortanixSDKMSTokenization(loginData.Company_mobile, keyUsername, keyPassword)
		errChan <- err
	}()
	go func() {
		var err error
		cipherCompanyAlias, err = encrypt.SendToFortanixSDKMSTokenization(loginData.Company_alias, keyUsername, keyPassword)
		errChan <- err
	}()
	go func() {
		var err error
		cipherTitle, err = encrypt.SendToFortanixSDKMSTokenization(loginData.Title, keyUsername, keyPassword)
		errChan <- err
	}()
	go func() {
		var err error
		cipherUsername, err = encrypt.SendToFortanixSDKMSTokenizationEmailForMasking(loginData.Username, keyUsername, keyPassword)
		errChan <- err
	}()
	go func() {
		var err error
		cipherFirstnameEn, err = encrypt.SendToFortanixSDKMSTokenization(loginData.Firstname_en, keyUsername, keyPassword)
		errChan <- err
	}()
	go func() {
		var err error
		cipherSurnamEn, err = encrypt.SendToFortanixSDKMSTokenization(loginData.Surname_en, keyUsername, keyPassword)
		errChan <- err
	}()
	go func() {
		var err error
		cipherMobilePhone, err = encrypt.SendToFortanixSDKMSTokenizationPhoneForMasking(loginData.Mobile_phone, keyUsername, keyPassword)
		errChan <- err
	}()
	go func() {
		var err error
		cipherAddressNo, err = encrypt.SendToFortanixSDKMSTokenization(loginData.Address_no, keyUsername, keyPassword)
		errChan <- err
	}()
	go func() {
		var err error
		cipherAddress1En, err = encrypt.SendToFortanixSDKMSTokenization(loginData.Address1_en, keyUsername, keyPassword)
		errChan <- err
	}()
	go func() {
		var err error
		cipherCompanyGeolo, err = encrypt.SendToFortanixSDKMSTokenization(loginData.Company_geolo, keyUsername, keyPassword)
		errChan <- err
	}()
	go func() {
		var err error
		cipherJobTitle, err = encrypt.SendToFortanixSDKMSTokenization(loginData.Job_title, keyUsername, keyPassword)
		errChan <- err
	}()
	go func() {
		var err error
		cipherWebsite, err = encrypt.SendToFortanixSDKMSTokenization(loginData.Website, keyUsername, keyPassword)
		errChan <- err
	}()
	// go func() { // not use yet remain for references for update
	// 	var err error
	// 	cipherUpdateLocation, err = encrypt.SendToFortanixSDKMSTokenization(loginData.Update_location, keyUsername, keyPassword)
	// 	errChan <- err
	// }()
	for i := 0; i < 20; i++ {
		if err := <-errChan; err != nil {
			log.Printf("variable not fullfill. Error:%v", err)
			return models.RegisterResponses{}, errors.New("please fill in all the required information")
		}
	}
	encryptData := models.EncryptedRegisterRequest{
		// CipherUpdate_Location: cipherUpdateLocation, // not use yet
		CipherProvince:        cipherProvince,
		CipherDistrict:        cipherDistrict,
		CipherSub_district:    cipherSubdistrict,
		CipherZipcode:         cipherZipcode,
		CipherCreate_Location: cipherCreateLocation,
		CipherUrl_logo:        cipherUrlLogo,
		CipherCompany_name_en: cipherCompanyNameEN,
		CipherCompany_domain:  cipherCompanyDomain,
		CipherCompany_mobile:  cipherCompanyMobile,
		CipherCompany_alias:   cipherCompanyAlias,
		CipherTitle:           cipherTitle,
		CipherUsername:        cipherUsername,
		CipherFirstname_en:    cipherFirstnameEn,
		CipherSurname_en:      cipherSurnamEn,
		CipherMobile_phone:    cipherMobilePhone,
		CipherAddress_no:      cipherAddressNo,
		CipherAddress1_en:     cipherAddress1En,
		CipherCompany_geolo:   cipherCompanyGeolo,
		CipherJob_title:       cipherJobTitle,
		HashPassword:          hashedPasswordString,
		CipherWebsite:         cipherWebsite,
	}
	if loginData.Job_title != "Manager" {
		responses, err = s.r.RegisterChicCRMSRepositoris(encryptData, loginData)
		if err != nil {
			log.Printf("Failed at RegisterChiccrmRepositories. Error:%v", err)
			return responses, err
		}
	} else {
		responses, err = s.r.RegisterChicCRMSRepositorisCase2(encryptData, loginData)
		if err != nil {
			log.Printf("Failed at RegisterChiccrmRepositoriesCase2. Error:%v", err)
			return responses, err
		}
		// return responses, nil // **fix email not send if manager == true
	}
	err = s.t.RegisterLogTransactions(loginData.Firstname_en, loginData.Surname_en, "Register successfully", loginData.Job_title, responses.CompanyID)
	if err != nil {
		log.Printf("Failed to insert transaction into acessLog(mongoDB). Error: %v", err)
		return responses, err
	}
	registrationComplete <- true
	return responses, nil
}
