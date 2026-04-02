package exec

import "org/utils"

func Run() {
	println("Organizing files...")
	defer println("Organized files")

	utils.OrgImagens()
	utils.OrgVideos()
	utils.OrgAudio()
	utils.OrgDocuments()
	utils.OrgSpreadsheets()
	utils.OrgPresentations()
	utils.OrgCompressed()
	utils.OrgExecutables()
	utils.OrgProgramming()
	utils.OrgWeb()
	utils.OrgDataBase()
	utils.OrgGameData()
	utils.OrgCertificates()
	utils.OrgRandom()
}
