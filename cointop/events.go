package cointop

// RowChanged is called when the row is updated
func (ct *Cointop) rowChanged() {
	ct.debuglog("rowChanged()")
	ct.RefreshRowLink()
}
