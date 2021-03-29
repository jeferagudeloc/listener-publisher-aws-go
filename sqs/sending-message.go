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
    Region      = "us-east-1"
    CredPath    = "/Users/jefersonagudelo/.aws/credentials"
    CredProfile = "jefersonagudeloc"
)


func GetQueueURL(sess *session.Session, queue *string) (*sqs.GetQueueUrlOutput, error) {
    svc := sqs.New(sess)

    result, err := svc.GetQueueUrl(&sqs.GetQueueUrlInput{
        QueueName: queue,
    })
    if err != nil {
        return nil, err
    }

    return result, nil
}

func SendMsg(sess *session.Session, queueURL *string) error {

    svc := sqs.New(sess)

    _, err := svc.SendMessage(&sqs.SendMessageInput{
        DelaySeconds: aws.Int64(10),
        MessageAttributes: map[string]*sqs.MessageAttributeValue{
            "Title": &sqs.MessageAttributeValue{
                DataType:    aws.String("String"),
                StringValue: aws.String("The Whistler"),
            },
            "Author": &sqs.MessageAttributeValue{
                DataType:    aws.String("String"),
                StringValue: aws.String("John Grisham"),
            },
            "WeeksOn": &sqs.MessageAttributeValue{
                DataType:    aws.String("Number"),
                StringValue: aws.String("6"),
            },
        },
        MessageBody: aws.String("Information about current NY Times fiction bestseller for week of 12/11/2016."),
        QueueUrl:    queueURL,
    })

    if err != nil {
        return err
    }

    return nil
}

func main() {

    queue := flag.String("q", "TaskQueue-test", "The name of the queue")
    flag.Parse()

    if *queue == "" {
        fmt.Println("You must supply the name of a queue (-q QUEUE)")
        return
    }

    sess, err := session.NewSession(&aws.Config{
        Region: aws.String(Region),
        Credentials: credentials.NewSharedCredentials(CredPath, CredProfile)},
    )
    
    result, err := GetQueueURL(sess, queue)
    if err != nil {
        fmt.Println("Got an error getting the queue URL:")
        fmt.Println(err)
        return
    }

    queueURL := result.QueueUrl

    err = SendMsg(sess, queueURL)
    if err != nil {
        fmt.Println("Got an error sending the message:")
        fmt.Println(err)
        return
    }

    fmt.Println("Sent message to queue ")
}