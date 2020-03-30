//
// Copyright 2017, Sander van Harmelen
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package gitlab

import "fmt"

// ApplicationsService handles communication with administrables applications
// of the Gitlab API.
//
// Gitlab API docs : https://docs.gitlab.com/ee/api/applications.html
type ApplicationsService struct {
	client *Client
}

// CreateApplicationOptions represents the available CreateApplication() options.
//
// GitLab API docs: https://docs.gitlab.com/ce/api/applications.html#create-an-application
type CreateApplicationOptions struct {
	Name         string `json:"name,omitempty"`
	RedirectURI  string `json:"redirect_uri,omitempty"`
	Scopes       string `json:"scopes,omitempty"`
	Confidential bool   `json:"confidential,omitempty"`
}

type ListApplicationsOptions struct {
}

type Application struct {
	ID              int    `json:"id"`
	ApplicationID   string `json:"application_id"`
	ApplicationName string `json:"application_name"`
	CallbackURL     string `json:"callback_url"`
	Confidential    bool   `json:"confidential"`
}

// CreateApplication creates a new application owned by the authenticated user.
//
// Gitlab API docs : https://docs.gitlab.com/ce/api/applications.html#create-an-application
func (s *ApplicationsService) CreateApplication(opt *CreateApplicationOptions, options ...OptionFunc) (*Application, *Response, error) {
	req, err := s.client.NewRequest("POST", "applications", opt, options)
	if err != nil {
		return nil, nil, err
	}

	a := new(Application)
	resp, err := s.client.Do(req, a)
	if err != nil {
		return nil, resp, err
	}

	return a, resp, err
}

// ListApplications get a list of administrables applications by the authenticated user
//
// Gitlab API docs : https://docs.gitlab.com/ce/api/applications.html#list-all-applications
func (s *ApplicationsService) ListApplications(options ...OptionFunc) ([]*Application, *Response, error) {

	req, err := s.client.NewRequest("GET", "applications", &ListApplicationsOptions{}, options)
	if err != nil {
		return nil, nil, err
	}

	var a []*Application
	resp, err := s.client.Do(req, &a)
	if err != nil {
		return nil, resp, err
	}

	return a, resp, err
}

// DeleteApplication removes a specific application
//
// GitLab API docs: https://docs.gitlab.com/ce/api/applications.html#delete-an-application
func (s *ApplicationsService) DeleteApplication(id interface{}, options ...OptionFunc) (*Response, error) {
	application, err := parseID(id)
	if err != nil {
		return nil, err
	}

	u := fmt.Sprintf("applications/%s", pathEscape(application))

	req, err := s.client.NewRequest("DELETE", u, nil, options)
	if err != nil {
		return nil, err
	}

	return s.client.Do(req, nil)
}
