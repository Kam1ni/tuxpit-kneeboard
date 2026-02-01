package kneeboardview

type colorCombo struct {
	backgroundColor string
	foregroundColor string
}

func createColorCombo(bg string, fg string) colorCombo {
	return colorCombo{backgroundColor: bg, foregroundColor: fg}
}

var colorCombos = []colorCombo{
	createColorCombo("#c04965", "#ffffff"),
	createColorCombo("#458991", "#000000"),
	createColorCombo("#d974b9", "#000000"),
	createColorCombo("#6260cb", "#ffffff"),
	createColorCombo("#e45f6c", "#000000"),
	createColorCombo("#31d768", "#000000"),
	createColorCombo("#41306f", "#ffffff"),
	createColorCombo("#c7889a", "#000000"),
	createColorCombo("#9789c2", "#000000"),
	createColorCombo("#c59589", "#000000"),
	createColorCombo("#8483c1", "#000000"),
	createColorCombo("#6577a5", "#000000"),
	createColorCombo("#b841cd", "#000000"),
	createColorCombo("#96c1ad", "#000000"),
	createColorCombo("#c93fc2", "#000000"),
	createColorCombo("#5caa7c", "#000000"),
	createColorCombo("#5f2aa0", "#ffffff"),
	createColorCombo("#91932d", "#000000"),
	createColorCombo("#864760", "#ffffff"),
	createColorCombo("#7f2885", "#ffffff"),
	createColorCombo("#599411", "#000000"),
	createColorCombo("#3cb036", "#000000"),
	createColorCombo("#64ebb5", "#000000"),
	createColorCombo("#86138e", "#ffffff"),
	createColorCombo("#50eaaf", "#000000"),
	createColorCombo("#ba9784", "#000000"),
	createColorCombo("#5b7c32", "#ffffff"),
	createColorCombo("#481fa2", "#ffffff"),
	createColorCombo("#6ad262", "#000000"),
	createColorCombo("#22a019", "#000000"),
	createColorCombo("#d47d8d", "#000000"),
	createColorCombo("#c17dc4", "#000000"),
	createColorCombo("#a9ad39", "#000000"),
	createColorCombo("#da65ea", "#000000"),
	createColorCombo("#62dc93", "#000000"),
	createColorCombo("#bb4745", "#ffffff"),
	createColorCombo("#b6615e", "#000000"),
	createColorCombo("#3e3ee2", "#ffffff"),
	createColorCombo("#9d39b8", "#ffffff"),
	createColorCombo("#b9829b", "#000000"),
	createColorCombo("#713e4c", "#ffffff"),
	createColorCombo("#72d6ae", "#000000"),
	createColorCombo("#c04cc4", "#000000"),
	createColorCombo("#87703b", "#ffffff"),
	createColorCombo("#b73db1", "#ffffff"),
	createColorCombo("#7f7436", "#ffffff"),
	createColorCombo("#14ac2b", "#000000"),
	createColorCombo("#bcb674", "#000000"),
	createColorCombo("#cb3659", "#ffffff"),
	createColorCombo("#404f62", "#ffffff"),
}
