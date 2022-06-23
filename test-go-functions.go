


func tagRds (arn String , tagkey String , tagvalue String ) {

    svc := rds.New(session.New())
    input := &rds.AddTagsToResourceInput{
        ResourceName: aws.String("arn:aws:rds:us-east-1:992648334831:og:mymysqloptiongroup"),
        Tags: []*rds.Tag{
            {
                Key:   aws.String("Staging"),
                Value: aws.String("LocationDB"),
            },
        },
    }

//     result, err := svc.AddTagsToResource(input)
//     if err != nil {
//         if aerr, ok := err.(awserr.Error); ok {
//             switch aerr.Code() {
//             case rds.ErrCodeDBInstanceNotFoundFault:
//                 fmt.Println(rds.ErrCodeDBInstanceNotFoundFault, aerr.Error())
//             case rds.ErrCodeDBClusterNotFoundFault:
//                 fmt.Println(rds.ErrCodeDBClusterNotFoundFault, aerr.Error())
//             case rds.ErrCodeDBSnapshotNotFoundFault:
//                 fmt.Println(rds.ErrCodeDBSnapshotNotFoundFault, aerr.Error())
//             case rds.ErrCodeDBProxyNotFoundFault:
//                 fmt.Println(rds.ErrCodeDBProxyNotFoundFault, aerr.Error())
//         case rds.ErrCodeDBProxyTargetGroupNotFoundFault:
//             fmt.Println(rds.ErrCodeDBProxyTargetGroupNotFoundFault, aerr.Error())
//         default:
//             fmt.Println(aerr.Error())
//         }
//     } else {
//         // Print the error, cast err to awserr.Error to get the Code and
//         // Message from an error.
//         fmt.Println(err.Error())
//     }
//     return
// }

// fmt.Println(result)
}