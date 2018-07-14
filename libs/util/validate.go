package util

import (
	"log"

	"gout/libs/setting"
)

func CheckClusterId(clusterId string) (bool, error) {
	form := map[string]string{
		"cluster_id": clusterId,
	}
	options := map[string]interface{}{
		"headers": map[string]string{
			"Content-Type": "application/json",
		},
		"body": form,
	}

	log.Printf("request license center url = %s", setting.LicenseRestValid)
	r := Request{URL: setting.LicenseRestValid, Options: options}
	json, err := r.Post()
	if err != nil {
		log.Println(err)
		return false, err
	}

	registered := json["registered"].(bool)
	expired := json["expired"].(bool)
	status := registered && !expired
	log.Printf("registered = %t, expired = %t", registered, expired)
	log.Printf("status = %t", status)

	return status, nil
}
