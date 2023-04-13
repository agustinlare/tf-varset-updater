package varsetupdate

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

type VariableSet struct {
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
		} `json:"relationships"`
	} `json:"data"`
}

type Variable struct {
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

type VarList struct {
	Data []struct {
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
	Links struct {
		Self  string      `json:"self"`
		First string      `json:"first"`
		Prev  interface{} `json:"prev"`
		Next  interface{} `json:"next"`
		Last  string      `json:"last"`
	} `json:"links"`
	Meta struct {
		Pagination struct {
			CurrentPage int         `json:"current-page"`
			PageSize    int         `json:"page-size"`
			PrevPage    interface{} `json:"prev-page"`
			NextPage    interface{} `json:"next-page"`
			TotalPages  int         `json:"total-pages"`
			TotalCount  int         `json:"total-count"`
		} `json:"pagination"`
	} `json:"meta"`
}

type NewVariable struct {
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

func getVarset(varsetid string, token string) (*VariableSet, error) {
	url := fmt.Sprintf("https://app.terraform.io/api/v2/varsets/%s", varsetid)
	client := &http.Client{}
	request, _ := http.NewRequest("GET", url, nil)
	request.Header.Add("Authorization", "Bearer "+token)
	request.Header.Add("Content-Type", "application/vnd.api+json")
	resp, err := client.Do(request)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	var result VariableSet

	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

func getVars(varsetid string, token string) (*VarList, error) {
	url := fmt.Sprintf("https://app.terraform.io/api/v2/varsets/%s/relationships/vars?page%ssize%s=100", varsetid, "%5B", "%5D")
	client := &http.Client{}
	request, _ := http.NewRequest("GET", url, nil)
	request.Header.Add("Authorization", "Bearer "+token)
	request.Header.Add("Content-Type", "application/vnd.api+json")
	resp, err := client.Do(request)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	var result VarList

	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

func appendVar(varsetid string, token string, category string, sensitive string, hcl string, key string, value string) error {
	varSet, err := getVarset(varsetid, token)

	if err != nil {
		return err
	}

	var newVar NewVariable
	newVar.Data.Type = varSet.Data.Type
	newVar.Data.Attributes.Category = category
	newVar.Data.Attributes.Key = key
	newVar.Data.Attributes.Value = value
	newVar.Data.Attributes.Sensitive = sensitive
	newVar.Data.Attributes.Hcl = hcl
	data, _ := json.Marshal(newVar)

	url := fmt.Sprintf("https://app.terraform.io/api/v2/varsets/%s/relationships/vars", varsetid)
	client := &http.Client{}

	request, _ := http.NewRequest("POST", url, strings.NewReader(string(data)))
	request.Header.Add("Authorization", "Bearer "+token)
	request.Header.Add("Content-Type", "application/vnd.api+json")
	resp, err := client.Do(request)

	if err != nil {
		return err
	}

	defer resp.Body.Close()

	return nil
}

func updateVar(varsetid string, key string, value string, token string) error {
	var varNew Variable
	listVars, err := getVars(varsetid, token)

	if err != nil {
		panic(err)
	}

	for i, v := range listVars.Data {
		if v.Attributes.Key == key {
			varNew.Data = listVars.Data[i]
			break
		}
	}

	varNew.Data.Attributes.Value = value
	data, _ := json.Marshal(varNew)
	url := fmt.Sprintf("https://app.terraform.io/api/v2/vars/%s", varNew.Data.ID)
	client := &http.Client{}
	request, _ := http.NewRequest("PATCH", url, strings.NewReader(string(data)))
	request.Header.Add("Authorization", "Bearer "+token)
	request.Header.Add("Content-Type", "application/vnd.api+json")
	resp, _ := client.Do(request)

	if resp.StatusCode != http.StatusOK {
		panic(resp)
	}

	return nil
}
