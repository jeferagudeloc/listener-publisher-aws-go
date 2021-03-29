package main

import (
	"fmt"
	"encoding/json"
	"context"
	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
  )


var (
    personalAccessToken = "28fa1447945bddeee00c64cfcec47e5cccaeb5ba"
)

type TokenSource struct {
    AccessToken string
}

func (t *TokenSource) Token() (*oauth2.Token, error) {
    token := &oauth2.Token{
        AccessToken: t.AccessToken,
    }
    return token, nil
}

func main(){
	ctx := context.Background()
    tokenSource := &TokenSource{
        AccessToken: personalAccessToken,
    }
    oauthClient := oauth2.NewClient(oauth2.NoContext, tokenSource)
    client := github.NewClient(oauthClient)
    user, _, err := client.Users.Get(ctx,"jeferagudeloc")
    if err != nil {
        fmt.Printf("client.Users.Get() faled with '%s'\n", err)
        return
    }
    d, err := json.MarshalIndent(user, "", "  ")
    if err != nil {
        fmt.Printf("json.MarshlIndent() failed with %s\n", err)
        return
    }
	fmt.Printf("User:\n%s\n", string(d))
	

	if err != nil {
		fmt.Print(err)
	}

	repo := &github.Repository{
		Name:    github.String("foo"),
		Private: github.Bool(true),
	}
	client.Repositories.Create(ctx, "", repo)
}