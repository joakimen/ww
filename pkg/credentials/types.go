package credentials

func NewCredentials() Credentials {
	return Credentials{
		Accounts: Accounts{
			AWS:    []AWSAccount{},
			GitHub: []GitHubAccount{},
			Jira:   []JiraAccount{},
		},
	}
}

type Credentials struct {
	Accounts Accounts `json:"accounts"`
}

type Accounts struct {
	AWS    []AWSAccount    `json:"aws"`
	GitHub []GitHubAccount `json:"github"`
	Jira   []JiraAccount   `json:"jira"`
}

type AWSAccount struct {
	AccessKeyID     string `json:"access_key_id"`
	SecretAccessKey string `json:"secret_access_key"`
	Region          string `json:"region"`
}

type GitHubAccount struct {
	Token string `json:"token"`
}

type JiraAccount struct {
	Email    string `json:"email"`
	APIToken string `json:"api_token"`
	Domain   string `json:"domain"`
}

func DeepCopy(creds Credentials) Credentials {
	var credsCopy Credentials

	// allocate memory for all accounts
	credsCopy.Accounts.GitHub = make([]GitHubAccount, len(creds.Accounts.GitHub))
	credsCopy.Accounts.AWS = make([]AWSAccount, len(creds.Accounts.AWS))
	credsCopy.Accounts.Jira = make([]JiraAccount, len(creds.Accounts.Jira))

	// copy accounts
	copy(credsCopy.Accounts.GitHub, creds.Accounts.GitHub)
	copy(credsCopy.Accounts.AWS, creds.Accounts.AWS)
	copy(credsCopy.Accounts.Jira, creds.Accounts.Jira)

	return credsCopy
}
