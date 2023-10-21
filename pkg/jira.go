package jira

import (
	"fmt"

	"github.com/andygrunwald/go-jira"
	"github.com/spf13/viper"
)

type Config struct {
	BaseURL  string `mapstructure:"base_url"`
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
}

type Issue struct {
	ID          string `json:"id"`
	Key         string `json:"key"`
	Description string `json:"description"`
	// Add other fields as needed
}

func ReadConfig(filename string) (Config, error) {
	var config Config

	viper.SetConfigFile(filename)
	if err := viper.ReadInConfig(); err != nil {
		return config, err
	}

	if err := viper.Unmarshal(&config); err != nil {
		return config, err
	}

	return config, nil
}

func NewJiraClient(config Config) (*jira.Client, error) {
	tp := jira.BasicAuthTransport{
		Username: config.Username,
		Password: config.Password,
	}

	client, err := jira.NewClient(tp.Client(), config.BaseURL)
	if err != nil {
		return nil, err
	}

	return client, nil
}

func GetIssues(client *jira.Client, projectKey string) ([]Issue, error) {
	// Create a Jira JQL query to get issues for a specific project (replace PROJECT_KEY with the actual project key)
	jqlQuery := fmt.Sprintf(`project = "%s" AND issuetype = "Bug"`, projectKey)

	// Execute the JQL query to get issues
	issues, _, err := client.Issue.Search(jqlQuery, nil)
	if err != nil {
		return nil, err
	}

	// Convert the issues to your custom Issue struct
	var customIssues []Issue
	for _, jiraIssue := range issues {
		customIssue := Issue{
			ID:          jiraIssue.ID,
			Key:         jiraIssue.Key,
			Description: jiraIssue.Fields.Description,
			// Add other fields as needed
		}
		customIssues = append(customIssues, customIssue)
	}

	return customIssues, nil
}

func UpdateIssueDescription(client *jira.Client, issueKey, newDescription string) error {
	// Create a Jira issue update request
	issueUpdate := &jira.Issue{
		Key: issueKey,
		Fields: &jira.IssueFields{
			Description: newDescription,
		},
	}

	_, err := client.Issue.Update(issueUpdate)
	if err != nil {
		return err
	}

	return nil
}
