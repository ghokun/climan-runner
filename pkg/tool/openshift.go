package tool

// import (
// 	"errors"
// 	"io"
// 	"log"
// 	"net/http"
// )

// func init() {
// 	tools, err := getOpenshiftTools()
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	for _, t := range tools {
// 		Tools[t.Name] = t
// 	}
// }

// func getOpenshiftTools() (tools []Tool, err error) {
// 	response, err := http.Get("https://mirror.openshift.com/pub/openshift-v4/clients/ocp/latest/release.txt")
// 	if err != nil {
// 		return tools, err
// 	}
// 	defer response.Body.Close()
// 	if response.StatusCode == http.StatusOK {
// 		bodyBytes, err := io.ReadAll(response.Body)
// 		if err != nil {
// 			return tools, err
// 		}
// 		version := ""
// 		tools = append(tools, Tool{
// 			Name:        "oc",
// 			Description: "Openshift command line interface",
// 			Supports:    713,
// 			Latest:      version,
// 		})
// 		tools = append(tools, Tool{
// 			Name:        "openshift-install",
// 			Description: "Openshift installer",
// 			Supports:    713,
// 			Latest:      version,
// 		})
// 		tools = append(tools, Tool{
// 			Name:        "opm",
// 			Description: "Openshift installer",
// 			Supports:    713,
// 			Latest:      version,
// 		})
// 		return tools, nil
// 	}
// 	return tools, errors.New("error while fetcing latest version of openshift tools")
// }

// TODO all
// TODO versions
