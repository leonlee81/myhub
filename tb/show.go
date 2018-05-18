package tb

import "strings"


const(
	SHOW_TABLES = "tables"
	SHOW_FIELDS = "fields"
	SHOW_KEYS = "keys"
	SHOW_CREATE = "create"
)

type Show struct {
	Full bool
	ExprStr string
	From string
}

func ParseShowStmt(query string) *Show{
	pShow := &Show{}
	query = strings.Replace(query,"`","",-1)
	query = strings.ToLower(query)
	tokens := strings.Split(query," ")
	//
	var findFrom bool;
	for _, token := range tokens{
		switch token {
		case "full":
			pShow.Full = true
		case SHOW_TABLES,SHOW_FIELDS,SHOW_KEYS,SHOW_CREATE:
			pShow.ExprStr = token
		case "from":
			findFrom = true;
		default:
			if findFrom && len(pShow.From) < 1{
				pShow.From = token
			}
		}
	}
	return pShow;
}
func (this *Show) IsShowTables() bool {
	if this.ExprStr == SHOW_TABLES{
		return true;
	}
	return false;
}
func (this *Show) IsShowFields() bool {
	if this.ExprStr == SHOW_FIELDS{
		return true;
	}
	return false;
}
func (this *Show) IsShowKeys() bool {
	if this.ExprStr == SHOW_KEYS{
		return true;
	}
	return false;
}
//
func (this *Show) GetFromDataBase() string {
	if this.IsShowTables(){
		return this.From;
	}
	arr := strings.Split(this.From,".")
	return arr[0];
}
//
func (this *Show) GetFromTable() string {
	if this.IsShowTables(){
		return "";
	}
	arr := strings.Split(this.From,".")
	if len(arr) > 1{
		return arr[1]
	}
	return arr[0];
}