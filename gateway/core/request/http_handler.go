package request

import (
	AppCore "api_gateway/gateway/core"
	AppAuth "api_gateway/gateway/core/auth"
	AppLogger "api_gateway/gateway/core/logger"
	"net/http"
	"fmt"
	"strings"
	"io/ioutil"
	"time"
)

func HttpHandler(w http.ResponseWriter, r *http.Request, router AppCore.Router) {
	originalPath := r.URL.Path
	// init logger
	logger := AppLogger.GetLogInstance()
	logger.InitLog(originalPath)

	service, _ := checkServiceExist(router, originalPath)

	msg, err := AppAuth.CheckAuth(r, service.TargetPath.Auth)
	if err != nil {
		AppLogger.DestroyLogInstance()

		AppCore.ShowError(w, err, http.StatusUnauthorized)
	}
	var req *http.Request

	defaultForwardPath := service.TargetPath

	req, err = createRequest(r, defaultForwardPath, originalPath, msg)
	if err != nil {
		logger.AddStep("HttpHandler", err.Error())
		AppLogger.DestroyLogInstance()

		AppCore.ShowError(w, err, http.StatusBadGateway)
	}

	fmt.Printf("forwarded to default :%v\n", req.URL)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		logger.AddStep("HttpHandler", err.Error())
		AppLogger.DestroyLogInstance()

		AppCore.ShowError(w, err, http.StatusBadGateway)
	}
	//
	defer resp.Body.Close()
	//
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logger.AddStep("HttpHandler", err.Error())
		AppLogger.DestroyLogInstance()
		AppCore.ShowError(w, err, http.StatusBadGateway)
	}

	headerResp := strings.Join(resp.Header["Content-Type"], "")
	w.Header().Set("Content-Type", headerResp)
	logger.AddStep("HttpHandler : Request Send Successfully", "")
	logger.EndTime = time.Now()
	logger.Status = true

	AppLogger.DestroyLogInstance()

	w.Write([]byte(body))

}

func createRequest(r *http.Request, forwardPath AppCore.TargetPath, originalReq string, msg string) (*http.Request, error) {
	logger := AppLogger.GetLogInstance()
	newPath := forwardPath.Path + originalReq
	req_content_type := r.Header.Get("Content-Type")
	req, err := http.NewRequest(r.Method, newPath, r.Body)
	if err != nil {
		logger.AddStep("createRequest", err.Error())
		return nil, err
	}

	req.Header.Set("Content-Type", req_content_type)
	req.Header.Set("msg-data", msg)
	req.Header.Set("Authorization", msg)

	logger.ForwardPath = newPath
	logger.AddStep("createRequest : Every Thing Is Good ", "")

	return req, nil
}

func checkServiceExist(router AppCore.Router, originalPath string) (AppCore.Services, error) {
	service_name_rray := strings.Split(originalPath, "/")
	service_prefix := service_name_rray[1]
	var service AppCore.Services
	var err error
	for _, v := range router.Services {

		if v.ServicePrefix == service_prefix {
			service = v
		}
	}
	if service.ServicePrefix == "" {
		err = fmt.Errorf("Service not Found")
	}
	// add loges steps
	errString := ""
	if err != nil {
		errString = err.Error()
	}

	logger := AppLogger.GetLogInstance()
	logger.AddStep("checkServiceExist : Every Thing Is Good", errString)
	return service, err

}
