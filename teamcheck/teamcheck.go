package main

import (
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

func run(url string, username string, password string, channelToCheck string) {

	args := os.Args
	log.Printf("Started TeamCheck\n")
	log.Printf("Recevived args: '%v'\n", args)

	if client := api.GetClient(url, username, password); client != nil {

		teams := api.GetTeams(client)
		if len(teams) == 0 {
			log.Fatal("No teams available for current user")
		}

		api.PrintTeams(teams)

		for _, team := range teams {

			log.Printf("Selecting Team: %v (%v)", team.DisplayName, team.Id)
			client.SetTeamId(team.Id)
			channels := api.GetChannels(client)
			api.PrintChannels(channels)

			r, e := client.GetTeamMembers(team.Id, 0, big.MaxExp)
			if e != nil {
				log.Fatal("Couldn't get team members")
			}

			log.Printf(r.RequestId)
			members := r.Data.([]*mattermost.TeamMember)
			log.Printf("Found %v team members", len(members))
			usersIds := make([]string, len(members))
			for i, member := range members {
				log.Printf("* %s", member.UserId)
				usersIds[i] = member.UserId
			}

			log.Print("Users ids: ", usersIds)
			r1, e := client.GetProfilesByIds(usersIds)
			if e != nil {
				log.Fatal("Couldn't get user profiles")
			}
			users := r1.Data.(map[string]*mattermost.User)
			log.Printf("Found %v profiles", len(users))
			for _, user := range users {
				log.Printf("* %s %s", user.Id, user.Username)
			}

			for _, channel := range *channels {

				if channelToCheck == channel.DisplayName {

					log.Printf("Checking todays post for channel %v", channel.DisplayName)
					posts := api.GetTodaysPosts(client, channel.Id)
					api.PrintPosts(posts)
				}
			}
		}
	}
}
