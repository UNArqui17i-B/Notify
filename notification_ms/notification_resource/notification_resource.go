package notification_resource

import (
        "github.com/ant0ine/go-json-rest/rest"
        "net/http"
        "notification_ms/notification_model"
        "notification_ms/notification_service"
        "io/ioutil"
        "encoding/json"
)

func PostSendNotificationResource(w rest.ResponseWriter, req *rest.Request) {
		var t *notification_model.Notification
		body, err := ioutil.ReadAll(req.Body)
    	err = json.Unmarshal(body, &t)
    	if err != nil {
        	panic(err)
    	}
    	t.SetFileId(req.PathParam("file_id"))
    	user_auth := notification_model.EmailUser{"Blinkbox Project","blinkboxunal@gmail.com", "bl1nkb0x","smtp.gmail.com", 587}
    	notification_service.SendEmail_Service(t,&user_auth,"You have been selected to test blinkbox new feature")
		w.WriteJson(&t)	
		w.WriteHeader(http.StatusAccepted)
}