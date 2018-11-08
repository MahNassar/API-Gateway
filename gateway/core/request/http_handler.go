package request

import (
	AppCore "api_gateway/gateway/core"
	AppAuth "api_gateway/gateway/core/auth"
	"net/http"
	"fmt"
	//"io/ioutil"
	"strings"
)

func HttpHandler(w http.ResponseWriter, r *http.Request, router AppCore.Router) {
	originalPath := r.URL.Path
	service_name_rray := strings.Split(originalPath, "/")
	service_prefix := service_name_rray[1]

	var service AppCore.Services

	for _, v := range router.Services {

		if v.ServicePrefix == service_prefix {
			service = v
		}
	}
	if service.ServicePrefix == "" {
		fmt.Printf("Service not defined\n")
		return
	}

	var req *http.Request

	defaultForwardPath := service.TargetPath

	req = convertRequest(r, defaultForwardPath, originalPath)

	fmt.Printf("forwarded to default :%v\n", req.URL)

	msg, err := AppAuth.CheckAuth(r, service.TargetPath.Auth)
	if err != nil {
		fmt.Printf("Not Auth Request")

		AppCore.ShowError(w, err, http.StatusUnauthorized)
	}
	fmt.Println(msg)
	//client := &http.Client{}
	//resp, err := client.Do(req)
	//AppCore.CheckErr(err)
	//
	//defer resp.Body.Close()
	//
	//body, err := ioutil.ReadAll(resp.Body)
	//AppCore.CheckErr(err)
	//
	////resp with original Content-Type
	//headerResp := strings.Join(resp.Header["Content-Type"], "")
	//w.Header().Set("Content-Type", headerResp)
	//w.Write([]byte(body))
	// fmt.Fprintf(w, string(body))

}

func convertRequest(r *http.Request, forwardPath AppCore.TargetPath, originalReq string) *http.Request {
	fmt.Println(forwardPath.Auth)
	newPath := forwardPath.Path + originalReq
	req_content_type := r.Header.Get("Content-Type")
	req, err := http.NewRequest(r.Method, newPath, r.Body)
	AppCore.CheckErr(err)
	fmt.Println("Req Done process")
	req.Header.Set("Content-Type", req_content_type)
	req.Header.Set("msg-data", "Mahmoud Nassar Programming yra7eb bekom")

	return req
}
