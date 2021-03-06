// Copyright 2019 Huawei Technologies Co.,Ltd.
// Licensed under the Apache License, Version 2.0 (the "License"); you may not use
// this file except in compliance with the License.  You may obtain a copy of the
// License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software distributed
// under the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR
// CONDITIONS OF ANY KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations under the License.

package obsClient

import (
	"fmt"
	"obsTest/examples"
	"os"
	"strings"
	"time"

	"github.com/Lincyaw/huaweicloud-sdk-go-obs/obs"
)

const (
	endpoint   = "https://obs.cn-south-1.myhuaweicloud.com"
	ak         = "LULDRGD41JYICKMEMDW8"
	sk         = "Rbq85F2XTSJdVRXPP2cmstc96xpHGkLXSoOrLNgF"
	bucketName = "bucket123234"
	objectKey  = "object345234"
	location   = "cn-south-1"
)

var obsClient *obs.ObsClient

func GetObsClient() *obs.ObsClient {
	var err error
	if obsClient == nil {
		obsClient, err = obs.New(ak, sk, endpoint)
		if err != nil {
			panic(err)
		}
	}
	return obsClient
}

func CreateBucket() {
	input := &obs.CreateBucketInput{}
	input.Bucket = bucketName
	input.Location = location
	_, err := GetObsClient().CreateBucket(input)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Create bucket:%s successfully!\n", bucketName)
	fmt.Println()
}

func ListBuckets() {
	input := &obs.ListBucketsInput{}
	input.QueryLocation = true
	output, err := GetObsClient().ListBuckets(input)
	if err == nil {
		fmt.Printf("StatusCode:%d, RequestId:%s\n", output.StatusCode, output.RequestId)
		fmt.Printf("Owner.DisplayName:%s, Owner.ID:%s\n", output.Owner.DisplayName, output.Owner.ID)
		for index, val := range output.Buckets {
			fmt.Printf("Bucket[%d]-Name:%s,CreationDate:%s,Location:%s\n", index, val.Name, val.CreationDate, val.Location)
		}
	} else {
		if obsError, ok := err.(obs.ObsError); ok {
			fmt.Println(obsError.StatusCode)
			fmt.Println(obsError.Code)
			fmt.Println(obsError.Message)
		} else {
			fmt.Println(err)
		}
	}
}

func SetBucketStoragePolicy() {
	input := &obs.SetBucketStoragePolicyInput{}
	input.Bucket = bucketName
	input.StorageClass = obs.StorageClassCold
	output, err := GetObsClient().SetBucketStoragePolicy(input)
	if err == nil {
		fmt.Printf("StatusCode:%d, RequestId:%s\n", output.StatusCode, output.RequestId)
	} else {
		if obsError, ok := err.(obs.ObsError); ok {
			fmt.Println(obsError.StatusCode)
			fmt.Println(obsError.Code)
			fmt.Println(obsError.Message)
		} else {
			fmt.Println(err)
		}
	}
}

func GetBucketStoragePolicy() {
	output, err := GetObsClient().GetBucketStoragePolicy(bucketName)
	if err == nil {
		fmt.Printf("StatusCode:%d, RequestId:%s\n", output.StatusCode, output.RequestId)
		fmt.Printf("StorageClass:%s\n", output.StorageClass)
	} else {
		if obsError, ok := err.(obs.ObsError); ok {
			fmt.Println(obsError.StatusCode)
			fmt.Println(obsError.Code)
			fmt.Println(obsError.Message)
		} else {
			fmt.Println(err)
		}
	}
}

func DeleteBucket() {
	output, err := GetObsClient().DeleteBucket(bucketName)
	if err == nil {
		fmt.Printf("StatusCode:%d, RequestId:%s\n", output.StatusCode, output.RequestId)
	} else {
		if obsError, ok := err.(obs.ObsError); ok {
			fmt.Println(obsError.StatusCode)
			fmt.Println(obsError.Code)
			fmt.Println(obsError.Message)
		} else {
			fmt.Println(err)
		}
	}
}

func ListObjects() {
	input := &obs.ListObjectsInput{}
	input.Bucket = bucketName
	input.MaxKeys = 10
	//	input.Prefix = "src/"
	output, err := GetObsClient().ListObjects(input)
	if err == nil {
		fmt.Printf("StatusCode:%d, RequestId:%s\n", output.StatusCode, output.RequestId)
		for index, val := range output.Contents {
			fmt.Printf("Content[%d]-OwnerId:%s, OwnerName:%s, ETag:%s, Key:%s, LastModified:%s, Size:%d, StorageClass:%s\n",
				index, val.Owner.ID, val.Owner.DisplayName, val.ETag, val.Key, val.LastModified, val.Size, val.StorageClass)
		}
	} else {
		if obsError, ok := err.(obs.ObsError); ok {
			fmt.Println(obsError.StatusCode)
			fmt.Println(obsError.Code)
			fmt.Println(obsError.Message)
		} else {
			fmt.Println(err)
		}
	}
}

func ListVersions() {
	input := &obs.ListVersionsInput{}
	input.Bucket = bucketName
	input.MaxKeys = 10
	//	input.Prefix = "src/"
	output, err := GetObsClient().ListVersions(input)
	if err == nil {
		fmt.Printf("StatusCode:%d, RequestId:%s\n", output.StatusCode, output.RequestId)
		for index, val := range output.Versions {
			fmt.Printf("Version[%d]-OwnerId:%s, OwnerName:%s, ETag:%s, Key:%s, VersionId:%s, LastModified:%s, Size:%d, StorageClass:%s\n",
				index, val.Owner.ID, val.Owner.DisplayName, val.ETag, val.Key, val.VersionId, val.LastModified, val.Size, val.StorageClass)
		}
		for index, val := range output.DeleteMarkers {
			fmt.Printf("DeleteMarker[%d]-OwnerId:%s, OwnerName:%s, Key:%s, VersionId:%s, LastModified:%s, StorageClass:%s\n",
				index, val.Owner.ID, val.Owner.DisplayName, val.Key, val.VersionId, val.LastModified, val.StorageClass)
		}
	} else {
		if obsError, ok := err.(obs.ObsError); ok {
			fmt.Println(obsError.StatusCode)
			fmt.Println(obsError.Code)
			fmt.Println(obsError.Message)
		} else {
			fmt.Println(err)
		}
	}
}

func SetBucketQuota() {
	input := &obs.SetBucketQuotaInput{}
	input.Bucket = bucketName
	input.Quota = 1024 * 1024 * 1024
	output, err := GetObsClient().SetBucketQuota(input)
	if err == nil {
		fmt.Printf("StatusCode:%d, RequestId:%s\n", output.StatusCode, output.RequestId)
	} else {
		if obsError, ok := err.(obs.ObsError); ok {
			fmt.Println(obsError.StatusCode)
			fmt.Println(obsError.Code)
			fmt.Println(obsError.Message)
		} else {
			fmt.Println(err)
		}
	}
}

func GetBucketQuota() {
	output, err := GetObsClient().GetBucketQuota(bucketName)
	if err == nil {
		fmt.Printf("StatusCode:%d, RequestId:%s\n", output.StatusCode, output.RequestId)
		fmt.Printf("Quota:%d\n", output.Quota)
	} else {
		if obsError, ok := err.(obs.ObsError); ok {
			fmt.Println(obsError.StatusCode)
			fmt.Println(obsError.Code)
			fmt.Println(obsError.Message)
		} else {
			fmt.Println(err)
		}
	}
}

func GetBucketStorageInfo() {
	output, err := GetObsClient().GetBucketStorageInfo(bucketName)
	if err == nil {
		fmt.Printf("StatusCode:%d, RequestId:%s\n", output.StatusCode, output.RequestId)
		fmt.Printf("Size:%d, ObjectNumber:%d\n", output.Size, output.ObjectNumber)
	} else {
		if obsError, ok := err.(obs.ObsError); ok {
			fmt.Println(obsError.StatusCode)
			fmt.Println(obsError.Code)
			fmt.Println(obsError.Message)
		} else {
			fmt.Println(err)
		}
	}
}

func GetBucketLocation() {
	output, err := GetObsClient().GetBucketLocation(bucketName)
	if err == nil {
		fmt.Printf("StatusCode:%d, RequestId:%s\n", output.StatusCode, output.RequestId)
		fmt.Printf("Location:%s\n", output.Location)
	} else {
		if obsError, ok := err.(obs.ObsError); ok {
			fmt.Println(obsError.StatusCode)
			fmt.Println(obsError.Code)
			fmt.Println(obsError.Message)
		} else {
			fmt.Println(err)
		}
	}
}

func SetBucketAcl() {
	input := &obs.SetBucketAclInput{}
	input.Bucket = bucketName
	//		input.ACL = obs.AclPublicRead
	input.Owner.ID = "ownerid"
	var grants [3]obs.Grant
	grants[0].Grantee.Type = obs.GranteeGroup
	grants[0].Grantee.URI = obs.GroupAuthenticatedUsers
	grants[0].Permission = obs.PermissionRead

	grants[1].Grantee.Type = obs.GranteeUser
	grants[1].Grantee.ID = "userid"
	grants[1].Permission = obs.PermissionWrite

	grants[2].Grantee.Type = obs.GranteeUser
	grants[2].Grantee.ID = "userid"
	grants[2].Grantee.DisplayName = "username"
	grants[2].Permission = obs.PermissionRead
	input.Grants = grants[0:3]
	output, err := GetObsClient().SetBucketAcl(input)
	if err == nil {
		fmt.Printf("StatusCode:%d, RequestId:%s\n", output.StatusCode, output.RequestId)
	} else {
		if obsError, ok := err.(obs.ObsError); ok {
			fmt.Println(obsError.StatusCode)
			fmt.Println(obsError.Code)
			fmt.Println(obsError.Message)
		} else {
			fmt.Println(err)
		}
	}
}

func GetBucketAcl() {
	output, err := GetObsClient().GetBucketAcl(bucketName)
	if err == nil {
		fmt.Printf("StatusCode:%d, RequestId:%s\n", output.StatusCode, output.RequestId)
		fmt.Printf("Owner.DisplayName:%s, Owner.ID:%s\n", output.Owner.DisplayName, output.Owner.ID)
		for index, grant := range output.Grants {
			fmt.Printf("Grant[%d]-Type:%s, ID:%s, URI:%s, Permission:%s\n", index, grant.Grantee.Type, grant.Grantee.ID, grant.Grantee.URI, grant.Permission)
		}
	} else {
		if obsError, ok := err.(obs.ObsError); ok {
			fmt.Println(obsError.StatusCode)
			fmt.Println(obsError.Code)
			fmt.Println(obsError.Message)
		} else {
			fmt.Println(err)
		}
	}
}

func SetBucketPolicy() {
	input := &obs.SetBucketPolicyInput{}
	input.Bucket = bucketName
	input.Policy = "your policy"
	output, err := GetObsClient().SetBucketPolicy(input)
	if err == nil {
		fmt.Printf("StatusCode:%d, RequestId:%s\n", output.StatusCode, output.RequestId)
	} else {
		if obsError, ok := err.(obs.ObsError); ok {
			fmt.Println(obsError.StatusCode)
			fmt.Println(obsError.Code)
			fmt.Println(obsError.Message)
		} else {
			fmt.Println(err)
		}
	}
}

func GetBucketPolicy() {
	output, err := GetObsClient().GetBucketPolicy(bucketName)
	if err == nil {
		fmt.Printf("StatusCode:%d, RequestId:%s\n", output.StatusCode, output.RequestId)
		fmt.Printf("Policy:%s\n", output.Policy)
	} else {
		if obsError, ok := err.(obs.ObsError); ok {
			fmt.Println(obsError.StatusCode)
			fmt.Println(obsError.Code)
			fmt.Println(obsError.Message)
		} else {
			fmt.Println(err)
		}
	}
}

func DeleteBucketPolicy() {
	output, err := GetObsClient().DeleteBucketPolicy(bucketName)
	if err == nil {
		fmt.Printf("StatusCode:%d, RequestId:%s\n", output.StatusCode, output.RequestId)
	} else {
		if obsError, ok := err.(obs.ObsError); ok {
			fmt.Println(obsError.StatusCode)
			fmt.Println(obsError.Code)
			fmt.Println(obsError.Message)
		} else {
			fmt.Println(err)
		}
	}
}

func SetBucketCors() {
	input := &obs.SetBucketCorsInput{}
	input.Bucket = bucketName

	var corsRules [2]obs.CorsRule
	corsRule0 := obs.CorsRule{}
	corsRule0.ID = "rule1"
	corsRule0.AllowedOrigin = []string{"http://www.a.com", "http://www.b.com"}
	corsRule0.AllowedMethod = []string{"GET", "PUT", "POST", "HEAD"}
	corsRule0.AllowedHeader = []string{"header1", "header2"}
	corsRule0.MaxAgeSeconds = 100
	corsRule0.ExposeHeader = []string{"obs-1", "obs-2"}
	corsRules[0] = corsRule0
	corsRule1 := obs.CorsRule{}

	corsRule1.ID = "rule2"
	corsRule1.AllowedOrigin = []string{"http://www.c.com", "http://www.d.com"}
	corsRule1.AllowedMethod = []string{"GET", "PUT", "POST", "HEAD"}
	corsRule1.AllowedHeader = []string{"header3", "header4"}
	corsRule1.MaxAgeSeconds = 50
	corsRule1.ExposeHeader = []string{"obs-3", "obs-4"}
	corsRules[1] = corsRule1
	input.CorsRules = corsRules[:]
	output, err := GetObsClient().SetBucketCors(input)
	if err == nil {
		fmt.Printf("StatusCode:%d, RequestId:%s\n", output.StatusCode, output.RequestId)
	} else {
		if obsError, ok := err.(obs.ObsError); ok {
			fmt.Println(obsError.StatusCode)
			fmt.Println(obsError.Code)
			fmt.Println(obsError.Message)
		} else {
			fmt.Println(err)
		}
	}
}

func GetBucketCors() {
	output, err := GetObsClient().GetBucketCors(bucketName)
	if err == nil {
		fmt.Printf("StatusCode:%d, RequestId:%s\n", output.StatusCode, output.RequestId)
		for _, corsRule := range output.CorsRules {
			fmt.Printf("ID:%s, AllowedOrigin:%s, AllowedMethod:%s, AllowedHeader:%s, MaxAgeSeconds:%d, ExposeHeader:%s\n",
				corsRule.ID, strings.Join(corsRule.AllowedOrigin, "|"), strings.Join(corsRule.AllowedMethod, "|"),
				strings.Join(corsRule.AllowedHeader, "|"), corsRule.MaxAgeSeconds, strings.Join(corsRule.ExposeHeader, "|"))
		}
	} else {
		if obsError, ok := err.(obs.ObsError); ok {
			fmt.Println(obsError.StatusCode)
			fmt.Println(obsError.Code)
			fmt.Println(obsError.Message)
		} else {
			fmt.Println(err)
		}
	}
}

func DeleteBucketCors() {
	output, err := GetObsClient().DeleteBucketCors(bucketName)
	if err == nil {
		fmt.Printf("StatusCode:%d, RequestId:%s\n", output.StatusCode, output.RequestId)
	} else {
		if obsError, ok := err.(obs.ObsError); ok {
			fmt.Println(obsError.StatusCode)
			fmt.Println(obsError.Code)
			fmt.Println(obsError.Message)
		} else {
			fmt.Println(err)
		}
	}
}

func SetBucketVersioning() {
	input := &obs.SetBucketVersioningInput{}
	input.Bucket = bucketName
	input.Status = obs.VersioningStatusEnabled
	output, err := GetObsClient().SetBucketVersioning(input)
	if err == nil {
		fmt.Printf("StatusCode:%d, RequestId:%s\n", output.StatusCode, output.RequestId)
	} else {
		if obsError, ok := err.(obs.ObsError); ok {
			fmt.Println(obsError.StatusCode)
			fmt.Println(obsError.Code)
			fmt.Println(obsError.Message)
		} else {
			fmt.Println(err)
		}
	}
}

func GetBucketVersioning() {
	output, err := GetObsClient().GetBucketVersioning(bucketName)
	if err == nil {
		fmt.Printf("StatusCode:%d, RequestId:%s\n", output.StatusCode, output.RequestId)
		fmt.Printf("Status:%s\n", output.Status)
	} else {
		if obsError, ok := err.(obs.ObsError); ok {
			fmt.Println(obsError.StatusCode)
			fmt.Println(obsError.Code)
			fmt.Println(obsError.Message)
		} else {
			fmt.Println(err)
		}
	}
}

func HeadBucket() {
	output, err := GetObsClient().HeadBucket(bucketName)
	if err == nil {
		fmt.Printf("StatusCode:%d, RequestId:%s\n", output.StatusCode, output.RequestId)
	} else {
		if obsError, ok := err.(obs.ObsError); ok {
			fmt.Println(obsError.StatusCode)
		} else {
			fmt.Println(err)
		}
	}
}

func GetBucketMetadata() {
	input := &obs.GetBucketMetadataInput{}
	input.Bucket = bucketName
	output, err := GetObsClient().GetBucketMetadata(input)
	if err == nil {
		fmt.Printf("StatusCode:%d, RequestId:%s\n", output.StatusCode, output.RequestId)
		fmt.Printf("StorageClass:%s\n", output.StorageClass)
	} else {
		if obsError, ok := err.(obs.ObsError); ok {
			fmt.Printf("StatusCode:%d\n", obsError.StatusCode)
		} else {
			fmt.Println(err)
		}
	}
}

func SetBucketLoggingConfiguration() {
	input := &obs.SetBucketLoggingConfigurationInput{}
	input.Bucket = bucketName
	input.TargetBucket = "target-bucket"
	input.TargetPrefix = "prefix"
	var grants [3]obs.Grant
	grants[0].Grantee.Type = obs.GranteeGroup
	grants[0].Grantee.URI = obs.GroupAuthenticatedUsers
	grants[0].Permission = obs.PermissionRead

	grants[1].Grantee.Type = obs.GranteeUser
	grants[1].Grantee.ID = "userid"
	grants[1].Permission = obs.PermissionWrite

	grants[2].Grantee.Type = obs.GranteeUser
	grants[2].Grantee.ID = "userid"
	grants[2].Grantee.DisplayName = "username"
	grants[2].Permission = obs.PermissionRead
	input.TargetGrants = grants[0:3]
	output, err := GetObsClient().SetBucketLoggingConfiguration(input)
	if err == nil {
		fmt.Printf("StatusCode:%d, RequestId:%s\n", output.StatusCode, output.RequestId)
	} else {
		if obsError, ok := err.(obs.ObsError); ok {
			fmt.Println(obsError.StatusCode)
			fmt.Println(obsError.Code)
			fmt.Println(obsError.Message)
		} else {
			fmt.Println(err)
		}
	}
}

func GetBucketLoggingConfiguration() {
	output, err := GetObsClient().GetBucketLoggingConfiguration(bucketName)
	if err == nil {
		fmt.Printf("StatusCode:%d, RequestId:%s\n", output.StatusCode, output.RequestId)
		fmt.Printf("TargetBucket:%s, TargetPrefix:%s\n", output.TargetBucket, output.TargetPrefix)
		for index, grant := range output.TargetGrants {
			fmt.Printf("Grant[%d]-Type:%s, ID:%s, URI:%s, Permission:%s\n", index, grant.Grantee.Type, grant.Grantee.ID, grant.Grantee.URI, grant.Permission)
		}
	} else {
		if obsError, ok := err.(obs.ObsError); ok {
			fmt.Println(obsError.StatusCode)
			fmt.Println(obsError.Code)
			fmt.Println(obsError.Message)
		} else {
			fmt.Println(err)
		}
	}
}

func SetBucketWebsiteConfiguration() {
	input := &obs.SetBucketWebsiteConfigurationInput{}
	input.Bucket = bucketName
	//	input.RedirectAllRequestsTo.HostName = "www.a.com"
	//	input.RedirectAllRequestsTo.Protocol = obs.ProtocolHttp
	input.IndexDocument.Suffix = "suffix"
	input.ErrorDocument.Key = "key"

	var routingRules [2]obs.RoutingRule
	routingRule0 := obs.RoutingRule{}

	routingRule0.Redirect.HostName = "www.a.com"
	routingRule0.Redirect.Protocol = obs.ProtocolHttp
	routingRule0.Redirect.ReplaceKeyPrefixWith = "prefix"
	routingRule0.Redirect.HttpRedirectCode = "304"
	routingRules[0] = routingRule0

	routingRule1 := obs.RoutingRule{}

	routingRule1.Redirect.HostName = "www.b.com"
	routingRule1.Redirect.Protocol = obs.ProtocolHttps
	routingRule1.Redirect.ReplaceKeyWith = "replaceKey"
	routingRule1.Redirect.HttpRedirectCode = "304"

	routingRule1.Condition.HttpErrorCodeReturnedEquals = "404"
	routingRule1.Condition.KeyPrefixEquals = "prefix"

	routingRules[1] = routingRule1

	input.RoutingRules = routingRules[:]
	output, err := GetObsClient().SetBucketWebsiteConfiguration(input)
	if err == nil {
		fmt.Printf("StatusCode:%d, RequestId:%s\n", output.StatusCode, output.RequestId)
	} else {
		if obsError, ok := err.(obs.ObsError); ok {
			fmt.Println(obsError.StatusCode)
			fmt.Println(obsError.Code)
			fmt.Println(obsError.Message)
		} else {
			fmt.Println(err)
		}
	}
}

func GetBucketWebsiteConfiguration() {
	output, err := GetObsClient().GetBucketWebsiteConfiguration(bucketName)
	if err == nil {
		fmt.Printf("StatusCode:%d, RequestId:%s\n", output.StatusCode, output.RequestId)
		fmt.Printf("RedirectAllRequestsTo.HostName:%s,RedirectAllRequestsTo.Protocol:%s\n", output.RedirectAllRequestsTo.HostName, output.RedirectAllRequestsTo.Protocol)
		fmt.Printf("Suffix:%s\n", output.IndexDocument.Suffix)
		fmt.Printf("Key:%s\n", output.ErrorDocument.Key)
		for index, routingRule := range output.RoutingRules {
			fmt.Printf("Condition[%d]-KeyPrefixEquals:%s, HttpErrorCodeReturnedEquals:%s\n", index, routingRule.Condition.KeyPrefixEquals, routingRule.Condition.HttpErrorCodeReturnedEquals)
			fmt.Printf("Redirect[%d]-Protocol:%s, HostName:%s, ReplaceKeyPrefixWith:%s, HttpRedirectCode:%s\n",
				index, routingRule.Redirect.Protocol, routingRule.Redirect.HostName, routingRule.Redirect.ReplaceKeyPrefixWith, routingRule.Redirect.HttpRedirectCode)
		}
	} else {
		if obsError, ok := err.(obs.ObsError); ok {
			fmt.Println(obsError.StatusCode)
			fmt.Println(obsError.Code)
			fmt.Println(obsError.Message)
		} else {
			fmt.Println(err)
		}
	}
}

func DeleteBucketWebsiteConfiguration() {
	output, err := GetObsClient().DeleteBucketWebsiteConfiguration(bucketName)
	if err == nil {
		fmt.Printf("StatusCode:%d, RequestId:%s\n", output.StatusCode, output.RequestId)
	} else {
		if obsError, ok := err.(obs.ObsError); ok {
			fmt.Println(obsError.StatusCode)
			fmt.Println(obsError.Code)
			fmt.Println(obsError.Message)
		} else {
			fmt.Println(err)
		}
	}
}

func SetBucketLifecycleConfiguration() {
	input := &obs.SetBucketLifecycleConfigurationInput{}
	input.Bucket = bucketName

	var lifecycleRules [2]obs.LifecycleRule
	lifecycleRule0 := obs.LifecycleRule{}
	lifecycleRule0.ID = "rule0"
	lifecycleRule0.Prefix = "prefix0"
	lifecycleRule0.Status = obs.RuleStatusEnabled

	var transitions [2]obs.Transition
	transitions[0] = obs.Transition{}
	transitions[0].Days = 30
	transitions[0].StorageClass = obs.StorageClassWarm

	transitions[1] = obs.Transition{}
	transitions[1].Days = 60
	transitions[1].StorageClass = obs.StorageClassCold
	lifecycleRule0.Transitions = transitions[:]

	lifecycleRule0.Expiration.Days = 100
	lifecycleRule0.NoncurrentVersionExpiration.NoncurrentDays = 20

	lifecycleRules[0] = lifecycleRule0

	lifecycleRule1 := obs.LifecycleRule{}
	lifecycleRule1.Status = obs.RuleStatusEnabled
	lifecycleRule1.ID = "rule1"
	lifecycleRule1.Prefix = "prefix1"
	lifecycleRule1.Expiration.Date = time.Now().Add(time.Duration(24) * time.Hour)

	var noncurrentTransitions [2]obs.NoncurrentVersionTransition
	noncurrentTransitions[0] = obs.NoncurrentVersionTransition{}
	noncurrentTransitions[0].NoncurrentDays = 30
	noncurrentTransitions[0].StorageClass = obs.StorageClassWarm

	noncurrentTransitions[1] = obs.NoncurrentVersionTransition{}
	noncurrentTransitions[1].NoncurrentDays = 60
	noncurrentTransitions[1].StorageClass = obs.StorageClassCold
	lifecycleRule1.NoncurrentVersionTransitions = noncurrentTransitions[:]
	lifecycleRules[1] = lifecycleRule1

	input.LifecycleRules = lifecycleRules[:]

	output, err := GetObsClient().SetBucketLifecycleConfiguration(input)
	if err == nil {
		fmt.Printf("StatusCode:%d, RequestId:%s\n", output.StatusCode, output.RequestId)
	} else {
		if obsError, ok := err.(obs.ObsError); ok {
			fmt.Println(obsError.StatusCode)
			fmt.Println(obsError.Code)
			fmt.Println(obsError.Message)
		} else {
			fmt.Println(err)
		}
	}
}

func GetBucketLifecycleConfiguration() {
	output, err := GetObsClient().GetBucketLifecycleConfiguration(bucketName)
	if err == nil {
		fmt.Printf("StatusCode:%d, RequestId:%s\n", output.StatusCode, output.RequestId)
		for index, lifecycleRule := range output.LifecycleRules {
			fmt.Printf("LifecycleRule[%d]:\n", index)
			fmt.Printf("ID:%s, Prefix:%s, Status:%s\n", lifecycleRule.ID, lifecycleRule.Prefix, lifecycleRule.Status)

			date := ""
			for _, transition := range lifecycleRule.Transitions {
				if !transition.Date.IsZero() {
					date = transition.Date.String()
				}
				fmt.Printf("transition.StorageClass:%s, Transition.Date:%s, Transition.Days:%d\n", transition.StorageClass, date, transition.Days)
			}

			date = ""
			if !lifecycleRule.Expiration.Date.IsZero() {
				date = lifecycleRule.Expiration.Date.String()
			}
			fmt.Printf("Expiration.Date:%s, Expiration.Days:%d\n", lifecycleRule.Expiration.Date, lifecycleRule.Expiration.Days)

			for _, noncurrentVersionTransition := range lifecycleRule.NoncurrentVersionTransitions {
				fmt.Printf("noncurrentVersionTransition.StorageClass:%s, noncurrentVersionTransition.NoncurrentDays:%d\n",
					noncurrentVersionTransition.StorageClass, noncurrentVersionTransition.NoncurrentDays)
			}
			fmt.Printf("NoncurrentVersionExpiration.NoncurrentDays:%d\n", lifecycleRule.NoncurrentVersionExpiration.NoncurrentDays)
		}
	} else {
		if obsError, ok := err.(obs.ObsError); ok {
			fmt.Println(obsError.StatusCode)
			fmt.Println(obsError.Code)
			fmt.Println(obsError.Message)
		} else {
			fmt.Println(err)
		}
	}
}

func DeleteBucketLifecycleConfiguration() {
	output, err := GetObsClient().DeleteBucketLifecycleConfiguration(bucketName)
	if err == nil {
		fmt.Printf("StatusCode:%d, RequestId:%s\n", output.StatusCode, output.RequestId)
	} else {
		if obsError, ok := err.(obs.ObsError); ok {
			fmt.Println(obsError.StatusCode)
			fmt.Println(obsError.Code)
			fmt.Println(obsError.Message)
		} else {
			fmt.Println(err)
		}
	}
}

func SetBucketTagging() {
	input := &obs.SetBucketTaggingInput{}
	input.Bucket = bucketName

	var tags [2]obs.Tag
	tags[0] = obs.Tag{Key: "key0", Value: "value0"}
	tags[1] = obs.Tag{Key: "key1", Value: "value1"}
	input.Tags = tags[:]
	output, err := GetObsClient().SetBucketTagging(input)
	if err == nil {
		fmt.Printf("StatusCode:%d, RequestId:%s\n", output.StatusCode, output.RequestId)
	} else {
		if obsError, ok := err.(obs.ObsError); ok {
			fmt.Println(obsError.StatusCode)
			fmt.Println(obsError.Code)
			fmt.Println(obsError.Message)
		} else {
			fmt.Println(err)
		}
	}
}

func GetBucketTagging() {
	output, err := GetObsClient().GetBucketTagging(bucketName)
	if err == nil {
		fmt.Printf("StatusCode:%d, RequestId:%s\n", output.StatusCode, output.RequestId)
		for index, tag := range output.Tags {
			fmt.Printf("Tag[%d]-Key:%s, Value:%s\n", index, tag.Key, tag.Value)
		}
	} else {
		if obsError, ok := err.(obs.ObsError); ok {
			fmt.Println(obsError.StatusCode)
			fmt.Println(obsError.Code)
			fmt.Println(obsError.Message)
		} else {
			fmt.Println(err)
		}
	}
}

func DeleteBucketTagging() {
	output, err := GetObsClient().DeleteBucketTagging(bucketName)
	if err == nil {
		fmt.Printf("StatusCode:%d, RequestId:%s\n", output.StatusCode, output.RequestId)
	} else {
		if obsError, ok := err.(obs.ObsError); ok {
			fmt.Println(obsError.StatusCode)
			fmt.Println(obsError.Code)
			fmt.Println(obsError.Message)
		} else {
			fmt.Println(err)
		}
	}
}

func SetBucketNotification() {
	input := &obs.SetBucketNotificationInput{}
	input.Bucket = bucketName
	var topicConfigurations [1]obs.TopicConfiguration
	topicConfigurations[0] = obs.TopicConfiguration{}
	topicConfigurations[0].ID = "001"
	topicConfigurations[0].Topic = "your topic"
	topicConfigurations[0].Events = []obs.EventType{obs.ObjectCreatedAll}

	var filterRules [2]obs.FilterRule

	filterRules[0] = obs.FilterRule{Name: "prefix", Value: "smn"}
	filterRules[1] = obs.FilterRule{Name: "suffix", Value: ".jpg"}
	topicConfigurations[0].FilterRules = filterRules[:]

	input.TopicConfigurations = topicConfigurations[:]
	output, err := GetObsClient().SetBucketNotification(input)
	if err == nil {
		fmt.Printf("StatusCode:%d, RequestId:%s\n", output.StatusCode, output.RequestId)
	} else {
		if obsError, ok := err.(obs.ObsError); ok {
			fmt.Println(obsError.StatusCode)
			fmt.Println(obsError.Code)
			fmt.Println(obsError.Message)
		} else {
			fmt.Println(err)
		}
	}
}

func GetBucketNotification() {
	output, err := GetObsClient().GetBucketNotification(bucketName)
	if err == nil {
		fmt.Printf("StatusCode:%d, RequestId:%s\n", output.StatusCode, output.RequestId)
		for index, topicConfiguration := range output.TopicConfigurations {
			fmt.Printf("TopicConfiguration[%d]\n", index)
			fmt.Printf("ID:%s, Topic:%s, Events:%v\n", topicConfiguration.ID, topicConfiguration.Topic, topicConfiguration.Events)
		}
	} else {
		if obsError, ok := err.(obs.ObsError); ok {
			fmt.Println(obsError.StatusCode)
			fmt.Println(obsError.Code)
			fmt.Println(obsError.Message)
		} else {
			fmt.Println(err)
		}
	}
}

func SetBucketEncryption() {
	input := &obs.SetBucketEncryptionInput{}
	input.Bucket = bucketName
	input.SSEAlgorithm = obs.DEFAULT_SSE_KMS_ENCRYPTION

	output, err := GetObsClient().SetBucketEncryption(input)
	if err == nil {
		fmt.Printf("StatusCode:%d, RequestId:%s\n", output.StatusCode, output.RequestId)
	} else {
		if obsError, ok := err.(obs.ObsError); ok {
			fmt.Println(obsError.StatusCode)
			fmt.Println(obsError.Code)
			fmt.Println(obsError.Message)
		} else {
			fmt.Println(err)
		}
	}
}

func GetBucketEncryption() {
	output, err := GetObsClient().GetBucketEncryption(bucketName)
	if err == nil {
		fmt.Printf("StatusCode:%d, RequestId:%s\n", output.StatusCode, output.RequestId)
		if output.KMSMasterKeyID == "" {
			fmt.Printf("KMSMasterKeyID: default master key.\n")
		} else {
			fmt.Printf("KMSMasterKeyID: %s\n", output.KMSMasterKeyID)
		}
	} else {
		if obsError, ok := err.(obs.ObsError); ok {
			fmt.Println(obsError.StatusCode)
			fmt.Println(obsError.Code)
			fmt.Println(obsError.Message)
		} else {
			fmt.Println(err)
		}
	}
}

func DeleteBucketEncryption() {
	output, err := GetObsClient().DeleteBucketEncryption(bucketName)
	if err == nil {
		fmt.Printf("StatusCode:%d, RequestId:%s\n", output.StatusCode, output.RequestId)
	} else {
		if obsError, ok := err.(obs.ObsError); ok {
			fmt.Println(obsError.StatusCode)
			fmt.Println(obsError.Code)
			fmt.Println(obsError.Message)
		} else {
			fmt.Println(err)
		}
	}
}

func ListMultipartUploads() {
	input := &obs.ListMultipartUploadsInput{}
	input.Bucket = bucketName
	input.MaxUploads = 10
	output, err := GetObsClient().ListMultipartUploads(input)
	if err == nil {
		fmt.Printf("StatusCode:%d, RequestId:%s\n", output.StatusCode, output.RequestId)
		for index, upload := range output.Uploads {
			fmt.Printf("Upload[%d]-OwnerId:%s, OwnerName:%s, UploadId:%s, Key:%s, Initiated:%s,StorageClass:%s\n",
				index, upload.Owner.ID, upload.Owner.DisplayName, upload.UploadId, upload.Key, upload.Initiated, upload.StorageClass)
		}
	} else {
		if obsError, ok := err.(obs.ObsError); ok {
			fmt.Println(obsError.StatusCode)
			fmt.Println(obsError.Code)
			fmt.Println(obsError.Message)
		} else {
			fmt.Println(err)
		}
	}
}

func DeleteObject() {
	input := &obs.DeleteObjectInput{}
	input.Bucket = bucketName
	input.Key = objectKey
	output, err := GetObsClient().DeleteObject(input)
	if err == nil {
		fmt.Printf("StatusCode:%d, RequestId:%s\n", output.StatusCode, output.RequestId)
		fmt.Printf("VersionId:%s, DeleteMarker:%v\n", output.VersionId, output.DeleteMarker)
	} else {
		if obsError, ok := err.(obs.ObsError); ok {
			fmt.Println(obsError.StatusCode)
			fmt.Println(obsError.Code)
			fmt.Println(obsError.Message)
		} else {
			fmt.Println(err)
		}
	}
}

func DeleteObjects() {
	input := &obs.DeleteObjectsInput{}
	input.Bucket = bucketName
	var objects [3]obs.ObjectToDelete
	objects[0] = obs.ObjectToDelete{Key: "key1"}
	objects[1] = obs.ObjectToDelete{Key: "key2"}
	objects[2] = obs.ObjectToDelete{Key: "key3"}

	input.Objects = objects[:]
	output, err := GetObsClient().DeleteObjects(input)
	if err == nil {
		fmt.Printf("StatusCode:%d, RequestId:%s\n", output.StatusCode, output.RequestId)
		for index, deleted := range output.Deleteds {
			fmt.Printf("Deleted[%d]-Key:%s, VersionId:%s\n", index, deleted.Key, deleted.VersionId)
		}
		for index, err := range output.Errors {
			fmt.Printf("Error[%d]-Key:%s, Code:%s\n", index, err.Key, err.Code)
		}
	} else {
		if obsError, ok := err.(obs.ObsError); ok {
			fmt.Println(obsError.StatusCode)
			fmt.Println(obsError.Code)
			fmt.Println(obsError.Message)
		} else {
			fmt.Println(err)
		}
	}
}

func SetObjectAcl() {
	input := &obs.SetObjectAclInput{}
	input.Bucket = bucketName
	input.Key = objectKey
	// input.ACL = obs.AclPublicRead
	input.Owner.ID = "ownerid"
	var grants [3]obs.Grant
	grants[0].Grantee.Type = obs.GranteeGroup
	grants[0].Grantee.URI = obs.GroupAuthenticatedUsers
	grants[0].Permission = obs.PermissionRead

	grants[1].Grantee.Type = obs.GranteeUser
	grants[1].Grantee.ID = "userid"
	grants[1].Permission = obs.PermissionWrite

	grants[2].Grantee.Type = obs.GranteeUser
	grants[2].Grantee.ID = "userid"
	grants[2].Grantee.DisplayName = "username"
	grants[2].Permission = obs.PermissionRead
	input.Grants = grants[0:3]
	output, err := GetObsClient().SetObjectAcl(input)
	if err == nil {
		fmt.Printf("StatusCode:%d, RequestId:%s\n", output.StatusCode, output.RequestId)
	} else {
		if obsError, ok := err.(obs.ObsError); ok {
			fmt.Println(obsError.StatusCode)
			fmt.Println(obsError.Code)
			fmt.Println(obsError.Message)
		} else {
			fmt.Println(err)
		}
	}
}

func GetObjectAcl() {
	input := &obs.GetObjectAclInput{}
	input.Bucket = bucketName
	input.Key = objectKey
	output, err := GetObsClient().GetObjectAcl(input)
	if err == nil {
		fmt.Printf("StatusCode:%d, RequestId:%s\n", output.StatusCode, output.RequestId)
		fmt.Printf("Owner.DisplayName:%s, Owner.ID:%s\n", output.Owner.DisplayName, output.Owner.ID)
		for index, grant := range output.Grants {
			fmt.Printf("Grant[%d]-Type:%s, ID:%s, URI:%s, Permission:%s\n", index, grant.Grantee.Type, grant.Grantee.ID, grant.Grantee.URI, grant.Permission)
		}
	} else {
		if obsError, ok := err.(obs.ObsError); ok {
			fmt.Println(obsError.StatusCode)
			fmt.Println(obsError.Code)
			fmt.Println(obsError.Message)
		} else {
			fmt.Println(err)
		}
	}
}

func RestoreObject() {
	input := &obs.RestoreObjectInput{}
	input.Bucket = bucketName
	input.Key = objectKey
	input.Days = 1
	input.Tier = obs.RestoreTierExpedited
	output, err := GetObsClient().RestoreObject(input)
	if err == nil {
		fmt.Printf("StatusCode:%d, RequestId:%s\n", output.StatusCode, output.RequestId)
	} else {
		if obsError, ok := err.(obs.ObsError); ok {
			fmt.Println(obsError.StatusCode)
			fmt.Println(obsError.Code)
			fmt.Println(obsError.Message)
		} else {
			fmt.Println(err)
		}
	}
}

func GetObjectMetadata() {
	input := &obs.GetObjectMetadataInput{}
	input.Bucket = bucketName
	input.Key = objectKey
	output, err := GetObsClient().GetObjectMetadata(input)
	if err == nil {
		fmt.Printf("StatusCode:%d, RequestId:%s\n", output.StatusCode, output.RequestId)
		fmt.Printf("StorageClass:%s, ETag:%s, ContentType:%s, ContentLength:%d, LastModified:%s\n",
			output.StorageClass, output.ETag, output.ContentType, output.ContentLength, output.LastModified)
	} else {
		if obsError, ok := err.(obs.ObsError); ok {
			fmt.Printf("StatusCode:%d\n", obsError.StatusCode)
		} else {
			fmt.Println(err)
		}
	}
}

func CopyObject() {
	input := &obs.CopyObjectInput{}
	input.Bucket = bucketName
	input.Key = objectKey
	input.CopySourceBucket = bucketName
	input.CopySourceKey = objectKey + "-back"
	input.Metadata = map[string]string{"meta": "value"}
	input.MetadataDirective = obs.ReplaceMetadata

	output, err := GetObsClient().CopyObject(input)
	if err == nil {
		fmt.Printf("StatusCode:%d, RequestId:%s\n", output.StatusCode, output.RequestId)
		fmt.Printf("ETag:%s, LastModified:%s\n",
			output.ETag, output.LastModified)
	} else {
		if obsError, ok := err.(obs.ObsError); ok {
			fmt.Println(obsError.StatusCode)
			fmt.Println(obsError.Code)
			fmt.Println(obsError.Message)
		} else {
			fmt.Println(err)
		}
	}
}

func InitiateMultipartUpload() {
	input := &obs.InitiateMultipartUploadInput{}
	input.Bucket = bucketName
	input.Key = objectKey
	input.Metadata = map[string]string{"meta": "value"}
	output, err := GetObsClient().InitiateMultipartUpload(input)
	if err == nil {
		fmt.Printf("StatusCode:%d, RequestId:%s\n", output.StatusCode, output.RequestId)
		fmt.Printf("Bucket:%s, Key:%s, UploadId:%s\n", output.Bucket, output.Key, output.UploadId)
	} else {
		if obsError, ok := err.(obs.ObsError); ok {
			fmt.Println(obsError.StatusCode)
			fmt.Println(obsError.Code)
			fmt.Println(obsError.Message)
		} else {
			fmt.Println(err)
		}
	}
}

func AbortMultipartUpload() {
	input := &obs.ListMultipartUploadsInput{}
	input.Bucket = bucketName
	output, err := GetObsClient().ListMultipartUploads(input)
	if err == nil {
		for _, upload := range output.Uploads {
			input := &obs.AbortMultipartUploadInput{Bucket: bucketName}
			input.UploadId = upload.UploadId
			input.Key = upload.Key
			output, err := GetObsClient().AbortMultipartUpload(input)
			if err == nil {
				fmt.Printf("Abort uploadId[%s] successfully\n", input.UploadId)
				fmt.Printf("StatusCode:%d, RequestId:%s\n", output.StatusCode, output.RequestId)
			}
		}
	} else {
		if obsError, ok := err.(obs.ObsError); ok {
			fmt.Println(obsError.StatusCode)
			fmt.Println(obsError.Code)
			fmt.Println(obsError.Message)
		} else {
			fmt.Println(err)
		}
	}
}

func PutObject() {
	input := &obs.PutObjectInput{}
	input.Bucket = bucketName
	input.Key = objectKey
	input.Metadata = map[string]string{"meta": "value"}
	input.Body = strings.NewReader("Hello OBS")
	output, err := GetObsClient().PutObject(input)
	if err == nil {
		fmt.Printf("StatusCode:%d, RequestId:%s\n", output.StatusCode, output.RequestId)
		fmt.Printf("ETag:%s, StorageClass:%s\n", output.ETag, output.StorageClass)
	} else {
		if obsError, ok := err.(obs.ObsError); ok {
			fmt.Println(obsError.StatusCode)
			fmt.Println(obsError.Code)
			fmt.Println(obsError.Message)
		} else {
			fmt.Println(err)
		}
	}
}

func PutFile() {
	input := &obs.PutFileInput{}
	input.Bucket = bucketName
	input.Key = objectKey
	input.SourceFile = "localfile"
	output, err := GetObsClient().PutFile(input)
	if err == nil {
		fmt.Printf("StatusCode:%d, RequestId:%s\n", output.StatusCode, output.RequestId)
		fmt.Printf("ETag:%s, StorageClass:%s\n", output.ETag, output.StorageClass)
	} else {
		if obsError, ok := err.(obs.ObsError); ok {
			fmt.Println(obsError.StatusCode)
			fmt.Println(obsError.Code)
			fmt.Println(obsError.Message)
		} else {
			fmt.Println(err)
		}
	}
}

func UploadPart() {
	sourceFile := "localfile"
	var partSize int64 = 1024 * 1024 * 5
	fileInfo, statErr := os.Stat(sourceFile)
	if statErr != nil {
		panic(statErr)
	}
	partCount := fileInfo.Size() / partSize
	if fileInfo.Size()%partSize > 0 {
		partCount++
	}
	var i int64
	for i = 0; i < partCount; i++ {
		input := &obs.UploadPartInput{}
		input.Bucket = bucketName
		input.Key = objectKey
		input.UploadId = "uploadid"
		input.PartNumber = int(i + 1)
		input.Offset = i * partSize
		if i == partCount-1 {
			input.PartSize = fileInfo.Size()
		} else {
			input.PartSize = partSize
		}
		input.SourceFile = sourceFile
		output, err := GetObsClient().UploadPart(input)
		if err == nil {
			fmt.Printf("StatusCode:%d, RequestId:%s\n", output.StatusCode, output.RequestId)
			fmt.Printf("ETag:%s\n", output.ETag)
		} else {
			if obsError, ok := err.(obs.ObsError); ok {
				fmt.Println(obsError.StatusCode)
				fmt.Println(obsError.Code)
				fmt.Println(obsError.Message)
			} else {
				fmt.Println(err)
			}
		}
	}
}

func ListParts() {
	input := &obs.ListPartsInput{}
	input.Bucket = bucketName
	input.Key = objectKey
	input.UploadId = "uploadid"
	output, err := GetObsClient().ListParts(input)
	if err == nil {
		fmt.Printf("StatusCode:%d, RequestId:%s\n", output.StatusCode, output.RequestId)
		for index, part := range output.Parts {
			fmt.Printf("Part[%d]-ETag:%s, PartNumber:%d, LastModified:%s, Size:%d\n", index, part.ETag,
				part.PartNumber, part.LastModified, part.Size)
		}
	} else {
		if obsError, ok := err.(obs.ObsError); ok {
			fmt.Println(obsError.StatusCode)
			fmt.Println(obsError.Code)
			fmt.Println(obsError.Message)
		} else {
			fmt.Println(err)
		}
	}
}

func CompleteMultipartUpload() {
	input := &obs.CompleteMultipartUploadInput{}
	input.Bucket = bucketName
	input.Key = objectKey
	input.UploadId = "uploadid"
	input.Parts = []obs.Part{
		obs.Part{PartNumber: 1, ETag: "etag1"},
		obs.Part{PartNumber: 2, ETag: "etag2"},
		obs.Part{PartNumber: 3, ETag: "etag3"},
	}
	output, err := GetObsClient().CompleteMultipartUpload(input)
	if err == nil {
		fmt.Printf("StatusCode:%d, RequestId:%s\n", output.StatusCode, output.RequestId)
		fmt.Printf("Location:%s, Bucket:%s, Key:%s, ETag:%s\n", output.Location, output.Bucket, output.Key, output.ETag)
	} else {
		if obsError, ok := err.(obs.ObsError); ok {
			fmt.Println(obsError.StatusCode)
			fmt.Println(obsError.Code)
			fmt.Println(obsError.Message)
		} else {
			fmt.Println(err)
		}
	}
}

func CopyPart() {

	sourceBucket := "source-bucket"
	sourceKey := "source-key"
	input := &obs.GetObjectMetadataInput{}
	input.Bucket = sourceBucket
	input.Key = sourceKey
	output, err := GetObsClient().GetObjectMetadata(input)
	if err == nil {
		objectSize := output.ContentLength
		var partSize int64 = 5 * 1024 * 1024
		partCount := objectSize / partSize
		if objectSize%partSize > 0 {
			partCount++
		}
		var i int64
		for i = 0; i < partCount; i++ {
			input := &obs.CopyPartInput{}
			input.Bucket = bucketName
			input.Key = objectKey
			input.UploadId = "uploadid"
			input.PartNumber = int(i + 1)
			input.CopySourceBucket = sourceBucket
			input.CopySourceKey = sourceKey
			input.CopySourceRangeStart = i * partSize
			if i == partCount-1 {
				input.CopySourceRangeEnd = objectSize - 1
			} else {
				input.CopySourceRangeEnd = (i+1)*partSize - 1
			}
			output, err := GetObsClient().CopyPart(input)
			if err == nil {
				fmt.Printf("StatusCode:%d, RequestId:%s\n", output.StatusCode, output.RequestId)
				fmt.Printf("ETag:%s, PartNumber:%d\n", output.ETag, output.PartNumber)
			} else {
				if obsError, ok := err.(obs.ObsError); ok {
					fmt.Println(obsError.StatusCode)
					fmt.Println(obsError.Code)
					fmt.Println(obsError.Message)
				} else {
					fmt.Println(err)
				}
			}
		}
	}
}

func GetObject() {
	input := &obs.GetObjectInput{}
	input.Bucket = bucketName
	input.Key = objectKey
	output, err := GetObsClient().GetObject(input)
	if err == nil {
		defer output.Body.Close()
		fmt.Printf("StatusCode:%d, RequestId:%s\n", output.StatusCode, output.RequestId)
		fmt.Printf("StorageClass:%s, ETag:%s, ContentType:%s, ContentLength:%d, LastModified:%s\n",
			output.StorageClass, output.ETag, output.ContentType, output.ContentLength, output.LastModified)
		p := make([]byte, 1024)
		var readErr error
		var readCount int
		for {
			readCount, readErr = output.Body.Read(p)
			if readCount > 0 {
				fmt.Printf("%s", p[:readCount])
			}
			if readErr != nil {
				break
			}
		}
	} else {
		if obsError, ok := err.(obs.ObsError); ok {
			fmt.Println(obsError.StatusCode)
			fmt.Println(obsError.Code)
			fmt.Println(obsError.Message)
		} else {
			fmt.Println(err)
		}
	}
}

func runExamples() {
	examples.RunBucketOperationsSample()
	//	examples.RunObjectOperationsSample()
	//	examples.RunDownloadSample()
	//	examples.RunCreateFolderSample()
	//	examples.RunDeleteObjectsSample()
	//	examples.RunListObjectsSample()
	//	examples.RunListVersionsSample()
	//	examples.RunListObjectsInFolderSample()
	//	examples.RunConcurrentCopyPartSample()
	//	examples.RunConcurrentDownloadObjectSample()
	//	examples.RunConcurrentUploadPartSample()
	//	examples.RunRestoreObjectSample()

	//	examples.RunSimpleMultipartUploadSample()
	//	examples.RunObjectMetaSample()
	//	examples.RunTemporarySignatureSample()
}
