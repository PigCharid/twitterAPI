package main

import (
	"context"
	"fmt"

	"github.com/michimani/gotwi"
	"github.com/michimani/gotwi/fields"
	"github.com/michimani/gotwi/tweet/timeline"
	LTtype "github.com/michimani/gotwi/tweet/timeline/types"
	"github.com/michimani/gotwi/tweet/tweetlookup"
	GTtype "github.com/michimani/gotwi/tweet/tweetlookup/types"
	"github.com/michimani/gotwi/user/userlookup"
	Utype "github.com/michimani/gotwi/user/userlookup/types"
)

func main() {
	in := &gotwi.NewClientWithAccessTokenInput{
		AccessToken: "AAAAAAAAAAAAAAAAAAAAAC0ggwEAAAAAdfBsKie9Yih4VDxgY5tT3CZ9ozE%3DEyW9HnvBs00tMXtyKyTjikqJoikrbRdRjZn8cW6vXtgwHs2RaB",
	}

	c, err := gotwi.NewClientWithAccessToken(in)
	if err != nil {
		panic(err)
	}

	userInput := &Utype.GetByUsernameInput{
		Username: "CryptoTigerBtc",
	}

	userOutput, err := userlookup.GetByUsername(context.Background(), c, userInput)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("ID:", gotwi.StringValue(userOutput.Data.ID))

	tweetsInput := &LTtype.ListTweetsInput{
		ID:         gotwi.StringValue(userOutput.Data.ID),
		MaxResults: 30,
	}

	tweetsOutput, err := timeline.ListTweets(context.Background(), c, tweetsInput)
	if err != nil {
		fmt.Println(err)
	}
	// tweetsOutput是个数组，里面有30
	fmt.Println(len(tweetsOutput.Data))
	fmt.Println(gotwi.StringValue(tweetsOutput.Data[0].ID))

	a := make([]string, 0)

	for i := 0; i < len(tweetsOutput.Data); i++ {
		a = append(a, gotwi.StringValue(tweetsOutput.Data[i].ID))
	}

	fmt.Println(a)

	tweetsinfoInput := &GTtype.ListInput{
		IDs: a,
		TweetFields: fields.TweetFieldList{
			fields.TweetFieldPublicMetrics,
		},
	}

	tweetsinfoOutput, err := tweetlookup.List(context.Background(), c, tweetsinfoInput)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(len(tweetsinfoOutput.Data))
	fmt.Println("GorillaCryptoo近30篇的推文数据")
	for k, v := range tweetsinfoOutput.Data {
		fmt.Printf("第%d篇: ", k+1)
		fmt.Print("----转推率: ", *v.PublicMetrics.RetweetCount)
		fmt.Print("----引用率: ", *v.PublicMetrics.QuoteCount)
		fmt.Print("----点赞数量: ", *v.PublicMetrics.LikeCount)
		fmt.Println("----回复量: ", *v.PublicMetrics.ReplyCount)
	}
	// fmt.Println(len(tweetsinfoOutput.Data))

}
