const ingredientArea = document.getElementById("ingredients");
ingredientArea.addEventListener("blur", (e) => {
	// Match lines that don't start with "## " or "* " and are non-empty
	if (ingredientArea.value.match(/^(?!(?:## |\* ).+|$)/m)) {
		ingredientArea.setCustomValidity("Error: Incorrect formatting");
	}
});
ingredientArea.addEventListener("focus", (e) => {
	ingredientArea.setCustomValidity("");
});
