package google_doc

import (
	"log"
	"sort"

	"google.golang.org/api/docs/v1"
)

const (
	DOCS_ID  = "1uGTHV9hq17Y5AkOro7wWeAHVFnVhW-i9jdR76oQSU_o"
	ROW_STEP = 5
	COL_STEP = 2
)

func MakeTableOnGoogleDocs(contentTable [][]string) {
	srv, err := googleAuthorization()
	if err != nil {
		log.Fatal(err)
	}

	docInfo, err := srv.Documents.Get(DOCS_ID).Do()
	if err != nil {
		log.Fatal(err.Error())
	}
	endIndex := docInfo.Body.Content[len(docInfo.Body.Content)-1].EndIndex - 1

	docUpdateReq := createTableDocumentRequest(contentTable, endIndex)
	_, err = srv.Documents.BatchUpdate(DOCS_ID, docUpdateReq).Do()
	if err != nil {
		log.Fatal(err.Error())
	}
}

func createTableDocumentRequest(contentTable [][]string, endIndex int64) *docs.BatchUpdateDocumentRequest {
	tableWithCells := []*docs.Request{
		{
			DeleteContentRange: &docs.DeleteContentRangeRequest{
				Range: &docs.Range{
					StartIndex: 1,
					EndIndex:   endIndex,
				},
			},
		},

		{
			InsertTable: &docs.InsertTableRequest{
				Rows:    int64(len(contentTable[0])),
				Columns: int64(len(contentTable)),
				Location: &docs.Location{
					Index: 1,
				},
			},
		},
	}

	cells := fillTableCells(contentTable)
	tableWithCells = append(tableWithCells, cells...)

	docUpdateReq := &docs.BatchUpdateDocumentRequest{
		Requests: tableWithCells,
	}
	return docUpdateReq
}

func fillTableCells(contentTable [][]string) []*docs.Request {
	cells := []*docs.Request{}
	for i := len(contentTable) - 1; i >= 0; i = i - 1 {
		for j := len(contentTable[i]) - 1; j >= 0; j = j - 1 {
			cellIndex := (j+1)*ROW_STEP + i*COL_STEP
			cell := &docs.Request{InsertText: &docs.InsertTextRequest{
				Text: contentTable[i][j],
				Location: &docs.Location{
					Index: int64(cellIndex),
				},
			}}

			cells = append(cells, cell)
		}
	}

	sort.SliceStable(cells, func(i, j int) bool {
		return cells[i].InsertText.Location.Index > cells[j].InsertText.Location.Index
	})

	return cells
}
