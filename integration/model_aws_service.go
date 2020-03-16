/*
 * Integrations API
 *
 * APIs for creating, retrieving, updating, and deleting SignalFx integrations to the systems you use.<br> An integration provides SignalFx with information from the external system that you're connecting to. You'll need to retrieve this information from the external system before you use the API. Each external system is different, so to see a summary of its requirements and procedures, view its request body description. # Authentication To create, update, delete, or validate an integration, you need to authenticate your request using a session token associated with a SignalFx administrator. To **retrieve** an integration, your session token doesn't need to be associated with an administrator. You can also retrieve integrations using an org token.<br> In the web UI, session tokens are known as <strong>user access</strong> tokens, and org tokens are known as <strong>access tokens</strong>. <br> To learn more about authentication tokens, see the topic [Authentication Tokens](https://developers.signalfx.com/administration/access_tokens_overview.html) in the Developers Guide. # Supported service types SignalFx offers integrations for the following:<br>   * Data collection from other monitoring systems such as AWS CloudWatch   * Authentication using your existing Single Sign-On (**SSO**) system   * Sending alerts using your preferred messaging, chat, or incident management service <br> To use one of these integrations, you first register it with SignalFx. After that, you configure the integration to communicate between the system you're using and SignalFx. ## Data collection SignalFx integrations APIs support data collection for the following services:<br>   * Amazon Web Services (**AWS**)   * Google Cloud Platform (**GCP**)   * Microsoft Azure   * NewRelic  ## Authentication using SSO SignalFx integration APIs support SAML-based SSO integrations for the following services:<br>   * Microsoft Active Directory Federation Services (**ADFS**)   * Bitium   * Okta   * OneLogin   * PingOne  ## Alerts using message, chat, or incident management services SignalFx integration APIs support alert notifications using the following services: <br>   * BigPanda   * Office 365   * Opsgenie   * PagerDuty   * ServiceNow   * Slack   * VictorOps   * Webhook   * xMatters<br>  **NOTE:** You can't create Office 365 integrations using the API, and your ability to update them in a **PUT** request is limited, but you can retrieve their data or delete them. To create an Office 365 integration, use the the web UI. <br> # Viewing request body documentation The *request* body format for the following operations depends on the type of integration you use:<br>   * POST `/integration`   * PUT `/integration/{id}`<br>  The *response* body format for the following operations also depends on the type of integration you use:<br>   * GET `/integration`   * GET `/integration/{id}`  <br>  To see the request or response body format for an integration: <br>   1. Find the endpoint and method.   2. For a request body, find the section *REQUEST BODY SCHEMA*. For a     response body, find the section *RESPONSE SCHEMA*.   3. Scroll down to the `type` property.   4. At the end of the description for `type`, find the dropdown box that      contains the integration type. By default, it's set to *AWSCloudWatch*.   5. To see a complete list of integrations, click the down arrow. A list      with a vertical scroll bar appears.   6. Select the integration type from the list. The request body properties      for this integration type now appear.
 *
 * API version: 3.3.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package integration

// AwsService : An AWS service that you want SignalFx to collect data from. SignalFx supports the following AWS services:<br>   * AWS/ApiGateway   * AWS/AppStream   * AWS/AutoScaling   * AWS/Billing   * AWS/CloudFront   * AWS/CloudSearch   * AWS/Events   * AWS/Logs   * AWS/Connect   * AWS/DMS   * AWS/DX   * AWS/DynamoDB   * AWS/EC2   * AWS/EC2Spot   * AWS/ECS   * AWS/ElasticBeanstalk   * AWS/EBS   * AWS/EFS   * AWS/ELB   * AWS/ApplicationELB   * AWS/NetworkELB   * AWS/ElasticTranscoder   * AWS/ElastiCache   * AWS/ES   * AWS/ElasticMapReduce   * AWS/GameLift   * AWS/Inspector   * AWS/IoT   * AWS/KMS   * AWS/KinesisAnalytics   * AWS/Firehose   * AWS/Kinesis   * AWS/KinesisVideo   * AWS/Lambda   * AWS/Lex   * AWS/ML   * AWS/OpsWorks   * AWS/Polly   * AWS/Redshift   * AWS/RDS   * AWS/Route53   * AWS/SageMaker   * AWS/DDoSProtection   * AWS/SES   * AWS/SNS   * AWS/SQS   * AWS/S3   * AWS/SWF   * AWS/States   * AWS/StorageGateway   * AWS/Translate   * AWS/NATGateway   * AWS/VPN   * WAF   * AWS/WorkSpaces
type AwsService string

// List of AWSService
const (
	AWSAPI_GATEWAY               AwsService = "AWS/ApiGateway"
	AWSAPP_STREAM                AwsService = "AWS/AppStream"
	AWSATHENA                    AwsService = "AWS/Athena"
	AWSAUTO_SCALING              AwsService = "AWS/AutoScaling"
	AWSBILLING                   AwsService = "AWS/Billing"
	AWSACP_PRIVATE_CA            AwsService = "AWS/ACMPrivateCA"
	AWSCLOUD_FRONT               AwsService = "AWS/CloudFront"
	AWSCLOUD_HSM                 AwsService = "AWS/CloudHSM"
	AWSCLOUD_SEARCH              AwsService = "AWS/CloudSearch"
	AWSEVENTS                    AwsService = "AWS/Events"
	AWSLOGS                      AwsService = "AWS/Logs"
	AWSCODEBUILD                 AwsService = "AWS/CodeBuild"
	AWSCOGNITO                   AwsService = "AWS/Cognito"
	AWSCONNECT                   AwsService = "AWS/Connect"
	AWSDMS                       AwsService = "AWS/DMS"
	AWSDX                        AwsService = "AWS/DX"
	AWSDOCDB                     AwsService = "AWS/DocDB"
	AWSDYNAMO_DB                 AwsService = "AWS/DynamoDB"
	AWSEC2                       AwsService = "AWS/EC2"
	AWSEC2_SPOT                  AwsService = "AWS/EC2Spot"
	AWSECS                       AwsService = "AWS/ECS"
	AWSELASTIC_BEANSTALK         AwsService = "AWS/ElasticBeanstalk"
	AWSELASTIC_INTERFACE         AwsService = "AWS/ElasticInterface"
	AWSEBS                       AwsService = "AWS/EBS"
	AWSEFS                       AwsService = "AWS/EFS"
	AWSELB                       AwsService = "AWS/ELB"
	AWSAPPLICATION_ELB           AwsService = "AWS/ApplicationELB"
	AWSNETWORK_ELB               AwsService = "AWS/NetworkELB"
	AWSELASTIC_TRANSCODER        AwsService = "AWS/ElasticTranscoder"
	AWSELASTI_CACHE              AwsService = "AWS/ElastiCache"
	AWSES                        AwsService = "AWS/ES"
	AWSELASTIC_MAP_REDUCE        AwsService = "AWS/ElasticMapReduce"
	AWSFSX                       AwsService = "AWS/FSx"
	AWSGAME_LIFT                 AwsService = "AWS/GameLift"
	AWSGLUE                      AwsService = "AWS/Glue"
	AWSINSPECTOR                 AwsService = "AWS/Inspector"
	AWSIO_T                      AwsService = "AWS/IoT"
	AWSIO_T_ANALYTICS            AwsService = "AWS/IoTAnalytics"
	AWSKAFKA                     AwsService = "AWS/Kafka"
	AWSKMS                       AwsService = "AWS/KMS"
	AWSKINESIS_ANALYTICS         AwsService = "AWS/KinesisAnalytics"
	AWSFIREHOSE                  AwsService = "AWS/Firehose"
	AWSKINESIS                   AwsService = "AWS/Kinesis"
	AWSKINESIS_VIDEO             AwsService = "AWS/KinesisVideo"
	AWSLAMBDA                    AwsService = "AWS/Lambda"
	AWSLEX                       AwsService = "AWS/Lex"
	AWSMEDIA_CONNECT             AwsService = "AWS/MediaConnect"
	AWSMEDIA_CONVERT             AwsService = "AWS/MediaConvert"
	AWSMEDIA_PACKAGE             AwsService = "AWS/MediaPackage"
	AWSMEDIA_TAILOR              AwsService = "AWS/MediaTailor"
	AWSML                        AwsService = "AWS/ML"
	AWSAMAZON_MQ                 AwsService = "AWS/AmazonMQ"
	AWSOPS_WORKS                 AwsService = "AWS/OpsWorks"
	AWSPOLLY                     AwsService = "AWS/Polly"
	AWSREDSHIFT                  AwsService = "AWS/Redshift"
	AWSRDS                       AwsService = "AWS/RDS"
	AWSROBOMAKER                 AwsService = "AWS/RoboMaker"
	AWSROUTE53                   AwsService = "AWS/Route53"
	AWSSAGE_MAKER                AwsService = "AWS/SageMaker"
	AWSSAGE_MAKER_TRAINING_JOBS  AwsService = "aws/sagemaker/TrainingJobs"
	AWSSAGE_MAKER_TRANSFORM_JOBS AwsService = "aws/sagemaker/TransformJobs"
	AWSSAGE_MAKER_ENDPOINTS      AwsService = "aws/sagemaker/Endpoints"
	AWSSDK_METRICS               AwsService = "AWS/SDKMetrics"
	AWSD_DO_S_PROTECTION         AwsService = "AWS/DDoSProtection"
	AWSSES                       AwsService = "AWS/SES"
	AWSSNS                       AwsService = "AWS/SNS"
	AWSSQS                       AwsService = "AWS/SQS"
	AWSS3                        AwsService = "AWS/S3"
	AWSSWF                       AwsService = "AWS/SWF"
	AWSSTATES                    AwsService = "AWS/States"
	AWSSTORAGE_GATEWAY           AwsService = "AWS/StorageGateway"
	AWSTEXTRACT                  AwsService = "AWS/Textract"
	AWSTHINGS_GRAPH              AwsService = "AWS/ThingsGraph"
	AWSTRANSLATE                 AwsService = "AWS/Translate"
	AWSTRUSTED_ADVISOR           AwsService = "AWS/TrustedAdvisor"
	AWSNAT_GATEWAY               AwsService = "AWS/NATGateway"
	AWSVPN                       AwsService = "AWS/VPN"
	WAF                          AwsService = "WAF"
	AWSWORK_MAIL                 AwsService = "AWS/WorkMail"
	AWSWORK_SPACES               AwsService = "AWS/WorkSpaces"
	AWSNEPTUNE                   AwsService = "AWS/Neptune"
	AWSMEDIA_LIVE                AwsService = "AWS/MediaLive"
	AWS_SYSTEM_LINUX             AwsService = "System/Linux"
)

var AWSServiceNames = map[string]AwsService{
	"AWS/ApiGateway":              AWSAPI_GATEWAY,
	"AWS/AppStream":               AWSAPP_STREAM,
	"AWS/Athena":                  AWSATHENA,
	"AWS/AutoScaling":             AWSAUTO_SCALING,
	"AWS/Billing":                 AWSBILLING,
	"AWS/ACMPrivateCA":            AWSACP_PRIVATE_CA,
	"AWS/CloudFront":              AWSCLOUD_FRONT,
	"AWS/CloudHSM":                AWSCLOUD_HSM,
	"AWS/CloudSearch":             AWSCLOUD_SEARCH,
	"AWS/Events":                  AWSEVENTS,
	"AWS/Logs":                    AWSLOGS,
	"AWS/CodeBuild":               AWSCODEBUILD,
	"AWS/Cognito":                 AWSCOGNITO,
	"AWS/Connect":                 AWSCONNECT,
	"AWS/DMS":                     AWSDMS,
	"AWS/DX":                      AWSDX,
	"AWS/DocDB":                   AWSDOCDB,
	"AWS/DynamoDB":                AWSDYNAMO_DB,
	"AWS/EC2":                     AWSEC2,
	"AWS/EC2Spot":                 AWSEC2_SPOT,
	"AWS/ECS":                     AWSECS,
	"AWS/ElasticBeanstalk":        AWSELASTIC_BEANSTALK,
	"AWS/ElasticInterface":        AWSELASTIC_INTERFACE,
	"AWS/EBS":                     AWSEBS,
	"AWS/EFS":                     AWSEFS,
	"AWS/ELB":                     AWSELB,
	"AWS/ApplicationELB":          AWSAPPLICATION_ELB,
	"AWS/NetworkELB":              AWSNETWORK_ELB,
	"AWS/ElasticTranscoder":       AWSELASTIC_TRANSCODER,
	"AWS/ElastiCache":             AWSELASTI_CACHE,
	"AWS/ES":                      AWSES,
	"AWS/ElasticMapReduce":        AWSELASTIC_MAP_REDUCE,
	"AWS/FSx":                     AWSFSX,
	"AWS/GameLift":                AWSGAME_LIFT,
	"AWS/Glue":                    AWSGLUE,
	"AWS/Inspector":               AWSINSPECTOR,
	"AWS/IoT":                     AWSIO_T,
	"AWS/IoTAnalytics":            AWSIO_T_ANALYTICS,
	"AWS/Kafka":                   AWSKAFKA,
	"AWS/KMS":                     AWSKMS,
	"AWS/KinesisAnalytics":        AWSKINESIS_ANALYTICS,
	"AWS/Firehose":                AWSFIREHOSE,
	"AWS/Kinesis":                 AWSKINESIS,
	"AWS/KinesisVideo":            AWSKINESIS_VIDEO,
	"AWS/Lambda":                  AWSLAMBDA,
	"AWS/Lex":                     AWSLEX,
	"AWS/MediaConnect":            AWSMEDIA_CONNECT,
	"AWS/MediaConvert":            AWSMEDIA_CONVERT,
	"AWS/MediaPackage":            AWSMEDIA_PACKAGE,
	"AWS/MediaTailor":             AWSMEDIA_TAILOR,
	"AWS/ML":                      AWSML,
	"AWS/AmazonMQ":                AWSAMAZON_MQ,
	"AWS/OpsWorks":                AWSOPS_WORKS,
	"AWS/Polly":                   AWSPOLLY,
	"AWS/Redshift":                AWSREDSHIFT,
	"AWS/RDS":                     AWSRDS,
	"AWS/RoboMaker":               AWSROBOMAKER,
	"AWS/Route53":                 AWSROUTE53,
	"AWS/SageMaker":               AWSSAGE_MAKER,
	"aws/sagemaker/TrainingJobs":  AWSSAGE_MAKER_TRAINING_JOBS,
	"aws/sagemaker/TransformJobs": AWSSAGE_MAKER_TRANSFORM_JOBS,
	"aws/sagemaker/Endpoints":     AWSSAGE_MAKER_ENDPOINTS,
	"AWS/SDKMetrics":              AWSSDK_METRICS,
	"AWS/DDoSProtection":          AWSD_DO_S_PROTECTION,
	"AWS/SES":                     AWSSES,
	"AWS/SNS":                     AWSSNS,
	"AWS/SQS":                     AWSSQS,
	"AWS/S3":                      AWSS3,
	"AWS/SWF":                     AWSSWF,
	"AWS/States":                  AWSSTATES,
	"AWS/StorageGateway":          AWSSTORAGE_GATEWAY,
	"AWS/Textract":                AWSTEXTRACT,
	"AWS/ThingsGraph":             AWSTHINGS_GRAPH,
	"AWS/Translate":               AWSTRANSLATE,
	"AWS/TrustedAdvisor":          AWSTRUSTED_ADVISOR,
	"AWS/NATGateway":              AWSNAT_GATEWAY,
	"AWS/VPN":                     AWSVPN,
	"WAF":                         WAF,
	"AWS/WorkMail":                AWSWORK_MAIL,
	"AWS/WorkSpaces":              AWSWORK_SPACES,
	"AWS/Neptune":                 AWSNEPTUNE,
	"AWS/MediaLive":               AWSMEDIA_LIVE,
	"System/Linux":                AWS_SYSTEM_LINUX,
}
