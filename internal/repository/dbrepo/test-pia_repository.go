package dbrepo

// GetAllAssets ruft alle PIA-Assets ab
func (r *testPiaRepository) GetAllAssets() string {
	return "Test data fetched using environment: " + r.App.Env.String()
}
