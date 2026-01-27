How to choose an AWS Region

**Compliance** with data governance and legal requirements
**Proximity** to customer
**Available service** within a Region
**Pricing**

AWS Availability Zones

가용지점은 나눠져있다.. 재난 상황에 대비하여
초지연적 네트워킹과 머시기로 연결되어있음.. 
하나의 가용 지점 (ex southern asia 2, 안에 3개 지점 존재, 지점 별 데이터센터 여러개 존재(지점 별로 다름), 커넥션이 있음)

---

IAM : Users & Groups - **Global Service**
IAM = Identity and Access Management, Global Service

Root account created by default, shouldn't be used or shared
Users are people within your organization, and can be grouped
Groups only contain users, not other groups
Users don't have to belong to a group, and user can belong to multiple groups

IAM: Permission
Users or Groups can be assigned JSON documents called policies
These policies define the permissions of the users
In AWS you apply the least privilege principle: don't give more permissions than a user needs

AWS는 권한을 막 부여하지 않는다. 최소한으로 부여한다.
IAM -> 그룹별 상속 가능
inline -> 개별 사용자에 대한..

IAM Policies Structure
Consists of
- Version: policy language version, always include "2012-10-17"
- ID: an identifier for the policy (optional)
- Statement: one or more individual statements(required)
Statements consists of
- Sid: an identifier for the statement (optional)
- Effect: whether the statement allows or denies access(Allow, Deny)
- Principal: account/user/role to which this policy applied to
- Actions: list of actions this policy allows or denies
- Resource: list of resources to which the actions applied to
- Condition: conditions for when this policy is in effect(optional)


MFA
- Users have access to your account and can possibly change configurations or delete resources in your AWS account
- You want to protect your Root Accounts and IAM users
- MFA = password you know + security device you own

Main benefit of MFA :
	if a password is stolen or hacked, the account is not compromised

MFA devices options in AWS
	Virtual MFA device
		- Google Authenticator(phone only)
		- Authy (phone only)
		- U2F (physical device)
		- Hardware Key Fob MFA Device
		- AWS GovCloud(US)

How can users access AWS?
To access AWS, you have three options:
- AWS Management Console(protected by password + MFA)
- AWS Command Line Interface(CLI)
- AWS Software Developer Kit(SDK)
Access Keys are generated through the AWS Console
User manage their own access key
Access Keys are secret, just like a password. Don't share them
Access Key ID ~= username
Secret Access Key ~= password

What's the AWS CLI -> 터미널 커맨드를 의미하는가? -> ㅇㅇ 아마 터미널 이용ㅇ해서 AWS를 다룰 수 있다는 뜻인듯, Console 대신해서 사용할 수 있다니까.
Command line interface
 A tool that enables you to interact with AWS services using commands in your command-line shell
Direct access to the public APIs of AWS services
You can develop scripts to manage your resources
It's open-source https://github.com/aws/aws-cli
Alternative to using AWS Management Console

What's the AWS SDK
AWS Software Development Kit
Language-specific APIs (set of libraries)
Enables you to access and manage AWS services programmatically
Embedded within your application
Supports
- SDKs (programming language)
- Mobile SDKs
- IoT Device SDKs 
Example:AWS CLI is built on AWS SDK for Python