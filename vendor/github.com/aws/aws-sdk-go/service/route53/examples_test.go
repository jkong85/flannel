// THIS FILE IS AUTOMATICALLY GENERATED. DO NOT EDIT.

package route53_test

import (
	"bytes"
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/route53"
)

var _ time.Duration
var _ bytes.Buffer

func ExampleRoute53_AssociateVPCWithHostedZone() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := route53.New(sess)

	params := &route53.AssociateVPCWithHostedZoneInput{
		HostedZoneId: aws.String("ResourceId"), // Required
		VPC: &route53.VPC{ // Required
			VPCId:     aws.String("VPCId"),
			VPCRegion: aws.String("VPCRegion"),
		},
		Comment: aws.String("AssociateVPCComment"),
	}
	resp, err := svc.AssociateVPCWithHostedZone(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleRoute53_ChangeResourceRecordSets() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := route53.New(sess)

	params := &route53.ChangeResourceRecordSetsInput{
		ChangeBatch: &route53.ChangeBatch{ // Required
			Changes: []*route53.Change{ // Required
				{ // Required
					Action: aws.String("ChangeAction"), // Required
					ResourceRecordSet: &route53.ResourceRecordSet{ // Required
						Name: aws.String("DNSName"), // Required
						Type: aws.String("RRType"),  // Required
						AliasTarget: &route53.AliasTarget{
							DNSName:              aws.String("DNSName"),    // Required
							EvaluateTargetHealth: aws.Bool(true),           // Required
							HostedZoneId:         aws.String("ResourceId"), // Required
						},
						Failover: aws.String("ResourceRecordSetFailover"),
						GeoLocation: &route53.GeoLocation{
							ContinentCode:   aws.String("GeoLocationContinentCode"),
							CountryCode:     aws.String("GeoLocationCountryCode"),
							SubdivisionCode: aws.String("GeoLocationSubdivisionCode"),
						},
						HealthCheckId: aws.String("HealthCheckId"),
						Region:        aws.String("ResourceRecordSetRegion"),
						ResourceRecords: []*route53.ResourceRecord{
							{ // Required
								Value: aws.String("RData"), // Required
							},
							// More values...
						},
						SetIdentifier: aws.String("ResourceRecordSetIdentifier"),
						TTL:           aws.Int64(1),
						TrafficPolicyInstanceId: aws.String("TrafficPolicyInstanceId"),
						Weight:                  aws.Int64(1),
					},
				},
				// More values...
			},
			Comment: aws.String("ResourceDescription"),
		},
		HostedZoneId: aws.String("ResourceId"), // Required
	}
	resp, err := svc.ChangeResourceRecordSets(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleRoute53_ChangeTagsForResource() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := route53.New(sess)

	params := &route53.ChangeTagsForResourceInput{
		ResourceId:   aws.String("TagResourceId"),   // Required
		ResourceType: aws.String("TagResourceType"), // Required
		AddTags: []*route53.Tag{
			{ // Required
				Key:   aws.String("TagKey"),
				Value: aws.String("TagValue"),
			},
			// More values...
		},
		RemoveTagKeys: []*string{
			aws.String("TagKey"), // Required
			// More values...
		},
	}
	resp, err := svc.ChangeTagsForResource(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleRoute53_CreateHealthCheck() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := route53.New(sess)

	params := &route53.CreateHealthCheckInput{
		CallerReference: aws.String("HealthCheckNonce"), // Required
		HealthCheckConfig: &route53.HealthCheckConfig{ // Required
			Type: aws.String("HealthCheckType"), // Required
			AlarmIdentifier: &route53.AlarmIdentifier{
				Name:   aws.String("AlarmName"),        // Required
				Region: aws.String("CloudWatchRegion"), // Required
			},
			ChildHealthChecks: []*string{
				aws.String("HealthCheckId"), // Required
				// More values...
			},
			EnableSNI:                    aws.Bool(true),
			FailureThreshold:             aws.Int64(1),
			FullyQualifiedDomainName:     aws.String("FullyQualifiedDomainName"),
			HealthThreshold:              aws.Int64(1),
			IPAddress:                    aws.String("IPAddress"),
			InsufficientDataHealthStatus: aws.String("InsufficientDataHealthStatus"),
			Inverted:                     aws.Bool(true),
			MeasureLatency:               aws.Bool(true),
			Port:                         aws.Int64(1),
			Regions: []*string{
				aws.String("HealthCheckRegion"), // Required
				// More values...
			},
			RequestInterval: aws.Int64(1),
			ResourcePath:    aws.String("ResourcePath"),
			SearchString:    aws.String("SearchString"),
		},
	}
	resp, err := svc.CreateHealthCheck(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleRoute53_CreateHostedZone() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := route53.New(sess)

	params := &route53.CreateHostedZoneInput{
		CallerReference: aws.String("Nonce"),   // Required
		Name:            aws.String("DNSName"), // Required
		DelegationSetId: aws.String("ResourceId"),
		HostedZoneConfig: &route53.HostedZoneConfig{
			Comment:     aws.String("ResourceDescription"),
			PrivateZone: aws.Bool(true),
		},
		VPC: &route53.VPC{
			VPCId:     aws.String("VPCId"),
			VPCRegion: aws.String("VPCRegion"),
		},
	}
	resp, err := svc.CreateHostedZone(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleRoute53_CreateReusableDelegationSet() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := route53.New(sess)

	params := &route53.CreateReusableDelegationSetInput{
		CallerReference: aws.String("Nonce"), // Required
		HostedZoneId:    aws.String("ResourceId"),
	}
	resp, err := svc.CreateReusableDelegationSet(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleRoute53_CreateTrafficPolicy() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := route53.New(sess)

	params := &route53.CreateTrafficPolicyInput{
		Document: aws.String("TrafficPolicyDocument"), // Required
		Name:     aws.String("TrafficPolicyName"),     // Required
		Comment:  aws.String("TrafficPolicyComment"),
	}
	resp, err := svc.CreateTrafficPolicy(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleRoute53_CreateTrafficPolicyInstance() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := route53.New(sess)

	params := &route53.CreateTrafficPolicyInstanceInput{
		HostedZoneId:         aws.String("ResourceId"),      // Required
		Name:                 aws.String("DNSName"),         // Required
		TTL:                  aws.Int64(1),                  // Required
		TrafficPolicyId:      aws.String("TrafficPolicyId"), // Required
		TrafficPolicyVersion: aws.Int64(1),                  // Required
	}
	resp, err := svc.CreateTrafficPolicyInstance(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleRoute53_CreateTrafficPolicyVersion() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := route53.New(sess)

	params := &route53.CreateTrafficPolicyVersionInput{
		Document: aws.String("TrafficPolicyDocument"), // Required
		Id:       aws.String("TrafficPolicyId"),       // Required
		Comment:  aws.String("TrafficPolicyComment"),
	}
	resp, err := svc.CreateTrafficPolicyVersion(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleRoute53_CreateVPCAssociationAuthorization() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := route53.New(sess)

	params := &route53.CreateVPCAssociationAuthorizationInput{
		HostedZoneId: aws.String("ResourceId"), // Required
		VPC: &route53.VPC{ // Required
			VPCId:     aws.String("VPCId"),
			VPCRegion: aws.String("VPCRegion"),
		},
	}
	resp, err := svc.CreateVPCAssociationAuthorization(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleRoute53_DeleteHealthCheck() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := route53.New(sess)

	params := &route53.DeleteHealthCheckInput{
		HealthCheckId: aws.String("HealthCheckId"), // Required
	}
	resp, err := svc.DeleteHealthCheck(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleRoute53_DeleteHostedZone() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := route53.New(sess)

	params := &route53.DeleteHostedZoneInput{
		Id: aws.String("ResourceId"), // Required
	}
	resp, err := svc.DeleteHostedZone(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleRoute53_DeleteReusableDelegationSet() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := route53.New(sess)

	params := &route53.DeleteReusableDelegationSetInput{
		Id: aws.String("ResourceId"), // Required
	}
	resp, err := svc.DeleteReusableDelegationSet(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleRoute53_DeleteTrafficPolicy() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := route53.New(sess)

	params := &route53.DeleteTrafficPolicyInput{
		Id:      aws.String("TrafficPolicyId"), // Required
		Version: aws.Int64(1),                  // Required
	}
	resp, err := svc.DeleteTrafficPolicy(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleRoute53_DeleteTrafficPolicyInstance() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := route53.New(sess)

	params := &route53.DeleteTrafficPolicyInstanceInput{
		Id: aws.String("TrafficPolicyInstanceId"), // Required
	}
	resp, err := svc.DeleteTrafficPolicyInstance(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleRoute53_DeleteVPCAssociationAuthorization() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := route53.New(sess)

	params := &route53.DeleteVPCAssociationAuthorizationInput{
		HostedZoneId: aws.String("ResourceId"), // Required
		VPC: &route53.VPC{ // Required
			VPCId:     aws.String("VPCId"),
			VPCRegion: aws.String("VPCRegion"),
		},
	}
	resp, err := svc.DeleteVPCAssociationAuthorization(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleRoute53_DisassociateVPCFromHostedZone() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := route53.New(sess)

	params := &route53.DisassociateVPCFromHostedZoneInput{
		HostedZoneId: aws.String("ResourceId"), // Required
		VPC: &route53.VPC{ // Required
			VPCId:     aws.String("VPCId"),
			VPCRegion: aws.String("VPCRegion"),
		},
		Comment: aws.String("DisassociateVPCComment"),
	}
	resp, err := svc.DisassociateVPCFromHostedZone(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleRoute53_GetChange() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := route53.New(sess)

	params := &route53.GetChangeInput{
		Id: aws.String("ResourceId"), // Required
	}
	resp, err := svc.GetChange(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleRoute53_GetCheckerIpRanges() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := route53.New(sess)

	var params *route53.GetCheckerIpRangesInput
	resp, err := svc.GetCheckerIpRanges(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleRoute53_GetGeoLocation() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := route53.New(sess)

	params := &route53.GetGeoLocationInput{
		ContinentCode:   aws.String("GeoLocationContinentCode"),
		CountryCode:     aws.String("GeoLocationCountryCode"),
		SubdivisionCode: aws.String("GeoLocationSubdivisionCode"),
	}
	resp, err := svc.GetGeoLocation(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleRoute53_GetHealthCheck() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := route53.New(sess)

	params := &route53.GetHealthCheckInput{
		HealthCheckId: aws.String("HealthCheckId"), // Required
	}
	resp, err := svc.GetHealthCheck(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleRoute53_GetHealthCheckCount() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := route53.New(sess)

	var params *route53.GetHealthCheckCountInput
	resp, err := svc.GetHealthCheckCount(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleRoute53_GetHealthCheckLastFailureReason() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := route53.New(sess)

	params := &route53.GetHealthCheckLastFailureReasonInput{
		HealthCheckId: aws.String("HealthCheckId"), // Required
	}
	resp, err := svc.GetHealthCheckLastFailureReason(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleRoute53_GetHealthCheckStatus() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := route53.New(sess)

	params := &route53.GetHealthCheckStatusInput{
		HealthCheckId: aws.String("HealthCheckId"), // Required
	}
	resp, err := svc.GetHealthCheckStatus(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleRoute53_GetHostedZone() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := route53.New(sess)

	params := &route53.GetHostedZoneInput{
		Id: aws.String("ResourceId"), // Required
	}
	resp, err := svc.GetHostedZone(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleRoute53_GetHostedZoneCount() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := route53.New(sess)

	var params *route53.GetHostedZoneCountInput
	resp, err := svc.GetHostedZoneCount(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleRoute53_GetReusableDelegationSet() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := route53.New(sess)

	params := &route53.GetReusableDelegationSetInput{
		Id: aws.String("ResourceId"), // Required
	}
	resp, err := svc.GetReusableDelegationSet(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleRoute53_GetTrafficPolicy() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := route53.New(sess)

	params := &route53.GetTrafficPolicyInput{
		Id:      aws.String("TrafficPolicyId"), // Required
		Version: aws.Int64(1),                  // Required
	}
	resp, err := svc.GetTrafficPolicy(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleRoute53_GetTrafficPolicyInstance() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := route53.New(sess)

	params := &route53.GetTrafficPolicyInstanceInput{
		Id: aws.String("TrafficPolicyInstanceId"), // Required
	}
	resp, err := svc.GetTrafficPolicyInstance(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleRoute53_GetTrafficPolicyInstanceCount() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := route53.New(sess)

	var params *route53.GetTrafficPolicyInstanceCountInput
	resp, err := svc.GetTrafficPolicyInstanceCount(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleRoute53_ListGeoLocations() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := route53.New(sess)

	params := &route53.ListGeoLocationsInput{
		MaxItems:             aws.String("PageMaxItems"),
		StartContinentCode:   aws.String("GeoLocationContinentCode"),
		StartCountryCode:     aws.String("GeoLocationCountryCode"),
		StartSubdivisionCode: aws.String("GeoLocationSubdivisionCode"),
	}
	resp, err := svc.ListGeoLocations(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleRoute53_ListHealthChecks() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := route53.New(sess)

	params := &route53.ListHealthChecksInput{
		Marker:   aws.String("PageMarker"),
		MaxItems: aws.String("PageMaxItems"),
	}
	resp, err := svc.ListHealthChecks(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleRoute53_ListHostedZones() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := route53.New(sess)

	params := &route53.ListHostedZonesInput{
		DelegationSetId: aws.String("ResourceId"),
		Marker:          aws.String("PageMarker"),
		MaxItems:        aws.String("PageMaxItems"),
	}
	resp, err := svc.ListHostedZones(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleRoute53_ListHostedZonesByName() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := route53.New(sess)

	params := &route53.ListHostedZonesByNameInput{
		DNSName:      aws.String("DNSName"),
		HostedZoneId: aws.String("ResourceId"),
		MaxItems:     aws.String("PageMaxItems"),
	}
	resp, err := svc.ListHostedZonesByName(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleRoute53_ListResourceRecordSets() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := route53.New(sess)

	params := &route53.ListResourceRecordSetsInput{
		HostedZoneId:          aws.String("ResourceId"), // Required
		MaxItems:              aws.String("PageMaxItems"),
		StartRecordIdentifier: aws.String("ResourceRecordSetIdentifier"),
		StartRecordName:       aws.String("DNSName"),
		StartRecordType:       aws.String("RRType"),
	}
	resp, err := svc.ListResourceRecordSets(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleRoute53_ListReusableDelegationSets() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := route53.New(sess)

	params := &route53.ListReusableDelegationSetsInput{
		Marker:   aws.String("PageMarker"),
		MaxItems: aws.String("PageMaxItems"),
	}
	resp, err := svc.ListReusableDelegationSets(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleRoute53_ListTagsForResource() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := route53.New(sess)

	params := &route53.ListTagsForResourceInput{
		ResourceId:   aws.String("TagResourceId"),   // Required
		ResourceType: aws.String("TagResourceType"), // Required
	}
	resp, err := svc.ListTagsForResource(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleRoute53_ListTagsForResources() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := route53.New(sess)

	params := &route53.ListTagsForResourcesInput{
		ResourceIds: []*string{ // Required
			aws.String("TagResourceId"), // Required
			// More values...
		},
		ResourceType: aws.String("TagResourceType"), // Required
	}
	resp, err := svc.ListTagsForResources(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleRoute53_ListTrafficPolicies() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := route53.New(sess)

	params := &route53.ListTrafficPoliciesInput{
		MaxItems:              aws.String("PageMaxItems"),
		TrafficPolicyIdMarker: aws.String("TrafficPolicyId"),
	}
	resp, err := svc.ListTrafficPolicies(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleRoute53_ListTrafficPolicyInstances() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := route53.New(sess)

	params := &route53.ListTrafficPolicyInstancesInput{
		HostedZoneIdMarker:              aws.String("ResourceId"),
		MaxItems:                        aws.String("PageMaxItems"),
		TrafficPolicyInstanceNameMarker: aws.String("DNSName"),
		TrafficPolicyInstanceTypeMarker: aws.String("RRType"),
	}
	resp, err := svc.ListTrafficPolicyInstances(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleRoute53_ListTrafficPolicyInstancesByHostedZone() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := route53.New(sess)

	params := &route53.ListTrafficPolicyInstancesByHostedZoneInput{
		HostedZoneId:                    aws.String("ResourceId"), // Required
		MaxItems:                        aws.String("PageMaxItems"),
		TrafficPolicyInstanceNameMarker: aws.String("DNSName"),
		TrafficPolicyInstanceTypeMarker: aws.String("RRType"),
	}
	resp, err := svc.ListTrafficPolicyInstancesByHostedZone(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleRoute53_ListTrafficPolicyInstancesByPolicy() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := route53.New(sess)

	params := &route53.ListTrafficPolicyInstancesByPolicyInput{
		TrafficPolicyId:                 aws.String("TrafficPolicyId"), // Required
		TrafficPolicyVersion:            aws.Int64(1),                  // Required
		HostedZoneIdMarker:              aws.String("ResourceId"),
		MaxItems:                        aws.String("PageMaxItems"),
		TrafficPolicyInstanceNameMarker: aws.String("DNSName"),
		TrafficPolicyInstanceTypeMarker: aws.String("RRType"),
	}
	resp, err := svc.ListTrafficPolicyInstancesByPolicy(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleRoute53_ListTrafficPolicyVersions() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := route53.New(sess)

	params := &route53.ListTrafficPolicyVersionsInput{
		Id:                         aws.String("TrafficPolicyId"), // Required
		MaxItems:                   aws.String("PageMaxItems"),
		TrafficPolicyVersionMarker: aws.String("TrafficPolicyVersionMarker"),
	}
	resp, err := svc.ListTrafficPolicyVersions(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleRoute53_ListVPCAssociationAuthorizations() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := route53.New(sess)

	params := &route53.ListVPCAssociationAuthorizationsInput{
		HostedZoneId: aws.String("ResourceId"), // Required
		MaxResults:   aws.String("MaxResults"),
		NextToken:    aws.String("PaginationToken"),
	}
	resp, err := svc.ListVPCAssociationAuthorizations(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleRoute53_TestDNSAnswer() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := route53.New(sess)

	params := &route53.TestDNSAnswerInput{
		HostedZoneId:          aws.String("ResourceId"), // Required
		RecordName:            aws.String("DNSName"),    // Required
		RecordType:            aws.String("RRType"),     // Required
		EDNS0ClientSubnetIP:   aws.String("IPAddress"),
		EDNS0ClientSubnetMask: aws.String("SubnetMask"),
		ResolverIP:            aws.String("IPAddress"),
	}
	resp, err := svc.TestDNSAnswer(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleRoute53_UpdateHealthCheck() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := route53.New(sess)

	params := &route53.UpdateHealthCheckInput{
		HealthCheckId: aws.String("HealthCheckId"), // Required
		AlarmIdentifier: &route53.AlarmIdentifier{
			Name:   aws.String("AlarmName"),        // Required
			Region: aws.String("CloudWatchRegion"), // Required
		},
		ChildHealthChecks: []*string{
			aws.String("HealthCheckId"), // Required
			// More values...
		},
		EnableSNI:                    aws.Bool(true),
		FailureThreshold:             aws.Int64(1),
		FullyQualifiedDomainName:     aws.String("FullyQualifiedDomainName"),
		HealthCheckVersion:           aws.Int64(1),
		HealthThreshold:              aws.Int64(1),
		IPAddress:                    aws.String("IPAddress"),
		InsufficientDataHealthStatus: aws.String("InsufficientDataHealthStatus"),
		Inverted:                     aws.Bool(true),
		Port:                         aws.Int64(1),
		Regions: []*string{
			aws.String("HealthCheckRegion"), // Required
			// More values...
		},
		ResourcePath: aws.String("ResourcePath"),
		SearchString: aws.String("SearchString"),
	}
	resp, err := svc.UpdateHealthCheck(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleRoute53_UpdateHostedZoneComment() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := route53.New(sess)

	params := &route53.UpdateHostedZoneCommentInput{
		Id:      aws.String("ResourceId"), // Required
		Comment: aws.String("ResourceDescription"),
	}
	resp, err := svc.UpdateHostedZoneComment(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleRoute53_UpdateTrafficPolicyComment() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := route53.New(sess)

	params := &route53.UpdateTrafficPolicyCommentInput{
		Comment: aws.String("TrafficPolicyComment"), // Required
		Id:      aws.String("TrafficPolicyId"),      // Required
		Version: aws.Int64(1),                       // Required
	}
	resp, err := svc.UpdateTrafficPolicyComment(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleRoute53_UpdateTrafficPolicyInstance() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := route53.New(sess)

	params := &route53.UpdateTrafficPolicyInstanceInput{
		Id:                   aws.String("TrafficPolicyInstanceId"), // Required
		TTL:                  aws.Int64(1),                          // Required
		TrafficPolicyId:      aws.String("TrafficPolicyId"),         // Required
		TrafficPolicyVersion: aws.Int64(1),                          // Required
	}
	resp, err := svc.UpdateTrafficPolicyInstance(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}
