package database

// GetRealGDP returns the real GDP data for the given date range
func (ctx *DBContext) GetRealGDP(startDate string, endDate string) ([]map[string]interface{}, error) {
	query := `
		select Date, GDPC1
		from main.real_gdp
		where Date between cast(? as date) and cast(? as date)
	`
	return ctx.ExecuteQuery(query, []interface{}{startDate, endDate})
}

// GetRealGDPPerCapita returns the real GDP per capita data for the given date range
func (ctx *DBContext) GetRealGDPPerCapita(startDate string, endDate string) ([]map[string]interface{}, error) {
	query := `
		select Date, A939RX0Q048SBEA
		from main.real_gdp_per_capita
		where Date between cast(? as date) and cast(? as date)
	`
	return ctx.ExecuteQuery(query, []interface{}{startDate, endDate})
}

// GetFederalFundsEffectiveRate returns the federal funds effective rate data for the given date range
func (ctx *DBContext) GetFederalFundsEffectiveRate(startDate string, endDate string) ([]map[string]interface{}, error) {
	query := `
		select Date, DFF
		from main.federal_funds_effective_rate
		where Date between cast(? as date) and cast(? as date)
	`
	return ctx.ExecuteQuery(query, []interface{}{startDate, endDate})
}

// GetLaborForceParticipationRate returns the labor force participation rate data for the given date range
func (ctx *DBContext) GetLaborForceParticipationRate(startDate string, endDate string) ([]map[string]interface{}, error) {
	query := `
		select Date, CIVPART
		from main.labor_force_participation_rate
		where Date between cast(? as date) and cast(? as date)
	`
	return ctx.ExecuteQuery(query, []interface{}{startDate, endDate})
}

// GetUnemploymentRate returns the unemployment rate data for the given date range
func (ctx *DBContext) GetUnemploymentRate(startDate string, endDate string) ([]map[string]interface{}, error) {
	query := `
		select Date, UNRATE
		from main.unemployment_rate
		where Date between cast(? as date) and cast(? as date)
	`
	return ctx.ExecuteQuery(query, []interface{}{startDate, endDate})
}

// GetRealMedianPersonalIncome returns the real median personal income data for the given date range
func (ctx *DBContext) GetRealMedianPersonalIncome(startDate string, endDate string) ([]map[string]interface{}, error) {
	query := `
		select Date, MEPAINUSA672N
		from main.real_median_personal_income
		where Date between cast(? as date) and cast(? as date)
	`
	return ctx.ExecuteQuery(query, []interface{}{startDate, endDate})
}

// GetMeanUnemploymentRate returns the mean unemployment rate for the given date range
func (ctx *DBContext) GetMeanUnemploymentRate(startDate string, endDate string) ([]map[string]interface{}, error) {
	query := `
		select AVG(UNRATE) as MeanUnemploymentRate
		from main.unemployment_rate
		where Date between cast(? as date) and cast(? as date)
	`
	return ctx.ExecuteQuery(query, []interface{}{startDate, endDate})
}
