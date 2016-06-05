package index

import "github.com/blevesearch/bleve"

// Index abstracts the underlying index technology
// TODO: AWS CloudSearch support
type Index struct {
	index bleve.Index
}

// NewIndex connects to or creates the index persistence
func NewIndex(indexPath string) (*Index, error) {

	var err error
	// try to open the persistence file
	index, err := bleve.Open(indexPath)
	if err != nil {
		// create a new mapping file and create a new index
		mapping := bleve.NewIndexMapping()
		index, err = bleve.New(indexPath, mapping)
		if err != nil {
			return nil, err
		}
	}

	return &Index{index: index}, nil
}

// Add uses the underlying index tech to analyze, index or store mapped data
// fields. The supplied identifier is bound to analyzed data and will be
// retrieved by search requests.
func (idx *Index) Add(id string, data interface{}) error {
	err := idx.index.Index(id, data)
	return err
}

// Search creates a new Query used for finding documents that satisfy a
// query string.
func (idx *Index) Search(term string) (*bleve.SearchResult, error) {
	// search for some text

	//query := bleve.NewMatchQuery(term)
	query := bleve.NewQueryStringQuery(term)
	searchQuery := bleve.NewSearchRequest(query)

	// Return all fields...
	searchQuery.Fields = append(searchQuery.Fields, "*")

	result, err := idx.index.Search(searchQuery)

	return result, err
}
