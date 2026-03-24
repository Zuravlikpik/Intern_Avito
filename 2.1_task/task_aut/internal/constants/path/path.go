package path

import "fmt"

const (
	CreateItemPath   = "/api/1/item"
	GetItemByIDPath  = "/api/1/item/%s"
	GetItemsBySeller = "/api/1/%d/item"
	GetStatisticByID = "/api/1/statistic/%s"
)

// возвращает отформатированный путь для получения объявлений по sellerID
func FormatGetItemsBySeller(sellerID int) string {
	return fmt.Sprintf(GetItemsBySeller, sellerID)
}
