package api

import (
	mattermost "github.com/mattermost/platform/model"
	"log"
	"time"
)

func GetClient(url string, username string, password string) *mattermost.Client {

	client := mattermost.NewClient(url)
	r, e := client.Login(username, password)
	if e != nil {
		log.Fatal("Couldn't login: ", e)
	}
	log.Printf("Client logged in to '%v'. Auth Token: %s.", url, client.AuthToken)
	user := r.Data.(*mattermost.User)
	log.Printf("User information: %s", user.ToJson())
	return client
}

func PrintTeams(teamMap map[string]*mattermost.Team) {
	log.Printf("Found %v teams:", len(teamMap))
	for teamId, team := range teamMap {
		log.Printf("* %s -> %s", teamId, team.DisplayName)
	}
}

func PrintChannels(channelList *mattermost.ChannelList) {
	log.Printf("Found %v channels:", len(*channelList))
	for _, channel := range *channelList {
		log.Printf("* %s -> %s", channel.Id, channel.DisplayName)
	}
}

func PrintPosts(posts *mattermost.PostList) {
	log.Print("Posts:")
	for _, post := range posts.Posts {
		log.Printf("%s -> '%s' by %v @ %v", post.Id, post.Message, post.UserId, post.CreateAt)
	}
}

func GetTeams(client *mattermost.Client) map[string]*mattermost.Team {
	r, e := client.GetAllTeams()
	if e != nil {
		log.Fatal("Couldn't get team map: ", e)
	}
	teamMap := r.Data.(map[string]*mattermost.Team)
	return teamMap
}

func GetChannels(client *mattermost.Client) *mattermost.ChannelList {
	r, e := client.GetChannels("")
	if e != nil {
		log.Fatal("Couldn't get channels: ", e)
	}
	return r.Data.(*mattermost.ChannelList)
}

func GetPosts(client *mattermost.Client, channelId string, start int, max int) *mattermost.PostList {
	r, e := client.GetPosts(channelId, start, max, "")
	if e != nil {
		log.Fatal("Couldn't get posts: ", e)
	}
	posts := r.Data.(*mattermost.PostList)
	return posts
}

func GetTodaysPosts(client *mattermost.Client, channelId string) *mattermost.PostList {
	n := time.Now()
	today := time.Date(n.Year(), n.Month(), n.Day(), 0, 0, 0, 0, n.Location()).Unix()
	r, e := client.GetPostsSince(channelId, today)
	if e != nil {
		log.Fatal("Couldn't get posts: ", e)
	}
	posts := r.Data.(*mattermost.PostList)
	return posts
}

func GetUsersFromPosts(posts *mattermost.PostList) {

	//var m map[string]
	//
	//for _, post := range posts.Posts {
	//
	//}
}
