package dbrepo

// GetAllAssets ruft alle PIA-Assets ab
func (r *piaRepositoryImpl) GetAllAssets() string {
	return "Data fetched using environment: " + r.App.Env.String()
}
