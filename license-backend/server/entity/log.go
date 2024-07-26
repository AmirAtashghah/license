package entity

type Log struct {
	ID          string `json:"id"`
	ClientHash  string `json:"client_hash"`
	Time        string `json:"time"`
	Level       string `json:"level"`
	Message     string `json:"msg"`
	Group       string `json:"group"`
	Path        string `json:"path"`
	Line        int    `json:"line"`
	Function    string `json:"function"`
	RequestBody string `json:"request_body"`
	Code        int    `json:"code"`
}
