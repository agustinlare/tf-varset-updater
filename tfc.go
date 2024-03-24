package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

func SendRequest(method, url, token string, body io.Reader) (*http.Response, error) {
	client := &http.Client{}
	request, _ := http.NewRequest(method, url, body)
	request.Header.Add("Authorization", "Bearer "+token)
	request.Header.Add("Content-Type", "application/vnd.api+json")

	return client.Do(request)
}

// @Summary Retrive list of a variable set
// @Description Gather a list of variable sets with method "GET"
// @Tags TerraformCloud
// @Accept json
// @Produce json
// @Param token header string true "Token"
// @Param organization header string true "Organization Name"
// @Success 200 {object} VariablesetList "OK"
// @Router /listvariablesets [get]
func ListVariableSet(c *gin.Context) {
	token := c.GetHeader("token")
	organization := c.GetHeader("organization")

	resp, err := SendRequest("GET", fmt.Sprintf("%s/organizations/%s/varsets", TerraformApi, organization), token, nil)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if resp.StatusCode != http.StatusOK {
		c.JSON(resp.StatusCode, gin.H{"error": resp.Status})
		return
	}

	defer resp.Body.Close()

	var varsetListResponse VariablesetList
	if err := json.NewDecoder(resp.Body).Decode(&varsetListResponse); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, varsetListResponse)
}

// @Summary Shows the details of a specify variableset
// @Description Gather details of a variableset
// @Tags TerraformCloud
// @Accept json
// @Produce json
// @Param token header string true "Token"
// @Param varset_id path string true "Variableset ID"
// @Success 200 {object} VariableSetDetail "OK"
// @Router /showvariableset/{varset_id} [get]
func ShowVariableSet(c *gin.Context) {
	token := c.GetHeader("token")
	id := c.Param("varset_id")

	resp, err := SendRequest("GET", fmt.Sprintf("%s/varsets/%s", TerraformApi, id), token, nil)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if resp.StatusCode != http.StatusOK {
		c.JSON(resp.StatusCode, gin.H{"error": resp.Status})
		return
	}

	defer resp.Body.Close()

	var variableSetDetail VariableSetDetail
	if err := json.NewDecoder(resp.Body).Decode(&variableSetDetail); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, variableSetDetail)
}

// @Summary Shows details ofa specific variable
// @Description Gather details of a variable
// @Tags TerraformCloud
// @Accept json
// @Produce json
// @Param token header string true "Token"
// @Param var_id path string true "ID del conjunto de variables"
// @Success 200 {object} VariableSetDetail "OK"
// @Router /showvariable/{var_id} [get]
func ShowVariable(c *gin.Context) {
	token := c.GetHeader("token")
	id := c.Param("var_id")
	resp, err := SendRequest("GET", fmt.Sprintf("%s/vars/%s", TerraformApi, id), token, nil)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if resp.StatusCode != http.StatusOK {
		c.JSON(resp.StatusCode, gin.H{"error": resp.Status})
		return
	}

	defer resp.Body.Close()
	var VariableDetail VariableDetail
	if err := json.NewDecoder(resp.Body).Decode(&VariableDetail); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, VariableDetail)
}

// @Summary Agrega una nueva variable al conjunto de variables especificado
// @Description Agrega una nueva variable al conjunto de variables especificado
// @Tags TerraformCloud
// @Accept x-www-form-urlencoded
// @Param token header string true "Token de autenticación"
// @Param varset_id path string true "ID del conjunto de variables"
// @Param category formData string true "Categoría de la variable (env/terraform)"
// @Param sensitive formData string true "Indica si la variable es sensible (true/false)"
// @Param hcl formData string true "Indica si la variable es de tipo HCL (true/false)"
// @Param key formData string true "Clave de la variable"
// @Param value formData string true "Valor de la variable"
// @Success 200 {string} string "OK"
// @Router /appendvariable/{varset_id} [post]
func AppendVariable(c *gin.Context) {
	token := c.GetHeader("token")
	varsetid := c.Param("varset_id")
	category := c.PostForm("category")
	sensitive := c.PostForm("sensitive")
	hcl := c.PostForm("hcl")
	key := c.PostForm("key")
	value := c.PostForm("value")

	variableSet, err := SendRequest("GET", fmt.Sprintf("%s/varsets/%s", TerraformApi, varsetid), token, nil)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if variableSet.StatusCode != http.StatusOK {
		c.JSON(variableSet.StatusCode, gin.H{"error": variableSet.Status})
		return
	}

	defer variableSet.Body.Close()

	var variableSetDetail VariableSetDetail
	if err := json.NewDecoder(variableSet.Body).Decode(&variableSetDetail); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var appendableVariable AppendVarDetails
	appendableVariable.Data.Type = variableSetDetail.Data.Type
	appendableVariable.Data.Attributes.Category = category
	appendableVariable.Data.Attributes.Key = key
	appendableVariable.Data.Attributes.Value = value
	appendableVariable.Data.Attributes.Sensitive = sensitive
	appendableVariable.Data.Attributes.Hcl = hcl

	body, err := json.Marshal(appendableVariable)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	resp, err := SendRequest("POST", fmt.Sprintf("%s/varsets/%s/relationships/vars", TerraformApi, varsetid), token, strings.NewReader(string(body)))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		c.JSON(http.StatusBadRequest, gin.H{"error": resp.Status})
		return
	}

	c.String(http.StatusOK, "OK")
}

// @Summary Update a variable value
// @Description Update a variable value
// @Tags TerraformCloud
// @Accept x-www-form-urlencoded
// @Param token header string true "Token"
// @Param varset_id path string true "Variableset ID"
// @Param key formData string true "Variable key"
// @Param value formData string true "Variable value"
// @Success 200 {string} string "OK"
// @Router /updatevariable/{varset_id} [patch]
func UpdateVariable(c *gin.Context) {
	token := c.GetHeader("token")
	varsetid := c.Param("varset_id")
	key := c.PostForm("key")
	value := c.PostForm("value")

	variableSet, err := SendRequest("GET", fmt.Sprintf("%s/varsets/%s", TerraformApi, varsetid), token, nil)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer variableSet.Body.Close()

	if variableSet.StatusCode != http.StatusOK {
		c.JSON(variableSet.StatusCode, gin.H{"error": variableSet.Status})
		return
	}

	var result VariableList
	if err := json.NewDecoder(variableSet.Body).Decode(&result); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	for _, v := range result.Data.Relationships.Vars.Data {
		var VariableDetail VariableDetail

		resp, err := SendRequest("GET", fmt.Sprintf("%s/vars/%s", TerraformApi, v.ID), token, nil)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			c.JSON(resp.StatusCode, gin.H{"error": resp.Status})
			return
		}

		if err := json.NewDecoder(resp.Body).Decode(&VariableDetail); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		if VariableDetail.Data.Attributes.Key == key {
			VariableDetail.Data.Attributes.Value = value

			var requestData AppendVarDetails
			requestData.Data.Type = "vars"
			requestData.Data.Attributes.Category = VariableDetail.Data.Attributes.Category
			requestData.Data.Attributes.Key = VariableDetail.Data.Attributes.Key
			requestData.Data.Attributes.Value = value
			requestData.Data.Attributes.Description = VariableDetail.Data.Attributes.Description
			requestData.Data.Attributes.Sensitive = strconv.FormatBool(VariableDetail.Data.Attributes.Sensitive)
			requestData.Data.Attributes.Hcl = strconv.FormatBool(VariableDetail.Data.Attributes.Hcl)

			jsonData, err := json.Marshal(requestData)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}

			resp, err := SendRequest("PATCH", fmt.Sprintf("%s/vars/%s", TerraformApi, v.ID), token, bytes.NewReader(jsonData))
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
			defer resp.Body.Close()

			if resp.StatusCode != http.StatusOK {
				c.JSON(http.StatusBadRequest, gin.H{"error": resp.Status})
				return
			}
		}
	}

	c.String(http.StatusOK, "OK")
}
