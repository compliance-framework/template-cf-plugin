package main

//GITHUB TEMPLATE INSTRUCTIONS: Go through all lines with the prefix 'GITHUB TEMPLATE INSTRUCTIONS' and do what it tells you to. When happy, remove the comment lines beginning 'GITHUB TEMPLATE INSTRUCTIONS'.

import (
	"fmt"

	"time"

	. "github.com/compliance-framework/assessment-runtime/provider"
	"github.com/google/uuid"
	"gopkg.in/yaml.v2"
)

//GITHUB TEMPLATE INSTRUCTIONS: Replace 'Template' throughout this file with your chosen plugin name
type TemplateProvider struct {
    message string
}

type YamlConfig struct {
	// GITHUB TEMPLATE INSTRUCTIONS: Replace this struct with config passed in through the 'yaml' item
    SomeConfigItem string `json:"config_item" yaml:"config_item"`
}

func (p *TemplateProvider) Evaluate(input *EvaluateInput) (*EvaluateResult, error) {
    var yaml_config YamlConfig

	// Extract the passed-in YAML into code here.
    yamlString, ok := input.Configuration["yaml"]
    if !ok {
        return nil, fmt.Errorf("yaml parameter is missing")
    }
    err := yaml.Unmarshal([]byte(yamlString), &yaml_config)
    if err != nil {
        return nil, fmt.Errorf("Error unmarshalling YAML: %v\n", err)
    }
    log.Printf("yamlString: %s", yamlString)

	// GITHUB TEMPLATE INSTRUCTIONS: If config item is needed from external call, change it here
	//someConfigItem := yaml_config.SomeConfigItem

	// GITHUB TEMPLATE INSTRUCTIONS: If env var secrets are needed, process them here.
    //// Get environment variable for the secret
    //secret := os.Getenv("SECRET")

    //if secret == "" {
    //    return nil, fmt.Errorf("One or more environment variables are not set")
    //}
    //if !ok {
    //    return nil, fmt.Errorf("secret is missing")
    //}

	// There can be an array of subjects if needed, but here we have only one
	subjects := make([]*Subject, 0)
	subject_id := fmt.Sprintf("Subject identifier: %s", someConfigItem)  // GITHUB TEMPLATE INSTRUCTIONS: Create an identifier for the subject of the compliance activity
	subjects = append(subjects, &Subject{
		Id:    subject_id,
		Type:  SubjectType_INVENTORY_ITEM,
		Title: "Subject title",  // GITHUB TEMPLATE INSTRUCTIONS: Replace this with your subject title
		Props: map[string]string{
			"id": subject_id,
		},
	})

	// Return the result with subjects and additional props if necessary
	return &EvaluateResult{
		Subjects: subjects,
	}, nil
}

func (p TemplateProvider) Execute(input *ExecuteInput) (*ExecuteResult, error) {
	start_time := time.Now().Format(time.RFC3339)

	var obs *Observation
	var fndngs *Finding

	observations := []*Observation{}
	findings := []*Finding{}

	obs_id := uuid.New().String()

	//GITHUB TEMPLATE INSTRUCTIONS: put the code that determines whether we are in
	//                              compliance or not here, replacing 'true' with the logic required
	compliant := true

	if (!compliant) {
		// observation and finding
		obs = &Observation{
			Id:               obs_id,
			Title:            "Title of Observation",  // GITHUB TEMPLATE INSTRUCTIONS: Replace with your title
			Description:      "Description of the observation that did not succeed",  // GITHUB TEMPLATE INSTRUCTIONS: Replace with your description
			Collected:        time.Now().Format(time.RFC3339),
			Expires:          time.Now().AddDate(0, 1, 0).Format(time.RFC3339),  // GITHUB TEMPLATE INSTRUCTIONS: Add a time period for the expiration
			Links:            []*Link{},
			Props:            []*Property{
				// GITHUB TEMPLATE INSTRUCTIONS: Add any properties that may be passed in the non-secret 'configuration'
				//                               of the plan that is passed into the configuration-manager
				//{
				//	Name:  "Command",
				//	Value: fmt.Sprintf("%s", command),
				//},
			},
			RelevantEvidence: []*Evidence{
				{
					Description: "Description of the evidence that resulted in this finding",  // GITHUB TEMPLATE INSTRUCTIONS: Replace with your description
				},
			},
			Remarks:          "Any relevant remarks",  // GITHUB TEMPLATE INSTRUCTIONS: Replace with your remarks
		}
		fndngs = &Finding{
			Id:                  uuid.New().String(),
			Title:               "Title of finding",  // GITHUB TEMPLATE INSTRUCTIONS: Replace with your finding title
			Description:         "Description of finding",  //GITHUB TEMPLATE INSTRUCTIONS: Replace with your finding description
			Remarks:             "Any relevant remarks",  //GITHUB TEMPLATE INSTRUCTIONS: Replace with your remarks
			RelatedObservations: []string{obs_id},
		}
		observations = append(observations, obs)
		findings = append(findings, fndngs)
	} else {
		// observation only
		obs = &Observation{
			Id:          obs_id,
			Title:       "Title of Observation",  // GITHUB TEMPLATE INSTRUCTIONS: Replace with your title
			Description: "Description of the observation that succeeded",  // GITHUB TEMPLATE INSTRUCTIONS: Replace with your description
			Collected:   time.Now().Format(time.RFC3339),
			Expires:     time.Now().AddDate(0, 1, 0).Format(time.RFC3339),  // GITHUB TEMPLATE INSTRUCTIONS: Add a time period for the expiration
			Links:       []*Link{},
			Props: []*Property{
				// GITHUB TEMPLATE INSTRUCTIONS: Add any properties that may be passed in the non-secret 'configuration'
				//                               of the plan that is passed into the configuration-manager
				//{
				//	Name:  "Command",
				//	Value: fmt.Sprintf("%s", command),
				//},
			},
			RelevantEvidence: []*Evidence{
				{
					Description: "Description of the evidence that resulted in this finding",  // GITHUB TEMPLATE INSTRUCTIONS: Replace with your description
				},
			},
			Remarks: "All OK.",  // GITHUB TEMPLATE INSTRUCTIONS: Replace with your remarks
		}
		observations = append(observations, obs)
	}

	// Log that the check has successfully run
	logEntry := &LogEntry{
		Title:       "Log entry title",  // GITHUB TEMPLATE INSTRUCTIONS: Replace with your log entry title
		Description: "Log entry description",  // GITHUB TEMPLATE INSTRUCTIONS: Replace with your log entry description
		Start:       start_time,
		End:         time.Now().Format(time.RFC3339),
	}

	// Return the result
	return &ExecuteResult{
		Status:       ExecutionStatus_SUCCESS,
		Observations: observations,
		Findings:     findings,
		Logs:         []*LogEntry{logEntry},
	}, nil
}

func main() {
	Register(&TemplateProvider{
		message: "Template completed",
	})
}
