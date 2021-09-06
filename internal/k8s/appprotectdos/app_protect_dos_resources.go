package appprotectdos

import (
	"fmt"
	"net"
	"regexp"
	"strconv"
	"strings"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
)

var appProtectDosPolicyRequiredFields = [][]string{
	{"spec"},
}

var appProtectDosLogConfRequiredFields = [][]string{
	{"spec", "content"},
	{"spec", "filter"},
}

func validateRequiredFields(policy *unstructured.Unstructured, fieldsList [][]string) error {
	for _, fields := range fieldsList {
		field, found, err := unstructured.NestedMap(policy.Object, fields...)
		if err != nil {
			return fmt.Errorf("Error checking for required field %v: %w", field, err)
		}
		if !found {
			return fmt.Errorf("Required field %v not found", field)
		}
	}
	return nil
}

// ValidateAppDosProtectLogConf validates LogConfiguration resource
func ValidateAppProtectDosLogConf(logConf *unstructured.Unstructured) error {
	lcName := logConf.GetName()
	err := validateRequiredFields(logConf, appProtectDosLogConfRequiredFields)
	if err != nil {
		return fmt.Errorf("Error validating App Protect Dos Log Configuration %v: %w", lcName, err)
	}

	return nil
}

var (
	logDstEx     = regexp.MustCompile(`(?:syslog:server=((?:\d{1,3}\.){3}\d{1,3}|localhost):\d{1,5})|stderr|(?:\/[\S]+)+`)
	logDstFileEx = regexp.MustCompile(`(?:\/[\S]+)+`)
)

// ValidateAppProtectDosLogDestination validates destination for log configuration
func ValidateAppProtectDosLogDestination(dstAntn string) error {
	errormsg := "Error parsing App Protect Dos Log config: Destination must follow format: syslog:server=<ip-address | localhost>:<port> or stderr or absolute path to file"
	if !logDstEx.MatchString(dstAntn) {
		return fmt.Errorf("%s Log Destination did not follow format", errormsg)
	}
	if dstAntn == "stderr" {
		return nil
	}

	if logDstFileEx.MatchString(dstAntn) {
		return nil
	}

	dstchunks := strings.Split(dstAntn, ":")

	// This error can be ignored since the regex check ensures this string will be parsable
	port, _ := strconv.Atoi(dstchunks[2])

	if port > 65535 || port < 1 {
		return fmt.Errorf("Error parsing port: %v not a valid port number", port)
	}

	ipstr := strings.Split(dstchunks[1], "=")[1]
	if ipstr == "localhost" {
		return nil
	}

	if net.ParseIP(ipstr) == nil {
		return fmt.Errorf("Error parsing host: %v is not a valid ip address", ipstr)
	}

	return nil
}

var accessLog = regexp.MustCompile(`^(((\d{1,3}\.){3}\d{1,3}):\d{1,5})$`)

// ValidateAppProtectDosAccessLog validates destination for access log configuration
func ValidateAppProtectDosAccessLogDest(accessLogDest string) error {
	errormsg := "Error parsing App Protect Dos Access Log Dest config: Destination must follow format: <ip-address>:<port>"
	if !accessLog.MatchString(accessLogDest) {
		return fmt.Errorf("%s Log Destination did not follow format", errormsg)
	}

	dstchunks := strings.Split(accessLogDest, ":")

	// This error can be ignored since the regex check ensures this string will be parsable
	port, _ := strconv.Atoi(dstchunks[1])

	if port > 65535 || port < 1 {
		return fmt.Errorf("Error parsing port: %v not a valid port number", port)
	}

	if net.ParseIP(dstchunks[0]) == nil {
		return fmt.Errorf("Error parsing host: %v is not a valid ip address", dstchunks[0])
	}

	return nil
}

// ValidateAppProtectDosPolicy validates Policy resource
func ValidateAppProtectDosPolicy(policy *unstructured.Unstructured) error {
	polName := policy.GetName()

	err := validateRequiredFields(policy, appProtectDosPolicyRequiredFields)
	if err != nil {
		return fmt.Errorf("Error validating App Protect Dos Policy %v: %w", polName, err)
	}

	return nil
}

// ParseResourceReferenceAnnotation returns a namespace/name string
func ParseResourceReferenceAnnotation(ns, antn string) string {
	if !strings.Contains(antn, "/") {
		return ns + "/" + antn
	}
	return antn
}

// GetNsName gets the key of a resource in the format: "resNamespace/resName"
func GetNsName(obj *unstructured.Unstructured) string {
	return obj.GetNamespace() + "/" + obj.GetName()
}
