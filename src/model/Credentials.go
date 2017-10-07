package model


type Cred struct {
	Uid          string `json:"uid"`
	Url          string `json:"url"`
	SubmitUrl    string `json:"submitUrl"`
	Username     string `json:"username"`
	Password     string `json:"password"`
	Name         string `json:"name"`
	StringFields string `json:"stringFields"`
}

func (creds *Cred)Update(credsNew *Cred)(*Cred,error){

	if credsNew.Uid !=""{
		creds.Uid= credsNew.Uid
	}
	if credsNew.Url !=""{
		creds.Url = credsNew.Url
	}
	if credsNew.SubmitUrl !=""{
		creds.SubmitUrl = credsNew.SubmitUrl
	}
	if credsNew.Username !=""{
		creds.Username = credsNew.Username
	}
	if credsNew.Password != ""{
		creds.Password  = credsNew.Password
	}
	if credsNew.Name !=""{
		creds.Name = credsNew.Name
	}
	if credsNew.StringFields !=""{
		creds.StringFields = credsNew.StringFields
	}
	return creds,nil
}