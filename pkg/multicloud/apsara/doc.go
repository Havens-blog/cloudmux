// Copyright 2019 Yunion
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package apsara // import "github.com/Havens-blog/cloudmux/pkg/multicloud/apsara"

/*
接口相关调用
ecs相关:
ApplyAutoSnapshotPolicy
AssociateEipAddress
AttachDisk
AttachKeyPair
AuthorizeSecurityGroup
AuthorizeSecurityGroupEgress
CancelAutoSnapshotPolicy
CancelTask
ConvertNatPublicIpToEip
CreateAutoSnapshotPolicy
CreateDisk
CreateImage
CreateInstance
CreateSecurityGroup
CreateSnapshot
CreateVpc
DeleteAutoSnapshotPolicy
DeleteDisk
DeleteImage
DeleteSecurityGroup
DeleteSnapshot
DeleteVpc
DescribeAccountAttributes
DescribeAutoSnapshotPolicyEx
DescribeDisks
DescribeImages
DescribeInstanceAutoRenewAttribute
DescribeInstanceVncUrl
DescribeInstances
DescribeKeyPairs
DescribeNetworkInterfaces
DescribeRegions
DescribeRouteTables
DescribeSecurityGroupAttribute
DescribeSecurityGroups
DescribeSnapshots
DescribeTaskAttribute
DescribeTasks
DescribeVRouters
DescribeVpcs
DescribeZones
DetachDisk
DetachKeyPair
ExportImage
ImportImage
ImportKeyPair
JoinSecurityGroup
LeaveSecurityGroup
ModifyAutoSnapshotPolicyEx
ModifyEipAddressAttribute
ModifyInstanceAttribute
ModifyInstanceAutoRenewAttribute
ModifyInstanceVncPasswd
ModifySecurityGroupAttribute
ModifySecurityGroupRule
ReInitDisk
ReleaseEipAddress
RenewInstance
ReplaceSystemDisk
ResetDisk
ResizeDisk
RevokeSecurityGroup
RevokeSecurityGroupEgress
UnassociateEipAddress

vpc相关:
AllocateEipAddress
AssociateRouteTable
CreateForwardEntry
CreateSnatEntry
CreateVSwitch
DeleteForwardEntry
DeleteSnatEntry
DeleteVSwitch
DescribeEipAddresses
DescribeForwardTableEntries
DescribeIpv6GatewayAttribute
DescribeIpv6Gateways
DescribeNatGateways
DescribeRouteTableList
DescribeSnatTableEntries
DescribeVSwitchAttributes
DescribeVSwitches
UnassociateRouteTable

oss相关:
GetBucketAcl
GetBucketStorageCapacity
GetService
GetBucketInfo

elb相关:
AddBackendServers
AddVServerGroupBackendServers
CreateAccessControlList
CreateLoadBalancer
CreateLoadBalancerHTTPListener
CreateLoadBalancerHTTPSListener
CreateLoadBalancerTCPListener
CreateLoadBalancerUDPListener
CreateMasterSlaveServerGroup
CreateRules
CreateVServerGroup
DeleteAccessControlList
DeleteLoadBalancer
DeleteLoadBalancerListener
DeleteMasterSlaveServerGroup
DeleteRules
DeleteServerCertificate
DeleteVServerGroup
DescribeAccessControlListAttribute
DescribeAccessControlLists
DescribeLoadBalancerAttribute
DescribeLoadBalancerHTTPListenerAttribute
DescribeLoadBalancerHTTPSListenerAttribute
DescribeLoadBalancerTCPListenerAttribute
DescribeLoadBalancerUDPListenerAttribute
DescribeLoadBalancers
DescribeMasterSlaveServerGroupAttribute
DescribeMasterSlaveServerGroups
DescribeRuleAttribute
DescribeRules
DescribeServerCertificates
DescribeVServerGroupAttribute
DescribeVServerGroups
RemoveAccessControlListEntry
RemoveBackendServers
RemoveVServerGroupBackendServers
SetAccessControlListAttribute
SetBackendServers
SetLoadBalancerHTTPListenerAttribute
SetLoadBalancerHTTPSListenerAttribute
SetLoadBalancerStatus
SetLoadBalancerTCPListenerAttribute
SetLoadBalancerUDPListenerAttribute
SetServerCertificateName
SetVServerGroupAttribute
StartLoadBalancerListener
StopLoadBalancerListener
UploadServerCertificate

rds相关:
AllocateInstancePublicConnection
CreateAccount
CreateBackup
CreateDatabase
DeleteAccount
DeleteBackup
DeleteDBInstance
DeleteDatabase
DescribeAccounts
DescribeBackupTasks
DescribeBackups
DescribeDBInstanceAttribute
DescribeDBInstanceNetInfo
DescribeDBInstances
DescribeDatabases
DescribeParameters
GrantAccountPrivilege
ModifyDBInstanceDescription
ModifyDBInstanceSpec
RecoveryDBInstance
ReleaseInstancePublicConnection
RenewInstance
RestartDBInstance
RevokeAccountPrivilege
*/
