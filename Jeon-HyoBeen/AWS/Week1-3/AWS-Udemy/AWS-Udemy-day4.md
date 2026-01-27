Introduction to Security Groups

Security Groups are the fundamental of network security in AWS
They control how traffic is allowed into or out of our EC2 Instances.

Security groups **only contain allows rules**
Security groups rules can reference by IP or by security group

Security Grouops
Deeper Dive
- Security groups are acting as a 'firewall' on EC2 instances
- They regulate:
	- Access to Ports
	- Authorised IP ranges -IPv4 and IPv6
	- Control of inbound network (from other to the instance)
	- Control of outbound network (from the instance to other)
-  0.0.0.0/0 은 모두를 의미

Security Groups Diagram
<!-- ![[Pasted image 20260116154016.png]] -->
Security Groups
Good to know
- Can be attached to multiple instances
- Locked down to a region / VPC combination
- Does live "outside" the EC2 - if traffic is blocked the EC2 instance won't see it
- It's good to maintain one separate security group for SSH access
- If your application is not accessible (time out), then it's a security group issue
- If your application gives a "connection refused" error, then it's an application error or it's not launched
- All inbound traffic is blocked by default
- All outbound traffic is authorised by default

Referencing other security groups Diagram
<!-- ![[Pasted image 20260116155001.png]] -->

Classic Ports to know
22 = SSH (Secure Shell) - log into a Linux instance
21 = FTp (File Transfer Protocol) - upload files into a file share
22 = SFTP (Secure File Transfer Protocol) - upload files using SSH
80 = HTTP - access unsecured website
443 = HTTPS - access secured websites
3389 = RDP (Remote Desktop Protocol) - log into a Windows instance


EC2 Network & Security
에 보안그룹(Security Groups)
에 들어가서 볼 수 있고, 각각 ID (식별자)가 존재한다.

인바운드 - 바깥에서 인스턴스로 접근할때의 규칙 
타임아웃이 보일때마다 100 % EC2 보안그룹의 원인이다.

아웃바운드 

보안 그룹이 필요할때마다 추가해서 할 수 있다. 하나의 보안 그룹에 보안 규칙을 많이 설정할수도있고

SSH Summary Table
SSH는 명령줄 인터페이스 유틸리티로 Mac Linux Windows>= 10 에 사용가능,
Windows <10 -> Putty (10이상도 사용가능)
EC2 Instance Connect는 전부 가능ㅇ

SSH
SSH is one of the most important function. It allows you to control a remote machine, all using the command line.
인스턴스 껏다 킬때마다 공용 IP 바뀔 수 있는 거 확인하고
chmod 400 "* .pem" 
ssh -i "* .pem" ec2-user@공용 IP
ssh ec2-user@공용 IP 했을때 connecting 에 대해서 yes 해라 - 첫 연결시 진짜 하냐고 묻는거임

Use the default option "Connect using EC2 Instance Connect"
임시 SSH 키를 줘서 여는거임 브라우저에서 CLI 를 다루네,
보안그룹에서 SSH를 열어야함,. IPv6를 사용해야함..
r그냥 CLI네

인스턴스에 연결해서 Access key로 입력하지마라 절대!! 다 볼 수있으니까!!!!!!
IAM Role를 이용해라

configure 해서 받으려고 하지말고, IAM Role을 부여해서 권한관리해라!!!!!
자격증명 액세스 키로 하지마!!!!!!절대 !!!!!!!!!!!!!!!!!!! IAM Role 부여한다고ㅓ 바론 안되지만 금방 된다 ^ㅡ^

EC2 Instance Purchasing Option
- On-Demand Instances - short workload, predictable pricing, pay by second
- Reserved (1 & 3 years)
	- Reserved Instances - long workload
	- Convertible Reserved Instances - long workloads with flexible instances
- Saving Plans (1 & 3 years ) -commitment to an amount of usage, long workload
- Spot Instances - short workloads, cheap, can lose instances(less reliable)
- Dedicated Host - book an entire physical server, control instance placement
- Dedicated Instances - no other customers will share your hardware
- Capacity Reservations - reserve capacity in a specific AZ for any duration

EC2 On Demand
- Pay for what you use:
	- Linux or Windows - billing per second, after the first minute
	- All other operating systems - billing per hour
- Has the highest cost but no upfront payment
- No long-term commitment

- Recommended for short-term and un-interrupted workloads, where you can't predict how the application will behave

EC2 Reserved Instances - 할인율은 달라질수 있다.
- UP tp 72%(변동성 존재) discount compared to On-demand
- You reserve a specific instance attributes (Instance Type, Region, Tenancy, OS)
- Reservation Period - 1 (+ discount) year or 3 year (+++ discount)
- Payment Options - No Upfront(+), Partial Upfront(++), All Upfront(+++)
- Reserved Instance's Scope - Regional or Zonal (reserve capacity in an AZ)
- Recommended for steady-state usage applications (think database)
- You can buy and sell in the Reserved Instance Marketplace

- Convertible Reserved Instance
	- Can change the EC2 instance type, instance family, OS, scope and tenancy
	- Up to 66% discount

EC2 Saving Plans
- Get a discount based on long-term usage (up to 72% - same as Rls)
- Commit to a certain type of usage ($10/hour for 1 or 3 years)
- Usage beyond EC2 Savings Plans is billed at the On-Demand price

- Locked to a specific instance family & AWS region (e.g., M5 in us-east-1)
- Flexible across:
	- Instance Size(e.g., m5.xlarge, m5.2xlarge)
	- OS (e.g., Linux , Windows)
	- Tenancy(Host,Dedicated,Default)

EC2 Spot Instances
- Can get a discount of up to 90% 
- Instances that you can "lose" at any point of time if your max price is less than the current spot price
- The MOST cost-efficient instances in AWS

- Useful for workloads that are resilient to failure
	- Batch jobs
	- Data analysis
	- Image processing
	- Any distributed workloads
	- Workloads with a flexible start and end time
- Not suitable for critical jobs or databases

EC2 Dedicated Hosts
- A physical server with EC2 instance capacity fully dedicated to your use
- Allows you address compliance requirements and use your existing server-bound software licenses (per-socket, per-core, pe--VM software licenses)
- Purchasing Options:
	-  On-demand - pay per second for active Dedicated Host
	- Reserved - 1 or 3 years (No Upfront, Partial Upfront, All Upfront)
- The most expensive option

- Useful for software that have complicated licensing model (BYOL - Bring Your Own License)
- Or for companies that have strong regulatory or compliance needs

EC2 Dedicated Instances
- Instances run on hardware that's dedicated to you
- May share hardware with other instances in same account
- No control over instance placement (can move hardware after Stop/Start)
  <!-- ![[Pasted image 20260116171249.png]]여기서 x표 된게 특징임 (구분 잘해야겠지~~ ??) -->

EC2 Capacity Reservations
- Reserve On-Demand instances capacity in a specific 
- You always have access to EC2 capacity when you need it
- No time commitment (create/cancel anytime), no billing discount
- Combine with Regional Reserved Instances and Savings Plans to benefit from billing discount
- You're charged at On-Demand rate whether you run instances or not

- Suitable for short-term, uninterrupted workloads that needs to be in a specific AZ

Which purchasing option is right for me?
- On demand: coming and staying in resort whenever we like, we pay the full price
- Reserved: like planning ahead and if we plan to stay for a long time, we may get a good discount.
- Savings Plans: pay a certain amount per hour for certain period and stay in any room type(e.g., King, Suite, Sea View...)
- Spot instances: the hotel allows people to bid for the empty rooms and the highest bidder keeps the rooms, You can get kicked out at any time
- Dedicated Hosts: We book an entire building of the resort
- Capacity Reservations: you book a room for a period with full price even you don't stay in it

Example -m4.large - us-east-1  ![[Pasted image 20260116172357.png]]

