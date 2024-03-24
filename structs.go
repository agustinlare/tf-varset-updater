package main

import "time"

type VariablesetList struct {
	Data []struct {
		ID         string `json:"id"`
		Type       string `json:"type"`
		Attributes struct {
			Name           string    `json:"name"`
			Description    string    `json:"description"`
			Global         bool      `json:"global"`
			UpdatedAt      time.Time `json:"updated-at"`
			VarCount       int       `json:"var-count"`
			WorkspaceCount int       `json:"workspace-count"`
			ProjectCount   int       `json:"project-count"`
			Priority       bool      `json:"priority"`
		} `json:"attributes"`
		Relationships struct {
			Organization struct {
				Data struct {
					ID   string `json:"id"`
					Type string `json:"type"`
				} `json:"data"`
			} `json:"organization"`
			Vars struct {
				Data []struct {
					ID   string `json:"id"`
					Type string `json:"type"`
				} `json:"data"`
			} `json:"vars"`
		} `json:"relationships,omitempty"`
		Relationships0 struct {
			Organization struct {
				Data struct {
					ID   string `json:"id"`
					Type string `json:"type"`
				} `json:"data"`
			} `json:"organization"`
			Vars struct {
				Data []struct {
					ID   string `json:"id"`
					Type string `json:"type"`
				} `json:"data"`
			} `json:"vars"`
			Workspaces struct {
				Data []struct {
					ID   string `json:"id"`
					Type string `json:"type"`
				} `json:"data"`
			} `json:"workspaces"`
			Projects struct {
				Data []any `json:"data"`
			} `json:"projects"`
		} `json:"relationships,omitempty"`
		Relationships1 struct {
			Organization struct {
				Data struct {
					ID   string `json:"id"`
					Type string `json:"type"`
				} `json:"data"`
			} `json:"organization"`
			Vars struct {
				Data []struct {
					ID   string `json:"id"`
					Type string `json:"type"`
				} `json:"data"`
			} `json:"vars"`
			Workspaces struct {
				Data []struct {
					ID   string `json:"id"`
					Type string `json:"type"`
				} `json:"data"`
			} `json:"workspaces"`
			Projects struct {
				Data []struct {
					ID   string `json:"id"`
					Type string `json:"type"`
				} `json:"data"`
			} `json:"projects"`
		} `json:"relationships,omitempty"`
	} `json:"data"`
	Links struct {
		Self  string `json:"self"`
		First string `json:"first"`
		Prev  any    `json:"prev"`
		Next  any    `json:"next"`
		Last  string `json:"last"`
	} `json:"links"`
	Meta struct {
		Pagination struct {
			CurrentPage int `json:"current-page"`
			PageSize    int `json:"page-size"`
			PrevPage    any `json:"prev-page"`
			NextPage    any `json:"next-page"`
			TotalPages  int `json:"total-pages"`
			TotalCount  int `json:"total-count"`
		} `json:"pagination"`
	} `json:"meta"`
}

type VariableSetDetail struct {
	Data struct {
		ID         string `json:"id"`
		Type       string `json:"type"`
		Attributes struct {
			Name           string    `json:"name"`
			Description    string    `json:"description"`
			Global         bool      `json:"global"`
			UpdatedAt      time.Time `json:"updated-at"`
			VarCount       int       `json:"var-count"`
			WorkspaceCount int       `json:"workspace-count"`
			ProjectCount   int       `json:"project-count"`
			Priority       bool      `json:"priority"`
		} `json:"attributes"`
		Relationships struct {
			Organization struct {
				Data struct {
					ID   string `json:"id"`
					Type string `json:"type"`
				} `json:"data"`
			} `json:"organization"`
			Vars struct {
				Data []struct {
					ID   string `json:"id"`
					Type string `json:"type"`
				} `json:"data"`
			} `json:"vars"`
			Workspaces struct {
				Data []struct {
					ID   string `json:"id"`
					Type string `json:"type"`
				} `json:"data"`
			} `json:"workspaces"`
			Projects struct {
				Data []any `json:"data"`
			} `json:"projects"`
		} `json:"relationships"`
	} `json:"data"`
}

type VariableDetail struct {
	Data struct {
		ID         string `json:"id"`
		Type       string `json:"type"`
		Attributes struct {
			Key         string    `json:"key"`
			Value       any       `json:"value"`
			Sensitive   bool      `json:"sensitive"`
			Category    string    `json:"category"`
			Hcl         bool      `json:"hcl"`
			CreatedAt   time.Time `json:"created-at"`
			Description string    `json:"description"`
			VersionID   string    `json:"version-id"`
		} `json:"attributes"`
		Relationships struct {
			Varset struct {
				Data struct {
					ID   string `json:"id"`
					Type string `json:"type"`
				} `json:"data"`
				Links struct {
					Related string `json:"related"`
				} `json:"links"`
			} `json:"varset"`
		} `json:"relationships"`
		Links struct {
			Self string `json:"self"`
		} `json:"links"`
	} `json:"data"`
}

type AppendVarDetails struct {
	Data struct {
		Type       string `json:"type"`
		Attributes struct {
			Category    string `json:"category"`
			Key         string `json:"key"`
			Value       string `json:"value"`
			Description string `json:"description"`
			Sensitive   string `json:"sensitive"`
			Hcl         string `json:"hcl"`
		} `json:"attributes"`
	} `json:"data"`
}

type VariableList struct {
	Data struct {
		ID         string `json:"id"`
		Type       string `json:"type"`
		Attributes struct {
			Name           string    `json:"name"`
			Description    string    `json:"description"`
			Global         bool      `json:"global"`
			UpdatedAt      time.Time `json:"updated-at"`
			VarCount       int       `json:"var-count"`
			WorkspaceCount int       `json:"workspace-count"`
			ProjectCount   int       `json:"project-count"`
			Priority       bool      `json:"priority"`
		} `json:"attributes"`
		Relationships struct {
			Organization struct {
				Data struct {
					ID   string `json:"id"`
					Type string `json:"type"`
				} `json:"data"`
			} `json:"organization"`
			Vars struct {
				Data []struct {
					ID   string `json:"id"`
					Type string `json:"type"`
				} `json:"data"`
			} `json:"vars"`
			Workspaces struct {
				Data []any `json:"data"`
			} `json:"workspaces"`
			Projects struct {
				Data []any `json:"data"`
			} `json:"projects"`
		} `json:"relationships"`
	} `json:"data"`
}

type UpdateVariableDetails struct {
	Data struct {
		ID         string `json:"id"`
		Type       string `json:"type"`
		Attributes struct {
			Key         string      `json:"key"`
			Value       interface{} `json:"value"`
			Sensitive   bool        `json:"sensitive"`
			Category    string      `json:"category"`
			Hcl         bool        `json:"hcl"`
			CreatedAt   time.Time   `json:"created-at"`
			Description interface{} `json:"description"`
		} `json:"attributes"`
		Relationships struct {
			Varset struct {
				Data struct {
					ID   string `json:"id"`
					Type string `json:"type"`
				} `json:"data"`
				Links struct {
					Related string `json:"related"`
				} `json:"links"`
			} `json:"varset"`
		} `json:"relationships"`
		Links struct {
			Self string `json:"self"`
		} `json:"links"`
	} `json:"data"`
}
