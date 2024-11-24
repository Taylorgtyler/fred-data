package database

// GetRealGDP returns the real GDP data for the given date range
func (ctx *DBContext) GetRealGDP(startDate string, endDate string) ([]map[string]interface{}, error) {
	query := `
		select Date, GDPC1
		from main.real_gdp
		where Date between cast(? as date) and cast(? as date)
	`
	return ctx.ExecuteQuery(query, startDate, endDate)
}

// GetRealGDPPerCapita returns the real GDP per capita data for the given date range
func (ctx *DBContext) GetRealGDPPerCapita(startDate string, endDate string) ([]map[string]interface{}, error) {
	query := `
		select Date, A939RX0Q048SBEA
		from main.real_gdp_per_capita
		where Date between cast(? as date) and cast(? as date)
	`
	return ctx.ExecuteQuery(query, startDate, endDate)
}
