package credentials

import "testing"

func TestMaskToken(t *testing.T) {
	tests := []struct {
		name     string
		token    string
		expected string
	}{
		{
			name:     "empty token",
			token:    "",
			expected: "",
		},
		{
			name:     "token shorter than 8 chars",
			token:    "1234567",
			expected: "*******",
		},
		{
			name:     "token exactly 8 chars",
			token:    "12345678",
			expected: "********",
		},
		{
			name:     "token longer than 8 chars",
			token:    "1234567890",
			expected: "1234**7890",
		},
		{
			name:     "long token with special chars",
			token:    "a^*d-1_3+-5__8-efgh",
			expected: "a^*d***********efgh",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := MaskValue(tt.token)
			if got != tt.expected {
				t.Errorf("got = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestMaskCredentials(t *testing.T) {
	creds := Credentials{
		Accounts: Accounts{
			AWS: []AWSAccount{
				{
					AccessKeyID:     "AKIA1234567890ABCDEF",
					SecretAccessKey: "abcd1234efgh5678ijkl9012mnop3456qrst",
					Region:          "us-west-2",
				},
				{
					AccessKeyID:     "AKIA0987654321ZYXWVU",
					SecretAccessKey: "zzzz9999yyyy8888xxxx7777wwww6666vvvv",
					Region:          "eu-central-1",
				},
			},
			GitHub: []GitHubAccount{
				{
					Token: "ghp_123456789abcdefghijklmnop",
				},
				{
					Token: "ghp_987654321zyxwvutsrqponm",
				},
			},
			Jira: []JiraAccount{
				{
					Email:    "user1@company.com",
					APIToken: "ATATT3xFfGF0123456789abcdefghijklmnop",
					Domain:   "company.atlassian.net",
				},
				{
					Email:    "user2@company.com",
					APIToken: "ATATT3xFfGF0987654321zyxwvutsrqponm",
					Domain:   "company.atlassian.net",
				},
			},
		},
	}

	masked := MaskCredentials(creds)
	for i, maskedAcc := range masked.Accounts.AWS {
		if maskedAcc.SecretAccessKey == creds.Accounts.AWS[i].SecretAccessKey {
			t.Errorf("AWS SecretAccessKey not masked for account %d", i)
		}
	}
	for i, maskedAcc := range masked.Accounts.GitHub {
		if maskedAcc.Token == creds.Accounts.GitHub[i].Token {
			t.Errorf("GitHub Token not masked for account %d", i)
		}
	}
	for i, maskedAcc := range masked.Accounts.Jira {
		if maskedAcc.APIToken == creds.Accounts.Jira[i].APIToken {
			t.Errorf("Jira APIToken not masked for account %d", i)
		}
	}
}
