package omnilayer

import (
	"bytes"
	"container/list"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"sync"
	"sync/atomic"
	"time"
)

const (
	contentType     = "Content-Type"
	contentTypeJSON = "application/json"
)

type Client struct {
	id         uint64
	config     *ConnConfig
	httpClient *http.Client

	requestLock sync.Mutex
	requestMap  map[uint64]*list.Element
	requestList *list.List

	sendChan     chan []byte
	sendPostChan chan *sendPostDetails
	disconnect   chan struct{}
	shutdown     chan struct{}
	wg           sync.WaitGroup
}

func (c *Client) WaitForShutdown() {
	c.wg.Wait()
}

func (c *Client) addRequest(jReq *jsonRequest) error {
	c.requestLock.Lock()
	defer c.requestLock.Unlock()

	select {
	case <-c.shutdown:
		return errClientShutdown()
	default:
		element := c.requestList.PushBack(jReq)
		c.requestMap[jReq.id] = element
		return nil
	}
}

func (c *Client) do(cmd command) chan *response {
	// Get the method associated with the command.
	method := cmd.Method()

	// Marshal the command.
	id := c.NextID()

	marshalledJSON, err := marshalCmd(cmd)
	if err != nil {
		return newFutureError(err)
	}

	// Generate the request and send it along with a channel to respond on.
	responseChan := make(chan *response, 1)
	jReq := &jsonRequest{
		id:             id,
		method:         method,
		cmd:            cmd,
		marshalledJSON: marshalledJSON,
		responseChan:   responseChan,
	}

	c.sendPost(jReq)

	return responseChan
}

func (c *Client) sendPost(jReq *jsonRequest) {
	req, err := http.NewRequest(http.MethodPost, "http://"+c.config.Host, bytes.NewReader(jReq.marshalledJSON))
	if err != nil {
		jReq.responseChan <- &response{result: nil, err: err}
		return
	}

	req.Close = true
	req.Header.Set(contentType, contentTypeJSON)
	req.SetBasicAuth(c.config.User, c.config.Pass)

	c.post(req, jReq)
}

func (c *Client) post(httpReq *http.Request, jReq *jsonRequest) {
	// Don't send the message if shutting down.
	select {
	case <-c.shutdown:
		jReq.responseChan <- &response{result: nil, err: errClientShutdown()}
	default:
	}

	c.sendPostChan <- &sendPostDetails{
		jsonRequest: jReq,
		httpRequest: httpReq,
	}
}

func New(config *ConnConfig) *Client {
	httpClient := newHTTPClient()

	client := &Client{
		config:       config,
		httpClient:   httpClient,
		requestMap:   make(map[uint64]*list.Element),
		requestList:  list.New(),
		sendChan:     make(chan []byte, sendBufferSize),
		sendPostChan: make(chan *sendPostDetails, sendPostBufferSize),
		disconnect:   make(chan struct{}),
		shutdown:     make(chan struct{}),
	}

	client.asyncStart()

	return client
}

func (c *Client) asyncStart() {
	c.wg.Add(1)
	go c.sendPostHandler()
}

func (c *Client) sendPostHandler() {
out:
	for {
		select {
		case details := <-c.sendPostChan:
			c.handleSendPostMessage(details)

		case <-c.shutdown:
			break out
		}
	}

cleanup:
	for {
		select {
		case details := <-c.sendPostChan:
			details.jsonRequest.responseChan <- &response{
				result: nil,
				err:    errClientShutdown(),
			}

		default:
			break cleanup
		}
	}

	c.wg.Done()
}

func (c *Client) handleSendPostMessage(details *sendPostDetails) {
	jReq := details.jsonRequest
	httpResponse, err := c.httpClient.Do(details.httpRequest)
	if err != nil {
		jReq.responseChan <- &response{err: err}
		return
	}

	respBytes, err := ioutil.ReadAll(httpResponse.Body)
	if err != nil {
		jReq.responseChan <- &response{err: err}
		return
	}
	err = httpResponse.Body.Close()
	if err != nil {
		jReq.responseChan <- &response{err: err}
		return
	}

	// Try to unmarshal the response as a regular JSON-RPC response.
	var resp rawResponse
	err = json.Unmarshal(respBytes, &resp)
	if err != nil {
		jReq.responseChan <- &response{err: err}
		return
	}

	res, err := resp.result()
	jReq.responseChan <- &response{result: res, err: err}
}

func newHTTPClient() *http.Client {
	return &http.Client{
		Transport: &http.Transport{
			ResponseHeaderTimeout: 5 * time.Second,
			ExpectContinueTimeout: 4 * time.Second,
			IdleConnTimeout:       5 * 60 * time.Second,
		},
	}
}

func (c *Client) NextID() uint64 {
	return atomic.AddUint64(&c.id, 1)
}

func (c *Client) Shutdown() {
	c.requestLock.Lock()
	defer c.requestLock.Unlock()

	select {
	case <-c.shutdown:
		return
	default:
	}

	close(c.shutdown)

	for e := c.requestList.Front(); e != nil; e = e.Next() {
		req, ok := e.Value.(*jsonRequest)
		if !ok {
			continue
		}

		req.responseChan <- &response{
			result: nil,
			err:    errClientShutdown(),
		}
	}

	c.removeAllRequests()
}

func (c *Client) removeAllRequests() {
	c.requestMap = make(map[uint64]*list.Element)
	c.requestList.Init()
}
