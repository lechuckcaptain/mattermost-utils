package api

import (
	mattermost "github.com/mattermost/platform/model"
	"log"
	"time"
	"math"
	"math/big"
)

func Login(client *mattermost.Client4, username string, password string) () {

	user, response := client.Login(username, password)

	if response.Error != nil {
		log.Fatal("Couldn't login: ", response.Error)
	}

	log.Printf("User '%v' logged in to '%v'. Auth Token: %s.", user.Username, client.ApiUrl, client.AuthToken)
	log.Printf("User information: %s", user.ToJson())
}

func GetTeams(client *mattermost.Client4) []*mattermost.Team {
	teams, response := client.GetAllTeams("",0,math.MaxInt64)
	if response.Error != nil {
		log.Fatal("Couldn't get team map: ", response.Error)
	}
	return teams
}

func GetChannels(client *mattermost.Client4, teamId string) []*mattermost.Channel {
	channels, response := client.GetPublicChannelsForTeam(teamId, 0,math.MaxInt64,"")
	if response.Error != nil {
		log.Fatal("Couldn't get channels: ", response.Error)
	}
	return channels
}

func GetTeamMemmbers(client *mattermost.Client4, teamId string) []*mattermost.TeamMember {
	tm, response := client.GetTeamMembers(teamId, 0, big.MaxExp, "")
	if response.Error != nil {
		log.Fatal("Couldn't get team members: ", response.Error)
	}
	return tm
}

func GetPosts(client *mattermost.Client, channelId string, start int, max int) *mattermost.PostList {
	r, e := client.GetPosts(channelId, start, max, "")
	if e != nil {
		log.Fatal("Couldn't get posts: ", e)
	}
	posts := r.Data.(*mattermost.PostList)
	return posts
}

func GetTodaysPosts(client *mattermost.Client4, channelId string) *mattermost.PostList {
	n := time.Now()
	today := time.Date(n.Year(), n.Month(), n.Day(), 0, 0, 0, 0, n.Location()).Unix()
	posts, response := client.GetPostsSince(channelId, today)
	if response.Error != nil {
		log.Fatal("Couldn't get posts: ", response.Error)
	}
	return posts
}

func GetUsersFromPosts(posts *mattermost.PostList) {

	//var m map[string]
	//
	//for _, post := range posts.Posts {
	//
	//}
}
