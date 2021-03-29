package main

import (
    "flag"
    "fmt"
    "github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/aws/credentials"
    "github.com/aws/aws-sdk-go/service/sqs"
)

const (
    QueueUrl    = "https://sqs.us-east-1.amazonaws.com/486892814007/TaskQueue-test"
    Region      = "us-east-1"
    CredPath    = "/Users/jefersonagudelo/.aws/credentials"
    CredProfile = "jefersonagudeloc"
)

func GetQueueURL(sess *session.Session, queue *string) (*sqs.GetQueueUrlOutput, error) {

    svc := sqs.New(sess)

    urlResult, err := svc.GetQueueUrl(&sqs.GetQueueUrlInput{
        QueueName: queue,
    })

    if err != nil {
        return nil, err
    }

    return urlResult, nil
}


func GetMessages(sess *session.Session, queueURL *string, timeout *int64) (*sqs.ReceiveMessageOutput, error) {
    svc := sqs.New(sess)

    msgResult, err := svc.ReceiveMessage(&sqs.ReceiveMessageInput{
        AttributeNames: []*string{
            aws.String(sqs.MessageSystemAttributeNameSentTimestamp),
        },
        MessageAttributeNames: []*string{
            aws.String(sqs.QueueAttributeNameAll),
        },
        QueueUrl:            queueURL,
        MaxNumberOfMessages: aws.Int64(1),
        VisibilityTimeout:   timeout,
	})
	
    if err != nil {
        return nil, err
    }

    return msgResult, nil
}

func main() {

    queue := flag.String("q", "ProcessQueue-test", "The name of the queue")
    timeout := flag.Int64("t", 5, "How long, in seconds, that the message is hidden from others")
    flag.Parse()

    if *queue == "" {
        fmt.Println("You must supply the name of a queue (-q QUEUE)")
        return
    }

    if *timeout < 0 {
        *timeout = 0
    }

    if *timeout > 12*60*60 {
        *timeout = 12 * 60 * 60
    }

    sess, err := session.NewSession(&aws.Config{
        Region: aws.String(Region),
        Credentials: credentials.NewSharedCredentials(CredPath, CredProfile)},
    )

    urlResult, err := GetQueueURL(sess, queue)
    if err != nil {
        fmt.Println("Got an error getting the queue URL:")
        fmt.Println(err)
        return
    }

    queueURL := urlResult.QueueUrl
	
	fmt.Println(urlResult.QueueUrl)

    msgResult, err := GetMessages(sess, queueURL, timeout)
    if err != nil {
        fmt.Println("Got an error receiving messages:")
        fmt.Println(err)
        return
    }

    fmt.Println("Message ID:     " + *msgResult.Messages[0].MessageId)
    fmt.Println("Message Handle: " + *msgResult.Messages[0].ReceiptHandle)
}