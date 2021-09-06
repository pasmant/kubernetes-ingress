package appprotectdos

import (
	"strings"
	"testing"
)

func TestValidateAppProtectDsLogDestinationAnnotation(t *testing.T) {
	// Positive test cases
	posDstAntns := []string{"stderr", "syslog:server=localhost:9000", "syslog:server=10.1.1.2:9000", "/var/log/ap.log"}

	// Negative test cases item, expected error message
	negDstAntns := [][]string{
		{"stdout", "Log Destination did not follow format"},
		{"syslog:server=localhost:99999", "not a valid port number"},
		{"syslog:server=999.99.99.99:5678", "is not a valid ip address"},
	}

	for _, tCase := range posDstAntns {
		err := ValidateAppProtectDosLogDestination(tCase)
		if err != nil {
			t.Errorf("got %v expected nil", err)
		}
	}
	for _, nTCase := range negDstAntns {
		err := ValidateAppProtectDosLogDestination(nTCase[0])
		if err == nil {
			t.Errorf("got no error expected error containing %s", nTCase[1])
		} else {
			if !strings.Contains(err.Error(), nTCase[1]) {
				t.Errorf("got %v expected to contain: %s", err, nTCase[1])
			}
		}
	}
}

func TestValidateAppProtectDosAccessLogDest(t *testing.T) {
	// Positive test cases
	posDstAntns := []string{"10.10.1.1:514"}

	// Negative test cases item, expected error message
	negDstAntns := [][]string{
		{"NotValid", "Error parsing App Protect Dos Access Log Dest config: Destination must follow format: <ip-address>:<port> Log Destination did not follow format"},
		{"10.10.1.1:99999", "not a valid port number"},
		{"999.99.99.99:5678", "is not a valid ip address"},
	}

	for _, tCase := range posDstAntns {
		err := ValidateAppProtectDosAccessLogDest(tCase)
		if err != nil {
			t.Errorf("got %v expected nil", err)
		}
	}

	for _, nTCase := range negDstAntns {
		err := ValidateAppProtectDosAccessLogDest(nTCase[0])
		if err == nil {
			t.Errorf("got no error expected error containing %s", nTCase[1])
		} else {
			if !strings.Contains(err.Error(), nTCase[1]) {
				t.Errorf("got %v expected to contain: %s", err, nTCase[1])
			}
		}
	}
	
}
