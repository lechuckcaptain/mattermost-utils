package main

import (
	"net/url"
	"log"
	mattermost "github.com/mattermost/platform/model"
)

func CheckServer(server string) bool {

	_, err := url.ParseRequestURI(server)
	if err != nil {
		return false
	} else {
		return true
	}
}

func PrintTeams(teams []*mattermost.Team) {
	log.Printf("Found %v teams:", len(teams))
	for i, team := range teams {
		log.Printf("* Team #%d: %s -> %s", i, team.Id, team.DisplayName)
	}
}

func PrintChannels(channels []*mattermost.Channel) {
	log.Printf("Found %v channels:", len(channels))
	for i, channel := range channels {
		log.Printf("* Channeld #%d: %s -> %s", i, channel.Id, channel.DisplayName)
	}
}

func PrintPosts(posts *mattermost.PostList) {
	log.Print("Posts:")
	for _, post := range posts.Posts {
		log.Printf("%s -> '%s' by %v @ %v", post.Id, post.Message, post.UserId, post.CreateAt)
	}
}

