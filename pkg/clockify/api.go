package clockify

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"

	"encoding/json"

	"github.com/sirupsen/logrus"
)

// APIClient for api requests with rate limit counter
type APIClient struct {
	BaseURL    *url.URL
	ReportURL  *url.URL
	UserAgent  string
	httpClient *http.Client
	key        string
	log        *logrus.Entry

	rateLimit        time.Duration
	rateLimitChannel chan time.Time

	User      *UserService
	UserGroup *UserGroupService
	Client    *ClientService
	TimeEntry *TimeEntryService
	Tag       *TagService
	Project   *ProjectService
	Task      *TaskService
	Workspace *WorkspaceService
	Report    *ReportService
}

// NewAPIClient return a client with all services and ratelimit channel
func NewAPIClient(endpoint string, reportEndpoint string, accessKey string, httpClient *http.Client, log *logrus.Entry) (*APIClient, error) {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}

	endpointURL, err := url.Parse(fmt.Sprintf("%s/", endpoint))
	if err != nil {
		log.Fatal(err)
		return nil, NewInternalClientError("endpoint url parsing failed.").WithInternalError(err)
	}

	reportURL, err := url.Parse(fmt.Sprintf("%s/", reportEndpoint))
	if err != nil {
		log.Fatal(err)
		return nil, NewInternalClientError("report endpoint url parsing failed.").WithInternalError(err)
	}

	log.Trace("init client..")

	c := &APIClient{
		httpClient: httpClient,
		key:        accessKey,
		log:        log,

		rateLimit:        time.Second / 10,
		rateLimitChannel: make(chan time.Time, 10),

		BaseURL:   endpointURL,
		ReportURL: reportURL,
		UserAgent: "pkuebler/clockify-cli",
	}

	log.Trace("add services..")
	c.User = &UserService{client: c}
	c.UserGroup = &UserGroupService{client: c}
	c.Client = &ClientService{client: c}
	c.TimeEntry = &TimeEntryService{client: c}
	c.Tag = &TagService{client: c}
	c.Project = &ProjectService{client: c}
	c.Task = &TaskService{client: c}
	c.Workspace = &WorkspaceService{client: c}
	c.Report = &ReportService{client: c}

	return c, nil
}

// StartRatelimit adds new api requests
func (c *APIClient) StartRatelimit(ctx context.Context) {
	c.log.Debug("start ratelimiter..")

	go func() {
		for i := 0; i < 10; i++ {
			c.rateLimitChannel <- time.Now()
		}

		ticker := time.NewTicker(c.rateLimit)
		defer ticker.Stop()

		for t := range ticker.C {
			select {
			case c.rateLimitChannel <- t:
				c.log.Trace("add ratelimit +1 free request..")
			case <-ctx.Done():
				c.log.Debug("stopped ratelimiter..")
				return
			}
		}
	}()
}

func (c *APIClient) newReportRequest(method string, path string, query string, body interface{}) (*http.Request, error) {
	return c.newRequest(c.ReportURL, method, path, query, body)
}

func (c *APIClient) newAPIRequest(method string, path string, query string, body interface{}) (*http.Request, error) {
	return c.newRequest(c.BaseURL, method, path, query, body)
}

func (c *APIClient) newRequest(baseURL *url.URL, method string, path string, query string, body interface{}) (*http.Request, error) {
	rel := &url.URL{Path: path, RawQuery: query}
	u := baseURL.ResolveReference(rel)

	var buf io.ReadWriter
	if body != nil {
		buf = new(bytes.Buffer)
		err := json.NewEncoder(buf).Encode(body)
		if err != nil {
			c.log.Fatal(err)
			return nil, err
		}
	}

	req, err := http.NewRequest(method, u.String(), buf)
	if err != nil {
		c.log.Fatal(err)
		return nil, err
	}
	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("User-Agent", c.UserAgent)
	req.Header.Set("X-Api-Key", c.key)

	return req, nil
}

func (c *APIClient) do(req *http.Request, v interface{}) (*http.Response, error) {
	<-c.rateLimitChannel

	res, err := c.httpClient.Do(req)
	if err != nil {
		c.log.Fatal(err)
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode < 200 || res.StatusCode >= 300 {
		return res, errors.New(http.StatusText(res.StatusCode))
	}

	err = json.NewDecoder(res.Body).Decode(v)
	return res, err
}
