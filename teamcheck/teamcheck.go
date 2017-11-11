package main

import (
	mattermost "github.com/mattermost/platform/model"
	"github.com/lechuckcaptain/mattermost-utils/api"
	mattermost "github.com/mattermost/platform/model"
	"log"
	"math/big"
	"os"
)

func main() {

	executable := GetExecutable()
	app := BuildCmd(executable, run)
	app.Run(os.Args)
}

func GetExecutable() string {

	return "teamcheck"
}

func run(url string, username string, password string, channelToCheck string) {

	args := os.Args
	log.Printf("Started TeamCheck\n")
	log.Printf("Recevived args: '%v'\n", args)
  
	client := mattermost.NewAPIv4Client(url)

	api.Login(client, username, password)

	teams := api.GetTeams(client)
	if len(teams) == 0 {
		log.Fatal("No teams available for current user")
	}

	PrintTeams(teams)

	for _, team := range teams {

		log.Printf("Using Team: %v (%v)", team.DisplayName, team.Id)
		channels := api.GetChannels(client,team.Id)
		PrintChannels(channels)

		teamMembers := api.GetTeamMemmbers(client,team.Id)
		log.Printf("Found %v team members", len(teamMembers))
		usersIds := make([]string, len(teamMembers))
		for i, member := range teamMembers {
			log.Printf("* %s", member.UserId)
			usersIds[i] = member.UserId
		}

		log.Print("Users ids: ", usersIds)
		users, response := client.GetUsersByIds(usersIds)
		if response.Error != nil {
			log.Fatal("Couldn't get user profiles: ", response.Error)
		}

		log.Printf("Found %v profiles", len(users))
		for _, user := range users {
			log.Printf("* %s %s", user.Id, user.Username)
		}

		for _, channel := range channels {

			if channelToCheck == channel.DisplayName {

				log.Printf("Checking todays post for channel %v", channel.DisplayName)
				posts := api.GetTodaysPosts(client, channel.Id)
				PrintPosts(posts)
			}
		}
	}
}
