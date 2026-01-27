EBS Volume?
- An EBS(Elastic Block Store) Volume is a network drive you can attach to your instances while they run
- It allows your instances to persist data, even after their termination
- They can only be mounted to one instance at a time (at the CCP level)
- They are bound to a specific availability zone
- Analogy: Think of them as a "network USB stick"
-  -> EBS 는 1:1 연결이지만, io1/io2 타입은 Multi-Attach 로 1:N 연결이 가능하다.

- It's a network drive (i.e. not a physical drive)
	- It uses the network to communicate the instance, which means there might be a bit of latency
	- It can be detached from an EC2 instance and attached to another one quickly
- It's locked to an Availability Zone(AZ)
	- An EBS Volume in us-east-1 a cannot be attached to us-east-1b
	- To move a volume across, you first need to snapshot it
- Have a provisioned capacity (size in GBs, and IOPS)
	- You get billed for all the provisioned capacity
예시,
하나의 EC2에 두개의 EBS를 연결할 수 있다.

EBS - Delete on Termination attribute
EBS를 만들때 Delete on Termination 체크박스가 있다.  -> 인스턴스를 만드는데 고급에서 루트 볼륨에 대해서 삭제할지 말지 결정 가능
-> Controls the EBS behavior when an EC2 instance terminates
	- By default, the root EBS volume is deleted (attribute enabled)
	- By default, any other attached EBS volume is not deleted(attribute disable)
- This can be controlled by the AWS console / AWS CLI
- Use case: preserve root volume when instance is terminated

EBS Snapshots
- Make a backup (snapshot) of your EBS volume at a point in time
- Not necessary to detach volume to do snapshot, but recommended
- Can copy snapshots across AZ or region

EBS Snapshot Features
- EBS Snapshot Archive
	- Move a Snapshot to an "archive tier" that is 75% cheaper
	- Takes within 24 to 72 hours for restoring the archive
- Recycle Bin for EBS Snapshots
	- Setup rules to retain deleted snapshots so you can recover them after an accidental deletion
	- Specify retention (from 1 day to 1 year)
- Fast Snapshot Restore 
	- Force full initialization of snapshot to have no latency on the first use($ $ $)
- 스냅샷 복사로 AZ 또는 지역을 옮길수도 스냅샷을 이용하여 볼륨을 다른 AZ로 만들 수도 있다. 
- 이후 스냅샷은 지워질텐데, 이걸 Recycle Bin에서 복구 가능
- 스냅샷은 중분 방식이다 -> 첫번째는 전체 저장, 두번째부터는 변경된 블록만 저장.


AMI = Amazon Machine Image
AMI는 라이브러리를 만드는거랑 같은건가? - ㅇㅇ..비슷함
AMI are a customization of an EC2 instance
- You add your own software, configuration, operating system, monitoring
- Faster boot / configuration time because all your software is pre-packaged
AMI are built for a specific region (and can be copied across regions)
You can launch EC2 instances from:
- A public AMI: AWS provided
- Your own AMI: you make and maintain them yourself
- An AWS Marketplace AMI: an AMI someone else made(and potentially sells)

AMI Process ( from an EC2 instance)
- Start an EC2 instance and customize it
- Stop the instance (for data integrity)
- Build an AMI - this will also create EBS snapshots
- Launch instances from other AMIs
e.g.) EC2 in us-east-1a -> create AMI -> Custom AMI -> Launch from AMI -> Other EC2 in us-east-1b

인스턴스에서 우클릭해서 Image and templates -> create Image 
이후 왼쪽에서 Images 에 AMI로 ㄱㄱ 이후 클릭해서 Launch instance from AMI 혹은 인스턴스 생성에서도 가능하다. (OS 선택하는곳에서 My AMIs 로)
훨씬 부팅이 빠르다잇

EC2 Instance Store
- EBS volumes are network drives with good but "limited" performance
- If you need a high-performance hardware disk, use EC2 Instance Store

- Better I/O performance  (쓰로풋도 좋음)
- EC2 Instance Store lose their storage if they're stopped (ephemeral)
- Good for buffer/ cache/ scratch data/ temporary content
- Risk of data loss if hardware fails
- Backups and Replication are your responsibility
-> 장기 저장 목표 XX, 단기적으로 퍼포먼스 좋을 필요가 있으면 ㄱㄱ, 단 데이터 손실을 주의해야함.
i로 쓰인 놈들인듯? (i3.large, i4i.xlarge 등) -> 인스턴스 타입
물리적 하드웨어를 사용하긴하나 root Volume이 생성 안되는건 아님

EBS Volume Types
- come in 6types
	- gp2/gp3(SSD) : General purpose SSD volume that balances price and performance for a wide variety of workloads
	- io 1/io 2 Block Express (SSD) : Highest-performance SSD volume for mission-critical low-latency or high-throughput workloads
	- st 1 (HDD): Low cost HDD volume designed for frequently accessed, throughput-intensive workloads
	- sc 1 (HDD): Lowest cost HDD volume designed for less frequently accessed workloads
- EBS  Volumes are characterized in Size | Throughput | IOPS(I/O Ops per Sec)
- When in doubt always consult the AWS documentation - it's good!
- Only gp2/gp3 and io 1/ io2 Block Express can be used as boot volumes

General purpose SSD
- Cost effective storage, low-latency
- System boot volumes, Virtual desktops, Development and test environments
- 1Gib - 16TiB
- gp3:
	- Baseline of 3,000 IOPS and throughput of 125 MiB/s
	- Can increase IOPS up to 16,000 and throughput up to 1000 MiB/s independently
- gp2:
	- Small gp2 volumes can burst IOPS to 3,000
	- Size of the volume and IOPS are linked, max IOPS is 16,000
	- 3 IOPS per GiB, means at 5,334 GiB we are at the max IOPS
- gp3와 gp2의 차이점을 알아보자
	- 연결되있고 안되있고 차인데,,, , 내가 원하는 만큼의 IOPS와 처리량을 독립적으로 설정ㅎ할 수 있으니 gp3가 더 권장됨

Provisioned IOPS (PIOPS) SSD
- Critical business applications with sustained IOPS performance
- Or applications that need more than 16,000 IOPS
- Great for databases workloads (sensitive to storage perf and consistency)
- io 1(4 GiB - 16 TiB):
	- Max PIOPS: 64,000 for Nitro EC2 instances & 32,000 for other
	- Can increase PIOPS independently from storage size
- io 2 Block Express (4 GiB - 64 TiB):
	- Sub-millisecond latency
	- Max PIOPS:256,000 with an IOPS:GiB ratio of 1,000: 1
- Supports EBS Multi-attach

EBS Volume Types Use cases
Hard Disk Drives (HDD)
- Cannot be a boot volume
- 125 GiB to 16 TiB
- Throughput Optimized HDD (st 1)
	- Big Data, Data Warehouses, Log Processing
	- Max throughput 500 MiB/s - max IOPS 500
- Cold HDD (sc 1):
	- For data that is infrequently accessed
	- Scenarios where lowest cost is important
	- Max throughput 250MiB/s - max IOPS 250

세부사항을 기억할 필요는 없지만, 높은 IOPS 를 얻으려면 EC2 Nitro와 io1 혹은 io2가 필요하다를 기억하면 된다.

EBS Multi-Attach - io1/ io2 family
- Attach the same EBS volume to multiple EC2 instances in the same AZ
- Each instance has full read & write permissions to the high - performance volume
- Use case:
	- Achieve higher application availability in clustered Linux applications (ex: Teradata)
	- Applications must manage concurrent write operations
- Up to **16** EC2 Instances at a time -> 16 숫자 기억하자
- Must use a file system that's cluster-aware (not XFS, EXT4, etc...) 클러스터화 된 파일 시스템만 된다~~ 일반 파일 시스템 안됨

EBS Encryption
- When you create an encrypted EBS volume, you get the following:
	- Data at rest is encrypted inside the volume
	- All the data in flight moving between the instance and the volume is encrypted
	- All snapshots are encrypted
	- All volumes created from the snapshot
- Encryption and decryption are handled transparently ( you have nothing to do)
- Encryption has a minimal impact on latency
- EBS Encryption leverages keys from KMS (AES-256) //AES-256 기억하랬음..
- Copying an unencrypted snapshot allows encryption
- Snapshots of encrypted volumes are encrypted

Encryption: encrypt an unencrypted EBS volume
EBS 스냅샷만들고, 암호화 걸고, 스냅샷을 통해서 다시 EBS 만들고 (자동 암호화 될거임)
이후에 오리지널 인스턴스에 다시 붙이면 된다.

인스턴스 중지하고 루트 볼륨을 디테치하고 새로 만든애를 루트로 다시 어태치하면 된다.