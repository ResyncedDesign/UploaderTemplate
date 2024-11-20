package types

type JSONResponse struct {
	Success bool   `json:"success"`
	Files   []File `json:"files"`
}

type File struct {
	Name      string `json:"name"`
	URL       string `json:"url"`
	DeleteURL string `json:"delete_url"` // If you use sharex to upload files, this can be useful since it allows the uploader to delete the file
}
